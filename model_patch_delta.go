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
)

// PatchDelta struct for PatchDelta
type PatchDelta struct {
	Action *string `json:"action,omitempty"`
	RelationTuple *InternalRelationTuple `json:"relation_tuple,omitempty"`
}

// NewPatchDelta instantiates a new PatchDelta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDelta() *PatchDelta {
	this := PatchDelta{}
	return &this
}

// NewPatchDeltaWithDefaults instantiates a new PatchDelta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDeltaWithDefaults() *PatchDelta {
	this := PatchDelta{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchDelta) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDelta) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchDelta) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PatchDelta) SetAction(v string) {
	o.Action = &v
}

// GetRelationTuple returns the RelationTuple field value if set, zero value otherwise.
func (o *PatchDelta) GetRelationTuple() InternalRelationTuple {
	if o == nil || o.RelationTuple == nil {
		var ret InternalRelationTuple
		return ret
	}
	return *o.RelationTuple
}

// GetRelationTupleOk returns a tuple with the RelationTuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDelta) GetRelationTupleOk() (*InternalRelationTuple, bool) {
	if o == nil || o.RelationTuple == nil {
		return nil, false
	}
	return o.RelationTuple, true
}

// HasRelationTuple returns a boolean if a field has been set.
func (o *PatchDelta) HasRelationTuple() bool {
	if o != nil && o.RelationTuple != nil {
		return true
	}

	return false
}

// SetRelationTuple gets a reference to the given InternalRelationTuple and assigns it to the RelationTuple field.
func (o *PatchDelta) SetRelationTuple(v InternalRelationTuple) {
	o.RelationTuple = &v
}

func (o PatchDelta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.RelationTuple != nil {
		toSerialize["relation_tuple"] = o.RelationTuple
	}
	return json.Marshal(toSerialize)
}

type NullablePatchDelta struct {
	value *PatchDelta
	isSet bool
}

func (v NullablePatchDelta) Get() *PatchDelta {
	return v.value
}

func (v *NullablePatchDelta) Set(val *PatchDelta) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDelta) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDelta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDelta(val *PatchDelta) *NullablePatchDelta {
	return &NullablePatchDelta{value: val, isSet: true}
}

func (v NullablePatchDelta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchDelta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


