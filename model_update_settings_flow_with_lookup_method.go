/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateSettingsFlowWithLookupMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSettingsFlowWithLookupMethod{}

// UpdateSettingsFlowWithLookupMethod Update Settings Flow with Lookup Method
type UpdateSettingsFlowWithLookupMethod struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// If set to true will save the regenerated lookup secrets
	LookupSecretConfirm *bool `json:"lookup_secret_confirm,omitempty"`
	// Disables this method if true.
	LookupSecretDisable *bool `json:"lookup_secret_disable,omitempty"`
	// If set to true will regenerate the lookup secrets
	LookupSecretRegenerate *bool `json:"lookup_secret_regenerate,omitempty"`
	// If set to true will reveal the lookup secrets
	LookupSecretReveal *bool `json:"lookup_secret_reveal,omitempty"`
	// Method  Should be set to \"lookup\" when trying to add, update, or remove a lookup pairing.
	Method string `json:"method"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSettingsFlowWithLookupMethod UpdateSettingsFlowWithLookupMethod

// NewUpdateSettingsFlowWithLookupMethod instantiates a new UpdateSettingsFlowWithLookupMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingsFlowWithLookupMethod(method string) *UpdateSettingsFlowWithLookupMethod {
	this := UpdateSettingsFlowWithLookupMethod{}
	this.Method = method
	return &this
}

// NewUpdateSettingsFlowWithLookupMethodWithDefaults instantiates a new UpdateSettingsFlowWithLookupMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingsFlowWithLookupMethodWithDefaults() *UpdateSettingsFlowWithLookupMethod {
	this := UpdateSettingsFlowWithLookupMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateSettingsFlowWithLookupMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetLookupSecretConfirm returns the LookupSecretConfirm field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretConfirm() bool {
	if o == nil || IsNil(o.LookupSecretConfirm) {
		var ret bool
		return ret
	}
	return *o.LookupSecretConfirm
}

// GetLookupSecretConfirmOk returns a tuple with the LookupSecretConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretConfirmOk() (*bool, bool) {
	if o == nil || IsNil(o.LookupSecretConfirm) {
		return nil, false
	}
	return o.LookupSecretConfirm, true
}

// HasLookupSecretConfirm returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasLookupSecretConfirm() bool {
	if o != nil && !IsNil(o.LookupSecretConfirm) {
		return true
	}

	return false
}

// SetLookupSecretConfirm gets a reference to the given bool and assigns it to the LookupSecretConfirm field.
func (o *UpdateSettingsFlowWithLookupMethod) SetLookupSecretConfirm(v bool) {
	o.LookupSecretConfirm = &v
}

// GetLookupSecretDisable returns the LookupSecretDisable field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretDisable() bool {
	if o == nil || IsNil(o.LookupSecretDisable) {
		var ret bool
		return ret
	}
	return *o.LookupSecretDisable
}

// GetLookupSecretDisableOk returns a tuple with the LookupSecretDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.LookupSecretDisable) {
		return nil, false
	}
	return o.LookupSecretDisable, true
}

// HasLookupSecretDisable returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasLookupSecretDisable() bool {
	if o != nil && !IsNil(o.LookupSecretDisable) {
		return true
	}

	return false
}

// SetLookupSecretDisable gets a reference to the given bool and assigns it to the LookupSecretDisable field.
func (o *UpdateSettingsFlowWithLookupMethod) SetLookupSecretDisable(v bool) {
	o.LookupSecretDisable = &v
}

// GetLookupSecretRegenerate returns the LookupSecretRegenerate field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretRegenerate() bool {
	if o == nil || IsNil(o.LookupSecretRegenerate) {
		var ret bool
		return ret
	}
	return *o.LookupSecretRegenerate
}

// GetLookupSecretRegenerateOk returns a tuple with the LookupSecretRegenerate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretRegenerateOk() (*bool, bool) {
	if o == nil || IsNil(o.LookupSecretRegenerate) {
		return nil, false
	}
	return o.LookupSecretRegenerate, true
}

// HasLookupSecretRegenerate returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasLookupSecretRegenerate() bool {
	if o != nil && !IsNil(o.LookupSecretRegenerate) {
		return true
	}

	return false
}

// SetLookupSecretRegenerate gets a reference to the given bool and assigns it to the LookupSecretRegenerate field.
func (o *UpdateSettingsFlowWithLookupMethod) SetLookupSecretRegenerate(v bool) {
	o.LookupSecretRegenerate = &v
}

// GetLookupSecretReveal returns the LookupSecretReveal field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretReveal() bool {
	if o == nil || IsNil(o.LookupSecretReveal) {
		var ret bool
		return ret
	}
	return *o.LookupSecretReveal
}

// GetLookupSecretRevealOk returns a tuple with the LookupSecretReveal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetLookupSecretRevealOk() (*bool, bool) {
	if o == nil || IsNil(o.LookupSecretReveal) {
		return nil, false
	}
	return o.LookupSecretReveal, true
}

// HasLookupSecretReveal returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasLookupSecretReveal() bool {
	if o != nil && !IsNil(o.LookupSecretReveal) {
		return true
	}

	return false
}

// SetLookupSecretReveal gets a reference to the given bool and assigns it to the LookupSecretReveal field.
func (o *UpdateSettingsFlowWithLookupMethod) SetLookupSecretReveal(v bool) {
	o.LookupSecretReveal = &v
}

// GetMethod returns the Method field value
func (o *UpdateSettingsFlowWithLookupMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateSettingsFlowWithLookupMethod) SetMethod(v string) {
	o.Method = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithLookupMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithLookupMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithLookupMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateSettingsFlowWithLookupMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateSettingsFlowWithLookupMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSettingsFlowWithLookupMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if !IsNil(o.LookupSecretConfirm) {
		toSerialize["lookup_secret_confirm"] = o.LookupSecretConfirm
	}
	if !IsNil(o.LookupSecretDisable) {
		toSerialize["lookup_secret_disable"] = o.LookupSecretDisable
	}
	if !IsNil(o.LookupSecretRegenerate) {
		toSerialize["lookup_secret_regenerate"] = o.LookupSecretRegenerate
	}
	if !IsNil(o.LookupSecretReveal) {
		toSerialize["lookup_secret_reveal"] = o.LookupSecretReveal
	}
	toSerialize["method"] = o.Method
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSettingsFlowWithLookupMethod) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateSettingsFlowWithLookupMethod := _UpdateSettingsFlowWithLookupMethod{}

	err = json.Unmarshal(data, &varUpdateSettingsFlowWithLookupMethod)

	if err != nil {
		return err
	}

	*o = UpdateSettingsFlowWithLookupMethod(varUpdateSettingsFlowWithLookupMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "lookup_secret_confirm")
		delete(additionalProperties, "lookup_secret_disable")
		delete(additionalProperties, "lookup_secret_regenerate")
		delete(additionalProperties, "lookup_secret_reveal")
		delete(additionalProperties, "method")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSettingsFlowWithLookupMethod struct {
	value *UpdateSettingsFlowWithLookupMethod
	isSet bool
}

func (v NullableUpdateSettingsFlowWithLookupMethod) Get() *UpdateSettingsFlowWithLookupMethod {
	return v.value
}

func (v *NullableUpdateSettingsFlowWithLookupMethod) Set(val *UpdateSettingsFlowWithLookupMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingsFlowWithLookupMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingsFlowWithLookupMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingsFlowWithLookupMethod(val *UpdateSettingsFlowWithLookupMethod) *NullableUpdateSettingsFlowWithLookupMethod {
	return &NullableUpdateSettingsFlowWithLookupMethod{value: val, isSet: true}
}

func (v NullableUpdateSettingsFlowWithLookupMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingsFlowWithLookupMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


