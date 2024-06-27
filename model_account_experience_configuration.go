/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.11.12
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AccountExperienceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountExperienceConfiguration{}

// AccountExperienceConfiguration struct for AccountExperienceConfiguration
type AccountExperienceConfiguration struct {
	AccountExperienceThemeStylesheet *string `json:"account_experience_theme_stylesheet,omitempty"`
	KratosSelfserviceFlowsRecoveryEnabled *bool `json:"kratos_selfservice_flows_recovery_enabled,omitempty"`
	KratosSelfserviceFlowsRegistrationEnabled *bool `json:"kratos_selfservice_flows_registration_enabled,omitempty"`
	KratosSelfserviceFlowsVerificationEnabled *bool `json:"kratos_selfservice_flows_verification_enabled,omitempty"`
	OrganizationMap *map[string]string `json:"organization_map,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountExperienceConfiguration AccountExperienceConfiguration

// NewAccountExperienceConfiguration instantiates a new AccountExperienceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountExperienceConfiguration() *AccountExperienceConfiguration {
	this := AccountExperienceConfiguration{}
	return &this
}

// NewAccountExperienceConfigurationWithDefaults instantiates a new AccountExperienceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountExperienceConfigurationWithDefaults() *AccountExperienceConfiguration {
	this := AccountExperienceConfiguration{}
	return &this
}

// GetAccountExperienceThemeStylesheet returns the AccountExperienceThemeStylesheet field value if set, zero value otherwise.
func (o *AccountExperienceConfiguration) GetAccountExperienceThemeStylesheet() string {
	if o == nil || IsNil(o.AccountExperienceThemeStylesheet) {
		var ret string
		return ret
	}
	return *o.AccountExperienceThemeStylesheet
}

// GetAccountExperienceThemeStylesheetOk returns a tuple with the AccountExperienceThemeStylesheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExperienceConfiguration) GetAccountExperienceThemeStylesheetOk() (*string, bool) {
	if o == nil || IsNil(o.AccountExperienceThemeStylesheet) {
		return nil, false
	}
	return o.AccountExperienceThemeStylesheet, true
}

// HasAccountExperienceThemeStylesheet returns a boolean if a field has been set.
func (o *AccountExperienceConfiguration) HasAccountExperienceThemeStylesheet() bool {
	if o != nil && !IsNil(o.AccountExperienceThemeStylesheet) {
		return true
	}

	return false
}

// SetAccountExperienceThemeStylesheet gets a reference to the given string and assigns it to the AccountExperienceThemeStylesheet field.
func (o *AccountExperienceConfiguration) SetAccountExperienceThemeStylesheet(v string) {
	o.AccountExperienceThemeStylesheet = &v
}

// GetKratosSelfserviceFlowsRecoveryEnabled returns the KratosSelfserviceFlowsRecoveryEnabled field value if set, zero value otherwise.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsRecoveryEnabled() bool {
	if o == nil || IsNil(o.KratosSelfserviceFlowsRecoveryEnabled) {
		var ret bool
		return ret
	}
	return *o.KratosSelfserviceFlowsRecoveryEnabled
}

// GetKratosSelfserviceFlowsRecoveryEnabledOk returns a tuple with the KratosSelfserviceFlowsRecoveryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsRecoveryEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KratosSelfserviceFlowsRecoveryEnabled) {
		return nil, false
	}
	return o.KratosSelfserviceFlowsRecoveryEnabled, true
}

// HasKratosSelfserviceFlowsRecoveryEnabled returns a boolean if a field has been set.
func (o *AccountExperienceConfiguration) HasKratosSelfserviceFlowsRecoveryEnabled() bool {
	if o != nil && !IsNil(o.KratosSelfserviceFlowsRecoveryEnabled) {
		return true
	}

	return false
}

// SetKratosSelfserviceFlowsRecoveryEnabled gets a reference to the given bool and assigns it to the KratosSelfserviceFlowsRecoveryEnabled field.
func (o *AccountExperienceConfiguration) SetKratosSelfserviceFlowsRecoveryEnabled(v bool) {
	o.KratosSelfserviceFlowsRecoveryEnabled = &v
}

// GetKratosSelfserviceFlowsRegistrationEnabled returns the KratosSelfserviceFlowsRegistrationEnabled field value if set, zero value otherwise.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsRegistrationEnabled() bool {
	if o == nil || IsNil(o.KratosSelfserviceFlowsRegistrationEnabled) {
		var ret bool
		return ret
	}
	return *o.KratosSelfserviceFlowsRegistrationEnabled
}

// GetKratosSelfserviceFlowsRegistrationEnabledOk returns a tuple with the KratosSelfserviceFlowsRegistrationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsRegistrationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KratosSelfserviceFlowsRegistrationEnabled) {
		return nil, false
	}
	return o.KratosSelfserviceFlowsRegistrationEnabled, true
}

// HasKratosSelfserviceFlowsRegistrationEnabled returns a boolean if a field has been set.
func (o *AccountExperienceConfiguration) HasKratosSelfserviceFlowsRegistrationEnabled() bool {
	if o != nil && !IsNil(o.KratosSelfserviceFlowsRegistrationEnabled) {
		return true
	}

	return false
}

// SetKratosSelfserviceFlowsRegistrationEnabled gets a reference to the given bool and assigns it to the KratosSelfserviceFlowsRegistrationEnabled field.
func (o *AccountExperienceConfiguration) SetKratosSelfserviceFlowsRegistrationEnabled(v bool) {
	o.KratosSelfserviceFlowsRegistrationEnabled = &v
}

// GetKratosSelfserviceFlowsVerificationEnabled returns the KratosSelfserviceFlowsVerificationEnabled field value if set, zero value otherwise.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsVerificationEnabled() bool {
	if o == nil || IsNil(o.KratosSelfserviceFlowsVerificationEnabled) {
		var ret bool
		return ret
	}
	return *o.KratosSelfserviceFlowsVerificationEnabled
}

// GetKratosSelfserviceFlowsVerificationEnabledOk returns a tuple with the KratosSelfserviceFlowsVerificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExperienceConfiguration) GetKratosSelfserviceFlowsVerificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KratosSelfserviceFlowsVerificationEnabled) {
		return nil, false
	}
	return o.KratosSelfserviceFlowsVerificationEnabled, true
}

// HasKratosSelfserviceFlowsVerificationEnabled returns a boolean if a field has been set.
func (o *AccountExperienceConfiguration) HasKratosSelfserviceFlowsVerificationEnabled() bool {
	if o != nil && !IsNil(o.KratosSelfserviceFlowsVerificationEnabled) {
		return true
	}

	return false
}

// SetKratosSelfserviceFlowsVerificationEnabled gets a reference to the given bool and assigns it to the KratosSelfserviceFlowsVerificationEnabled field.
func (o *AccountExperienceConfiguration) SetKratosSelfserviceFlowsVerificationEnabled(v bool) {
	o.KratosSelfserviceFlowsVerificationEnabled = &v
}

// GetOrganizationMap returns the OrganizationMap field value if set, zero value otherwise.
func (o *AccountExperienceConfiguration) GetOrganizationMap() map[string]string {
	if o == nil || IsNil(o.OrganizationMap) {
		var ret map[string]string
		return ret
	}
	return *o.OrganizationMap
}

// GetOrganizationMapOk returns a tuple with the OrganizationMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountExperienceConfiguration) GetOrganizationMapOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.OrganizationMap) {
		return nil, false
	}
	return o.OrganizationMap, true
}

// HasOrganizationMap returns a boolean if a field has been set.
func (o *AccountExperienceConfiguration) HasOrganizationMap() bool {
	if o != nil && !IsNil(o.OrganizationMap) {
		return true
	}

	return false
}

// SetOrganizationMap gets a reference to the given map[string]string and assigns it to the OrganizationMap field.
func (o *AccountExperienceConfiguration) SetOrganizationMap(v map[string]string) {
	o.OrganizationMap = &v
}

func (o AccountExperienceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountExperienceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountExperienceThemeStylesheet) {
		toSerialize["account_experience_theme_stylesheet"] = o.AccountExperienceThemeStylesheet
	}
	if !IsNil(o.KratosSelfserviceFlowsRecoveryEnabled) {
		toSerialize["kratos_selfservice_flows_recovery_enabled"] = o.KratosSelfserviceFlowsRecoveryEnabled
	}
	if !IsNil(o.KratosSelfserviceFlowsRegistrationEnabled) {
		toSerialize["kratos_selfservice_flows_registration_enabled"] = o.KratosSelfserviceFlowsRegistrationEnabled
	}
	if !IsNil(o.KratosSelfserviceFlowsVerificationEnabled) {
		toSerialize["kratos_selfservice_flows_verification_enabled"] = o.KratosSelfserviceFlowsVerificationEnabled
	}
	if !IsNil(o.OrganizationMap) {
		toSerialize["organization_map"] = o.OrganizationMap
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountExperienceConfiguration) UnmarshalJSON(data []byte) (err error) {
	varAccountExperienceConfiguration := _AccountExperienceConfiguration{}

	err = json.Unmarshal(data, &varAccountExperienceConfiguration)

	if err != nil {
		return err
	}

	*o = AccountExperienceConfiguration(varAccountExperienceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_experience_theme_stylesheet")
		delete(additionalProperties, "kratos_selfservice_flows_recovery_enabled")
		delete(additionalProperties, "kratos_selfservice_flows_registration_enabled")
		delete(additionalProperties, "kratos_selfservice_flows_verification_enabled")
		delete(additionalProperties, "organization_map")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountExperienceConfiguration struct {
	value *AccountExperienceConfiguration
	isSet bool
}

func (v NullableAccountExperienceConfiguration) Get() *AccountExperienceConfiguration {
	return v.value
}

func (v *NullableAccountExperienceConfiguration) Set(val *AccountExperienceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountExperienceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountExperienceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountExperienceConfiguration(val *AccountExperienceConfiguration) *NullableAccountExperienceConfiguration {
	return &NullableAccountExperienceConfiguration{value: val, isSet: true}
}

func (v NullableAccountExperienceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountExperienceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


