/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.38
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GenericUsage struct for GenericUsage
type GenericUsage struct {
	// AdditionalPrice is the price per-unit exceeding IncludedUsage. A price of 0 means that no other items can be consumed.
	AdditionalPrice int64 `json:"additional_price"`
	// IncludedUsage is the number of included items.
	IncludedUsage int64 `json:"included_usage"`
	AdditionalProperties map[string]interface{}
}

type _GenericUsage GenericUsage

// NewGenericUsage instantiates a new GenericUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericUsage(additionalPrice int64, includedUsage int64) *GenericUsage {
	this := GenericUsage{}
	this.AdditionalPrice = additionalPrice
	this.IncludedUsage = includedUsage
	return &this
}

// NewGenericUsageWithDefaults instantiates a new GenericUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericUsageWithDefaults() *GenericUsage {
	this := GenericUsage{}
	return &this
}

// GetAdditionalPrice returns the AdditionalPrice field value
func (o *GenericUsage) GetAdditionalPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdditionalPrice
}

// GetAdditionalPriceOk returns a tuple with the AdditionalPrice field value
// and a boolean to check if the value has been set.
func (o *GenericUsage) GetAdditionalPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalPrice, true
}

// SetAdditionalPrice sets field value
func (o *GenericUsage) SetAdditionalPrice(v int64) {
	o.AdditionalPrice = v
}

// GetIncludedUsage returns the IncludedUsage field value
func (o *GenericUsage) GetIncludedUsage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IncludedUsage
}

// GetIncludedUsageOk returns a tuple with the IncludedUsage field value
// and a boolean to check if the value has been set.
func (o *GenericUsage) GetIncludedUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludedUsage, true
}

// SetIncludedUsage sets field value
func (o *GenericUsage) SetIncludedUsage(v int64) {
	o.IncludedUsage = v
}

func (o GenericUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["additional_price"] = o.AdditionalPrice
	}
	if true {
		toSerialize["included_usage"] = o.IncludedUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GenericUsage) UnmarshalJSON(bytes []byte) (err error) {
	varGenericUsage := _GenericUsage{}

	if err = json.Unmarshal(bytes, &varGenericUsage); err == nil {
		*o = GenericUsage(varGenericUsage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_price")
		delete(additionalProperties, "included_usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericUsage struct {
	value *GenericUsage
	isSet bool
}

func (v NullableGenericUsage) Get() *GenericUsage {
	return v.value
}

func (v *NullableGenericUsage) Set(val *GenericUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericUsage(val *GenericUsage) *NullableGenericUsage {
	return &NullableGenericUsage{value: val, isSet: true}
}

func (v NullableGenericUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


