/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.19
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateOAuth2ClientLifespans The OIDC Hybrid grant type inherits token lifespan configuration from the implicit grant.
type UpdateOAuth2ClientLifespans struct {
	AuthorizationCodeGrantAccessTokenLifespan NullableString `json:"authorization_code_grant_access_token_lifespan,omitempty"`
	AuthorizationCodeGrantIdTokenLifespan NullableString `json:"authorization_code_grant_id_token_lifespan,omitempty"`
	AuthorizationCodeGrantRefreshTokenLifespan NullableString `json:"authorization_code_grant_refresh_token_lifespan,omitempty"`
	ClientCredentialsGrantAccessTokenLifespan NullableString `json:"client_credentials_grant_access_token_lifespan,omitempty"`
	ImplicitGrantAccessTokenLifespan NullableString `json:"implicit_grant_access_token_lifespan,omitempty"`
	ImplicitGrantIdTokenLifespan NullableString `json:"implicit_grant_id_token_lifespan,omitempty"`
	JwtBearerGrantAccessTokenLifespan NullableString `json:"jwt_bearer_grant_access_token_lifespan,omitempty"`
	PasswordGrantAccessTokenLifespan NullableString `json:"password_grant_access_token_lifespan,omitempty"`
	PasswordGrantRefreshTokenLifespan NullableString `json:"password_grant_refresh_token_lifespan,omitempty"`
	RefreshTokenGrantAccessTokenLifespan NullableString `json:"refresh_token_grant_access_token_lifespan,omitempty"`
	RefreshTokenGrantIdTokenLifespan NullableString `json:"refresh_token_grant_id_token_lifespan,omitempty"`
	RefreshTokenGrantRefreshTokenLifespan NullableString `json:"refresh_token_grant_refresh_token_lifespan,omitempty"`
}

// NewUpdateOAuth2ClientLifespans instantiates a new UpdateOAuth2ClientLifespans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOAuth2ClientLifespans() *UpdateOAuth2ClientLifespans {
	this := UpdateOAuth2ClientLifespans{}
	return &this
}

// NewUpdateOAuth2ClientLifespansWithDefaults instantiates a new UpdateOAuth2ClientLifespans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOAuth2ClientLifespansWithDefaults() *UpdateOAuth2ClientLifespans {
	this := UpdateOAuth2ClientLifespans{}
	return &this
}

// GetAuthorizationCodeGrantAccessTokenLifespan returns the AuthorizationCodeGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantAccessTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantAccessTokenLifespan.Get()
}

// GetAuthorizationCodeGrantAccessTokenLifespanOk returns a tuple with the AuthorizationCodeGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantAccessTokenLifespan.Get(), o.AuthorizationCodeGrantAccessTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasAuthorizationCodeGrantAccessTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantAccessTokenLifespan(v string) {
	o.AuthorizationCodeGrantAccessTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantAccessTokenLifespanNil sets the value for AuthorizationCodeGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantAccessTokenLifespanNil() {
	o.AuthorizationCodeGrantAccessTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantAccessTokenLifespan ensures that no value is present for AuthorizationCodeGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetAuthorizationCodeGrantAccessTokenLifespan() {
	o.AuthorizationCodeGrantAccessTokenLifespan.Unset()
}

// GetAuthorizationCodeGrantIdTokenLifespan returns the AuthorizationCodeGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantIdTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantIdTokenLifespan.Get()
}

// GetAuthorizationCodeGrantIdTokenLifespanOk returns a tuple with the AuthorizationCodeGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantIdTokenLifespan.Get(), o.AuthorizationCodeGrantIdTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasAuthorizationCodeGrantIdTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantIdTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantIdTokenLifespan(v string) {
	o.AuthorizationCodeGrantIdTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantIdTokenLifespanNil sets the value for AuthorizationCodeGrantIdTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantIdTokenLifespanNil() {
	o.AuthorizationCodeGrantIdTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantIdTokenLifespan ensures that no value is present for AuthorizationCodeGrantIdTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetAuthorizationCodeGrantIdTokenLifespan() {
	o.AuthorizationCodeGrantIdTokenLifespan.Unset()
}

// GetAuthorizationCodeGrantRefreshTokenLifespan returns the AuthorizationCodeGrantRefreshTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantRefreshTokenLifespan() string {
	if o == nil || o.AuthorizationCodeGrantRefreshTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationCodeGrantRefreshTokenLifespan.Get()
}

