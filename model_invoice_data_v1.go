/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the InvoiceDataV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDataV1{}

// InvoiceDataV1 struct for InvoiceDataV1
type InvoiceDataV1 struct {
	BillingPeriod TimeInterval `json:"billing_period"`
	// The currency of the invoice.
	Currency string `json:"currency"`
	// The items that are part of this invoice.
	Items []LineItemV1 `json:"items"`
	// The plan that this invoice is based on.
	Plan *string `json:"plan,omitempty"`
	StripeInvoiceItem *string `json:"stripe_invoice_item,omitempty"`
	// The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
	StripeInvoiceStatus *string `json:"stripe_invoice_status,omitempty"`
	// An optional link to the invoice on Stripe.
	StripeLink *string `json:"stripe_link,omitempty"`
	// The subtitle of the invoice.
	Subtitle *string `json:"subtitle,omitempty"`
	Tax *TaxLineItem `json:"tax,omitempty"`
	// The title of the invoice.
	Title string `json:"title"`
	TotalInCent int64 `json:"total_in_cent"`
	AdditionalProperties map[string]interface{}
}

type _InvoiceDataV1 InvoiceDataV1

// NewInvoiceDataV1 instantiates a new InvoiceDataV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDataV1(billingPeriod TimeInterval, currency string, items []LineItemV1, title string, totalInCent int64) *InvoiceDataV1 {
	this := InvoiceDataV1{}
	this.BillingPeriod = billingPeriod
	this.Currency = currency
	this.Items = items
	this.Title = title
	this.TotalInCent = totalInCent
	return &this
}

// NewInvoiceDataV1WithDefaults instantiates a new InvoiceDataV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDataV1WithDefaults() *InvoiceDataV1 {
	this := InvoiceDataV1{}
	return &this
}

// GetBillingPeriod returns the BillingPeriod field value
func (o *InvoiceDataV1) GetBillingPeriod() TimeInterval {
	if o == nil {
		var ret TimeInterval
		return ret
	}

	return o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetBillingPeriodOk() (*TimeInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingPeriod, true
}

// SetBillingPeriod sets field value
func (o *InvoiceDataV1) SetBillingPeriod(v TimeInterval) {
	o.BillingPeriod = v
}

// GetCurrency returns the Currency field value
func (o *InvoiceDataV1) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InvoiceDataV1) SetCurrency(v string) {
	o.Currency = v
}

// GetItems returns the Items field value
func (o *InvoiceDataV1) GetItems() []LineItemV1 {
	if o == nil {
		var ret []LineItemV1
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetItemsOk() ([]LineItemV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InvoiceDataV1) SetItems(v []LineItemV1) {
	o.Items = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetPlan() string {
	if o == nil || IsNil(o.Plan) {
		var ret string
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetPlanOk() (*string, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given string and assigns it to the Plan field.
func (o *InvoiceDataV1) SetPlan(v string) {
	o.Plan = &v
}

// GetStripeInvoiceItem returns the StripeInvoiceItem field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetStripeInvoiceItem() string {
	if o == nil || IsNil(o.StripeInvoiceItem) {
		var ret string
		return ret
	}
	return *o.StripeInvoiceItem
}

// GetStripeInvoiceItemOk returns a tuple with the StripeInvoiceItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetStripeInvoiceItemOk() (*string, bool) {
	if o == nil || IsNil(o.StripeInvoiceItem) {
		return nil, false
	}
	return o.StripeInvoiceItem, true
}

// HasStripeInvoiceItem returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasStripeInvoiceItem() bool {
	if o != nil && !IsNil(o.StripeInvoiceItem) {
		return true
	}

	return false
}

// SetStripeInvoiceItem gets a reference to the given string and assigns it to the StripeInvoiceItem field.
func (o *InvoiceDataV1) SetStripeInvoiceItem(v string) {
	o.StripeInvoiceItem = &v
}

// GetStripeInvoiceStatus returns the StripeInvoiceStatus field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetStripeInvoiceStatus() string {
	if o == nil || IsNil(o.StripeInvoiceStatus) {
		var ret string
		return ret
	}
	return *o.StripeInvoiceStatus
}

// GetStripeInvoiceStatusOk returns a tuple with the StripeInvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetStripeInvoiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.StripeInvoiceStatus) {
		return nil, false
	}
	return o.StripeInvoiceStatus, true
}

