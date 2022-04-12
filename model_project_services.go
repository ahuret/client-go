/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.164
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectServices struct for ProjectServices
type ProjectServices struct {
	Identity ProjectServiceIdentity `json:"identity"`
}

// NewProjectServices instantiates a new ProjectServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServices(identity ProjectServiceIdentity) *ProjectServices {
	this := ProjectServices{}
	this.Identity = identity
	return &this
}

// NewProjectServicesWithDefaults instantiates a new ProjectServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServicesWithDefaults() *ProjectServices {
	this := ProjectServices{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *ProjectServices) GetIdentity() ProjectServiceIdentity {
	if o == nil {
		var ret ProjectServiceIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *ProjectServices) GetIdentityOk() (*ProjectServiceIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *ProjectServices) SetIdentity(v ProjectServiceIdentity) {
	o.Identity = v
}

func (o ProjectServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	return json.Marshal(toSerialize)
}

type NullableProjectServices struct {
	value *ProjectServices
	isSet bool
}

func (v NullableProjectServices) Get() *ProjectServices {
	return v.value
}

func (v *NullableProjectServices) Set(val *ProjectServices) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServices) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServices(val *ProjectServices) *NullableProjectServices {
	return &NullableProjectServices{value: val, isSet: true}
}

func (v NullableProjectServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


