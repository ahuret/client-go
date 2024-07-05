/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateSettingsFlowWithTotpMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSettingsFlowWithTotpMethod{}

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
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
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
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
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
	if o == nil || IsNil(o.TotpCode) {
		var ret string
		return ret
	}
	return *o.TotpCode
}

// GetTotpCodeOk returns a tuple with the TotpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TotpCode) {
		return nil, false
	}
	return o.TotpCode, true
}

// HasTotpCode returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasTotpCode() bool {
	if o != nil && !IsNil(o.TotpCode) {
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
	if o == nil || IsNil(o.TotpUnlink) {
		var ret bool
		return ret
	}
	return *o.TotpUnlink
}

// GetTotpUnlinkOk returns a tuple with the TotpUnlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetTotpUnlinkOk() (*bool, bool) {
	if o == nil || IsNil(o.TotpUnlink) {
		return nil, false
	}
	return o.TotpUnlink, true
}

// HasTotpUnlink returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasTotpUnlink() bool {
	if o != nil && !IsNil(o.TotpUnlink) {
		return true
	}

	return false
}

// SetTotpUnlink gets a reference to the given bool and assigns it to the TotpUnlink field.
func (o *UpdateSettingsFlowWithTotpMethod) SetTotpUnlink(v bool) {
	o.TotpUnlink = &v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithTotpMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithTotpMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithTotpMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateSettingsFlowWithTotpMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateSettingsFlowWithTotpMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSettingsFlowWithTotpMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.TotpCode) {
		toSerialize["totp_code"] = o.TotpCode
	}
	if !IsNil(o.TotpUnlink) {
		toSerialize["totp_unlink"] = o.TotpUnlink
	}
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSettingsFlowWithTotpMethod) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateSettingsFlowWithTotpMethod := _UpdateSettingsFlowWithTotpMethod{}

	err = json.Unmarshal(data, &varUpdateSettingsFlowWithTotpMethod)

	if err != nil {
		return err
	}

	*o = UpdateSettingsFlowWithTotpMethod(varUpdateSettingsFlowWithTotpMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "totp_code")
		delete(additionalProperties, "totp_unlink")
		delete(additionalProperties, "transient_payload")
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


