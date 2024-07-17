/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.14.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the TaxLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxLineItem{}

// TaxLineItem struct for TaxLineItem
type TaxLineItem struct {
	AmountInCent *int64 `json:"amount_in_cent,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaxLineItem TaxLineItem

// NewTaxLineItem instantiates a new TaxLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxLineItem() *TaxLineItem {
	this := TaxLineItem{}
	return &this
}

// NewTaxLineItemWithDefaults instantiates a new TaxLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxLineItemWithDefaults() *TaxLineItem {
	this := TaxLineItem{}
	return &this
}

// GetAmountInCent returns the AmountInCent field value if set, zero value otherwise.
func (o *TaxLineItem) GetAmountInCent() int64 {
	if o == nil || IsNil(o.AmountInCent) {
		var ret int64
		return ret
	}
	return *o.AmountInCent
}

// GetAmountInCentOk returns a tuple with the AmountInCent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxLineItem) GetAmountInCentOk() (*int64, bool) {
	if o == nil || IsNil(o.AmountInCent) {
		return nil, false
	}
	return o.AmountInCent, true
}

// HasAmountInCent returns a boolean if a field has been set.
func (o *TaxLineItem) HasAmountInCent() bool {
	if o != nil && !IsNil(o.AmountInCent) {
		return true
	}

	return false
}

// SetAmountInCent gets a reference to the given int64 and assigns it to the AmountInCent field.
func (o *TaxLineItem) SetAmountInCent(v int64) {
	o.AmountInCent = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TaxLineItem) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxLineItem) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TaxLineItem) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TaxLineItem) SetTitle(v string) {
	o.Title = &v
}

func (o TaxLineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountInCent) {
		toSerialize["amount_in_cent"] = o.AmountInCent
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TaxLineItem) UnmarshalJSON(data []byte) (err error) {
	varTaxLineItem := _TaxLineItem{}

	err = json.Unmarshal(data, &varTaxLineItem)

	if err != nil {
		return err
	}

	*o = TaxLineItem(varTaxLineItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount_in_cent")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaxLineItem struct {
	value *TaxLineItem
	isSet bool
}

func (v NullableTaxLineItem) Get() *TaxLineItem {
	return v.value
}

func (v *NullableTaxLineItem) Set(val *TaxLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxLineItem(val *TaxLineItem) *NullableTaxLineItem {
	return &NullableTaxLineItem{value: val, isSet: true}
}

func (v NullableTaxLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


