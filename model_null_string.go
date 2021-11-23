/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.28
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NullString var s NullString err := db.QueryRow(\"SELECT name FROM foo WHERE id=?\", id).Scan(&s) ... if s.Valid { use s.String } else { NULL value }
type NullString struct {
	String *string `json:"String,omitempty"`
	Valid *bool `json:"Valid,omitempty"`
}

// NewNullString instantiates a new NullString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullString() *NullString {
	this := NullString{}
	return &this
}

// NewNullStringWithDefaults instantiates a new NullString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullStringWithDefaults() *NullString {
	this := NullString{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *NullString) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullString) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *NullString) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *NullString) SetString(v string) {
	o.String = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *NullString) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullString) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *NullString) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *NullString) SetValid(v bool) {
	o.Valid = &v
}

func (o NullString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.String != nil {
		toSerialize["String"] = o.String
	}
	if o.Valid != nil {
		toSerialize["Valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableNullString struct {
	value *NullString
	isSet bool
}

func (v NullableNullString) Get() *NullString {
	return v.value
}

func (v *NullableNullString) Set(val *NullString) {
	v.value = val
	v.isSet = true
}

func (v NullableNullString) IsSet() bool {
	return v.isSet
}

func (v *NullableNullString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullString(val *NullString) *NullableNullString {
	return &NullableNullString{value: val, isSet: true}
}

func (v NullableNullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


