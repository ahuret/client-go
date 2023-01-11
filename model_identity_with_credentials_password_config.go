/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.5
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentityWithCredentialsPasswordConfig Create Identity and Import Password Credentials Configuration
type IdentityWithCredentialsPasswordConfig struct {
	// The hashed password in [PHC format]( https://www.ory.sh/docs/kratos/concepts/credentials/username-email-password#hashed-password-format)
	HashedPassword *string `json:"hashed_password,omitempty"`
	// The password in plain text if no hash is available.
	Password *string `json:"password,omitempty"`
}

// NewIdentityWithCredentialsPasswordConfig instantiates a new IdentityWithCredentialsPasswordConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsPasswordConfig() *IdentityWithCredentialsPasswordConfig {
	this := IdentityWithCredentialsPasswordConfig{}
	return &this
}

// NewIdentityWithCredentialsPasswordConfigWithDefaults instantiates a new IdentityWithCredentialsPasswordConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsPasswordConfigWithDefaults() *IdentityWithCredentialsPasswordConfig {
	this := IdentityWithCredentialsPasswordConfig{}
	return &this
}

// GetHashedPassword returns the HashedPassword field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPasswordConfig) GetHashedPassword() string {
	if o == nil || o.HashedPassword == nil {
		var ret string
		return ret
	}
	return *o.HashedPassword
}

// GetHashedPasswordOk returns a tuple with the HashedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPasswordConfig) GetHashedPasswordOk() (*string, bool) {
	if o == nil || o.HashedPassword == nil {
		return nil, false
	}
	return o.HashedPassword, true
}

// HasHashedPassword returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPasswordConfig) HasHashedPassword() bool {
	if o != nil && o.HashedPassword != nil {
		return true
	}

	return false
}

// SetHashedPassword gets a reference to the given string and assigns it to the HashedPassword field.
func (o *IdentityWithCredentialsPasswordConfig) SetHashedPassword(v string) {
	o.HashedPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPasswordConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPasswordConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPasswordConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IdentityWithCredentialsPasswordConfig) SetPassword(v string) {
	o.Password = &v
}

func (o IdentityWithCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HashedPassword != nil {
		toSerialize["hashed_password"] = o.HashedPassword
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityWithCredentialsPasswordConfig struct {
	value *IdentityWithCredentialsPasswordConfig
	isSet bool
}

func (v NullableIdentityWithCredentialsPasswordConfig) Get() *IdentityWithCredentialsPasswordConfig {
	return v.value
}

func (v *NullableIdentityWithCredentialsPasswordConfig) Set(val *IdentityWithCredentialsPasswordConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsPasswordConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsPasswordConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsPasswordConfig(val *IdentityWithCredentialsPasswordConfig) *NullableIdentityWithCredentialsPasswordConfig {
	return &NullableIdentityWithCredentialsPasswordConfig{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsPasswordConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


