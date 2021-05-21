/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.8
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ContainerTopOKBody ContainerTopOKBody OK response to ContainerTop operation
type ContainerTopOKBody struct {
	// Each process running in the container, where each is process is an array of values corresponding to the titles
	Processes [][]string `json:"Processes"`
	// The ps column titles
	Titles []string `json:"Titles"`
}

// NewContainerTopOKBody instantiates a new ContainerTopOKBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerTopOKBody(processes [][]string, titles []string) *ContainerTopOKBody {
	this := ContainerTopOKBody{}
	this.Processes = processes
	this.Titles = titles
	return &this
}

// NewContainerTopOKBodyWithDefaults instantiates a new ContainerTopOKBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTopOKBodyWithDefaults() *ContainerTopOKBody {
	this := ContainerTopOKBody{}
	return &this
}

// GetProcesses returns the Processes field value
func (o *ContainerTopOKBody) GetProcesses() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
func (o *ContainerTopOKBody) GetProcessesOk() ([][]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ContainerTopOKBody) SetProcesses(v [][]string) {
	o.Processes = v
}

// GetTitles returns the Titles field value
func (o *ContainerTopOKBody) GetTitles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Titles
}

// GetTitlesOk returns a tuple with the Titles field value
// and a boolean to check if the value has been set.
func (o *ContainerTopOKBody) GetTitlesOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Titles, true
}

// SetTitles sets field value
func (o *ContainerTopOKBody) SetTitles(v []string) {
	o.Titles = v
}

func (o ContainerTopOKBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Processes"] = o.Processes
	}
	if true {
		toSerialize["Titles"] = o.Titles
	}
	return json.Marshal(toSerialize)
}

type NullableContainerTopOKBody struct {
	value *ContainerTopOKBody
	isSet bool
}

func (v NullableContainerTopOKBody) Get() *ContainerTopOKBody {
	return v.value
}

func (v *NullableContainerTopOKBody) Set(val *ContainerTopOKBody) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerTopOKBody) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerTopOKBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerTopOKBody(val *ContainerTopOKBody) *NullableContainerTopOKBody {
	return &NullableContainerTopOKBody{value: val, isSet: true}
}

func (v NullableContainerTopOKBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerTopOKBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


