/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.9
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GraphDriverData struct for GraphDriverData
type GraphDriverData struct {
	// data
	Data map[string]string `json:"Data"`
	// name
	Name string `json:"Name"`
}

// NewGraphDriverData instantiates a new GraphDriverData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphDriverData(data map[string]string, name string) *GraphDriverData {
	this := GraphDriverData{}
	this.Data = data
	this.Name = name
	return &this
}

// NewGraphDriverDataWithDefaults instantiates a new GraphDriverData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphDriverDataWithDefaults() *GraphDriverData {
	this := GraphDriverData{}
	return &this
}

// GetData returns the Data field value
func (o *GraphDriverData) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GraphDriverData) GetDataOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GraphDriverData) SetData(v map[string]string) {
	o.Data = v
}

// GetName returns the Name field value
func (o *GraphDriverData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GraphDriverData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GraphDriverData) SetName(v string) {
	o.Name = v
}

func (o GraphDriverData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Data"] = o.Data
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGraphDriverData struct {
	value *GraphDriverData
	isSet bool
}

func (v NullableGraphDriverData) Get() *GraphDriverData {
	return v.value
}

func (v *NullableGraphDriverData) Set(val *GraphDriverData) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphDriverData) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphDriverData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphDriverData(val *GraphDriverData) *NullableGraphDriverData {
	return &NullableGraphDriverData{value: val, isSet: true}
}

func (v NullableGraphDriverData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphDriverData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


