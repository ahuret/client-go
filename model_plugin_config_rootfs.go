/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.10
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginConfigRootfs PluginConfigRootfs plugin config rootfs
type PluginConfigRootfs struct {
	// diff ids
	DiffIds []string `json:"diff_ids,omitempty"`
	// type
	Type *string `json:"type,omitempty"`
}

// NewPluginConfigRootfs instantiates a new PluginConfigRootfs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigRootfs() *PluginConfigRootfs {
	this := PluginConfigRootfs{}
	return &this
}

// NewPluginConfigRootfsWithDefaults instantiates a new PluginConfigRootfs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigRootfsWithDefaults() *PluginConfigRootfs {
	this := PluginConfigRootfs{}
	return &this
}

// GetDiffIds returns the DiffIds field value if set, zero value otherwise.
func (o *PluginConfigRootfs) GetDiffIds() []string {
	if o == nil || o.DiffIds == nil {
		var ret []string
		return ret
	}
	return o.DiffIds
}

// GetDiffIdsOk returns a tuple with the DiffIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigRootfs) GetDiffIdsOk() ([]string, bool) {
	if o == nil || o.DiffIds == nil {
		return nil, false
	}
	return o.DiffIds, true
}

// HasDiffIds returns a boolean if a field has been set.
func (o *PluginConfigRootfs) HasDiffIds() bool {
	if o != nil && o.DiffIds != nil {
		return true
	}

	return false
}

// SetDiffIds gets a reference to the given []string and assigns it to the DiffIds field.
func (o *PluginConfigRootfs) SetDiffIds(v []string) {
	o.DiffIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PluginConfigRootfs) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigRootfs) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PluginConfigRootfs) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PluginConfigRootfs) SetType(v string) {
	o.Type = &v
}

func (o PluginConfigRootfs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiffIds != nil {
		toSerialize["diff_ids"] = o.DiffIds
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePluginConfigRootfs struct {
	value *PluginConfigRootfs
	isSet bool
}

func (v NullablePluginConfigRootfs) Get() *PluginConfigRootfs {
	return v.value
}

func (v *NullablePluginConfigRootfs) Set(val *PluginConfigRootfs) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigRootfs) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigRootfs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigRootfs(val *PluginConfigRootfs) *NullablePluginConfigRootfs {
	return &NullablePluginConfigRootfs{value: val, isSet: true}
}

func (v NullablePluginConfigRootfs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigRootfs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


