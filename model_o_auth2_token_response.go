/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.20
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OAuth2TokenResponse OAuth2 Token Response
type OAuth2TokenResponse struct {
	// The access token issued by the authorization server.
	AccessToken *string `json:"access_token,omitempty"`
	// The lifetime in seconds of the access token.  For example, the value \"3600\" denotes that the access token will expire in one hour from the time the response was generated.
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// To retrieve a refresh token request the id_token scope.
	IdToken *int64 `json:"id_token,omitempty"`
	// The refresh token, which can be used to obtain new access tokens. To retrieve it add the scope \"offline\" to your access token request.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scope of the access token
	Scope *int64 `json:"scope,omitempty"`
	// The type of the token issued
	TokenType *string `json:"token_type,omitempty"`
}

// NewOAuth2TokenResponse instantiates a new OAuth2TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2TokenResponse() *OAuth2TokenResponse {
	this := OAuth2TokenResponse{}
	return &this
}

// NewOAuth2TokenResponseWithDefaults instantiates a new OAuth2TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2TokenResponseWithDefaults() *OAuth2TokenResponse {
	this := OAuth2TokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OAuth2TokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetExpiresIn() int64 {
	if o == nil || o.ExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *OAuth2TokenResponse) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetIdToken() int64 {
	if o == nil || o.IdToken == nil {
		var ret int64
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetIdTokenOk() (*int64, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given int64 and assigns it to the IdToken field.
func (o *OAuth2TokenResponse) SetIdToken(v int64) {
	o.IdToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *OAuth2TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetScope() int64 {
	if o == nil || o.Scope == nil {
		var ret int64
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetScopeOk() (*int64, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given int64 and assigns it to the Scope field.
func (o *OAuth2TokenResponse) SetScope(v int64) {
	o.Scope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *OAuth2TokenResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *OAuth2TokenResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *OAuth2TokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

func (o OAuth2TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2TokenResponse struct {
	value *OAuth2TokenResponse
	isSet bool
}

func (v NullableOAuth2TokenResponse) Get() *OAuth2TokenResponse {
	return v.value
}

func (v *NullableOAuth2TokenResponse) Set(val *OAuth2TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2TokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2TokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2TokenResponse(val *OAuth2TokenResponse) *NullableOAuth2TokenResponse {
	return &NullableOAuth2TokenResponse{value: val, isSet: true}
}

func (v NullableOAuth2TokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2TokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


