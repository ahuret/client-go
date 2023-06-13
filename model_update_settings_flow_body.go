/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.35
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UpdateSettingsFlowBody - Update Settings Flow Request Body
type UpdateSettingsFlowBody struct {
	UpdateSettingsFlowWithLookupMethod *UpdateSettingsFlowWithLookupMethod
	UpdateSettingsFlowWithOidcMethod *UpdateSettingsFlowWithOidcMethod
	UpdateSettingsFlowWithPasswordMethod *UpdateSettingsFlowWithPasswordMethod
	UpdateSettingsFlowWithProfileMethod *UpdateSettingsFlowWithProfileMethod
	UpdateSettingsFlowWithTotpMethod *UpdateSettingsFlowWithTotpMethod
	UpdateSettingsFlowWithWebAuthnMethod *UpdateSettingsFlowWithWebAuthnMethod
}

// UpdateSettingsFlowWithLookupMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithLookupMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithLookupMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithLookupMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithLookupMethod: v,
	}
}

// UpdateSettingsFlowWithOidcMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithOidcMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithOidcMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithOidcMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithOidcMethod: v,
	}
}

// UpdateSettingsFlowWithPasswordMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithPasswordMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithPasswordMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithPasswordMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithPasswordMethod: v,
	}
}

// UpdateSettingsFlowWithProfileMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithProfileMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithProfileMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithProfileMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithProfileMethod: v,
	}
}

// UpdateSettingsFlowWithTotpMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithTotpMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithTotpMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithTotpMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithTotpMethod: v,
	}
}

// UpdateSettingsFlowWithWebAuthnMethodAsUpdateSettingsFlowBody is a convenience function that returns UpdateSettingsFlowWithWebAuthnMethod wrapped in UpdateSettingsFlowBody
func UpdateSettingsFlowWithWebAuthnMethodAsUpdateSettingsFlowBody(v *UpdateSettingsFlowWithWebAuthnMethod) UpdateSettingsFlowBody {
	return UpdateSettingsFlowBody{
		UpdateSettingsFlowWithWebAuthnMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateSettingsFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateSettingsFlowWithLookupMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithLookupMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithLookupMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithLookupMethod)
		if string(jsonUpdateSettingsFlowWithLookupMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithLookupMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithLookupMethod = nil
	}

	// try to unmarshal data into UpdateSettingsFlowWithOidcMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithOidcMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithOidcMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithOidcMethod)
		if string(jsonUpdateSettingsFlowWithOidcMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithOidcMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithOidcMethod = nil
	}

	// try to unmarshal data into UpdateSettingsFlowWithPasswordMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithPasswordMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithPasswordMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithPasswordMethod)
		if string(jsonUpdateSettingsFlowWithPasswordMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithPasswordMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithPasswordMethod = nil
	}

	// try to unmarshal data into UpdateSettingsFlowWithProfileMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithProfileMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithProfileMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithProfileMethod)
		if string(jsonUpdateSettingsFlowWithProfileMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithProfileMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithProfileMethod = nil
	}

	// try to unmarshal data into UpdateSettingsFlowWithTotpMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithTotpMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithTotpMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithTotpMethod)
		if string(jsonUpdateSettingsFlowWithTotpMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithTotpMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithTotpMethod = nil
	}

	// try to unmarshal data into UpdateSettingsFlowWithWebAuthnMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateSettingsFlowWithWebAuthnMethod)
	if err == nil {
		jsonUpdateSettingsFlowWithWebAuthnMethod, _ := json.Marshal(dst.UpdateSettingsFlowWithWebAuthnMethod)
		if string(jsonUpdateSettingsFlowWithWebAuthnMethod) == "{}" { // empty struct
			dst.UpdateSettingsFlowWithWebAuthnMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateSettingsFlowWithWebAuthnMethod = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateSettingsFlowWithLookupMethod = nil
		dst.UpdateSettingsFlowWithOidcMethod = nil
		dst.UpdateSettingsFlowWithPasswordMethod = nil
		dst.UpdateSettingsFlowWithProfileMethod = nil
		dst.UpdateSettingsFlowWithTotpMethod = nil
		dst.UpdateSettingsFlowWithWebAuthnMethod = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UpdateSettingsFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateSettingsFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateSettingsFlowBody) MarshalJSON() ([]byte, error) {
	if src.UpdateSettingsFlowWithLookupMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithLookupMethod)
	}

	if src.UpdateSettingsFlowWithOidcMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithOidcMethod)
	}

	if src.UpdateSettingsFlowWithPasswordMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithPasswordMethod)
	}

	if src.UpdateSettingsFlowWithProfileMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithProfileMethod)
	}

	if src.UpdateSettingsFlowWithTotpMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithTotpMethod)
	}

	if src.UpdateSettingsFlowWithWebAuthnMethod != nil {
		return json.Marshal(&src.UpdateSettingsFlowWithWebAuthnMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateSettingsFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateSettingsFlowWithLookupMethod != nil {
		return obj.UpdateSettingsFlowWithLookupMethod
	}

	if obj.UpdateSettingsFlowWithOidcMethod != nil {
		return obj.UpdateSettingsFlowWithOidcMethod
	}

	if obj.UpdateSettingsFlowWithPasswordMethod != nil {
		return obj.UpdateSettingsFlowWithPasswordMethod
	}

	if obj.UpdateSettingsFlowWithProfileMethod != nil {
		return obj.UpdateSettingsFlowWithProfileMethod
	}

	if obj.UpdateSettingsFlowWithTotpMethod != nil {
		return obj.UpdateSettingsFlowWithTotpMethod
	}

	if obj.UpdateSettingsFlowWithWebAuthnMethod != nil {
		return obj.UpdateSettingsFlowWithWebAuthnMethod
	}

	// all schemas are nil
	return nil
}

type NullableUpdateSettingsFlowBody struct {
	value *UpdateSettingsFlowBody
	isSet bool
}

func (v NullableUpdateSettingsFlowBody) Get() *UpdateSettingsFlowBody {
	return v.value
}

func (v *NullableUpdateSettingsFlowBody) Set(val *UpdateSettingsFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingsFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingsFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingsFlowBody(val *UpdateSettingsFlowBody) *NullableUpdateSettingsFlowBody {
	return &NullableUpdateSettingsFlowBody{value: val, isSet: true}
}

func (v NullableUpdateSettingsFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingsFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


