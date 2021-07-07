/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.12
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UiNodeTextAttributes struct for UiNodeTextAttributes
type UiNodeTextAttributes struct {
	Text UiText `json:"text"`
}

// NewUiNodeTextAttributes instantiates a new UiNodeTextAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeTextAttributes(text UiText) *UiNodeTextAttributes {
	this := UiNodeTextAttributes{}
	this.Text = text
	return &this
}

// NewUiNodeTextAttributesWithDefaults instantiates a new UiNodeTextAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeTextAttributesWithDefaults() *UiNodeTextAttributes {
	this := UiNodeTextAttributes{}
	return &this
}

// GetText returns the Text field value
func (o *UiNodeTextAttributes) GetText() UiText {
	if o == nil {
		var ret UiText
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *UiNodeTextAttributes) GetTextOk() (*UiText, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *UiNodeTextAttributes) SetText(v UiText) {
	o.Text = v
}

func (o UiNodeTextAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableUiNodeTextAttributes struct {
	value *UiNodeTextAttributes
	isSet bool
}

func (v NullableUiNodeTextAttributes) Get() *UiNodeTextAttributes {
	return v.value
}

func (v *NullableUiNodeTextAttributes) Set(val *UiNodeTextAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeTextAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeTextAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeTextAttributes(val *UiNodeTextAttributes) *NullableUiNodeTextAttributes {
	return &NullableUiNodeTextAttributes{value: val, isSet: true}
}

func (v NullableUiNodeTextAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeTextAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


