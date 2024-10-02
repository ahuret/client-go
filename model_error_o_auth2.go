/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.5
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ErrorOAuth2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorOAuth2{}

// ErrorOAuth2 Error
type ErrorOAuth2 struct {
	// Error
	Error *string `json:"error,omitempty"`
	// Error Debug Information  Only available in dev mode.
	ErrorDebug *string `json:"error_debug,omitempty"`
	// Error Description
	ErrorDescription *string `json:"error_description,omitempty"`
	// Error Hint  Helps the user identify the error cause.
	ErrorHint *string `json:"error_hint,omitempty"`
	// HTTP Status Code
	StatusCode *int64 `json:"status_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorOAuth2 ErrorOAuth2

// NewErrorOAuth2 instantiates a new ErrorOAuth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorOAuth2() *ErrorOAuth2 {
	this := ErrorOAuth2{}
	return &this
}

// NewErrorOAuth2WithDefaults instantiates a new ErrorOAuth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorOAuth2WithDefaults() *ErrorOAuth2 {
	this := ErrorOAuth2{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorOAuth2) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorOAuth2) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorOAuth2) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ErrorOAuth2) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *ErrorOAuth2) GetErrorDebug() string {
	if o == nil || IsNil(o.ErrorDebug) {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorOAuth2) GetErrorDebugOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDebug) {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *ErrorOAuth2) HasErrorDebug() bool {
	if o != nil && !IsNil(o.ErrorDebug) {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *ErrorOAuth2) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *ErrorOAuth2) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorOAuth2) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *ErrorOAuth2) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *ErrorOAuth2) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorHint returns the ErrorHint field value if set, zero value otherwise.
func (o *ErrorOAuth2) GetErrorHint() string {
	if o == nil || IsNil(o.ErrorHint) {
		var ret string
		return ret
	}
	return *o.ErrorHint
}

// GetErrorHintOk returns a tuple with the ErrorHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorOAuth2) GetErrorHintOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorHint) {
		return nil, false
	}
	return o.ErrorHint, true
}

// HasErrorHint returns a boolean if a field has been set.
func (o *ErrorOAuth2) HasErrorHint() bool {
	if o != nil && !IsNil(o.ErrorHint) {
		return true
	}

	return false
}

// SetErrorHint gets a reference to the given string and assigns it to the ErrorHint field.
func (o *ErrorOAuth2) SetErrorHint(v string) {
	o.ErrorHint = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ErrorOAuth2) GetStatusCode() int64 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorOAuth2) GetStatusCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ErrorOAuth2) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *ErrorOAuth2) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o ErrorOAuth2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorOAuth2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorDebug) {
		toSerialize["error_debug"] = o.ErrorDebug
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if !IsNil(o.ErrorHint) {
		toSerialize["error_hint"] = o.ErrorHint
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorOAuth2) UnmarshalJSON(data []byte) (err error) {
	varErrorOAuth2 := _ErrorOAuth2{}

	err = json.Unmarshal(data, &varErrorOAuth2)

	if err != nil {
		return err
	}

	*o = ErrorOAuth2(varErrorOAuth2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "error_debug")
		delete(additionalProperties, "error_description")
		delete(additionalProperties, "error_hint")
		delete(additionalProperties, "status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorOAuth2 struct {
	value *ErrorOAuth2
	isSet bool
}

func (v NullableErrorOAuth2) Get() *ErrorOAuth2 {
	return v.value
}

func (v *NullableErrorOAuth2) Set(val *ErrorOAuth2) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorOAuth2) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorOAuth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorOAuth2(val *ErrorOAuth2) *NullableErrorOAuth2 {
	return &NullableErrorOAuth2{value: val, isSet: true}
}

func (v NullableErrorOAuth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorOAuth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


