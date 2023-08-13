/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.45
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PaginationHeaders struct for PaginationHeaders
type PaginationHeaders struct {
	// The link header contains pagination links.  For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).  in: header
	Link *string `json:"link,omitempty"`
	// The total number of clients.  in: header
	XTotalCount *string `json:"x-total-count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginationHeaders PaginationHeaders

// NewPaginationHeaders instantiates a new PaginationHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationHeaders() *PaginationHeaders {
	this := PaginationHeaders{}
	return &this
}

// NewPaginationHeadersWithDefaults instantiates a new PaginationHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationHeadersWithDefaults() *PaginationHeaders {
	this := PaginationHeaders{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *PaginationHeaders) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationHeaders) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *PaginationHeaders) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *PaginationHeaders) SetLink(v string) {
	o.Link = &v
}

// GetXTotalCount returns the XTotalCount field value if set, zero value otherwise.
func (o *PaginationHeaders) GetXTotalCount() string {
	if o == nil || o.XTotalCount == nil {
		var ret string
		return ret
	}
	return *o.XTotalCount
}

// GetXTotalCountOk returns a tuple with the XTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationHeaders) GetXTotalCountOk() (*string, bool) {
	if o == nil || o.XTotalCount == nil {
		return nil, false
	}
	return o.XTotalCount, true
}

// HasXTotalCount returns a boolean if a field has been set.
func (o *PaginationHeaders) HasXTotalCount() bool {
	if o != nil && o.XTotalCount != nil {
		return true
	}

	return false
}

// SetXTotalCount gets a reference to the given string and assigns it to the XTotalCount field.
func (o *PaginationHeaders) SetXTotalCount(v string) {
	o.XTotalCount = &v
}

func (o PaginationHeaders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.XTotalCount != nil {
		toSerialize["x-total-count"] = o.XTotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaginationHeaders) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationHeaders := _PaginationHeaders{}

	if err = json.Unmarshal(bytes, &varPaginationHeaders); err == nil {
		*o = PaginationHeaders(varPaginationHeaders)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "link")
		delete(additionalProperties, "x-total-count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginationHeaders struct {
	value *PaginationHeaders
	isSet bool
}

func (v NullablePaginationHeaders) Get() *PaginationHeaders {
	return v.value
}

func (v *NullablePaginationHeaders) Set(val *PaginationHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationHeaders(val *PaginationHeaders) *NullablePaginationHeaders {
	return &NullablePaginationHeaders{value: val, isSet: true}
}

func (v NullablePaginationHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


