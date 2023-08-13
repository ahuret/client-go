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

// UpdateSettingsFlowWithPasswordMethod Update Settings Flow with Password Method
type UpdateSettingsFlowWithPasswordMethod struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to password when trying to update a password.
	Method string `json:"method"`
	// Password is the updated password
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSettingsFlowWithPasswordMethod UpdateSettingsFlowWithPasswordMethod

// NewUpdateSettingsFlowWithPasswordMethod instantiates a new UpdateSettingsFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingsFlowWithPasswordMethod(method string, password string) *UpdateSettingsFlowWithPasswordMethod {
	this := UpdateSettingsFlowWithPasswordMethod{}
	this.Method = method
	this.Password = password
	return &this
}

// NewUpdateSettingsFlowWithPasswordMethodWithDefaults instantiates a new UpdateSettingsFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingsFlowWithPasswordMethodWithDefaults() *UpdateSettingsFlowWithPasswordMethod {
	this := UpdateSettingsFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateSettingsFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *UpdateSettingsFlowWithPasswordMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateSettingsFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = v
}

// GetPassword returns the Password field value
func (o *UpdateSettingsFlowWithPasswordMethod) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateSettingsFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = v
}

func (o UpdateSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateSettingsFlowWithPasswordMethod) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateSettingsFlowWithPasswordMethod := _UpdateSettingsFlowWithPasswordMethod{}

	if err = json.Unmarshal(bytes, &varUpdateSettingsFlowWithPasswordMethod); err == nil {
		*o = UpdateSettingsFlowWithPasswordMethod(varUpdateSettingsFlowWithPasswordMethod)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csrf_token")
		delete(additionalProperties, "method")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSettingsFlowWithPasswordMethod struct {
	value *UpdateSettingsFlowWithPasswordMethod
	isSet bool
}

func (v NullableUpdateSettingsFlowWithPasswordMethod) Get() *UpdateSettingsFlowWithPasswordMethod {
	return v.value
}

func (v *NullableUpdateSettingsFlowWithPasswordMethod) Set(val *UpdateSettingsFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingsFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingsFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingsFlowWithPasswordMethod(val *UpdateSettingsFlowWithPasswordMethod) *NullableUpdateSettingsFlowWithPasswordMethod {
	return &NullableUpdateSettingsFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableUpdateSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingsFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


