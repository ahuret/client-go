/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.36
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IdentityPatch Payload for patching an identity
type IdentityPatch struct {
	Create *CreateIdentityBody `json:"create,omitempty"`
	// The ID of this patch.  The patch ID is optional. If specified, the ID will be returned in the response, so consumers of this API can correlate the response with the patch.
	PatchId *string `json:"patch_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityPatch IdentityPatch

// NewIdentityPatch instantiates a new IdentityPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPatch() *IdentityPatch {
	this := IdentityPatch{}
	return &this
}

// NewIdentityPatchWithDefaults instantiates a new IdentityPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPatchWithDefaults() *IdentityPatch {
	this := IdentityPatch{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *IdentityPatch) GetCreate() CreateIdentityBody {
	if o == nil || o.Create == nil {
		var ret CreateIdentityBody
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPatch) GetCreateOk() (*CreateIdentityBody, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *IdentityPatch) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given CreateIdentityBody and assigns it to the Create field.
func (o *IdentityPatch) SetCreate(v CreateIdentityBody) {
	o.Create = &v
}

// GetPatchId returns the PatchId field value if set, zero value otherwise.
func (o *IdentityPatch) GetPatchId() string {
	if o == nil || o.PatchId == nil {
		var ret string
		return ret
	}
	return *o.PatchId
}

// GetPatchIdOk returns a tuple with the PatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPatch) GetPatchIdOk() (*string, bool) {
	if o == nil || o.PatchId == nil {
		return nil, false
	}
	return o.PatchId, true
}

// HasPatchId returns a boolean if a field has been set.
func (o *IdentityPatch) HasPatchId() bool {
	if o != nil && o.PatchId != nil {
		return true
	}

	return false
}

// SetPatchId gets a reference to the given string and assigns it to the PatchId field.
func (o *IdentityPatch) SetPatchId(v string) {
	o.PatchId = &v
}

func (o IdentityPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.PatchId != nil {
		toSerialize["patch_id"] = o.PatchId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityPatch) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityPatch := _IdentityPatch{}

	if err = json.Unmarshal(bytes, &varIdentityPatch); err == nil {
		*o = IdentityPatch(varIdentityPatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "create")
		delete(additionalProperties, "patch_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityPatch struct {
	value *IdentityPatch
	isSet bool
}

func (v NullableIdentityPatch) Get() *IdentityPatch {
	return v.value
}

func (v *NullableIdentityPatch) Set(val *IdentityPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPatch(val *IdentityPatch) *NullableIdentityPatch {
	return &NullableIdentityPatch{value: val, isSet: true}
}

func (v NullableIdentityPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


