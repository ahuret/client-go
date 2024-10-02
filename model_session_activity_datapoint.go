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
	"fmt"
)

// checks if the SessionActivityDatapoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionActivityDatapoint{}

// SessionActivityDatapoint struct for SessionActivityDatapoint
type SessionActivityDatapoint struct {
	// Country of the events
	Country string `json:"country"`
	// Number of events that failed in the given timeframe
	Failed int64 `json:"failed"`
	// Number of events that succeeded in the given timeframe
	Succeeded int64 `json:"succeeded"`
	AdditionalProperties map[string]interface{}
}

type _SessionActivityDatapoint SessionActivityDatapoint

// NewSessionActivityDatapoint instantiates a new SessionActivityDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionActivityDatapoint(country string, failed int64, succeeded int64) *SessionActivityDatapoint {
	this := SessionActivityDatapoint{}
	this.Country = country
	this.Failed = failed
	this.Succeeded = succeeded
	return &this
}

// NewSessionActivityDatapointWithDefaults instantiates a new SessionActivityDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionActivityDatapointWithDefaults() *SessionActivityDatapoint {
	this := SessionActivityDatapoint{}
	return &this
}

// GetCountry returns the Country field value
func (o *SessionActivityDatapoint) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *SessionActivityDatapoint) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *SessionActivityDatapoint) SetCountry(v string) {
	o.Country = v
}

// GetFailed returns the Failed field value
func (o *SessionActivityDatapoint) GetFailed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *SessionActivityDatapoint) GetFailedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value
func (o *SessionActivityDatapoint) SetFailed(v int64) {
	o.Failed = v
}

// GetSucceeded returns the Succeeded field value
func (o *SessionActivityDatapoint) GetSucceeded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value
// and a boolean to check if the value has been set.
func (o *SessionActivityDatapoint) GetSucceededOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Succeeded, true
}

// SetSucceeded sets field value
func (o *SessionActivityDatapoint) SetSucceeded(v int64) {
	o.Succeeded = v
}

func (o SessionActivityDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionActivityDatapoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	toSerialize["failed"] = o.Failed
	toSerialize["succeeded"] = o.Succeeded

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionActivityDatapoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
		"failed",
		"succeeded",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSessionActivityDatapoint := _SessionActivityDatapoint{}

	err = json.Unmarshal(data, &varSessionActivityDatapoint)

	if err != nil {
		return err
	}

	*o = SessionActivityDatapoint(varSessionActivityDatapoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "succeeded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionActivityDatapoint struct {
	value *SessionActivityDatapoint
	isSet bool
}

func (v NullableSessionActivityDatapoint) Get() *SessionActivityDatapoint {
	return v.value
}

func (v *NullableSessionActivityDatapoint) Set(val *SessionActivityDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionActivityDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionActivityDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionActivityDatapoint(val *SessionActivityDatapoint) *NullableSessionActivityDatapoint {
	return &NullableSessionActivityDatapoint{value: val, isSet: true}
}

func (v NullableSessionActivityDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionActivityDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


