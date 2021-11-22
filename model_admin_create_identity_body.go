/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.24
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AdminCreateIdentityBody struct for AdminCreateIdentityBody
type AdminCreateIdentityBody struct {
	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	SchemaId string `json:"schema_id"`
	State *IdentityState `json:"state,omitempty"`
	// Traits represent an identity's traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in `schema_url`.
	Traits map[string]interface{} `json:"traits"`
}

// NewAdminCreateIdentityBody instantiates a new AdminCreateIdentityBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateIdentityBody(schemaId string, traits map[string]interface{}) *AdminCreateIdentityBody {
	this := AdminCreateIdentityBody{}
	this.SchemaId = schemaId
	this.Traits = traits
	return &this
}

// NewAdminCreateIdentityBodyWithDefaults instantiates a new AdminCreateIdentityBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateIdentityBodyWithDefaults() *AdminCreateIdentityBody {
	this := AdminCreateIdentityBody{}
	return &this
}

// GetSchemaId returns the SchemaId field value
func (o *AdminCreateIdentityBody) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityBody) GetSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *AdminCreateIdentityBody) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AdminCreateIdentityBody) GetState() IdentityState {
	if o == nil || o.State == nil {
		var ret IdentityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityBody) GetStateOk() (*IdentityState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AdminCreateIdentityBody) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given IdentityState and assigns it to the State field.
func (o *AdminCreateIdentityBody) SetState(v IdentityState) {
	o.State = &v
}

// GetTraits returns the Traits field value
func (o *AdminCreateIdentityBody) GetTraits() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *AdminCreateIdentityBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *AdminCreateIdentityBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o AdminCreateIdentityBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["traits"] = o.Traits
	}
	return json.Marshal(toSerialize)
}

type NullableAdminCreateIdentityBody struct {
	value *AdminCreateIdentityBody
	isSet bool
}

func (v NullableAdminCreateIdentityBody) Get() *AdminCreateIdentityBody {
	return v.value
}

func (v *NullableAdminCreateIdentityBody) Set(val *AdminCreateIdentityBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateIdentityBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateIdentityBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateIdentityBody(val *AdminCreateIdentityBody) *NullableAdminCreateIdentityBody {
	return &NullableAdminCreateIdentityBody{value: val, isSet: true}
}

func (v NullableAdminCreateIdentityBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateIdentityBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


