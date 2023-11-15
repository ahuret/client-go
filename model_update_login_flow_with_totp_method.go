/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateLoginFlowWithTotpMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLoginFlowWithTotpMethod{}

// UpdateLoginFlowWithTotpMethod Update Login Flow with TOTP Method
type UpdateLoginFlowWithTotpMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method should be set to \"totp\" when logging in using the TOTP strategy.
	Method string `json:"method"`
	// The TOTP code.
	TotpCode string `json:"totp_code"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLoginFlowWithTotpMethod UpdateLoginFlowWithTotpMethod

// NewUpdateLoginFlowWithTotpMethod instantiates a new UpdateLoginFlowWithTotpMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoginFlowWithTotpMethod(method string, totpCode string) *UpdateLoginFlowWithTotpMethod {
	this := UpdateLoginFlowWithTotpMethod{}
	this.Method = method
	this.TotpCode = totpCode
	return &this
}

// NewUpdateLoginFlowWithTotpMethodWithDefaults instantiates a new UpdateLoginFlowWithTotpMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoginFlowWithTotpMethodWithDefaults() *UpdateLoginFlowWithTotpMethod {
	this := UpdateLoginFlowWithTotpMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithTotpMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithTotpMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithTotpMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateLoginFlowWithTotpMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateLoginFlowWithTotpMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithTotpMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateLoginFlowWithTotpMethod) SetMethod(v string) {
	o.Method = v
}

// GetTotpCode returns the TotpCode field value
func (o *UpdateLoginFlowWithTotpMethod) GetTotpCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotpCode
}

// GetTotpCodeOk returns a tuple with the TotpCode field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithTotpMethod) GetTotpCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotpCode, true
}

// SetTotpCode sets field value
func (o *UpdateLoginFlowWithTotpMethod) SetTotpCode(v string) {
	o.TotpCode = v
}

func (o UpdateLoginFlowWithTotpMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLoginFlowWithTotpMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	toSerialize["totp_code"] = o.TotpCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLoginFlowWithTotpMethod) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"totp_code",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateLoginFlowWithTotpMethod := _UpdateLoginFlowWithTotpMethod{}

	err = json.Unmarshal(bytes, &varUpdateLoginFlowWithTotpMethod)

	if err != nil {
		return err
	}

	*o = UpdateLoginFlowWithTotpMethod(varUpdateLoginFlowWithTotpMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "totp_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLoginFlowWithTotpMethod struct {
	value *UpdateLoginFlowWithTotpMethod
	isSet bool
}

func (v NullableUpdateLoginFlowWithTotpMethod) Get() *UpdateLoginFlowWithTotpMethod {
	return v.value
}

func (v *NullableUpdateLoginFlowWithTotpMethod) Set(val *UpdateLoginFlowWithTotpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowWithTotpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowWithTotpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowWithTotpMethod(val *UpdateLoginFlowWithTotpMethod) *NullableUpdateLoginFlowWithTotpMethod {
	return &NullableUpdateLoginFlowWithTotpMethod{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowWithTotpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowWithTotpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


