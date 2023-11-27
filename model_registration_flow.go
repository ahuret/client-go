/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the RegistrationFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationFlow{}

// RegistrationFlow struct for RegistrationFlow
type RegistrationFlow struct {
	Active *IdentityCredentialsType `json:"active,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the flow's unique ID. When performing the registration flow, this represents the id in the registration ui's query parameter: http://<selfservice.flows.registration.ui_url>/?flow=<id>
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the flow occurred.
	IssuedAt time.Time `json:"issued_at"`
	// Ory OAuth 2.0 Login Challenge.  This value is set using the `login_challenge` query parameter of the registration and login endpoints. If set will cooperate with Ory OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.
	Oauth2LoginChallenge *string `json:"oauth2_login_challenge,omitempty"`
	Oauth2LoginRequest *OAuth2LoginRequest `json:"oauth2_login_request,omitempty"`
	OrganizationId NullableString `json:"organization_id,omitempty"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// SessionTokenExchangeCode holds the secret code that the client can use to retrieve a session token after the flow has been completed. This is only set if the client has requested a session token exchange code, and if the flow is of type \"api\", and only on creating the flow.
	SessionTokenExchangeCode *string `json:"session_token_exchange_code,omitempty"`
	// State represents the state of this request:  choose_method: ask the user to choose a method (e.g. registration with email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the registration challenge was passed.
	State interface{} `json:"state"`
	// TransientPayload is used to pass data from the registration to a webhook
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationFlow RegistrationFlow

// NewRegistrationFlow instantiates a new RegistrationFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, state interface{}, type_ string, ui UiContainer) *RegistrationFlow {
	this := RegistrationFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewRegistrationFlowWithDefaults instantiates a new RegistrationFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationFlowWithDefaults() *RegistrationFlow {
	this := RegistrationFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RegistrationFlow) GetActive() IdentityCredentialsType {
	if o == nil || IsNil(o.Active) {
		var ret IdentityCredentialsType
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetActiveOk() (*IdentityCredentialsType, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RegistrationFlow) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given IdentityCredentialsType and assigns it to the Active field.
func (o *RegistrationFlow) SetActive(v IdentityCredentialsType) {
	o.Active = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *RegistrationFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *RegistrationFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *RegistrationFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegistrationFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *RegistrationFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *RegistrationFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetOauth2LoginChallenge returns the Oauth2LoginChallenge field value if set, zero value otherwise.
func (o *RegistrationFlow) GetOauth2LoginChallenge() string {
	if o == nil || IsNil(o.Oauth2LoginChallenge) {
		var ret string
		return ret
	}
	return *o.Oauth2LoginChallenge
}

// GetOauth2LoginChallengeOk returns a tuple with the Oauth2LoginChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetOauth2LoginChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.Oauth2LoginChallenge) {
		return nil, false
	}
	return o.Oauth2LoginChallenge, true
}

// HasOauth2LoginChallenge returns a boolean if a field has been set.
func (o *RegistrationFlow) HasOauth2LoginChallenge() bool {
	if o != nil && !IsNil(o.Oauth2LoginChallenge) {
		return true
	}

	return false
}

// SetOauth2LoginChallenge gets a reference to the given string and assigns it to the Oauth2LoginChallenge field.
func (o *RegistrationFlow) SetOauth2LoginChallenge(v string) {
	o.Oauth2LoginChallenge = &v
}

// GetOauth2LoginRequest returns the Oauth2LoginRequest field value if set, zero value otherwise.
func (o *RegistrationFlow) GetOauth2LoginRequest() OAuth2LoginRequest {
	if o == nil || IsNil(o.Oauth2LoginRequest) {
		var ret OAuth2LoginRequest
		return ret
	}
	return *o.Oauth2LoginRequest
}

// GetOauth2LoginRequestOk returns a tuple with the Oauth2LoginRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetOauth2LoginRequestOk() (*OAuth2LoginRequest, bool) {
	if o == nil || IsNil(o.Oauth2LoginRequest) {
		return nil, false
	}
	return o.Oauth2LoginRequest, true
}

// HasOauth2LoginRequest returns a boolean if a field has been set.
func (o *RegistrationFlow) HasOauth2LoginRequest() bool {
	if o != nil && !IsNil(o.Oauth2LoginRequest) {
		return true
	}

	return false
}

// SetOauth2LoginRequest gets a reference to the given OAuth2LoginRequest and assigns it to the Oauth2LoginRequest field.
func (o *RegistrationFlow) SetOauth2LoginRequest(v OAuth2LoginRequest) {
	o.Oauth2LoginRequest = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegistrationFlow) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistrationFlow) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RegistrationFlow) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableString and assigns it to the OrganizationId field.
