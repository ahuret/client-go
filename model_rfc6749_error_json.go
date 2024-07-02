/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.12.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the RFC6749ErrorJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RFC6749ErrorJson{}

// RFC6749ErrorJson struct for RFC6749ErrorJson
type RFC6749ErrorJson struct {
	Error *string `json:"error,omitempty"`
	ErrorDebug *string `json:"error_debug,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
	ErrorHint *string `json:"error_hint,omitempty"`
	StatusCode *int64 `json:"status_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RFC6749ErrorJson RFC6749ErrorJson

// NewRFC6749ErrorJson instantiates a new RFC6749ErrorJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRFC6749ErrorJson() *RFC6749ErrorJson {
	this := RFC6749ErrorJson{}
	return &this
}

// NewRFC6749ErrorJsonWithDefaults instantiates a new RFC6749ErrorJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRFC6749ErrorJsonWithDefaults() *RFC6749ErrorJson {
	this := RFC6749ErrorJson{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RFC6749ErrorJson) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RFC6749ErrorJson) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RFC6749ErrorJson) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RFC6749ErrorJson) SetError(v string) {
	o.Error = &v
}

// GetErrorDebug returns the ErrorDebug field value if set, zero value otherwise.
func (o *RFC6749ErrorJson) GetErrorDebug() string {
	if o == nil || IsNil(o.ErrorDebug) {
		var ret string
		return ret
	}
	return *o.ErrorDebug
}

// GetErrorDebugOk returns a tuple with the ErrorDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RFC6749ErrorJson) GetErrorDebugOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDebug) {
		return nil, false
	}
	return o.ErrorDebug, true
}

// HasErrorDebug returns a boolean if a field has been set.
func (o *RFC6749ErrorJson) HasErrorDebug() bool {
	if o != nil && !IsNil(o.ErrorDebug) {
		return true
	}

	return false
}

// SetErrorDebug gets a reference to the given string and assigns it to the ErrorDebug field.
func (o *RFC6749ErrorJson) SetErrorDebug(v string) {
	o.ErrorDebug = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *RFC6749ErrorJson) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RFC6749ErrorJson) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *RFC6749ErrorJson) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *RFC6749ErrorJson) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorHint returns the ErrorHint field value if set, zero value otherwise.
func (o *RFC6749ErrorJson) GetErrorHint() string {
	if o == nil || IsNil(o.ErrorHint) {
		var ret string
		return ret
	}
	return *o.ErrorHint
}

// GetErrorHintOk returns a tuple with the ErrorHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RFC6749ErrorJson) GetErrorHintOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorHint) {
		return nil, false
	}
	return o.ErrorHint, true
}

// HasErrorHint returns a boolean if a field has been set.
func (o *RFC6749ErrorJson) HasErrorHint() bool {
	if o != nil && !IsNil(o.ErrorHint) {
		return true
	}

	return false
}

// SetErrorHint gets a reference to the given string and assigns it to the ErrorHint field.
func (o *RFC6749ErrorJson) SetErrorHint(v string) {
	o.ErrorHint = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *RFC6749ErrorJson) GetStatusCode() int64 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RFC6749ErrorJson) GetStatusCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *RFC6749ErrorJson) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *RFC6749ErrorJson) SetStatusCode(v int64) {
	o.StatusCode = &v
}

func (o RFC6749ErrorJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RFC6749ErrorJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorDebug) {
		toSerialize["error_debug"] = o.ErrorDebug
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if !IsNil(o.ErrorHint) {
		toSerialize["error_hint"] = o.ErrorHint
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RFC6749ErrorJson) UnmarshalJSON(data []byte) (err error) {
	varRFC6749ErrorJson := _RFC6749ErrorJson{}

	err = json.Unmarshal(data, &varRFC6749ErrorJson)

	if err != nil {
		return err
	}

	*o = RFC6749ErrorJson(varRFC6749ErrorJson)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "error_debug")
		delete(additionalProperties, "error_description")
		delete(additionalProperties, "error_hint")
		delete(additionalProperties, "status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRFC6749ErrorJson struct {
	value *RFC6749ErrorJson
	isSet bool
}

func (v NullableRFC6749ErrorJson) Get() *RFC6749ErrorJson {
	return v.value
}

func (v *NullableRFC6749ErrorJson) Set(val *RFC6749ErrorJson) {
	v.value = val
	v.isSet = true
}

func (v NullableRFC6749ErrorJson) IsSet() bool {
	return v.isSet
}

func (v *NullableRFC6749ErrorJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRFC6749ErrorJson(val *RFC6749ErrorJson) *NullableRFC6749ErrorJson {
	return &NullableRFC6749ErrorJson{value: val, isSet: true}
}

func (v NullableRFC6749ErrorJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRFC6749ErrorJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


