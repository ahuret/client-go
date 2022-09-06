/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.28
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SubmitSelfServiceRegistrationFlowBody - struct for SubmitSelfServiceRegistrationFlowBody
type SubmitSelfServiceRegistrationFlowBody struct {
	SubmitSelfServiceRegistrationFlowWithOidcMethodBody *SubmitSelfServiceRegistrationFlowWithOidcMethodBody
	SubmitSelfServiceRegistrationFlowWithPasswordMethodBody *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
	SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
}

// SubmitSelfServiceRegistrationFlowWithOidcMethodBodyAsSubmitSelfServiceRegistrationFlowBody is a convenience function that returns SubmitSelfServiceRegistrationFlowWithOidcMethodBody wrapped in SubmitSelfServiceRegistrationFlowBody
func SubmitSelfServiceRegistrationFlowWithOidcMethodBodyAsSubmitSelfServiceRegistrationFlowBody(v *SubmitSelfServiceRegistrationFlowWithOidcMethodBody) SubmitSelfServiceRegistrationFlowBody {
	return SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithOidcMethodBody: v,
	}
}

// SubmitSelfServiceRegistrationFlowWithPasswordMethodBodyAsSubmitSelfServiceRegistrationFlowBody is a convenience function that returns SubmitSelfServiceRegistrationFlowWithPasswordMethodBody wrapped in SubmitSelfServiceRegistrationFlowBody
func SubmitSelfServiceRegistrationFlowWithPasswordMethodBodyAsSubmitSelfServiceRegistrationFlowBody(v *SubmitSelfServiceRegistrationFlowWithPasswordMethodBody) SubmitSelfServiceRegistrationFlowBody {
	return SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithPasswordMethodBody: v,
	}
}

// SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBodyAsSubmitSelfServiceRegistrationFlowBody is a convenience function that returns SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody wrapped in SubmitSelfServiceRegistrationFlowBody
func SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBodyAsSubmitSelfServiceRegistrationFlowBody(v *SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) SubmitSelfServiceRegistrationFlowBody {
	return SubmitSelfServiceRegistrationFlowBody{
		SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceRegistrationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceRegistrationFlowWithOidcMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceRegistrationFlowWithOidcMethodBody)
	if err == nil {
		jsonSubmitSelfServiceRegistrationFlowWithOidcMethodBody, _ := json.Marshal(dst.SubmitSelfServiceRegistrationFlowWithOidcMethodBody)
		if string(jsonSubmitSelfServiceRegistrationFlowWithOidcMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceRegistrationFlowWithOidcMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceRegistrationFlowWithOidcMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody)
	if err == nil {
		jsonSubmitSelfServiceRegistrationFlowWithPasswordMethodBody, _ := json.Marshal(dst.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody)
		if string(jsonSubmitSelfServiceRegistrationFlowWithPasswordMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody = nil
	}

	// try to unmarshal data into SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody)
	if err == nil {
		jsonSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody, _ := json.Marshal(dst.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody)
		if string(jsonSubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody) == "{}" { // empty struct
			dst.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceRegistrationFlowWithOidcMethodBody = nil
		dst.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody = nil
		dst.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceRegistrationFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceRegistrationFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceRegistrationFlowWithOidcMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceRegistrationFlowWithOidcMethodBody)
	}

	if src.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody)
	}

	if src.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody != nil {
		return json.Marshal(&src.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceRegistrationFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceRegistrationFlowWithOidcMethodBody != nil {
		return obj.SubmitSelfServiceRegistrationFlowWithOidcMethodBody
	}

	if obj.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody != nil {
		return obj.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody
	}

	if obj.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody != nil {
		return obj.SubmitSelfServiceRegistrationFlowWithWebAuthnMethodBody
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceRegistrationFlowBody struct {
	value *SubmitSelfServiceRegistrationFlowBody
	isSet bool
}

func (v NullableSubmitSelfServiceRegistrationFlowBody) Get() *SubmitSelfServiceRegistrationFlowBody {
	return v.value
}

func (v *NullableSubmitSelfServiceRegistrationFlowBody) Set(val *SubmitSelfServiceRegistrationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceRegistrationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceRegistrationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceRegistrationFlowBody(val *SubmitSelfServiceRegistrationFlowBody) *NullableSubmitSelfServiceRegistrationFlowBody {
	return &NullableSubmitSelfServiceRegistrationFlowBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceRegistrationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