// HasStripeInvoiceStatus returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasStripeInvoiceStatus() bool {
	if o != nil && !IsNil(o.StripeInvoiceStatus) {
		return true
	}

	return false
}

// SetStripeInvoiceStatus gets a reference to the given string and assigns it to the StripeInvoiceStatus field.
func (o *InvoiceDataV1) SetStripeInvoiceStatus(v string) {
	o.StripeInvoiceStatus = &v
}

// GetStripeLink returns the StripeLink field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetStripeLink() string {
	if o == nil || IsNil(o.StripeLink) {
		var ret string
		return ret
	}
	return *o.StripeLink
}

// GetStripeLinkOk returns a tuple with the StripeLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetStripeLinkOk() (*string, bool) {
	if o == nil || IsNil(o.StripeLink) {
		return nil, false
	}
	return o.StripeLink, true
}

// HasStripeLink returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasStripeLink() bool {
	if o != nil && !IsNil(o.StripeLink) {
		return true
	}

	return false
}

// SetStripeLink gets a reference to the given string and assigns it to the StripeLink field.
func (o *InvoiceDataV1) SetStripeLink(v string) {
	o.StripeLink = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *InvoiceDataV1) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *InvoiceDataV1) GetTax() TaxLineItem {
	if o == nil || IsNil(o.Tax) {
		var ret TaxLineItem
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetTaxOk() (*TaxLineItem, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *InvoiceDataV1) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given TaxLineItem and assigns it to the Tax field.
func (o *InvoiceDataV1) SetTax(v TaxLineItem) {
	o.Tax = &v
}

// GetTitle returns the Title field value
func (o *InvoiceDataV1) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *InvoiceDataV1) SetTitle(v string) {
	o.Title = v
}

// GetTotalInCent returns the TotalInCent field value
func (o *InvoiceDataV1) GetTotalInCent() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalInCent
}

// GetTotalInCentOk returns a tuple with the TotalInCent field value
// and a boolean to check if the value has been set.
func (o *InvoiceDataV1) GetTotalInCentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInCent, true
}

// SetTotalInCent sets field value
func (o *InvoiceDataV1) SetTotalInCent(v int64) {
	o.TotalInCent = v
}

func (o InvoiceDataV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDataV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billing_period"] = o.BillingPeriod
	toSerialize["currency"] = o.Currency
	toSerialize["items"] = o.Items
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.StripeInvoiceItem) {
		toSerialize["stripe_invoice_item"] = o.StripeInvoiceItem
	}
	if !IsNil(o.StripeInvoiceStatus) {
		toSerialize["stripe_invoice_status"] = o.StripeInvoiceStatus
	}
	if !IsNil(o.StripeLink) {
		toSerialize["stripe_link"] = o.StripeLink
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	toSerialize["title"] = o.Title
	toSerialize["total_in_cent"] = o.TotalInCent

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoiceDataV1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billing_period",
		"currency",
		"items",
		"title",
		"total_in_cent",
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

	varInvoiceDataV1 := _InvoiceDataV1{}

	err = json.Unmarshal(data, &varInvoiceDataV1)

	if err != nil {
		return err
	}

	*o = InvoiceDataV1(varInvoiceDataV1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_period")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "items")
		delete(additionalProperties, "plan")
		delete(additionalProperties, "stripe_invoice_item")
		delete(additionalProperties, "stripe_invoice_status")
		delete(additionalProperties, "stripe_link")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "tax")
		delete(additionalProperties, "title")
		delete(additionalProperties, "total_in_cent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoiceDataV1 struct {
	value *InvoiceDataV1
	isSet bool
}

func (v NullableInvoiceDataV1) Get() *InvoiceDataV1 {
	return v.value
}

func (v *NullableInvoiceDataV1) Set(val *InvoiceDataV1) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDataV1) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDataV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDataV1(val *InvoiceDataV1) *NullableInvoiceDataV1 {
	return &NullableInvoiceDataV1{value: val, isSet: true}
}

func (v NullableInvoiceDataV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDataV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


