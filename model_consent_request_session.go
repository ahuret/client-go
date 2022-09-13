/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.40
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ConsentRequestSession struct for ConsentRequestSession
type ConsentRequestSession struct {
	// AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection. If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
	AccessToken interface{} `json:"access_token,omitempty"`
	// IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable by anyone that has access to the ID Challenge. Use with care!
	IdToken interface{} `json:"id_token,omitempty"`
}

// NewConsentRequestSession instantiates a new ConsentRequestSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentRequestSession() *ConsentRequestSession {
	this := ConsentRequestSession{}
	return &this
}

// NewConsentRequestSessionWithDefaults instantiates a new ConsentRequestSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentRequestSessionWithDefaults() *ConsentRequestSession {
	this := ConsentRequestSession{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsentRequestSession) GetAccessToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsentRequestSession) GetAccessTokenOk() (*interface{}, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *ConsentRequestSession) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given interface{} and assigns it to the AccessToken field.
func (o *ConsentRequestSession) SetAccessToken(v interface{}) {
	o.AccessToken = v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsentRequestSession) GetIdToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsentRequestSession) GetIdTokenOk() (*interface{}, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *ConsentRequestSession) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given interface{} and assigns it to the IdToken field.
func (o *ConsentRequestSession) SetIdToken(v interface{}) {
	o.IdToken = v
}

func (o ConsentRequestSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	return json.Marshal(toSerialize)
}

type NullableConsentRequestSession struct {
	value *ConsentRequestSession
	isSet bool
}

func (v NullableConsentRequestSession) Get() *ConsentRequestSession {
	return v.value
}

func (v *NullableConsentRequestSession) Set(val *ConsentRequestSession) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentRequestSession) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentRequestSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentRequestSession(val *ConsentRequestSession) *NullableConsentRequestSession {
	return &NullableConsentRequestSession{value: val, isSet: true}
}

func (v NullableConsentRequestSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentRequestSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


