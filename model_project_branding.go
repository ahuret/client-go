/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ProjectBranding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectBranding{}

// ProjectBranding struct for ProjectBranding
type ProjectBranding struct {
	// The Customization Creation Date
	CreatedAt time.Time `json:"created_at"`
	DefaultTheme ProjectBrandingTheme `json:"default_theme"`
	// The customization ID.
	Id string `json:"id"`
	// The Project's ID this customization is associated with
	ProjectId string `json:"project_id"`
	Themes []ProjectBrandingTheme `json:"themes"`
	// Last Time Branding was Updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _ProjectBranding ProjectBranding

// NewProjectBranding instantiates a new ProjectBranding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBranding(createdAt time.Time, defaultTheme ProjectBrandingTheme, id string, projectId string, themes []ProjectBrandingTheme, updatedAt time.Time) *ProjectBranding {
	this := ProjectBranding{}
	this.CreatedAt = createdAt
	this.DefaultTheme = defaultTheme
	this.Id = id
	this.ProjectId = projectId
	this.Themes = themes
	this.UpdatedAt = updatedAt
	return &this
}

// NewProjectBrandingWithDefaults instantiates a new ProjectBranding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBrandingWithDefaults() *ProjectBranding {
	this := ProjectBranding{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectBranding) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectBranding) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefaultTheme returns the DefaultTheme field value
func (o *ProjectBranding) GetDefaultTheme() ProjectBrandingTheme {
	if o == nil {
		var ret ProjectBrandingTheme
		return ret
	}

	return o.DefaultTheme
}

// GetDefaultThemeOk returns a tuple with the DefaultTheme field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetDefaultThemeOk() (*ProjectBrandingTheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTheme, true
}

// SetDefaultTheme sets field value
func (o *ProjectBranding) SetDefaultTheme(v ProjectBrandingTheme) {
	o.DefaultTheme = v
}

// GetId returns the Id field value
func (o *ProjectBranding) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectBranding) SetId(v string) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectBranding) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectBranding) SetProjectId(v string) {
	o.ProjectId = v
}

// GetThemes returns the Themes field value
func (o *ProjectBranding) GetThemes() []ProjectBrandingTheme {
	if o == nil {
		var ret []ProjectBrandingTheme
		return ret
	}

	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetThemesOk() ([]ProjectBrandingTheme, bool) {
	if o == nil {
		return nil, false
	}
	return o.Themes, true
}

// SetThemes sets field value
func (o *ProjectBranding) SetThemes(v []ProjectBrandingTheme) {
	o.Themes = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectBranding) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectBranding) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectBranding) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ProjectBranding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectBranding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["default_theme"] = o.DefaultTheme
	toSerialize["id"] = o.Id
	toSerialize["project_id"] = o.ProjectId
	toSerialize["themes"] = o.Themes
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectBranding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"default_theme",
		"id",
		"project_id",
		"themes",
		"updated_at",
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

	varProjectBranding := _ProjectBranding{}

	err = json.Unmarshal(data, &varProjectBranding)

	if err != nil {
		return err
	}

	*o = ProjectBranding(varProjectBranding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "default_theme")
		delete(additionalProperties, "id")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "themes")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectBranding struct {
	value *ProjectBranding
	isSet bool
}

func (v NullableProjectBranding) Get() *ProjectBranding {
	return v.value
}

func (v *NullableProjectBranding) Set(val *ProjectBranding) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBranding) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBranding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBranding(val *ProjectBranding) *NullableProjectBranding {
	return &NullableProjectBranding{value: val, isSet: true}
}

func (v NullableProjectBranding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBranding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


