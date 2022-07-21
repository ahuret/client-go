/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.1.0-alpha.12
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectServiceIdentity struct for ProjectServiceIdentity
type ProjectServiceIdentity struct {
	Config map[string]interface{} `json:"config"`
}

// NewProjectServiceIdentity instantiates a new ProjectServiceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServiceIdentity(config map[string]interface{}) *ProjectServiceIdentity {
	this := ProjectServiceIdentity{}
	this.Config = config
	return &this
}

// NewProjectServiceIdentityWithDefaults instantiates a new ProjectServiceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServiceIdentityWithDefaults() *ProjectServiceIdentity {
	this := ProjectServiceIdentity{}
	return &this
}

// GetConfig returns the Config field value
func (o *ProjectServiceIdentity) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceIdentity) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *ProjectServiceIdentity) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o ProjectServiceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableProjectServiceIdentity struct {
	value *ProjectServiceIdentity
	isSet bool
}

func (v NullableProjectServiceIdentity) Get() *ProjectServiceIdentity {
	return v.value
}

func (v *NullableProjectServiceIdentity) Set(val *ProjectServiceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServiceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServiceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServiceIdentity(val *ProjectServiceIdentity) *NullableProjectServiceIdentity {
	return &NullableProjectServiceIdentity{value: val, isSet: true}
}

func (v NullableProjectServiceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServiceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


