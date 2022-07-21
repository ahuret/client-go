/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.1.0-alpha.11
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// IdentityState The state can either be `active` or `inactive`.
type IdentityState string

// List of identityState
const (
	IDENTITYSTATE_ACTIVE IdentityState = "active"
	IDENTITYSTATE_INACTIVE IdentityState = "inactive"
)

var allowedIdentityStateEnumValues = []IdentityState{
	"active",
	"inactive",
}

func (v *IdentityState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityState(value)
	for _, existing := range allowedIdentityStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityState", value)
}

// NewIdentityStateFromValue returns a pointer to a valid IdentityState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityStateFromValue(v string) (*IdentityState, error) {
	ev := IdentityState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityState: valid values are %v", v, allowedIdentityStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityState) IsValid() bool {
	for _, existing := range allowedIdentityStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to identityState value
func (v IdentityState) Ptr() *IdentityState {
	return &v
}

type NullableIdentityState struct {
	value *IdentityState
	isSet bool
}

func (v NullableIdentityState) Get() *IdentityState {
	return v.value
}

func (v *NullableIdentityState) Set(val *IdentityState) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityState) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityState(val *IdentityState) *NullableIdentityState {
	return &NullableIdentityState{value: val, isSet: true}
}

func (v NullableIdentityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

