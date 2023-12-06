/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ListMyWorkspacesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMyWorkspacesResponse{}

// ListMyWorkspacesResponse struct for ListMyWorkspacesResponse
type ListMyWorkspacesResponse struct {
	HasNextPage bool `json:"has_next_page"`
	NextPageToken string `json:"next_page_token"`
	Workspaces []Workspace `json:"workspaces"`
	AdditionalProperties map[string]interface{}
}

type _ListMyWorkspacesResponse ListMyWorkspacesResponse

// NewListMyWorkspacesResponse instantiates a new ListMyWorkspacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMyWorkspacesResponse(hasNextPage bool, nextPageToken string, workspaces []Workspace) *ListMyWorkspacesResponse {
	this := ListMyWorkspacesResponse{}
	this.HasNextPage = hasNextPage
	this.NextPageToken = nextPageToken
	this.Workspaces = workspaces
	return &this
}

// NewListMyWorkspacesResponseWithDefaults instantiates a new ListMyWorkspacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMyWorkspacesResponseWithDefaults() *ListMyWorkspacesResponse {
	this := ListMyWorkspacesResponse{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *ListMyWorkspacesResponse) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *ListMyWorkspacesResponse) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *ListMyWorkspacesResponse) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *ListMyWorkspacesResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ListMyWorkspacesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ListMyWorkspacesResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

// GetWorkspaces returns the Workspaces field value
func (o *ListMyWorkspacesResponse) GetWorkspaces() []Workspace {
	if o == nil {
		var ret []Workspace
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *ListMyWorkspacesResponse) GetWorkspacesOk() ([]Workspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *ListMyWorkspacesResponse) SetWorkspaces(v []Workspace) {
	o.Workspaces = v
}

func (o ListMyWorkspacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMyWorkspacesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_next_page"] = o.HasNextPage
	toSerialize["next_page_token"] = o.NextPageToken
	toSerialize["workspaces"] = o.Workspaces

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListMyWorkspacesResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_next_page",
		"next_page_token",
		"workspaces",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListMyWorkspacesResponse := _ListMyWorkspacesResponse{}

	err = json.Unmarshal(bytes, &varListMyWorkspacesResponse)

	if err != nil {
		return err
	}

	*o = ListMyWorkspacesResponse(varListMyWorkspacesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "has_next_page")
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "workspaces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListMyWorkspacesResponse struct {
	value *ListMyWorkspacesResponse
	isSet bool
}

func (v NullableListMyWorkspacesResponse) Get() *ListMyWorkspacesResponse {
	return v.value
}

func (v *NullableListMyWorkspacesResponse) Set(val *ListMyWorkspacesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMyWorkspacesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMyWorkspacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMyWorkspacesResponse(val *ListMyWorkspacesResponse) *NullableListMyWorkspacesResponse {
	return &NullableListMyWorkspacesResponse{value: val, isSet: true}
}

func (v NullableListMyWorkspacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMyWorkspacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


