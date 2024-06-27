/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the HealthNotReadyStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthNotReadyStatus{}

// HealthNotReadyStatus struct for HealthNotReadyStatus
type HealthNotReadyStatus struct {
	// Errors contains a list of errors that caused the not ready status.
	Errors *map[string]string `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HealthNotReadyStatus HealthNotReadyStatus

// NewHealthNotReadyStatus instantiates a new HealthNotReadyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthNotReadyStatus() *HealthNotReadyStatus {
	this := HealthNotReadyStatus{}
	return &this
}

// NewHealthNotReadyStatusWithDefaults instantiates a new HealthNotReadyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthNotReadyStatusWithDefaults() *HealthNotReadyStatus {
	this := HealthNotReadyStatus{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *HealthNotReadyStatus) GetErrors() map[string]string {
	if o == nil || IsNil(o.Errors) {
		var ret map[string]string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthNotReadyStatus) GetErrorsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *HealthNotReadyStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]string and assigns it to the Errors field.
func (o *HealthNotReadyStatus) SetErrors(v map[string]string) {
	o.Errors = &v
}

func (o HealthNotReadyStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthNotReadyStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HealthNotReadyStatus) UnmarshalJSON(data []byte) (err error) {
	varHealthNotReadyStatus := _HealthNotReadyStatus{}

	err = json.Unmarshal(data, &varHealthNotReadyStatus)

	if err != nil {
		return err
	}

	*o = HealthNotReadyStatus(varHealthNotReadyStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHealthNotReadyStatus struct {
	value *HealthNotReadyStatus
	isSet bool
}

func (v NullableHealthNotReadyStatus) Get() *HealthNotReadyStatus {
	return v.value
}

func (v *NullableHealthNotReadyStatus) Set(val *HealthNotReadyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthNotReadyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthNotReadyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthNotReadyStatus(val *HealthNotReadyStatus) *NullableHealthNotReadyStatus {
	return &NullableHealthNotReadyStatus{value: val, isSet: true}
}

func (v NullableHealthNotReadyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthNotReadyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


