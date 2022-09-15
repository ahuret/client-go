/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.44
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// TrustJwtGrantIssuerBody struct for TrustJwtGrantIssuerBody
type TrustJwtGrantIssuerBody struct {
	// The \"allow_any_subject\" indicates that the issuer is allowed to have any principal as the subject of the JWT.
	AllowAnySubject *bool `json:"allow_any_subject,omitempty"`
	// The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".
	ExpiresAt time.Time `json:"expires_at"`
	// The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).
	Issuer string `json:"issuer"`
	Jwk JSONWebKey `json:"jwk"`
	// The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
	Scope []string `json:"scope"`
	// The \"subject\" identifies the principal that is the subject of the JWT.
	Subject *string `json:"subject,omitempty"`
}

// NewTrustJwtGrantIssuerBody instantiates a new TrustJwtGrantIssuerBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustJwtGrantIssuerBody(expiresAt time.Time, issuer string, jwk JSONWebKey, scope []string) *TrustJwtGrantIssuerBody {
	this := TrustJwtGrantIssuerBody{}
	this.ExpiresAt = expiresAt
	this.Issuer = issuer
	this.Jwk = jwk
	this.Scope = scope
	return &this
}

// NewTrustJwtGrantIssuerBodyWithDefaults instantiates a new TrustJwtGrantIssuerBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustJwtGrantIssuerBodyWithDefaults() *TrustJwtGrantIssuerBody {
	this := TrustJwtGrantIssuerBody{}
	return &this
}

// GetAllowAnySubject returns the AllowAnySubject field value if set, zero value otherwise.
func (o *TrustJwtGrantIssuerBody) GetAllowAnySubject() bool {
	if o == nil || o.AllowAnySubject == nil {
		var ret bool
		return ret
	}
	return *o.AllowAnySubject
}

// GetAllowAnySubjectOk returns a tuple with the AllowAnySubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetAllowAnySubjectOk() (*bool, bool) {
	if o == nil || o.AllowAnySubject == nil {
		return nil, false
	}
	return o.AllowAnySubject, true
}

// HasAllowAnySubject returns a boolean if a field has been set.
func (o *TrustJwtGrantIssuerBody) HasAllowAnySubject() bool {
	if o != nil && o.AllowAnySubject != nil {
		return true
	}

	return false
}

// SetAllowAnySubject gets a reference to the given bool and assigns it to the AllowAnySubject field.
func (o *TrustJwtGrantIssuerBody) SetAllowAnySubject(v bool) {
	o.AllowAnySubject = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *TrustJwtGrantIssuerBody) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *TrustJwtGrantIssuerBody) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetIssuer returns the Issuer field value
func (o *TrustJwtGrantIssuerBody) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *TrustJwtGrantIssuerBody) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwk returns the Jwk field value
func (o *TrustJwtGrantIssuerBody) GetJwk() JSONWebKey {
	if o == nil {
		var ret JSONWebKey
		return ret
	}

	return o.Jwk
}

// GetJwkOk returns a tuple with the Jwk field value
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetJwkOk() (*JSONWebKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwk, true
}

// SetJwk sets field value
func (o *TrustJwtGrantIssuerBody) SetJwk(v JSONWebKey) {
	o.Jwk = v
}

// GetScope returns the Scope field value
func (o *TrustJwtGrantIssuerBody) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *TrustJwtGrantIssuerBody) SetScope(v []string) {
	o.Scope = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TrustJwtGrantIssuerBody) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustJwtGrantIssuerBody) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TrustJwtGrantIssuerBody) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TrustJwtGrantIssuerBody) SetSubject(v string) {
	o.Subject = &v
}

func (o TrustJwtGrantIssuerBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowAnySubject != nil {
		toSerialize["allow_any_subject"] = o.AllowAnySubject
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["jwk"] = o.Jwk
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableTrustJwtGrantIssuerBody struct {
	value *TrustJwtGrantIssuerBody
	isSet bool
}

func (v NullableTrustJwtGrantIssuerBody) Get() *TrustJwtGrantIssuerBody {
	return v.value
}

func (v *NullableTrustJwtGrantIssuerBody) Set(val *TrustJwtGrantIssuerBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustJwtGrantIssuerBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustJwtGrantIssuerBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustJwtGrantIssuerBody(val *TrustJwtGrantIssuerBody) *NullableTrustJwtGrantIssuerBody {
	return &NullableTrustJwtGrantIssuerBody{value: val, isSet: true}
}

func (v NullableTrustJwtGrantIssuerBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustJwtGrantIssuerBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


