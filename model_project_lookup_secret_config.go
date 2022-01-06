/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.41
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectLookupSecretConfig struct for ProjectLookupSecretConfig
type ProjectLookupSecretConfig struct {
	// Set to true to enable the WebAuthn authentication method.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewProjectLookupSecretConfig instantiates a new ProjectLookupSecretConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLookupSecretConfig() *ProjectLookupSecretConfig {
	this := ProjectLookupSecretConfig{}
	return &this
}

// NewProjectLookupSecretConfigWithDefaults instantiates a new ProjectLookupSecretConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLookupSecretConfigWithDefaults() *ProjectLookupSecretConfig {
	this := ProjectLookupSecretConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProjectLookupSecretConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLookupSecretConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProjectLookupSecretConfig) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProjectLookupSecretConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ProjectLookupSecretConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableProjectLookupSecretConfig struct {
	value *ProjectLookupSecretConfig
	isSet bool
}

func (v NullableProjectLookupSecretConfig) Get() *ProjectLookupSecretConfig {
	return v.value
}

func (v *NullableProjectLookupSecretConfig) Set(val *ProjectLookupSecretConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLookupSecretConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLookupSecretConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLookupSecretConfig(val *ProjectLookupSecretConfig) *NullableProjectLookupSecretConfig {
	return &NullableProjectLookupSecretConfig{value: val, isSet: true}
}

func (v NullableProjectLookupSecretConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLookupSecretConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


