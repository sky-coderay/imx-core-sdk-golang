package stark

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"

	"github.com/aarbt/hdkeys"
	"github.com/dontpanicdao/caigo"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// package-global variable to represent our standard stark curve
	curve *caigo.StarkCurve

	//go:embed pedersen_params.json
	pedersenParamsBytes []byte

	//go:embed *.json
	_ embed.FS // Unused but required to import the embed module
)

// GenerateKey generates a random key that can be used to create StarkSigner.
// On creation save this key for future usage as this key will be required to reuse your stark signer.
// @return Randomly generated private key.
func GenerateKey() (string, error) {
	if curve == nil {
		var err error
		curve, err = loadCurve()
		if err != nil {
			return "", err
		}
	}
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return "", err
	}
	return hexutil.EncodeBig(grindKeyV100Beta1(privKey.D)), nil
}

const (
	DefaultSeedMessage = "Only sign this request if you’ve initiated an action with Immutable X."
	ApplicationName    = "immutablex"
	LayerName          = "starkex"
	Index              = "1"
)

// GenerateLegacyKey Generates a deterministic Stark private key from the provided signer.
// @return Deterministically generated private key.
func GenerateLegacyKey(signer imx.L1Signer) (string, error) {
	seed, err := generateSeed(signer, DefaultSeedMessage)
	if err != nil {
		return "", fmt.Errorf("failed to generate seed using l1signer: %v", err)
	}

	starkPath := getStarkPath(LayerName, ApplicationName, signer.GetAddress(), Index)
	childKey, err := hdkeys.NewMasterKey(seed).Chain(starkPath)
	if err != nil {
		return "", err
	}

	// Last 32 bits
	childPrivateKey := childKey.Serialize()[46:]
	keyBigInt := new(big.Int).SetBytes(childPrivateKey)
	starkPrivateKey := hexutil.EncodeBig(grindKey(keyBigInt))

	// The following logic is added in response to bug found in DX-2167 and DX-2184.
	// To provide a backwards compatible way to fetch keys across link/js SDK and Core SDK.

	// Note: The issue we are addressing here is that if the hashKeyWithIndex value is above the limit, then we perform
	//     hash again with index until it comes below the limit. But for the first time the index is not incremented in
	//     imx-sdk-js where as it was in core-sdk that was the difference. For this reason it would have worked for cases
	//     where the hashed value was less than the limit and fails only when it is above the limit.
	//     Refer, https://immutable.atlassian.net/browse/DX-2167

	// Same code as grindKey, required to check if the generated private key
	// goes above the limit when hashed for first time to identify if we need to do backwards
	// compatibility check:

	// The bug only exists if the hashed value of given seed is above the stark curve limit.
	if !checkIfHashedKeyIsAboveLimit(keyBigInt) {
		return starkPrivateKey, nil
	}

	// Check if the generated stark public key matches with the existing account value for that user.
	// We are only validating for Production environment.
	// For Sandbox account/key mismatch, solution is to discard the old account and create a new one.
	registeredStarkPublicKey, accountNotFound, err := getStarkPublicKeyFromImx(context.Background(), signer.GetAddress())
	if err != nil {
		return "", fmt.Errorf("error in obtaining the registered public key: %w", err)
	}

	// If the account is not found or account matches we just return the key pair at the end of this method.
	// Only need to so alternative method if the account is found but the stark public key does not match.

	// If the account is not found it is a new account, just return the Stark Private Key that is generated by grindKey function.
	if accountNotFound {
		return starkPrivateKey, nil
	}

	starkSigner, err := NewSigner(starkPrivateKey)
	if err != nil {
		return "", fmt.Errorf("creating stark signer: %w", err)
	}
	// If the user account matches with generated stark public key user, just return Stark Private Key.
	if registeredStarkPublicKey == starkSigner.GetPublicKey() {
		return starkPrivateKey, nil
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// This is backwards compatible crypto (core-sdk) version 2.0.1

	// If we are here, we found the account but did not match with the recorded user account.
	// Lets try to use grindKeyV100Beta1 method from backwards compatible logic to generate a key and see if that matches.
	starkPrivateKeyV100Beta1Compatible := hexutil.EncodeBig(grindKeyV100Beta1(keyBigInt))
	starkSigner, err = NewSigner(starkPrivateKeyV100Beta1Compatible)
	if err != nil {
		return "", fmt.Errorf("creating stark signer: %w", err)
	}
	if registeredStarkPublicKey == starkSigner.GetPublicKey() {
		return starkPrivateKeyV100Beta1Compatible, nil
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// This section is legacy crypto, compatible with imx-sdk-js released before 1.43.5

	starkPrivateKeyLegacy := hexutil.EncodeBig(grindKeyLegacy(keyBigInt))
	starkSigner, err = NewSigner(starkPrivateKeyLegacy)
	if err != nil {
		return "", fmt.Errorf("creating stark signer: %w", err)
	}
	if registeredStarkPublicKey == starkSigner.GetPublicKey() {
		return starkPrivateKeyLegacy, nil
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	// Account is found, but did not match with stark public keys generated by either grindKey or grindKeyV1 method.
	// Will have to contact support for further investigation.
	return "", errors.New("can not deterministically generate stark private key - please contact support")
}

// Initialises the Stark Elliptic Curve.
func loadCurve() (*caigo.StarkCurve, error) {
	caigo.PedersenParamsRaw = pedersenParamsBytes
	return &caigo.Curve, nil
}

// Create a hash from a key + an index
func hashKeyWithIndex(key *big.Int, index byte) *big.Int {
	bytes := sha256.Sum256(append(key.Bytes(), index))
	return new(big.Int).SetBytes(bytes[:])
}

/*
grindKeyV100Beta1 receives a key seed and produces an appropriate StarkEx key from a uniform distribution.

Although it is possible to define a StarkEx key as a residue (mod) between the StarkEx EC order and a
random 256bit digest value, the result would be a biased key. In order to prevent this bias, we
deterministically search (by applying more hashes, AKA grinding) for a value lower than the largest
256bit multiple of StarkEx EC order.

https://github.com/starkware-libs/starkware-crypto-utils/blob/dev/src/js/key_derivation.js#L119
*/
func grindKeyV100Beta1(key *big.Int) *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))

	//index is 0, 0, 1, 2...
	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		key = hashKeyWithIndex(key, i)
		i += 1
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

// grindKeyLegacy is compatible with legacy imx-sdk-js versions 1.45.6 and prior.
func grindKeyLegacy(key *big.Int) *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))

	//index is 0, 0, 0, ...
	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		key = hashKeyWithIndex(key, i)
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

