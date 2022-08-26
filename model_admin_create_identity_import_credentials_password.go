/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.21
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdminCreateIdentityImportCredentialsPassword struct for AdminCreateIdentityImportCredentialsPassword
type AdminCreateIdentityImportCredentialsPassword struct {
	Config *AdminCreateIdentityImportCredentialsPasswordConfig `json:"config,omitempty"`
}

// NewAdminCreateIdentityImportCredentialsPassword instantiates a new AdminCreateIdentityImportCredentialsPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateIdentityImportCredentialsPassword() *AdminCreateIdentityImportCredentialsPassword {
	this := AdminCreateIdentityImportCredentialsPassword{}
	return &this
}

// NewAdminCreateIdentityImportCredentialsPasswordWithDefaults instantiates a new AdminCreateIdentityImportCredentialsPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateIdentityImportCredentialsPasswordWithDefaults() *AdminCreateIdentityImportCredentialsPassword {
	this := AdminCreateIdentityImportCredentialsPassword{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AdminCreateIdentityImportCredentialsPassword) GetConfig() AdminCreateIdentityImportCredentialsPasswordConfig {
	if o == nil || o.Config == nil {
		var ret AdminCreateIdentityImportCredentialsPasswordConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityImportCredentialsPassword) GetConfigOk() (*AdminCreateIdentityImportCredentialsPasswordConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AdminCreateIdentityImportCredentialsPassword) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AdminCreateIdentityImportCredentialsPasswordConfig and assigns it to the Config field.
func (o *AdminCreateIdentityImportCredentialsPassword) SetConfig(v AdminCreateIdentityImportCredentialsPasswordConfig) {
	o.Config = &v
}

func (o AdminCreateIdentityImportCredentialsPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableAdminCreateIdentityImportCredentialsPassword struct {
	value *AdminCreateIdentityImportCredentialsPassword
	isSet bool
}

func (v NullableAdminCreateIdentityImportCredentialsPassword) Get() *AdminCreateIdentityImportCredentialsPassword {
	return v.value
}

func (v *NullableAdminCreateIdentityImportCredentialsPassword) Set(val *AdminCreateIdentityImportCredentialsPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateIdentityImportCredentialsPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateIdentityImportCredentialsPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateIdentityImportCredentialsPassword(val *AdminCreateIdentityImportCredentialsPassword) *NullableAdminCreateIdentityImportCredentialsPassword {
	return &NullableAdminCreateIdentityImportCredentialsPassword{value: val, isSet: true}
}

func (v NullableAdminCreateIdentityImportCredentialsPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateIdentityImportCredentialsPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


