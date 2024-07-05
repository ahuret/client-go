/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SetProjectBrandingThemeBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetProjectBrandingThemeBody{}

// SetProjectBrandingThemeBody struct for SetProjectBrandingThemeBody
type SetProjectBrandingThemeBody struct {
	// Favicon Type
	FaviconType *string `json:"favicon_type,omitempty"`
	// Favicon URL
	FaviconUrl *string `json:"favicon_url,omitempty"`
	// Logo type
	LogoType *string `json:"logo_type,omitempty"`
	// Logo URL
	LogoUrl *string `json:"logo_url,omitempty"`
	// Branding name
	Name *string `json:"name,omitempty"`
	Theme *ProjectBrandingColors `json:"theme,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetProjectBrandingThemeBody SetProjectBrandingThemeBody

// NewSetProjectBrandingThemeBody instantiates a new SetProjectBrandingThemeBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetProjectBrandingThemeBody() *SetProjectBrandingThemeBody {
	this := SetProjectBrandingThemeBody{}
	return &this
}

// NewSetProjectBrandingThemeBodyWithDefaults instantiates a new SetProjectBrandingThemeBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetProjectBrandingThemeBodyWithDefaults() *SetProjectBrandingThemeBody {
	this := SetProjectBrandingThemeBody{}
	return &this
}

// GetFaviconType returns the FaviconType field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetFaviconType() string {
	if o == nil || IsNil(o.FaviconType) {
		var ret string
		return ret
	}
	return *o.FaviconType
}

// GetFaviconTypeOk returns a tuple with the FaviconType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetFaviconTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FaviconType) {
		return nil, false
	}
	return o.FaviconType, true
}

// HasFaviconType returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasFaviconType() bool {
	if o != nil && !IsNil(o.FaviconType) {
		return true
	}

	return false
}

// SetFaviconType gets a reference to the given string and assigns it to the FaviconType field.
func (o *SetProjectBrandingThemeBody) SetFaviconType(v string) {
	o.FaviconType = &v
}

// GetFaviconUrl returns the FaviconUrl field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetFaviconUrl() string {
	if o == nil || IsNil(o.FaviconUrl) {
		var ret string
		return ret
	}
	return *o.FaviconUrl
}

// GetFaviconUrlOk returns a tuple with the FaviconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetFaviconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FaviconUrl) {
		return nil, false
	}
	return o.FaviconUrl, true
}

// HasFaviconUrl returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasFaviconUrl() bool {
	if o != nil && !IsNil(o.FaviconUrl) {
		return true
	}

	return false
}

// SetFaviconUrl gets a reference to the given string and assigns it to the FaviconUrl field.
func (o *SetProjectBrandingThemeBody) SetFaviconUrl(v string) {
	o.FaviconUrl = &v
}

// GetLogoType returns the LogoType field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetLogoType() string {
	if o == nil || IsNil(o.LogoType) {
		var ret string
		return ret
	}
	return *o.LogoType
}

// GetLogoTypeOk returns a tuple with the LogoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetLogoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LogoType) {
		return nil, false
	}
	return o.LogoType, true
}

// HasLogoType returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasLogoType() bool {
	if o != nil && !IsNil(o.LogoType) {
		return true
	}

	return false
}

// SetLogoType gets a reference to the given string and assigns it to the LogoType field.
func (o *SetProjectBrandingThemeBody) SetLogoType(v string) {
	o.LogoType = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *SetProjectBrandingThemeBody) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetProjectBrandingThemeBody) SetName(v string) {
	o.Name = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetTheme() ProjectBrandingColors {
	if o == nil || IsNil(o.Theme) {
		var ret ProjectBrandingColors
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetThemeOk() (*ProjectBrandingColors, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given ProjectBrandingColors and assigns it to the Theme field.
func (o *SetProjectBrandingThemeBody) SetTheme(v ProjectBrandingColors) {
	o.Theme = &v
}

func (o SetProjectBrandingThemeBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetProjectBrandingThemeBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FaviconType) {
		toSerialize["favicon_type"] = o.FaviconType
	}
	if !IsNil(o.FaviconUrl) {
		toSerialize["favicon_url"] = o.FaviconUrl
	}
	if !IsNil(o.LogoType) {
		toSerialize["logo_type"] = o.LogoType
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetProjectBrandingThemeBody) UnmarshalJSON(data []byte) (err error) {
	varSetProjectBrandingThemeBody := _SetProjectBrandingThemeBody{}

	err = json.Unmarshal(data, &varSetProjectBrandingThemeBody)

	if err != nil {
		return err
	}

	*o = SetProjectBrandingThemeBody(varSetProjectBrandingThemeBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "favicon_type")
		delete(additionalProperties, "favicon_url")
		delete(additionalProperties, "logo_type")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "theme")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetProjectBrandingThemeBody struct {
	value *SetProjectBrandingThemeBody
	isSet bool
}

func (v NullableSetProjectBrandingThemeBody) Get() *SetProjectBrandingThemeBody {
	return v.value
}

func (v *NullableSetProjectBrandingThemeBody) Set(val *SetProjectBrandingThemeBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetProjectBrandingThemeBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetProjectBrandingThemeBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetProjectBrandingThemeBody(val *SetProjectBrandingThemeBody) *NullableSetProjectBrandingThemeBody {
	return &NullableSetProjectBrandingThemeBody{value: val, isSet: true}
}

func (v NullableSetProjectBrandingThemeBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetProjectBrandingThemeBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


