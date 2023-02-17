/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.17
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// LoginFlow This object represents a login flow. A login flow is initiated at the \"Initiate Login API / Browser Flow\" endpoint by a client.  Once a login flow is completed successfully, a session cookie or session token will be issued.
type LoginFlow struct {
	Active *IdentityCredentialsType `json:"active,omitempty"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated.
	ExpiresAt time.Time `json:"expires_at"`
	// ID represents the flow's unique ID. When performing the login flow, this represents the id in the login UI's query parameter: http://<selfservice.flows.login.ui_url>/?flow=<flow_id>
	Id string `json:"id"`
	// IssuedAt is the time (UTC) when the flow started.
	IssuedAt time.Time `json:"issued_at"`
	Oauth2LoginChallenge NullableString `json:"oauth2_login_challenge,omitempty"`
	Oauth2LoginRequest *OAuth2LoginRequest `json:"oauth2_login_request,omitempty"`
	// Refresh stores whether this login flow should enforce re-authentication.
	Refresh *bool `json:"refresh,omitempty"`
	// RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL's path or query for example.
	RequestUrl string `json:"request_url"`
	RequestedAal *AuthenticatorAssuranceLevel `json:"requested_aal,omitempty"`
	// ReturnTo contains the requested return_to URL.
	ReturnTo *string `json:"return_to,omitempty"`
	// The flow type can either be `api` or `browser`.
	Type string `json:"type"`
	Ui UiContainer `json:"ui"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewLoginFlow instantiates a new LoginFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, type_ string, ui UiContainer) *LoginFlow {
	this := LoginFlow{}
	this.ExpiresAt = expiresAt
	this.Id = id
	this.IssuedAt = issuedAt
	this.RequestUrl = requestUrl
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
func (o *LoginFlow) GetActive() IdentityCredentialsType {
	if o == nil || o.Active == nil {
		var ret IdentityCredentialsType
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetActiveOk() (*IdentityCredentialsType, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *LoginFlow) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given IdentityCredentialsType and assigns it to the Active field.
func (o *LoginFlow) SetActive(v IdentityCredentialsType) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoginFlow) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoginFlow) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
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

// GetOauth2LoginChallenge returns the Oauth2LoginChallenge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginFlow) GetOauth2LoginChallenge() string {
	if o == nil || o.Oauth2LoginChallenge.Get() == nil {
		var ret string
		return ret
	}
	return *o.Oauth2LoginChallenge.Get()
}

// GetOauth2LoginChallengeOk returns a tuple with the Oauth2LoginChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginFlow) GetOauth2LoginChallengeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Oauth2LoginChallenge.Get(), o.Oauth2LoginChallenge.IsSet()
}

// HasOauth2LoginChallenge returns a boolean if a field has been set.
func (o *LoginFlow) HasOauth2LoginChallenge() bool {
	if o != nil && o.Oauth2LoginChallenge.IsSet() {
		return true
	}

	return false
}

// SetOauth2LoginChallenge gets a reference to the given NullableString and assigns it to the Oauth2LoginChallenge field.
func (o *LoginFlow) SetOauth2LoginChallenge(v string) {
	o.Oauth2LoginChallenge.Set(&v)
}
// SetOauth2LoginChallengeNil sets the value for Oauth2LoginChallenge to be an explicit nil
func (o *LoginFlow) SetOauth2LoginChallengeNil() {
	o.Oauth2LoginChallenge.Set(nil)
}

// UnsetOauth2LoginChallenge ensures that no value is present for Oauth2LoginChallenge, not even an explicit nil
func (o *LoginFlow) UnsetOauth2LoginChallenge() {
	o.Oauth2LoginChallenge.Unset()
}

// GetOauth2LoginRequest returns the Oauth2LoginRequest field value if set, zero value otherwise.
func (o *LoginFlow) GetOauth2LoginRequest() OAuth2LoginRequest {
	if o == nil || o.Oauth2LoginRequest == nil {
		var ret OAuth2LoginRequest
		return ret
	}
	return *o.Oauth2LoginRequest
}

// GetOauth2LoginRequestOk returns a tuple with the Oauth2LoginRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetOauth2LoginRequestOk() (*OAuth2LoginRequest, bool) {
	if o == nil || o.Oauth2LoginRequest == nil {
		return nil, false
	}
	return o.Oauth2LoginRequest, true
}

// HasOauth2LoginRequest returns a boolean if a field has been set.
func (o *LoginFlow) HasOauth2LoginRequest() bool {
	if o != nil && o.Oauth2LoginRequest != nil {
		return true
	}

	return false
}

// SetOauth2LoginRequest gets a reference to the given OAuth2LoginRequest and assigns it to the Oauth2LoginRequest field.
func (o *LoginFlow) SetOauth2LoginRequest(v OAuth2LoginRequest) {
	o.Oauth2LoginRequest = &v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *LoginFlow) GetRefresh() bool {
	if o == nil || o.Refresh == nil {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRefreshOk() (*bool, bool) {
	if o == nil || o.Refresh == nil {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *LoginFlow) HasRefresh() bool {
	if o != nil && o.Refresh != nil {
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
	if o == nil || o.RequestedAal == nil {
		var ret AuthenticatorAssuranceLevel
		return ret
	}
	return *o.RequestedAal
}

// GetRequestedAalOk returns a tuple with the RequestedAal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetRequestedAalOk() (*AuthenticatorAssuranceLevel, bool) {
	if o == nil || o.RequestedAal == nil {
		return nil, false
	}
	return o.RequestedAal, true
}

// HasRequestedAal returns a boolean if a field has been set.
func (o *LoginFlow) HasRequestedAal() bool {
	if o != nil && o.RequestedAal != nil {
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
	if o == nil || o.ReturnTo == nil {
		var ret string
		return ret
	}
	return *o.ReturnTo
}

// GetReturnToOk returns a tuple with the ReturnTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetReturnToOk() (*string, bool) {
	if o == nil || o.ReturnTo == nil {
		return nil, false
	}
	return o.ReturnTo, true
}

// HasReturnTo returns a boolean if a field has been set.
func (o *LoginFlow) HasReturnTo() bool {
	if o != nil && o.ReturnTo != nil {
		return true
	}

	return false
}

// SetReturnTo gets a reference to the given string and assigns it to the ReturnTo field.
func (o *LoginFlow) SetReturnTo(v string) {
	o.ReturnTo = &v
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
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginFlow) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoginFlow) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoginFlow) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o LoginFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
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
	if o.Oauth2LoginChallenge.IsSet() {
		toSerialize["oauth2_login_challenge"] = o.Oauth2LoginChallenge.Get()
	}
	if o.Oauth2LoginRequest != nil {
		toSerialize["oauth2_login_request"] = o.Oauth2LoginRequest
	}
	if o.Refresh != nil {
		toSerialize["refresh"] = o.Refresh
	}
	if true {
		toSerialize["request_url"] = o.RequestUrl
	}
	if o.RequestedAal != nil {
		toSerialize["requested_aal"] = o.RequestedAal
	}
	if o.ReturnTo != nil {
		toSerialize["return_to"] = o.ReturnTo
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["ui"] = o.Ui
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
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


