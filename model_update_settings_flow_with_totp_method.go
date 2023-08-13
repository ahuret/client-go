/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.45
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateSettingsFlowWithTotpMethod Update Settings Flow with TOTP Method
type UpdateSettingsFlowWithTotpMethod struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"totp\" when trying to add, update, or remove a totp pairing.
	Method string `json:"method"`
	// ValidationTOTP must contain a valid TOTP based on the
	TotpCode *string `json:"totp_code,omitempty"`
	// UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.
	TotpUnlink *bool `json:"totp_unlink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSettingsFlowWithTotpMethod UpdateSettingsFlowWithTotpMethod

// NewUpdateSettingsFlowWithTotpMethod instantiates a new UpdateSettingsFlowWithTotpMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingsFlowWithTotpMethod(method string) *UpdateSettingsFlowWithTotpMethod {
	this := UpdateSettingsFlowWithTotpMethod{}
	this.Method = method
	return &this
}

// NewUpdateSettingsFlowWithTotpMethodWithDefaults instantiates a new UpdateSettingsFlowWithTotpMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingsFlowWithTotpMethodWithDefaults() *UpdateSettingsFlowWithTotpMethod {
	this := UpdateSettingsFlowWithTotpMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithTotpMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateSettingsFlowWithTotpMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateSettingsFlowWithTotpMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateSettingsFlowWithTotpMethod) SetMethod(v string) {
	o.Method = v
}

// GetTotpCode returns the TotpCode field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpCode() string {
	if o == nil || o.TotpCode == nil {
		var ret string
		return ret
	}
	return *o.TotpCode
}

// GetTotpCodeOk returns a tuple with the TotpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpCodeOk() (*string, bool) {
	if o == nil || o.TotpCode == nil {
		return nil, false
	}
	return o.TotpCode, true
}

// HasTotpCode returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasTotpCode() bool {
	if o != nil && o.TotpCode != nil {
		return true
	}

	return false
}

// SetTotpCode gets a reference to the given string and assigns it to the TotpCode field.
func (o *UpdateSettingsFlowWithTotpMethod) SetTotpCode(v string) {
	o.TotpCode = &v
}

// GetTotpUnlink returns the TotpUnlink field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpUnlink() bool {
	if o == nil || o.TotpUnlink == nil {
		var ret bool
		return ret
	}
	return *o.TotpUnlink
}

// GetTotpUnlinkOk returns a tuple with the TotpUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpUnlinkOk() (*bool, bool) {
	if o == nil || o.TotpUnlink == nil {
		return nil, false
	}
	return o.TotpUnlink, true
}

// HasTotpUnlink returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasTotpUnlink() bool {
	if o != nil && o.TotpUnlink != nil {
		return true
	}

	return false
}

// SetTotpUnlink gets a reference to the given bool and assigns it to the TotpUnlink field.
func (o *UpdateSettingsFlowWithTotpMethod) SetTotpUnlink(v bool) {
	o.TotpUnlink = &v
}

func (o UpdateSettingsFlowWithTotpMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.TotpCode != nil {
		toSerialize["totp_code"] = o.TotpCode
	}
	if o.TotpUnlink != nil {
		toSerialize["totp_unlink"] = o.TotpUnlink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateSettingsFlowWithTotpMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateSettingsFlowWithTotpMethod := _UpdateSettingsFlowWithTotpMethod{}

	if err = json.Unmarshal(bytes, &varUpdateSettingsFlowWithTotpMethod); err == nil {
		*o = UpdateSettingsFlowWithTotpMethod(varUpdateSettingsFlowWithTotpMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "totp_code")
		delete(additionalProperties, "totp_unlink")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSettingsFlowWithTotpMethod struct {
	value *UpdateSettingsFlowWithTotpMethod
	isSet bool
}

func (v NullableUpdateSettingsFlowWithTotpMethod) Get() *UpdateSettingsFlowWithTotpMethod {
	return v.value
}

func (v *NullableUpdateSettingsFlowWithTotpMethod) Set(val *UpdateSettingsFlowWithTotpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingsFlowWithTotpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingsFlowWithTotpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingsFlowWithTotpMethod(val *UpdateSettingsFlowWithTotpMethod) *NullableUpdateSettingsFlowWithTotpMethod {
	return &NullableUpdateSettingsFlowWithTotpMethod{value: val, isSet: true}
}

func (v NullableUpdateSettingsFlowWithTotpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingsFlowWithTotpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


