/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.1
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the NeedsPrivilegedSessionError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeedsPrivilegedSessionError{}

// NeedsPrivilegedSessionError struct for NeedsPrivilegedSessionError
type NeedsPrivilegedSessionError struct {
	Error *GenericError `json:"error,omitempty"`
	// Points to where to redirect the user to next.
	RedirectBrowserTo string `json:"redirect_browser_to"`
	AdditionalProperties map[string]interface{}
}

type _NeedsPrivilegedSessionError NeedsPrivilegedSessionError

// NewNeedsPrivilegedSessionError instantiates a new NeedsPrivilegedSessionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeedsPrivilegedSessionError(redirectBrowserTo string) *NeedsPrivilegedSessionError {
	this := NeedsPrivilegedSessionError{}
	this.RedirectBrowserTo = redirectBrowserTo
	return &this
}

// NewNeedsPrivilegedSessionErrorWithDefaults instantiates a new NeedsPrivilegedSessionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeedsPrivilegedSessionErrorWithDefaults() *NeedsPrivilegedSessionError {
	this := NeedsPrivilegedSessionError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *NeedsPrivilegedSessionError) GetError() GenericError {
	if o == nil || IsNil(o.Error) {
		var ret GenericError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeedsPrivilegedSessionError) GetErrorOk() (*GenericError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *NeedsPrivilegedSessionError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given GenericError and assigns it to the Error field.
func (o *NeedsPrivilegedSessionError) SetError(v GenericError) {
	o.Error = &v
}

// GetRedirectBrowserTo returns the RedirectBrowserTo field value
func (o *NeedsPrivilegedSessionError) GetRedirectBrowserTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectBrowserTo
}

// GetRedirectBrowserToOk returns a tuple with the RedirectBrowserTo field value
// and a boolean to check if the value has been set.
func (o *NeedsPrivilegedSessionError) GetRedirectBrowserToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectBrowserTo, true
}

// SetRedirectBrowserTo sets field value
func (o *NeedsPrivilegedSessionError) SetRedirectBrowserTo(v string) {
	o.RedirectBrowserTo = v
}

func (o NeedsPrivilegedSessionError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeedsPrivilegedSessionError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["redirect_browser_to"] = o.RedirectBrowserTo

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NeedsPrivilegedSessionError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"redirect_browser_to",
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

	varNeedsPrivilegedSessionError := _NeedsPrivilegedSessionError{}

	err = json.Unmarshal(data, &varNeedsPrivilegedSessionError)

	if err != nil {
		return err
	}

	*o = NeedsPrivilegedSessionError(varNeedsPrivilegedSessionError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "redirect_browser_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNeedsPrivilegedSessionError struct {
	value *NeedsPrivilegedSessionError
	isSet bool
}

func (v NullableNeedsPrivilegedSessionError) Get() *NeedsPrivilegedSessionError {
	return v.value
}

func (v *NullableNeedsPrivilegedSessionError) Set(val *NeedsPrivilegedSessionError) {
	v.value = val
	v.isSet = true
}

func (v NullableNeedsPrivilegedSessionError) IsSet() bool {
	return v.isSet
}

func (v *NullableNeedsPrivilegedSessionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeedsPrivilegedSessionError(val *NeedsPrivilegedSessionError) *NullableNeedsPrivilegedSessionError {
	return &NullableNeedsPrivilegedSessionError{value: val, isSet: true}
}

func (v NullableNeedsPrivilegedSessionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeedsPrivilegedSessionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


