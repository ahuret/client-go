/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.56
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// SessionAuthenticationMethod A singular authenticator used during authentication / login.
type SessionAuthenticationMethod struct {
	// When the authentication challenge was completed.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Method *string `json:"method,omitempty"`
}

// NewSessionAuthenticationMethod instantiates a new SessionAuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionAuthenticationMethod() *SessionAuthenticationMethod {
	this := SessionAuthenticationMethod{}
	return &this
}

// NewSessionAuthenticationMethodWithDefaults instantiates a new SessionAuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionAuthenticationMethodWithDefaults() *SessionAuthenticationMethod {
	this := SessionAuthenticationMethod{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *SessionAuthenticationMethod) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAuthenticationMethod) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *SessionAuthenticationMethod) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *SessionAuthenticationMethod) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SessionAuthenticationMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAuthenticationMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SessionAuthenticationMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SessionAuthenticationMethod) SetMethod(v string) {
	o.Method = &v
}

func (o SessionAuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompletedAt != nil {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableSessionAuthenticationMethod struct {
	value *SessionAuthenticationMethod
	isSet bool
}

func (v NullableSessionAuthenticationMethod) Get() *SessionAuthenticationMethod {
	return v.value
}

func (v *NullableSessionAuthenticationMethod) Set(val *SessionAuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionAuthenticationMethod(val *SessionAuthenticationMethod) *NullableSessionAuthenticationMethod {
	return &NullableSessionAuthenticationMethod{value: val, isSet: true}
}

func (v NullableSessionAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


