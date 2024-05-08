/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateLoginFlowWithPasskeyMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLoginFlowWithPasskeyMethod{}

// UpdateLoginFlowWithPasskeyMethod Update Login Flow with Passkey Method
type UpdateLoginFlowWithPasskeyMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method should be set to \"passkey\" when logging in using the Passkey strategy.
	Method string `json:"method"`
	// Login a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
	PasskeyLogin *string `json:"passkey_login,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLoginFlowWithPasskeyMethod UpdateLoginFlowWithPasskeyMethod

// NewUpdateLoginFlowWithPasskeyMethod instantiates a new UpdateLoginFlowWithPasskeyMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoginFlowWithPasskeyMethod(method string) *UpdateLoginFlowWithPasskeyMethod {
	this := UpdateLoginFlowWithPasskeyMethod{}
	this.Method = method
	return &this
}

// NewUpdateLoginFlowWithPasskeyMethodWithDefaults instantiates a new UpdateLoginFlowWithPasskeyMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoginFlowWithPasskeyMethodWithDefaults() *UpdateLoginFlowWithPasskeyMethod {
	this := UpdateLoginFlowWithPasskeyMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithPasskeyMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithPasskeyMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithPasskeyMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateLoginFlowWithPasskeyMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateLoginFlowWithPasskeyMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithPasskeyMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateLoginFlowWithPasskeyMethod) SetMethod(v string) {
	o.Method = v
}

// GetPasskeyLogin returns the PasskeyLogin field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithPasskeyMethod) GetPasskeyLogin() string {
	if o == nil || IsNil(o.PasskeyLogin) {
		var ret string
		return ret
	}
	return *o.PasskeyLogin
}

// GetPasskeyLoginOk returns a tuple with the PasskeyLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithPasskeyMethod) GetPasskeyLoginOk() (*string, bool) {
	if o == nil || IsNil(o.PasskeyLogin) {
		return nil, false
	}
	return o.PasskeyLogin, true
}

// HasPasskeyLogin returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithPasskeyMethod) HasPasskeyLogin() bool {
	if o != nil && !IsNil(o.PasskeyLogin) {
		return true
	}

	return false
}

// SetPasskeyLogin gets a reference to the given string and assigns it to the PasskeyLogin field.
func (o *UpdateLoginFlowWithPasskeyMethod) SetPasskeyLogin(v string) {
	o.PasskeyLogin = &v
}

func (o UpdateLoginFlowWithPasskeyMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLoginFlowWithPasskeyMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.PasskeyLogin) {
		toSerialize["passkey_login"] = o.PasskeyLogin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLoginFlowWithPasskeyMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateLoginFlowWithPasskeyMethod := _UpdateLoginFlowWithPasskeyMethod{}

	err = json.Unmarshal(data, &varUpdateLoginFlowWithPasskeyMethod)

	if err != nil {
		return err
	}

	*o = UpdateLoginFlowWithPasskeyMethod(varUpdateLoginFlowWithPasskeyMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "passkey_login")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLoginFlowWithPasskeyMethod struct {
	value *UpdateLoginFlowWithPasskeyMethod
	isSet bool
}

func (v NullableUpdateLoginFlowWithPasskeyMethod) Get() *UpdateLoginFlowWithPasskeyMethod {
	return v.value
}

func (v *NullableUpdateLoginFlowWithPasskeyMethod) Set(val *UpdateLoginFlowWithPasskeyMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowWithPasskeyMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowWithPasskeyMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowWithPasskeyMethod(val *UpdateLoginFlowWithPasskeyMethod) *NullableUpdateLoginFlowWithPasskeyMethod {
	return &NullableUpdateLoginFlowWithPasskeyMethod{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowWithPasskeyMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowWithPasskeyMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


