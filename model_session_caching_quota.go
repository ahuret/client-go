/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.15
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SessionCachingQuota Session Caching Quota
type SessionCachingQuota struct {
	CanUse *bool `json:"can_use,omitempty"`
}

// NewSessionCachingQuota instantiates a new SessionCachingQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionCachingQuota() *SessionCachingQuota {
	this := SessionCachingQuota{}
	return &this
}

// NewSessionCachingQuotaWithDefaults instantiates a new SessionCachingQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionCachingQuotaWithDefaults() *SessionCachingQuota {
	this := SessionCachingQuota{}
	return &this
}

// GetCanUse returns the CanUse field value if set, zero value otherwise.
func (o *SessionCachingQuota) GetCanUse() bool {
	if o == nil || o.CanUse == nil {
		var ret bool
		return ret
	}
	return *o.CanUse
}

// GetCanUseOk returns a tuple with the CanUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionCachingQuota) GetCanUseOk() (*bool, bool) {
	if o == nil || o.CanUse == nil {
		return nil, false
	}
	return o.CanUse, true
}

// HasCanUse returns a boolean if a field has been set.
func (o *SessionCachingQuota) HasCanUse() bool {
	if o != nil && o.CanUse != nil {
		return true
	}

	return false
}

// SetCanUse gets a reference to the given bool and assigns it to the CanUse field.
func (o *SessionCachingQuota) SetCanUse(v bool) {
	o.CanUse = &v
}

func (o SessionCachingQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanUse != nil {
		toSerialize["can_use"] = o.CanUse
	}
	return json.Marshal(toSerialize)
}

type NullableSessionCachingQuota struct {
	value *SessionCachingQuota
	isSet bool
}

func (v NullableSessionCachingQuota) Get() *SessionCachingQuota {
	return v.value
}

func (v *NullableSessionCachingQuota) Set(val *SessionCachingQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionCachingQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionCachingQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionCachingQuota(val *SessionCachingQuota) *NullableSessionCachingQuota {
	return &NullableSessionCachingQuota{value: val, isSet: true}
}

func (v NullableSessionCachingQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionCachingQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


