/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.3.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the BatchPatchIdentitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchPatchIdentitiesResponse{}

// BatchPatchIdentitiesResponse Patch identities response
type BatchPatchIdentitiesResponse struct {
	// The patch responses for the individual identities.
	Identities []IdentityPatchResponse `json:"identities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BatchPatchIdentitiesResponse BatchPatchIdentitiesResponse

// NewBatchPatchIdentitiesResponse instantiates a new BatchPatchIdentitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchPatchIdentitiesResponse() *BatchPatchIdentitiesResponse {
	this := BatchPatchIdentitiesResponse{}
	return &this
}

// NewBatchPatchIdentitiesResponseWithDefaults instantiates a new BatchPatchIdentitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchPatchIdentitiesResponseWithDefaults() *BatchPatchIdentitiesResponse {
	this := BatchPatchIdentitiesResponse{}
	return &this
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *BatchPatchIdentitiesResponse) GetIdentities() []IdentityPatchResponse {
	if o == nil || IsNil(o.Identities) {
		var ret []IdentityPatchResponse
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchPatchIdentitiesResponse) GetIdentitiesOk() ([]IdentityPatchResponse, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *BatchPatchIdentitiesResponse) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []IdentityPatchResponse and assigns it to the Identities field.
func (o *BatchPatchIdentitiesResponse) SetIdentities(v []IdentityPatchResponse) {
	o.Identities = v
}

func (o BatchPatchIdentitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchPatchIdentitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchPatchIdentitiesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBatchPatchIdentitiesResponse := _BatchPatchIdentitiesResponse{}

	err = json.Unmarshal(bytes, &varBatchPatchIdentitiesResponse)

	if err != nil {
		return err
	}

	*o = BatchPatchIdentitiesResponse(varBatchPatchIdentitiesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchPatchIdentitiesResponse struct {
	value *BatchPatchIdentitiesResponse
	isSet bool
}

func (v NullableBatchPatchIdentitiesResponse) Get() *BatchPatchIdentitiesResponse {
	return v.value
}

func (v *NullableBatchPatchIdentitiesResponse) Set(val *BatchPatchIdentitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchPatchIdentitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchPatchIdentitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchPatchIdentitiesResponse(val *BatchPatchIdentitiesResponse) *NullableBatchPatchIdentitiesResponse {
	return &NullableBatchPatchIdentitiesResponse{value: val, isSet: true}
}

func (v NullableBatchPatchIdentitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchPatchIdentitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


