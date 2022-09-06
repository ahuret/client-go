/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.28
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetRelationTuplesResponse struct for GetRelationTuplesResponse
type GetRelationTuplesResponse struct {
	// The opaque token to provide in a subsequent request to get the next page. It is the empty string iff this is the last page.
	NextPageToken *string `json:"next_page_token,omitempty"`
	RelationTuples []RelationTuple `json:"relation_tuples,omitempty"`
}

// NewGetRelationTuplesResponse instantiates a new GetRelationTuplesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRelationTuplesResponse() *GetRelationTuplesResponse {
	this := GetRelationTuplesResponse{}
	return &this
}

// NewGetRelationTuplesResponseWithDefaults instantiates a new GetRelationTuplesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRelationTuplesResponseWithDefaults() *GetRelationTuplesResponse {
	this := GetRelationTuplesResponse{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *GetRelationTuplesResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRelationTuplesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *GetRelationTuplesResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *GetRelationTuplesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetRelationTuples returns the RelationTuples field value if set, zero value otherwise.
func (o *GetRelationTuplesResponse) GetRelationTuples() []RelationTuple {
	if o == nil || o.RelationTuples == nil {
		var ret []RelationTuple
		return ret
	}
	return o.RelationTuples
}

// GetRelationTuplesOk returns a tuple with the RelationTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRelationTuplesResponse) GetRelationTuplesOk() ([]RelationTuple, bool) {
	if o == nil || o.RelationTuples == nil {
		return nil, false
	}
	return o.RelationTuples, true
}

// HasRelationTuples returns a boolean if a field has been set.
func (o *GetRelationTuplesResponse) HasRelationTuples() bool {
	if o != nil && o.RelationTuples != nil {
		return true
	}

	return false
}

// SetRelationTuples gets a reference to the given []RelationTuple and assigns it to the RelationTuples field.
func (o *GetRelationTuplesResponse) SetRelationTuples(v []RelationTuple) {
	o.RelationTuples = v
}

func (o GetRelationTuplesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	if o.RelationTuples != nil {
		toSerialize["relation_tuples"] = o.RelationTuples
	}
	return json.Marshal(toSerialize)
}

type NullableGetRelationTuplesResponse struct {
	value *GetRelationTuplesResponse
	isSet bool
}

func (v NullableGetRelationTuplesResponse) Get() *GetRelationTuplesResponse {
	return v.value
}

func (v *NullableGetRelationTuplesResponse) Set(val *GetRelationTuplesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRelationTuplesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRelationTuplesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRelationTuplesResponse(val *GetRelationTuplesResponse) *NullableGetRelationTuplesResponse {
	return &NullableGetRelationTuplesResponse{value: val, isSet: true}
}

func (v NullableGetRelationTuplesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRelationTuplesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


