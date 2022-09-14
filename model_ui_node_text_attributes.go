/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.43
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UiNodeTextAttributes struct for UiNodeTextAttributes
type UiNodeTextAttributes struct {
	// A unique identifier
	Id string `json:"id"`
	// NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"text\".
	NodeType string `json:"node_type"`
	Text UiText `json:"text"`
}

// NewUiNodeTextAttributes instantiates a new UiNodeTextAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeTextAttributes(id string, nodeType string, text UiText) *UiNodeTextAttributes {
	this := UiNodeTextAttributes{}
	this.Id = id
	this.NodeType = nodeType
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

// GetId returns the Id field value
func (o *UiNodeTextAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UiNodeTextAttributes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UiNodeTextAttributes) SetId(v string) {
	o.Id = v
}

// GetNodeType returns the NodeType field value
func (o *UiNodeTextAttributes) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *UiNodeTextAttributes) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *UiNodeTextAttributes) SetNodeType(v string) {
	o.NodeType = v
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
	if o == nil {
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
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_type"] = o.NodeType
	}
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


