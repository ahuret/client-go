/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.14.5
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the RecoveryFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryFlow{}

// RecoveryFlow This request is used when an identity wants to recover their account.  We recommend reading the [Account Recovery Documentation](../self-service/flows/password-reset-account-recovery)
type RecoveryFlow struct {
	// Active, if set, contains the recovery method that is being used. It is initially not set.
	Active *string `json:"active,omitempty"`
	// Contains possible actions that could follow this flow
	ContinueWith []ContinueWith `json:"continue_with,omitempty"`
	// ExpiresAt is the time (UTC) when the request expires. If the user still wishes to update the setting, a new request has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the request's unique ID. When performing the recovery flow, this represents the id in the recovery ui's query parameter: http://<selfservice.flows.recovery.ui_url>?request=<id>
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the request occurred.
	IssuedAt time.Time `json:"issued_at"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// State represents the state of this request:  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
	State interface{} `json:"state"`
	// TransientPayload is used to pass data from the recovery flow to hooks and email templates
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryFlow RecoveryFlow

// NewRecoveryFlow instantiates a new RecoveryFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, state interface{}, type_ string, ui UiContainer) *RecoveryFlow {
	this := RecoveryFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewRecoveryFlowWithDefaults instantiates a new RecoveryFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryFlowWithDefaults() *RecoveryFlow {
	this := RecoveryFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RecoveryFlow) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RecoveryFlow) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *RecoveryFlow) SetActive(v string) {
	o.Active = &v
}

// GetContinueWith returns the ContinueWith field value if set, zero value otherwise.
func (o *RecoveryFlow) GetContinueWith() []ContinueWith {
	if o == nil || IsNil(o.ContinueWith) {
		var ret []ContinueWith
		return ret
	}
	return o.ContinueWith
}

// GetContinueWithOk returns a tuple with the ContinueWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetContinueWithOk() ([]ContinueWith, bool) {
	if o == nil || IsNil(o.ContinueWith) {
		return nil, false
	}
	return o.ContinueWith, true
}

// HasContinueWith returns a boolean if a field has been set.
func (o *RecoveryFlow) HasContinueWith() bool {
	if o != nil && !IsNil(o.ContinueWith) {
		return true
	}

	return false
}

// SetContinueWith gets a reference to the given []ContinueWith and assigns it to the ContinueWith field.
func (o *RecoveryFlow) SetContinueWith(v []ContinueWith) {
	o.ContinueWith = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *RecoveryFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *RecoveryFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *RecoveryFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecoveryFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *RecoveryFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *RecoveryFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetRequestUrl returns the RequestUrl field value
func (o *RecoveryFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *RecoveryFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *RecoveryFlow) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *RecoveryFlow) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *RecoveryFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RecoveryFlow) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryFlow) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RecoveryFlow) SetState(v interface{}) {
	o.State = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *RecoveryFlow) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *RecoveryFlow) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *RecoveryFlow) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

// GetType returns the Type field value
func (o *RecoveryFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RecoveryFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *RecoveryFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *RecoveryFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *RecoveryFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o RecoveryFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ContinueWith) {
		toSerialize["continue_with"] = o.ContinueWith
	}
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["id"] = o.Id
	toSerialize["issued_at"] = o.IssuedAt
	toSerialize["request_url"] = o.RequestUrl
	if !IsNil(o.ReturnTo) {
		toSerialize["return_to"] = o.ReturnTo
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TransientPayload) {
		toSerialize["transient_payload"] = o.TransientPayload
	}
	toSerialize["type"] = o.Type
	toSerialize["ui"] = o.Ui

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryFlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expires_at",
		"id",
		"issued_at",
		"request_url",
		"state",
		"type",
		"ui",
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

	varRecoveryFlow := _RecoveryFlow{}

	err = json.Unmarshal(data, &varRecoveryFlow)

	if err != nil {
		return err
	}

	*o = RecoveryFlow(varRecoveryFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "continue_with")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "request_url")
		delete(additionalProperties, "return_to")
		delete(additionalProperties, "state")
		delete(additionalProperties, "transient_payload")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ui")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryFlow struct {
	value *RecoveryFlow
	isSet bool
}

func (v NullableRecoveryFlow) Get() *RecoveryFlow {
	return v.value
}

func (v *NullableRecoveryFlow) Set(val *RecoveryFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryFlow(val *RecoveryFlow) *NullableRecoveryFlow {
	return &NullableRecoveryFlow{value: val, isSet: true}
}

func (v NullableRecoveryFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


