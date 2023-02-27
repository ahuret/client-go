/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.20
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OAuth2LogoutRequest struct for OAuth2LogoutRequest
type OAuth2LogoutRequest struct {
	// Challenge is the identifier (\"logout challenge\") of the logout authentication request. It is used to identify the session.
	Challenge *string `json:"challenge,omitempty"`
	Client *OAuth2Client `json:"client,omitempty"`
	// RequestURL is the original Logout URL requested.
	RequestUrl *string `json:"request_url,omitempty"`
	// RPInitiated is set to true if the request was initiated by a Relying Party (RP), also known as an OAuth 2.0 Client.
	RpInitiated *bool `json:"rp_initiated,omitempty"`
	// SessionID is the login session ID that was requested to log out.
	Sid *string `json:"sid,omitempty"`
	// Subject is the user for whom the logout was request.
	Subject *string `json:"subject,omitempty"`
}

// NewOAuth2LogoutRequest instantiates a new OAuth2LogoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2LogoutRequest() *OAuth2LogoutRequest {
	this := OAuth2LogoutRequest{}
	return &this
}

// NewOAuth2LogoutRequestWithDefaults instantiates a new OAuth2LogoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2LogoutRequestWithDefaults() *OAuth2LogoutRequest {
	this := OAuth2LogoutRequest{}
	return &this
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetChallenge() string {
	if o == nil || o.Challenge == nil {
		var ret string
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetChallengeOk() (*string, bool) {
	if o == nil || o.Challenge == nil {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasChallenge() bool {
	if o != nil && o.Challenge != nil {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given string and assigns it to the Challenge field.
func (o *OAuth2LogoutRequest) SetChallenge(v string) {
	o.Challenge = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetClient() OAuth2Client {
	if o == nil || o.Client == nil {
		var ret OAuth2Client
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetClientOk() (*OAuth2Client, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given OAuth2Client and assigns it to the Client field.
func (o *OAuth2LogoutRequest) SetClient(v OAuth2Client) {
	o.Client = &v
}

// GetRequestUrl returns the RequestUrl field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetRequestUrl() string {
	if o == nil || o.RequestUrl == nil {
		var ret string
		return ret
	}
	return *o.RequestUrl
}

// GetRequestUrlOk returns a tuple with the RequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetRequestUrlOk() (*string, bool) {
	if o == nil || o.RequestUrl == nil {
		return nil, false
	}
	return o.RequestUrl, true
}

// HasRequestUrl returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasRequestUrl() bool {
	if o != nil && o.RequestUrl != nil {
		return true
	}

	return false
}

// SetRequestUrl gets a reference to the given string and assigns it to the RequestUrl field.
func (o *OAuth2LogoutRequest) SetRequestUrl(v string) {
	o.RequestUrl = &v
}

// GetRpInitiated returns the RpInitiated field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetRpInitiated() bool {
	if o == nil || o.RpInitiated == nil {
		var ret bool
		return ret
	}
	return *o.RpInitiated
}

// GetRpInitiatedOk returns a tuple with the RpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetRpInitiatedOk() (*bool, bool) {
	if o == nil || o.RpInitiated == nil {
		return nil, false
	}
	return o.RpInitiated, true
}

// HasRpInitiated returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasRpInitiated() bool {
	if o != nil && o.RpInitiated != nil {
		return true
	}

	return false
}

// SetRpInitiated gets a reference to the given bool and assigns it to the RpInitiated field.
func (o *OAuth2LogoutRequest) SetRpInitiated(v bool) {
	o.RpInitiated = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetSid() string {
	if o == nil || o.Sid == nil {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetSidOk() (*string, bool) {
	if o == nil || o.Sid == nil {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasSid() bool {
	if o != nil && o.Sid != nil {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *OAuth2LogoutRequest) SetSid(v string) {
	o.Sid = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *OAuth2LogoutRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2LogoutRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *OAuth2LogoutRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *OAuth2LogoutRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o OAuth2LogoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Challenge != nil {
		toSerialize["challenge"] = o.Challenge
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.RequestUrl != nil {
		toSerialize["request_url"] = o.RequestUrl
	}
	if o.RpInitiated != nil {
		toSerialize["rp_initiated"] = o.RpInitiated
	}
	if o.Sid != nil {
		toSerialize["sid"] = o.Sid
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2LogoutRequest struct {
	value *OAuth2LogoutRequest
	isSet bool
}

func (v NullableOAuth2LogoutRequest) Get() *OAuth2LogoutRequest {
	return v.value
}

func (v *NullableOAuth2LogoutRequest) Set(val *OAuth2LogoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2LogoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2LogoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2LogoutRequest(val *OAuth2LogoutRequest) *NullableOAuth2LogoutRequest {
	return &NullableOAuth2LogoutRequest{value: val, isSet: true}
}

func (v NullableOAuth2LogoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2LogoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


