/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the IsOwnerForProjectBySlug type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsOwnerForProjectBySlug{}

// IsOwnerForProjectBySlug struct for IsOwnerForProjectBySlug
type IsOwnerForProjectBySlug struct {
	// ProjectSlug is the project's slug.
	ProjectSlug string `json:"ProjectSlug"`
	// Subject is the subject from the API Token.
	Subject string `json:"Subject"`
	AdditionalProperties map[string]interface{}
}

type _IsOwnerForProjectBySlug IsOwnerForProjectBySlug

// NewIsOwnerForProjectBySlug instantiates a new IsOwnerForProjectBySlug object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsOwnerForProjectBySlug(projectSlug string, subject string) *IsOwnerForProjectBySlug {
	this := IsOwnerForProjectBySlug{}
	this.ProjectSlug = projectSlug
	this.Subject = subject
	return &this
}

// NewIsOwnerForProjectBySlugWithDefaults instantiates a new IsOwnerForProjectBySlug object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsOwnerForProjectBySlugWithDefaults() *IsOwnerForProjectBySlug {
	this := IsOwnerForProjectBySlug{}
	return &this
}

// GetProjectSlug returns the ProjectSlug field value
func (o *IsOwnerForProjectBySlug) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *IsOwnerForProjectBySlug) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *IsOwnerForProjectBySlug) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetSubject returns the Subject field value
func (o *IsOwnerForProjectBySlug) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *IsOwnerForProjectBySlug) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *IsOwnerForProjectBySlug) SetSubject(v string) {
	o.Subject = v
}

func (o IsOwnerForProjectBySlug) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsOwnerForProjectBySlug) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProjectSlug"] = o.ProjectSlug
	toSerialize["Subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IsOwnerForProjectBySlug) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectSlug",
		"Subject",
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

	varIsOwnerForProjectBySlug := _IsOwnerForProjectBySlug{}

	err = json.Unmarshal(data, &varIsOwnerForProjectBySlug)

	if err != nil {
		return err
	}

	*o = IsOwnerForProjectBySlug(varIsOwnerForProjectBySlug)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ProjectSlug")
		delete(additionalProperties, "Subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIsOwnerForProjectBySlug struct {
	value *IsOwnerForProjectBySlug
	isSet bool
}

func (v NullableIsOwnerForProjectBySlug) Get() *IsOwnerForProjectBySlug {
	return v.value
}

func (v *NullableIsOwnerForProjectBySlug) Set(val *IsOwnerForProjectBySlug) {
	v.value = val
	v.isSet = true
}

func (v NullableIsOwnerForProjectBySlug) IsSet() bool {
	return v.isSet
}

func (v *NullableIsOwnerForProjectBySlug) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsOwnerForProjectBySlug(val *IsOwnerForProjectBySlug) *NullableIsOwnerForProjectBySlug {
	return &NullableIsOwnerForProjectBySlug{value: val, isSet: true}
}

func (v NullableIsOwnerForProjectBySlug) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsOwnerForProjectBySlug) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


