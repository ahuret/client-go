/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the VerificationFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationFlow{}

// VerificationFlow Used to verify an out-of-band communication channel such as an email address or a phone number.  For more information head over to: https://www.ory.sh/docs/kratos/self-service/flows/verify-email-account-activation
type VerificationFlow struct {
	// Active, if set, contains the registration method that is being used. It is initially not set.
	Active *string `json:"active,omitempty"`
	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to verify the address, a new request has to be initiated.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ID represents the request's unique ID. When performing the verification flow, this represents the id in the verify ui's query parameter: http://<selfservice.flows.verification.ui_url>?request=<id>  type: string format: uuid
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the request occurred.
	IssuedAt *time.Time `json:"issued_at,omitempty"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl *string `json:"request_url,omitempty"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// State represents the state of this request:  choose_method: ask the user to choose a method (e.g. verify your email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the verification challenge was passed.
	State interface{} `json:"state"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	AdditionalProperties map[string]interface{}
}

type _VerificationFlow VerificationFlow

// NewVerificationFlow instantiates a new VerificationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationFlow(id string, state interface{}, type_ string, ui UiContainer) *VerificationFlow {
	this := VerificationFlow{}
	this.Id = id
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewVerificationFlowWithDefaults instantiates a new VerificationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationFlowWithDefaults() *VerificationFlow {
	this := VerificationFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *VerificationFlow) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *VerificationFlow) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *VerificationFlow) SetActive(v string) {
	o.Active = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *VerificationFlow) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *VerificationFlow) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *VerificationFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value
func (o *VerificationFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VerificationFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise.
func (o *VerificationFlow) GetIssuedAt() time.Time {
	if o == nil || IsNil(o.IssuedAt) {
		var ret time.Time
		return ret
	}
	return *o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.IssuedAt) {
		return nil, false
	}
	return o.IssuedAt, true
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *VerificationFlow) HasIssuedAt() bool {
	if o != nil && !IsNil(o.IssuedAt) {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given time.Time and assigns it to the IssuedAt field.
func (o *VerificationFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = &v
}

// GetRequestUrl returns the RequestUrl field value if set, zero value otherwise.
func (o *VerificationFlow) GetRequestUrl() string {
	if o == nil || IsNil(o.RequestUrl) {
		var ret string
		return ret
	}
	return *o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RequestUrl) {
		return nil, false
	}
	return o.RequestUrl, true
}

// HasRequestUrl returns a boolean if a field has been set.
func (o *VerificationFlow) HasRequestUrl() bool {
	if o != nil && !IsNil(o.RequestUrl) {
		return true
	}

	return false
}

// SetRequestUrl gets a reference to the given string and assigns it to the RequestUrl field.
func (o *VerificationFlow) SetRequestUrl(v string) {
	o.RequestUrl = &v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *VerificationFlow) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *VerificationFlow) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *VerificationFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VerificationFlow) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerificationFlow) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VerificationFlow) SetState(v interface{}) {
	o.State = v
}

// GetType returns the Type field value
func (o *VerificationFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VerificationFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *VerificationFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *VerificationFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *VerificationFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o VerificationFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.IssuedAt) {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if !IsNil(o.RequestUrl) {
		toSerialize["request_url"] = o.RequestUrl
	}
	if !IsNil(o.ReturnTo) {
		toSerialize["return_to"] = o.ReturnTo
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	toSerialize["type"] = o.Type
	toSerialize["ui"] = o.Ui

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerificationFlow) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"type",
		"ui",
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

	varVerificationFlow := _VerificationFlow{}

	err = json.Unmarshal(bytes, &varVerificationFlow)

	if err != nil {
		return err
	}

	*o = VerificationFlow(varVerificationFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "request_url")
		delete(additionalProperties, "return_to")
		delete(additionalProperties, "state")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ui")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationFlow struct {
	value *VerificationFlow
	isSet bool
}

func (v NullableVerificationFlow) Get() *VerificationFlow {
	return v.value
}

func (v *NullableVerificationFlow) Set(val *VerificationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationFlow(val *VerificationFlow) *NullableVerificationFlow {
	return &NullableVerificationFlow{value: val, isSet: true}
}

func (v NullableVerificationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


