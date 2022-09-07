/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.32
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// WellKnown It includes links to several endpoints (e.g. /oauth2/token) and exposes information on supported signature algorithms among others.
type WellKnown struct {
	// URL of the OP's OAuth 2.0 Authorization Endpoint.
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// Boolean value specifying whether the OP can pass a sid (session ID) Claim in the Logout Token to identify the RP session with the OP. If supported, the sid Claim is also included in ID Tokens issued by the OP
	BackchannelLogoutSessionSupported *bool `json:"backchannel_logout_session_supported,omitempty"`
	// Boolean value specifying whether the OP supports back-channel logout, with true indicating support.
	BackchannelLogoutSupported *bool `json:"backchannel_logout_supported,omitempty"`
	// Boolean value specifying whether the OP supports use of the claims parameter, with true indicating support.
	ClaimsParameterSupported *bool `json:"claims_parameter_supported,omitempty"`
	// JSON array containing a list of the Claim Names of the Claims that the OpenID Provider MAY be able to supply values for. Note that for privacy or other reasons, this might not be an exhaustive list.
	ClaimsSupported []string `json:"claims_supported,omitempty"`
	// JSON array containing a list of Proof Key for Code Exchange (PKCE) [RFC7636] code challenge methods supported by this authorization server.
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,omitempty"`
	// URL at the OP to which an RP can perform a redirect to request that the End-User be logged out at the OP.
	EndSessionEndpoint *string `json:"end_session_endpoint,omitempty"`
	// Boolean value specifying whether the OP can pass iss (issuer) and sid (session ID) query parameters to identify the RP session with the OP when the frontchannel_logout_uri is used. If supported, the sid Claim is also included in ID Tokens issued by the OP.
	FrontchannelLogoutSessionSupported *bool `json:"frontchannel_logout_session_supported,omitempty"`
	// Boolean value specifying whether the OP supports HTTP-based logout, with true indicating support.
	FrontchannelLogoutSupported *bool `json:"frontchannel_logout_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 Grant Type values that this OP supports.
	GrantTypesSupported []string `json:"grant_types_supported,omitempty"`
	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the OP for the ID Token to encode the Claims in a JWT.
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	// URL using the https scheme with no query or fragment component that the OP asserts as its IssuerURL Identifier. If IssuerURL discovery is supported , this value MUST be identical to the issuer value returned by WebFinger. This also MUST be identical to the iss Claim value in ID Tokens issued from this IssuerURL.
	Issuer string `json:"issuer"`
	// URL of the OP's JSON Web Key Set [JWK] document. This contains the signing key(s) the RP uses to validate signatures from the OP. The JWK Set MAY also contain the Server's encryption key(s), which are used by RPs to encrypt requests to the Server. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.
	JwksUri string `json:"jwks_uri"`
	// URL of the OP's Dynamic Client Registration Endpoint.
	RegistrationEndpoint *string `json:"registration_endpoint,omitempty"`
	// JSON array containing a list of the JWS signing algorithms (alg values) supported by the OP for Request Objects, which are described in Section 6.1 of OpenID Connect Core 1.0 [OpenID.Core]. These algorithms are used both when the Request Object is passed by value (using the request parameter) and when it is passed by reference (using the request_uri parameter).
	RequestObjectSigningAlgValuesSupported []string `json:"request_object_signing_alg_values_supported,omitempty"`
	// Boolean value specifying whether the OP supports use of the request parameter, with true indicating support.
	RequestParameterSupported *bool `json:"request_parameter_supported,omitempty"`
	// Boolean value specifying whether the OP supports use of the request_uri parameter, with true indicating support.
	RequestUriParameterSupported *bool `json:"request_uri_parameter_supported,omitempty"`
	// Boolean value specifying whether the OP requires any request_uri values used to be pre-registered using the request_uris registration parameter.
	RequireRequestUriRegistration *bool `json:"require_request_uri_registration,omitempty"`
	// JSON array containing a list of the OAuth 2.0 response_mode values that this OP supports.
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 response_type values that this OP supports. Dynamic OpenID Providers MUST support the code, id_token, and the token id_token Response Type values.
	ResponseTypesSupported []string `json:"response_types_supported"`
	// URL of the authorization server's OAuth 2.0 revocation endpoint.
	RevocationEndpoint *string `json:"revocation_endpoint,omitempty"`
	// SON array containing a list of the OAuth 2.0 [RFC6749] scope values that this server supports. The server MUST support the openid scope value. Servers MAY choose not to advertise some supported scope values even when this parameter is used
	ScopesSupported []string `json:"scopes_supported,omitempty"`
	// JSON array containing a list of the Subject Identifier types that this OP supports. Valid types include pairwise and public.
	SubjectTypesSupported []string `json:"subject_types_supported"`
	// URL of the OP's OAuth 2.0 Token Endpoint
	TokenEndpoint string `json:"token_endpoint"`
	// JSON array containing a list of Client Authentication methods supported by this Token Endpoint. The options are client_secret_post, client_secret_basic, client_secret_jwt, and private_key_jwt, as described in Section 9 of OpenID Connect Core 1.0
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	// URL of the OP's UserInfo Endpoint.
	UserinfoEndpoint *string `json:"userinfo_endpoint,omitempty"`
	// JSON array containing a list of the JWS [JWS] signing algorithms (alg values) [JWA] supported by the UserInfo Endpoint to encode the Claims in a JWT [JWT].
	UserinfoSigningAlgValuesSupported []string `json:"userinfo_signing_alg_values_supported,omitempty"`
}

// NewWellKnown instantiates a new WellKnown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnown(authorizationEndpoint string, idTokenSigningAlgValuesSupported []string, issuer string, jwksUri string, responseTypesSupported []string, subjectTypesSupported []string, tokenEndpoint string) *WellKnown {
	this := WellKnown{}
	this.AuthorizationEndpoint = authorizationEndpoint
	this.IdTokenSigningAlgValuesSupported = idTokenSigningAlgValuesSupported
	this.Issuer = issuer
	this.JwksUri = jwksUri
	this.ResponseTypesSupported = responseTypesSupported
	this.SubjectTypesSupported = subjectTypesSupported
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewWellKnownWithDefaults instantiates a new WellKnown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownWithDefaults() *WellKnown {
	this := WellKnown{}
	return &this
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value
func (o *WellKnown) GetAuthorizationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint sets field value
func (o *WellKnown) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = v
}

// GetBackchannelLogoutSessionSupported returns the BackchannelLogoutSessionSupported field value if set, zero value otherwise.
func (o *WellKnown) GetBackchannelLogoutSessionSupported() bool {
	if o == nil || o.BackchannelLogoutSessionSupported == nil {
		var ret bool
		return ret
	}
	return *o.BackchannelLogoutSessionSupported
}

// GetBackchannelLogoutSessionSupportedOk returns a tuple with the BackchannelLogoutSessionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetBackchannelLogoutSessionSupportedOk() (*bool, bool) {
	if o == nil || o.BackchannelLogoutSessionSupported == nil {
		return nil, false
	}
	return o.BackchannelLogoutSessionSupported, true
}

// HasBackchannelLogoutSessionSupported returns a boolean if a field has been set.
func (o *WellKnown) HasBackchannelLogoutSessionSupported() bool {
	if o != nil && o.BackchannelLogoutSessionSupported != nil {
		return true
	}

	return false
}

// SetBackchannelLogoutSessionSupported gets a reference to the given bool and assigns it to the BackchannelLogoutSessionSupported field.
func (o *WellKnown) SetBackchannelLogoutSessionSupported(v bool) {
	o.BackchannelLogoutSessionSupported = &v
}

// GetBackchannelLogoutSupported returns the BackchannelLogoutSupported field value if set, zero value otherwise.
func (o *WellKnown) GetBackchannelLogoutSupported() bool {
	if o == nil || o.BackchannelLogoutSupported == nil {
		var ret bool
		return ret
	}
	return *o.BackchannelLogoutSupported
}

// GetBackchannelLogoutSupportedOk returns a tuple with the BackchannelLogoutSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetBackchannelLogoutSupportedOk() (*bool, bool) {
	if o == nil || o.BackchannelLogoutSupported == nil {
		return nil, false
	}
	return o.BackchannelLogoutSupported, true
}

// HasBackchannelLogoutSupported returns a boolean if a field has been set.
func (o *WellKnown) HasBackchannelLogoutSupported() bool {
	if o != nil && o.BackchannelLogoutSupported != nil {
		return true
	}

	return false
}

// SetBackchannelLogoutSupported gets a reference to the given bool and assigns it to the BackchannelLogoutSupported field.
func (o *WellKnown) SetBackchannelLogoutSupported(v bool) {
	o.BackchannelLogoutSupported = &v
}

// GetClaimsParameterSupported returns the ClaimsParameterSupported field value if set, zero value otherwise.
func (o *WellKnown) GetClaimsParameterSupported() bool {
	if o == nil || o.ClaimsParameterSupported == nil {
		var ret bool
		return ret
	}
	return *o.ClaimsParameterSupported
}

// GetClaimsParameterSupportedOk returns a tuple with the ClaimsParameterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetClaimsParameterSupportedOk() (*bool, bool) {
	if o == nil || o.ClaimsParameterSupported == nil {
		return nil, false
	}
	return o.ClaimsParameterSupported, true
}

// HasClaimsParameterSupported returns a boolean if a field has been set.
func (o *WellKnown) HasClaimsParameterSupported() bool {
	if o != nil && o.ClaimsParameterSupported != nil {
		return true
	}

	return false
}

// SetClaimsParameterSupported gets a reference to the given bool and assigns it to the ClaimsParameterSupported field.
func (o *WellKnown) SetClaimsParameterSupported(v bool) {
	o.ClaimsParameterSupported = &v
}

// GetClaimsSupported returns the ClaimsSupported field value if set, zero value otherwise.
func (o *WellKnown) GetClaimsSupported() []string {
	if o == nil || o.ClaimsSupported == nil {
		var ret []string
		return ret
	}
	return o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetClaimsSupportedOk() ([]string, bool) {
	if o == nil || o.ClaimsSupported == nil {
		return nil, false
	}
	return o.ClaimsSupported, true
}

// HasClaimsSupported returns a boolean if a field has been set.
func (o *WellKnown) HasClaimsSupported() bool {
	if o != nil && o.ClaimsSupported != nil {
		return true
	}

	return false
}

// SetClaimsSupported gets a reference to the given []string and assigns it to the ClaimsSupported field.
func (o *WellKnown) SetClaimsSupported(v []string) {
	o.ClaimsSupported = v
}

// GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field value if set, zero value otherwise.
func (o *WellKnown) GetCodeChallengeMethodsSupported() []string {
	if o == nil || o.CodeChallengeMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.CodeChallengeMethodsSupported
}

// GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetCodeChallengeMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.CodeChallengeMethodsSupported == nil {
		return nil, false
	}
	return o.CodeChallengeMethodsSupported, true
}

// HasCodeChallengeMethodsSupported returns a boolean if a field has been set.
func (o *WellKnown) HasCodeChallengeMethodsSupported() bool {
	if o != nil && o.CodeChallengeMethodsSupported != nil {
		return true
	}

	return false
}

// SetCodeChallengeMethodsSupported gets a reference to the given []string and assigns it to the CodeChallengeMethodsSupported field.
func (o *WellKnown) SetCodeChallengeMethodsSupported(v []string) {
	o.CodeChallengeMethodsSupported = v
}

// GetEndSessionEndpoint returns the EndSessionEndpoint field value if set, zero value otherwise.
func (o *WellKnown) GetEndSessionEndpoint() string {
	if o == nil || o.EndSessionEndpoint == nil {
		var ret string
		return ret
	}
	return *o.EndSessionEndpoint
}

// GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetEndSessionEndpointOk() (*string, bool) {
	if o == nil || o.EndSessionEndpoint == nil {
		return nil, false
	}
	return o.EndSessionEndpoint, true
}

// HasEndSessionEndpoint returns a boolean if a field has been set.
func (o *WellKnown) HasEndSessionEndpoint() bool {
	if o != nil && o.EndSessionEndpoint != nil {
		return true
	}

	return false
}

// SetEndSessionEndpoint gets a reference to the given string and assigns it to the EndSessionEndpoint field.
func (o *WellKnown) SetEndSessionEndpoint(v string) {
	o.EndSessionEndpoint = &v
}

// GetFrontchannelLogoutSessionSupported returns the FrontchannelLogoutSessionSupported field value if set, zero value otherwise.
func (o *WellKnown) GetFrontchannelLogoutSessionSupported() bool {
	if o == nil || o.FrontchannelLogoutSessionSupported == nil {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogoutSessionSupported
}

// GetFrontchannelLogoutSessionSupportedOk returns a tuple with the FrontchannelLogoutSessionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetFrontchannelLogoutSessionSupportedOk() (*bool, bool) {
	if o == nil || o.FrontchannelLogoutSessionSupported == nil {
		return nil, false
	}
	return o.FrontchannelLogoutSessionSupported, true
}

// HasFrontchannelLogoutSessionSupported returns a boolean if a field has been set.
func (o *WellKnown) HasFrontchannelLogoutSessionSupported() bool {
	if o != nil && o.FrontchannelLogoutSessionSupported != nil {
		return true
	}

	return false
}

// SetFrontchannelLogoutSessionSupported gets a reference to the given bool and assigns it to the FrontchannelLogoutSessionSupported field.
func (o *WellKnown) SetFrontchannelLogoutSessionSupported(v bool) {
	o.FrontchannelLogoutSessionSupported = &v
}

// GetFrontchannelLogoutSupported returns the FrontchannelLogoutSupported field value if set, zero value otherwise.
func (o *WellKnown) GetFrontchannelLogoutSupported() bool {
	if o == nil || o.FrontchannelLogoutSupported == nil {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogoutSupported
}

// GetFrontchannelLogoutSupportedOk returns a tuple with the FrontchannelLogoutSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetFrontchannelLogoutSupportedOk() (*bool, bool) {
	if o == nil || o.FrontchannelLogoutSupported == nil {
		return nil, false
	}
	return o.FrontchannelLogoutSupported, true
}

// HasFrontchannelLogoutSupported returns a boolean if a field has been set.
func (o *WellKnown) HasFrontchannelLogoutSupported() bool {
	if o != nil && o.FrontchannelLogoutSupported != nil {
		return true
	}

	return false
}

// SetFrontchannelLogoutSupported gets a reference to the given bool and assigns it to the FrontchannelLogoutSupported field.
func (o *WellKnown) SetFrontchannelLogoutSupported(v bool) {
	o.FrontchannelLogoutSupported = &v
}

// GetGrantTypesSupported returns the GrantTypesSupported field value if set, zero value otherwise.
func (o *WellKnown) GetGrantTypesSupported() []string {
	if o == nil || o.GrantTypesSupported == nil {
		var ret []string
		return ret
	}
	return o.GrantTypesSupported
}

// GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetGrantTypesSupportedOk() ([]string, bool) {
	if o == nil || o.GrantTypesSupported == nil {
		return nil, false
	}
	return o.GrantTypesSupported, true
}

// HasGrantTypesSupported returns a boolean if a field has been set.
func (o *WellKnown) HasGrantTypesSupported() bool {
	if o != nil && o.GrantTypesSupported != nil {
		return true
	}

	return false
}

// SetGrantTypesSupported gets a reference to the given []string and assigns it to the GrantTypesSupported field.
func (o *WellKnown) SetGrantTypesSupported(v []string) {
	o.GrantTypesSupported = v
}

// GetIdTokenSigningAlgValuesSupported returns the IdTokenSigningAlgValuesSupported field value
func (o *WellKnown) GetIdTokenSigningAlgValuesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IdTokenSigningAlgValuesSupported
}

// GetIdTokenSigningAlgValuesSupportedOk returns a tuple with the IdTokenSigningAlgValuesSupported field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetIdTokenSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdTokenSigningAlgValuesSupported, true
}

// SetIdTokenSigningAlgValuesSupported sets field value
func (o *WellKnown) SetIdTokenSigningAlgValuesSupported(v []string) {
	o.IdTokenSigningAlgValuesSupported = v
}

// GetIssuer returns the Issuer field value
func (o *WellKnown) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *WellKnown) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwksUri returns the JwksUri field value
func (o *WellKnown) GetJwksUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetJwksUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksUri, true
}

// SetJwksUri sets field value
func (o *WellKnown) SetJwksUri(v string) {
	o.JwksUri = v
}

// GetRegistrationEndpoint returns the RegistrationEndpoint field value if set, zero value otherwise.
func (o *WellKnown) GetRegistrationEndpoint() string {
	if o == nil || o.RegistrationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.RegistrationEndpoint
}

// GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRegistrationEndpointOk() (*string, bool) {
	if o == nil || o.RegistrationEndpoint == nil {
		return nil, false
	}
	return o.RegistrationEndpoint, true
}

// HasRegistrationEndpoint returns a boolean if a field has been set.
func (o *WellKnown) HasRegistrationEndpoint() bool {
	if o != nil && o.RegistrationEndpoint != nil {
		return true
	}

	return false
}

// SetRegistrationEndpoint gets a reference to the given string and assigns it to the RegistrationEndpoint field.
func (o *WellKnown) SetRegistrationEndpoint(v string) {
	o.RegistrationEndpoint = &v
}

// GetRequestObjectSigningAlgValuesSupported returns the RequestObjectSigningAlgValuesSupported field value if set, zero value otherwise.
func (o *WellKnown) GetRequestObjectSigningAlgValuesSupported() []string {
	if o == nil || o.RequestObjectSigningAlgValuesSupported == nil {
		var ret []string
		return ret
	}
	return o.RequestObjectSigningAlgValuesSupported
}

// GetRequestObjectSigningAlgValuesSupportedOk returns a tuple with the RequestObjectSigningAlgValuesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRequestObjectSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil || o.RequestObjectSigningAlgValuesSupported == nil {
		return nil, false
	}
	return o.RequestObjectSigningAlgValuesSupported, true
}

// HasRequestObjectSigningAlgValuesSupported returns a boolean if a field has been set.
func (o *WellKnown) HasRequestObjectSigningAlgValuesSupported() bool {
	if o != nil && o.RequestObjectSigningAlgValuesSupported != nil {
		return true
	}

	return false
}

// SetRequestObjectSigningAlgValuesSupported gets a reference to the given []string and assigns it to the RequestObjectSigningAlgValuesSupported field.
func (o *WellKnown) SetRequestObjectSigningAlgValuesSupported(v []string) {
	o.RequestObjectSigningAlgValuesSupported = v
}

// GetRequestParameterSupported returns the RequestParameterSupported field value if set, zero value otherwise.
func (o *WellKnown) GetRequestParameterSupported() bool {
	if o == nil || o.RequestParameterSupported == nil {
		var ret bool
		return ret
	}
	return *o.RequestParameterSupported
}

// GetRequestParameterSupportedOk returns a tuple with the RequestParameterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRequestParameterSupportedOk() (*bool, bool) {
	if o == nil || o.RequestParameterSupported == nil {
		return nil, false
	}
	return o.RequestParameterSupported, true
}

// HasRequestParameterSupported returns a boolean if a field has been set.
func (o *WellKnown) HasRequestParameterSupported() bool {
	if o != nil && o.RequestParameterSupported != nil {
		return true
	}

	return false
}

// SetRequestParameterSupported gets a reference to the given bool and assigns it to the RequestParameterSupported field.
func (o *WellKnown) SetRequestParameterSupported(v bool) {
	o.RequestParameterSupported = &v
}

// GetRequestUriParameterSupported returns the RequestUriParameterSupported field value if set, zero value otherwise.
func (o *WellKnown) GetRequestUriParameterSupported() bool {
	if o == nil || o.RequestUriParameterSupported == nil {
		var ret bool
		return ret
	}
	return *o.RequestUriParameterSupported
}

// GetRequestUriParameterSupportedOk returns a tuple with the RequestUriParameterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRequestUriParameterSupportedOk() (*bool, bool) {
	if o == nil || o.RequestUriParameterSupported == nil {
		return nil, false
	}
	return o.RequestUriParameterSupported, true
}

// HasRequestUriParameterSupported returns a boolean if a field has been set.
func (o *WellKnown) HasRequestUriParameterSupported() bool {
	if o != nil && o.RequestUriParameterSupported != nil {
		return true
	}

	return false
}

// SetRequestUriParameterSupported gets a reference to the given bool and assigns it to the RequestUriParameterSupported field.
func (o *WellKnown) SetRequestUriParameterSupported(v bool) {
	o.RequestUriParameterSupported = &v
}

// GetRequireRequestUriRegistration returns the RequireRequestUriRegistration field value if set, zero value otherwise.
func (o *WellKnown) GetRequireRequestUriRegistration() bool {
	if o == nil || o.RequireRequestUriRegistration == nil {
		var ret bool
		return ret
	}
	return *o.RequireRequestUriRegistration
}

// GetRequireRequestUriRegistrationOk returns a tuple with the RequireRequestUriRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRequireRequestUriRegistrationOk() (*bool, bool) {
	if o == nil || o.RequireRequestUriRegistration == nil {
		return nil, false
	}
	return o.RequireRequestUriRegistration, true
}

// HasRequireRequestUriRegistration returns a boolean if a field has been set.
func (o *WellKnown) HasRequireRequestUriRegistration() bool {
	if o != nil && o.RequireRequestUriRegistration != nil {
		return true
	}

	return false
}

// SetRequireRequestUriRegistration gets a reference to the given bool and assigns it to the RequireRequestUriRegistration field.
func (o *WellKnown) SetRequireRequestUriRegistration(v bool) {
	o.RequireRequestUriRegistration = &v
}

// GetResponseModesSupported returns the ResponseModesSupported field value if set, zero value otherwise.
func (o *WellKnown) GetResponseModesSupported() []string {
	if o == nil || o.ResponseModesSupported == nil {
		var ret []string
		return ret
	}
	return o.ResponseModesSupported
}

// GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetResponseModesSupportedOk() ([]string, bool) {
	if o == nil || o.ResponseModesSupported == nil {
		return nil, false
	}
	return o.ResponseModesSupported, true
}

// HasResponseModesSupported returns a boolean if a field has been set.
func (o *WellKnown) HasResponseModesSupported() bool {
	if o != nil && o.ResponseModesSupported != nil {
		return true
	}

	return false
}

// SetResponseModesSupported gets a reference to the given []string and assigns it to the ResponseModesSupported field.
func (o *WellKnown) SetResponseModesSupported(v []string) {
	o.ResponseModesSupported = v
}

// GetResponseTypesSupported returns the ResponseTypesSupported field value
func (o *WellKnown) GetResponseTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResponseTypesSupported
}

// GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetResponseTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseTypesSupported, true
}

// SetResponseTypesSupported sets field value
func (o *WellKnown) SetResponseTypesSupported(v []string) {
	o.ResponseTypesSupported = v
}

// GetRevocationEndpoint returns the RevocationEndpoint field value if set, zero value otherwise.
func (o *WellKnown) GetRevocationEndpoint() string {
	if o == nil || o.RevocationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.RevocationEndpoint
}

// GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetRevocationEndpointOk() (*string, bool) {
	if o == nil || o.RevocationEndpoint == nil {
		return nil, false
	}
	return o.RevocationEndpoint, true
}

// HasRevocationEndpoint returns a boolean if a field has been set.
func (o *WellKnown) HasRevocationEndpoint() bool {
	if o != nil && o.RevocationEndpoint != nil {
		return true
	}

	return false
}

// SetRevocationEndpoint gets a reference to the given string and assigns it to the RevocationEndpoint field.
func (o *WellKnown) SetRevocationEndpoint(v string) {
	o.RevocationEndpoint = &v
}

// GetScopesSupported returns the ScopesSupported field value if set, zero value otherwise.
func (o *WellKnown) GetScopesSupported() []string {
	if o == nil || o.ScopesSupported == nil {
		var ret []string
		return ret
	}
	return o.ScopesSupported
}

// GetScopesSupportedOk returns a tuple with the ScopesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetScopesSupportedOk() ([]string, bool) {
	if o == nil || o.ScopesSupported == nil {
		return nil, false
	}
	return o.ScopesSupported, true
}

// HasScopesSupported returns a boolean if a field has been set.
func (o *WellKnown) HasScopesSupported() bool {
	if o != nil && o.ScopesSupported != nil {
		return true
	}

	return false
}

// SetScopesSupported gets a reference to the given []string and assigns it to the ScopesSupported field.
func (o *WellKnown) SetScopesSupported(v []string) {
	o.ScopesSupported = v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value
func (o *WellKnown) GetSubjectTypesSupported() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetSubjectTypesSupportedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectTypesSupported, true
}

// SetSubjectTypesSupported sets field value
func (o *WellKnown) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *WellKnown) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *WellKnown) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *WellKnown) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field value if set, zero value otherwise.
func (o *WellKnown) GetTokenEndpointAuthMethodsSupported() []string {
	if o == nil || o.TokenEndpointAuthMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.TokenEndpointAuthMethodsSupported
}

// GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetTokenEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.TokenEndpointAuthMethodsSupported == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethodsSupported, true
}

// HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.
func (o *WellKnown) HasTokenEndpointAuthMethodsSupported() bool {
	if o != nil && o.TokenEndpointAuthMethodsSupported != nil {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethodsSupported gets a reference to the given []string and assigns it to the TokenEndpointAuthMethodsSupported field.
func (o *WellKnown) SetTokenEndpointAuthMethodsSupported(v []string) {
	o.TokenEndpointAuthMethodsSupported = v
}

// GetUserinfoEndpoint returns the UserinfoEndpoint field value if set, zero value otherwise.
func (o *WellKnown) GetUserinfoEndpoint() string {
	if o == nil || o.UserinfoEndpoint == nil {
		var ret string
		return ret
	}
	return *o.UserinfoEndpoint
}

// GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetUserinfoEndpointOk() (*string, bool) {
	if o == nil || o.UserinfoEndpoint == nil {
		return nil, false
	}
	return o.UserinfoEndpoint, true
}

// HasUserinfoEndpoint returns a boolean if a field has been set.
func (o *WellKnown) HasUserinfoEndpoint() bool {
	if o != nil && o.UserinfoEndpoint != nil {
		return true
	}

	return false
}

// SetUserinfoEndpoint gets a reference to the given string and assigns it to the UserinfoEndpoint field.
func (o *WellKnown) SetUserinfoEndpoint(v string) {
	o.UserinfoEndpoint = &v
}

// GetUserinfoSigningAlgValuesSupported returns the UserinfoSigningAlgValuesSupported field value if set, zero value otherwise.
func (o *WellKnown) GetUserinfoSigningAlgValuesSupported() []string {
	if o == nil || o.UserinfoSigningAlgValuesSupported == nil {
		var ret []string
		return ret
	}
	return o.UserinfoSigningAlgValuesSupported
}

// GetUserinfoSigningAlgValuesSupportedOk returns a tuple with the UserinfoSigningAlgValuesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnown) GetUserinfoSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil || o.UserinfoSigningAlgValuesSupported == nil {
		return nil, false
	}
	return o.UserinfoSigningAlgValuesSupported, true
}

// HasUserinfoSigningAlgValuesSupported returns a boolean if a field has been set.
func (o *WellKnown) HasUserinfoSigningAlgValuesSupported() bool {
	if o != nil && o.UserinfoSigningAlgValuesSupported != nil {
		return true
	}

	return false
}

// SetUserinfoSigningAlgValuesSupported gets a reference to the given []string and assigns it to the UserinfoSigningAlgValuesSupported field.
func (o *WellKnown) SetUserinfoSigningAlgValuesSupported(v []string) {
	o.UserinfoSigningAlgValuesSupported = v
}

func (o WellKnown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	}
	if o.BackchannelLogoutSessionSupported != nil {
		toSerialize["backchannel_logout_session_supported"] = o.BackchannelLogoutSessionSupported
	}
	if o.BackchannelLogoutSupported != nil {
		toSerialize["backchannel_logout_supported"] = o.BackchannelLogoutSupported
	}
	if o.ClaimsParameterSupported != nil {
		toSerialize["claims_parameter_supported"] = o.ClaimsParameterSupported
	}
	if o.ClaimsSupported != nil {
		toSerialize["claims_supported"] = o.ClaimsSupported
	}
	if o.CodeChallengeMethodsSupported != nil {
		toSerialize["code_challenge_methods_supported"] = o.CodeChallengeMethodsSupported
	}
	if o.EndSessionEndpoint != nil {
		toSerialize["end_session_endpoint"] = o.EndSessionEndpoint
	}
	if o.FrontchannelLogoutSessionSupported != nil {
		toSerialize["frontchannel_logout_session_supported"] = o.FrontchannelLogoutSessionSupported
	}
	if o.FrontchannelLogoutSupported != nil {
		toSerialize["frontchannel_logout_supported"] = o.FrontchannelLogoutSupported
	}
	if o.GrantTypesSupported != nil {
		toSerialize["grant_types_supported"] = o.GrantTypesSupported
	}
	if true {
		toSerialize["id_token_signing_alg_values_supported"] = o.IdTokenSigningAlgValuesSupported
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.RegistrationEndpoint != nil {
		toSerialize["registration_endpoint"] = o.RegistrationEndpoint
	}
	if o.RequestObjectSigningAlgValuesSupported != nil {
		toSerialize["request_object_signing_alg_values_supported"] = o.RequestObjectSigningAlgValuesSupported
	}
	if o.RequestParameterSupported != nil {
		toSerialize["request_parameter_supported"] = o.RequestParameterSupported
	}
	if o.RequestUriParameterSupported != nil {
		toSerialize["request_uri_parameter_supported"] = o.RequestUriParameterSupported
	}
	if o.RequireRequestUriRegistration != nil {
		toSerialize["require_request_uri_registration"] = o.RequireRequestUriRegistration
	}
	if o.ResponseModesSupported != nil {
		toSerialize["response_modes_supported"] = o.ResponseModesSupported
	}
	if true {
		toSerialize["response_types_supported"] = o.ResponseTypesSupported
	}
	if o.RevocationEndpoint != nil {
		toSerialize["revocation_endpoint"] = o.RevocationEndpoint
	}
	if o.ScopesSupported != nil {
		toSerialize["scopes_supported"] = o.ScopesSupported
	}
	if true {
		toSerialize["subject_types_supported"] = o.SubjectTypesSupported
	}
	if true {
		toSerialize["token_endpoint"] = o.TokenEndpoint
	}
	if o.TokenEndpointAuthMethodsSupported != nil {
		toSerialize["token_endpoint_auth_methods_supported"] = o.TokenEndpointAuthMethodsSupported
	}
	if o.UserinfoEndpoint != nil {
		toSerialize["userinfo_endpoint"] = o.UserinfoEndpoint
	}
	if o.UserinfoSigningAlgValuesSupported != nil {
		toSerialize["userinfo_signing_alg_values_supported"] = o.UserinfoSigningAlgValuesSupported
	}
	return json.Marshal(toSerialize)
}

type NullableWellKnown struct {
	value *WellKnown
	isSet bool
}

func (v NullableWellKnown) Get() *WellKnown {
	return v.value
}

func (v *NullableWellKnown) Set(val *WellKnown) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnown) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnown(val *WellKnown) *NullableWellKnown {
	return &NullableWellKnown{value: val, isSet: true}
}

func (v NullableWellKnown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


