/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectServicePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectServicePermission{}

// ProjectServicePermission struct for ProjectServicePermission
type ProjectServicePermission struct {
	Config map[string]interface{} `json:"config"`
	AdditionalProperties map[string]interface{}
}

type _ProjectServicePermission ProjectServicePermission

// NewProjectServicePermission instantiates a new ProjectServicePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServicePermission(config map[string]interface{}) *ProjectServicePermission {
	this := ProjectServicePermission{}
	this.Config = config
	return &this
}

// NewProjectServicePermissionWithDefaults instantiates a new ProjectServicePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServicePermissionWithDefaults() *ProjectServicePermission {
	this := ProjectServicePermission{}
	return &this
}

// GetConfig returns the Config field value
func (o *ProjectServicePermission) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ProjectServicePermission) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *ProjectServicePermission) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o ProjectServicePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectServicePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectServicePermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectServicePermission := _ProjectServicePermission{}

	err = json.Unmarshal(data, &varProjectServicePermission)

	if err != nil {
		return err
	}

	*o = ProjectServicePermission(varProjectServicePermission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectServicePermission struct {
	value *ProjectServicePermission
	isSet bool
}

func (v NullableProjectServicePermission) Get() *ProjectServicePermission {
	return v.value
}

func (v *NullableProjectServicePermission) Set(val *ProjectServicePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServicePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServicePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServicePermission(val *ProjectServicePermission) *NullableProjectServicePermission {
	return &NullableProjectServicePermission{value: val, isSet: true}
}

func (v NullableProjectServicePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServicePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


