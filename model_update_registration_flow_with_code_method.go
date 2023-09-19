/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.8
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateRegistrationFlowWithCodeMethod Update Registration Flow with Code Method
type UpdateRegistrationFlowWithCodeMethod struct {
	// The OTP Code sent to the user
	Code *string `json:"code,omitempty"`
	// The CSRF Token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method to use  This field must be set to `code` when using the code method.
	Method string `json:"method"`
	// Resend restarts the flow with a new code
	Resend *string `json:"resend,omitempty"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRegistrationFlowWithCodeMethod UpdateRegistrationFlowWithCodeMethod

// NewUpdateRegistrationFlowWithCodeMethod instantiates a new UpdateRegistrationFlowWithCodeMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegistrationFlowWithCodeMethod(method string, traits map[string]interface{}) *UpdateRegistrationFlowWithCodeMethod {
	this := UpdateRegistrationFlowWithCodeMethod{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewUpdateRegistrationFlowWithCodeMethodWithDefaults instantiates a new UpdateRegistrationFlowWithCodeMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegistrationFlowWithCodeMethodWithDefaults() *UpdateRegistrationFlowWithCodeMethod {
	this := UpdateRegistrationFlowWithCodeMethod{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithCodeMethod) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateRegistrationFlowWithCodeMethod) SetCode(v string) {
	o.Code = &v
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithCodeMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRegistrationFlowWithCodeMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateRegistrationFlowWithCodeMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRegistrationFlowWithCodeMethod) SetMethod(v string) {
	o.Method = v
}

// GetResend returns the Resend field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithCodeMethod) GetResend() string {
	if o == nil || o.Resend == nil {
		var ret string
		return ret
	}
	return *o.Resend
}

// GetResendOk returns a tuple with the Resend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetResendOk() (*string, bool) {
	if o == nil || o.Resend == nil {
		return nil, false
	}
	return o.Resend, true
}

// HasResend returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) HasResend() bool {
	if o != nil && o.Resend != nil {
		return true
	}

	return false
}

// SetResend gets a reference to the given string and assigns it to the Resend field.
func (o *UpdateRegistrationFlowWithCodeMethod) SetResend(v string) {
	o.Resend = &v
}

// GetTraits returns the Traits field value
func (o *UpdateRegistrationFlowWithCodeMethod) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *UpdateRegistrationFlowWithCodeMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRegistrationFlowWithCodeMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || o.TransientPayload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || o.TransientPayload == nil {
		return nil, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRegistrationFlowWithCodeMethod) HasTransientPayload() bool {
	if o != nil && o.TransientPayload != nil {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRegistrationFlowWithCodeMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateRegistrationFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Resend != nil {
		toSerialize["resend"] = o.Resend
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	if o.TransientPayload != nil {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateRegistrationFlowWithCodeMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateRegistrationFlowWithCodeMethod := _UpdateRegistrationFlowWithCodeMethod{}

	if err = json.Unmarshal(bytes, &varUpdateRegistrationFlowWithCodeMethod); err == nil {
		*o = UpdateRegistrationFlowWithCodeMethod(varUpdateRegistrationFlowWithCodeMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "resend")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRegistrationFlowWithCodeMethod struct {
	value *UpdateRegistrationFlowWithCodeMethod
	isSet bool
}

func (v NullableUpdateRegistrationFlowWithCodeMethod) Get() *UpdateRegistrationFlowWithCodeMethod {
	return v.value
}

func (v *NullableUpdateRegistrationFlowWithCodeMethod) Set(val *UpdateRegistrationFlowWithCodeMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowWithCodeMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowWithCodeMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowWithCodeMethod(val *UpdateRegistrationFlowWithCodeMethod) *NullableUpdateRegistrationFlowWithCodeMethod {
	return &NullableUpdateRegistrationFlowWithCodeMethod{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowWithCodeMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


