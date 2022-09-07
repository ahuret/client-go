/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.33
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RejectRequest struct for RejectRequest
type RejectRequest struct {
	// The error should follow the OAuth2 error format (e.g. `invalid_request`, `login_required`).  Defaults to `request_denied`.
	Error *string `json:"error,omitempty"`
	// Debug contains information to help resolve the problem as a developer. Usually not exposed to the public but only in the server logs.
	ErrorDebug *string `json:"error_debug,omitempty"`
	// Description of the error in a human readable format.
	ErrorDescription *string `json:"error_description,omitempty"`
	// Hint to help resolve the error.
	ErrorHint *string `json:"error_hint,omitempty"`
	// Represents the HTTP status code of the error (e.g. 401 or 403)  Defaults to 400
	StatusCode *int64 `json:"status_code,omitempty"`
}

// NewRejectRequest instantiates a new RejectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectRequest() *RejectRequest {
	this := RejectRequest{}
	return &this
}

// NewRejectRequestWithDefaults instantiates a new RejectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectRequestWithDefaults() *RejectRequest {
	this := RejectRequest{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RejectRequest) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequest) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RejectRequest) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RejectRequest) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *RejectRequest) GetErrorDebug() string {
	if o == nil || o.ErrorDebug == nil {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequest) GetErrorDebugOk() (*string, bool) {
	if o == nil || o.ErrorDebug == nil {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *RejectRequest) HasErrorDebug() bool {
	if o != nil && o.ErrorDebug != nil {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *RejectRequest) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *RejectRequest) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequest) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *RejectRequest) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *RejectRequest) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorHint returns the ErrorHint field value if set, zero value otherwise.
func (o *RejectRequest) GetErrorHint() string {
	if o == nil || o.ErrorHint == nil {
		var ret string
		return ret
	}
	return *o.ErrorHint
}

// GetErrorHintOk returns a tuple with the ErrorHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequest) GetErrorHintOk() (*string, bool) {
	if o == nil || o.ErrorHint == nil {
		return nil, false
	}
	return o.ErrorHint, true
}

// HasErrorHint returns a boolean if a field has been set.
func (o *RejectRequest) HasErrorHint() bool {
	if o != nil && o.ErrorHint != nil {
		return true
	}

	return false
}

// SetErrorHint gets a reference to the given string and assigns it to the ErrorHint field.
func (o *RejectRequest) SetErrorHint(v string) {
	o.ErrorHint = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *RejectRequest) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectRequest) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *RejectRequest) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *RejectRequest) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o RejectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.ErrorDebug != nil {
		toSerialize["error_debug"] = o.ErrorDebug
	}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.ErrorHint != nil {
		toSerialize["error_hint"] = o.ErrorHint
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableRejectRequest struct {
	value *RejectRequest
	isSet bool
}

func (v NullableRejectRequest) Get() *RejectRequest {
	return v.value
}

func (v *NullableRejectRequest) Set(val *RejectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectRequest(val *RejectRequest) *NullableRejectRequest {
	return &NullableRejectRequest{value: val, isSet: true}
}

func (v NullableRejectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


