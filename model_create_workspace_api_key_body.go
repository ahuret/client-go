/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.14.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateWorkspaceApiKeyBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceApiKeyBody{}

// CreateWorkspaceApiKeyBody struct for CreateWorkspaceApiKeyBody
type CreateWorkspaceApiKeyBody struct {
	// The API Key Name  A descriptive name for the API key.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreateWorkspaceApiKeyBody CreateWorkspaceApiKeyBody

// NewCreateWorkspaceApiKeyBody instantiates a new CreateWorkspaceApiKeyBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceApiKeyBody(name string) *CreateWorkspaceApiKeyBody {
	this := CreateWorkspaceApiKeyBody{}
	this.Name = name
	return &this
}

// NewCreateWorkspaceApiKeyBodyWithDefaults instantiates a new CreateWorkspaceApiKeyBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceApiKeyBodyWithDefaults() *CreateWorkspaceApiKeyBody {
	this := CreateWorkspaceApiKeyBody{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWorkspaceApiKeyBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceApiKeyBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWorkspaceApiKeyBody) SetName(v string) {
	o.Name = v
}

func (o CreateWorkspaceApiKeyBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceApiKeyBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateWorkspaceApiKeyBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateWorkspaceApiKeyBody := _CreateWorkspaceApiKeyBody{}

	err = json.Unmarshal(data, &varCreateWorkspaceApiKeyBody)

	if err != nil {
		return err
	}

	*o = CreateWorkspaceApiKeyBody(varCreateWorkspaceApiKeyBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWorkspaceApiKeyBody struct {
	value *CreateWorkspaceApiKeyBody
	isSet bool
}

func (v NullableCreateWorkspaceApiKeyBody) Get() *CreateWorkspaceApiKeyBody {
	return v.value
}

func (v *NullableCreateWorkspaceApiKeyBody) Set(val *CreateWorkspaceApiKeyBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceApiKeyBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceApiKeyBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceApiKeyBody(val *CreateWorkspaceApiKeyBody) *NullableCreateWorkspaceApiKeyBody {
	return &NullableCreateWorkspaceApiKeyBody{value: val, isSet: true}
}

func (v NullableCreateWorkspaceApiKeyBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceApiKeyBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


