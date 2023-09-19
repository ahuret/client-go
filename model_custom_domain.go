/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// CustomDomain Custom Hostname
type CustomDomain struct {
	CookieDomain *string `json:"cookie_domain,omitempty"`
	CorsAllowedOrigins []string `json:"cors_allowed_origins,omitempty"`
	CorsEnabled *bool `json:"cors_enabled,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CustomUiBaseUrl *string `json:"custom_ui_base_url,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Id *string `json:"id,omitempty"`
	SslStatus *string `json:"ssl_status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	VerificationErrors []string `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomDomain CustomDomain

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *CustomDomain) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *CustomDomain) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *CustomDomain) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value if set, zero value otherwise.
func (o *CustomDomain) GetCorsAllowedOrigins() []string {
	if o == nil || o.CorsAllowedOrigins == nil {
		var ret []string
		return ret
	}
	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCorsAllowedOriginsOk() ([]string, bool) {
	if o == nil || o.CorsAllowedOrigins == nil {
		return nil, false
	}
	return o.CorsAllowedOrigins, true
}

// HasCorsAllowedOrigins returns a boolean if a field has been set.
func (o *CustomDomain) HasCorsAllowedOrigins() bool {
	if o != nil && o.CorsAllowedOrigins != nil {
		return true
	}

	return false
}

// SetCorsAllowedOrigins gets a reference to the given []string and assigns it to the CorsAllowedOrigins field.
func (o *CustomDomain) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

// GetCorsEnabled returns the CorsEnabled field value if set, zero value otherwise.
func (o *CustomDomain) GetCorsEnabled() bool {
	if o == nil || o.CorsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CorsEnabled
}

// GetCorsEnabledOk returns a tuple with the CorsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCorsEnabledOk() (*bool, bool) {
	if o == nil || o.CorsEnabled == nil {
		return nil, false
	}
	return o.CorsEnabled, true
}

// HasCorsEnabled returns a boolean if a field has been set.
func (o *CustomDomain) HasCorsEnabled() bool {
	if o != nil && o.CorsEnabled != nil {
		return true
	}

	return false
}

// SetCorsEnabled gets a reference to the given bool and assigns it to the CorsEnabled field.
func (o *CustomDomain) SetCorsEnabled(v bool) {
	o.CorsEnabled = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomDomain) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomDomain) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomDomain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomUiBaseUrl returns the CustomUiBaseUrl field value if set, zero value otherwise.
func (o *CustomDomain) GetCustomUiBaseUrl() string {
	if o == nil || o.CustomUiBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomUiBaseUrl
}

// GetCustomUiBaseUrlOk returns a tuple with the CustomUiBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCustomUiBaseUrlOk() (*string, bool) {
	if o == nil || o.CustomUiBaseUrl == nil {
		return nil, false
	}
	return o.CustomUiBaseUrl, true
}

// HasCustomUiBaseUrl returns a boolean if a field has been set.
func (o *CustomDomain) HasCustomUiBaseUrl() bool {
	if o != nil && o.CustomUiBaseUrl != nil {
		return true
	}

	return false
}

// SetCustomUiBaseUrl gets a reference to the given string and assigns it to the CustomUiBaseUrl field.
func (o *CustomDomain) SetCustomUiBaseUrl(v string) {
	o.CustomUiBaseUrl = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CustomDomain) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CustomDomain) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CustomDomain) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomDomain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomDomain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomDomain) SetId(v string) {
	o.Id = &v
}

// GetSslStatus returns the SslStatus field value if set, zero value otherwise.
func (o *CustomDomain) GetSslStatus() string {
	if o == nil || o.SslStatus == nil {
		var ret string
		return ret
	}
	return *o.SslStatus
}

// GetSslStatusOk returns a tuple with the SslStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetSslStatusOk() (*string, bool) {
	if o == nil || o.SslStatus == nil {
		return nil, false
	}
	return o.SslStatus, true
}

// HasSslStatus returns a boolean if a field has been set.
func (o *CustomDomain) HasSslStatus() bool {
	if o != nil && o.SslStatus != nil {
		return true
	}

	return false
}

// SetSslStatus gets a reference to the given string and assigns it to the SslStatus field.
func (o *CustomDomain) SetSslStatus(v string) {
	o.SslStatus = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomDomain) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomDomain) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomDomain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerificationErrors returns the VerificationErrors field value if set, zero value otherwise.
func (o *CustomDomain) GetVerificationErrors() []string {
	if o == nil || o.VerificationErrors == nil {
		var ret []string
		return ret
	}
	return o.VerificationErrors
}

// GetVerificationErrorsOk returns a tuple with the VerificationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetVerificationErrorsOk() ([]string, bool) {
	if o == nil || o.VerificationErrors == nil {
		return nil, false
	}
	return o.VerificationErrors, true
}

// HasVerificationErrors returns a boolean if a field has been set.
func (o *CustomDomain) HasVerificationErrors() bool {
	if o != nil && o.VerificationErrors != nil {
		return true
	}

	return false
}

// SetVerificationErrors gets a reference to the given []string and assigns it to the VerificationErrors field.
func (o *CustomDomain) SetVerificationErrors(v []string) {
	o.VerificationErrors = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *CustomDomain) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *CustomDomain) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *CustomDomain) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CustomUiBaseUrl != nil {
		toSerialize["custom_ui_base_url"] = o.CustomUiBaseUrl
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SslStatus != nil {
		toSerialize["ssl_status"] = o.SslStatus
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.VerificationErrors != nil {
		toSerialize["verification_errors"] = o.VerificationErrors
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CustomDomain) UnmarshalJSON(bytes []byte) (err error) {
	varCustomDomain := _CustomDomain{}

	if err = json.Unmarshal(bytes, &varCustomDomain); err == nil {
		*o = CustomDomain(varCustomDomain)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cookie_domain")
		delete(additionalProperties, "cors_allowed_origins")
		delete(additionalProperties, "cors_enabled")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "custom_ui_base_url")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ssl_status")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "verification_errors")
		delete(additionalProperties, "verification_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