// grindKey is the default and correct one to use.
func grindKey(key *big.Int) *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))

	//index is 0, 1, 2, ...
	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		i += 1
		key = hashKeyWithIndex(key, i)
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

func maxAllowedValue() *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)
	return new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))
}

// checkIfHashedKeyIsAboveLimit checks if the hash value of the the given PrivateKey falls above the starkEcOrder limit.
// This function is only serving the context of DX-2184, used to determine if we need to validate the generated key
// against the one recorded in IMX servers.
func checkIfHashedKeyIsAboveLimit(key *big.Int) bool {
	maxAllowedVal := maxAllowedValue()
	// The key passed to hashKeyWithIndex must have a length of 64 characters
	// to ensure that the correct number of leading zeroes are used as input
	// to the hashing loop
	hashedKey := hashKeyWithIndex(key, 0)
	return (hashedKey.Cmp(maxAllowedVal) >= 0)
}

// generateSeed generates the seed value for the given seed message.
func generateSeed(signer imx.L1Signer, seedMessage string) ([]byte, error) {
	signature, err := signer.SignMessage(seedMessage)
	if err != nil {
		return nil, err
	}
	return signature[32:64], nil
}

func getStarkPath(layerName, applicationName, ethereumAddress, index string) string {
	// Starkware keys are derived with the following BIP43-compatible derivation path, with direct inspiration from BIP44:
	//
	// m / purpose' / layer' / application' / eth_address_1' / eth_address_2' / index
	// where:
	//
	// m 			- the seed.
	// purpose 		- 2645 (the number of this EIP).
	// layer 		- the 31 lowest bits of sha256 on the layer name. Serve as a domain separator between different technologies. In the context of starkex, the value would be 579218131.
	// application 	- the 31 lowest bits of sha256 of the application name. Serve as a domain separator between different applications.
	// 					In the context of DeversiFi in June 2020, it is the 31 lowest bits of sha256(starkexdvf) and the value would be 1393043894.
	// eth_address_1 / eth_address_2 - the first and second 31 lowest bits of the corresponding eth_address.
	// index 		- to allow multiple keys per eth_address.
	//
	// See for more info regarding path derivation https://docs.starkware.co/starkex/key-derivation.html and https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2645.md

	layerHash := sha256.Sum256([]byte(layerName))
	appHash := sha256.Sum256([]byte(applicationName))
	layer := binary.BigEndian.Uint32(layerHash[28:]) & (1<<31 - 1)
	application := binary.BigEndian.Uint32(appHash[28:]) & (1<<31 - 1)

	ethereumAddressInBytes := hexutil.MustDecode(ethereumAddress)
	ethAddress1 := binary.BigEndian.Uint32(ethereumAddressInBytes[16:]) & (1<<31 - 1)
	// Mask of 31 binary 1 digits, at the position 32 from the end (counting from 1)
	ethAddress2 := (binary.BigEndian.Uint64(ethereumAddressInBytes[12:]) & ((1<<31 - 1) << 31)) >> 31

	return fmt.Sprintf("m/2645'/%d'/%d'/%d'/%d'/%s", layer, application, ethAddress1, ethAddress2, index)
}

// getStarkPublicKeyFromImx gets the account (stark public key) value of the requested user (ethAddress) for Production environment only.
func getStarkPublicKeyFromImx(ctx context.Context, ethAddress string) (starkPublicKey string, accountNotFound bool, err error) {
	apiConfig := api.NewConfiguration()
	apiConfig.Servers[0] = api.ServerConfiguration{
		URL:         imx.Mainnet.BaseAPIPath,
		Description: "Prod Environment",
	}
	usersApi := api.NewAPIClient(apiConfig).UsersApi

	// Query existing account value for the given user (ethAddress).
	response, httpResponse, err := usersApi.GetUsers(ctx, ethAddress).Execute()
	defer httpResponse.Body.Close()
	if err != nil {
		if httpResponse.StatusCode == http.StatusNotFound {
			return "", true, nil // This means a new account. So lets use the value from default GrindKey function.
		}
		return "", false, imx.NewIMXError(httpResponse, err)
	}
	//err is nil from here on
	if response == nil {
		return "", false, imx.NewIMXError(httpResponse, errors.New("nil response"))
	}
	accounts := response.GetAccounts()
	if len(accounts) == 0 {
		return "", false, imx.NewIMXError(httpResponse, errors.New("successful API call, but 0 accounts returned"))
	}
	return accounts[0], false, nil
}
