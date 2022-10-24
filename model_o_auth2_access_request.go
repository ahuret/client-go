/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.60
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OAuth2AccessRequest struct for OAuth2AccessRequest
type OAuth2AccessRequest struct {
	// ClientID is the identifier of the OAuth 2.0 client.
	ClientId *string `json:"client_id,omitempty"`
	// GrantTypes is the requests grant types.
	GrantTypes []string `json:"grant_types,omitempty"`
	// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
	GrantedAudience []string `json:"granted_audience,omitempty"`
	// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
	GrantedScopes []string `json:"granted_scopes,omitempty"`
}

// NewOAuth2AccessRequest instantiates a new OAuth2AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2AccessRequest() *OAuth2AccessRequest {
	this := OAuth2AccessRequest{}
	return &this
}

// NewOAuth2AccessRequestWithDefaults instantiates a new OAuth2AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2AccessRequestWithDefaults() *OAuth2AccessRequest {
	this := OAuth2AccessRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuth2AccessRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuth2AccessRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuth2AccessRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *OAuth2AccessRequest) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRequest) GetGrantTypesOk() ([]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *OAuth2AccessRequest) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *OAuth2AccessRequest) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetGrantedAudience returns the GrantedAudience field value if set, zero value otherwise.
func (o *OAuth2AccessRequest) GetGrantedAudience() []string {
	if o == nil || o.GrantedAudience == nil {
		var ret []string
		return ret
	}
	return o.GrantedAudience
}

// GetGrantedAudienceOk returns a tuple with the GrantedAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRequest) GetGrantedAudienceOk() ([]string, bool) {
	if o == nil || o.GrantedAudience == nil {
		return nil, false
	}
	return o.GrantedAudience, true
}

// HasGrantedAudience returns a boolean if a field has been set.
func (o *OAuth2AccessRequest) HasGrantedAudience() bool {
	if o != nil && o.GrantedAudience != nil {
		return true
	}

	return false
}

// SetGrantedAudience gets a reference to the given []string and assigns it to the GrantedAudience field.
func (o *OAuth2AccessRequest) SetGrantedAudience(v []string) {
	o.GrantedAudience = v
}

// GetGrantedScopes returns the GrantedScopes field value if set, zero value otherwise.
func (o *OAuth2AccessRequest) GetGrantedScopes() []string {
	if o == nil || o.GrantedScopes == nil {
		var ret []string
		return ret
	}
	return o.GrantedScopes
}

// GetGrantedScopesOk returns a tuple with the GrantedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2AccessRequest) GetGrantedScopesOk() ([]string, bool) {
	if o == nil || o.GrantedScopes == nil {
		return nil, false
	}
	return o.GrantedScopes, true
}

// HasGrantedScopes returns a boolean if a field has been set.
func (o *OAuth2AccessRequest) HasGrantedScopes() bool {
	if o != nil && o.GrantedScopes != nil {
		return true
	}

	return false
}

// SetGrantedScopes gets a reference to the given []string and assigns it to the GrantedScopes field.
func (o *OAuth2AccessRequest) SetGrantedScopes(v []string) {
	o.GrantedScopes = v
}

func (o OAuth2AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.GrantedAudience != nil {
		toSerialize["granted_audience"] = o.GrantedAudience
	}
	if o.GrantedScopes != nil {
		toSerialize["granted_scopes"] = o.GrantedScopes
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2AccessRequest struct {
	value *OAuth2AccessRequest
	isSet bool
}

func (v NullableOAuth2AccessRequest) Get() *OAuth2AccessRequest {
	return v.value
}

func (v *NullableOAuth2AccessRequest) Set(val *OAuth2AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2AccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2AccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2AccessRequest(val *OAuth2AccessRequest) *NullableOAuth2AccessRequest {
	return &NullableOAuth2AccessRequest{value: val, isSet: true}
}

func (v NullableOAuth2AccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2AccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


