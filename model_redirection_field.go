/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.55
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RedirectionField struct for RedirectionField
type RedirectionField struct {
	Main *string `json:"main,omitempty"`
	Oidc *string `json:"oidc,omitempty"`
	Password *string `json:"password,omitempty"`
	Profile *string `json:"profile,omitempty"`
}

// NewRedirectionField instantiates a new RedirectionField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectionField() *RedirectionField {
	this := RedirectionField{}
	return &this
}

// NewRedirectionFieldWithDefaults instantiates a new RedirectionField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectionFieldWithDefaults() *RedirectionField {
	this := RedirectionField{}
	return &this
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *RedirectionField) GetMain() string {
	if o == nil || o.Main == nil {
		var ret string
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionField) GetMainOk() (*string, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *RedirectionField) HasMain() bool {
	if o != nil && o.Main != nil {
		return true
	}

	return false
}

// SetMain gets a reference to the given string and assigns it to the Main field.
func (o *RedirectionField) SetMain(v string) {
	o.Main = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *RedirectionField) GetOidc() string {
	if o == nil || o.Oidc == nil {
		var ret string
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionField) GetOidcOk() (*string, bool) {
	if o == nil || o.Oidc == nil {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *RedirectionField) HasOidc() bool {
	if o != nil && o.Oidc != nil {
		return true
	}

	return false
}

// SetOidc gets a reference to the given string and assigns it to the Oidc field.
func (o *RedirectionField) SetOidc(v string) {
	o.Oidc = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RedirectionField) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionField) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RedirectionField) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RedirectionField) SetPassword(v string) {
	o.Password = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *RedirectionField) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionField) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *RedirectionField) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *RedirectionField) SetProfile(v string) {
	o.Profile = &v
}

func (o RedirectionField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Main != nil {
		toSerialize["main"] = o.Main
	}
	if o.Oidc != nil {
		toSerialize["oidc"] = o.Oidc
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectionField struct {
	value *RedirectionField
	isSet bool
}

func (v NullableRedirectionField) Get() *RedirectionField {
	return v.value
}

func (v *NullableRedirectionField) Set(val *RedirectionField) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectionField) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectionField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectionField(val *RedirectionField) *NullableRedirectionField {
	return &NullableRedirectionField{value: val, isSet: true}
}

func (v NullableRedirectionField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectionField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


