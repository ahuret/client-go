/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.31
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SelfServiceRecoveryFlowState The state represents the state of the recovery flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
type SelfServiceRecoveryFlowState string

// List of selfServiceRecoveryFlowState
const (
	SELFSERVICERECOVERYFLOWSTATE_CHOOSE_METHOD SelfServiceRecoveryFlowState = "choose_method"
	SELFSERVICERECOVERYFLOWSTATE_SENT_EMAIL SelfServiceRecoveryFlowState = "sent_email"
	SELFSERVICERECOVERYFLOWSTATE_PASSED_CHALLENGE SelfServiceRecoveryFlowState = "passed_challenge"
)

// All allowed values of SelfServiceRecoveryFlowState enum
var AllowedSelfServiceRecoveryFlowStateEnumValues = []SelfServiceRecoveryFlowState{
	"choose_method",
	"sent_email",
	"passed_challenge",
}

func (v *SelfServiceRecoveryFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfServiceRecoveryFlowState(value)
	for _, existing := range AllowedSelfServiceRecoveryFlowStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfServiceRecoveryFlowState", value)
}

// NewSelfServiceRecoveryFlowStateFromValue returns a pointer to a valid SelfServiceRecoveryFlowState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfServiceRecoveryFlowStateFromValue(v string) (*SelfServiceRecoveryFlowState, error) {
	ev := SelfServiceRecoveryFlowState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SelfServiceRecoveryFlowState: valid values are %v", v, AllowedSelfServiceRecoveryFlowStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfServiceRecoveryFlowState) IsValid() bool {
	for _, existing := range AllowedSelfServiceRecoveryFlowStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to selfServiceRecoveryFlowState value
func (v SelfServiceRecoveryFlowState) Ptr() *SelfServiceRecoveryFlowState {
	return &v
}

type NullableSelfServiceRecoveryFlowState struct {
	value *SelfServiceRecoveryFlowState
	isSet bool
}

func (v NullableSelfServiceRecoveryFlowState) Get() *SelfServiceRecoveryFlowState {
	return v.value
}

func (v *NullableSelfServiceRecoveryFlowState) Set(val *SelfServiceRecoveryFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRecoveryFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRecoveryFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRecoveryFlowState(val *SelfServiceRecoveryFlowState) *NullableSelfServiceRecoveryFlowState {
	return &NullableSelfServiceRecoveryFlowState{value: val, isSet: true}
}

func (v NullableSelfServiceRecoveryFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRecoveryFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

