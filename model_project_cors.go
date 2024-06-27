/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectCors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCors{}

// ProjectCors struct for ProjectCors
type ProjectCors struct {
	// Whether CORS is enabled for this endpoint.
	Enabled *bool `json:"enabled,omitempty"`
	// The allowed origins. Use `*` to allow all origins. A wildcard can also be used in the subdomain, i.e. `https://_*.example.com` will allow all origins on all subdomains of `example.com`.
	Origins []string `json:"origins,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectCors ProjectCors

// NewProjectCors instantiates a new ProjectCors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCors() *ProjectCors {
	this := ProjectCors{}
	return &this
}

// NewProjectCorsWithDefaults instantiates a new ProjectCors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCorsWithDefaults() *ProjectCors {
	this := ProjectCors{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProjectCors) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCors) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProjectCors) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProjectCors) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *ProjectCors) GetOrigins() []string {
	if o == nil || IsNil(o.Origins) {
		var ret []string
		return ret
	}
	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCors) GetOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *ProjectCors) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given []string and assigns it to the Origins field.
func (o *ProjectCors) SetOrigins(v []string) {
	o.Origins = v
}

func (o ProjectCors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Origins) {
		toSerialize["origins"] = o.Origins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectCors) UnmarshalJSON(data []byte) (err error) {
	varProjectCors := _ProjectCors{}

	err = json.Unmarshal(data, &varProjectCors)

	if err != nil {
		return err
	}

	*o = ProjectCors(varProjectCors)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "origins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectCors struct {
	value *ProjectCors
	isSet bool
}

func (v NullableProjectCors) Get() *ProjectCors {
	return v.value
}

func (v *NullableProjectCors) Set(val *ProjectCors) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCors) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCors(val *ProjectCors) *NullableProjectCors {
	return &NullableProjectCors{value: val, isSet: true}
}

func (v NullableProjectCors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


