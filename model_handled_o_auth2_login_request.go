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

// HandledOAuth2LoginRequest struct for HandledOAuth2LoginRequest
type HandledOAuth2LoginRequest struct {
	// Original request URL to which you should redirect the user if request was already handled.
	RedirectTo string `json:"redirect_to"`
}

// NewHandledOAuth2LoginRequest instantiates a new HandledOAuth2LoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandledOAuth2LoginRequest(redirectTo string) *HandledOAuth2LoginRequest {
	this := HandledOAuth2LoginRequest{}
	this.RedirectTo = redirectTo
	return &this
}

// NewHandledOAuth2LoginRequestWithDefaults instantiates a new HandledOAuth2LoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandledOAuth2LoginRequestWithDefaults() *HandledOAuth2LoginRequest {
	this := HandledOAuth2LoginRequest{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *HandledOAuth2LoginRequest) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *HandledOAuth2LoginRequest) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *HandledOAuth2LoginRequest) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o HandledOAuth2LoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableHandledOAuth2LoginRequest struct {
	value *HandledOAuth2LoginRequest
	isSet bool
}

func (v NullableHandledOAuth2LoginRequest) Get() *HandledOAuth2LoginRequest {
	return v.value
}

func (v *NullableHandledOAuth2LoginRequest) Set(val *HandledOAuth2LoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandledOAuth2LoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandledOAuth2LoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandledOAuth2LoginRequest(val *HandledOAuth2LoginRequest) *NullableHandledOAuth2LoginRequest {
	return &NullableHandledOAuth2LoginRequest{value: val, isSet: true}
}

func (v NullableHandledOAuth2LoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandledOAuth2LoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


