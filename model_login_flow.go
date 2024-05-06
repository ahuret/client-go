/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.6
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the LoginFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginFlow{}

// LoginFlow This object represents a login flow. A login flow is initiated at the \"Initiate Login API / Browser Flow\" endpoint by a client.  Once a login flow is completed successfully, a session cookie or session token will be issued.
type LoginFlow struct {
	// The active login method  If set contains the login method used. If the flow is new, it is unset. password CredentialsTypePassword oidc CredentialsTypeOIDC totp CredentialsTypeTOTP lookup_secret CredentialsTypeLookup webauthn CredentialsTypeWebAuthn code CredentialsTypeCodeAuth passkey CredentialsTypePasskey profile CredentialsTypeProfile link_recovery CredentialsTypeRecoveryLink  CredentialsTypeRecoveryLink is a special credential type linked to the link strategy (recovery flow).  It is not used within the credentials object itself. code_recovery CredentialsTypeRecoveryCode
	Active *string `json:"active,omitempty"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the flow's unique ID. When performing the login flow, this represents the id in the login UI's query parameter: http://<selfservice.flows.login.ui_url>/?flow=<flow_id>
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the flow started.
	IssuedAt time.Time `json:"issued_at"`
	// Ory OAuth 2.0 Login Challenge.  This value is set using the `login_challenge` query parameter of the registration and login endpoints. If set will cooperate with Ory OAuth2 and OpenID to act as an OAuth2 server / OpenID Provider.
	Oauth2LoginChallenge *string `json:"oauth2_login_challenge,omitempty"`
	Oauth2LoginRequest *OAuth2LoginRequest `json:"oauth2_login_request,omitempty"`
	OrganizationId NullableString `json:"organization_id,omitempty"`
	// Refresh stores whether this login flow should enforce re-authentication.
	Refresh *bool `json:"refresh,omitempty"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	RequestedAal *AuthenticatorAssuranceLevel `json:"requested_aal,omitempty"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// SessionTokenExchangeCode holds the secret code that the client can use to retrieve a session token after the login flow has been completed. This is only set if the client has requested a session token exchange code, and if the flow is of type \"api\", and only on creating the login flow.
	SessionTokenExchangeCode *string `json:"session_token_exchange_code,omitempty"`
	// State represents the state of this request:  choose_method: ask the user to choose a method to sign in with sent_email: the email has been sent to the user passed_challenge: the request was successful and the login challenge was passed.
	State interface{} `json:"state"`
	// TransientPayload is used to pass data from the login to hooks and email templates
	TransientPayload map[string]interface{} `json:"transient_payload,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginFlow LoginFlow

// NewLoginFlow instantiates a new LoginFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, state interface{}, type_ string, ui UiContainer) *LoginFlow {
	this := LoginFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
	this.State = state
	this.Type = type_
	this.Ui = ui
	return &this
}

// NewLoginFlowWithDefaults instantiates a new LoginFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginFlowWithDefaults() *LoginFlow {
	this := LoginFlow{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *LoginFlow) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LoginFlow) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *LoginFlow) SetActive(v string) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoginFlow) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoginFlow) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoginFlow) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *LoginFlow) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *LoginFlow) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *LoginFlow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoginFlow) SetId(v string) {
	o.Id = v
}

// GetIssuedAt returns the IssuedAt field value
func (o *LoginFlow) GetIssuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.IssuedAt
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuedAt, true
}

// SetIssuedAt sets field value
func (o *LoginFlow) SetIssuedAt(v time.Time) {
	o.IssuedAt = v
}

// GetOauth2LoginChallenge returns the Oauth2LoginChallenge field value if set, zero value otherwise.
func (o *LoginFlow) GetOauth2LoginChallenge() string {
	if o == nil || IsNil(o.Oauth2LoginChallenge) {
		var ret string
		return ret
	}
	return *o.Oauth2LoginChallenge
}

