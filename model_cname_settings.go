/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.160
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// CnameSettings struct for CnameSettings
type CnameSettings struct {
	CookieDomain *string `json:"cookie_domain,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Id *string `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	VerificationErrors []string `json:"verification_errors,omitempty"`
	// CustomHostnameStatus is the enumeration of valid state values in the CustomHostnameSSL
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// NewCnameSettings instantiates a new CnameSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCnameSettings() *CnameSettings {
	this := CnameSettings{}
	return &this
}

// NewCnameSettingsWithDefaults instantiates a new CnameSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCnameSettingsWithDefaults() *CnameSettings {
	this := CnameSettings{}
	return &this
}

// GetCookieDomain returns the CookieDomain field value if set, zero value otherwise.
func (o *CnameSettings) GetCookieDomain() string {
	if o == nil || o.CookieDomain == nil {
		var ret string
		return ret
	}
	return *o.CookieDomain
}

// GetCookieDomainOk returns a tuple with the CookieDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetCookieDomainOk() (*string, bool) {
	if o == nil || o.CookieDomain == nil {
		return nil, false
	}
	return o.CookieDomain, true
}

// HasCookieDomain returns a boolean if a field has been set.
func (o *CnameSettings) HasCookieDomain() bool {
	if o != nil && o.CookieDomain != nil {
		return true
	}

	return false
}

// SetCookieDomain gets a reference to the given string and assigns it to the CookieDomain field.
func (o *CnameSettings) SetCookieDomain(v string) {
	o.CookieDomain = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CnameSettings) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CnameSettings) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CnameSettings) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *CnameSettings) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *CnameSettings) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *CnameSettings) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CnameSettings) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CnameSettings) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CnameSettings) SetId(v string) {
	o.Id = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CnameSettings) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CnameSettings) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CnameSettings) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerificationErrors returns the VerificationErrors field value if set, zero value otherwise.
func (o *CnameSettings) GetVerificationErrors() []string {
	if o == nil || o.VerificationErrors == nil {
		var ret []string
		return ret
	}
	return o.VerificationErrors
}

// GetVerificationErrorsOk returns a tuple with the VerificationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetVerificationErrorsOk() ([]string, bool) {
	if o == nil || o.VerificationErrors == nil {
		return nil, false
	}
	return o.VerificationErrors, true
}

// HasVerificationErrors returns a boolean if a field has been set.
func (o *CnameSettings) HasVerificationErrors() bool {
	if o != nil && o.VerificationErrors != nil {
		return true
	}

	return false
}

// SetVerificationErrors gets a reference to the given []string and assigns it to the VerificationErrors field.
func (o *CnameSettings) SetVerificationErrors(v []string) {
	o.VerificationErrors = v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *CnameSettings) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CnameSettings) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *CnameSettings) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *CnameSettings) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o CnameSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CookieDomain != nil {
		toSerialize["cookie_domain"] = o.CookieDomain
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	return json.Marshal(toSerialize)
}

type NullableCnameSettings struct {
	value *CnameSettings
	isSet bool
}

func (v NullableCnameSettings) Get() *CnameSettings {
	return v.value
}

func (v *NullableCnameSettings) Set(val *CnameSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCnameSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCnameSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnameSettings(val *CnameSettings) *NullableCnameSettings {
	return &NullableCnameSettings{value: val, isSet: true}
}

func (v NullableCnameSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnameSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


