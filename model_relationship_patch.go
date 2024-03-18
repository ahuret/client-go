/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.9.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the RelationshipPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelationshipPatch{}

// RelationshipPatch Payload for patching a relationship
type RelationshipPatch struct {
	Action *string `json:"action,omitempty"`
	RelationTuple *Relationship `json:"relation_tuple,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RelationshipPatch RelationshipPatch

// NewRelationshipPatch instantiates a new RelationshipPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipPatch() *RelationshipPatch {
	this := RelationshipPatch{}
	return &this
}

// NewRelationshipPatchWithDefaults instantiates a new RelationshipPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipPatchWithDefaults() *RelationshipPatch {
	this := RelationshipPatch{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RelationshipPatch) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipPatch) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RelationshipPatch) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RelationshipPatch) SetAction(v string) {
	o.Action = &v
}

// GetRelationTuple returns the RelationTuple field value if set, zero value otherwise.
func (o *RelationshipPatch) GetRelationTuple() Relationship {
	if o == nil || IsNil(o.RelationTuple) {
		var ret Relationship
		return ret
	}
	return *o.RelationTuple
}

// GetRelationTupleOk returns a tuple with the RelationTuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipPatch) GetRelationTupleOk() (*Relationship, bool) {
	if o == nil || IsNil(o.RelationTuple) {
		return nil, false
	}
	return o.RelationTuple, true
}

// HasRelationTuple returns a boolean if a field has been set.
func (o *RelationshipPatch) HasRelationTuple() bool {
	if o != nil && !IsNil(o.RelationTuple) {
		return true
	}

	return false
}

// SetRelationTuple gets a reference to the given Relationship and assigns it to the RelationTuple field.
func (o *RelationshipPatch) SetRelationTuple(v Relationship) {
	o.RelationTuple = &v
}

func (o RelationshipPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelationshipPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.RelationTuple) {
		toSerialize["relation_tuple"] = o.RelationTuple
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RelationshipPatch) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipPatch := _RelationshipPatch{}

	err = json.Unmarshal(bytes, &varRelationshipPatch)

	if err != nil {
		return err
	}

	*o = RelationshipPatch(varRelationshipPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "relation_tuple")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelationshipPatch struct {
	value *RelationshipPatch
	isSet bool
}

func (v NullableRelationshipPatch) Get() *RelationshipPatch {
	return v.value
}

func (v *NullableRelationshipPatch) Set(val *RelationshipPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipPatch(val *RelationshipPatch) *NullableRelationshipPatch {
	return &NullableRelationshipPatch{value: val, isSet: true}
}

func (v NullableRelationshipPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


