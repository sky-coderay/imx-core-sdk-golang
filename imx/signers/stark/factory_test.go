package stark

import (
	"math/big"
	"testing"

	"github.com/immutable/imx-core-sdk-golang/imx/signers/ethereum"

	"github.com/stretchr/testify/assert"
)

func TestStarkSignerFactory_GenerateRandomKey(t *testing.T) {
	key1, err := GenerateKey()
	assert.NoError(t, err)
	key2, err := GenerateKey()
	assert.NoError(t, err)

	assert.NotEqualf(t, key1, key2, "Generated random keys are not same")
}

func TestStarkSignerFactory_Grinding(t *testing.T) {
	type test struct {
		name                 string
		privateKey           string
		wantKeyAfterGrinding string
	}
	tests := []test{
		{
			name:                 "correct ground key",
			privateKey:           "86F3E7293141F20A8BAFF320E8EE4ACCB9D4A4BF2B4D295E8CEE784DB46E0519",
			wantKeyAfterGrinding: "5c8c8683596c732541a59e03007b2d30dbbbb873556fe65b5fb63c16688f941",
		},
		{
			name:                 "private key is above the key limit",
			privateKey:           "a978531943ad2e2a8af34e0e2a7d306dc99516d489be16e4ea2ee74c90a9d88f",
			wantKeyAfterGrinding: "1e8108d99e74b769d6b998a5a41ff2745f0607496f2eed39abfd161837408e7",
		},
		{
			name:                 "private key starts with zero",
			privateKey:           "086F3E7293141F20A8BAFF320E8EE4ACCB9D4A4BF2B4D295E8CEE784DB46E051",
			wantKeyAfterGrinding: "2b2c6db790a95ce05426c3d67247547f1a72d104fd5af24553d42b7557ab082",
		},
	}
	for _, tt := range tests {
		privateKey, ok := new(big.Int).SetString(tt.privateKey, 16)
		assert.True(t, ok)
		wantKeyAfterGrinding, ok := new(big.Int).SetString(tt.wantKeyAfterGrinding, 16)
		assert.True(t, ok)
		assert.Equalf(t, wantKeyAfterGrinding, grind(privateKey), "Verify grinding logic")
	}
}

func TestStarkSignerFactory_GenerateLegacyKey(t *testing.T) {
	l1Signer, err := ethereum.NewSigner("5c7b4b5cad9a3fc7b1ba235a49cd74e615488a18b0d6a531739fd1062935104d", big.NewInt(5))
	assert.NoError(t, err)
	key1, err := GenerateLegacyKey(l1Signer)
	assert.NoError(t, err)
	assert.Equalf(t, "0x556413893a023efd75f62cd4eca825f2be7e918b5188f1db06cbec12d7d1b88", key1, "Check the generated key matches")

	key2, err := GenerateLegacyKey(l1Signer)
	assert.NoError(t, err)

	assert.Equalf(t, key1, key2, "Generated keys are same")
}

func TestStarkSignerFactory_GenerateLegacyKey_GrindkeyFix(t *testing.T) {
	type test struct {
		name          string
		privateKey    string
		wantPublicKey string
	}
	tests := []test{
		{
			name:          "old test",
			privateKey:    "5c7b4b5cad9a3fc7b1ba235a49cd74e615488a18b0d6a531739fd1062935104d",
			wantPublicKey: "0x556413893a023efd75f62cd4eca825f2be7e918b5188f1db06cbec12d7d1b88",
		},
		// {
		// 	name:          "grindKeyfix", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "5c7b4b5cad9a3fc7b1ba235a49cd74e615488a18b0d6a531739fd1062935104d",
		// 	wantPublicKey: "0x0579f97e8084dfbbead9bffd750df780e06d8c09a3ba7f40ebe51d46b47df043",
		// },
		// {
		// 	name:          "grindKeyfix hashed once", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0x1a245f2fa7c4f04a65d45a3877ad00b1423d081490dcc1a7050c8d7c11ec5c8f",
		// 	wantPublicKey: "0x07ca61905954bdd858ae63704b111da5ca52b30a21fa9aaeee9aef9b24e89607",
		// },
		// {
		// 	name:          "grindKeyfix hashed more than once", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0xe516a5b715dc53ba4bf06ed98cbe50897921f45f76d4138f8a7d4ba02a89e10d",
		// 	wantPublicKey: "0x0157c6efcb2c69604c1b46ffeb568b3b10ea7093da762e1daf66b59d9b63b5f9",
		// },
		// {
		// 	name:          "grindKeyfix is backwards compatible", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "ba3c969f4957e6bf24e5cf8a931bdba4f90d27c01bb7dff738e4593142826db7",
		// 	wantPublicKey: "0x04eb684f7318d2b90ef8703e6cd77e362414f4997934ee3c99ee50db3411dfef",
		// },
		// {
		// 	name:          "grindKeyfix hashed once is backwards compatible", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0xa25022c521d884d195c119b705ac515fabb48971d3cf2369f188c96b0e6404ef",
		// 	wantPublicKey: "0x03aa9edf4a5b33f623d3dd37a6df62e5590f172fd4d1a9483532464aa13bcbb9",
		// },
		// {
		// 	name:          "grindKeyfix hashed more than once is backwards compatible", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0x719b531dfdbf5e327646ccd992923c9d8846d60767b49d0589d631c1f54d1f12",
		// 	wantPublicKey: "0x057a8278637277ebaaa63b000185507eb847eec8f69c7758fc10a490d0226e7d",
		// },
		// {
		// 	name:          "grindKeyfix hashed once generates correct legacy key", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0xe7ecb8f91175446248a6cfa45a9526c85bc4cd7cbd9427e3e65a82d3f5fb8cdc",
		// 	wantPublicKey: "0x05f8f00b03e896cb8c02a936bab73623fc3651fddabfe0d28fbc309c63642c10",
		// },
		// {
		// 	name:          "grindKeyfix hashed more than once generates correct legacy key", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0x82930dfa052b198828f70529148b02178b03515e910217dab32bf9fc046bc9d9",
		// 	wantPublicKey: "0x04d34c9f519b2a3538b02c4e939188b591065eb8c2f210277c6ed1de85b16fab",
		// },
		// {
		// 	name:          "random eth private key with no index generates correct legacy key", //from imx-sdk-js, starkCurve.test.ts
		// 	privateKey:    "0x0ac63e2143a49a14a87d865c8d7993806ff16ac1c3288ff97101569881f0d306",
		// 	wantPublicKey: "0x06e6ac4bb44f3295b881532452f90eccb5601314fafe306db17684b47aa388bd",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1Signer, err := ethereum.NewSigner(tt.privateKey, big.NewInt(5))
			assert.NoError(t, err)
			key1, err := GenerateLegacyKey(l1Signer)
			assert.NoError(t, err)
			assert.Equalf(t, tt.wantPublicKey, key1, "Check the generated key matches")
			key2, err := GenerateLegacyKey(l1Signer)
			assert.NoError(t, err)
			assert.Equalf(t, key1, key2, "Generated keys are same")
		})
	}
}
