/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.11
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginEnv PluginEnv plugin env
type PluginEnv struct {
	// description
	Description string `json:"Description"`
	// name
	Name string `json:"Name"`
	// settable
	Settable []string `json:"Settable"`
	// value
	Value string `json:"Value"`
}

// NewPluginEnv instantiates a new PluginEnv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginEnv(description string, name string, settable []string, value string) *PluginEnv {
	this := PluginEnv{}
	this.Description = description
	this.Name = name
	this.Settable = settable
	this.Value = value
	return &this
}

// NewPluginEnvWithDefaults instantiates a new PluginEnv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginEnvWithDefaults() *PluginEnv {
	this := PluginEnv{}
	return &this
}

// GetDescription returns the Description field value
func (o *PluginEnv) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginEnv) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginEnv) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *PluginEnv) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PluginEnv) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PluginEnv) SetName(v string) {
	o.Name = v
}

// GetSettable returns the Settable field value
func (o *PluginEnv) GetSettable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value
// and a boolean to check if the value has been set.
func (o *PluginEnv) GetSettableOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settable, true
}

// SetSettable sets field value
func (o *PluginEnv) SetSettable(v []string) {
	o.Settable = v
}

// GetValue returns the Value field value
func (o *PluginEnv) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PluginEnv) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PluginEnv) SetValue(v string) {
	o.Value = v
}

func (o PluginEnv) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Settable"] = o.Settable
	}
	if true {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePluginEnv struct {
	value *PluginEnv
	isSet bool
}

func (v NullablePluginEnv) Get() *PluginEnv {
	return v.value
}

func (v *NullablePluginEnv) Set(val *PluginEnv) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginEnv) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginEnv(val *PluginEnv) *NullablePluginEnv {
	return &NullablePluginEnv{value: val, isSet: true}
}

func (v NullablePluginEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


