/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.12
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IdentityWithCredentialsPassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityWithCredentialsPassword{}

// IdentityWithCredentialsPassword Create Identity and Import Password Credentials
type IdentityWithCredentialsPassword struct {
	Config *IdentityWithCredentialsPasswordConfig `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityWithCredentialsPassword IdentityWithCredentialsPassword

// NewIdentityWithCredentialsPassword instantiates a new IdentityWithCredentialsPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsPassword() *IdentityWithCredentialsPassword {
	this := IdentityWithCredentialsPassword{}
	return &this
}

// NewIdentityWithCredentialsPasswordWithDefaults instantiates a new IdentityWithCredentialsPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsPasswordWithDefaults() *IdentityWithCredentialsPassword {
	this := IdentityWithCredentialsPassword{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPassword) GetConfig() IdentityWithCredentialsPasswordConfig {
	if o == nil || IsNil(o.Config) {
		var ret IdentityWithCredentialsPasswordConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPassword) GetConfigOk() (*IdentityWithCredentialsPasswordConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPassword) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentityWithCredentialsPasswordConfig and assigns it to the Config field.
func (o *IdentityWithCredentialsPassword) SetConfig(v IdentityWithCredentialsPasswordConfig) {
	o.Config = &v
}

func (o IdentityWithCredentialsPassword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityWithCredentialsPassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityWithCredentialsPassword) UnmarshalJSON(data []byte) (err error) {
	varIdentityWithCredentialsPassword := _IdentityWithCredentialsPassword{}

	err = json.Unmarshal(data, &varIdentityWithCredentialsPassword)

	if err != nil {
		return err
	}

	*o = IdentityWithCredentialsPassword(varIdentityWithCredentialsPassword)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityWithCredentialsPassword struct {
	value *IdentityWithCredentialsPassword
	isSet bool
}

func (v NullableIdentityWithCredentialsPassword) Get() *IdentityWithCredentialsPassword {
	return v.value
}

func (v *NullableIdentityWithCredentialsPassword) Set(val *IdentityWithCredentialsPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsPassword(val *IdentityWithCredentialsPassword) *NullableIdentityWithCredentialsPassword {
	return &NullableIdentityWithCredentialsPassword{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


