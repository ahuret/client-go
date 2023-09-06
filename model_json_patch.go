/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JsonPatch A JSONPatch document as defined by RFC 6902
type JsonPatch struct {
	// This field is used together with operation \"move\" and uses JSON Pointer notation.  Learn more [about JSON Pointers](https://datatracker.ietf.org/doc/html/rfc6901#section-5).
	From *string `json:"from,omitempty"`
	// The operation to be performed. One of \"add\", \"remove\", \"replace\", \"move\", \"copy\", or \"test\".
	Op string `json:"op"`
	// The path to the target path. Uses JSON pointer notation.  Learn more [about JSON Pointers](https://datatracker.ietf.org/doc/html/rfc6901#section-5).
	Path string `json:"path"`
	// The value to be used within the operations.  Learn more [about JSON Pointers](https://datatracker.ietf.org/doc/html/rfc6901#section-5).
	Value interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonPatch JsonPatch

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

// GetFrom returns the From field value if set, zero value otherwise.
func (o *JsonPatch) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatch) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *JsonPatch) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *JsonPatch) SetFrom(v string) {
	o.From = &v
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
	if o == nil {
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
	if o == nil {
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
	if o == nil {
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
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JsonPatch) UnmarshalJSON(bytes []byte) (err error) {
	varJsonPatch := _JsonPatch{}

	if err = json.Unmarshal(bytes, &varJsonPatch); err == nil {
		*o = JsonPatch(varJsonPatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


