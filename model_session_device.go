/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.8
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SessionDevice Device corresponding to a Session
type SessionDevice struct {
	// Device record ID
	Id string `json:"id"`
	// IPAddress of the client
	IpAddress *string `json:"ip_address,omitempty"`
	// Geo Location corresponding to the IP Address
	Location *string `json:"location,omitempty"`
	// UserAgent of the client
	UserAgent *string `json:"user_agent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionDevice SessionDevice

// NewSessionDevice instantiates a new SessionDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionDevice(id string) *SessionDevice {
	this := SessionDevice{}
	this.Id = id
	return &this
}

// NewSessionDeviceWithDefaults instantiates a new SessionDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionDeviceWithDefaults() *SessionDevice {
	this := SessionDevice{}
	return &this
}

// GetId returns the Id field value
func (o *SessionDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SessionDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SessionDevice) SetId(v string) {
	o.Id = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SessionDevice) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionDevice) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SessionDevice) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SessionDevice) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SessionDevice) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionDevice) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SessionDevice) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SessionDevice) SetLocation(v string) {
	o.Location = &v
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
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionDevice) UnmarshalJSON(bytes []byte) (err error) {
	varSessionDevice := _SessionDevice{}

	if err = json.Unmarshal(bytes, &varSessionDevice); err == nil {
		*o = SessionDevice(varSessionDevice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ip_address")
		delete(additionalProperties, "location")
		delete(additionalProperties, "user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


