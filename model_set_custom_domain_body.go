/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.45
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetCustomDomainBody Update Custom Hostname Body
type SetCustomDomainBody struct {
	// The domain where cookies will be set. Has to be a parent domain of the custom hostname to work.
	CookieDomain *string `json:"cookie_domain,omitempty"`
	// CORS Allowed origins for the custom hostname.
	CorsAllowedOrigins []string `json:"cors_allowed_origins,omitempty"`
	// CORS Enabled for the custom hostname.
	CorsEnabled *bool `json:"cors_enabled,omitempty"`
	// The custom UI base URL where the UI will be exposed.
	CustomUiBaseUrl *string `json:"custom_ui_base_url,omitempty"`
	// The custom hostname where the API will be exposed.
	Hostname *string `json:"hostname,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SetCustomDomainBody SetCustomDomainBody

// NewSetCustomDomainBody instantiates a new SetCustomDomainBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCustomDomainBody() *SetCustomDomainBody {
	this := SetCustomDomainBody{}
	return &this
}

// NewSetCustomDomainBodyWithDefaults instantiates a new SetCustomDomainBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCustomDomainBodyWithDefaults() *SetCustomDomainBody {
	this := SetCustomDomainBody{}
	return &this
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *SetCustomDomainBody) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomDomainBody) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *SetCustomDomainBody) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *SetCustomDomainBody) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value if set, zero value otherwise.
func (o *SetCustomDomainBody) GetCorsAllowedOrigins() []string {
	if o == nil || o.CorsAllowedOrigins == nil {
		var ret []string
		return ret
	}
	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomDomainBody) GetCorsAllowedOriginsOk() ([]string, bool) {
	if o == nil || o.CorsAllowedOrigins == nil {
		return nil, false
	}
	return o.CorsAllowedOrigins, true
}

// HasCorsAllowedOrigins returns a boolean if a field has been set.
func (o *SetCustomDomainBody) HasCorsAllowedOrigins() bool {
	if o != nil && o.CorsAllowedOrigins != nil {
		return true
	}

	return false
}

// SetCorsAllowedOrigins gets a reference to the given []string and assigns it to the CorsAllowedOrigins field.
func (o *SetCustomDomainBody) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

// GetCorsEnabled returns the CorsEnabled field value if set, zero value otherwise.
func (o *SetCustomDomainBody) GetCorsEnabled() bool {
	if o == nil || o.CorsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CorsEnabled
}

// GetCorsEnabledOk returns a tuple with the CorsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomDomainBody) GetCorsEnabledOk() (*bool, bool) {
	if o == nil || o.CorsEnabled == nil {
		return nil, false
	}
	return o.CorsEnabled, true
}

// HasCorsEnabled returns a boolean if a field has been set.
func (o *SetCustomDomainBody) HasCorsEnabled() bool {
	if o != nil && o.CorsEnabled != nil {
		return true
	}

	return false
}

// SetCorsEnabled gets a reference to the given bool and assigns it to the CorsEnabled field.
func (o *SetCustomDomainBody) SetCorsEnabled(v bool) {
	o.CorsEnabled = &v
}

// GetCustomUiBaseUrl returns the CustomUiBaseUrl field value if set, zero value otherwise.
func (o *SetCustomDomainBody) GetCustomUiBaseUrl() string {
	if o == nil || o.CustomUiBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomUiBaseUrl
}

// GetCustomUiBaseUrlOk returns a tuple with the CustomUiBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomDomainBody) GetCustomUiBaseUrlOk() (*string, bool) {
	if o == nil || o.CustomUiBaseUrl == nil {
		return nil, false
	}
	return o.CustomUiBaseUrl, true
}

// HasCustomUiBaseUrl returns a boolean if a field has been set.
func (o *SetCustomDomainBody) HasCustomUiBaseUrl() bool {
	if o != nil && o.CustomUiBaseUrl != nil {
		return true
	}

	return false
}

// SetCustomUiBaseUrl gets a reference to the given string and assigns it to the CustomUiBaseUrl field.
func (o *SetCustomDomainBody) SetCustomUiBaseUrl(v string) {
	o.CustomUiBaseUrl = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SetCustomDomainBody) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomDomainBody) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SetCustomDomainBody) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SetCustomDomainBody) SetHostname(v string) {
	o.Hostname = &v
}

func (o SetCustomDomainBody) MarshalJSON() ([]byte, error) {
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
	if o.CustomUiBaseUrl != nil {
		toSerialize["custom_ui_base_url"] = o.CustomUiBaseUrl
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SetCustomDomainBody) UnmarshalJSON(bytes []byte) (err error) {
	varSetCustomDomainBody := _SetCustomDomainBody{}

	if err = json.Unmarshal(bytes, &varSetCustomDomainBody); err == nil {
		*o = SetCustomDomainBody(varSetCustomDomainBody)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cookie_domain")
		delete(additionalProperties, "cors_allowed_origins")
		delete(additionalProperties, "cors_enabled")
		delete(additionalProperties, "custom_ui_base_url")
		delete(additionalProperties, "hostname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetCustomDomainBody struct {
	value *SetCustomDomainBody
	isSet bool
}

func (v NullableSetCustomDomainBody) Get() *SetCustomDomainBody {
	return v.value
}

func (v *NullableSetCustomDomainBody) Set(val *SetCustomDomainBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCustomDomainBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCustomDomainBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCustomDomainBody(val *SetCustomDomainBody) *NullableSetCustomDomainBody {
	return &NullableSetCustomDomainBody{value: val, isSet: true}
}

func (v NullableSetCustomDomainBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCustomDomainBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


