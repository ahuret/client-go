/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Usage struct for Usage
type Usage struct {
	GenericUsage *GenericUsage `json:"GenericUsage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Usage Usage

// NewUsage instantiates a new Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsage() *Usage {
	this := Usage{}
	return &this
}

// NewUsageWithDefaults instantiates a new Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageWithDefaults() *Usage {
	this := Usage{}
	return &this
}

// GetGenericUsage returns the GenericUsage field value if set, zero value otherwise.
func (o *Usage) GetGenericUsage() GenericUsage {
	if o == nil || o.GenericUsage == nil {
		var ret GenericUsage
		return ret
	}
	return *o.GenericUsage
}

// GetGenericUsageOk returns a tuple with the GenericUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usage) GetGenericUsageOk() (*GenericUsage, bool) {
	if o == nil || o.GenericUsage == nil {
		return nil, false
	}
	return o.GenericUsage, true
}

// HasGenericUsage returns a boolean if a field has been set.
func (o *Usage) HasGenericUsage() bool {
	if o != nil && o.GenericUsage != nil {
		return true
	}

	return false
}

// SetGenericUsage gets a reference to the given GenericUsage and assigns it to the GenericUsage field.
func (o *Usage) SetGenericUsage(v GenericUsage) {
	o.GenericUsage = &v
}

func (o Usage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GenericUsage != nil {
		toSerialize["GenericUsage"] = o.GenericUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Usage) UnmarshalJSON(bytes []byte) (err error) {
	varUsage := _Usage{}

	if err = json.Unmarshal(bytes, &varUsage); err == nil {
		*o = Usage(varUsage)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "GenericUsage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsage struct {
	value *Usage
	isSet bool
}

func (v NullableUsage) Get() *Usage {
	return v.value
}

func (v *NullableUsage) Set(val *Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsage(val *Usage) *NullableUsage {
	return &NullableUsage{value: val, isSet: true}
}

func (v NullableUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


