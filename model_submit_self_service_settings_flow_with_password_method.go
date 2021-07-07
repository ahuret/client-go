/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.13
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceSettingsFlowWithPasswordMethod struct for SubmitSelfServiceSettingsFlowWithPasswordMethod
type SubmitSelfServiceSettingsFlowWithPasswordMethod struct {
	// CSRFToken is the anti-CSRF token  type: string
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to password when trying to update a password.  type: string
	Method *string `json:"method,omitempty"`
	// Password is the updated password  type: string
	Password string `json:"password"`
}

// NewSubmitSelfServiceSettingsFlowWithPasswordMethod instantiates a new SubmitSelfServiceSettingsFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithPasswordMethod(password string) *SubmitSelfServiceSettingsFlowWithPasswordMethod {
	this := SubmitSelfServiceSettingsFlowWithPasswordMethod{}
	this.Password = password
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithPasswordMethodWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithPasswordMethodWithDefaults() *SubmitSelfServiceSettingsFlowWithPasswordMethod {
	this := SubmitSelfServiceSettingsFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = &v
}

// GetPassword returns the Password field value
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SubmitSelfServiceSettingsFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = v
}

func (o SubmitSelfServiceSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithPasswordMethod struct {
	value *SubmitSelfServiceSettingsFlowWithPasswordMethod
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) Get() *SubmitSelfServiceSettingsFlowWithPasswordMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) Set(val *SubmitSelfServiceSettingsFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithPasswordMethod(val *SubmitSelfServiceSettingsFlowWithPasswordMethod) *NullableSubmitSelfServiceSettingsFlowWithPasswordMethod {
	return &NullableSubmitSelfServiceSettingsFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


