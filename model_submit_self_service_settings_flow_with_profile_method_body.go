/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.133
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceSettingsFlowWithProfileMethodBody nolint:deadcode,unused
type SubmitSelfServiceSettingsFlowWithProfileMethodBody struct {
	// The Anti-CSRF Token  This token is only required when performing browser flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to profile when trying to update a profile.
	Method string `json:"method"`
	// Traits contains all of the identity's traits.
	Traits map[string]interface{} `json:"traits"`
}

// NewSubmitSelfServiceSettingsFlowWithProfileMethodBody instantiates a new SubmitSelfServiceSettingsFlowWithProfileMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithProfileMethodBody(method string, traits map[string]interface{}) *SubmitSelfServiceSettingsFlowWithProfileMethodBody {
	this := SubmitSelfServiceSettingsFlowWithProfileMethodBody{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithProfileMethodBodyWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithProfileMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithProfileMethodBodyWithDefaults() *SubmitSelfServiceSettingsFlowWithProfileMethodBody {
	this := SubmitSelfServiceSettingsFlowWithProfileMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetTraits returns the Traits field value
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *SubmitSelfServiceSettingsFlowWithProfileMethodBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o SubmitSelfServiceSettingsFlowWithProfileMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody struct {
	value *SubmitSelfServiceSettingsFlowWithProfileMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) Get() *SubmitSelfServiceSettingsFlowWithProfileMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) Set(val *SubmitSelfServiceSettingsFlowWithProfileMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithProfileMethodBody(val *SubmitSelfServiceSettingsFlowWithProfileMethodBody) *NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody {
	return &NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithProfileMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


