/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.154
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdminIdentityImportCredentials struct for AdminIdentityImportCredentials
type AdminIdentityImportCredentials struct {
	Oidc *AdminCreateIdentityImportCredentialsOidc `json:"oidc,omitempty"`
	Password *AdminCreateIdentityImportCredentialsPassword `json:"password,omitempty"`
}

// NewAdminIdentityImportCredentials instantiates a new AdminIdentityImportCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminIdentityImportCredentials() *AdminIdentityImportCredentials {
	this := AdminIdentityImportCredentials{}
	return &this
}

// NewAdminIdentityImportCredentialsWithDefaults instantiates a new AdminIdentityImportCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminIdentityImportCredentialsWithDefaults() *AdminIdentityImportCredentials {
	this := AdminIdentityImportCredentials{}
	return &this
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *AdminIdentityImportCredentials) GetOidc() AdminCreateIdentityImportCredentialsOidc {
	if o == nil || o.Oidc == nil {
		var ret AdminCreateIdentityImportCredentialsOidc
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminIdentityImportCredentials) GetOidcOk() (*AdminCreateIdentityImportCredentialsOidc, bool) {
	if o == nil || o.Oidc == nil {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *AdminIdentityImportCredentials) HasOidc() bool {
	if o != nil && o.Oidc != nil {
		return true
	}

	return false
}

// SetOidc gets a reference to the given AdminCreateIdentityImportCredentialsOidc and assigns it to the Oidc field.
func (o *AdminIdentityImportCredentials) SetOidc(v AdminCreateIdentityImportCredentialsOidc) {
	o.Oidc = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AdminIdentityImportCredentials) GetPassword() AdminCreateIdentityImportCredentialsPassword {
	if o == nil || o.Password == nil {
		var ret AdminCreateIdentityImportCredentialsPassword
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminIdentityImportCredentials) GetPasswordOk() (*AdminCreateIdentityImportCredentialsPassword, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AdminIdentityImportCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given AdminCreateIdentityImportCredentialsPassword and assigns it to the Password field.
func (o *AdminIdentityImportCredentials) SetPassword(v AdminCreateIdentityImportCredentialsPassword) {
	o.Password = &v
}

func (o AdminIdentityImportCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Oidc != nil {
		toSerialize["oidc"] = o.Oidc
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableAdminIdentityImportCredentials struct {
	value *AdminIdentityImportCredentials
	isSet bool
}

func (v NullableAdminIdentityImportCredentials) Get() *AdminIdentityImportCredentials {
	return v.value
}

func (v *NullableAdminIdentityImportCredentials) Set(val *AdminIdentityImportCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminIdentityImportCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminIdentityImportCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminIdentityImportCredentials(val *AdminIdentityImportCredentials) *NullableAdminIdentityImportCredentials {
	return &NullableAdminIdentityImportCredentials{value: val, isSet: true}
}

func (v NullableAdminIdentityImportCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminIdentityImportCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


