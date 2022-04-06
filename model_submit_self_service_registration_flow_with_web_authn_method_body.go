/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.157
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody struct for SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
type SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"webauthn\" when trying to add, update, or remove a webAuthn pairing.
	Method string `json:"method"`
	// The identity's traits
	Traits map[string]interface{} `json:"traits"`
	// Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here.
	WebauthnRegister *string `json:"webauthn_register,omitempty"`
	// Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added.
	WebauthnRegisterDisplayname *string `json:"webauthn_register_displayname,omitempty"`
}

// NewSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody instantiates a new SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody(method string, traits map[string]interface{}) *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
	this := SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody{}
	this.Method = method
	this.Traits = traits
	return &this
}

// NewSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBodyWithDefaults instantiates a new SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBodyWithDefaults() *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
	this := SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetTraits returns the Traits field value
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetWebauthnRegister returns the WebauthnRegister field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetWebauthnRegister() string {
	if o == nil || o.WebauthnRegister == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegister
}

// GetWebauthnRegisterOk returns a tuple with the WebauthnRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetWebauthnRegisterOk() (*string, bool) {
	if o == nil || o.WebauthnRegister == nil {
		return nil, false
	}
	return o.WebauthnRegister, true
}

// HasWebauthnRegister returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) HasWebauthnRegister() bool {
	if o != nil && o.WebauthnRegister != nil {
		return true
	}

	return false
}

// SetWebauthnRegister gets a reference to the given string and assigns it to the WebauthnRegister field.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SetWebauthnRegister(v string) {
	o.WebauthnRegister = &v
}

// GetWebauthnRegisterDisplayname returns the WebauthnRegisterDisplayname field value if set, zero value otherwise.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetWebauthnRegisterDisplayname() string {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		var ret string
		return ret
	}
	return *o.WebauthnRegisterDisplayname
}

// GetWebauthnRegisterDisplaynameOk returns a tuple with the WebauthnRegisterDisplayname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) GetWebauthnRegisterDisplaynameOk() (*string, bool) {
	if o == nil || o.WebauthnRegisterDisplayname == nil {
		return nil, false
	}
	return o.WebauthnRegisterDisplayname, true
}

// HasWebauthnRegisterDisplayname returns a boolean if a field has been set.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) HasWebauthnRegisterDisplayname() bool {
	if o != nil && o.WebauthnRegisterDisplayname != nil {
		return true
	}

	return false
}

// SetWebauthnRegisterDisplayname gets a reference to the given string and assigns it to the WebauthnRegisterDisplayname field.
func (o *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SetWebauthnRegisterDisplayname(v string) {
	o.WebauthnRegisterDisplayname = &v
}

func (o SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) MarshalJSON() ([]byte, error) {
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
	if o.WebauthnRegister != nil {
		toSerialize["webauthn_register"] = o.WebauthnRegister
	}
	if o.WebauthnRegisterDisplayname != nil {
		toSerialize["webauthn_register_displayname"] = o.WebauthnRegisterDisplayname
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody struct {
	value *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) Get() *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) Set(val *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody(val *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) *NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody {
	return &NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


