/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.77
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// ProjectInvite struct for ProjectInvite
type ProjectInvite struct {
	// The Project's Revision Creation Date
	CreatedAt time.Time `json:"created_at"`
	Id string `json:"id"`
	// The invitee's email
	InviteeEmail string `json:"invitee_email"`
	InviteeId *string `json:"invitee_id,omitempty"`
	// The invite owner's email Usually the project's owner email
	OwnerEmail string `json:"owner_email"`
	OwnerId string `json:"owner_id"`
	ProjectId string `json:"project_id"`
	// The invite's status Keeps track of the invites status such as pending, accepted, declined, expired
	Status string `json:"status"`
	// Last Time Project's Revision was Updated
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProjectInvite instantiates a new ProjectInvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectInvite(createdAt time.Time, id string, inviteeEmail string, ownerEmail string, ownerId string, projectId string, status string, updatedAt time.Time) *ProjectInvite {
	this := ProjectInvite{}
	this.CreatedAt = createdAt
	this.Id = id
	this.InviteeEmail = inviteeEmail
	this.OwnerEmail = ownerEmail
	this.OwnerId = ownerId
	this.ProjectId = projectId
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectInviteWithDefaults instantiates a new ProjectInvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectInviteWithDefaults() *ProjectInvite {
	this := ProjectInvite{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectInvite) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectInvite) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ProjectInvite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectInvite) SetId(v string) {
	o.Id = v
}

// GetInviteeEmail returns the InviteeEmail field value
func (o *ProjectInvite) GetInviteeEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteeEmail
}

// GetInviteeEmailOk returns a tuple with the InviteeEmail field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetInviteeEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InviteeEmail, true
}

// SetInviteeEmail sets field value
func (o *ProjectInvite) SetInviteeEmail(v string) {
	o.InviteeEmail = v
}

// GetInviteeId returns the InviteeId field value if set, zero value otherwise.
func (o *ProjectInvite) GetInviteeId() string {
	if o == nil || o.InviteeId == nil {
		var ret string
		return ret
	}
	return *o.InviteeId
}

// GetInviteeIdOk returns a tuple with the InviteeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetInviteeIdOk() (*string, bool) {
	if o == nil || o.InviteeId == nil {
		return nil, false
	}
	return o.InviteeId, true
}

// HasInviteeId returns a boolean if a field has been set.
func (o *ProjectInvite) HasInviteeId() bool {
	if o != nil && o.InviteeId != nil {
		return true
	}

	return false
}

// SetInviteeId gets a reference to the given string and assigns it to the InviteeId field.
func (o *ProjectInvite) SetInviteeId(v string) {
	o.InviteeId = &v
}

// GetOwnerEmail returns the OwnerEmail field value
func (o *ProjectInvite) GetOwnerEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetOwnerEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OwnerEmail, true
}

// SetOwnerEmail sets field value
func (o *ProjectInvite) SetOwnerEmail(v string) {
	o.OwnerEmail = v
}

// GetOwnerId returns the OwnerId field value
func (o *ProjectInvite) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ProjectInvite) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectInvite) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectInvite) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStatus returns the Status field value
func (o *ProjectInvite) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProjectInvite) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectInvite) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectInvite) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectInvite) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ProjectInvite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["invitee_email"] = o.InviteeEmail
	}
	if o.InviteeId != nil {
		toSerialize["invitee_id"] = o.InviteeId
	}
	if true {
		toSerialize["owner_email"] = o.OwnerEmail
	}
	if true {
		toSerialize["owner_id"] = o.OwnerId
	}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectInvite struct {
	value *ProjectInvite
	isSet bool
}

func (v NullableProjectInvite) Get() *ProjectInvite {
	return v.value
}

func (v *NullableProjectInvite) Set(val *ProjectInvite) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectInvite(val *ProjectInvite) *NullableProjectInvite {
	return &NullableProjectInvite{value: val, isSet: true}
}

func (v NullableProjectInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


