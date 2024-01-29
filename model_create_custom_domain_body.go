/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.5.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateCustomDomainBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomDomainBody{}

// CreateCustomDomainBody Create Custom Hostname Request Body
type CreateCustomDomainBody struct {
	// The domain where cookies will be set. Has to be a parent domain of the custom hostname to work.
	CookieDomain *string `json:"cookie_domain,omitempty"`
	// CORS Allowed origins for the custom hostname.
	CorsAllowedOrigins []string `json:"cors_allowed_origins,omitempty"`
	// CORS Enabled for the custom hostname.
	CorsEnabled *bool `json:"cors_enabled,omitempty"`
	// The base URL where the custom user interface will be exposed.
	CustomUiBaseUrl *string `json:"custom_ui_base_url,omitempty"`
	// The custom hostname where the API will be exposed.
	Hostname *string `json:"hostname,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateCustomDomainBody CreateCustomDomainBody

// NewCreateCustomDomainBody instantiates a new CreateCustomDomainBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomDomainBody() *CreateCustomDomainBody {
	this := CreateCustomDomainBody{}
	return &this
}

// NewCreateCustomDomainBodyWithDefaults instantiates a new CreateCustomDomainBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomDomainBodyWithDefaults() *CreateCustomDomainBody {
	this := CreateCustomDomainBody{}
	return &this
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *CreateCustomDomainBody) GetCookieDomain() string {
	if o == nil || IsNil(o.CookieDomain) {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainBody) GetCookieDomainOk() (*string, bool) {
	if o == nil || IsNil(o.CookieDomain) {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *CreateCustomDomainBody) HasCookieDomain() bool {
	if o != nil && !IsNil(o.CookieDomain) {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *CreateCustomDomainBody) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value if set, zero value otherwise.
func (o *CreateCustomDomainBody) GetCorsAllowedOrigins() []string {
	if o == nil || IsNil(o.CorsAllowedOrigins) {
		var ret []string
		return ret
	}
	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainBody) GetCorsAllowedOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.CorsAllowedOrigins) {
		return nil, false
	}
	return o.CorsAllowedOrigins, true
}

// HasCorsAllowedOrigins returns a boolean if a field has been set.
func (o *CreateCustomDomainBody) HasCorsAllowedOrigins() bool {
	if o != nil && !IsNil(o.CorsAllowedOrigins) {
		return true
	}

	return false
}

// SetCorsAllowedOrigins gets a reference to the given []string and assigns it to the CorsAllowedOrigins field.
func (o *CreateCustomDomainBody) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

// GetCorsEnabled returns the CorsEnabled field value if set, zero value otherwise.
func (o *CreateCustomDomainBody) GetCorsEnabled() bool {
	if o == nil || IsNil(o.CorsEnabled) {
		var ret bool
		return ret
	}
	return *o.CorsEnabled
}

// GetCorsEnabledOk returns a tuple with the CorsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainBody) GetCorsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CorsEnabled) {
		return nil, false
	}
	return o.CorsEnabled, true
}

// HasCorsEnabled returns a boolean if a field has been set.
func (o *CreateCustomDomainBody) HasCorsEnabled() bool {
	if o != nil && !IsNil(o.CorsEnabled) {
		return true
	}

	return false
}

// SetCorsEnabled gets a reference to the given bool and assigns it to the CorsEnabled field.
func (o *CreateCustomDomainBody) SetCorsEnabled(v bool) {
	o.CorsEnabled = &v
}

// GetCustomUiBaseUrl returns the CustomUiBaseUrl field value if set, zero value otherwise.
func (o *CreateCustomDomainBody) GetCustomUiBaseUrl() string {
	if o == nil || IsNil(o.CustomUiBaseUrl) {
		var ret string
		return ret
	}
	return *o.CustomUiBaseUrl
}

// GetCustomUiBaseUrlOk returns a tuple with the CustomUiBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainBody) GetCustomUiBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomUiBaseUrl) {
		return nil, false
	}
	return o.CustomUiBaseUrl, true
}

// HasCustomUiBaseUrl returns a boolean if a field has been set.
func (o *CreateCustomDomainBody) HasCustomUiBaseUrl() bool {
	if o != nil && !IsNil(o.CustomUiBaseUrl) {
		return true
	}

	return false
}

// SetCustomUiBaseUrl gets a reference to the given string and assigns it to the CustomUiBaseUrl field.
func (o *CreateCustomDomainBody) SetCustomUiBaseUrl(v string) {
	o.CustomUiBaseUrl = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CreateCustomDomainBody) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomDomainBody) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CreateCustomDomainBody) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CreateCustomDomainBody) SetHostname(v string) {
	o.Hostname = &v
}

func (o CreateCustomDomainBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomDomainBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CookieDomain) {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if !IsNil(o.CorsAllowedOrigins) {
		toSerialize["cors_allowed_origins"] = o.CorsAllowedOrigins
	}
	if !IsNil(o.CorsEnabled) {
		toSerialize["cors_enabled"] = o.CorsEnabled
	}
	if !IsNil(o.CustomUiBaseUrl) {
		toSerialize["custom_ui_base_url"] = o.CustomUiBaseUrl
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCustomDomainBody) UnmarshalJSON(bytes []byte) (err error) {
	varCreateCustomDomainBody := _CreateCustomDomainBody{}

	err = json.Unmarshal(bytes, &varCreateCustomDomainBody)

	if err != nil {
		return err
	}

	*o = CreateCustomDomainBody(varCreateCustomDomainBody)

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

type NullableCreateCustomDomainBody struct {
	value *CreateCustomDomainBody
	isSet bool
}

func (v NullableCreateCustomDomainBody) Get() *CreateCustomDomainBody {
	return v.value
}

func (v *NullableCreateCustomDomainBody) Set(val *CreateCustomDomainBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomDomainBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomDomainBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomDomainBody(val *CreateCustomDomainBody) *NullableCreateCustomDomainBody {
	return &NullableCreateCustomDomainBody{value: val, isSet: true}
}

func (v NullableCreateCustomDomainBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomDomainBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


