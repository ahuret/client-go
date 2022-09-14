/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.41
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// Subscription struct for Subscription
type Subscription struct {
	CreatedAt time.Time `json:"created_at"`
	// The currently active plan of the subscription unknown Unknown free Free start_up_monthly StartUpMonthly start_up_yearly StartUpYearly business_monthly BusinessMonthly business_yearly BusinessYearly custom Custom
	CurrentPlan string `json:"current_plan"`
	// The ID of the stripe customer
	CustomerId string `json:"customer_id"`
	// The ID of the subscription
	Id string `json:"id"`
	OngoingStripeCheckoutId NullableString `json:"ongoing_stripe_checkout_id,omitempty"`
	// Until when the subscription is payed
	PayedUntil time.Time `json:"payed_until"`
	PlanChangesAt *time.Time `json:"plan_changes_at,omitempty"`
	PlanChangesTo NullPlan `json:"plan_changes_to"`
	// For `collection_method=charge_automatically` a subscription moves into `incomplete` if the initial payment attempt fails. A subscription in this state can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an `active` state. If the first invoice is not paid within 23 hours, the subscription transitions to `incomplete_expired`. This is a terminal state, the open invoice will be voided and no further invoices will be generated.  A subscription that is currently in a trial period is `trialing` and moves to `active` when the trial period is over.  If subscription `collection_method=charge_automatically` it becomes `past_due` when payment to renew it fails and `canceled` or `unpaid` (depending on your subscriptions settings) when Stripe has exhausted all payment retry attempts.  If subscription `collection_method=send_invoice` it becomes `past_due` when its invoice is not paid by the due date, and `canceled` or `unpaid` if it is still not paid by an additional deadline after that. Note that when a subscription has a status of `unpaid`, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.
	Status string `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(createdAt time.Time, currentPlan string, customerId string, id string, payedUntil time.Time, planChangesTo NullPlan, status string, updatedAt time.Time) *Subscription {
	this := Subscription{}
	this.CreatedAt = createdAt
	this.CurrentPlan = currentPlan
	this.CustomerId = customerId
	this.Id = id
	this.PayedUntil = payedUntil
	this.PlanChangesTo = planChangesTo
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Subscription) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Subscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrentPlan returns the CurrentPlan field value
func (o *Subscription) GetCurrentPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPlan
}

// GetCurrentPlanOk returns a tuple with the CurrentPlan field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPlan, true
}

// SetCurrentPlan sets field value
func (o *Subscription) SetCurrentPlan(v string) {
	o.CurrentPlan = v
}

// GetCustomerId returns the CustomerId field value
func (o *Subscription) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Subscription) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *Subscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Subscription) SetId(v string) {
	o.Id = v
}

// GetOngoingStripeCheckoutId returns the OngoingStripeCheckoutId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetOngoingStripeCheckoutId() string {
	if o == nil || o.OngoingStripeCheckoutId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OngoingStripeCheckoutId.Get()
}

// GetOngoingStripeCheckoutIdOk returns a tuple with the OngoingStripeCheckoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetOngoingStripeCheckoutIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OngoingStripeCheckoutId.Get(), o.OngoingStripeCheckoutId.IsSet()
}

// HasOngoingStripeCheckoutId returns a boolean if a field has been set.
func (o *Subscription) HasOngoingStripeCheckoutId() bool {
	if o != nil && o.OngoingStripeCheckoutId.IsSet() {
		return true
	}

	return false
}

// SetOngoingStripeCheckoutId gets a reference to the given NullableString and assigns it to the OngoingStripeCheckoutId field.
func (o *Subscription) SetOngoingStripeCheckoutId(v string) {
	o.OngoingStripeCheckoutId.Set(&v)
}
// SetOngoingStripeCheckoutIdNil sets the value for OngoingStripeCheckoutId to be an explicit nil
func (o *Subscription) SetOngoingStripeCheckoutIdNil() {
	o.OngoingStripeCheckoutId.Set(nil)
}

// UnsetOngoingStripeCheckoutId ensures that no value is present for OngoingStripeCheckoutId, not even an explicit nil
func (o *Subscription) UnsetOngoingStripeCheckoutId() {
	o.OngoingStripeCheckoutId.Unset()
}

// GetPayedUntil returns the PayedUntil field value
func (o *Subscription) GetPayedUntil() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PayedUntil
}

// GetPayedUntilOk returns a tuple with the PayedUntil field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetPayedUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayedUntil, true
}

// SetPayedUntil sets field value
func (o *Subscription) SetPayedUntil(v time.Time) {
	o.PayedUntil = v
}

// GetPlanChangesAt returns the PlanChangesAt field value if set, zero value otherwise.
func (o *Subscription) GetPlanChangesAt() time.Time {
	if o == nil || o.PlanChangesAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PlanChangesAt
}

// GetPlanChangesAtOk returns a tuple with the PlanChangesAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPlanChangesAtOk() (*time.Time, bool) {
	if o == nil || o.PlanChangesAt == nil {
		return nil, false
	}
	return o.PlanChangesAt, true
}

// HasPlanChangesAt returns a boolean if a field has been set.
func (o *Subscription) HasPlanChangesAt() bool {
	if o != nil && o.PlanChangesAt != nil {
		return true
	}

	return false
}

// SetPlanChangesAt gets a reference to the given time.Time and assigns it to the PlanChangesAt field.
func (o *Subscription) SetPlanChangesAt(v time.Time) {
	o.PlanChangesAt = &v
}

// GetPlanChangesTo returns the PlanChangesTo field value
func (o *Subscription) GetPlanChangesTo() NullPlan {
	if o == nil {
		var ret NullPlan
		return ret
	}

	return o.PlanChangesTo
}

// GetPlanChangesToOk returns a tuple with the PlanChangesTo field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetPlanChangesToOk() (*NullPlan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanChangesTo, true
}

// SetPlanChangesTo sets field value
func (o *Subscription) SetPlanChangesTo(v NullPlan) {
	o.PlanChangesTo = v
}

// GetStatus returns the Status field value
func (o *Subscription) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Subscription) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Subscription) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Subscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["current_plan"] = o.CurrentPlan
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.OngoingStripeCheckoutId.IsSet() {
		toSerialize["ongoing_stripe_checkout_id"] = o.OngoingStripeCheckoutId.Get()
	}
	if true {
		toSerialize["payed_until"] = o.PayedUntil
	}
	if o.PlanChangesAt != nil {
		toSerialize["plan_changes_at"] = o.PlanChangesAt
	}
	if true {
		toSerialize["plan_changes_to"] = o.PlanChangesTo
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


