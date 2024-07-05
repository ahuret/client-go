/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PermissionsOnWorkspace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsOnWorkspace{}

// PermissionsOnWorkspace Get Permissions on Project Request Parameters
type PermissionsOnWorkspace struct {
	Permissions *map[string]bool `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionsOnWorkspace PermissionsOnWorkspace

// NewPermissionsOnWorkspace instantiates a new PermissionsOnWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsOnWorkspace() *PermissionsOnWorkspace {
	this := PermissionsOnWorkspace{}
	return &this
}

// NewPermissionsOnWorkspaceWithDefaults instantiates a new PermissionsOnWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsOnWorkspaceWithDefaults() *PermissionsOnWorkspace {
	this := PermissionsOnWorkspace{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PermissionsOnWorkspace) GetPermissions() map[string]bool {
	if o == nil || IsNil(o.Permissions) {
		var ret map[string]bool
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnWorkspace) GetPermissionsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PermissionsOnWorkspace) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]bool and assigns it to the Permissions field.
func (o *PermissionsOnWorkspace) SetPermissions(v map[string]bool) {
	o.Permissions = &v
}

func (o PermissionsOnWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsOnWorkspace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PermissionsOnWorkspace) UnmarshalJSON(data []byte) (err error) {
	varPermissionsOnWorkspace := _PermissionsOnWorkspace{}

	err = json.Unmarshal(data, &varPermissionsOnWorkspace)

	if err != nil {
		return err
	}

	*o = PermissionsOnWorkspace(varPermissionsOnWorkspace)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissionsOnWorkspace struct {
	value *PermissionsOnWorkspace
	isSet bool
}

func (v NullablePermissionsOnWorkspace) Get() *PermissionsOnWorkspace {
	return v.value
}

func (v *NullablePermissionsOnWorkspace) Set(val *PermissionsOnWorkspace) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsOnWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsOnWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsOnWorkspace(val *PermissionsOnWorkspace) *NullablePermissionsOnWorkspace {
	return &NullablePermissionsOnWorkspace{value: val, isSet: true}
}

func (v NullablePermissionsOnWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsOnWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


