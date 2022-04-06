/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.156
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SessionDevice struct for SessionDevice
type SessionDevice struct {
	// UserAgent of this device
	UserAgent *string `json:"user_agent,omitempty"`
}

// NewSessionDevice instantiates a new SessionDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionDevice() *SessionDevice {
	this := SessionDevice{}
	return &this
}

// NewSessionDeviceWithDefaults instantiates a new SessionDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionDeviceWithDefaults() *SessionDevice {
	this := SessionDevice{}
	return &this
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *SessionDevice) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionDevice) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *SessionDevice) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *SessionDevice) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o SessionDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}
	return json.Marshal(toSerialize)
}

type NullableSessionDevice struct {
	value *SessionDevice
	isSet bool
}

func (v NullableSessionDevice) Get() *SessionDevice {
	return v.value
}

func (v *NullableSessionDevice) Set(val *SessionDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionDevice(val *SessionDevice) *NullableSessionDevice {
	return &NullableSessionDevice{value: val, isSet: true}
}

func (v NullableSessionDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


