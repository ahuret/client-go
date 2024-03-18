/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.9.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PermissionsOnWorkpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsOnWorkpaceResponse{}

// PermissionsOnWorkpaceResponse Get Permissions on Project Request Parameters
type PermissionsOnWorkpaceResponse struct {
	Permissions *map[string]bool `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionsOnWorkpaceResponse PermissionsOnWorkpaceResponse

// NewPermissionsOnWorkpaceResponse instantiates a new PermissionsOnWorkpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsOnWorkpaceResponse() *PermissionsOnWorkpaceResponse {
	this := PermissionsOnWorkpaceResponse{}
	return &this
}

// NewPermissionsOnWorkpaceResponseWithDefaults instantiates a new PermissionsOnWorkpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsOnWorkpaceResponseWithDefaults() *PermissionsOnWorkpaceResponse {
	this := PermissionsOnWorkpaceResponse{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PermissionsOnWorkpaceResponse) GetPermissions() map[string]bool {
	if o == nil || IsNil(o.Permissions) {
		var ret map[string]bool
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnWorkpaceResponse) GetPermissionsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PermissionsOnWorkpaceResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given map[string]bool and assigns it to the Permissions field.
func (o *PermissionsOnWorkpaceResponse) SetPermissions(v map[string]bool) {
	o.Permissions = &v
}

func (o PermissionsOnWorkpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsOnWorkpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PermissionsOnWorkpaceResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPermissionsOnWorkpaceResponse := _PermissionsOnWorkpaceResponse{}

	err = json.Unmarshal(bytes, &varPermissionsOnWorkpaceResponse)

	if err != nil {
		return err
	}

	*o = PermissionsOnWorkpaceResponse(varPermissionsOnWorkpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissionsOnWorkpaceResponse struct {
	value *PermissionsOnWorkpaceResponse
	isSet bool
}

func (v NullablePermissionsOnWorkpaceResponse) Get() *PermissionsOnWorkpaceResponse {
	return v.value
}

func (v *NullablePermissionsOnWorkpaceResponse) Set(val *PermissionsOnWorkpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsOnWorkpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsOnWorkpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsOnWorkpaceResponse(val *PermissionsOnWorkpaceResponse) *NullablePermissionsOnWorkpaceResponse {
	return &NullablePermissionsOnWorkpaceResponse{value: val, isSet: true}
}

func (v NullablePermissionsOnWorkpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsOnWorkpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