// GetAuthorizationCodeGrantRefreshTokenLifespanOk returns a tuple with the AuthorizationCodeGrantRefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetAuthorizationCodeGrantRefreshTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCodeGrantRefreshTokenLifespan.Get(), o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet()
}

// HasAuthorizationCodeGrantRefreshTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasAuthorizationCodeGrantRefreshTokenLifespan() bool {
	if o != nil && o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantRefreshTokenLifespan gets a reference to the given NullableString and assigns it to the AuthorizationCodeGrantRefreshTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantRefreshTokenLifespan(v string) {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Set(&v)
}
// SetAuthorizationCodeGrantRefreshTokenLifespanNil sets the value for AuthorizationCodeGrantRefreshTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetAuthorizationCodeGrantRefreshTokenLifespanNil() {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Set(nil)
}

// UnsetAuthorizationCodeGrantRefreshTokenLifespan ensures that no value is present for AuthorizationCodeGrantRefreshTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetAuthorizationCodeGrantRefreshTokenLifespan() {
	o.AuthorizationCodeGrantRefreshTokenLifespan.Unset()
}

// GetClientCredentialsGrantAccessTokenLifespan returns the ClientCredentialsGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetClientCredentialsGrantAccessTokenLifespan() string {
	if o == nil || o.ClientCredentialsGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientCredentialsGrantAccessTokenLifespan.Get()
}

// GetClientCredentialsGrantAccessTokenLifespanOk returns a tuple with the ClientCredentialsGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetClientCredentialsGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCredentialsGrantAccessTokenLifespan.Get(), o.ClientCredentialsGrantAccessTokenLifespan.IsSet()
}

// HasClientCredentialsGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasClientCredentialsGrantAccessTokenLifespan() bool {
	if o != nil && o.ClientCredentialsGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetClientCredentialsGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the ClientCredentialsGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetClientCredentialsGrantAccessTokenLifespan(v string) {
	o.ClientCredentialsGrantAccessTokenLifespan.Set(&v)
}
// SetClientCredentialsGrantAccessTokenLifespanNil sets the value for ClientCredentialsGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetClientCredentialsGrantAccessTokenLifespanNil() {
	o.ClientCredentialsGrantAccessTokenLifespan.Set(nil)
}

// UnsetClientCredentialsGrantAccessTokenLifespan ensures that no value is present for ClientCredentialsGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetClientCredentialsGrantAccessTokenLifespan() {
	o.ClientCredentialsGrantAccessTokenLifespan.Unset()
}

// GetImplicitGrantAccessTokenLifespan returns the ImplicitGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetImplicitGrantAccessTokenLifespan() string {
	if o == nil || o.ImplicitGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImplicitGrantAccessTokenLifespan.Get()
}

// GetImplicitGrantAccessTokenLifespanOk returns a tuple with the ImplicitGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetImplicitGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplicitGrantAccessTokenLifespan.Get(), o.ImplicitGrantAccessTokenLifespan.IsSet()
}

// HasImplicitGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasImplicitGrantAccessTokenLifespan() bool {
	if o != nil && o.ImplicitGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetImplicitGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the ImplicitGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetImplicitGrantAccessTokenLifespan(v string) {
	o.ImplicitGrantAccessTokenLifespan.Set(&v)
}
// SetImplicitGrantAccessTokenLifespanNil sets the value for ImplicitGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetImplicitGrantAccessTokenLifespanNil() {
	o.ImplicitGrantAccessTokenLifespan.Set(nil)
}

// UnsetImplicitGrantAccessTokenLifespan ensures that no value is present for ImplicitGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetImplicitGrantAccessTokenLifespan() {
	o.ImplicitGrantAccessTokenLifespan.Unset()
}

