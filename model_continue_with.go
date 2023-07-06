/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.41
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ContinueWith - struct for ContinueWith
type ContinueWith struct {
	ContinueWithSetOrySessionToken *ContinueWithSetOrySessionToken
	ContinueWithVerificationUi *ContinueWithVerificationUi
}

// ContinueWithSetOrySessionTokenAsContinueWith is a convenience function that returns ContinueWithSetOrySessionToken wrapped in ContinueWith
func ContinueWithSetOrySessionTokenAsContinueWith(v *ContinueWithSetOrySessionToken) ContinueWith {
	return ContinueWith{
		ContinueWithSetOrySessionToken: v,
	}
}

// ContinueWithVerificationUiAsContinueWith is a convenience function that returns ContinueWithVerificationUi wrapped in ContinueWith
func ContinueWithVerificationUiAsContinueWith(v *ContinueWithVerificationUi) ContinueWith {
	return ContinueWith{
		ContinueWithVerificationUi: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContinueWith) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContinueWithSetOrySessionToken
	err = newStrictDecoder(data).Decode(&dst.ContinueWithSetOrySessionToken)
	if err == nil {
		jsonContinueWithSetOrySessionToken, _ := json.Marshal(dst.ContinueWithSetOrySessionToken)
		if string(jsonContinueWithSetOrySessionToken) == "{}" { // empty struct
			dst.ContinueWithSetOrySessionToken = nil
		} else {
			match++
		}
	} else {
		dst.ContinueWithSetOrySessionToken = nil
	}

	// try to unmarshal data into ContinueWithVerificationUi
	err = newStrictDecoder(data).Decode(&dst.ContinueWithVerificationUi)
	if err == nil {
		jsonContinueWithVerificationUi, _ := json.Marshal(dst.ContinueWithVerificationUi)
		if string(jsonContinueWithVerificationUi) == "{}" { // empty struct
			dst.ContinueWithVerificationUi = nil
		} else {
			match++
		}
	} else {
		dst.ContinueWithVerificationUi = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContinueWithSetOrySessionToken = nil
		dst.ContinueWithVerificationUi = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ContinueWith)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ContinueWith)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContinueWith) MarshalJSON() ([]byte, error) {
	if src.ContinueWithSetOrySessionToken != nil {
		return json.Marshal(&src.ContinueWithSetOrySessionToken)
	}

	if src.ContinueWithVerificationUi != nil {
		return json.Marshal(&src.ContinueWithVerificationUi)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContinueWith) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContinueWithSetOrySessionToken != nil {
		return obj.ContinueWithSetOrySessionToken
	}

	if obj.ContinueWithVerificationUi != nil {
		return obj.ContinueWithVerificationUi
	}

	// all schemas are nil
	return nil
}

type NullableContinueWith struct {
	value *ContinueWith
	isSet bool
}

func (v NullableContinueWith) Get() *ContinueWith {
	return v.value
}

func (v *NullableContinueWith) Set(val *ContinueWith) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWith) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWith) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWith(val *ContinueWith) *NullableContinueWith {
	return &NullableContinueWith{value: val, isSet: true}
}

func (v NullableContinueWith) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWith) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


