/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.48
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UpdateLoginFlowBody - struct for UpdateLoginFlowBody
type UpdateLoginFlowBody struct {
	UpdateLoginFlowWithLookupSecretMethod *UpdateLoginFlowWithLookupSecretMethod
	UpdateLoginFlowWithOidcMethod *UpdateLoginFlowWithOidcMethod
	UpdateLoginFlowWithPasswordMethod *UpdateLoginFlowWithPasswordMethod
	UpdateLoginFlowWithTotpMethod *UpdateLoginFlowWithTotpMethod
	UpdateLoginFlowWithWebAuthnMethod *UpdateLoginFlowWithWebAuthnMethod
}

// UpdateLoginFlowWithLookupSecretMethodAsUpdateLoginFlowBody is a convenience function that returns UpdateLoginFlowWithLookupSecretMethod wrapped in UpdateLoginFlowBody
func UpdateLoginFlowWithLookupSecretMethodAsUpdateLoginFlowBody(v *UpdateLoginFlowWithLookupSecretMethod) UpdateLoginFlowBody {
	return UpdateLoginFlowBody{
		UpdateLoginFlowWithLookupSecretMethod: v,
	}
}

// UpdateLoginFlowWithOidcMethodAsUpdateLoginFlowBody is a convenience function that returns UpdateLoginFlowWithOidcMethod wrapped in UpdateLoginFlowBody
func UpdateLoginFlowWithOidcMethodAsUpdateLoginFlowBody(v *UpdateLoginFlowWithOidcMethod) UpdateLoginFlowBody {
	return UpdateLoginFlowBody{
		UpdateLoginFlowWithOidcMethod: v,
	}
}

// UpdateLoginFlowWithPasswordMethodAsUpdateLoginFlowBody is a convenience function that returns UpdateLoginFlowWithPasswordMethod wrapped in UpdateLoginFlowBody
func UpdateLoginFlowWithPasswordMethodAsUpdateLoginFlowBody(v *UpdateLoginFlowWithPasswordMethod) UpdateLoginFlowBody {
	return UpdateLoginFlowBody{
		UpdateLoginFlowWithPasswordMethod: v,
	}
}

// UpdateLoginFlowWithTotpMethodAsUpdateLoginFlowBody is a convenience function that returns UpdateLoginFlowWithTotpMethod wrapped in UpdateLoginFlowBody
func UpdateLoginFlowWithTotpMethodAsUpdateLoginFlowBody(v *UpdateLoginFlowWithTotpMethod) UpdateLoginFlowBody {
	return UpdateLoginFlowBody{
		UpdateLoginFlowWithTotpMethod: v,
	}
}

