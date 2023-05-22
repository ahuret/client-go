/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.32
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetProject struct for SetProject
type SetProject struct {
	// The name of the project.
	Name string `json:"name"`
	Services ProjectServices `json:"services"`
}

// NewSetProject instantiates a new SetProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetProject(name string, services ProjectServices) *SetProject {
	this := SetProject{}
	this.Name = name
	this.Services = services
	return &this
}

// NewSetProjectWithDefaults instantiates a new SetProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetProjectWithDefaults() *SetProject {
	this := SetProject{}
	return &this
}

// GetName returns the Name field value
func (o *SetProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SetProject) SetName(v string) {
	o.Name = v
}

// GetServices returns the Services field value
func (o *SetProject) GetServices() ProjectServices {
	if o == nil {
		var ret ProjectServices
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *SetProject) GetServicesOk() (*ProjectServices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *SetProject) SetServices(v ProjectServices) {
	o.Services = v
}

func (o SetProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableSetProject struct {
	value *SetProject
	isSet bool
}

func (v NullableSetProject) Get() *SetProject {
	return v.value
}

func (v *NullableSetProject) Set(val *SetProject) {
	v.value = val
	v.isSet = true
}

func (v NullableSetProject) IsSet() bool {
	return v.isSet
}

func (v *NullableSetProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetProject(val *SetProject) *NullableSetProject {
	return &NullableSetProject{value: val, isSet: true}
}

func (v NullableSetProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


