/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.145
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// ProjectMetadata struct for ProjectMetadata
type ProjectMetadata struct {
	// The Project's Creation Date
	CreatedAt time.Time `json:"created_at"`
	Hosts []string `json:"hosts"`
	Id string `json:"id"`
	// The project's name if set
	Name string `json:"name"`
	// The project's slug
	Slug *string `json:"slug,omitempty"`
	// The state of the project.
	State string `json:"state"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// Last Time Project was Updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProjectMetadata instantiates a new ProjectMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMetadata(createdAt time.Time, hosts []string, id string, name string, state string, updatedAt time.Time) *ProjectMetadata {
	this := ProjectMetadata{}
	this.CreatedAt = createdAt
	this.Hosts = hosts
	this.Id = id
	this.Name = name
	this.State = state
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectMetadataWithDefaults instantiates a new ProjectMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetadataWithDefaults() *ProjectMetadata {
	this := ProjectMetadata{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectMetadata) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetHosts returns the Hosts field value
func (o *ProjectMetadata) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetHostsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *ProjectMetadata) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value
func (o *ProjectMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMetadata) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ProjectMetadata) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ProjectMetadata) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ProjectMetadata) SetSlug(v string) {
	o.Slug = &v
}

// GetState returns the State field value
func (o *ProjectMetadata) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ProjectMetadata) SetState(v string) {
	o.State = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *ProjectMetadata) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *ProjectMetadata) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *ProjectMetadata) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectMetadata) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectMetadata) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectMetadata) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ProjectMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["hosts"] = o.Hosts
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
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

type NullableProjectMetadata struct {
	value *ProjectMetadata
	isSet bool
}

func (v NullableProjectMetadata) Get() *ProjectMetadata {
	return v.value
}

func (v *NullableProjectMetadata) Set(val *ProjectMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMetadata(val *ProjectMetadata) *NullableProjectMetadata {
	return &NullableProjectMetadata{value: val, isSet: true}
}

func (v NullableProjectMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


