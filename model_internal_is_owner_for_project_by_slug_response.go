/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.9
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InternalIsOwnerForProjectBySlugResponse struct for InternalIsOwnerForProjectBySlugResponse
type InternalIsOwnerForProjectBySlugResponse struct {
	// ProjectID is the project's ID.
	ProjectId string `json:"project_id"`
	AdditionalProperties map[string]interface{}
}

type _InternalIsOwnerForProjectBySlugResponse InternalIsOwnerForProjectBySlugResponse

// NewInternalIsOwnerForProjectBySlugResponse instantiates a new InternalIsOwnerForProjectBySlugResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalIsOwnerForProjectBySlugResponse(projectId string) *InternalIsOwnerForProjectBySlugResponse {
	this := InternalIsOwnerForProjectBySlugResponse{}
	this.ProjectId = projectId
	return &this
}

// NewInternalIsOwnerForProjectBySlugResponseWithDefaults instantiates a new InternalIsOwnerForProjectBySlugResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalIsOwnerForProjectBySlugResponseWithDefaults() *InternalIsOwnerForProjectBySlugResponse {
	this := InternalIsOwnerForProjectBySlugResponse{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *InternalIsOwnerForProjectBySlugResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *InternalIsOwnerForProjectBySlugResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *InternalIsOwnerForProjectBySlugResponse) SetProjectId(v string) {
	o.ProjectId = v
}

func (o InternalIsOwnerForProjectBySlugResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InternalIsOwnerForProjectBySlugResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInternalIsOwnerForProjectBySlugResponse := _InternalIsOwnerForProjectBySlugResponse{}

	if err = json.Unmarshal(bytes, &varInternalIsOwnerForProjectBySlugResponse); err == nil {
		*o = InternalIsOwnerForProjectBySlugResponse(varInternalIsOwnerForProjectBySlugResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInternalIsOwnerForProjectBySlugResponse struct {
	value *InternalIsOwnerForProjectBySlugResponse
	isSet bool
}

func (v NullableInternalIsOwnerForProjectBySlugResponse) Get() *InternalIsOwnerForProjectBySlugResponse {
	return v.value
}

func (v *NullableInternalIsOwnerForProjectBySlugResponse) Set(val *InternalIsOwnerForProjectBySlugResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalIsOwnerForProjectBySlugResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalIsOwnerForProjectBySlugResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalIsOwnerForProjectBySlugResponse(val *InternalIsOwnerForProjectBySlugResponse) *NullableInternalIsOwnerForProjectBySlugResponse {
	return &NullableInternalIsOwnerForProjectBySlugResponse{value: val, isSet: true}
}

func (v NullableInternalIsOwnerForProjectBySlugResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalIsOwnerForProjectBySlugResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


