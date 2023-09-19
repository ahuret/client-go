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
)

// CORS struct for CORS
type CORS struct {
	// Whether CORS is enabled for this endpoint.
	Enabled bool `json:"enabled"`
	// The allowed origins. Use `*` to allow all origins. A wildcard can also be used in the subdomain, i.e. `https://_*.example.com` will allow all origins on all subdomains of `example.com`.
	Origins []string `json:"origins"`
	AdditionalProperties map[string]interface{}
}

type _CORS CORS

// NewCORS instantiates a new CORS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCORS(enabled bool, origins []string) *CORS {
	this := CORS{}
	this.Enabled = enabled
	this.Origins = origins
	return &this
}

// NewCORSWithDefaults instantiates a new CORS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCORSWithDefaults() *CORS {
	this := CORS{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *CORS) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CORS) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CORS) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOrigins returns the Origins field value
func (o *CORS) GetOrigins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value
// and a boolean to check if the value has been set.
func (o *CORS) GetOriginsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origins, true
}

// SetOrigins sets field value
func (o *CORS) SetOrigins(v []string) {
	o.Origins = v
}

func (o CORS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["origins"] = o.Origins
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CORS) UnmarshalJSON(bytes []byte) (err error) {
	varCORS := _CORS{}

	if err = json.Unmarshal(bytes, &varCORS); err == nil {
		*o = CORS(varCORS)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "origins")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCORS struct {
	value *CORS
	isSet bool
}

func (v NullableCORS) Get() *CORS {
	return v.value
}

func (v *NullableCORS) Set(val *CORS) {
	v.value = val
	v.isSet = true
}

func (v NullableCORS) IsSet() bool {
	return v.isSet
}

func (v *NullableCORS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCORS(val *CORS) *NullableCORS {
	return &NullableCORS{value: val, isSet: true}
}

func (v NullableCORS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCORS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


