/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.33
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OAuth2ConsentRequestOpenIDConnectContext struct for OAuth2ConsentRequestOpenIDConnectContext
type OAuth2ConsentRequestOpenIDConnectContext struct {
	// ACRValues is the Authentication AuthorizationContext Class Reference requested in the OAuth 2.0 Authorization request. It is a parameter defined by OpenID Connect and expresses which level of authentication (e.g. 2FA) is required.  OpenID Connect defines it as follows: > Requested Authentication AuthorizationContext Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. The Authentication AuthorizationContext Class satisfied by the authentication performed is returned as the acr Claim Value, as specified in Section 2. The acr Claim is requested as a Voluntary Claim by this parameter.
	AcrValues []string `json:"acr_values,omitempty"`
	// Display is a string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. The defined values are: page: The Authorization Server SHOULD display the authentication and consent UI consistent with a full User Agent page view. If the display parameter is not specified, this is the default display mode. popup: The Authorization Server SHOULD display the authentication and consent UI consistent with a popup User Agent window. The popup User Agent window should be of an appropriate size for a login-focused dialog and should not obscure the entire window that it is popping up over. touch: The Authorization Server SHOULD display the authentication and consent UI consistent with a device that leverages a touch interface. wap: The Authorization Server SHOULD display the authentication and consent UI consistent with a \"feature phone\" type display.  The Authorization Server MAY also attempt to detect the capabilities of the User Agent and present an appropriate display.
	Display *string `json:"display,omitempty"`
	// IDTokenHintClaims are the claims of the ID Token previously issued by the Authorization Server being passed as a hint about the End-User's current or past authenticated session with the Client.
	IdTokenHintClaims map[string]interface{} `json:"id_token_hint_claims,omitempty"`
	// LoginHint hints about the login identifier the End-User might use to log in (if necessary). This hint can be used by an RP if it first asks the End-User for their e-mail address (or other identifier) and then wants to pass that value as a hint to the discovered authorization service. This value MAY also be a phone number in the format specified for the phone_number Claim. The use of this parameter is optional.
	LoginHint *string `json:"login_hint,omitempty"`
	// UILocales is the End-User'id preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. For instance, the value \"fr-CA fr en\" represents a preference for French as spoken in Canada, then French (without a region designation), followed by English (without a region designation). An error SHOULD NOT result if some or all of the requested locales are not supported by the OpenID Provider.
	UiLocales []string `json:"ui_locales,omitempty"`
}

// NewOAuth2ConsentRequestOpenIDConnectContext instantiates a new OAuth2ConsentRequestOpenIDConnectContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ConsentRequestOpenIDConnectContext() *OAuth2ConsentRequestOpenIDConnectContext {
	this := OAuth2ConsentRequestOpenIDConnectContext{}
	return &this
}

// NewOAuth2ConsentRequestOpenIDConnectContextWithDefaults instantiates a new OAuth2ConsentRequestOpenIDConnectContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConsentRequestOpenIDConnectContextWithDefaults() *OAuth2ConsentRequestOpenIDConnectContext {
	this := OAuth2ConsentRequestOpenIDConnectContext{}
	return &this
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetAcrValues() []string {
	if o == nil || o.AcrValues == nil {
		var ret []string
		return ret
	}
	return o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetAcrValuesOk() ([]string, bool) {
	if o == nil || o.AcrValues == nil {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) HasAcrValues() bool {
	if o != nil && o.AcrValues != nil {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given []string and assigns it to the AcrValues field.
func (o *OAuth2ConsentRequestOpenIDConnectContext) SetAcrValues(v []string) {
	o.AcrValues = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *OAuth2ConsentRequestOpenIDConnectContext) SetDisplay(v string) {
	o.Display = &v
}

// GetIdTokenHintClaims returns the IdTokenHintClaims field value if set, zero value otherwise.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetIdTokenHintClaims() map[string]interface{} {
	if o == nil || o.IdTokenHintClaims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IdTokenHintClaims
}

// GetIdTokenHintClaimsOk returns a tuple with the IdTokenHintClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetIdTokenHintClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.IdTokenHintClaims == nil {
		return nil, false
	}
	return o.IdTokenHintClaims, true
}

// HasIdTokenHintClaims returns a boolean if a field has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) HasIdTokenHintClaims() bool {
	if o != nil && o.IdTokenHintClaims != nil {
		return true
	}

	return false
}

