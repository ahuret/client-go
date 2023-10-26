/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.14
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateLoginFlowWithCodeMethod Update Login flow using the code method
type UpdateLoginFlowWithCodeMethod struct {
	// Code is the 6 digits code sent to the user
	Code *string `json:"code,omitempty"`
	// CSRFToken is the anti-CSRF token
	CsrfToken string `json:"csrf_token"`
	// Identifier is the code identifier The identifier requires that the user has already completed the registration or settings with code flow.
	Identifier *string `json:"identifier,omitempty"`
	// Method should be set to \"code\" when logging in using the code strategy.
	Method string `json:"method"`
	// Resend is set when the user wants to resend the code
	Resend *string `json:"resend,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLoginFlowWithCodeMethod UpdateLoginFlowWithCodeMethod

// NewUpdateLoginFlowWithCodeMethod instantiates a new UpdateLoginFlowWithCodeMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoginFlowWithCodeMethod(csrfToken string, method string) *UpdateLoginFlowWithCodeMethod {
	this := UpdateLoginFlowWithCodeMethod{}
	this.CsrfToken = csrfToken
	this.Method = method
	return &this
}

// NewUpdateLoginFlowWithCodeMethodWithDefaults instantiates a new UpdateLoginFlowWithCodeMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoginFlowWithCodeMethodWithDefaults() *UpdateLoginFlowWithCodeMethod {
	this := UpdateLoginFlowWithCodeMethod{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithCodeMethod) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithCodeMethod) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithCodeMethod) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateLoginFlowWithCodeMethod) SetCode(v string) {
	o.Code = &v
}

// GetCsrfToken returns the CsrfToken field value
func (o *UpdateLoginFlowWithCodeMethod) GetCsrfToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithCodeMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsrfToken, true
}

// SetCsrfToken sets field value
func (o *UpdateLoginFlowWithCodeMethod) SetCsrfToken(v string) {
	o.CsrfToken = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithCodeMethod) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithCodeMethod) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithCodeMethod) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *UpdateLoginFlowWithCodeMethod) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMethod returns the Method field value
func (o *UpdateLoginFlowWithCodeMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithCodeMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateLoginFlowWithCodeMethod) SetMethod(v string) {
	o.Method = v
}

// GetResend returns the Resend field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithCodeMethod) GetResend() string {
	if o == nil || o.Resend == nil {
		var ret string
		return ret
	}
	return *o.Resend
}

// GetResendOk returns a tuple with the Resend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithCodeMethod) GetResendOk() (*string, bool) {
	if o == nil || o.Resend == nil {
		return nil, false
	}
	return o.Resend, true
}

// HasResend returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithCodeMethod) HasResend() bool {
	if o != nil && o.Resend != nil {
		return true
	}

	return false
}

// SetResend gets a reference to the given string and assigns it to the Resend field.
func (o *UpdateLoginFlowWithCodeMethod) SetResend(v string) {
	o.Resend = &v
}

func (o UpdateLoginFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Resend != nil {
		toSerialize["resend"] = o.Resend
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateLoginFlowWithCodeMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateLoginFlowWithCodeMethod := _UpdateLoginFlowWithCodeMethod{}

	if err = json.Unmarshal(bytes, &varUpdateLoginFlowWithCodeMethod); err == nil {
		*o = UpdateLoginFlowWithCodeMethod(varUpdateLoginFlowWithCodeMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "method")
		delete(additionalProperties, "resend")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLoginFlowWithCodeMethod struct {
	value *UpdateLoginFlowWithCodeMethod
	isSet bool
}

func (v NullableUpdateLoginFlowWithCodeMethod) Get() *UpdateLoginFlowWithCodeMethod {
	return v.value
}

func (v *NullableUpdateLoginFlowWithCodeMethod) Set(val *UpdateLoginFlowWithCodeMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowWithCodeMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowWithCodeMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowWithCodeMethod(val *UpdateLoginFlowWithCodeMethod) *NullableUpdateLoginFlowWithCodeMethod {
	return &NullableUpdateLoginFlowWithCodeMethod{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowWithCodeMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowWithCodeMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


