/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.10
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization B2B SSO Organization
type Organization struct {
	Domains []string `json:"domains"`
	// The organization's ID.
	Id string `json:"id"`
	// The organization's human-readable label.
	Label string `json:"label"`
	// The project's ID.
	ProjectId string `json:"project_id"`
	AdditionalProperties map[string]interface{}
}

type _Organization Organization

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(domains []string, id string, label string, projectId string) *Organization {
	this := Organization{}
	this.Domains = domains
	this.Id = id
	this.Label = label
	this.ProjectId = projectId
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetDomains returns the Domains field value
func (o *Organization) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *Organization) GetDomainsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *Organization) SetDomains(v []string) {
	o.Domains = v
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *Organization) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Organization) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Organization) SetLabel(v string) {
	o.Label = v
}

// GetProjectId returns the ProjectId field value
func (o *Organization) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Organization) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Organization) SetProjectId(v string) {
	o.ProjectId = v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domains"] = o.Domains
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Organization) UnmarshalJSON(bytes []byte) (err error) {
	varOrganization := _Organization{}

	err = json.Unmarshal(bytes, &varOrganization)

	if err != nil {
		return err
	}

	*o = Organization(varOrganization)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "domains")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