// SetIdTokenHintClaims gets a reference to the given map[string]interface{} and assigns it to the IdTokenHintClaims field.
func (o *OAuth2ConsentRequestOpenIDConnectContext) SetIdTokenHintClaims(v map[string]interface{}) {
	o.IdTokenHintClaims = v
}

// GetLoginHint returns the LoginHint field value if set, zero value otherwise.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetLoginHint() string {
	if o == nil || o.LoginHint == nil {
		var ret string
		return ret
	}
	return *o.LoginHint
}

// GetLoginHintOk returns a tuple with the LoginHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetLoginHintOk() (*string, bool) {
	if o == nil || o.LoginHint == nil {
		return nil, false
	}
	return o.LoginHint, true
}

// HasLoginHint returns a boolean if a field has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) HasLoginHint() bool {
	if o != nil && o.LoginHint != nil {
		return true
	}

	return false
}

// SetLoginHint gets a reference to the given string and assigns it to the LoginHint field.
func (o *OAuth2ConsentRequestOpenIDConnectContext) SetLoginHint(v string) {
	o.LoginHint = &v
}

// GetUiLocales returns the UiLocales field value if set, zero value otherwise.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetUiLocales() []string {
	if o == nil || o.UiLocales == nil {
		var ret []string
		return ret
	}
	return o.UiLocales
}

// GetUiLocalesOk returns a tuple with the UiLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) GetUiLocalesOk() ([]string, bool) {
	if o == nil || o.UiLocales == nil {
		return nil, false
	}
	return o.UiLocales, true
}

// HasUiLocales returns a boolean if a field has been set.
func (o *OAuth2ConsentRequestOpenIDConnectContext) HasUiLocales() bool {
	if o != nil && o.UiLocales != nil {
		return true
	}

	return false
}

// SetUiLocales gets a reference to the given []string and assigns it to the UiLocales field.
func (o *OAuth2ConsentRequestOpenIDConnectContext) SetUiLocales(v []string) {
	o.UiLocales = v
}

func (o OAuth2ConsentRequestOpenIDConnectContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcrValues != nil {
		toSerialize["acr_values"] = o.AcrValues
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.IdTokenHintClaims != nil {
		toSerialize["id_token_hint_claims"] = o.IdTokenHintClaims
	}
	if o.LoginHint != nil {
		toSerialize["login_hint"] = o.LoginHint
	}
	if o.UiLocales != nil {
		toSerialize["ui_locales"] = o.UiLocales
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2ConsentRequestOpenIDConnectContext struct {
	value *OAuth2ConsentRequestOpenIDConnectContext
	isSet bool
}

func (v NullableOAuth2ConsentRequestOpenIDConnectContext) Get() *OAuth2ConsentRequestOpenIDConnectContext {
	return v.value
}

func (v *NullableOAuth2ConsentRequestOpenIDConnectContext) Set(val *OAuth2ConsentRequestOpenIDConnectContext) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ConsentRequestOpenIDConnectContext) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ConsentRequestOpenIDConnectContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ConsentRequestOpenIDConnectContext(val *OAuth2ConsentRequestOpenIDConnectContext) *NullableOAuth2ConsentRequestOpenIDConnectContext {
	return &NullableOAuth2ConsentRequestOpenIDConnectContext{value: val, isSet: true}
}

func (v NullableOAuth2ConsentRequestOpenIDConnectContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ConsentRequestOpenIDConnectContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