// GetOauth2LoginChallengeOk returns a tuple with the Oauth2LoginChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetOauth2LoginChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.Oauth2LoginChallenge) {
		return nil, false
	}
	return o.Oauth2LoginChallenge, true
}

// HasOauth2LoginChallenge returns a boolean if a field has been set.
func (o *LoginFlow) HasOauth2LoginChallenge() bool {
	if o != nil && !IsNil(o.Oauth2LoginChallenge) {
		return true
	}

	return false
}

// SetOauth2LoginChallenge gets a reference to the given string and assigns it to the Oauth2LoginChallenge field.
func (o *LoginFlow) SetOauth2LoginChallenge(v string) {
	o.Oauth2LoginChallenge = &v
}

// GetOauth2LoginRequest returns the Oauth2LoginRequest field value if set, zero value otherwise.
func (o *LoginFlow) GetOauth2LoginRequest() OAuth2LoginRequest {
	if o == nil || IsNil(o.Oauth2LoginRequest) {
		var ret OAuth2LoginRequest
		return ret
	}
	return *o.Oauth2LoginRequest
}

// GetOauth2LoginRequestOk returns a tuple with the Oauth2LoginRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetOauth2LoginRequestOk() (*OAuth2LoginRequest, bool) {
	if o == nil || IsNil(o.Oauth2LoginRequest) {
		return nil, false
	}
	return o.Oauth2LoginRequest, true
}

// HasOauth2LoginRequest returns a boolean if a field has been set.
func (o *LoginFlow) HasOauth2LoginRequest() bool {
	if o != nil && !IsNil(o.Oauth2LoginRequest) {
		return true
	}

	return false
}

// SetOauth2LoginRequest gets a reference to the given OAuth2LoginRequest and assigns it to the Oauth2LoginRequest field.
func (o *LoginFlow) SetOauth2LoginRequest(v OAuth2LoginRequest) {
	o.Oauth2LoginRequest = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginFlow) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginFlow) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *LoginFlow) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableString and assigns it to the OrganizationId field.
func (o *LoginFlow) SetOrganizationId(v string) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *LoginFlow) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *LoginFlow) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *LoginFlow) GetRefresh() bool {
	if o == nil || IsNil(o.Refresh) {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRefreshOk() (*bool, bool) {
	if o == nil || IsNil(o.Refresh) {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *LoginFlow) HasRefresh() bool {
	if o != nil && !IsNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *LoginFlow) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetRequestUrl returns the RequestUrl field value
func (o *LoginFlow) GetRequestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRequestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUrl, true
}

// SetRequestUrl sets field value
func (o *LoginFlow) SetRequestUrl(v string) {
	o.RequestUrl = v
}

// GetRequestedAal returns the RequestedAal field value if set, zero value otherwise.
func (o *LoginFlow) GetRequestedAal() AuthenticatorAssuranceLevel {
	if o == nil || IsNil(o.RequestedAal) {
		var ret AuthenticatorAssuranceLevel
		return ret
	}
	return *o.RequestedAal
}

// GetRequestedAalOk returns a tuple with the RequestedAal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRequestedAalOk() (*AuthenticatorAssuranceLevel, bool) {
	if o == nil || IsNil(o.RequestedAal) {
		return nil, false
	}
	return o.RequestedAal, true
}

// HasRequestedAal returns a boolean if a field has been set.
func (o *LoginFlow) HasRequestedAal() bool {
	if o != nil && !IsNil(o.RequestedAal) {
		return true
	}

	return false
}

// SetRequestedAal gets a reference to the given AuthenticatorAssuranceLevel and assigns it to the RequestedAal field.
func (o *LoginFlow) SetRequestedAal(v AuthenticatorAssuranceLevel) {
	o.RequestedAal = &v
}

// GetReturnTo returns the ReturnTo field value if set, zero value otherwise.
func (o *LoginFlow) GetReturnTo() string {
	if o == nil || IsNil(o.ReturnTo) {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetReturnToOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnTo) {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *LoginFlow) HasReturnTo() bool {
	if o != nil && !IsNil(o.ReturnTo) {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *LoginFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
}

// GetSessionTokenExchangeCode returns the SessionTokenExchangeCode field value if set, zero value otherwise.
func (o *LoginFlow) GetSessionTokenExchangeCode() string {
	if o == nil || IsNil(o.SessionTokenExchangeCode) {
		var ret string
		return ret
	}
	return *o.SessionTokenExchangeCode
}

// GetSessionTokenExchangeCodeOk returns a tuple with the SessionTokenExchangeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetSessionTokenExchangeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SessionTokenExchangeCode) {
		return nil, false
	}
	return o.SessionTokenExchangeCode, true
}

// HasSessionTokenExchangeCode returns a boolean if a field has been set.
func (o *LoginFlow) HasSessionTokenExchangeCode() bool {
	if o != nil && !IsNil(o.SessionTokenExchangeCode) {
		return true
	}

	return false
}

// SetSessionTokenExchangeCode gets a reference to the given string and assigns it to the SessionTokenExchangeCode field.
func (o *LoginFlow) SetSessionTokenExchangeCode(v string) {
	o.SessionTokenExchangeCode = &v
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *LoginFlow) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginFlow) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *LoginFlow) SetState(v interface{}) {
	o.State = v
}

