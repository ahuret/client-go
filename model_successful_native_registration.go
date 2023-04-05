/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.24
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SuccessfulNativeRegistration The Response for Registration Flows via API
type SuccessfulNativeRegistration struct {
	// Contains a list of actions, that could follow this flow  It can, for example, this will contain a reference to the verification flow, created as part of the user's registration or the token of the session.
	ContinueWith []ContinueWith `json:"continue_with,omitempty"`
	Identity Identity `json:"identity"`
	Session *Session `json:"session,omitempty"`
	// The Session Token  This field is only set when the session hook is configured as a post-registration hook.  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
	SessionToken *string `json:"session_token,omitempty"`
}

// NewSuccessfulNativeRegistration instantiates a new SuccessfulNativeRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulNativeRegistration(identity Identity) *SuccessfulNativeRegistration {
	this := SuccessfulNativeRegistration{}
	this.Identity = identity
	return &this
}

// NewSuccessfulNativeRegistrationWithDefaults instantiates a new SuccessfulNativeRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulNativeRegistrationWithDefaults() *SuccessfulNativeRegistration {
	this := SuccessfulNativeRegistration{}
	return &this
}

// GetContinueWith returns the ContinueWith field value if set, zero value otherwise.
func (o *SuccessfulNativeRegistration) GetContinueWith() []ContinueWith {
	if o == nil || o.ContinueWith == nil {
		var ret []ContinueWith
		return ret
	}
	return o.ContinueWith
}

// GetContinueWithOk returns a tuple with the ContinueWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeRegistration) GetContinueWithOk() ([]ContinueWith, bool) {
	if o == nil || o.ContinueWith == nil {
		return nil, false
	}
	return o.ContinueWith, true
}

// HasContinueWith returns a boolean if a field has been set.
func (o *SuccessfulNativeRegistration) HasContinueWith() bool {
	if o != nil && o.ContinueWith != nil {
		return true
	}

	return false
}

// SetContinueWith gets a reference to the given []ContinueWith and assigns it to the ContinueWith field.
func (o *SuccessfulNativeRegistration) SetContinueWith(v []ContinueWith) {
	o.ContinueWith = v
}

// GetIdentity returns the Identity field value
func (o *SuccessfulNativeRegistration) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeRegistration) GetIdentityOk() (*Identity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *SuccessfulNativeRegistration) SetIdentity(v Identity) {
	o.Identity = v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *SuccessfulNativeRegistration) GetSession() Session {
	if o == nil || o.Session == nil {
		var ret Session
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeRegistration) GetSessionOk() (*Session, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *SuccessfulNativeRegistration) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given Session and assigns it to the Session field.
func (o *SuccessfulNativeRegistration) SetSession(v Session) {
	o.Session = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *SuccessfulNativeRegistration) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulNativeRegistration) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *SuccessfulNativeRegistration) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *SuccessfulNativeRegistration) SetSessionToken(v string) {
	o.SessionToken = &v
}

func (o SuccessfulNativeRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContinueWith != nil {
		toSerialize["continue_with"] = o.ContinueWith
	}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	if o.SessionToken != nil {
		toSerialize["session_token"] = o.SessionToken
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessfulNativeRegistration struct {
	value *SuccessfulNativeRegistration
	isSet bool
}

func (v NullableSuccessfulNativeRegistration) Get() *SuccessfulNativeRegistration {
	return v.value
}

func (v *NullableSuccessfulNativeRegistration) Set(val *SuccessfulNativeRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulNativeRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulNativeRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulNativeRegistration(val *SuccessfulNativeRegistration) *NullableSuccessfulNativeRegistration {
	return &NullableSuccessfulNativeRegistration{value: val, isSet: true}
}

func (v NullableSuccessfulNativeRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulNativeRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