// GetImplicitGrantIdTokenLifespan returns the ImplicitGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetImplicitGrantIdTokenLifespan() string {
	if o == nil || o.ImplicitGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImplicitGrantIdTokenLifespan.Get()
}

// GetImplicitGrantIdTokenLifespanOk returns a tuple with the ImplicitGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetImplicitGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplicitGrantIdTokenLifespan.Get(), o.ImplicitGrantIdTokenLifespan.IsSet()
}

// HasImplicitGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasImplicitGrantIdTokenLifespan() bool {
	if o != nil && o.ImplicitGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetImplicitGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the ImplicitGrantIdTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetImplicitGrantIdTokenLifespan(v string) {
	o.ImplicitGrantIdTokenLifespan.Set(&v)
}
// SetImplicitGrantIdTokenLifespanNil sets the value for ImplicitGrantIdTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetImplicitGrantIdTokenLifespanNil() {
	o.ImplicitGrantIdTokenLifespan.Set(nil)
}

// UnsetImplicitGrantIdTokenLifespan ensures that no value is present for ImplicitGrantIdTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetImplicitGrantIdTokenLifespan() {
	o.ImplicitGrantIdTokenLifespan.Unset()
}

// GetJwtBearerGrantAccessTokenLifespan returns the JwtBearerGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetJwtBearerGrantAccessTokenLifespan() string {
	if o == nil || o.JwtBearerGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.JwtBearerGrantAccessTokenLifespan.Get()
}

// GetJwtBearerGrantAccessTokenLifespanOk returns a tuple with the JwtBearerGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetJwtBearerGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JwtBearerGrantAccessTokenLifespan.Get(), o.JwtBearerGrantAccessTokenLifespan.IsSet()
}

// HasJwtBearerGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasJwtBearerGrantAccessTokenLifespan() bool {
	if o != nil && o.JwtBearerGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetJwtBearerGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the JwtBearerGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetJwtBearerGrantAccessTokenLifespan(v string) {
	o.JwtBearerGrantAccessTokenLifespan.Set(&v)
}
// SetJwtBearerGrantAccessTokenLifespanNil sets the value for JwtBearerGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetJwtBearerGrantAccessTokenLifespanNil() {
	o.JwtBearerGrantAccessTokenLifespan.Set(nil)
}

// UnsetJwtBearerGrantAccessTokenLifespan ensures that no value is present for JwtBearerGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetJwtBearerGrantAccessTokenLifespan() {
	o.JwtBearerGrantAccessTokenLifespan.Unset()
}

// GetPasswordGrantAccessTokenLifespan returns the PasswordGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetPasswordGrantAccessTokenLifespan() string {
	if o == nil || o.PasswordGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordGrantAccessTokenLifespan.Get()
}

// GetPasswordGrantAccessTokenLifespanOk returns a tuple with the PasswordGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetPasswordGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordGrantAccessTokenLifespan.Get(), o.PasswordGrantAccessTokenLifespan.IsSet()
}

// HasPasswordGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasPasswordGrantAccessTokenLifespan() bool {
	if o != nil && o.PasswordGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetPasswordGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the PasswordGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetPasswordGrantAccessTokenLifespan(v string) {
	o.PasswordGrantAccessTokenLifespan.Set(&v)
}
// SetPasswordGrantAccessTokenLifespanNil sets the value for PasswordGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetPasswordGrantAccessTokenLifespanNil() {
	o.PasswordGrantAccessTokenLifespan.Set(nil)
}

// UnsetPasswordGrantAccessTokenLifespan ensures that no value is present for PasswordGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetPasswordGrantAccessTokenLifespan() {
	o.PasswordGrantAccessTokenLifespan.Unset()
}

// GetPasswordGrantRefreshTokenLifespan returns the PasswordGrantRefreshTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetPasswordGrantRefreshTokenLifespan() string {
	if o == nil || o.PasswordGrantRefreshTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.PasswordGrantRefreshTokenLifespan.Get()
}

// GetPasswordGrantRefreshTokenLifespanOk returns a tuple with the PasswordGrantRefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetPasswordGrantRefreshTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordGrantRefreshTokenLifespan.Get(), o.PasswordGrantRefreshTokenLifespan.IsSet()
}

// HasPasswordGrantRefreshTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasPasswordGrantRefreshTokenLifespan() bool {
	if o != nil && o.PasswordGrantRefreshTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetPasswordGrantRefreshTokenLifespan gets a reference to the given NullableString and assigns it to the PasswordGrantRefreshTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetPasswordGrantRefreshTokenLifespan(v string) {
	o.PasswordGrantRefreshTokenLifespan.Set(&v)
}
// SetPasswordGrantRefreshTokenLifespanNil sets the value for PasswordGrantRefreshTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetPasswordGrantRefreshTokenLifespanNil() {
	o.PasswordGrantRefreshTokenLifespan.Set(nil)
}

// UnsetPasswordGrantRefreshTokenLifespan ensures that no value is present for PasswordGrantRefreshTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetPasswordGrantRefreshTokenLifespan() {
	o.PasswordGrantRefreshTokenLifespan.Unset()
}

// GetRefreshTokenGrantAccessTokenLifespan returns the RefreshTokenGrantAccessTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantAccessTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantAccessTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantAccessTokenLifespan.Get()
}

// GetRefreshTokenGrantAccessTokenLifespanOk returns a tuple with the RefreshTokenGrantAccessTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantAccessTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantAccessTokenLifespan.Get(), o.RefreshTokenGrantAccessTokenLifespan.IsSet()
}

// HasRefreshTokenGrantAccessTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasRefreshTokenGrantAccessTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantAccessTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantAccessTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantAccessTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantAccessTokenLifespan(v string) {
	o.RefreshTokenGrantAccessTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantAccessTokenLifespanNil sets the value for RefreshTokenGrantAccessTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantAccessTokenLifespanNil() {
	o.RefreshTokenGrantAccessTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantAccessTokenLifespan ensures that no value is present for RefreshTokenGrantAccessTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetRefreshTokenGrantAccessTokenLifespan() {
	o.RefreshTokenGrantAccessTokenLifespan.Unset()
}

// GetRefreshTokenGrantIdTokenLifespan returns the RefreshTokenGrantIdTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantIdTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantIdTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantIdTokenLifespan.Get()
}

// GetRefreshTokenGrantIdTokenLifespanOk returns a tuple with the RefreshTokenGrantIdTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantIdTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantIdTokenLifespan.Get(), o.RefreshTokenGrantIdTokenLifespan.IsSet()
}

// HasRefreshTokenGrantIdTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasRefreshTokenGrantIdTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantIdTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantIdTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantIdTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantIdTokenLifespan(v string) {
	o.RefreshTokenGrantIdTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantIdTokenLifespanNil sets the value for RefreshTokenGrantIdTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantIdTokenLifespanNil() {
	o.RefreshTokenGrantIdTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantIdTokenLifespan ensures that no value is present for RefreshTokenGrantIdTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetRefreshTokenGrantIdTokenLifespan() {
	o.RefreshTokenGrantIdTokenLifespan.Unset()
}

// GetRefreshTokenGrantRefreshTokenLifespan returns the RefreshTokenGrantRefreshTokenLifespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantRefreshTokenLifespan() string {
	if o == nil || o.RefreshTokenGrantRefreshTokenLifespan.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefreshTokenGrantRefreshTokenLifespan.Get()
}

// GetRefreshTokenGrantRefreshTokenLifespanOk returns a tuple with the RefreshTokenGrantRefreshTokenLifespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOAuth2ClientLifespans) GetRefreshTokenGrantRefreshTokenLifespanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshTokenGrantRefreshTokenLifespan.Get(), o.RefreshTokenGrantRefreshTokenLifespan.IsSet()
}

// HasRefreshTokenGrantRefreshTokenLifespan returns a boolean if a field has been set.
func (o *UpdateOAuth2ClientLifespans) HasRefreshTokenGrantRefreshTokenLifespan() bool {
	if o != nil && o.RefreshTokenGrantRefreshTokenLifespan.IsSet() {
		return true
	}

	return false
}

