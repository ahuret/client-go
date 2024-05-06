/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.6
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// RegistrationFlowState choose_method: ask the user to choose a method (e.g. registration with email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the registration challenge was passed.
type RegistrationFlowState string

// List of registrationFlowState
const (
	REGISTRATIONFLOWSTATE_CHOOSE_METHOD RegistrationFlowState = "choose_method"
	REGISTRATIONFLOWSTATE_SENT_EMAIL RegistrationFlowState = "sent_email"
	REGISTRATIONFLOWSTATE_PASSED_CHALLENGE RegistrationFlowState = "passed_challenge"
)

// All allowed values of RegistrationFlowState enum
var AllowedRegistrationFlowStateEnumValues = []RegistrationFlowState{
	"choose_method",
	"sent_email",
	"passed_challenge",
}

func (v *RegistrationFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistrationFlowState(value)
	for _, existing := range AllowedRegistrationFlowStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistrationFlowState", value)
}

// NewRegistrationFlowStateFromValue returns a pointer to a valid RegistrationFlowState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistrationFlowStateFromValue(v string) (*RegistrationFlowState, error) {
	ev := RegistrationFlowState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistrationFlowState: valid values are %v", v, AllowedRegistrationFlowStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistrationFlowState) IsValid() bool {
	for _, existing := range AllowedRegistrationFlowStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to registrationFlowState value
func (v RegistrationFlowState) Ptr() *RegistrationFlowState {
	return &v
}

type NullableRegistrationFlowState struct {
	value *RegistrationFlowState
	isSet bool
}

func (v NullableRegistrationFlowState) Get() *RegistrationFlowState {
	return v.value
}

func (v *NullableRegistrationFlowState) Set(val *RegistrationFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationFlowState(val *RegistrationFlowState) *NullableRegistrationFlowState {
	return &NullableRegistrationFlowState{value: val, isSet: true}
}

func (v NullableRegistrationFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

