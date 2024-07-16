/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.14.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateInviteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInviteResponse{}

// CreateInviteResponse struct for CreateInviteResponse
type CreateInviteResponse struct {
	// A list of all invites for this resource
	AllInvites []MemberInvite `json:"all_invites"`
	CreatedInvite MemberInvite `json:"created_invite"`
	AdditionalProperties map[string]interface{}
}

type _CreateInviteResponse CreateInviteResponse

// NewCreateInviteResponse instantiates a new CreateInviteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInviteResponse(allInvites []MemberInvite, createdInvite MemberInvite) *CreateInviteResponse {
	this := CreateInviteResponse{}
	this.AllInvites = allInvites
	this.CreatedInvite = createdInvite
	return &this
}

// NewCreateInviteResponseWithDefaults instantiates a new CreateInviteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInviteResponseWithDefaults() *CreateInviteResponse {
	this := CreateInviteResponse{}
	return &this
}

// GetAllInvites returns the AllInvites field value
func (o *CreateInviteResponse) GetAllInvites() []MemberInvite {
	if o == nil {
		var ret []MemberInvite
		return ret
	}

	return o.AllInvites
}

// GetAllInvitesOk returns a tuple with the AllInvites field value
// and a boolean to check if the value has been set.
func (o *CreateInviteResponse) GetAllInvitesOk() ([]MemberInvite, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllInvites, true
}

// SetAllInvites sets field value
func (o *CreateInviteResponse) SetAllInvites(v []MemberInvite) {
	o.AllInvites = v
}

// GetCreatedInvite returns the CreatedInvite field value
func (o *CreateInviteResponse) GetCreatedInvite() MemberInvite {
	if o == nil {
		var ret MemberInvite
		return ret
	}

	return o.CreatedInvite
}

// GetCreatedInviteOk returns a tuple with the CreatedInvite field value
// and a boolean to check if the value has been set.
func (o *CreateInviteResponse) GetCreatedInviteOk() (*MemberInvite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedInvite, true
}

// SetCreatedInvite sets field value
func (o *CreateInviteResponse) SetCreatedInvite(v MemberInvite) {
	o.CreatedInvite = v
}

func (o CreateInviteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInviteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all_invites"] = o.AllInvites
	toSerialize["created_invite"] = o.CreatedInvite

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateInviteResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all_invites",
		"created_invite",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateInviteResponse := _CreateInviteResponse{}

	err = json.Unmarshal(data, &varCreateInviteResponse)

	if err != nil {
		return err
	}

	*o = CreateInviteResponse(varCreateInviteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "all_invites")
		delete(additionalProperties, "created_invite")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateInviteResponse struct {
	value *CreateInviteResponse
	isSet bool
}

func (v NullableCreateInviteResponse) Get() *CreateInviteResponse {
	return v.value
}

func (v *NullableCreateInviteResponse) Set(val *CreateInviteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInviteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInviteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInviteResponse(val *CreateInviteResponse) *NullableCreateInviteResponse {
	return &NullableCreateInviteResponse{value: val, isSet: true}
}

func (v NullableCreateInviteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInviteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


