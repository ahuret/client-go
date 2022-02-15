/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.95
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// NullPlan the model 'NullPlan'
type NullPlan string

// List of NullPlan
const (
	NULLPLAN_UNKNOWN NullPlan = "unknown"
	NULLPLAN_FREE NullPlan = "free"
	NULLPLAN_START_UP_MONTHLY NullPlan = "start_up_monthly"
	NULLPLAN_START_UP_YEARLY NullPlan = "start_up_yearly"
	NULLPLAN_CUSTOM NullPlan = "custom"
)

var allowedNullPlanEnumValues = []NullPlan{
	"unknown",
	"free",
	"start_up_monthly",
	"start_up_yearly",
	"custom",
}

func (v *NullPlan) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NullPlan(value)
	for _, existing := range allowedNullPlanEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NullPlan", value)
}

// NewNullPlanFromValue returns a pointer to a valid NullPlan
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNullPlanFromValue(v string) (*NullPlan, error) {
	ev := NullPlan(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NullPlan: valid values are %v", v, allowedNullPlanEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NullPlan) IsValid() bool {
	for _, existing := range allowedNullPlanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NullPlan value
func (v NullPlan) Ptr() *NullPlan {
	return &v
}

type NullableNullPlan struct {
	value *NullPlan
	isSet bool
}

func (v NullableNullPlan) Get() *NullPlan {
	return v.value
}

func (v *NullableNullPlan) Set(val *NullPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableNullPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableNullPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullPlan(val *NullPlan) *NullableNullPlan {
	return &NullableNullPlan{value: val, isSet: true}
}

func (v NullableNullPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

