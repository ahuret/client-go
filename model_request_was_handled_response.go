/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: latest
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RequestWasHandledResponse struct for RequestWasHandledResponse
type RequestWasHandledResponse struct {
	// Original request URL to which you should redirect the user if request was already handled.
	RedirectTo string `json:"redirect_to"`
}

// NewRequestWasHandledResponse instantiates a new RequestWasHandledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestWasHandledResponse(redirectTo string) *RequestWasHandledResponse {
	this := RequestWasHandledResponse{}
	this.RedirectTo = redirectTo
	return &this
}

// NewRequestWasHandledResponseWithDefaults instantiates a new RequestWasHandledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWasHandledResponseWithDefaults() *RequestWasHandledResponse {
	this := RequestWasHandledResponse{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *RequestWasHandledResponse) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *RequestWasHandledResponse) GetRedirectToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *RequestWasHandledResponse) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o RequestWasHandledResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}
	return json.Marshal(toSerialize)
}

type NullableRequestWasHandledResponse struct {
	value *RequestWasHandledResponse
	isSet bool
}

func (v NullableRequestWasHandledResponse) Get() *RequestWasHandledResponse {
	return v.value
}

func (v *NullableRequestWasHandledResponse) Set(val *RequestWasHandledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestWasHandledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestWasHandledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestWasHandledResponse(val *RequestWasHandledResponse) *NullableRequestWasHandledResponse {
	return &NullableRequestWasHandledResponse{value: val, isSet: true}
}

func (v NullableRequestWasHandledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestWasHandledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


