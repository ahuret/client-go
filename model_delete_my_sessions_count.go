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
)

// checks if the DeleteMySessionsCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMySessionsCount{}

// DeleteMySessionsCount Deleted Session Count
type DeleteMySessionsCount struct {
	// The number of sessions that were revoked.
	Count *int64 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteMySessionsCount DeleteMySessionsCount

// NewDeleteMySessionsCount instantiates a new DeleteMySessionsCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMySessionsCount() *DeleteMySessionsCount {
	this := DeleteMySessionsCount{}
	return &this
}

// NewDeleteMySessionsCountWithDefaults instantiates a new DeleteMySessionsCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMySessionsCountWithDefaults() *DeleteMySessionsCount {
	this := DeleteMySessionsCount{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DeleteMySessionsCount) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMySessionsCount) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DeleteMySessionsCount) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *DeleteMySessionsCount) SetCount(v int64) {
	o.Count = &v
}

func (o DeleteMySessionsCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMySessionsCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteMySessionsCount) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteMySessionsCount := _DeleteMySessionsCount{}

	err = json.Unmarshal(bytes, &varDeleteMySessionsCount)

	if err != nil {
		return err
	}

	*o = DeleteMySessionsCount(varDeleteMySessionsCount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteMySessionsCount struct {
	value *DeleteMySessionsCount
	isSet bool
}

func (v NullableDeleteMySessionsCount) Get() *DeleteMySessionsCount {
	return v.value
}

func (v *NullableDeleteMySessionsCount) Set(val *DeleteMySessionsCount) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMySessionsCount) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMySessionsCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMySessionsCount(val *DeleteMySessionsCount) *NullableDeleteMySessionsCount {
	return &NullableDeleteMySessionsCount{value: val, isSet: true}
}

func (v NullableDeleteMySessionsCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMySessionsCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


