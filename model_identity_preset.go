/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.98
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentityPreset struct for IdentityPreset
type IdentityPreset struct {
	// Schema is the Identity JSON Schema
	Schema map[string]interface{} `json:"schema"`
	// URL is the preset identifier
	Url string `json:"url"`
}

// NewIdentityPreset instantiates a new IdentityPreset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPreset(schema map[string]interface{}, url string) *IdentityPreset {
	this := IdentityPreset{}
	this.Schema = schema
	this.Url = url
	return &this
}

// NewIdentityPresetWithDefaults instantiates a new IdentityPreset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPresetWithDefaults() *IdentityPreset {
	this := IdentityPreset{}
	return &this
}

// GetSchema returns the Schema field value
func (o *IdentityPreset) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *IdentityPreset) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *IdentityPreset) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetUrl returns the Url field value
func (o *IdentityPreset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IdentityPreset) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IdentityPreset) SetUrl(v string) {
	o.Url = v
}

func (o IdentityPreset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schema"] = o.Schema
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityPreset struct {
	value *IdentityPreset
	isSet bool
}

func (v NullableIdentityPreset) Get() *IdentityPreset {
	return v.value
}

func (v *NullableIdentityPreset) Set(val *IdentityPreset) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPreset) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPreset(val *IdentityPreset) *NullableIdentityPreset {
	return &NullableIdentityPreset{value: val, isSet: true}
}

func (v NullableIdentityPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


