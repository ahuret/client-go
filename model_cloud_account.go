/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAccount{}

// CloudAccount struct for CloudAccount
type CloudAccount struct {
	Email string `json:"email"`
	Id string `json:"id"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CloudAccount CloudAccount

// NewCloudAccount instantiates a new CloudAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAccount(email string, id string, name string) *CloudAccount {
	this := CloudAccount{}
	this.Email = email
	this.Id = id
	this.Name = name
	return &this
}

// NewCloudAccountWithDefaults instantiates a new CloudAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAccountWithDefaults() *CloudAccount {
	this := CloudAccount{}
	return &this
}

// GetEmail returns the Email field value
func (o *CloudAccount) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CloudAccount) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CloudAccount) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *CloudAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloudAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloudAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CloudAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudAccount) SetName(v string) {
	o.Name = v
}

func (o CloudAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"id",
		"name",
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

	varCloudAccount := _CloudAccount{}

	err = json.Unmarshal(data, &varCloudAccount)

	if err != nil {
		return err
	}

	*o = CloudAccount(varCloudAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAccount struct {
	value *CloudAccount
	isSet bool
}

func (v NullableCloudAccount) Get() *CloudAccount {
	return v.value
}

func (v *NullableCloudAccount) Set(val *CloudAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAccount(val *CloudAccount) *NullableCloudAccount {
	return &NullableCloudAccount{value: val, isSet: true}
}

func (v NullableCloudAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


