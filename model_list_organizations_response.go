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
	"fmt"
)

// checks if the ListOrganizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationsResponse{}

// ListOrganizationsResponse B2B SSO Organization List
type ListOrganizationsResponse struct {
	HasNextPage bool `json:"has_next_page"`
	NextPageToken string `json:"next_page_token"`
	// The list of organizations
	Organizations []Organization `json:"organizations"`
	AdditionalProperties map[string]interface{}
}

type _ListOrganizationsResponse ListOrganizationsResponse

// NewListOrganizationsResponse instantiates a new ListOrganizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationsResponse(hasNextPage bool, nextPageToken string, organizations []Organization) *ListOrganizationsResponse {
	this := ListOrganizationsResponse{}
	this.HasNextPage = hasNextPage
	this.NextPageToken = nextPageToken
	this.Organizations = organizations
	return &this
}

// NewListOrganizationsResponseWithDefaults instantiates a new ListOrganizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationsResponseWithDefaults() *ListOrganizationsResponse {
	this := ListOrganizationsResponse{}
	return &this
}

// GetHasNextPage returns the HasNextPage field value
func (o *ListOrganizationsResponse) GetHasNextPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNextPage
}

// GetHasNextPageOk returns a tuple with the HasNextPage field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponse) GetHasNextPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNextPage, true
}

// SetHasNextPage sets field value
func (o *ListOrganizationsResponse) SetHasNextPage(v bool) {
	o.HasNextPage = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *ListOrganizationsResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ListOrganizationsResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

// GetOrganizations returns the Organizations field value
func (o *ListOrganizationsResponse) GetOrganizations() []Organization {
	if o == nil {
		var ret []Organization
		return ret
	}

	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationsResponse) GetOrganizationsOk() ([]Organization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organizations, true
}

// SetOrganizations sets field value
func (o *ListOrganizationsResponse) SetOrganizations(v []Organization) {
	o.Organizations = v
}

func (o ListOrganizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_next_page"] = o.HasNextPage
	toSerialize["next_page_token"] = o.NextPageToken
	toSerialize["organizations"] = o.Organizations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListOrganizationsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"has_next_page",
		"next_page_token",
		"organizations",
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

	varListOrganizationsResponse := _ListOrganizationsResponse{}

	err = json.Unmarshal(data, &varListOrganizationsResponse)

	if err != nil {
		return err
	}

	*o = ListOrganizationsResponse(varListOrganizationsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "has_next_page")
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "organizations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListOrganizationsResponse struct {
	value *ListOrganizationsResponse
	isSet bool
}

func (v NullableListOrganizationsResponse) Get() *ListOrganizationsResponse {
	return v.value
}

func (v *NullableListOrganizationsResponse) Set(val *ListOrganizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationsResponse(val *ListOrganizationsResponse) *NullableListOrganizationsResponse {
	return &NullableListOrganizationsResponse{value: val, isSet: true}
}

func (v NullableListOrganizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


