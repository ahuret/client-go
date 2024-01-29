/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.5.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the TrustOAuth2JwtGrantIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustOAuth2JwtGrantIssuer{}

// TrustOAuth2JwtGrantIssuer Trust OAuth2 JWT Bearer Grant Type Issuer Request Body
type TrustOAuth2JwtGrantIssuer struct {
	// The \"allow_any_subject\" indicates that the issuer is allowed to have any principal as the subject of the JWT.
	AllowAnySubject *bool `json:"allow_any_subject,omitempty"`
	// The \"expires_at\" indicates, when grant will expire, so we will reject assertion from \"issuer\" targeting \"subject\".
	ExpiresAt time.Time `json:"expires_at"`
	// The \"issuer\" identifies the principal that issued the JWT assertion (same as \"iss\" claim in JWT).
	Issuer string `json:"issuer"`
	Jwk JsonWebKey `json:"jwk"`
	// The \"scope\" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
	Scope []string `json:"scope"`
	// The \"subject\" identifies the principal that is the subject of the JWT.
	Subject *string `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrustOAuth2JwtGrantIssuer TrustOAuth2JwtGrantIssuer

// NewTrustOAuth2JwtGrantIssuer instantiates a new TrustOAuth2JwtGrantIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustOAuth2JwtGrantIssuer(expiresAt time.Time, issuer string, jwk JsonWebKey, scope []string) *TrustOAuth2JwtGrantIssuer {
	this := TrustOAuth2JwtGrantIssuer{}
	this.ExpiresAt = expiresAt
	this.Issuer = issuer
	this.Jwk = jwk
	this.Scope = scope
	return &this
}

// NewTrustOAuth2JwtGrantIssuerWithDefaults instantiates a new TrustOAuth2JwtGrantIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustOAuth2JwtGrantIssuerWithDefaults() *TrustOAuth2JwtGrantIssuer {
	this := TrustOAuth2JwtGrantIssuer{}
	return &this
}

// GetAllowAnySubject returns the AllowAnySubject field value if set, zero value otherwise.
func (o *TrustOAuth2JwtGrantIssuer) GetAllowAnySubject() bool {
	if o == nil || IsNil(o.AllowAnySubject) {
		var ret bool
		return ret
	}
	return *o.AllowAnySubject
}

// GetAllowAnySubjectOk returns a tuple with the AllowAnySubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetAllowAnySubjectOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAnySubject) {
		return nil, false
	}
	return o.AllowAnySubject, true
}

// HasAllowAnySubject returns a boolean if a field has been set.
func (o *TrustOAuth2JwtGrantIssuer) HasAllowAnySubject() bool {
	if o != nil && !IsNil(o.AllowAnySubject) {
		return true
	}

	return false
}

// SetAllowAnySubject gets a reference to the given bool and assigns it to the AllowAnySubject field.
func (o *TrustOAuth2JwtGrantIssuer) SetAllowAnySubject(v bool) {
	o.AllowAnySubject = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *TrustOAuth2JwtGrantIssuer) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *TrustOAuth2JwtGrantIssuer) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetIssuer returns the Issuer field value
func (o *TrustOAuth2JwtGrantIssuer) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *TrustOAuth2JwtGrantIssuer) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwk returns the Jwk field value
func (o *TrustOAuth2JwtGrantIssuer) GetJwk() JsonWebKey {
	if o == nil {
		var ret JsonWebKey
		return ret
	}

	return o.Jwk
}

// GetJwkOk returns a tuple with the Jwk field value
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetJwkOk() (*JsonWebKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwk, true
}

// SetJwk sets field value
func (o *TrustOAuth2JwtGrantIssuer) SetJwk(v JsonWebKey) {
	o.Jwk = v
}

// GetScope returns the Scope field value
func (o *TrustOAuth2JwtGrantIssuer) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *TrustOAuth2JwtGrantIssuer) SetScope(v []string) {
	o.Scope = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TrustOAuth2JwtGrantIssuer) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustOAuth2JwtGrantIssuer) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TrustOAuth2JwtGrantIssuer) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TrustOAuth2JwtGrantIssuer) SetSubject(v string) {
	o.Subject = &v
}

func (o TrustOAuth2JwtGrantIssuer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustOAuth2JwtGrantIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowAnySubject) {
		toSerialize["allow_any_subject"] = o.AllowAnySubject
	}
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["issuer"] = o.Issuer
	toSerialize["jwk"] = o.Jwk
	toSerialize["scope"] = o.Scope
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TrustOAuth2JwtGrantIssuer) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expires_at",
		"issuer",
		"jwk",
		"scope",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTrustOAuth2JwtGrantIssuer := _TrustOAuth2JwtGrantIssuer{}

	err = json.Unmarshal(bytes, &varTrustOAuth2JwtGrantIssuer)

	if err != nil {
		return err
	}

	*o = TrustOAuth2JwtGrantIssuer(varTrustOAuth2JwtGrantIssuer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_any_subject")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "jwk")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTrustOAuth2JwtGrantIssuer struct {
	value *TrustOAuth2JwtGrantIssuer
	isSet bool
}

func (v NullableTrustOAuth2JwtGrantIssuer) Get() *TrustOAuth2JwtGrantIssuer {
	return v.value
}

func (v *NullableTrustOAuth2JwtGrantIssuer) Set(val *TrustOAuth2JwtGrantIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustOAuth2JwtGrantIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustOAuth2JwtGrantIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustOAuth2JwtGrantIssuer(val *TrustOAuth2JwtGrantIssuer) *NullableTrustOAuth2JwtGrantIssuer {
	return &NullableTrustOAuth2JwtGrantIssuer{value: val, isSet: true}
}

func (v NullableTrustOAuth2JwtGrantIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustOAuth2JwtGrantIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


