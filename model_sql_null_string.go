/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: latest
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SQLNullString struct for SQLNullString
type SQLNullString struct {
	String *string `json:"String,omitempty"`
	Valid *bool `json:"Valid,omitempty"`
}

// NewSQLNullString instantiates a new SQLNullString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSQLNullString() *SQLNullString {
	this := SQLNullString{}
	return &this
}

// NewSQLNullStringWithDefaults instantiates a new SQLNullString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSQLNullStringWithDefaults() *SQLNullString {
	this := SQLNullString{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *SQLNullString) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SQLNullString) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *SQLNullString) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *SQLNullString) SetString(v string) {
	o.String = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *SQLNullString) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SQLNullString) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *SQLNullString) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *SQLNullString) SetValid(v bool) {
	o.Valid = &v
}

func (o SQLNullString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.String != nil {
		toSerialize["String"] = o.String
	}
	if o.Valid != nil {
		toSerialize["Valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableSQLNullString struct {
	value *SQLNullString
	isSet bool
}

func (v NullableSQLNullString) Get() *SQLNullString {
	return v.value
}

func (v *NullableSQLNullString) Set(val *SQLNullString) {
	v.value = val
	v.isSet = true
}

func (v NullableSQLNullString) IsSet() bool {
	return v.isSet
}

func (v *NullableSQLNullString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSQLNullString(val *SQLNullString) *NullableSQLNullString {
	return &NullableSQLNullString{value: val, isSet: true}
}

func (v NullableSQLNullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSQLNullString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


