/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.29
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProvisionProjectPayload struct for ProvisionProjectPayload
type ProvisionProjectPayload struct {
	// The stripe subscription ID to use for provisioning the project.
	SubscriptionId string `json:"subscription_id"`
}

// NewProvisionProjectPayload instantiates a new ProvisionProjectPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionProjectPayload(subscriptionId string) *ProvisionProjectPayload {
	this := ProvisionProjectPayload{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewProvisionProjectPayloadWithDefaults instantiates a new ProvisionProjectPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionProjectPayloadWithDefaults() *ProvisionProjectPayload {
	this := ProvisionProjectPayload{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ProvisionProjectPayload) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ProvisionProjectPayload) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ProvisionProjectPayload) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o ProvisionProjectPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	return json.Marshal(toSerialize)
}

type NullableProvisionProjectPayload struct {
	value *ProvisionProjectPayload
	isSet bool
}

func (v NullableProvisionProjectPayload) Get() *ProvisionProjectPayload {
	return v.value
}

func (v *NullableProvisionProjectPayload) Set(val *ProvisionProjectPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionProjectPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionProjectPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionProjectPayload(val *ProvisionProjectPayload) *NullableProvisionProjectPayload {
	return &NullableProvisionProjectPayload{value: val, isSet: true}
}

func (v NullableProvisionProjectPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionProjectPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


