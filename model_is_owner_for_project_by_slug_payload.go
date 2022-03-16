/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.130
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// IsOwnerForProjectBySlugPayload struct for IsOwnerForProjectBySlugPayload
type IsOwnerForProjectBySlugPayload struct {
	// ProjectScope is the project_id resolved from the API Token.
	ProjectScope *string `json:"project_scope,omitempty"`
	// ProjectSlug is the project's slug.
	ProjectSlug string `json:"project_slug"`
	// Subject is the subject from the API Token.
	Subject string `json:"subject"`
}

// NewIsOwnerForProjectBySlugPayload instantiates a new IsOwnerForProjectBySlugPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsOwnerForProjectBySlugPayload(projectSlug string, subject string) *IsOwnerForProjectBySlugPayload {
	this := IsOwnerForProjectBySlugPayload{}
	this.ProjectSlug = projectSlug
	this.Subject = subject
	return &this
}

// NewIsOwnerForProjectBySlugPayloadWithDefaults instantiates a new IsOwnerForProjectBySlugPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsOwnerForProjectBySlugPayloadWithDefaults() *IsOwnerForProjectBySlugPayload {
	this := IsOwnerForProjectBySlugPayload{}
	return &this
}

// GetProjectScope returns the ProjectScope field value if set, zero value otherwise.
func (o *IsOwnerForProjectBySlugPayload) GetProjectScope() string {
	if o == nil || o.ProjectScope == nil {
		var ret string
		return ret
	}
	return *o.ProjectScope
}

// GetProjectScopeOk returns a tuple with the ProjectScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsOwnerForProjectBySlugPayload) GetProjectScopeOk() (*string, bool) {
	if o == nil || o.ProjectScope == nil {
		return nil, false
	}
	return o.ProjectScope, true
}

// HasProjectScope returns a boolean if a field has been set.
func (o *IsOwnerForProjectBySlugPayload) HasProjectScope() bool {
	if o != nil && o.ProjectScope != nil {
		return true
	}

	return false
}

// SetProjectScope gets a reference to the given string and assigns it to the ProjectScope field.
func (o *IsOwnerForProjectBySlugPayload) SetProjectScope(v string) {
	o.ProjectScope = &v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *IsOwnerForProjectBySlugPayload) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *IsOwnerForProjectBySlugPayload) GetProjectSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *IsOwnerForProjectBySlugPayload) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetSubject returns the Subject field value
func (o *IsOwnerForProjectBySlugPayload) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *IsOwnerForProjectBySlugPayload) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *IsOwnerForProjectBySlugPayload) SetSubject(v string) {
	o.Subject = v
}

func (o IsOwnerForProjectBySlugPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectScope != nil {
		toSerialize["project_scope"] = o.ProjectScope
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableIsOwnerForProjectBySlugPayload struct {
	value *IsOwnerForProjectBySlugPayload
	isSet bool
}

func (v NullableIsOwnerForProjectBySlugPayload) Get() *IsOwnerForProjectBySlugPayload {
	return v.value
}

func (v *NullableIsOwnerForProjectBySlugPayload) Set(val *IsOwnerForProjectBySlugPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableIsOwnerForProjectBySlugPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableIsOwnerForProjectBySlugPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsOwnerForProjectBySlugPayload(val *IsOwnerForProjectBySlugPayload) *NullableIsOwnerForProjectBySlugPayload {
	return &NullableIsOwnerForProjectBySlugPayload{value: val, isSet: true}
}

func (v NullableIsOwnerForProjectBySlugPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsOwnerForProjectBySlugPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


