/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.12
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateWorkspaceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWorkspaceBody{}

// UpdateWorkspaceBody struct for UpdateWorkspaceBody
type UpdateWorkspaceBody struct {
	// The name of the workspace.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _UpdateWorkspaceBody UpdateWorkspaceBody

// NewUpdateWorkspaceBody instantiates a new UpdateWorkspaceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkspaceBody(name string) *UpdateWorkspaceBody {
	this := UpdateWorkspaceBody{}
	this.Name = name
	return &this
}

// NewUpdateWorkspaceBodyWithDefaults instantiates a new UpdateWorkspaceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkspaceBodyWithDefaults() *UpdateWorkspaceBody {
	this := UpdateWorkspaceBody{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateWorkspaceBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateWorkspaceBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateWorkspaceBody) SetName(v string) {
	o.Name = v
}

func (o UpdateWorkspaceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWorkspaceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateWorkspaceBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varUpdateWorkspaceBody := _UpdateWorkspaceBody{}

	err = json.Unmarshal(data, &varUpdateWorkspaceBody)

	if err != nil {
		return err
	}

	*o = UpdateWorkspaceBody(varUpdateWorkspaceBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateWorkspaceBody struct {
	value *UpdateWorkspaceBody
	isSet bool
}

func (v NullableUpdateWorkspaceBody) Get() *UpdateWorkspaceBody {
	return v.value
}

func (v *NullableUpdateWorkspaceBody) Set(val *UpdateWorkspaceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkspaceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkspaceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkspaceBody(val *UpdateWorkspaceBody) *NullableUpdateWorkspaceBody {
	return &NullableUpdateWorkspaceBody{value: val, isSet: true}
}

func (v NullableUpdateWorkspaceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkspaceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

