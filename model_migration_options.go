/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.6.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the MigrationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigrationOptions{}

// MigrationOptions struct for MigrationOptions
type MigrationOptions struct {
	// The environment of the project in the workspace. Can be one of \"prod\" or \"dev\". Note that the number of projects in the \"prod\" environment is limited depending on the subscription. prod Production dev Development
	Environment string `json:"environment"`
	// The action to take with the project subscription. Can be one of \"migrate\", and \"ignore\". \"migrate\" will migrate the project subscription to the workspace. \"ignore\" will ignore the project subscription. migrate ProjectSubscriptionActionMigrate  ProjectSubscriptionActionMigrate will migrate the project subscription to the  workspace. ignore ProjectSubscriptionActionIgnore  ProjectSubscriptionActionIgnore will ignore the project subscription.
	ProjectSubscription string `json:"project_subscription"`
	AdditionalProperties map[string]interface{}
}

type _MigrationOptions MigrationOptions

// NewMigrationOptions instantiates a new MigrationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationOptions(environment string, projectSubscription string) *MigrationOptions {
	this := MigrationOptions{}
	this.Environment = environment
	this.ProjectSubscription = projectSubscription
	return &this
}

// NewMigrationOptionsWithDefaults instantiates a new MigrationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationOptionsWithDefaults() *MigrationOptions {
	this := MigrationOptions{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *MigrationOptions) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *MigrationOptions) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *MigrationOptions) SetEnvironment(v string) {
	o.Environment = v
}

// GetProjectSubscription returns the ProjectSubscription field value
func (o *MigrationOptions) GetProjectSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSubscription
}

// GetProjectSubscriptionOk returns a tuple with the ProjectSubscription field value
// and a boolean to check if the value has been set.
func (o *MigrationOptions) GetProjectSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSubscription, true
}

// SetProjectSubscription sets field value
func (o *MigrationOptions) SetProjectSubscription(v string) {
	o.ProjectSubscription = v
}

func (o MigrationOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigrationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["project_subscription"] = o.ProjectSubscription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MigrationOptions) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"project_subscription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMigrationOptions := _MigrationOptions{}

	err = json.Unmarshal(bytes, &varMigrationOptions)

	if err != nil {
		return err
	}

	*o = MigrationOptions(varMigrationOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "project_subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMigrationOptions struct {
	value *MigrationOptions
	isSet bool
}

func (v NullableMigrationOptions) Get() *MigrationOptions {
	return v.value
}

func (v *NullableMigrationOptions) Set(val *MigrationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationOptions(val *MigrationOptions) *NullableMigrationOptions {
	return &NullableMigrationOptions{value: val, isSet: true}
}

func (v NullableMigrationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


