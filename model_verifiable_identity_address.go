/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.21
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// VerifiableIdentityAddress VerifiableAddress is an identity's verifiable address
type VerifiableIdentityAddress struct {
	// When this entry was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id string `json:"id"`
	// VerifiableAddressStatus must not exceed 16 characters as that is the limitation in the SQL Schema
	Status string `json:"status"`
	// When this entry was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The address value  example foo@user.com
	Value string `json:"value"`
	// Indicates if the address has already been verified
	Verified bool `json:"verified"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
	// VerifiableAddressType must not exceed 16 characters as that is the limitation in the SQL Schema
	Via string `json:"via"`
}

// NewVerifiableIdentityAddress instantiates a new VerifiableIdentityAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiableIdentityAddress(id string, status string, value string, verified bool, via string) *VerifiableIdentityAddress {
	this := VerifiableIdentityAddress{}
	this.Id = id
	this.Status = status
	this.Value = value
	this.Verified = verified
	this.Via = via
	return &this
}

// NewVerifiableIdentityAddressWithDefaults instantiates a new VerifiableIdentityAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiableIdentityAddressWithDefaults() *VerifiableIdentityAddress {
	this := VerifiableIdentityAddress{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VerifiableIdentityAddress) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VerifiableIdentityAddress) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VerifiableIdentityAddress) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *VerifiableIdentityAddress) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VerifiableIdentityAddress) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *VerifiableIdentityAddress) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VerifiableIdentityAddress) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VerifiableIdentityAddress) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VerifiableIdentityAddress) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VerifiableIdentityAddress) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value
func (o *VerifiableIdentityAddress) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VerifiableIdentityAddress) SetValue(v string) {
	o.Value = v
}

// GetVerified returns the Verified field value
func (o *VerifiableIdentityAddress) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetVerifiedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *VerifiableIdentityAddress) SetVerified(v bool) {
	o.Verified = v
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise.
func (o *VerifiableIdentityAddress) GetVerifiedAt() time.Time {
	if o == nil || o.VerifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.VerifiedAt
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetVerifiedAtOk() (*time.Time, bool) {
	if o == nil || o.VerifiedAt == nil {
		return nil, false
	}
	return o.VerifiedAt, true
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *VerifiableIdentityAddress) HasVerifiedAt() bool {
	if o != nil && o.VerifiedAt != nil {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given time.Time and assigns it to the VerifiedAt field.
func (o *VerifiableIdentityAddress) SetVerifiedAt(v time.Time) {
	o.VerifiedAt = &v
}

// GetVia returns the Via field value
func (o *VerifiableIdentityAddress) GetVia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Via
}

// GetViaOk returns a tuple with the Via field value
// and a boolean to check if the value has been set.
func (o *VerifiableIdentityAddress) GetViaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Via, true
}

// SetVia sets field value
func (o *VerifiableIdentityAddress) SetVia(v string) {
	o.Via = v
}

func (o VerifiableIdentityAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if o.VerifiedAt != nil {
		toSerialize["verified_at"] = o.VerifiedAt
	}
	if true {
		toSerialize["via"] = o.Via
	}
	return json.Marshal(toSerialize)
}

type NullableVerifiableIdentityAddress struct {
	value *VerifiableIdentityAddress
	isSet bool
}

func (v NullableVerifiableIdentityAddress) Get() *VerifiableIdentityAddress {
	return v.value
}

func (v *NullableVerifiableIdentityAddress) Set(val *VerifiableIdentityAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiableIdentityAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiableIdentityAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiableIdentityAddress(val *VerifiableIdentityAddress) *NullableVerifiableIdentityAddress {
	return &NullableVerifiableIdentityAddress{value: val, isSet: true}
}

func (v NullableVerifiableIdentityAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiableIdentityAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


