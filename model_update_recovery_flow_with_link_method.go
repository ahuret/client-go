/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateRecoveryFlowWithLinkMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRecoveryFlowWithLinkMethod{}

// UpdateRecoveryFlowWithLinkMethod Update Recovery Flow with Link Method
type UpdateRecoveryFlowWithLinkMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
	Email string `json:"email"`
	// Method is the method that should be used for this recovery flow  Allowed values are `link` and `code` link RecoveryStrategyLink code RecoveryStrategyCode
	Method string `json:"method"`
	// Transient data to pass along to any webhooks
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRecoveryFlowWithLinkMethod UpdateRecoveryFlowWithLinkMethod

// NewUpdateRecoveryFlowWithLinkMethod instantiates a new UpdateRecoveryFlowWithLinkMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecoveryFlowWithLinkMethod(email string, method string) *UpdateRecoveryFlowWithLinkMethod {
	this := UpdateRecoveryFlowWithLinkMethod{}
	this.Email = email
	this.Method = method
	return &this
}

// NewUpdateRecoveryFlowWithLinkMethodWithDefaults instantiates a new UpdateRecoveryFlowWithLinkMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecoveryFlowWithLinkMethodWithDefaults() *UpdateRecoveryFlowWithLinkMethod {
	this := UpdateRecoveryFlowWithLinkMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateRecoveryFlowWithLinkMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateRecoveryFlowWithLinkMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value
func (o *UpdateRecoveryFlowWithLinkMethod) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UpdateRecoveryFlowWithLinkMethod) SetEmail(v string) {
	o.Email = v
}

// GetMethod returns the Method field value
func (o *UpdateRecoveryFlowWithLinkMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateRecoveryFlowWithLinkMethod) SetMethod(v string) {
	o.Method = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *UpdateRecoveryFlowWithLinkMethod) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *UpdateRecoveryFlowWithLinkMethod) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *UpdateRecoveryFlowWithLinkMethod) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

func (o UpdateRecoveryFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRecoveryFlowWithLinkMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["email"] = o.Email
	toSerialize["method"] = o.Method
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRecoveryFlowWithLinkMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUpdateRecoveryFlowWithLinkMethod := _UpdateRecoveryFlowWithLinkMethod{}

	err = json.Unmarshal(data, &varUpdateRecoveryFlowWithLinkMethod)

	if err != nil {
		return err
	}

	*o = UpdateRecoveryFlowWithLinkMethod(varUpdateRecoveryFlowWithLinkMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "email")
		delete(additionalProperties, "method")
		delete(additionalProperties, "transient_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRecoveryFlowWithLinkMethod struct {
	value *UpdateRecoveryFlowWithLinkMethod
	isSet bool
}

func (v NullableUpdateRecoveryFlowWithLinkMethod) Get() *UpdateRecoveryFlowWithLinkMethod {
	return v.value
}

func (v *NullableUpdateRecoveryFlowWithLinkMethod) Set(val *UpdateRecoveryFlowWithLinkMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecoveryFlowWithLinkMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecoveryFlowWithLinkMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecoveryFlowWithLinkMethod(val *UpdateRecoveryFlowWithLinkMethod) *NullableUpdateRecoveryFlowWithLinkMethod {
	return &NullableUpdateRecoveryFlowWithLinkMethod{value: val, isSet: true}
}

func (v NullableUpdateRecoveryFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecoveryFlowWithLinkMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


