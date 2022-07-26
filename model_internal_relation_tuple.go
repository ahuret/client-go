/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InternalRelationTuple struct for InternalRelationTuple
type InternalRelationTuple struct {
	// Namespace of the Relation Tuple
	Namespace string `json:"namespace"`
	// Object of the Relation Tuple
	Object string `json:"object"`
	// Relation of the Relation Tuple
	Relation string `json:"relation"`
	// SubjectID of the Relation Tuple  Either SubjectSet or SubjectID are required.
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// NewInternalRelationTuple instantiates a new InternalRelationTuple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalRelationTuple(namespace string, object string, relation string) *InternalRelationTuple {
	this := InternalRelationTuple{}
	this.Namespace = namespace
	this.Object = object
	this.Relation = relation
	return &this
}

// NewInternalRelationTupleWithDefaults instantiates a new InternalRelationTuple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalRelationTupleWithDefaults() *InternalRelationTuple {
	this := InternalRelationTuple{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *InternalRelationTuple) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *InternalRelationTuple) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *InternalRelationTuple) SetNamespace(v string) {
	o.Namespace = v
}

// GetObject returns the Object field value
func (o *InternalRelationTuple) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *InternalRelationTuple) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *InternalRelationTuple) SetObject(v string) {
	o.Object = v
}

// GetRelation returns the Relation field value
func (o *InternalRelationTuple) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *InternalRelationTuple) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *InternalRelationTuple) SetRelation(v string) {
	o.Relation = v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *InternalRelationTuple) GetSubjectId() string {
	if o == nil || o.SubjectId == nil {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalRelationTuple) GetSubjectIdOk() (*string, bool) {
	if o == nil || o.SubjectId == nil {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *InternalRelationTuple) HasSubjectId() bool {
	if o != nil && o.SubjectId != nil {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *InternalRelationTuple) SetSubjectId(v string) {
	o.SubjectId = &v
}

// GetSubjectSet returns the SubjectSet field value if set, zero value otherwise.
func (o *InternalRelationTuple) GetSubjectSet() SubjectSet {
	if o == nil || o.SubjectSet == nil {
		var ret SubjectSet
		return ret
	}
	return *o.SubjectSet
}

// GetSubjectSetOk returns a tuple with the SubjectSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalRelationTuple) GetSubjectSetOk() (*SubjectSet, bool) {
	if o == nil || o.SubjectSet == nil {
		return nil, false
	}
	return o.SubjectSet, true
}

// HasSubjectSet returns a boolean if a field has been set.
func (o *InternalRelationTuple) HasSubjectSet() bool {
	if o != nil && o.SubjectSet != nil {
		return true
	}

	return false
}

// SetSubjectSet gets a reference to the given SubjectSet and assigns it to the SubjectSet field.
func (o *InternalRelationTuple) SetSubjectSet(v SubjectSet) {
	o.SubjectSet = &v
}

func (o InternalRelationTuple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["object"] = o.Object
	}
	if true {
		toSerialize["relation"] = o.Relation
	}
	if o.SubjectId != nil {
		toSerialize["subject_id"] = o.SubjectId
	}
	if o.SubjectSet != nil {
		toSerialize["subject_set"] = o.SubjectSet
	}
	return json.Marshal(toSerialize)
}

type NullableInternalRelationTuple struct {
	value *InternalRelationTuple
	isSet bool
}

func (v NullableInternalRelationTuple) Get() *InternalRelationTuple {
	return v.value
}

func (v *NullableInternalRelationTuple) Set(val *InternalRelationTuple) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalRelationTuple) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalRelationTuple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalRelationTuple(val *InternalRelationTuple) *NullableInternalRelationTuple {
	return &NullableInternalRelationTuple{value: val, isSet: true}
}

func (v NullableInternalRelationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalRelationTuple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


