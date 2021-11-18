/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.23
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectPasswordConfig struct for ProjectPasswordConfig
type ProjectPasswordConfig struct {
	// Set to true to enable the password authentication method.
	Enabled *bool `json:"enabled,omitempty"`
	// Set to true to remove active sessions when the users logs in.
	RevokeActiveSessions *bool `json:"revoke_active_sessions,omitempty"`
}

// NewProjectPasswordConfig instantiates a new ProjectPasswordConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectPasswordConfig() *ProjectPasswordConfig {
	this := ProjectPasswordConfig{}
	return &this
}

// NewProjectPasswordConfigWithDefaults instantiates a new ProjectPasswordConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectPasswordConfigWithDefaults() *ProjectPasswordConfig {
	this := ProjectPasswordConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProjectPasswordConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPasswordConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProjectPasswordConfig) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProjectPasswordConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRevokeActiveSessions returns the RevokeActiveSessions field value if set, zero value otherwise.
func (o *ProjectPasswordConfig) GetRevokeActiveSessions() bool {
	if o == nil || o.RevokeActiveSessions == nil {
		var ret bool
		return ret
	}
	return *o.RevokeActiveSessions
}

// GetRevokeActiveSessionsOk returns a tuple with the RevokeActiveSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectPasswordConfig) GetRevokeActiveSessionsOk() (*bool, bool) {
	if o == nil || o.RevokeActiveSessions == nil {
		return nil, false
	}
	return o.RevokeActiveSessions, true
}

// HasRevokeActiveSessions returns a boolean if a field has been set.
func (o *ProjectPasswordConfig) HasRevokeActiveSessions() bool {
	if o != nil && o.RevokeActiveSessions != nil {
		return true
	}

	return false
}

// SetRevokeActiveSessions gets a reference to the given bool and assigns it to the RevokeActiveSessions field.
func (o *ProjectPasswordConfig) SetRevokeActiveSessions(v bool) {
	o.RevokeActiveSessions = &v
}

func (o ProjectPasswordConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.RevokeActiveSessions != nil {
		toSerialize["revoke_active_sessions"] = o.RevokeActiveSessions
	}
	return json.Marshal(toSerialize)
}

type NullableProjectPasswordConfig struct {
	value *ProjectPasswordConfig
	isSet bool
}

func (v NullableProjectPasswordConfig) Get() *ProjectPasswordConfig {
	return v.value
}

func (v *NullableProjectPasswordConfig) Set(val *ProjectPasswordConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectPasswordConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectPasswordConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectPasswordConfig(val *ProjectPasswordConfig) *NullableProjectPasswordConfig {
	return &NullableProjectPasswordConfig{value: val, isSet: true}
}

func (v NullableProjectPasswordConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectPasswordConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


