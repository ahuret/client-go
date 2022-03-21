/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.143
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Project struct for Project
type Project struct {
	Id string `json:"id"`
	// The name of the project.
	Name string `json:"name"`
	RevisionId string `json:"revision_id"`
	Services ProjectServices `json:"services"`
	// The project's slug
	Slug string `json:"slug"`
	// The state of the project.
	State string `json:"state"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(id string, name string, revisionId string, services ProjectServices, slug string, state string) *Project {
	this := Project{}
	this.Id = id
	this.Name = name
	this.RevisionId = revisionId
	this.Services = services
	this.Slug = slug
	this.State = state
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetId returns the Id field value
func (o *Project) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Project) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetRevisionId returns the RevisionId field value
func (o *Project) GetRevisionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value
// and a boolean to check if the value has been set.
func (o *Project) GetRevisionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RevisionId, true
}

// SetRevisionId sets field value
func (o *Project) SetRevisionId(v string) {
	o.RevisionId = v
}

// GetServices returns the Services field value
func (o *Project) GetServices() ProjectServices {
	if o == nil {
		var ret ProjectServices
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *Project) GetServicesOk() (*ProjectServices, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *Project) SetServices(v ProjectServices) {
	o.Services = v
}

// GetSlug returns the Slug field value
func (o *Project) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Project) GetSlugOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Project) SetSlug(v string) {
	o.Slug = v
}

// GetState returns the State field value
func (o *Project) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Project) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Project) SetState(v string) {
	o.State = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["revision_id"] = o.RevisionId
	}
	if true {
		toSerialize["services"] = o.Services
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


