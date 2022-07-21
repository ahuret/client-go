/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.1.0-alpha.12
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateCustomHostnameBody struct for UpdateCustomHostnameBody
type UpdateCustomHostnameBody struct {
	// The domain where cookies will be set. Has to be a parent domain of the custom hostname to work.
	CookieDomain *string `json:"cookie_domain,omitempty"`
	// CORS Allowed origins for the custom hostname.
	CorsAllowedOrigins []string `json:"cors_allowed_origins,omitempty"`
	// CORS Enabled for the custom hostname.
	CorsEnabled *bool `json:"cors_enabled,omitempty"`
	// The custom hostname where the API will be exposed.
	Hostname *string `json:"hostname,omitempty"`
}

// NewUpdateCustomHostnameBody instantiates a new UpdateCustomHostnameBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomHostnameBody() *UpdateCustomHostnameBody {
	this := UpdateCustomHostnameBody{}
	return &this
}

// NewUpdateCustomHostnameBodyWithDefaults instantiates a new UpdateCustomHostnameBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomHostnameBodyWithDefaults() *UpdateCustomHostnameBody {
	this := UpdateCustomHostnameBody{}
	return &this
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *UpdateCustomHostnameBody) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomHostnameBody) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *UpdateCustomHostnameBody) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *UpdateCustomHostnameBody) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value if set, zero value otherwise.
func (o *UpdateCustomHostnameBody) GetCorsAllowedOrigins() []string {
	if o == nil || o.CorsAllowedOrigins == nil {
		var ret []string
		return ret
	}
	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomHostnameBody) GetCorsAllowedOriginsOk() ([]string, bool) {
	if o == nil || o.CorsAllowedOrigins == nil {
		return nil, false
	}
	return o.CorsAllowedOrigins, true
}

// HasCorsAllowedOrigins returns a boolean if a field has been set.
func (o *UpdateCustomHostnameBody) HasCorsAllowedOrigins() bool {
	if o != nil && o.CorsAllowedOrigins != nil {
		return true
	}

	return false
}

// SetCorsAllowedOrigins gets a reference to the given []string and assigns it to the CorsAllowedOrigins field.
func (o *UpdateCustomHostnameBody) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

// GetCorsEnabled returns the CorsEnabled field value if set, zero value otherwise.
func (o *UpdateCustomHostnameBody) GetCorsEnabled() bool {
	if o == nil || o.CorsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CorsEnabled
}

// GetCorsEnabledOk returns a tuple with the CorsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomHostnameBody) GetCorsEnabledOk() (*bool, bool) {
	if o == nil || o.CorsEnabled == nil {
		return nil, false
	}
	return o.CorsEnabled, true
}

// HasCorsEnabled returns a boolean if a field has been set.
func (o *UpdateCustomHostnameBody) HasCorsEnabled() bool {
	if o != nil && o.CorsEnabled != nil {
		return true
	}

	return false
}

// SetCorsEnabled gets a reference to the given bool and assigns it to the CorsEnabled field.
func (o *UpdateCustomHostnameBody) SetCorsEnabled(v bool) {
	o.CorsEnabled = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *UpdateCustomHostnameBody) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomHostnameBody) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *UpdateCustomHostnameBody) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *UpdateCustomHostnameBody) SetHostname(v string) {
	o.Hostname = &v
}

func (o UpdateCustomHostnameBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if o.CorsAllowedOrigins != nil {
		toSerialize["cors_allowed_origins"] = o.CorsAllowedOrigins
	}
	if o.CorsEnabled != nil {
		toSerialize["cors_enabled"] = o.CorsEnabled
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCustomHostnameBody struct {
	value *UpdateCustomHostnameBody
	isSet bool
}

func (v NullableUpdateCustomHostnameBody) Get() *UpdateCustomHostnameBody {
	return v.value
}

func (v *NullableUpdateCustomHostnameBody) Set(val *UpdateCustomHostnameBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomHostnameBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomHostnameBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomHostnameBody(val *UpdateCustomHostnameBody) *NullableUpdateCustomHostnameBody {
	return &NullableUpdateCustomHostnameBody{value: val, isSet: true}
}

func (v NullableUpdateCustomHostnameBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomHostnameBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


