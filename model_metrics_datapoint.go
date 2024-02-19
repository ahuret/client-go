/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.6.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the MetricsDatapoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsDatapoint{}

// MetricsDatapoint Represents a single datapoint/bucket of a time series
type MetricsDatapoint struct {
	// The count of events that occured in this time
	Count int64 `json:"count"`
	// The time of the bucket
	Time time.Time `json:"time"`
	AdditionalProperties map[string]interface{}
}

type _MetricsDatapoint MetricsDatapoint

// NewMetricsDatapoint instantiates a new MetricsDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsDatapoint(count int64, time time.Time) *MetricsDatapoint {
	this := MetricsDatapoint{}
	this.Count = count
	this.Time = time
	return &this
}

// NewMetricsDatapointWithDefaults instantiates a new MetricsDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsDatapointWithDefaults() *MetricsDatapoint {
	this := MetricsDatapoint{}
	return &this
}

// GetCount returns the Count field value
func (o *MetricsDatapoint) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *MetricsDatapoint) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *MetricsDatapoint) SetCount(v int64) {
	o.Count = v
}

// GetTime returns the Time field value
func (o *MetricsDatapoint) GetTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *MetricsDatapoint) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *MetricsDatapoint) SetTime(v time.Time) {
	o.Time = v
}

func (o MetricsDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsDatapoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricsDatapoint) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"time",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMetricsDatapoint := _MetricsDatapoint{}

	err = json.Unmarshal(bytes, &varMetricsDatapoint)

	if err != nil {
		return err
	}

	*o = MetricsDatapoint(varMetricsDatapoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetricsDatapoint struct {
	value *MetricsDatapoint
	isSet bool
}

func (v NullableMetricsDatapoint) Get() *MetricsDatapoint {
	return v.value
}

func (v *NullableMetricsDatapoint) Set(val *MetricsDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsDatapoint(val *MetricsDatapoint) *NullableMetricsDatapoint {
	return &NullableMetricsDatapoint{value: val, isSet: true}
}

func (v NullableMetricsDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


