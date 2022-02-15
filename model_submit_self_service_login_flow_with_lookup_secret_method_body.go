/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.94
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceLoginFlowWithLookupSecretMethodBody struct for SubmitSelfServiceLoginFlowWithLookupSecretMethodBody
type SubmitSelfServiceLoginFlowWithLookupSecretMethodBody struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// The lookup secret.
	LookupSecret string `json:"lookup_secret"`
	// Method should be set to \"lookup_secret\" when logging in using the lookup_secret strategy.
	Method string `json:"method"`
}

// NewSubmitSelfServiceLoginFlowWithLookupSecretMethodBody instantiates a new SubmitSelfServiceLoginFlowWithLookupSecretMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceLoginFlowWithLookupSecretMethodBody(lookupSecret string, method string) *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
	this := SubmitSelfServiceLoginFlowWithLookupSecretMethodBody{}
	this.LookupSecret = lookupSecret
	this.Method = method
	return &this
}

// NewSubmitSelfServiceLoginFlowWithLookupSecretMethodBodyWithDefaults instantiates a new SubmitSelfServiceLoginFlowWithLookupSecretMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceLoginFlowWithLookupSecretMethodBodyWithDefaults() *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
	this := SubmitSelfServiceLoginFlowWithLookupSecretMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetLookupSecret returns the LookupSecret field value
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetLookupSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupSecret
}

// GetLookupSecretOk returns a tuple with the LookupSecret field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetLookupSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LookupSecret, true
}

// SetLookupSecret sets field value
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) SetLookupSecret(v string) {
	o.LookupSecret = v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) SetMethod(v string) {
	o.Method = v
}

func (o SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["lookup_secret"] = o.LookupSecret
	}
	if true {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody struct {
	value *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) Get() *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) Set(val *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody(val *SubmitSelfServiceLoginFlowWithLookupSecretMethodBody) *NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody {
	return &NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceLoginFlowWithLookupSecretMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


