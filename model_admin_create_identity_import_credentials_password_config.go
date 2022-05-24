/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.181
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdminCreateIdentityImportCredentialsPasswordConfig struct for AdminCreateIdentityImportCredentialsPasswordConfig
type AdminCreateIdentityImportCredentialsPasswordConfig struct {
	// The hashed password in [PHC format]( https://www.ory.sh/docs/kratos/concepts/credentials/username-email-password#hashed-password-format)
	HashedPassword *string `json:"hashed_password,omitempty"`
	// The password in plain text if no hash is available.
	Password *string `json:"password,omitempty"`
}

// NewAdminCreateIdentityImportCredentialsPasswordConfig instantiates a new AdminCreateIdentityImportCredentialsPasswordConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateIdentityImportCredentialsPasswordConfig() *AdminCreateIdentityImportCredentialsPasswordConfig {
	this := AdminCreateIdentityImportCredentialsPasswordConfig{}
	return &this
}

// NewAdminCreateIdentityImportCredentialsPasswordConfigWithDefaults instantiates a new AdminCreateIdentityImportCredentialsPasswordConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateIdentityImportCredentialsPasswordConfigWithDefaults() *AdminCreateIdentityImportCredentialsPasswordConfig {
	this := AdminCreateIdentityImportCredentialsPasswordConfig{}
	return &this
}

// GetHashedPassword returns the HashedPassword field value if set, zero value otherwise.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) GetHashedPassword() string {
	if o == nil || o.HashedPassword == nil {
		var ret string
		return ret
	}
	return *o.HashedPassword
}

// GetHashedPasswordOk returns a tuple with the HashedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) GetHashedPasswordOk() (*string, bool) {
	if o == nil || o.HashedPassword == nil {
		return nil, false
	}
	return o.HashedPassword, true
}

// HasHashedPassword returns a boolean if a field has been set.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) HasHashedPassword() bool {
	if o != nil && o.HashedPassword != nil {
		return true
	}

	return false
}

// SetHashedPassword gets a reference to the given string and assigns it to the HashedPassword field.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) SetHashedPassword(v string) {
	o.HashedPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AdminCreateIdentityImportCredentialsPasswordConfig) SetPassword(v string) {
	o.Password = &v
}

func (o AdminCreateIdentityImportCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HashedPassword != nil {
		toSerialize["hashed_password"] = o.HashedPassword
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableAdminCreateIdentityImportCredentialsPasswordConfig struct {
	value *AdminCreateIdentityImportCredentialsPasswordConfig
	isSet bool
}

func (v NullableAdminCreateIdentityImportCredentialsPasswordConfig) Get() *AdminCreateIdentityImportCredentialsPasswordConfig {
	return v.value
}

func (v *NullableAdminCreateIdentityImportCredentialsPasswordConfig) Set(val *AdminCreateIdentityImportCredentialsPasswordConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateIdentityImportCredentialsPasswordConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateIdentityImportCredentialsPasswordConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateIdentityImportCredentialsPasswordConfig(val *AdminCreateIdentityImportCredentialsPasswordConfig) *NullableAdminCreateIdentityImportCredentialsPasswordConfig {
	return &NullableAdminCreateIdentityImportCredentialsPasswordConfig{value: val, isSet: true}
}

func (v NullableAdminCreateIdentityImportCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateIdentityImportCredentialsPasswordConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


