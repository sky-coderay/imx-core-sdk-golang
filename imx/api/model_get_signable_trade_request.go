/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSignableTradeRequest struct for GetSignableTradeRequest
type GetSignableTradeRequest struct {
	// ExpirationTimestamp in Unix time. Note: will be rounded down to the nearest hour
	ExpirationTimestamp *int32 `json:"expiration_timestamp,omitempty"`
	// Inclusion of either maker or taker fees
	Fees []FeeEntry `json:"fees,omitempty"`
	// The ID of the maker order involved
	OrderId int32 `json:"order_id"`
	// Ethereum address of the submitting user
	User string `json:"user"`
}

// NewGetSignableTradeRequest instantiates a new GetSignableTradeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTradeRequest(orderId int32, user string) *GetSignableTradeRequest {
	this := GetSignableTradeRequest{}
	this.OrderId = orderId
	this.User = user
	return &this
}

// NewGetSignableTradeRequestWithDefaults instantiates a new GetSignableTradeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTradeRequestWithDefaults() *GetSignableTradeRequest {
	this := GetSignableTradeRequest{}
	return &this
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetSignableTradeRequest) GetExpirationTimestamp() int32 {
	if o == nil || o.ExpirationTimestamp == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableTradeRequest) GetExpirationTimestampOk() (*int32, bool) {
	if o == nil || o.ExpirationTimestamp == nil {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetSignableTradeRequest) HasExpirationTimestamp() bool {
	if o != nil && o.ExpirationTimestamp != nil {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int32 and assigns it to the ExpirationTimestamp field.
func (o *GetSignableTradeRequest) SetExpirationTimestamp(v int32) {
	o.ExpirationTimestamp = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *GetSignableTradeRequest) GetFees() []FeeEntry {
	if o == nil || o.Fees == nil {
		var ret []FeeEntry
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableTradeRequest) GetFeesOk() ([]FeeEntry, bool) {
	if o == nil || o.Fees == nil {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *GetSignableTradeRequest) HasFees() bool {
	if o != nil && o.Fees != nil {
		return true
	}

	return false
}

// SetFees gets a reference to the given []FeeEntry and assigns it to the Fees field.
func (o *GetSignableTradeRequest) SetFees(v []FeeEntry) {
	o.Fees = v
}

// GetOrderId returns the OrderId field value
func (o *GetSignableTradeRequest) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeRequest) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *GetSignableTradeRequest) SetOrderId(v int32) {
	o.OrderId = v
}

// GetUser returns the User field value
func (o *GetSignableTradeRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetSignableTradeRequest) SetUser(v string) {
	o.User = v
}

func (o GetSignableTradeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationTimestamp != nil {
		toSerialize["expiration_timestamp"] = o.ExpirationTimestamp
	}
	if o.Fees != nil {
		toSerialize["fees"] = o.Fees
	}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableTradeRequest struct {
	value *GetSignableTradeRequest
	isSet bool
}

func (v NullableGetSignableTradeRequest) Get() *GetSignableTradeRequest {
	return v.value
}

func (v *NullableGetSignableTradeRequest) Set(val *GetSignableTradeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTradeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTradeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTradeRequest(val *GetSignableTradeRequest) *NullableGetSignableTradeRequest {
	return &NullableGetSignableTradeRequest{value: val, isSet: true}
}

func (v NullableGetSignableTradeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTradeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

