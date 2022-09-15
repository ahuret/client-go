/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.44
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GenericErrorContent Error response
type GenericErrorContent struct {
	// Debug contains debug information. This is usually not available and has to be enabled.
	Debug *string `json:"debug,omitempty"`
	// Name is the error name.
	Error *string `json:"error,omitempty"`
	// Description contains further information on the nature of the error.
	ErrorDescription *string `json:"error_description,omitempty"`
	// Message contains the error message.
	Message *string `json:"message,omitempty"`
	// Code represents the error status code (404, 403, 401, ...).
	StatusCode *int64 `json:"status_code,omitempty"`
}

// NewGenericErrorContent instantiates a new GenericErrorContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericErrorContent() *GenericErrorContent {
	this := GenericErrorContent{}
	return &this
}

// NewGenericErrorContentWithDefaults instantiates a new GenericErrorContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorContentWithDefaults() *GenericErrorContent {
	this := GenericErrorContent{}
	return &this
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *GenericErrorContent) GetDebug() string {
	if o == nil || o.Debug == nil {
		var ret string
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorContent) GetDebugOk() (*string, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *GenericErrorContent) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given string and assigns it to the Debug field.
func (o *GenericErrorContent) SetDebug(v string) {
	o.Debug = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GenericErrorContent) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorContent) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GenericErrorContent) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GenericErrorContent) SetError(v string) {
	o.Error = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *GenericErrorContent) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorContent) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *GenericErrorContent) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *GenericErrorContent) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GenericErrorContent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorContent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GenericErrorContent) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GenericErrorContent) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *GenericErrorContent) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericErrorContent) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *GenericErrorContent) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *GenericErrorContent) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o GenericErrorContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableGenericErrorContent struct {
	value *GenericErrorContent
	isSet bool
}

func (v NullableGenericErrorContent) Get() *GenericErrorContent {
	return v.value
}

func (v *NullableGenericErrorContent) Set(val *GenericErrorContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericErrorContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericErrorContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericErrorContent(val *GenericErrorContent) *NullableGenericErrorContent {
	return &NullableGenericErrorContent{value: val, isSet: true}
}

func (v NullableGenericErrorContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericErrorContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


