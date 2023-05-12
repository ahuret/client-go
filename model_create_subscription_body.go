/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.28
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateSubscriptionBody Create Subscription Request Body
type CreateSubscriptionBody struct {
	//  monthly Monthly yearly Yearly
	Interval string `json:"interval"`
	Plan string `json:"plan"`
	ProvisionFirstProject string `json:"provision_first_project"`
	ReturnTo *string `json:"return_to,omitempty"`
}

// NewCreateSubscriptionBody instantiates a new CreateSubscriptionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionBody(interval string, plan string, provisionFirstProject string) *CreateSubscriptionBody {
	this := CreateSubscriptionBody{}
	this.Interval = interval
	this.Plan = plan
	this.ProvisionFirstProject = provisionFirstProject
	return &this
}

// NewCreateSubscriptionBodyWithDefaults instantiates a new CreateSubscriptionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionBodyWithDefaults() *CreateSubscriptionBody {
	this := CreateSubscriptionBody{}
	return &this
}

// GetInterval returns the Interval field value
func (o *CreateSubscriptionBody) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionBody) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *CreateSubscriptionBody) SetInterval(v string) {
	o.Interval = v
}

// GetPlan returns the Plan field value
func (o *CreateSubscriptionBody) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionBody) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CreateSubscriptionBody) SetPlan(v string) {
	o.Plan = v
}

// GetProvisionFirstProject returns the ProvisionFirstProject field value
func (o *CreateSubscriptionBody) GetProvisionFirstProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProvisionFirstProject
}

// GetProvisionFirstProjectOk returns a tuple with the ProvisionFirstProject field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionBody) GetProvisionFirstProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisionFirstProject, true
}

// SetProvisionFirstProject sets field value
func (o *CreateSubscriptionBody) SetProvisionFirstProject(v string) {
	o.ProvisionFirstProject = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *CreateSubscriptionBody) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionBody) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *CreateSubscriptionBody) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *CreateSubscriptionBody) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o CreateSubscriptionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["provision_first_project"] = o.ProvisionFirstProject
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSubscriptionBody struct {
	value *CreateSubscriptionBody
	isSet bool
}

func (v NullableCreateSubscriptionBody) Get() *CreateSubscriptionBody {
	return v.value
}

func (v *NullableCreateSubscriptionBody) Set(val *CreateSubscriptionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionBody(val *CreateSubscriptionBody) *NullableCreateSubscriptionBody {
	return &NullableCreateSubscriptionBody{value: val, isSet: true}
}

func (v NullableCreateSubscriptionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


