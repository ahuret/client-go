/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.112
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateSubscriptionPayload struct for CreateSubscriptionPayload
type CreateSubscriptionPayload struct {
	PlanOrPrice string `json:"plan_or_price"`
	ProvisionFirstProject string `json:"provision_first_project"`
	ReturnTo *string `json:"return_to,omitempty"`
}

// NewCreateSubscriptionPayload instantiates a new CreateSubscriptionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionPayload(planOrPrice string, provisionFirstProject string) *CreateSubscriptionPayload {
	this := CreateSubscriptionPayload{}
	this.PlanOrPrice = planOrPrice
	this.ProvisionFirstProject = provisionFirstProject
	return &this
}

// NewCreateSubscriptionPayloadWithDefaults instantiates a new CreateSubscriptionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionPayloadWithDefaults() *CreateSubscriptionPayload {
	this := CreateSubscriptionPayload{}
	return &this
}

// GetPlanOrPrice returns the PlanOrPrice field value
func (o *CreateSubscriptionPayload) GetPlanOrPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanOrPrice
}

// GetPlanOrPriceOk returns a tuple with the PlanOrPrice field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionPayload) GetPlanOrPriceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlanOrPrice, true
}

// SetPlanOrPrice sets field value
func (o *CreateSubscriptionPayload) SetPlanOrPrice(v string) {
	o.PlanOrPrice = v
}

// GetProvisionFirstProject returns the ProvisionFirstProject field value
func (o *CreateSubscriptionPayload) GetProvisionFirstProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisionFirstProject
}

// GetProvisionFirstProjectOk returns a tuple with the ProvisionFirstProject field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionPayload) GetProvisionFirstProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProvisionFirstProject, true
}

// SetProvisionFirstProject sets field value
func (o *CreateSubscriptionPayload) SetProvisionFirstProject(v string) {
	o.ProvisionFirstProject = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *CreateSubscriptionPayload) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionPayload) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *CreateSubscriptionPayload) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *CreateSubscriptionPayload) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o CreateSubscriptionPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plan_or_price"] = o.PlanOrPrice
	}
	if true {
		toSerialize["provision_first_project"] = o.ProvisionFirstProject
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSubscriptionPayload struct {
	value *CreateSubscriptionPayload
	isSet bool
}

func (v NullableCreateSubscriptionPayload) Get() *CreateSubscriptionPayload {
	return v.value
}

func (v *NullableCreateSubscriptionPayload) Set(val *CreateSubscriptionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionPayload(val *CreateSubscriptionPayload) *NullableCreateSubscriptionPayload {
	return &NullableCreateSubscriptionPayload{value: val, isSet: true}
}

func (v NullableCreateSubscriptionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


