/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.45
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Relationship Relationship
type Relationship struct {
	// Namespace of the Relation Tuple
	Namespace string `json:"namespace"`
	// Object of the Relation Tuple
	Object string `json:"object"`
	// Relation of the Relation Tuple
	Relation string `json:"relation"`
	// SubjectID of the Relation Tuple  Either SubjectSet or SubjectID can be provided.
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Relationship Relationship

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship(namespace string, object string, relation string) *Relationship {
	this := Relationship{}
	this.Namespace = namespace
	this.Object = object
	this.Relation = relation
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *Relationship) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *Relationship) SetNamespace(v string) {
	o.Namespace = v
}

// GetObject returns the Object field value
func (o *Relationship) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Relationship) SetObject(v string) {
	o.Object = v
}

// GetRelation returns the Relation field value
func (o *Relationship) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *Relationship) SetRelation(v string) {
	o.Relation = v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *Relationship) GetSubjectId() string {
	if o == nil || o.SubjectId == nil {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetSubjectIdOk() (*string, bool) {
	if o == nil || o.SubjectId == nil {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *Relationship) HasSubjectId() bool {
	if o != nil && o.SubjectId != nil {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *Relationship) SetSubjectId(v string) {
	o.SubjectId = &v
}

// GetSubjectSet returns the SubjectSet field value if set, zero value otherwise.
func (o *Relationship) GetSubjectSet() SubjectSet {
	if o == nil || o.SubjectSet == nil {
		var ret SubjectSet
		return ret
	}
	return *o.SubjectSet
}

// GetSubjectSetOk returns a tuple with the SubjectSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetSubjectSetOk() (*SubjectSet, bool) {
	if o == nil || o.SubjectSet == nil {
		return nil, false
	}
	return o.SubjectSet, true
}

// HasSubjectSet returns a boolean if a field has been set.
func (o *Relationship) HasSubjectSet() bool {
	if o != nil && o.SubjectSet != nil {
		return true
	}

	return false
}

// SetSubjectSet gets a reference to the given SubjectSet and assigns it to the SubjectSet field.
func (o *Relationship) SetSubjectSet(v SubjectSet) {
	o.SubjectSet = &v
}

func (o Relationship) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Relationship) UnmarshalJSON(bytes []byte) (err error) {
	varRelationship := _Relationship{}

	if err = json.Unmarshal(bytes, &varRelationship); err == nil {
		*o = Relationship(varRelationship)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "object")
		delete(additionalProperties, "relation")
		delete(additionalProperties, "subject_id")
		delete(additionalProperties, "subject_set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


