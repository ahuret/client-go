/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GetSessionActivityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionActivityResponse{}

// GetSessionActivityResponse Response of the getSessionActivity endpoint
type GetSessionActivityResponse struct {
	// The list of data points.
	Data []SessionActivityDatapoint `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _GetSessionActivityResponse GetSessionActivityResponse

// NewGetSessionActivityResponse instantiates a new GetSessionActivityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionActivityResponse(data []SessionActivityDatapoint) *GetSessionActivityResponse {
	this := GetSessionActivityResponse{}
	this.Data = data
	return &this
}

// NewGetSessionActivityResponseWithDefaults instantiates a new GetSessionActivityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionActivityResponseWithDefaults() *GetSessionActivityResponse {
	this := GetSessionActivityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetSessionActivityResponse) GetData() []SessionActivityDatapoint {
	if o == nil {
		var ret []SessionActivityDatapoint
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSessionActivityResponse) GetDataOk() ([]SessionActivityDatapoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetSessionActivityResponse) SetData(v []SessionActivityDatapoint) {
	o.Data = v
}

func (o GetSessionActivityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionActivityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSessionActivityResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetSessionActivityResponse := _GetSessionActivityResponse{}

	err = json.Unmarshal(data, &varGetSessionActivityResponse)

	if err != nil {
		return err
	}

	*o = GetSessionActivityResponse(varGetSessionActivityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSessionActivityResponse struct {
	value *GetSessionActivityResponse
	isSet bool
}

func (v NullableGetSessionActivityResponse) Get() *GetSessionActivityResponse {
	return v.value
}

func (v *NullableGetSessionActivityResponse) Set(val *GetSessionActivityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionActivityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionActivityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionActivityResponse(val *GetSessionActivityResponse) *NullableGetSessionActivityResponse {
	return &NullableGetSessionActivityResponse{value: val, isSet: true}
}

func (v NullableGetSessionActivityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionActivityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


