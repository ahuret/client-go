/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.22
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Headers Headers is the jwt headers
type Headers struct {
	Extra map[string]interface{} `json:"extra,omitempty"`
}

// NewHeaders instantiates a new Headers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaders() *Headers {
	this := Headers{}
	return &this
}

// NewHeadersWithDefaults instantiates a new Headers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadersWithDefaults() *Headers {
	this := Headers{}
	return &this
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Headers) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Headers) GetExtraOk() (map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Headers) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *Headers) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

func (o Headers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	return json.Marshal(toSerialize)
}

type NullableHeaders struct {
	value *Headers
	isSet bool
}

func (v NullableHeaders) Get() *Headers {
	return v.value
}

func (v *NullableHeaders) Set(val *Headers) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaders(val *Headers) *NullableHeaders {
	return &NullableHeaders{value: val, isSet: true}
}

func (v NullableHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


