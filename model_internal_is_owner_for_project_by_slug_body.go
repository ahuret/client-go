/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.9
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InternalIsOwnerForProjectBySlugBody Is Owner For Project By Slug Request Body
type InternalIsOwnerForProjectBySlugBody struct {
	// Namespace is the namespace of the subject.
	Namespace string `json:"namespace"`
	// ProjectScope is the project_id resolved from the API Token.
	ProjectScope *string `json:"project_scope,omitempty"`
	// ProjectSlug is the project's slug.
	ProjectSlug string `json:"project_slug"`
	// Subject is the subject acting (user or API key).
	Subject string `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _InternalIsOwnerForProjectBySlugBody InternalIsOwnerForProjectBySlugBody

// NewInternalIsOwnerForProjectBySlugBody instantiates a new InternalIsOwnerForProjectBySlugBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalIsOwnerForProjectBySlugBody(namespace string, projectSlug string, subject string) *InternalIsOwnerForProjectBySlugBody {
	this := InternalIsOwnerForProjectBySlugBody{}
	this.Namespace = namespace
	this.ProjectSlug = projectSlug
	this.Subject = subject
	return &this
}

// NewInternalIsOwnerForProjectBySlugBodyWithDefaults instantiates a new InternalIsOwnerForProjectBySlugBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalIsOwnerForProjectBySlugBodyWithDefaults() *InternalIsOwnerForProjectBySlugBody {
	this := InternalIsOwnerForProjectBySlugBody{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *InternalIsOwnerForProjectBySlugBody) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *InternalIsOwnerForProjectBySlugBody) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *InternalIsOwnerForProjectBySlugBody) SetNamespace(v string) {
	o.Namespace = v
}

// GetProjectScope returns the ProjectScope field value if set, zero value otherwise.
func (o *InternalIsOwnerForProjectBySlugBody) GetProjectScope() string {
	if o == nil || o.ProjectScope == nil {
		var ret string
		return ret
	}
	return *o.ProjectScope
}

// GetProjectScopeOk returns a tuple with the ProjectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalIsOwnerForProjectBySlugBody) GetProjectScopeOk() (*string, bool) {
	if o == nil || o.ProjectScope == nil {
		return nil, false
	}
	return o.ProjectScope, true
}

// HasProjectScope returns a boolean if a field has been set.
func (o *InternalIsOwnerForProjectBySlugBody) HasProjectScope() bool {
	if o != nil && o.ProjectScope != nil {
		return true
	}

	return false
}

// SetProjectScope gets a reference to the given string and assigns it to the ProjectScope field.
func (o *InternalIsOwnerForProjectBySlugBody) SetProjectScope(v string) {
	o.ProjectScope = &v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *InternalIsOwnerForProjectBySlugBody) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *InternalIsOwnerForProjectBySlugBody) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *InternalIsOwnerForProjectBySlugBody) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetSubject returns the Subject field value
func (o *InternalIsOwnerForProjectBySlugBody) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *InternalIsOwnerForProjectBySlugBody) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *InternalIsOwnerForProjectBySlugBody) SetSubject(v string) {
	o.Subject = v
}

func (o InternalIsOwnerForProjectBySlugBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if o.ProjectScope != nil {
		toSerialize["project_scope"] = o.ProjectScope
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if true {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InternalIsOwnerForProjectBySlugBody) UnmarshalJSON(bytes []byte) (err error) {
	varInternalIsOwnerForProjectBySlugBody := _InternalIsOwnerForProjectBySlugBody{}

	if err = json.Unmarshal(bytes, &varInternalIsOwnerForProjectBySlugBody); err == nil {
		*o = InternalIsOwnerForProjectBySlugBody(varInternalIsOwnerForProjectBySlugBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "project_scope")
		delete(additionalProperties, "project_slug")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalIsOwnerForProjectBySlugBody struct {
	value *InternalIsOwnerForProjectBySlugBody
	isSet bool
}

func (v NullableInternalIsOwnerForProjectBySlugBody) Get() *InternalIsOwnerForProjectBySlugBody {
	return v.value
}

func (v *NullableInternalIsOwnerForProjectBySlugBody) Set(val *InternalIsOwnerForProjectBySlugBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalIsOwnerForProjectBySlugBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalIsOwnerForProjectBySlugBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalIsOwnerForProjectBySlugBody(val *InternalIsOwnerForProjectBySlugBody) *NullableInternalIsOwnerForProjectBySlugBody {
	return &NullableInternalIsOwnerForProjectBySlugBody{value: val, isSet: true}
}

func (v NullableInternalIsOwnerForProjectBySlugBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalIsOwnerForProjectBySlugBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