// SetRefreshTokenGrantRefreshTokenLifespan gets a reference to the given NullableString and assigns it to the RefreshTokenGrantRefreshTokenLifespan field.
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantRefreshTokenLifespan(v string) {
	o.RefreshTokenGrantRefreshTokenLifespan.Set(&v)
}
// SetRefreshTokenGrantRefreshTokenLifespanNil sets the value for RefreshTokenGrantRefreshTokenLifespan to be an explicit nil
func (o *UpdateOAuth2ClientLifespans) SetRefreshTokenGrantRefreshTokenLifespanNil() {
	o.RefreshTokenGrantRefreshTokenLifespan.Set(nil)
}

// UnsetRefreshTokenGrantRefreshTokenLifespan ensures that no value is present for RefreshTokenGrantRefreshTokenLifespan, not even an explicit nil
func (o *UpdateOAuth2ClientLifespans) UnsetRefreshTokenGrantRefreshTokenLifespan() {
	o.RefreshTokenGrantRefreshTokenLifespan.Unset()
}

func (o UpdateOAuth2ClientLifespans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationCodeGrantAccessTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_access_token_lifespan"] = o.AuthorizationCodeGrantAccessTokenLifespan.Get()
	}
	if o.AuthorizationCodeGrantIdTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_id_token_lifespan"] = o.AuthorizationCodeGrantIdTokenLifespan.Get()
	}
	if o.AuthorizationCodeGrantRefreshTokenLifespan.IsSet() {
		toSerialize["authorization_code_grant_refresh_token_lifespan"] = o.AuthorizationCodeGrantRefreshTokenLifespan.Get()
	}
	if o.ClientCredentialsGrantAccessTokenLifespan.IsSet() {
		toSerialize["client_credentials_grant_access_token_lifespan"] = o.ClientCredentialsGrantAccessTokenLifespan.Get()
	}
	if o.ImplicitGrantAccessTokenLifespan.IsSet() {
		toSerialize["implicit_grant_access_token_lifespan"] = o.ImplicitGrantAccessTokenLifespan.Get()
	}
	if o.ImplicitGrantIdTokenLifespan.IsSet() {
		toSerialize["implicit_grant_id_token_lifespan"] = o.ImplicitGrantIdTokenLifespan.Get()
	}
	if o.JwtBearerGrantAccessTokenLifespan.IsSet() {
		toSerialize["jwt_bearer_grant_access_token_lifespan"] = o.JwtBearerGrantAccessTokenLifespan.Get()
	}
	if o.PasswordGrantAccessTokenLifespan.IsSet() {
		toSerialize["password_grant_access_token_lifespan"] = o.PasswordGrantAccessTokenLifespan.Get()
	}
	if o.PasswordGrantRefreshTokenLifespan.IsSet() {
		toSerialize["password_grant_refresh_token_lifespan"] = o.PasswordGrantRefreshTokenLifespan.Get()
	}
	if o.RefreshTokenGrantAccessTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_access_token_lifespan"] = o.RefreshTokenGrantAccessTokenLifespan.Get()
	}
	if o.RefreshTokenGrantIdTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_id_token_lifespan"] = o.RefreshTokenGrantIdTokenLifespan.Get()
	}
	if o.RefreshTokenGrantRefreshTokenLifespan.IsSet() {
		toSerialize["refresh_token_grant_refresh_token_lifespan"] = o.RefreshTokenGrantRefreshTokenLifespan.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOAuth2ClientLifespans struct {
	value *UpdateOAuth2ClientLifespans
	isSet bool
}

func (v NullableUpdateOAuth2ClientLifespans) Get() *UpdateOAuth2ClientLifespans {
	return v.value
}

func (v *NullableUpdateOAuth2ClientLifespans) Set(val *UpdateOAuth2ClientLifespans) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOAuth2ClientLifespans) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOAuth2ClientLifespans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOAuth2ClientLifespans(val *UpdateOAuth2ClientLifespans) *NullableUpdateOAuth2ClientLifespans {
	return &NullableUpdateOAuth2ClientLifespans{value: val, isSet: true}
}

func (v NullableUpdateOAuth2ClientLifespans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOAuth2ClientLifespans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


