/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.33
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectServices struct for ProjectServices
type ProjectServices struct {
	Identity *ProjectServiceIdentity `json:"identity,omitempty"`
	Oauth2 *ProjectServiceOAuth2 `json:"oauth2,omitempty"`
	Permission *ProjectServicePermission `json:"permission,omitempty"`
}

// NewProjectServices instantiates a new ProjectServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectServices() *ProjectServices {
	this := ProjectServices{}
	return &this
}

// NewProjectServicesWithDefaults instantiates a new ProjectServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectServicesWithDefaults() *ProjectServices {
	this := ProjectServices{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *ProjectServices) GetIdentity() ProjectServiceIdentity {
	if o == nil || o.Identity == nil {
		var ret ProjectServiceIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectServices) GetIdentityOk() (*ProjectServiceIdentity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *ProjectServices) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given ProjectServiceIdentity and assigns it to the Identity field.
func (o *ProjectServices) SetIdentity(v ProjectServiceIdentity) {
	o.Identity = &v
}

// GetOauth2 returns the Oauth2 field value if set, zero value otherwise.
func (o *ProjectServices) GetOauth2() ProjectServiceOAuth2 {
	if o == nil || o.Oauth2 == nil {
		var ret ProjectServiceOAuth2
		return ret
	}
	return *o.Oauth2
}

// GetOauth2Ok returns a tuple with the Oauth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectServices) GetOauth2Ok() (*ProjectServiceOAuth2, bool) {
	if o == nil || o.Oauth2 == nil {
		return nil, false
	}
	return o.Oauth2, true
}

// HasOauth2 returns a boolean if a field has been set.
func (o *ProjectServices) HasOauth2() bool {
	if o != nil && o.Oauth2 != nil {
		return true
	}

	return false
}

// SetOauth2 gets a reference to the given ProjectServiceOAuth2 and assigns it to the Oauth2 field.
func (o *ProjectServices) SetOauth2(v ProjectServiceOAuth2) {
	o.Oauth2 = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ProjectServices) GetPermission() ProjectServicePermission {
	if o == nil || o.Permission == nil {
		var ret ProjectServicePermission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectServices) GetPermissionOk() (*ProjectServicePermission, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ProjectServices) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given ProjectServicePermission and assigns it to the Permission field.
func (o *ProjectServices) SetPermission(v ProjectServicePermission) {
	o.Permission = &v
}

func (o ProjectServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Oauth2 != nil {
		toSerialize["oauth2"] = o.Oauth2
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableProjectServices struct {
	value *ProjectServices
	isSet bool
}

func (v NullableProjectServices) Get() *ProjectServices {
	return v.value
}

func (v *NullableProjectServices) Set(val *ProjectServices) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectServices) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectServices(val *ProjectServices) *NullableProjectServices {
	return &NullableProjectServices{value: val, isSet: true}
}

func (v NullableProjectServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


