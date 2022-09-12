/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.39
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// NormalizedProjectRevisionIdentitySchema struct for NormalizedProjectRevisionIdentitySchema
type NormalizedProjectRevisionIdentitySchema struct {
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The unique ID of this entry.
	Id *string `json:"id,omitempty"`
	IdentitySchema *ManagedIdentitySchema `json:"identity_schema,omitempty"`
	IdentitySchemaId NullableString `json:"identity_schema_id,omitempty"`
	// The imported (named) ID of the Identity Schema referenced in the Ory Kratos config.
	ImportId *string `json:"import_id,omitempty"`
	// The ImportURL can be used to import an Identity Schema from a bse64 encoded string. In the future, this key also support HTTPS and other sources!  If you import an Ory Kratos configuration, this would be akin to the `identity.schemas.#.url` key.  The configuration will always return the import URL when you fetch it from the API.
	ImportUrl *string `json:"import_url,omitempty"`
	// If true sets the default schema for identities  Only one schema can ever be the default schema. If you try to add two schemas with default to true, the request will fail.
	IsDefault *bool `json:"is_default,omitempty"`
	// Use a preset instead of a custom identity schema.
	Preset *string `json:"preset,omitempty"`
	// The Revision's ID this schema belongs to
	ProjectRevisionId *string `json:"project_revision_id,omitempty"`
	// Last Time Project's Revision was Updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewNormalizedProjectRevisionIdentitySchema instantiates a new NormalizedProjectRevisionIdentitySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionIdentitySchema() *NormalizedProjectRevisionIdentitySchema {
	this := NormalizedProjectRevisionIdentitySchema{}
	return &this
}

// NewNormalizedProjectRevisionIdentitySchemaWithDefaults instantiates a new NormalizedProjectRevisionIdentitySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionIdentitySchemaWithDefaults() *NormalizedProjectRevisionIdentitySchema {
	this := NormalizedProjectRevisionIdentitySchema{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionIdentitySchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalizedProjectRevisionIdentitySchema) SetId(v string) {
	o.Id = &v
}

// GetIdentitySchema returns the IdentitySchema field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetIdentitySchema() ManagedIdentitySchema {
	if o == nil || o.IdentitySchema == nil {
		var ret ManagedIdentitySchema
		return ret
	}
	return *o.IdentitySchema
}

// GetIdentitySchemaOk returns a tuple with the IdentitySchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetIdentitySchemaOk() (*ManagedIdentitySchema, bool) {
	if o == nil || o.IdentitySchema == nil {
		return nil, false
	}
	return o.IdentitySchema, true
}

// HasIdentitySchema returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasIdentitySchema() bool {
	if o != nil && o.IdentitySchema != nil {
		return true
	}

	return false
}

// SetIdentitySchema gets a reference to the given ManagedIdentitySchema and assigns it to the IdentitySchema field.
func (o *NormalizedProjectRevisionIdentitySchema) SetIdentitySchema(v ManagedIdentitySchema) {
	o.IdentitySchema = &v
}

// GetIdentitySchemaId returns the IdentitySchemaId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionIdentitySchema) GetIdentitySchemaId() string {
	if o == nil || o.IdentitySchemaId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdentitySchemaId.Get()
}

// GetIdentitySchemaIdOk returns a tuple with the IdentitySchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionIdentitySchema) GetIdentitySchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentitySchemaId.Get(), o.IdentitySchemaId.IsSet()
}

// HasIdentitySchemaId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasIdentitySchemaId() bool {
	if o != nil && o.IdentitySchemaId.IsSet() {
		return true
	}

	return false
}

// SetIdentitySchemaId gets a reference to the given NullableString and assigns it to the IdentitySchemaId field.
func (o *NormalizedProjectRevisionIdentitySchema) SetIdentitySchemaId(v string) {
	o.IdentitySchemaId.Set(&v)
}
// SetIdentitySchemaIdNil sets the value for IdentitySchemaId to be an explicit nil
func (o *NormalizedProjectRevisionIdentitySchema) SetIdentitySchemaIdNil() {
	o.IdentitySchemaId.Set(nil)
}

// UnsetIdentitySchemaId ensures that no value is present for IdentitySchemaId, not even an explicit nil
func (o *NormalizedProjectRevisionIdentitySchema) UnsetIdentitySchemaId() {
	o.IdentitySchemaId.Unset()
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetImportId() string {
	if o == nil || o.ImportId == nil {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetImportIdOk() (*string, bool) {
	if o == nil || o.ImportId == nil {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *NormalizedProjectRevisionIdentitySchema) SetImportId(v string) {
	o.ImportId = &v
}

// GetImportUrl returns the ImportUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetImportUrl() string {
	if o == nil || o.ImportUrl == nil {
		var ret string
		return ret
	}
	return *o.ImportUrl
}

// GetImportUrlOk returns a tuple with the ImportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetImportUrlOk() (*string, bool) {
	if o == nil || o.ImportUrl == nil {
		return nil, false
	}
	return o.ImportUrl, true
}

// HasImportUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasImportUrl() bool {
	if o != nil && o.ImportUrl != nil {
		return true
	}

	return false
}

// SetImportUrl gets a reference to the given string and assigns it to the ImportUrl field.
func (o *NormalizedProjectRevisionIdentitySchema) SetImportUrl(v string) {
	o.ImportUrl = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *NormalizedProjectRevisionIdentitySchema) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetPreset returns the Preset field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetPreset() string {
	if o == nil || o.Preset == nil {
		var ret string
		return ret
	}
	return *o.Preset
}

// GetPresetOk returns a tuple with the Preset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetPresetOk() (*string, bool) {
	if o == nil || o.Preset == nil {
		return nil, false
	}
	return o.Preset, true
}

// HasPreset returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasPreset() bool {
	if o != nil && o.Preset != nil {
		return true
	}

	return false
}

// SetPreset gets a reference to the given string and assigns it to the Preset field.
func (o *NormalizedProjectRevisionIdentitySchema) SetPreset(v string) {
	o.Preset = &v
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetProjectRevisionId() string {
	if o == nil || o.ProjectRevisionId == nil {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || o.ProjectRevisionId == nil {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasProjectRevisionId() bool {
	if o != nil && o.ProjectRevisionId != nil {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *NormalizedProjectRevisionIdentitySchema) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionIdentitySchema) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionIdentitySchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionIdentitySchema) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionIdentitySchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NormalizedProjectRevisionIdentitySchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentitySchema != nil {
		toSerialize["identity_schema"] = o.IdentitySchema
	}
	if o.IdentitySchemaId.IsSet() {
		toSerialize["identity_schema_id"] = o.IdentitySchemaId.Get()
	}
	if o.ImportId != nil {
		toSerialize["import_id"] = o.ImportId
	}
	if o.ImportUrl != nil {
		toSerialize["import_url"] = o.ImportUrl
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Preset != nil {
		toSerialize["preset"] = o.Preset
	}
	if o.ProjectRevisionId != nil {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableNormalizedProjectRevisionIdentitySchema struct {
	value *NormalizedProjectRevisionIdentitySchema
	isSet bool
}

func (v NullableNormalizedProjectRevisionIdentitySchema) Get() *NormalizedProjectRevisionIdentitySchema {
	return v.value
}

func (v *NullableNormalizedProjectRevisionIdentitySchema) Set(val *NormalizedProjectRevisionIdentitySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionIdentitySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionIdentitySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionIdentitySchema(val *NormalizedProjectRevisionIdentitySchema) *NullableNormalizedProjectRevisionIdentitySchema {
	return &NullableNormalizedProjectRevisionIdentitySchema{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionIdentitySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionIdentitySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


