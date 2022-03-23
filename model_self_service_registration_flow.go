/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.145
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// SelfServiceRegistrationFlow struct for SelfServiceRegistrationFlow
type SelfServiceRegistrationFlow struct {
	Active *IdentityCredentialsType `json:"active,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the flow occurred.
	IssuedAt time.Time `json:"issued_at"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type *string `json:"type,omitempty"`
	Ui UiContainer `json:"ui"`
}

// NewSelfServiceRegistrationFlow instantiates a new SelfServiceRegistrationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceRegistrationFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, ui UiContainer) *SelfServiceRegistrationFlow {
	this := SelfServiceRegistrationFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.Ui = ui
	return &this
}

// NewSelfServiceRegistrationFlowWithDefaults instantiates a new SelfServiceRegistrationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceRegistrationFlowWithDefaults() *SelfServiceRegistrationFlow {
	this := SelfServiceRegistrationFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SelfServiceRegistrationFlow) GetActive() IdentityCredentialsType {
	if o == nil || o.Active == nil {
		var ret IdentityCredentialsType
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetActiveOk() (*IdentityCredentialsType, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SelfServiceRegistrationFlow) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given IdentityCredentialsType and assigns it to the Active field.
func (o *SelfServiceRegistrationFlow) SetActive(v IdentityCredentialsType) {
	o.Active = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *SelfServiceRegistrationFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *SelfServiceRegistrationFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *SelfServiceRegistrationFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SelfServiceRegistrationFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *SelfServiceRegistrationFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *SelfServiceRegistrationFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetRequestUrl returns the RequestUrl field value
func (o *SelfServiceRegistrationFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *SelfServiceRegistrationFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *SelfServiceRegistrationFlow) GetReturnTo() string {
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *SelfServiceRegistrationFlow) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *SelfServiceRegistrationFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SelfServiceRegistrationFlow) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SelfServiceRegistrationFlow) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SelfServiceRegistrationFlow) SetType(v string) {
	o.Type = &v
}

// GetUi returns the Ui field value
func (o *SelfServiceRegistrationFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *SelfServiceRegistrationFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *SelfServiceRegistrationFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o SelfServiceRegistrationFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["issued_at"] = o.IssuedAt
	}
	if true {
		toSerialize["request_url"] = o.RequestUrl
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceRegistrationFlow struct {
	value *SelfServiceRegistrationFlow
	isSet bool
}

func (v NullableSelfServiceRegistrationFlow) Get() *SelfServiceRegistrationFlow {
	return v.value
}

func (v *NullableSelfServiceRegistrationFlow) Set(val *SelfServiceRegistrationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceRegistrationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceRegistrationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceRegistrationFlow(val *SelfServiceRegistrationFlow) *NullableSelfServiceRegistrationFlow {
	return &NullableSelfServiceRegistrationFlow{value: val, isSet: true}
}

func (v NullableSelfServiceRegistrationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceRegistrationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


