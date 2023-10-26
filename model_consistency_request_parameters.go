/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.14
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ConsistencyRequestParameters Control API consistency guarantees
type ConsistencyRequestParameters struct {
	// Read Consistency Level (preview)  The read consistency level determines the consistency guarantee for reads:  strong (slow): The read is guaranteed to return the most recent data committed at the start of the read. eventual (very fast): The result will return data that is about 4.8 seconds old.  The default consistency guarantee can be changed in the Ory Network Console or using the Ory CLI with `ory patch project --replace '/previews/default_read_consistency_level=\"strong\"'`.  Setting the default consistency level to `eventual` may cause regressions in the future as we add consistency controls to more APIs. Currently, the following APIs will be affected by this setting:  `GET /admin/identities`  This feature is in preview and only available in Ory Network.  ConsistencyLevelUnset  ConsistencyLevelUnset is the unset / default consistency level. strong ConsistencyLevelStrong  ConsistencyLevelStrong is the strong consistency level. eventual ConsistencyLevelEventual  ConsistencyLevelEventual is the eventual consistency level using follower read timestamps.
	Consistency *string `json:"consistency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsistencyRequestParameters ConsistencyRequestParameters

// NewConsistencyRequestParameters instantiates a new ConsistencyRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistencyRequestParameters() *ConsistencyRequestParameters {
	this := ConsistencyRequestParameters{}
	return &this
}

// NewConsistencyRequestParametersWithDefaults instantiates a new ConsistencyRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistencyRequestParametersWithDefaults() *ConsistencyRequestParameters {
	this := ConsistencyRequestParameters{}
	return &this
}

// GetConsistency returns the Consistency field value if set, zero value otherwise.
func (o *ConsistencyRequestParameters) GetConsistency() string {
	if o == nil || o.Consistency == nil {
		var ret string
		return ret
	}
	return *o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyRequestParameters) GetConsistencyOk() (*string, bool) {
	if o == nil || o.Consistency == nil {
		return nil, false
	}
	return o.Consistency, true
}

// HasConsistency returns a boolean if a field has been set.
func (o *ConsistencyRequestParameters) HasConsistency() bool {
	if o != nil && o.Consistency != nil {
		return true
	}

	return false
}

// SetConsistency gets a reference to the given string and assigns it to the Consistency field.
func (o *ConsistencyRequestParameters) SetConsistency(v string) {
	o.Consistency = &v
}

func (o ConsistencyRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consistency != nil {
		toSerialize["consistency"] = o.Consistency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConsistencyRequestParameters) UnmarshalJSON(bytes []byte) (err error) {
	varConsistencyRequestParameters := _ConsistencyRequestParameters{}

	if err = json.Unmarshal(bytes, &varConsistencyRequestParameters); err == nil {
		*o = ConsistencyRequestParameters(varConsistencyRequestParameters)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "consistency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsistencyRequestParameters struct {
	value *ConsistencyRequestParameters
	isSet bool
}

func (v NullableConsistencyRequestParameters) Get() *ConsistencyRequestParameters {
	return v.value
}

func (v *NullableConsistencyRequestParameters) Set(val *ConsistencyRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistencyRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistencyRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistencyRequestParameters(val *ConsistencyRequestParameters) *NullableConsistencyRequestParameters {
	return &NullableConsistencyRequestParameters{value: val, isSet: true}
}

func (v NullableConsistencyRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistencyRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


