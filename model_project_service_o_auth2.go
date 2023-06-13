/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.36
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectServiceOAuth2 struct for ProjectServiceOAuth2
type ProjectServiceOAuth2 struct {
	Config map[string]interface{} `json:"config"`
	AdditionalProperties map[string]interface{}
}

type _ProjectServiceOAuth2 ProjectServiceOAuth2

// NewProjectServiceOAuth2 instantiates a new ProjectServiceOAuth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServiceOAuth2(config map[string]interface{}) *ProjectServiceOAuth2 {
	this := ProjectServiceOAuth2{}
	this.Config = config
	return &this
}

// NewProjectServiceOAuth2WithDefaults instantiates a new ProjectServiceOAuth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServiceOAuth2WithDefaults() *ProjectServiceOAuth2 {
	this := ProjectServiceOAuth2{}
	return &this
}

// GetConfig returns the Config field value
func (o *ProjectServiceOAuth2) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ProjectServiceOAuth2) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *ProjectServiceOAuth2) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o ProjectServiceOAuth2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProjectServiceOAuth2) UnmarshalJSON(bytes []byte) (err error) {
	varProjectServiceOAuth2 := _ProjectServiceOAuth2{}

	if err = json.Unmarshal(bytes, &varProjectServiceOAuth2); err == nil {
		*o = ProjectServiceOAuth2(varProjectServiceOAuth2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectServiceOAuth2 struct {
	value *ProjectServiceOAuth2
	isSet bool
}

func (v NullableProjectServiceOAuth2) Get() *ProjectServiceOAuth2 {
	return v.value
}

func (v *NullableProjectServiceOAuth2) Set(val *ProjectServiceOAuth2) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServiceOAuth2) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServiceOAuth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServiceOAuth2(val *ProjectServiceOAuth2) *NullableProjectServiceOAuth2 {
	return &NullableProjectServiceOAuth2{value: val, isSet: true}
}

func (v NullableProjectServiceOAuth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServiceOAuth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


