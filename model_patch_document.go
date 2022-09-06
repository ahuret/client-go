/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v0.2.0-alpha.29
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PatchDocument A JSONPatch document as defined by RFC 6902
type PatchDocument struct {
	// A JSON-pointer
	From *string `json:"from,omitempty"`
	// The operation to be performed
	Op string `json:"op"`
	// A JSON-pointer
	Path string `json:"path"`
	// The value to be used within the operations
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewPatchDocument instantiates a new PatchDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDocument(op string, path string) *PatchDocument {
	this := PatchDocument{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchDocumentWithDefaults instantiates a new PatchDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDocumentWithDefaults() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchDocument) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchDocument) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchDocument) SetFrom(v string) {
	o.From = &v
}

// GetOp returns the Op field value
func (o *PatchDocument) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchDocument) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchDocument) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchDocument) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchDocument) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchDocument) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *PatchDocument) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o PatchDocument) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePatchDocument struct {
	value *PatchDocument
	isSet bool
}

func (v NullablePatchDocument) Get() *PatchDocument {
	return v.value
}

func (v *NullablePatchDocument) Set(val *PatchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDocument(val *PatchDocument) *NullablePatchDocument {
	return &NullablePatchDocument{value: val, isSet: true}
}

func (v NullablePatchDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


