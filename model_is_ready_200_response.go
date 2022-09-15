/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.45
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IsReady200Response struct for IsReady200Response
type IsReady200Response struct {
	// Always \"ok\".
	Status string `json:"status"`
}

// NewIsReady200Response instantiates a new IsReady200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsReady200Response(status string) *IsReady200Response {
	this := IsReady200Response{}
	this.Status = status
	return &this
}

// NewIsReady200ResponseWithDefaults instantiates a new IsReady200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsReady200ResponseWithDefaults() *IsReady200Response {
	this := IsReady200Response{}
	return &this
}

// GetStatus returns the Status field value
func (o *IsReady200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IsReady200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IsReady200Response) SetStatus(v string) {
	o.Status = v
}

func (o IsReady200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableIsReady200Response struct {
	value *IsReady200Response
	isSet bool
}

func (v NullableIsReady200Response) Get() *IsReady200Response {
	return v.value
}

func (v *NullableIsReady200Response) Set(val *IsReady200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIsReady200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIsReady200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsReady200Response(val *IsReady200Response) *NullableIsReady200Response {
	return &NullableIsReady200Response{value: val, isSet: true}
}

func (v NullableIsReady200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsReady200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


