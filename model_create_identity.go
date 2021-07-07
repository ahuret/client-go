/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.12
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateIdentity struct for CreateIdentity
type CreateIdentity struct {
	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	SchemaId string `json:"schema_id"`
	// Traits represent an identity's traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in `schema_url`.
	Traits map[string]interface{} `json:"traits"`
}

// NewCreateIdentity instantiates a new CreateIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIdentity(schemaId string, traits map[string]interface{}) *CreateIdentity {
	this := CreateIdentity{}
	this.SchemaId = schemaId
	this.Traits = traits
	return &this
}

// NewCreateIdentityWithDefaults instantiates a new CreateIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIdentityWithDefaults() *CreateIdentity {
	this := CreateIdentity{}
	return &this
}

// GetSchemaId returns the SchemaId field value
func (o *CreateIdentity) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity) GetSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *CreateIdentity) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetTraits returns the Traits field value
func (o *CreateIdentity) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *CreateIdentity) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *CreateIdentity) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o CreateIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schema_id"] = o.SchemaId
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableCreateIdentity struct {
	value *CreateIdentity
	isSet bool
}

func (v NullableCreateIdentity) Get() *CreateIdentity {
	return v.value
}

func (v *NullableCreateIdentity) Set(val *CreateIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIdentity(val *CreateIdentity) *NullableCreateIdentity {
	return &NullableCreateIdentity{value: val, isSet: true}
}

func (v NullableCreateIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