// UpdateLoginFlowWithWebAuthnMethodAsUpdateLoginFlowBody is a convenience function that returns UpdateLoginFlowWithWebAuthnMethod wrapped in UpdateLoginFlowBody
func UpdateLoginFlowWithWebAuthnMethodAsUpdateLoginFlowBody(v *UpdateLoginFlowWithWebAuthnMethod) UpdateLoginFlowBody {
	return UpdateLoginFlowBody{
		UpdateLoginFlowWithWebAuthnMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateLoginFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateLoginFlowWithLookupSecretMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateLoginFlowWithLookupSecretMethod)
	if err == nil {
		jsonUpdateLoginFlowWithLookupSecretMethod, _ := json.Marshal(dst.UpdateLoginFlowWithLookupSecretMethod)
		if string(jsonUpdateLoginFlowWithLookupSecretMethod) == "{}" { // empty struct
			dst.UpdateLoginFlowWithLookupSecretMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateLoginFlowWithLookupSecretMethod = nil
	}

	// try to unmarshal data into UpdateLoginFlowWithOidcMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateLoginFlowWithOidcMethod)
	if err == nil {
		jsonUpdateLoginFlowWithOidcMethod, _ := json.Marshal(dst.UpdateLoginFlowWithOidcMethod)
		if string(jsonUpdateLoginFlowWithOidcMethod) == "{}" { // empty struct
			dst.UpdateLoginFlowWithOidcMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateLoginFlowWithOidcMethod = nil
	}

	// try to unmarshal data into UpdateLoginFlowWithPasswordMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateLoginFlowWithPasswordMethod)
	if err == nil {
		jsonUpdateLoginFlowWithPasswordMethod, _ := json.Marshal(dst.UpdateLoginFlowWithPasswordMethod)
		if string(jsonUpdateLoginFlowWithPasswordMethod) == "{}" { // empty struct
			dst.UpdateLoginFlowWithPasswordMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateLoginFlowWithPasswordMethod = nil
	}

	// try to unmarshal data into UpdateLoginFlowWithTotpMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateLoginFlowWithTotpMethod)
	if err == nil {
		jsonUpdateLoginFlowWithTotpMethod, _ := json.Marshal(dst.UpdateLoginFlowWithTotpMethod)
		if string(jsonUpdateLoginFlowWithTotpMethod) == "{}" { // empty struct
			dst.UpdateLoginFlowWithTotpMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateLoginFlowWithTotpMethod = nil
	}

	// try to unmarshal data into UpdateLoginFlowWithWebAuthnMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateLoginFlowWithWebAuthnMethod)
	if err == nil {
		jsonUpdateLoginFlowWithWebAuthnMethod, _ := json.Marshal(dst.UpdateLoginFlowWithWebAuthnMethod)
		if string(jsonUpdateLoginFlowWithWebAuthnMethod) == "{}" { // empty struct
			dst.UpdateLoginFlowWithWebAuthnMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateLoginFlowWithWebAuthnMethod = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateLoginFlowWithLookupSecretMethod = nil
		dst.UpdateLoginFlowWithOidcMethod = nil
		dst.UpdateLoginFlowWithPasswordMethod = nil
		dst.UpdateLoginFlowWithTotpMethod = nil
		dst.UpdateLoginFlowWithWebAuthnMethod = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UpdateLoginFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateLoginFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateLoginFlowBody) MarshalJSON() ([]byte, error) {
	if src.UpdateLoginFlowWithLookupSecretMethod != nil {
		return json.Marshal(&src.UpdateLoginFlowWithLookupSecretMethod)
	}

	if src.UpdateLoginFlowWithOidcMethod != nil {
		return json.Marshal(&src.UpdateLoginFlowWithOidcMethod)
	}

	if src.UpdateLoginFlowWithPasswordMethod != nil {
		return json.Marshal(&src.UpdateLoginFlowWithPasswordMethod)
	}

	if src.UpdateLoginFlowWithTotpMethod != nil {
		return json.Marshal(&src.UpdateLoginFlowWithTotpMethod)
	}

	if src.UpdateLoginFlowWithWebAuthnMethod != nil {
		return json.Marshal(&src.UpdateLoginFlowWithWebAuthnMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateLoginFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateLoginFlowWithLookupSecretMethod != nil {
		return obj.UpdateLoginFlowWithLookupSecretMethod
	}

	if obj.UpdateLoginFlowWithOidcMethod != nil {
		return obj.UpdateLoginFlowWithOidcMethod
	}

	if obj.UpdateLoginFlowWithPasswordMethod != nil {
		return obj.UpdateLoginFlowWithPasswordMethod
	}

	if obj.UpdateLoginFlowWithTotpMethod != nil {
		return obj.UpdateLoginFlowWithTotpMethod
	}

	if obj.UpdateLoginFlowWithWebAuthnMethod != nil {
		return obj.UpdateLoginFlowWithWebAuthnMethod
	}

	// all schemas are nil
	return nil
}

type NullableUpdateLoginFlowBody struct {
	value *UpdateLoginFlowBody
	isSet bool
}

func (v NullableUpdateLoginFlowBody) Get() *UpdateLoginFlowBody {
	return v.value
}

func (v *NullableUpdateLoginFlowBody) Set(val *UpdateLoginFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowBody(val *UpdateLoginFlowBody) *NullableUpdateLoginFlowBody {
	return &NullableUpdateLoginFlowBody{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


