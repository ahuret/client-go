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

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan struct for Plan
type Plan struct {
	// Name is the name of the plan.
	Name string `json:"name"`
	// Version is the version of the plan. The combination of `name@version` must be unique.
	Version int64 `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _Plan Plan

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan(name string, version int64) *Plan {
	this := Plan{}
	this.Name = name
	this.Version = version
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetName returns the Name field value
func (o *Plan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Plan) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *Plan) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Plan) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Plan) SetVersion(v int64) {
	o.Version = v
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Plan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
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

	varPlan := _Plan{}

	err = json.Unmarshal(data, &varPlan)

	if err != nil {
		return err
	}

	*o = Plan(varPlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


