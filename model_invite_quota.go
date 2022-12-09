/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InviteQuota Project invite quota
type InviteQuota struct {
	ProjectId *string `json:"project_id,omitempty"`
	RemainingSeats *int64 `json:"remaining_seats,omitempty"`
	TotalSeats *int64 `json:"total_seats,omitempty"`
}

// NewInviteQuota instantiates a new InviteQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteQuota() *InviteQuota {
	this := InviteQuota{}
	return &this
}

// NewInviteQuotaWithDefaults instantiates a new InviteQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteQuotaWithDefaults() *InviteQuota {
	this := InviteQuota{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *InviteQuota) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteQuota) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *InviteQuota) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *InviteQuota) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRemainingSeats returns the RemainingSeats field value if set, zero value otherwise.
func (o *InviteQuota) GetRemainingSeats() int64 {
	if o == nil || o.RemainingSeats == nil {
		var ret int64
		return ret
	}
	return *o.RemainingSeats
}

// GetRemainingSeatsOk returns a tuple with the RemainingSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteQuota) GetRemainingSeatsOk() (*int64, bool) {
	if o == nil || o.RemainingSeats == nil {
		return nil, false
	}
	return o.RemainingSeats, true
}

// HasRemainingSeats returns a boolean if a field has been set.
func (o *InviteQuota) HasRemainingSeats() bool {
	if o != nil && o.RemainingSeats != nil {
		return true
	}

	return false
}

// SetRemainingSeats gets a reference to the given int64 and assigns it to the RemainingSeats field.
func (o *InviteQuota) SetRemainingSeats(v int64) {
	o.RemainingSeats = &v
}

// GetTotalSeats returns the TotalSeats field value if set, zero value otherwise.
func (o *InviteQuota) GetTotalSeats() int64 {
	if o == nil || o.TotalSeats == nil {
		var ret int64
		return ret
	}
	return *o.TotalSeats
}

// GetTotalSeatsOk returns a tuple with the TotalSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteQuota) GetTotalSeatsOk() (*int64, bool) {
	if o == nil || o.TotalSeats == nil {
		return nil, false
	}
	return o.TotalSeats, true
}

// HasTotalSeats returns a boolean if a field has been set.
func (o *InviteQuota) HasTotalSeats() bool {
	if o != nil && o.TotalSeats != nil {
		return true
	}

	return false
}

// SetTotalSeats gets a reference to the given int64 and assigns it to the TotalSeats field.
func (o *InviteQuota) SetTotalSeats(v int64) {
	o.TotalSeats = &v
}

func (o InviteQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.RemainingSeats != nil {
		toSerialize["remaining_seats"] = o.RemainingSeats
	}
	if o.TotalSeats != nil {
		toSerialize["total_seats"] = o.TotalSeats
	}
	return json.Marshal(toSerialize)
}

type NullableInviteQuota struct {
	value *InviteQuota
	isSet bool
}

func (v NullableInviteQuota) Get() *InviteQuota {
	return v.value
}

func (v *NullableInviteQuota) Set(val *InviteQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteQuota(val *InviteQuota) *NullableInviteQuota {
	return &NullableInviteQuota{value: val, isSet: true}
}

func (v NullableInviteQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


