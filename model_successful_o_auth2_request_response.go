/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.14
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SuccessfulOAuth2RequestResponse struct for SuccessfulOAuth2RequestResponse
type SuccessfulOAuth2RequestResponse struct {
	// RedirectURL is the URL which you should redirect the user to once the authentication process is completed.
	RedirectTo string `json:"redirect_to"`
}

// NewSuccessfulOAuth2RequestResponse instantiates a new SuccessfulOAuth2RequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulOAuth2RequestResponse(redirectTo string) *SuccessfulOAuth2RequestResponse {
	this := SuccessfulOAuth2RequestResponse{}
	this.RedirectTo = redirectTo
	return &this
}

// NewSuccessfulOAuth2RequestResponseWithDefaults instantiates a new SuccessfulOAuth2RequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulOAuth2RequestResponseWithDefaults() *SuccessfulOAuth2RequestResponse {
	this := SuccessfulOAuth2RequestResponse{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *SuccessfulOAuth2RequestResponse) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *SuccessfulOAuth2RequestResponse) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *SuccessfulOAuth2RequestResponse) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o SuccessfulOAuth2RequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessfulOAuth2RequestResponse struct {
	value *SuccessfulOAuth2RequestResponse
	isSet bool
}

func (v NullableSuccessfulOAuth2RequestResponse) Get() *SuccessfulOAuth2RequestResponse {
	return v.value
}

func (v *NullableSuccessfulOAuth2RequestResponse) Set(val *SuccessfulOAuth2RequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulOAuth2RequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulOAuth2RequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulOAuth2RequestResponse(val *SuccessfulOAuth2RequestResponse) *NullableSuccessfulOAuth2RequestResponse {
	return &NullableSuccessfulOAuth2RequestResponse{value: val, isSet: true}
}

func (v NullableSuccessfulOAuth2RequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulOAuth2RequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


