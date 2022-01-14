/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.46
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RedirectionConfig struct for RedirectionConfig
type RedirectionConfig struct {
	Global *RedirectionField `json:"global,omitempty"`
	Login *RedirectionField `json:"login,omitempty"`
	Logout *RedirectionField `json:"logout,omitempty"`
	Recovery *RedirectionField `json:"recovery,omitempty"`
	Registration *RedirectionField `json:"registration,omitempty"`
	Settings *RedirectionField `json:"settings,omitempty"`
	UrlAllowlist []string `json:"url_allowlist,omitempty"`
	Verification *RedirectionField `json:"verification,omitempty"`
}

// NewRedirectionConfig instantiates a new RedirectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectionConfig() *RedirectionConfig {
	this := RedirectionConfig{}
	return &this
}

// NewRedirectionConfigWithDefaults instantiates a new RedirectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectionConfigWithDefaults() *RedirectionConfig {
	this := RedirectionConfig{}
	return &this
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *RedirectionConfig) GetGlobal() RedirectionField {
	if o == nil || o.Global == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetGlobalOk() (*RedirectionField, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *RedirectionConfig) HasGlobal() bool {
	if o != nil && o.Global != nil {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given RedirectionField and assigns it to the Global field.
func (o *RedirectionConfig) SetGlobal(v RedirectionField) {
	o.Global = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *RedirectionConfig) GetLogin() RedirectionField {
	if o == nil || o.Login == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetLoginOk() (*RedirectionField, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *RedirectionConfig) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given RedirectionField and assigns it to the Login field.
func (o *RedirectionConfig) SetLogin(v RedirectionField) {
	o.Login = &v
}

// GetLogout returns the Logout field value if set, zero value otherwise.
func (o *RedirectionConfig) GetLogout() RedirectionField {
	if o == nil || o.Logout == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Logout
}

// GetLogoutOk returns a tuple with the Logout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetLogoutOk() (*RedirectionField, bool) {
	if o == nil || o.Logout == nil {
		return nil, false
	}
	return o.Logout, true
}

// HasLogout returns a boolean if a field has been set.
func (o *RedirectionConfig) HasLogout() bool {
	if o != nil && o.Logout != nil {
		return true
	}

	return false
}

// SetLogout gets a reference to the given RedirectionField and assigns it to the Logout field.
func (o *RedirectionConfig) SetLogout(v RedirectionField) {
	o.Logout = &v
}

// GetRecovery returns the Recovery field value if set, zero value otherwise.
func (o *RedirectionConfig) GetRecovery() RedirectionField {
	if o == nil || o.Recovery == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Recovery
}

// GetRecoveryOk returns a tuple with the Recovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetRecoveryOk() (*RedirectionField, bool) {
	if o == nil || o.Recovery == nil {
		return nil, false
	}
	return o.Recovery, true
}

// HasRecovery returns a boolean if a field has been set.
func (o *RedirectionConfig) HasRecovery() bool {
	if o != nil && o.Recovery != nil {
		return true
	}

	return false
}

// SetRecovery gets a reference to the given RedirectionField and assigns it to the Recovery field.
func (o *RedirectionConfig) SetRecovery(v RedirectionField) {
	o.Recovery = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *RedirectionConfig) GetRegistration() RedirectionField {
	if o == nil || o.Registration == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetRegistrationOk() (*RedirectionField, bool) {
	if o == nil || o.Registration == nil {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *RedirectionConfig) HasRegistration() bool {
	if o != nil && o.Registration != nil {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given RedirectionField and assigns it to the Registration field.
func (o *RedirectionConfig) SetRegistration(v RedirectionField) {
	o.Registration = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *RedirectionConfig) GetSettings() RedirectionField {
	if o == nil || o.Settings == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetSettingsOk() (*RedirectionField, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *RedirectionConfig) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given RedirectionField and assigns it to the Settings field.
func (o *RedirectionConfig) SetSettings(v RedirectionField) {
	o.Settings = &v
}

// GetUrlAllowlist returns the UrlAllowlist field value if set, zero value otherwise.
func (o *RedirectionConfig) GetUrlAllowlist() []string {
	if o == nil || o.UrlAllowlist == nil {
		var ret []string
		return ret
	}
	return o.UrlAllowlist
}

// GetUrlAllowlistOk returns a tuple with the UrlAllowlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetUrlAllowlistOk() ([]string, bool) {
	if o == nil || o.UrlAllowlist == nil {
		return nil, false
	}
	return o.UrlAllowlist, true
}

// HasUrlAllowlist returns a boolean if a field has been set.
func (o *RedirectionConfig) HasUrlAllowlist() bool {
	if o != nil && o.UrlAllowlist != nil {
		return true
	}

	return false
}

// SetUrlAllowlist gets a reference to the given []string and assigns it to the UrlAllowlist field.
func (o *RedirectionConfig) SetUrlAllowlist(v []string) {
	o.UrlAllowlist = v
}

// GetVerification returns the Verification field value if set, zero value otherwise.
func (o *RedirectionConfig) GetVerification() RedirectionField {
	if o == nil || o.Verification == nil {
		var ret RedirectionField
		return ret
	}
	return *o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectionConfig) GetVerificationOk() (*RedirectionField, bool) {
	if o == nil || o.Verification == nil {
		return nil, false
	}
	return o.Verification, true
}

// HasVerification returns a boolean if a field has been set.
func (o *RedirectionConfig) HasVerification() bool {
	if o != nil && o.Verification != nil {
		return true
	}

	return false
}

// SetVerification gets a reference to the given RedirectionField and assigns it to the Verification field.
func (o *RedirectionConfig) SetVerification(v RedirectionField) {
	o.Verification = &v
}

func (o RedirectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Global != nil {
		toSerialize["global"] = o.Global
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Logout != nil {
		toSerialize["logout"] = o.Logout
	}
	if o.Recovery != nil {
		toSerialize["recovery"] = o.Recovery
	}
	if o.Registration != nil {
		toSerialize["registration"] = o.Registration
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.UrlAllowlist != nil {
		toSerialize["url_allowlist"] = o.UrlAllowlist
	}
	if o.Verification != nil {
		toSerialize["verification"] = o.Verification
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectionConfig struct {
	value *RedirectionConfig
	isSet bool
}

func (v NullableRedirectionConfig) Get() *RedirectionConfig {
	return v.value
}

func (v *NullableRedirectionConfig) Set(val *RedirectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectionConfig(val *RedirectionConfig) *NullableRedirectionConfig {
	return &NullableRedirectionConfig{value: val, isSet: true}
}

func (v NullableRedirectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


