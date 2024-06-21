/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.11
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OAuth2TokenExchange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2TokenExchange{}

// OAuth2TokenExchange OAuth2 Token Exchange Result
type OAuth2TokenExchange struct {
	// The access token issued by the authorization server.
	AccessToken *string `json:"access_token,omitempty"`
	// The lifetime in seconds of the access token. For example, the value \"3600\" denotes that the access token will expire in one hour from the time the response was generated.
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// To retrieve a refresh token request the id_token scope.
	IdToken *string `json:"id_token,omitempty"`
	// The refresh token, which can be used to obtain new access tokens. To retrieve it add the scope \"offline\" to your access token request.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scope of the access token
	Scope *string `json:"scope,omitempty"`
	// The type of the token issued
	TokenType *string `json:"token_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2TokenExchange OAuth2TokenExchange

// NewOAuth2TokenExchange instantiates a new OAuth2TokenExchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2TokenExchange() *OAuth2TokenExchange {
	this := OAuth2TokenExchange{}
	return &this
}

// NewOAuth2TokenExchangeWithDefaults instantiates a new OAuth2TokenExchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2TokenExchangeWithDefaults() *OAuth2TokenExchange {
	this := OAuth2TokenExchange{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OAuth2TokenExchange) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetExpiresInOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int64 and assigns it to the ExpiresIn field.
func (o *OAuth2TokenExchange) SetExpiresIn(v int64) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetIdToken() string {
	if o == nil || IsNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetIdTokenOk() (*string, bool) {
	if o == nil || IsNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasIdToken() bool {
	if o != nil && !IsNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *OAuth2TokenExchange) SetIdToken(v string) {
	o.IdToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *OAuth2TokenExchange) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OAuth2TokenExchange) SetScope(v string) {
	o.Scope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *OAuth2TokenExchange) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2TokenExchange) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *OAuth2TokenExchange) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *OAuth2TokenExchange) SetTokenType(v string) {
	o.TokenType = &v
}

func (o OAuth2TokenExchange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2TokenExchange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !IsNil(o.IdToken) {
		toSerialize["id_token"] = o.IdToken
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2TokenExchange) UnmarshalJSON(data []byte) (err error) {
	varOAuth2TokenExchange := _OAuth2TokenExchange{}

	err = json.Unmarshal(data, &varOAuth2TokenExchange)

	if err != nil {
		return err
	}

	*o = OAuth2TokenExchange(varOAuth2TokenExchange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "id_token")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2TokenExchange struct {
	value *OAuth2TokenExchange
	isSet bool
}

func (v NullableOAuth2TokenExchange) Get() *OAuth2TokenExchange {
	return v.value
}

func (v *NullableOAuth2TokenExchange) Set(val *OAuth2TokenExchange) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2TokenExchange) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2TokenExchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2TokenExchange(val *OAuth2TokenExchange) *NullableOAuth2TokenExchange {
	return &NullableOAuth2TokenExchange{value: val, isSet: true}
}

func (v NullableOAuth2TokenExchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2TokenExchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


