/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.24
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RefreshTokenHookResponse struct for RefreshTokenHookResponse
type RefreshTokenHookResponse struct {
	Session *AcceptOAuth2ConsentRequestSession `json:"session,omitempty"`
}

// NewRefreshTokenHookResponse instantiates a new RefreshTokenHookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenHookResponse() *RefreshTokenHookResponse {
	this := RefreshTokenHookResponse{}
	return &this
}

// NewRefreshTokenHookResponseWithDefaults instantiates a new RefreshTokenHookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenHookResponseWithDefaults() *RefreshTokenHookResponse {
	this := RefreshTokenHookResponse{}
	return &this
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *RefreshTokenHookResponse) GetSession() AcceptOAuth2ConsentRequestSession {
	if o == nil || o.Session == nil {
		var ret AcceptOAuth2ConsentRequestSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenHookResponse) GetSessionOk() (*AcceptOAuth2ConsentRequestSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *RefreshTokenHookResponse) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given AcceptOAuth2ConsentRequestSession and assigns it to the Session field.
func (o *RefreshTokenHookResponse) SetSession(v AcceptOAuth2ConsentRequestSession) {
	o.Session = &v
}

func (o RefreshTokenHookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshTokenHookResponse struct {
	value *RefreshTokenHookResponse
	isSet bool
}

func (v NullableRefreshTokenHookResponse) Get() *RefreshTokenHookResponse {
	return v.value
}

func (v *NullableRefreshTokenHookResponse) Set(val *RefreshTokenHookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenHookResponse(val *RefreshTokenHookResponse) *NullableRefreshTokenHookResponse {
	return &NullableRefreshTokenHookResponse{value: val, isSet: true}
}

func (v NullableRefreshTokenHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


