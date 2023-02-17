/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.16
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentityWithCredentialsOidc Create Identity and Import Social Sign In Credentials
type IdentityWithCredentialsOidc struct {
	Config *IdentityWithCredentialsOidcConfig `json:"config,omitempty"`
}

// NewIdentityWithCredentialsOidc instantiates a new IdentityWithCredentialsOidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsOidc() *IdentityWithCredentialsOidc {
	this := IdentityWithCredentialsOidc{}
	return &this
}

// NewIdentityWithCredentialsOidcWithDefaults instantiates a new IdentityWithCredentialsOidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsOidcWithDefaults() *IdentityWithCredentialsOidc {
	this := IdentityWithCredentialsOidc{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentityWithCredentialsOidc) GetConfig() IdentityWithCredentialsOidcConfig {
	if o == nil || o.Config == nil {
		var ret IdentityWithCredentialsOidcConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsOidc) GetConfigOk() (*IdentityWithCredentialsOidcConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentityWithCredentialsOidc) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentityWithCredentialsOidcConfig and assigns it to the Config field.
func (o *IdentityWithCredentialsOidc) SetConfig(v IdentityWithCredentialsOidcConfig) {
	o.Config = &v
}

func (o IdentityWithCredentialsOidc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityWithCredentialsOidc struct {
	value *IdentityWithCredentialsOidc
	isSet bool
}

func (v NullableIdentityWithCredentialsOidc) Get() *IdentityWithCredentialsOidc {
	return v.value
}

func (v *NullableIdentityWithCredentialsOidc) Set(val *IdentityWithCredentialsOidc) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsOidc(val *IdentityWithCredentialsOidc) *NullableIdentityWithCredentialsOidc {
	return &NullableIdentityWithCredentialsOidc{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

