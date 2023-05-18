/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.30
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ContinueWithVerificationUiFlow struct for ContinueWithVerificationUiFlow
type ContinueWithVerificationUiFlow struct {
	// The ID of the verification flow
	Id string `json:"id"`
	// The URL of the verification flow
	Url *string `json:"url,omitempty"`
	// The address that should be verified in this flow
	VerifiableAddress string `json:"verifiable_address"`
}

// NewContinueWithVerificationUiFlow instantiates a new ContinueWithVerificationUiFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithVerificationUiFlow(id string, verifiableAddress string) *ContinueWithVerificationUiFlow {
	this := ContinueWithVerificationUiFlow{}
	this.Id = id
	this.VerifiableAddress = verifiableAddress
	return &this
}

// NewContinueWithVerificationUiFlowWithDefaults instantiates a new ContinueWithVerificationUiFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithVerificationUiFlowWithDefaults() *ContinueWithVerificationUiFlow {
	this := ContinueWithVerificationUiFlow{}
	return &this
}

// GetId returns the Id field value
func (o *ContinueWithVerificationUiFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContinueWithVerificationUiFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContinueWithVerificationUiFlow) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ContinueWithVerificationUiFlow) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinueWithVerificationUiFlow) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ContinueWithVerificationUiFlow) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ContinueWithVerificationUiFlow) SetUrl(v string) {
	o.Url = &v
}

// GetVerifiableAddress returns the VerifiableAddress field value
func (o *ContinueWithVerificationUiFlow) GetVerifiableAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifiableAddress
}

// GetVerifiableAddressOk returns a tuple with the VerifiableAddress field value
// and a boolean to check if the value has been set.
func (o *ContinueWithVerificationUiFlow) GetVerifiableAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifiableAddress, true
}

// SetVerifiableAddress sets field value
func (o *ContinueWithVerificationUiFlow) SetVerifiableAddress(v string) {
	o.VerifiableAddress = v
}

func (o ContinueWithVerificationUiFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["verifiable_address"] = o.VerifiableAddress
	}
	return json.Marshal(toSerialize)
}

type NullableContinueWithVerificationUiFlow struct {
	value *ContinueWithVerificationUiFlow
	isSet bool
}

func (v NullableContinueWithVerificationUiFlow) Get() *ContinueWithVerificationUiFlow {
	return v.value
}

func (v *NullableContinueWithVerificationUiFlow) Set(val *ContinueWithVerificationUiFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithVerificationUiFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithVerificationUiFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithVerificationUiFlow(val *ContinueWithVerificationUiFlow) *NullableContinueWithVerificationUiFlow {
	return &NullableContinueWithVerificationUiFlow{value: val, isSet: true}
}

func (v NullableContinueWithVerificationUiFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithVerificationUiFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


