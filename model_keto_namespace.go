/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.23
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// KetoNamespace struct for KetoNamespace
type KetoNamespace struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewKetoNamespace instantiates a new KetoNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKetoNamespace() *KetoNamespace {
	this := KetoNamespace{}
	return &this
}

// NewKetoNamespaceWithDefaults instantiates a new KetoNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKetoNamespaceWithDefaults() *KetoNamespace {
	this := KetoNamespace{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KetoNamespace) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KetoNamespace) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KetoNamespace) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *KetoNamespace) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KetoNamespace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KetoNamespace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KetoNamespace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KetoNamespace) SetName(v string) {
	o.Name = &v
}

func (o KetoNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableKetoNamespace struct {
	value *KetoNamespace
	isSet bool
}

func (v NullableKetoNamespace) Get() *KetoNamespace {
	return v.value
}

func (v *NullableKetoNamespace) Set(val *KetoNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableKetoNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableKetoNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKetoNamespace(val *KetoNamespace) *NullableKetoNamespace {
	return &NullableKetoNamespace{value: val, isSet: true}
}

func (v NullableKetoNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKetoNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


