/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.43
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OauthTokenResponse The token response
type OauthTokenResponse struct {
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

// NewOauthTokenResponse instantiates a new OauthTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthTokenResponse() *OauthTokenResponse {
	this := OauthTokenResponse{}
	return &this
}

// NewOauthTokenResponseWithDefaults instantiates a new OauthTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthTokenResponseWithDefaults() *OauthTokenResponse {
	this := OauthTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OauthTokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetExpiresIn() int64 {
	if o == nil || o.ExpiresIn == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *OauthTokenResponse) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetIdToken() int64 {
	if o == nil || o.IdToken == nil {
		var ret int64
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetIdTokenOk() (*int64, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given int64 and assigns it to the IdToken field.
func (o *OauthTokenResponse) SetIdToken(v int64) {
	o.IdToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *OauthTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetScope() int64 {
	if o == nil || o.Scope == nil {
		var ret int64
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetScopeOk() (*int64, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given int64 and assigns it to the Scope field.
func (o *OauthTokenResponse) SetScope(v int64) {
	o.Scope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *OauthTokenResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *OauthTokenResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *OauthTokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

func (o OauthTokenResponse) MarshalJSON() ([]byte, error) {
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

type NullableOauthTokenResponse struct {
	value *OauthTokenResponse
	isSet bool
}

func (v NullableOauthTokenResponse) Get() *OauthTokenResponse {
	return v.value
}

func (v *NullableOauthTokenResponse) Set(val *OauthTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthTokenResponse(val *OauthTokenResponse) *NullableOauthTokenResponse {
	return &NullableOauthTokenResponse{value: val, isSet: true}
}

func (v NullableOauthTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


