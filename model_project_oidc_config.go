/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.38
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectOidcConfig struct for ProjectOidcConfig
type ProjectOidcConfig struct {
	// AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when `provider` is set to `generic`.
	AuthUrl *string `json:"auth_url,omitempty"`
	// ClientID is the application's Client ID.
	ClientId *string `json:"client_id,omitempty"`
	// ClientSecret is the application's secret.
	ClientSecret *string `json:"client_secret,omitempty"`
	// ID is the provider's ID
	Id *string `json:"id,omitempty"`
	// IssuerURL is the OpenID Connect Server URL. You can leave this empty if `provider` is not set to `generic`. If set, neither `auth_url` nor `token_url` are required.
	IssuerUrl *string `json:"issuer_url,omitempty"`
	// Label represents an optional label which can be used in the UI generation.
	Label *string `json:"label,omitempty"`
	// Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider's data (e.g. GitHub or Google profile information) to hydrate the identity's data.  It can be either a URL (file://, http(s)://, base64://) or an inline JSONNet code snippet.
	MapperUrl *string `json:"mapper_url,omitempty"`
	// Provider is either \"generic\" for a generic OAuth 2.0 / OpenID Connect Provider or one of: generic google github gitlab microsoft discord slack facebook vk yandex
	Provider *string `json:"provider,omitempty"`
	// RequestedClaims string encoded json object that specifies claims and optionally their properties which should be included in the id_token or returned from the UserInfo Endpoint.  More information: https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter
	RequestedClaims map[string]interface{} `json:"requested_claims,omitempty"`
	// Scope specifies optional requested permissions.
	Scope []string `json:"scope,omitempty"`
	String *string `json:"string,omitempty"`
	// Tenant is the Azure AD Tenant to use for authentication, and must be set when `provider` is set to `microsoft`. Can be either `common`, `organizations`, `consumers` for a multitenant application or a specific tenant like `8eaef023-2b34-4da1-9baa-8bc8c9d6a490` or `contoso.onmicrosoft.com`.
	Tenant *string `json:"tenant,omitempty"`
	// TokenURL is the token url, typically something like: https://example.org/oauth2/token Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when `provider` is set to `generic`.
	TokenUrl *string `json:"token_url,omitempty"`
}

// NewProjectOidcConfig instantiates a new ProjectOidcConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectOidcConfig() *ProjectOidcConfig {
	this := ProjectOidcConfig{}
	return &this
}

// NewProjectOidcConfigWithDefaults instantiates a new ProjectOidcConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectOidcConfigWithDefaults() *ProjectOidcConfig {
	this := ProjectOidcConfig{}
	return &this
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *ProjectOidcConfig) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProjectOidcConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ProjectOidcConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectOidcConfig) SetId(v string) {
	o.Id = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl != nil {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *ProjectOidcConfig) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProjectOidcConfig) SetLabel(v string) {
	o.Label = &v
}

// GetMapperUrl returns the MapperUrl field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetMapperUrl() string {
	if o == nil || o.MapperUrl == nil {
		var ret string
		return ret
	}
	return *o.MapperUrl
}

// GetMapperUrlOk returns a tuple with the MapperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetMapperUrlOk() (*string, bool) {
	if o == nil || o.MapperUrl == nil {
		return nil, false
	}
	return o.MapperUrl, true
}

// HasMapperUrl returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasMapperUrl() bool {
	if o != nil && o.MapperUrl != nil {
		return true
	}

	return false
}

// SetMapperUrl gets a reference to the given string and assigns it to the MapperUrl field.
func (o *ProjectOidcConfig) SetMapperUrl(v string) {
	o.MapperUrl = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ProjectOidcConfig) SetProvider(v string) {
	o.Provider = &v
}

// GetRequestedClaims returns the RequestedClaims field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetRequestedClaims() map[string]interface{} {
	if o == nil || o.RequestedClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RequestedClaims
}

// GetRequestedClaimsOk returns a tuple with the RequestedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetRequestedClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.RequestedClaims == nil {
		return nil, false
	}
	return o.RequestedClaims, true
}

// HasRequestedClaims returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasRequestedClaims() bool {
	if o != nil && o.RequestedClaims != nil {
		return true
	}

	return false
}

// SetRequestedClaims gets a reference to the given map[string]interface{} and assigns it to the RequestedClaims field.
func (o *ProjectOidcConfig) SetRequestedClaims(v map[string]interface{}) {
	o.RequestedClaims = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetScope() []string {
	if o == nil || o.Scope == nil {
		var ret []string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetScopeOk() ([]string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given []string and assigns it to the Scope field.
func (o *ProjectOidcConfig) SetScope(v []string) {
	o.Scope = v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *ProjectOidcConfig) SetString(v string) {
	o.String = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *ProjectOidcConfig) SetTenant(v string) {
	o.Tenant = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *ProjectOidcConfig) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOidcConfig) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *ProjectOidcConfig) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *ProjectOidcConfig) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o ProjectOidcConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthUrl != nil {
		toSerialize["auth_url"] = o.AuthUrl
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MapperUrl != nil {
		toSerialize["mapper_url"] = o.MapperUrl
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.RequestedClaims != nil {
		toSerialize["requested_claims"] = o.RequestedClaims
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullableProjectOidcConfig struct {
	value *ProjectOidcConfig
	isSet bool
}

func (v NullableProjectOidcConfig) Get() *ProjectOidcConfig {
	return v.value
}

func (v *NullableProjectOidcConfig) Set(val *ProjectOidcConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectOidcConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectOidcConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectOidcConfig(val *ProjectOidcConfig) *NullableProjectOidcConfig {
	return &NullableProjectOidcConfig{value: val, isSet: true}
}

func (v NullableProjectOidcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectOidcConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


