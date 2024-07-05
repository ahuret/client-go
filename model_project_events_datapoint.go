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
	"time"
	"fmt"
)

// checks if the ProjectEventsDatapoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEventsDatapoint{}

// ProjectEventsDatapoint struct for ProjectEventsDatapoint
type ProjectEventsDatapoint struct {
	// Event attributes with details
	Attributes []Attribute `json:"attributes"`
	// Name of the event
	Name string `json:"name"`
	// Time of occurence
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ProjectEventsDatapoint ProjectEventsDatapoint

// NewProjectEventsDatapoint instantiates a new ProjectEventsDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEventsDatapoint(attributes []Attribute, name string, timestamp time.Time) *ProjectEventsDatapoint {
	this := ProjectEventsDatapoint{}
	this.Attributes = attributes
	this.Name = name
	this.Timestamp = timestamp
	return &this
}

// NewProjectEventsDatapointWithDefaults instantiates a new ProjectEventsDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEventsDatapointWithDefaults() *ProjectEventsDatapoint {
	this := ProjectEventsDatapoint{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *ProjectEventsDatapoint) GetAttributes() []Attribute {
	if o == nil {
		var ret []Attribute
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProjectEventsDatapoint) GetAttributesOk() ([]Attribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ProjectEventsDatapoint) SetAttributes(v []Attribute) {
	o.Attributes = v
}

// GetName returns the Name field value
func (o *ProjectEventsDatapoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectEventsDatapoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectEventsDatapoint) SetName(v string) {
	o.Name = v
}

// GetTimestamp returns the Timestamp field value
func (o *ProjectEventsDatapoint) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ProjectEventsDatapoint) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ProjectEventsDatapoint) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ProjectEventsDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEventsDatapoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["name"] = o.Name
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectEventsDatapoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
		"name",
		"timestamp",
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

	varProjectEventsDatapoint := _ProjectEventsDatapoint{}

	err = json.Unmarshal(data, &varProjectEventsDatapoint)

	if err != nil {
		return err
	}

	*o = ProjectEventsDatapoint(varProjectEventsDatapoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectEventsDatapoint struct {
	value *ProjectEventsDatapoint
	isSet bool
}

func (v NullableProjectEventsDatapoint) Get() *ProjectEventsDatapoint {
	return v.value
}

func (v *NullableProjectEventsDatapoint) Set(val *ProjectEventsDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEventsDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEventsDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEventsDatapoint(val *ProjectEventsDatapoint) *NullableProjectEventsDatapoint {
	return &NullableProjectEventsDatapoint{value: val, isSet: true}
}

func (v NullableProjectEventsDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEventsDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


