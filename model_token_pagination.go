/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.21
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// TokenPagination struct for TokenPagination
type TokenPagination struct {
	// Items per page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
	PageSize *int64 `json:"page_size,omitempty"`
	// Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
	PageToken *string `json:"page_token,omitempty"`
}

// NewTokenPagination instantiates a new TokenPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPagination() *TokenPagination {
	this := TokenPagination{}
	var pageSize int64 = 250
	this.PageSize = &pageSize
	var pageToken string = "1"
	this.PageToken = &pageToken
	return &this
}

// NewTokenPaginationWithDefaults instantiates a new TokenPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPaginationWithDefaults() *TokenPagination {
	this := TokenPagination{}
	var pageSize int64 = 250
	this.PageSize = &pageSize
	var pageToken string = "1"
	this.PageToken = &pageToken
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *TokenPagination) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPagination) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *TokenPagination) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *TokenPagination) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *TokenPagination) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPagination) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *TokenPagination) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *TokenPagination) SetPageToken(v string) {
	o.PageToken = &v
}

func (o TokenPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenPagination struct {
	value *TokenPagination
	isSet bool
}

func (v NullableTokenPagination) Get() *TokenPagination {
	return v.value
}

func (v *NullableTokenPagination) Set(val *TokenPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPagination(val *TokenPagination) *NullableTokenPagination {
	return &NullableTokenPagination{value: val, isSet: true}
}

func (v NullableTokenPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