// GetTransientPayload returns the TransientPayload field value if set, zero value otherwise.
func (o *LoginFlow) GetTransientPayload() map[string]interface{} {
	if o == nil || IsNil(o.TransientPayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.TransientPayload
}

// GetTransientPayloadOk returns a tuple with the TransientPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetTransientPayloadOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TransientPayload) {
		return map[string]interface{}{}, false
	}
	return o.TransientPayload, true
}

// HasTransientPayload returns a boolean if a field has been set.
func (o *LoginFlow) HasTransientPayload() bool {
	if o != nil && !IsNil(o.TransientPayload) {
		return true
	}

	return false
}

// SetTransientPayload gets a reference to the given map[string]interface{} and assigns it to the TransientPayload field.
func (o *LoginFlow) SetTransientPayload(v map[string]interface{}) {
	o.TransientPayload = v
}

// GetType returns the Type field value
func (o *LoginFlow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoginFlow) SetType(v string) {
	o.Type = v
}

// GetUi returns the Ui field value
func (o *LoginFlow) GetUi() UiContainer {
	if o == nil {
		var ret UiContainer
		return ret
	}

	return o.Ui
}

// GetUiOk returns a tuple with the Ui field value
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetUiOk() (*UiContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ui, true
}

// SetUi sets field value
func (o *LoginFlow) SetUi(v UiContainer) {
	o.Ui = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LoginFlow) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoginFlow) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoginFlow) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o LoginFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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
	if !IsNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	toSerialize["request_url"] = o.RequestUrl
	if !IsNil(o.RequestedAal) {
		toSerialize["requested_aal"] = o.RequestedAal
	}
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
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginFlow) UnmarshalJSON(data []byte) (err error) {
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

	varLoginFlow := _LoginFlow{}

	err = json.Unmarshal(data, &varLoginFlow)

	if err != nil {
		return err
	}

	*o = LoginFlow(varLoginFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "oauth2_login_challenge")
		delete(additionalProperties, "oauth2_login_request")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "refresh")
		delete(additionalProperties, "request_url")
		delete(additionalProperties, "requested_aal")
		delete(additionalProperties, "return_to")
		delete(additionalProperties, "session_token_exchange_code")
		delete(additionalProperties, "state")
		delete(additionalProperties, "transient_payload")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ui")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginFlow struct {
	value *LoginFlow
	isSet bool
}

func (v NullableLoginFlow) Get() *LoginFlow {
	return v.value
}

func (v *NullableLoginFlow) Set(val *LoginFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlow(val *LoginFlow) *NullableLoginFlow {
	return &NullableLoginFlow{value: val, isSet: true}
}

func (v NullableLoginFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


