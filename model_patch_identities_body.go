/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.6
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PatchIdentitiesBody Patch Identities Body
type PatchIdentitiesBody struct {
	// Identities holds the list of patches to apply  required
	Identities []IdentityPatch `json:"identities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchIdentitiesBody PatchIdentitiesBody

// NewPatchIdentitiesBody instantiates a new PatchIdentitiesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchIdentitiesBody() *PatchIdentitiesBody {
	this := PatchIdentitiesBody{}
	return &this
}

// NewPatchIdentitiesBodyWithDefaults instantiates a new PatchIdentitiesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchIdentitiesBodyWithDefaults() *PatchIdentitiesBody {
	this := PatchIdentitiesBody{}
	return &this
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *PatchIdentitiesBody) GetIdentities() []IdentityPatch {
	if o == nil || o.Identities == nil {
		var ret []IdentityPatch
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchIdentitiesBody) GetIdentitiesOk() ([]IdentityPatch, bool) {
	if o == nil || o.Identities == nil {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *PatchIdentitiesBody) HasIdentities() bool {
	if o != nil && o.Identities != nil {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []IdentityPatch and assigns it to the Identities field.
func (o *PatchIdentitiesBody) SetIdentities(v []IdentityPatch) {
	o.Identities = v
}

func (o PatchIdentitiesBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identities != nil {
		toSerialize["identities"] = o.Identities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PatchIdentitiesBody) UnmarshalJSON(bytes []byte) (err error) {
	varPatchIdentitiesBody := _PatchIdentitiesBody{}

	if err = json.Unmarshal(bytes, &varPatchIdentitiesBody); err == nil {
		*o = PatchIdentitiesBody(varPatchIdentitiesBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchIdentitiesBody struct {
	value *PatchIdentitiesBody
	isSet bool
}

func (v NullablePatchIdentitiesBody) Get() *PatchIdentitiesBody {
	return v.value
}

func (v *NullablePatchIdentitiesBody) Set(val *PatchIdentitiesBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchIdentitiesBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchIdentitiesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchIdentitiesBody(val *PatchIdentitiesBody) *NullablePatchIdentitiesBody {
	return &NullablePatchIdentitiesBody{value: val, isSet: true}
}

func (v NullablePatchIdentitiesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchIdentitiesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


