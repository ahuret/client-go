/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ProjectApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectApiKey{}

// ProjectApiKey struct for ProjectApiKey
type ProjectApiKey struct {
	// The token's creation date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The token's ID.
	Id string `json:"id"`
	// The Token's Name  Set this to help you remember, for example, where you use the token.
	Name string `json:"name"`
	// The token's owner
	OwnerId string `json:"owner_id"`
	// The Token's Project ID
	ProjectId *string `json:"project_id,omitempty"`
	// The token's last update date
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The token's value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectApiKey ProjectApiKey

// NewProjectApiKey instantiates a new ProjectApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectApiKey(id string, name string, ownerId string) *ProjectApiKey {
	this := ProjectApiKey{}
	this.Id = id
	this.Name = name
	this.OwnerId = ownerId
	return &this
}

// NewProjectApiKeyWithDefaults instantiates a new ProjectApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectApiKeyWithDefaults() *ProjectApiKey {
	this := ProjectApiKey{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectApiKey) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectApiKey) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectApiKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *ProjectApiKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectApiKey) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectApiKey) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *ProjectApiKey) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ProjectApiKey) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectApiKey) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectApiKey) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ProjectApiKey) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectApiKey) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectApiKey) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectApiKey) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProjectApiKey) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectApiKey) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectApiKey) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProjectApiKey) SetValue(v string) {
	o.Value = &v
}

func (o ProjectApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["owner_id"] = o.OwnerId
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectApiKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"owner_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectApiKey := _ProjectApiKey{}

	err = json.Unmarshal(data, &varProjectApiKey)

	if err != nil {
		return err
	}

	*o = ProjectApiKey(varProjectApiKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_id")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectApiKey struct {
	value *ProjectApiKey
	isSet bool
}

func (v NullableProjectApiKey) Get() *ProjectApiKey {
	return v.value
}

func (v *NullableProjectApiKey) Set(val *ProjectApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectApiKey(val *ProjectApiKey) *NullableProjectApiKey {
	return &NullableProjectApiKey{value: val, isSet: true}
}

func (v NullableProjectApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


