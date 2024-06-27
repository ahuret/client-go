/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the LogoutFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoutFlow{}

// LogoutFlow Logout Flow
type LogoutFlow struct {
	// LogoutToken can be used to perform logout using AJAX.
	LogoutToken string `json:"logout_token"`
	// LogoutURL can be opened in a browser to sign the user out.  format: uri
	LogoutUrl string `json:"logout_url"`
	AdditionalProperties map[string]interface{}
}

type _LogoutFlow LogoutFlow

// NewLogoutFlow instantiates a new LogoutFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoutFlow(logoutToken string, logoutUrl string) *LogoutFlow {
	this := LogoutFlow{}
	this.LogoutToken = logoutToken
	this.LogoutUrl = logoutUrl
	return &this
}

// NewLogoutFlowWithDefaults instantiates a new LogoutFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoutFlowWithDefaults() *LogoutFlow {
	this := LogoutFlow{}
	return &this
}

// GetLogoutToken returns the LogoutToken field value
func (o *LogoutFlow) GetLogoutToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutToken
}

// GetLogoutTokenOk returns a tuple with the LogoutToken field value
// and a boolean to check if the value has been set.
func (o *LogoutFlow) GetLogoutTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoutToken, true
}

// SetLogoutToken sets field value
func (o *LogoutFlow) SetLogoutToken(v string) {
	o.LogoutToken = v
}

// GetLogoutUrl returns the LogoutUrl field value
func (o *LogoutFlow) GetLogoutUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value
// and a boolean to check if the value has been set.
func (o *LogoutFlow) GetLogoutUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoutUrl, true
}

// SetLogoutUrl sets field value
func (o *LogoutFlow) SetLogoutUrl(v string) {
	o.LogoutUrl = v
}

func (o LogoutFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoutFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logout_token"] = o.LogoutToken
	toSerialize["logout_url"] = o.LogoutUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogoutFlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logout_token",
		"logout_url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLogoutFlow := _LogoutFlow{}

	err = json.Unmarshal(data, &varLogoutFlow)

	if err != nil {
		return err
	}

	*o = LogoutFlow(varLogoutFlow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "logout_token")
		delete(additionalProperties, "logout_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogoutFlow struct {
	value *LogoutFlow
	isSet bool
}

func (v NullableLogoutFlow) Get() *LogoutFlow {
	return v.value
}

func (v *NullableLogoutFlow) Set(val *LogoutFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoutFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoutFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoutFlow(val *LogoutFlow) *NullableLogoutFlow {
	return &NullableLogoutFlow{value: val, isSet: true}
}

func (v NullableLogoutFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoutFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


