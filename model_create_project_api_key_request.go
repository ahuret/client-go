/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.3.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateProjectApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectApiKeyRequest{}

// CreateProjectApiKeyRequest struct for CreateProjectApiKeyRequest
type CreateProjectApiKeyRequest struct {
	// The Token Name  A descriptive name for the token.  in: body
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreateProjectApiKeyRequest CreateProjectApiKeyRequest

// NewCreateProjectApiKeyRequest instantiates a new CreateProjectApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectApiKeyRequest(name string) *CreateProjectApiKeyRequest {
	this := CreateProjectApiKeyRequest{}
	this.Name = name
	return &this
}

// NewCreateProjectApiKeyRequestWithDefaults instantiates a new CreateProjectApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectApiKeyRequestWithDefaults() *CreateProjectApiKeyRequest {
	this := CreateProjectApiKeyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateProjectApiKeyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectApiKeyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectApiKeyRequest) SetName(v string) {
	o.Name = v
}

func (o CreateProjectApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProjectApiKeyRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateProjectApiKeyRequest := _CreateProjectApiKeyRequest{}

	err = json.Unmarshal(bytes, &varCreateProjectApiKeyRequest)

	if err != nil {
		return err
	}

	*o = CreateProjectApiKeyRequest(varCreateProjectApiKeyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProjectApiKeyRequest struct {
	value *CreateProjectApiKeyRequest
	isSet bool
}

func (v NullableCreateProjectApiKeyRequest) Get() *CreateProjectApiKeyRequest {
	return v.value
}

func (v *NullableCreateProjectApiKeyRequest) Set(val *CreateProjectApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectApiKeyRequest(val *CreateProjectApiKeyRequest) *NullableCreateProjectApiKeyRequest {
	return &NullableCreateProjectApiKeyRequest{value: val, isSet: true}
}

func (v NullableCreateProjectApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


