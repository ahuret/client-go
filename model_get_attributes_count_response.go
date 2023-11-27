/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GetAttributesCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAttributesCountResponse{}

// GetAttributesCountResponse Response of the getAttributesCount endpoint
type GetAttributesCountResponse struct {
	// The list of data points.
	Data []AttributesCountDatapoint `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _GetAttributesCountResponse GetAttributesCountResponse

// NewGetAttributesCountResponse instantiates a new GetAttributesCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAttributesCountResponse(data []AttributesCountDatapoint) *GetAttributesCountResponse {
	this := GetAttributesCountResponse{}
	this.Data = data
	return &this
}

// NewGetAttributesCountResponseWithDefaults instantiates a new GetAttributesCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAttributesCountResponseWithDefaults() *GetAttributesCountResponse {
	this := GetAttributesCountResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetAttributesCountResponse) GetData() []AttributesCountDatapoint {
	if o == nil {
		var ret []AttributesCountDatapoint
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAttributesCountResponse) GetDataOk() ([]AttributesCountDatapoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetAttributesCountResponse) SetData(v []AttributesCountDatapoint) {
	o.Data = v
}

func (o GetAttributesCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAttributesCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAttributesCountResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetAttributesCountResponse := _GetAttributesCountResponse{}

	err = json.Unmarshal(bytes, &varGetAttributesCountResponse)

	if err != nil {
		return err
	}

	*o = GetAttributesCountResponse(varGetAttributesCountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAttributesCountResponse struct {
	value *GetAttributesCountResponse
	isSet bool
}

func (v NullableGetAttributesCountResponse) Get() *GetAttributesCountResponse {
	return v.value
}

func (v *NullableGetAttributesCountResponse) Set(val *GetAttributesCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAttributesCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAttributesCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAttributesCountResponse(val *GetAttributesCountResponse) *NullableGetAttributesCountResponse {
	return &NullableGetAttributesCountResponse{value: val, isSet: true}
}

func (v NullableGetAttributesCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAttributesCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


