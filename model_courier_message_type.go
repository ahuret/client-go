/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// CourierMessageType It can either be `email` or `phone`
type CourierMessageType string

// List of courierMessageType
const (
	COURIERMESSAGETYPE_EMAIL CourierMessageType = "email"
	COURIERMESSAGETYPE_PHONE CourierMessageType = "phone"
)

// All allowed values of CourierMessageType enum
var AllowedCourierMessageTypeEnumValues = []CourierMessageType{
	"email",
	"phone",
}

func (v *CourierMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CourierMessageType(value)
	for _, existing := range AllowedCourierMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CourierMessageType", value)
}

// NewCourierMessageTypeFromValue returns a pointer to a valid CourierMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCourierMessageTypeFromValue(v string) (*CourierMessageType, error) {
	ev := CourierMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CourierMessageType: valid values are %v", v, AllowedCourierMessageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CourierMessageType) IsValid() bool {
	for _, existing := range AllowedCourierMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to courierMessageType value
func (v CourierMessageType) Ptr() *CourierMessageType {
	return &v
}

type NullableCourierMessageType struct {
	value *CourierMessageType
	isSet bool
}

func (v NullableCourierMessageType) Get() *CourierMessageType {
	return v.value
}

func (v *NullableCourierMessageType) Set(val *CourierMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableCourierMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableCourierMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCourierMessageType(val *CourierMessageType) *NullableCourierMessageType {
	return &NullableCourierMessageType{value: val, isSet: true}
}

func (v NullableCourierMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCourierMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

