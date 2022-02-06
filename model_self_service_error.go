/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.66
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// SelfServiceError struct for SelfServiceError
type SelfServiceError struct {
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Error map[string]interface{} `json:"error,omitempty"`
	Id string `json:"id"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewSelfServiceError instantiates a new SelfServiceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceError(id string) *SelfServiceError {
	this := SelfServiceError{}
	this.Id = id
	return &this
}

// NewSelfServiceErrorWithDefaults instantiates a new SelfServiceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceErrorWithDefaults() *SelfServiceError {
	this := SelfServiceError{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SelfServiceError) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceError) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SelfServiceError) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SelfServiceError) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SelfServiceError) GetError() map[string]interface{} {
	if o == nil || o.Error == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceError) GetErrorOk() (map[string]interface{}, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SelfServiceError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given map[string]interface{} and assigns it to the Error field.
func (o *SelfServiceError) SetError(v map[string]interface{}) {
	o.Error = v
}

// GetId returns the Id field value
func (o *SelfServiceError) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SelfServiceError) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SelfServiceError) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SelfServiceError) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceError) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SelfServiceError) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SelfServiceError) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o SelfServiceError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceError struct {
	value *SelfServiceError
	isSet bool
}

func (v NullableSelfServiceError) Get() *SelfServiceError {
	return v.value
}

func (v *NullableSelfServiceError) Set(val *SelfServiceError) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceError) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceError(val *SelfServiceError) *NullableSelfServiceError {
	return &NullableSelfServiceError{value: val, isSet: true}
}

func (v NullableSelfServiceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


