/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.39
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UiNodeMeta This might include a label and other information that can optionally be used to render UIs.
type UiNodeMeta struct {
	Label *UiText `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiNodeMeta UiNodeMeta

// NewUiNodeMeta instantiates a new UiNodeMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeMeta() *UiNodeMeta {
	this := UiNodeMeta{}
	return &this
}

// NewUiNodeMetaWithDefaults instantiates a new UiNodeMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeMetaWithDefaults() *UiNodeMeta {
	this := UiNodeMeta{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UiNodeMeta) GetLabel() UiText {
	if o == nil || o.Label == nil {
		var ret UiText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeMeta) GetLabelOk() (*UiText, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UiNodeMeta) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given UiText and assigns it to the Label field.
func (o *UiNodeMeta) SetLabel(v UiText) {
	o.Label = &v
}

func (o UiNodeMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UiNodeMeta) UnmarshalJSON(bytes []byte) (err error) {
	varUiNodeMeta := _UiNodeMeta{}

	if err = json.Unmarshal(bytes, &varUiNodeMeta); err == nil {
		*o = UiNodeMeta(varUiNodeMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiNodeMeta struct {
	value *UiNodeMeta
	isSet bool
}

func (v NullableUiNodeMeta) Get() *UiNodeMeta {
	return v.value
}

func (v *NullableUiNodeMeta) Set(val *UiNodeMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeMeta(val *UiNodeMeta) *NullableUiNodeMeta {
	return &NullableUiNodeMeta{value: val, isSet: true}
}

func (v NullableUiNodeMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


