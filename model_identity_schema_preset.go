/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the IdentitySchemaPreset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySchemaPreset{}

// IdentitySchemaPreset struct for IdentitySchemaPreset
type IdentitySchemaPreset struct {
	// Schema is the Identity JSON Schema
	Schema map[string]interface{} `json:"schema"`
	// URL is the preset identifier
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySchemaPreset IdentitySchemaPreset

// NewIdentitySchemaPreset instantiates a new IdentitySchemaPreset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySchemaPreset(schema map[string]interface{}, url string) *IdentitySchemaPreset {
	this := IdentitySchemaPreset{}
	this.Schema = schema
	this.Url = url
	return &this
}

// NewIdentitySchemaPresetWithDefaults instantiates a new IdentitySchemaPreset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySchemaPresetWithDefaults() *IdentitySchemaPreset {
	this := IdentitySchemaPreset{}
	return &this
}

// GetSchema returns the Schema field value
func (o *IdentitySchemaPreset) GetSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *IdentitySchemaPreset) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// SetSchema sets field value
func (o *IdentitySchemaPreset) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetUrl returns the Url field value
func (o *IdentitySchemaPreset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IdentitySchemaPreset) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IdentitySchemaPreset) SetUrl(v string) {
	o.Url = v
}

func (o IdentitySchemaPreset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySchemaPreset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySchemaPreset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdentitySchemaPreset := _IdentitySchemaPreset{}

	err = json.Unmarshal(data, &varIdentitySchemaPreset)

	if err != nil {
		return err
	}

	*o = IdentitySchemaPreset(varIdentitySchemaPreset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schema")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySchemaPreset struct {
	value *IdentitySchemaPreset
	isSet bool
}

func (v NullableIdentitySchemaPreset) Get() *IdentitySchemaPreset {
	return v.value
}

func (v *NullableIdentitySchemaPreset) Set(val *IdentitySchemaPreset) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySchemaPreset) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySchemaPreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySchemaPreset(val *IdentitySchemaPreset) *NullableIdentitySchemaPreset {
	return &NullableIdentitySchemaPreset{value: val, isSet: true}
}

func (v NullableIdentitySchemaPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySchemaPreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


