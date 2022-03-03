/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.112
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateProjectConfigConfig struct for UpdateProjectConfigConfig
type UpdateProjectConfigConfig struct {
	// The Ory Kratos config to import
	Identity map[string]interface{} `json:"identity"`
}

// NewUpdateProjectConfigConfig instantiates a new UpdateProjectConfigConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectConfigConfig(identity map[string]interface{}) *UpdateProjectConfigConfig {
	this := UpdateProjectConfigConfig{}
	this.Identity = identity
	return &this
}

// NewUpdateProjectConfigConfigWithDefaults instantiates a new UpdateProjectConfigConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectConfigConfigWithDefaults() *UpdateProjectConfigConfig {
	this := UpdateProjectConfigConfig{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *UpdateProjectConfigConfig) GetIdentity() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectConfigConfig) GetIdentityOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identity, true
}

// SetIdentity sets field value
func (o *UpdateProjectConfigConfig) SetIdentity(v map[string]interface{}) {
	o.Identity = v
}

func (o UpdateProjectConfigConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProjectConfigConfig struct {
	value *UpdateProjectConfigConfig
	isSet bool
}

func (v NullableUpdateProjectConfigConfig) Get() *UpdateProjectConfigConfig {
	return v.value
}

func (v *NullableUpdateProjectConfigConfig) Set(val *UpdateProjectConfigConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectConfigConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectConfigConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectConfigConfig(val *UpdateProjectConfigConfig) *NullableUpdateProjectConfigConfig {
	return &NullableUpdateProjectConfigConfig{value: val, isSet: true}
}

func (v NullableUpdateProjectConfigConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectConfigConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


