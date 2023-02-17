/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.15
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateVerificationFlowWithCodeMethodBody struct for UpdateVerificationFlowWithCodeMethodBody
type UpdateVerificationFlowWithCodeMethodBody struct {
	// The verification code
	Code *string `json:"code,omitempty"`
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
	Email *string `json:"email,omitempty"`
	// The id of the flow
	Flow *string `json:"flow,omitempty"`
	// Method is the recovery method
	Method *string `json:"method,omitempty"`
}

// NewUpdateVerificationFlowWithCodeMethodBody instantiates a new UpdateVerificationFlowWithCodeMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVerificationFlowWithCodeMethodBody() *UpdateVerificationFlowWithCodeMethodBody {
	this := UpdateVerificationFlowWithCodeMethodBody{}
	return &this
}

// NewUpdateVerificationFlowWithCodeMethodBodyWithDefaults instantiates a new UpdateVerificationFlowWithCodeMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVerificationFlowWithCodeMethodBodyWithDefaults() *UpdateVerificationFlowWithCodeMethodBody {
	this := UpdateVerificationFlowWithCodeMethodBody{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateVerificationFlowWithCodeMethodBody) SetCode(v string) {
	o.Code = &v
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateVerificationFlowWithCodeMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateVerificationFlowWithCodeMethodBody) SetEmail(v string) {
	o.Email = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetFlow() string {
	if o == nil || o.Flow == nil {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetFlowOk() (*string, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *UpdateVerificationFlowWithCodeMethodBody) SetFlow(v string) {
	o.Flow = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UpdateVerificationFlowWithCodeMethodBody) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *UpdateVerificationFlowWithCodeMethodBody) SetMethod(v string) {
	o.Method = &v
}

func (o UpdateVerificationFlowWithCodeMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVerificationFlowWithCodeMethodBody struct {
	value *UpdateVerificationFlowWithCodeMethodBody
	isSet bool
}

func (v NullableUpdateVerificationFlowWithCodeMethodBody) Get() *UpdateVerificationFlowWithCodeMethodBody {
	return v.value
}

func (v *NullableUpdateVerificationFlowWithCodeMethodBody) Set(val *UpdateVerificationFlowWithCodeMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVerificationFlowWithCodeMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVerificationFlowWithCodeMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVerificationFlowWithCodeMethodBody(val *UpdateVerificationFlowWithCodeMethodBody) *NullableUpdateVerificationFlowWithCodeMethodBody {
	return &NullableUpdateVerificationFlowWithCodeMethodBody{value: val, isSet: true}
}

func (v NullableUpdateVerificationFlowWithCodeMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVerificationFlowWithCodeMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


