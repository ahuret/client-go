/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.23
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SelfServiceBrowserLocationChangeRequiredError struct for SelfServiceBrowserLocationChangeRequiredError
type SelfServiceBrowserLocationChangeRequiredError struct {
	// The status code
	Code *int64 `json:"code,omitempty"`
	// Debug information  This field is often not exposed to protect against leaking sensitive information.
	Debug *string `json:"debug,omitempty"`
	// Further error details
	Details map[string]interface{} `json:"details,omitempty"`
	// The error ID  Useful when trying to identify various errors in application logic.
	Id *string `json:"id,omitempty"`
	// Error message  The error's message.
	Message string `json:"message"`
	// A human-readable reason for the error
	Reason *string `json:"reason,omitempty"`
	// Since when the flow has expired
	RedirectBrowserTo *string `json:"redirect_browser_to,omitempty"`
	// The request ID  The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID.
	Request *string `json:"request,omitempty"`
	// The status description
	Status *string `json:"status,omitempty"`
}

// NewSelfServiceBrowserLocationChangeRequiredError instantiates a new SelfServiceBrowserLocationChangeRequiredError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceBrowserLocationChangeRequiredError(message string) *SelfServiceBrowserLocationChangeRequiredError {
	this := SelfServiceBrowserLocationChangeRequiredError{}
	this.Message = message
	return &this
}

// NewSelfServiceBrowserLocationChangeRequiredErrorWithDefaults instantiates a new SelfServiceBrowserLocationChangeRequiredError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceBrowserLocationChangeRequiredErrorWithDefaults() *SelfServiceBrowserLocationChangeRequiredError {
	this := SelfServiceBrowserLocationChangeRequiredError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetCode() int64 {
	if o == nil || o.Code == nil {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetCodeOk() (*int64, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetCode(v int64) {
	o.Code = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetDebug() string {
	if o == nil || o.Debug == nil {
		var ret string
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetDebugOk() (*string, bool) {
	if o == nil || o.Debug == nil {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasDebug() bool {
	if o != nil && o.Debug != nil {
		return true
	}

	return false
}

// SetDebug gets a reference to the given string and assigns it to the Debug field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetDebug(v string) {
	o.Debug = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetDetails() map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value
func (o *SelfServiceBrowserLocationChangeRequiredError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SelfServiceBrowserLocationChangeRequiredError) SetMessage(v string) {
	o.Message = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetReason(v string) {
	o.Reason = &v
}

// GetRedirectBrowserTo returns the RedirectBrowserTo field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetRedirectBrowserTo() string {
	if o == nil || o.RedirectBrowserTo == nil {
		var ret string
		return ret
	}
	return *o.RedirectBrowserTo
}

// GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetRedirectBrowserToOk() (*string, bool) {
	if o == nil || o.RedirectBrowserTo == nil {
		return nil, false
	}
	return o.RedirectBrowserTo, true
}

// HasRedirectBrowserTo returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasRedirectBrowserTo() bool {
	if o != nil && o.RedirectBrowserTo != nil {
		return true
	}

	return false
}

// SetRedirectBrowserTo gets a reference to the given string and assigns it to the RedirectBrowserTo field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetRedirectBrowserTo(v string) {
	o.RedirectBrowserTo = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetRequest(v string) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SelfServiceBrowserLocationChangeRequiredError) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SelfServiceBrowserLocationChangeRequiredError) SetStatus(v string) {
	o.Status = &v
}

func (o SelfServiceBrowserLocationChangeRequiredError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Debug != nil {
		toSerialize["debug"] = o.Debug
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.RedirectBrowserTo != nil {
		toSerialize["redirect_browser_to"] = o.RedirectBrowserTo
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceBrowserLocationChangeRequiredError struct {
	value *SelfServiceBrowserLocationChangeRequiredError
	isSet bool
}

func (v NullableSelfServiceBrowserLocationChangeRequiredError) Get() *SelfServiceBrowserLocationChangeRequiredError {
	return v.value
}

func (v *NullableSelfServiceBrowserLocationChangeRequiredError) Set(val *SelfServiceBrowserLocationChangeRequiredError) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceBrowserLocationChangeRequiredError) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceBrowserLocationChangeRequiredError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceBrowserLocationChangeRequiredError(val *SelfServiceBrowserLocationChangeRequiredError) *NullableSelfServiceBrowserLocationChangeRequiredError {
	return &NullableSelfServiceBrowserLocationChangeRequiredError{value: val, isSet: true}
}

func (v NullableSelfServiceBrowserLocationChangeRequiredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceBrowserLocationChangeRequiredError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


