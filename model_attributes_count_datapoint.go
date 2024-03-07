/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.8.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the AttributesCountDatapoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributesCountDatapoint{}

// AttributesCountDatapoint struct for AttributesCountDatapoint
type AttributesCountDatapoint struct {
	// Count of the attribute value for given key
	Count int64 `json:"count"`
	// Name of the attribute value for given key
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AttributesCountDatapoint AttributesCountDatapoint

// NewAttributesCountDatapoint instantiates a new AttributesCountDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributesCountDatapoint(count int64, name string) *AttributesCountDatapoint {
	this := AttributesCountDatapoint{}
	this.Count = count
	this.Name = name
	return &this
}

// NewAttributesCountDatapointWithDefaults instantiates a new AttributesCountDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributesCountDatapointWithDefaults() *AttributesCountDatapoint {
	this := AttributesCountDatapoint{}
	return &this
}

// GetCount returns the Count field value
func (o *AttributesCountDatapoint) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AttributesCountDatapoint) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AttributesCountDatapoint) SetCount(v int64) {
	o.Count = v
}

// GetName returns the Name field value
func (o *AttributesCountDatapoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttributesCountDatapoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttributesCountDatapoint) SetName(v string) {
	o.Name = v
}

func (o AttributesCountDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributesCountDatapoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttributesCountDatapoint) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"name",
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

	varAttributesCountDatapoint := _AttributesCountDatapoint{}

	err = json.Unmarshal(bytes, &varAttributesCountDatapoint)

	if err != nil {
		return err
	}

	*o = AttributesCountDatapoint(varAttributesCountDatapoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttributesCountDatapoint struct {
	value *AttributesCountDatapoint
	isSet bool
}

func (v NullableAttributesCountDatapoint) Get() *AttributesCountDatapoint {
	return v.value
}

func (v *NullableAttributesCountDatapoint) Set(val *AttributesCountDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributesCountDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributesCountDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributesCountDatapoint(val *AttributesCountDatapoint) *NullableAttributesCountDatapoint {
	return &NullableAttributesCountDatapoint{value: val, isSet: true}
}

func (v NullableAttributesCountDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributesCountDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


