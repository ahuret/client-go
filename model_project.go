/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.117
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// Project struct for Project
type Project struct {
	// The Project's Creation Date
	CreatedAt time.Time `json:"created_at"`
	CurrentRevision ProjectRevision `json:"current_revision"`
	Hosts []string `json:"hosts"`
	Id string `json:"id"`
	Revisions []ProjectRevision `json:"revisions"`
	// The project's slug
	Slug string `json:"slug"`
	// The state of the project.
	State string `json:"state"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// Last Time Project was Updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(createdAt time.Time, currentRevision ProjectRevision, hosts []string, id string, revisions []ProjectRevision, slug string, state string, updatedAt time.Time) *Project {
	this := Project{}
	this.CreatedAt = createdAt
	this.CurrentRevision = currentRevision
	this.Hosts = hosts
	this.Id = id
	this.Revisions = revisions
	this.Slug = slug
	this.State = state
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Project) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrentRevision returns the CurrentRevision field value
func (o *Project) GetCurrentRevision() ProjectRevision {
	if o == nil {
		var ret ProjectRevision
		return ret
	}

	return o.CurrentRevision
}

// GetCurrentRevisionOk returns a tuple with the CurrentRevision field value
// and a boolean to check if the value has been set.
func (o *Project) GetCurrentRevisionOk() (*ProjectRevision, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentRevision, true
}

// SetCurrentRevision sets field value
func (o *Project) SetCurrentRevision(v ProjectRevision) {
	o.CurrentRevision = v
}

// GetHosts returns the Hosts field value
func (o *Project) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *Project) GetHostsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *Project) SetHosts(v []string) {
	o.Hosts = v
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

// GetRevisions returns the Revisions field value
func (o *Project) GetRevisions() []ProjectRevision {
	if o == nil {
		var ret []ProjectRevision
		return ret
	}

	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
func (o *Project) GetRevisionsOk() ([]ProjectRevision, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Revisions, true
}

// SetRevisions sets field value
func (o *Project) SetRevisions(v []ProjectRevision) {
	o.Revisions = v
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

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Project) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Project) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Project) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Project) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Project) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Project) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["current_revision"] = o.CurrentRevision
	}
	if true {
		toSerialize["hosts"] = o.Hosts
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["revisions"] = o.Revisions
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.SubscriptionId != nil {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
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


