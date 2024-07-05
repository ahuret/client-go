/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the RejectOAuth2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectOAuth2Request{}

// RejectOAuth2Request struct for RejectOAuth2Request
type RejectOAuth2Request struct {
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
	AdditionalProperties map[string]interface{}
}

type _RejectOAuth2Request RejectOAuth2Request

// NewRejectOAuth2Request instantiates a new RejectOAuth2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectOAuth2Request() *RejectOAuth2Request {
	this := RejectOAuth2Request{}
	return &this
}

// NewRejectOAuth2RequestWithDefaults instantiates a new RejectOAuth2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectOAuth2RequestWithDefaults() *RejectOAuth2Request {
	this := RejectOAuth2Request{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RejectOAuth2Request) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectOAuth2Request) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RejectOAuth2Request) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RejectOAuth2Request) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *RejectOAuth2Request) GetErrorDebug() string {
	if o == nil || IsNil(o.ErrorDebug) {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectOAuth2Request) GetErrorDebugOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDebug) {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *RejectOAuth2Request) HasErrorDebug() bool {
	if o != nil && !IsNil(o.ErrorDebug) {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *RejectOAuth2Request) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *RejectOAuth2Request) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectOAuth2Request) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *RejectOAuth2Request) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *RejectOAuth2Request) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorHint returns the ErrorHint field value if set, zero value otherwise.
func (o *RejectOAuth2Request) GetErrorHint() string {
	if o == nil || IsNil(o.ErrorHint) {
		var ret string
		return ret
	}
	return *o.ErrorHint
}

// GetErrorHintOk returns a tuple with the ErrorHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectOAuth2Request) GetErrorHintOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorHint) {
		return nil, false
	}
	return o.ErrorHint, true
}

// HasErrorHint returns a boolean if a field has been set.
func (o *RejectOAuth2Request) HasErrorHint() bool {
	if o != nil && !IsNil(o.ErrorHint) {
		return true
	}

	return false
}

// SetErrorHint gets a reference to the given string and assigns it to the ErrorHint field.
func (o *RejectOAuth2Request) SetErrorHint(v string) {
	o.ErrorHint = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *RejectOAuth2Request) GetStatusCode() int64 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectOAuth2Request) GetStatusCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *RejectOAuth2Request) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *RejectOAuth2Request) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o RejectOAuth2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RejectOAuth2Request) ToMap() (map[string]interface{}, error) {
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

func (o *RejectOAuth2Request) UnmarshalJSON(data []byte) (err error) {
	varRejectOAuth2Request := _RejectOAuth2Request{}

	err = json.Unmarshal(data, &varRejectOAuth2Request)

	if err != nil {
		return err
	}

	*o = RejectOAuth2Request(varRejectOAuth2Request)

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

type NullableRejectOAuth2Request struct {
	value *RejectOAuth2Request
	isSet bool
}

func (v NullableRejectOAuth2Request) Get() *RejectOAuth2Request {
	return v.value
}

func (v *NullableRejectOAuth2Request) Set(val *RejectOAuth2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectOAuth2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectOAuth2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectOAuth2Request(val *RejectOAuth2Request) *NullableRejectOAuth2Request {
	return &NullableRejectOAuth2Request{value: val, isSet: true}
}

func (v NullableRejectOAuth2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectOAuth2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


