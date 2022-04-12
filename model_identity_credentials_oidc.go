/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.161
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentityCredentialsOidc struct for IdentityCredentialsOidc
type IdentityCredentialsOidc struct {
	Providers []IdentityCredentialsOidcProvider `json:"providers,omitempty"`
}

// NewIdentityCredentialsOidc instantiates a new IdentityCredentialsOidc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCredentialsOidc() *IdentityCredentialsOidc {
	this := IdentityCredentialsOidc{}
	return &this
}

// NewIdentityCredentialsOidcWithDefaults instantiates a new IdentityCredentialsOidc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCredentialsOidcWithDefaults() *IdentityCredentialsOidc {
	this := IdentityCredentialsOidc{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *IdentityCredentialsOidc) GetProviders() []IdentityCredentialsOidcProvider {
	if o == nil || o.Providers == nil {
		var ret []IdentityCredentialsOidcProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityCredentialsOidc) GetProvidersOk() ([]IdentityCredentialsOidcProvider, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *IdentityCredentialsOidc) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []IdentityCredentialsOidcProvider and assigns it to the Providers field.
func (o *IdentityCredentialsOidc) SetProviders(v []IdentityCredentialsOidcProvider) {
	o.Providers = v
}

func (o IdentityCredentialsOidc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityCredentialsOidc struct {
	value *IdentityCredentialsOidc
	isSet bool
}

func (v NullableIdentityCredentialsOidc) Get() *IdentityCredentialsOidc {
	return v.value
}

func (v *NullableIdentityCredentialsOidc) Set(val *IdentityCredentialsOidc) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialsOidc) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialsOidc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialsOidc(val *IdentityCredentialsOidc) *NullableIdentityCredentialsOidc {
	return &NullableIdentityCredentialsOidc{value: val, isSet: true}
}

func (v NullableIdentityCredentialsOidc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialsOidc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


