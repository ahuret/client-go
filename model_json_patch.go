/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.139
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JsonPatch JSON Patch allows you to target individual keys in a JSON document for updates.  For more examples see: https://jsonpatch.com
type JsonPatch struct {
	// The JSON Patch operation
	Op string `json:"op"`
	// The JSON Pointer to the target key
	Path string `json:"path"`
	// The value to be used. Only available for `add` and `replace` operations.
	Value interface{} `json:"value,omitempty"`
}

// NewJsonPatch instantiates a new JsonPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatch(op string, path string) *JsonPatch {
	this := JsonPatch{}
	this.Op = op
	this.Path = path
	return &this
}

// NewJsonPatchWithDefaults instantiates a new JsonPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchWithDefaults() *JsonPatch {
	this := JsonPatch{}
	return &this
}

// GetOp returns the Op field value
func (o *JsonPatch) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetOpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *JsonPatch) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *JsonPatch) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *JsonPatch) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JsonPatch) GetValue() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JsonPatch) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JsonPatch) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *JsonPatch) SetValue(v interface{}) {
	o.Value = v
}

func (o JsonPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableJsonPatch struct {
	value *JsonPatch
	isSet bool
}

func (v NullableJsonPatch) Get() *JsonPatch {
	return v.value
}

func (v *NullableJsonPatch) Set(val *JsonPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatch(val *JsonPatch) *NullableJsonPatch {
	return &NullableJsonPatch{value: val, isSet: true}
}

func (v NullableJsonPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


