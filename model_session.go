/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.9
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// Session struct for Session
type Session struct {
	Active *bool `json:"active,omitempty"`
	AuthenticatedAt time.Time `json:"authenticated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Id string `json:"id"`
	Identity Identity `json:"identity"`
	IssuedAt time.Time `json:"issued_at"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(authenticatedAt time.Time, expiresAt time.Time, id string, identity Identity, issuedAt time.Time) *Session {
	this := Session{}
	this.AuthenticatedAt = authenticatedAt
	this.ExpiresAt = expiresAt
	this.Id = id
	this.Identity = identity
	this.IssuedAt = issuedAt
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Session) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Session) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Session) SetActive(v bool) {
	o.Active = &v
}

// GetAuthenticatedAt returns the AuthenticatedAt field value
func (o *Session) GetAuthenticatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.AuthenticatedAt
}

// GetAuthenticatedAtOk returns a tuple with the AuthenticatedAt field value
// and a boolean to check if the value has been set.
func (o *Session) GetAuthenticatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthenticatedAt, true
}

// SetAuthenticatedAt sets field value
func (o *Session) SetAuthenticatedAt(v time.Time) {
	o.AuthenticatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *Session) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *Session) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *Session) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetIdentity returns the Identity field value
func (o *Session) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdentityOk() (*Identity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *Session) SetIdentity(v Identity) {
	o.Identity = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *Session) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *Session) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *Session) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["authenticated_at"] = o.AuthenticatedAt
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["issued_at"] = o.IssuedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


