/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateWorkspaceSubscriptionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceSubscriptionBody{}

// CreateWorkspaceSubscriptionBody struct for CreateWorkspaceSubscriptionBody
type CreateWorkspaceSubscriptionBody struct {
	//  usd USD eur Euro
	Currency *string `json:"currency,omitempty"`
	//  monthly Monthly yearly Yearly
	Interval string `json:"interval"`
	Plan string `json:"plan"`
	ReturnTo *string `json:"return_to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateWorkspaceSubscriptionBody CreateWorkspaceSubscriptionBody

// NewCreateWorkspaceSubscriptionBody instantiates a new CreateWorkspaceSubscriptionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceSubscriptionBody(interval string, plan string) *CreateWorkspaceSubscriptionBody {
	this := CreateWorkspaceSubscriptionBody{}
	this.Interval = interval
	this.Plan = plan
	return &this
}

// NewCreateWorkspaceSubscriptionBodyWithDefaults instantiates a new CreateWorkspaceSubscriptionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceSubscriptionBodyWithDefaults() *CreateWorkspaceSubscriptionBody {
	this := CreateWorkspaceSubscriptionBody{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateWorkspaceSubscriptionBody) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceSubscriptionBody) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateWorkspaceSubscriptionBody) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateWorkspaceSubscriptionBody) SetCurrency(v string) {
	o.Currency = &v
}

// GetInterval returns the Interval field value
func (o *CreateWorkspaceSubscriptionBody) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceSubscriptionBody) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *CreateWorkspaceSubscriptionBody) SetInterval(v string) {
	o.Interval = v
}

// GetPlan returns the Plan field value
func (o *CreateWorkspaceSubscriptionBody) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceSubscriptionBody) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *CreateWorkspaceSubscriptionBody) SetPlan(v string) {
	o.Plan = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *CreateWorkspaceSubscriptionBody) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceSubscriptionBody) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *CreateWorkspaceSubscriptionBody) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *CreateWorkspaceSubscriptionBody) SetReturnTo(v string) {
	o.ReturnTo = &v
}

func (o CreateWorkspaceSubscriptionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceSubscriptionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["interval"] = o.Interval
	toSerialize["plan"] = o.Plan
	if !IsNil(o.ReturnTo) {
		toSerialize["return_to"] = o.ReturnTo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateWorkspaceSubscriptionBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"plan",
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

	varCreateWorkspaceSubscriptionBody := _CreateWorkspaceSubscriptionBody{}

	err = json.Unmarshal(data, &varCreateWorkspaceSubscriptionBody)

	if err != nil {
		return err
	}

	*o = CreateWorkspaceSubscriptionBody(varCreateWorkspaceSubscriptionBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "return_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWorkspaceSubscriptionBody struct {
	value *CreateWorkspaceSubscriptionBody
	isSet bool
}

func (v NullableCreateWorkspaceSubscriptionBody) Get() *CreateWorkspaceSubscriptionBody {
	return v.value
}

func (v *NullableCreateWorkspaceSubscriptionBody) Set(val *CreateWorkspaceSubscriptionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceSubscriptionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceSubscriptionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceSubscriptionBody(val *CreateWorkspaceSubscriptionBody) *NullableCreateWorkspaceSubscriptionBody {
	return &NullableCreateWorkspaceSubscriptionBody{value: val, isSet: true}
}

func (v NullableCreateWorkspaceSubscriptionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceSubscriptionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


