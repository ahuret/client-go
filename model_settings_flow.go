/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.10
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SettingsFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsFlow{}

// SettingsFlow This flow is used when an identity wants to update settings (e.g. profile data, passwords, ...) in a selfservice manner.  We recommend reading the [User Settings Documentation](../self-service/flows/user-settings)
type SettingsFlow struct {
	// Active, if set, contains the registration method that is being used. It is initially not set.
	Active *string `json:"active,omitempty"`
	// Contains a list of actions, that could follow this flow  It can, for example, contain a reference to the verification flow, created as part of the user's registration.
	ContinueWith []ContinueWith `json:"continue_with,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to update the setting, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the flow's unique ID. When performing the settings flow, this represents the id in the settings ui's query parameter: http://<selfservice.flows.settings.ui_url>?flow=<id>
	Id string `json:"id"`
	Identity Identity `json:"identity"`
	// IssuedAt is the time (UTC) when the flow occurred.
	IssuedAt time.Time `json:"issued_at"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// State represents the state of this flow. It knows two states:  show_form: No user data has been collected, or it is invalid, and thus the form should be shown. success: Indicates that the settings flow has been updated successfully with the provided data. Done will stay true when repeatedly checking. If set to true, done will revert back to false only when a flow with invalid (e.g. \"please use a valid phone number\") data was sent.
	State interface{} `json:"state"`
	// TransientPayload is used to pass data from the settings flow to hooks and email templates
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	AdditionalProperties map[string]interface{}
}

type _SettingsFlow SettingsFlow

// NewSettingsFlow instantiates a new SettingsFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsFlow(expiresAt time.Time, id string, identity Identity, issuedAt time.Time, requestUrl string, state interface{}, type_ string, ui UiContainer) *SettingsFlow {
	this := SettingsFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.Identity = identity
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewSettingsFlowWithDefaults instantiates a new SettingsFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsFlowWithDefaults() *SettingsFlow {
	this := SettingsFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SettingsFlow) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SettingsFlow) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *SettingsFlow) SetActive(v string) {
	o.Active = &v
}

// GetContinueWith returns the ContinueWith field value if set, zero value otherwise.
func (o *SettingsFlow) GetContinueWith() []ContinueWith {
	if o == nil || IsNil(o.ContinueWith) {
		var ret []ContinueWith
		return ret
	}
	return o.ContinueWith
}

// GetContinueWithOk returns a tuple with the ContinueWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetContinueWithOk() ([]ContinueWith, bool) {
	if o == nil || IsNil(o.ContinueWith) {
		return nil, false
	}
	return o.ContinueWith, true
}

// HasContinueWith returns a boolean if a field has been set.
func (o *SettingsFlow) HasContinueWith() bool {
	if o != nil && !IsNil(o.ContinueWith) {
		return true
	}

	return false
}

// SetContinueWith gets a reference to the given []ContinueWith and assigns it to the ContinueWith field.
func (o *SettingsFlow) SetContinueWith(v []ContinueWith) {
	o.ContinueWith = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *SettingsFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *SettingsFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *SettingsFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SettingsFlow) SetId(v string) {
	o.Id = v
}

// GetIdentity returns the Identity field value
func (o *SettingsFlow) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetIdentityOk() (*Identity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *SettingsFlow) SetIdentity(v Identity) {
	o.Identity = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *SettingsFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *SettingsFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetRequestUrl returns the RequestUrl field value
func (o *SettingsFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *SettingsFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *SettingsFlow) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *SettingsFlow) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *SettingsFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SettingsFlow) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingsFlow) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SettingsFlow) SetState(v interface{}) {
	o.State = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *SettingsFlow) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *SettingsFlow) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *SettingsFlow) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

// GetType returns the Type field value
func (o *SettingsFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SettingsFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *SettingsFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *SettingsFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *SettingsFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o SettingsFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ContinueWith) {
		toSerialize["continue_with"] = o.ContinueWith
	}
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["id"] = o.Id
	toSerialize["identity"] = o.Identity
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

func (o *SettingsFlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expires_at",
		"id",
		"identity",
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

	varSettingsFlow := _SettingsFlow{}

	err = json.Unmarshal(data, &varSettingsFlow)

	if err != nil {
		return err
	}

	*o = SettingsFlow(varSettingsFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "continue_with")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identity")
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

type NullableSettingsFlow struct {
	value *SettingsFlow
	isSet bool
}

func (v NullableSettingsFlow) Get() *SettingsFlow {
	return v.value
}

func (v *NullableSettingsFlow) Set(val *SettingsFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsFlow(val *SettingsFlow) *NullableSettingsFlow {
	return &NullableSettingsFlow{value: val, isSet: true}
}

func (v NullableSettingsFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