func (o *RegistrationFlow) SetOrganizationId(v string) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *RegistrationFlow) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *RegistrationFlow) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetRequestUrl returns the RequestUrl field value
func (o *RegistrationFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *RegistrationFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *RegistrationFlow) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *RegistrationFlow) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *RegistrationFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetSessionTokenExchangeCode returns the SessionTokenExchangeCode field value if set, zero value otherwise.
func (o *RegistrationFlow) GetSessionTokenExchangeCode() string {
	if o == nil || IsNil(o.SessionTokenExchangeCode) {
		var ret string
		return ret
	}
	return *o.SessionTokenExchangeCode
}

// GetSessionTokenExchangeCodeOk returns a tuple with the SessionTokenExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetSessionTokenExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SessionTokenExchangeCode) {
		return nil, false
	}
	return o.SessionTokenExchangeCode, true
}

// HasSessionTokenExchangeCode returns a boolean if a field has been set.
func (o *RegistrationFlow) HasSessionTokenExchangeCode() bool {
	if o != nil && !IsNil(o.SessionTokenExchangeCode) {
		return true
	}

	return false
}

// SetSessionTokenExchangeCode gets a reference to the given string and assigns it to the SessionTokenExchangeCode field.
func (o *RegistrationFlow) SetSessionTokenExchangeCode(v string) {
	o.SessionTokenExchangeCode = &v
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RegistrationFlow) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistrationFlow) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RegistrationFlow) SetState(v interface{}) {
	o.State = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *RegistrationFlow) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *RegistrationFlow) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *RegistrationFlow) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

// GetType returns the Type field value
func (o *RegistrationFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RegistrationFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *RegistrationFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *RegistrationFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *RegistrationFlow) SetUi(v UiContainer) {
	o.Ui = v
}

func (o RegistrationFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["id"] = o.Id
	toSerialize["issued_at"] = o.IssuedAt
	if !IsNil(o.Oauth2LoginChallenge) {
		toSerialize["oauth2_login_challenge"] = o.Oauth2LoginChallenge
	}
	if !IsNil(o.Oauth2LoginRequest) {
		toSerialize["oauth2_login_request"] = o.Oauth2LoginRequest
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organization_id"] = o.OrganizationId.Get()
	}
	toSerialize["request_url"] = o.RequestUrl
	if !IsNil(o.ReturnTo) {
		toSerialize["return_to"] = o.ReturnTo
	}
	if !IsNil(o.SessionTokenExchangeCode) {
		toSerialize["session_token_exchange_code"] = o.SessionTokenExchangeCode
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

func (o *RegistrationFlow) UnmarshalJSON(bytes []byte) (err error) {
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

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegistrationFlow := _RegistrationFlow{}

	err = json.Unmarshal(bytes, &varRegistrationFlow)

	if err != nil {
		return err
	}

	*o = RegistrationFlow(varRegistrationFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "oauth2_login_challenge")
		delete(additionalProperties, "oauth2_login_request")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "request_url")
		delete(additionalProperties, "return_to")
		delete(additionalProperties, "session_token_exchange_code")
		delete(additionalProperties, "state")
		delete(additionalProperties, "transient_payload")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ui")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationFlow struct {
	value *RegistrationFlow
	isSet bool
}

func (v NullableRegistrationFlow) Get() *RegistrationFlow {
	return v.value
}

func (v *NullableRegistrationFlow) Set(val *RegistrationFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationFlow(val *RegistrationFlow) *NullableRegistrationFlow {
	return &NullableRegistrationFlow{value: val, isSet: true}
}

func (v NullableRegistrationFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


