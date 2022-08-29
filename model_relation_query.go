/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: latest
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RelationQuery struct for RelationQuery
type RelationQuery struct {
	// Namespace of the Relation Tuple
	Namespace *string `json:"namespace,omitempty"`
	// Object of the Relation Tuple
	Object *string `json:"object,omitempty"`
	// Relation of the Relation Tuple
	Relation *string `json:"relation,omitempty"`
	// SubjectID of the Relation Tuple  Either SubjectSet or SubjectID can be provided.
	SubjectId *string `json:"subject_id,omitempty"`
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// NewRelationQuery instantiates a new RelationQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationQuery() *RelationQuery {
	this := RelationQuery{}
	return &this
}

// NewRelationQueryWithDefaults instantiates a new RelationQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationQueryWithDefaults() *RelationQuery {
	this := RelationQuery{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RelationQuery) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationQuery) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RelationQuery) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RelationQuery) SetNamespace(v string) {
	o.Namespace = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RelationQuery) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationQuery) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RelationQuery) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *RelationQuery) SetObject(v string) {
	o.Object = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *RelationQuery) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationQuery) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *RelationQuery) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *RelationQuery) SetRelation(v string) {
	o.Relation = &v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *RelationQuery) GetSubjectId() string {
	if o == nil || o.SubjectId == nil {
		var ret string
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationQuery) GetSubjectIdOk() (*string, bool) {
	if o == nil || o.SubjectId == nil {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *RelationQuery) HasSubjectId() bool {
	if o != nil && o.SubjectId != nil {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given string and assigns it to the SubjectId field.
func (o *RelationQuery) SetSubjectId(v string) {
	o.SubjectId = &v
}

// GetSubjectSet returns the SubjectSet field value if set, zero value otherwise.
func (o *RelationQuery) GetSubjectSet() SubjectSet {
	if o == nil || o.SubjectSet == nil {
		var ret SubjectSet
		return ret
	}
	return *o.SubjectSet
}

// GetSubjectSetOk returns a tuple with the SubjectSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationQuery) GetSubjectSetOk() (*SubjectSet, bool) {
	if o == nil || o.SubjectSet == nil {
		return nil, false
	}
	return o.SubjectSet, true
}

// HasSubjectSet returns a boolean if a field has been set.
func (o *RelationQuery) HasSubjectSet() bool {
	if o != nil && o.SubjectSet != nil {
		return true
	}

	return false
}

// SetSubjectSet gets a reference to the given SubjectSet and assigns it to the SubjectSet field.
func (o *RelationQuery) SetSubjectSet(v SubjectSet) {
	o.SubjectSet = &v
}

func (o RelationQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Relation != nil {
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

type NullableRelationQuery struct {
	value *RelationQuery
	isSet bool
}

func (v NullableRelationQuery) Get() *RelationQuery {
	return v.value
}

func (v *NullableRelationQuery) Set(val *RelationQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationQuery(val *RelationQuery) *NullableRelationQuery {
	return &NullableRelationQuery{value: val, isSet: true}
}

func (v NullableRelationQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


