/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.0.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// TokenPaginationRequestParameters The `Link` HTTP header contains multiple links (`first`, `next`, `last`, `previous`) formatted as: `<https://{project-slug}.projects.oryapis.com/admin/clients?limit={limit}&offset={offset}>; rel=\"{page}\"`  For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
type TokenPaginationRequestParameters struct {
	// Items per Page  This is the number of items per page to return. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
	PageSize *int64 `json:"page_size,omitempty"`
	// Next Page Token  The next page token. For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
	PageToken *string `json:"page_token,omitempty"`
}

// NewTokenPaginationRequestParameters instantiates a new TokenPaginationRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPaginationRequestParameters() *TokenPaginationRequestParameters {
	this := TokenPaginationRequestParameters{}
	var pageSize int64 = 250
	this.PageSize = &pageSize
	var pageToken string = "1"
	this.PageToken = &pageToken
	return &this
}

// NewTokenPaginationRequestParametersWithDefaults instantiates a new TokenPaginationRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPaginationRequestParametersWithDefaults() *TokenPaginationRequestParameters {
	this := TokenPaginationRequestParameters{}
	var pageSize int64 = 250
	this.PageSize = &pageSize
	var pageToken string = "1"
	this.PageToken = &pageToken
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *TokenPaginationRequestParameters) GetPageSize() int64 {
	if o == nil || o.PageSize == nil {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPaginationRequestParameters) GetPageSizeOk() (*int64, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *TokenPaginationRequestParameters) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *TokenPaginationRequestParameters) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *TokenPaginationRequestParameters) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPaginationRequestParameters) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *TokenPaginationRequestParameters) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *TokenPaginationRequestParameters) SetPageToken(v string) {
	o.PageToken = &v
}

func (o TokenPaginationRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenPaginationRequestParameters struct {
	value *TokenPaginationRequestParameters
	isSet bool
}

func (v NullableTokenPaginationRequestParameters) Get() *TokenPaginationRequestParameters {
	return v.value
}

func (v *NullableTokenPaginationRequestParameters) Set(val *TokenPaginationRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPaginationRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPaginationRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPaginationRequestParameters(val *TokenPaginationRequestParameters) *NullableTokenPaginationRequestParameters {
	return &NullableTokenPaginationRequestParameters{value: val, isSet: true}
}

func (v NullableTokenPaginationRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPaginationRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


