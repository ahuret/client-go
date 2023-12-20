/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.7
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the UiNodeInputAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiNodeInputAttributes{}

// UiNodeInputAttributes InputAttributes represents the attributes of an input node
type UiNodeInputAttributes struct {
	// The autocomplete attribute for the input. email InputAttributeAutocompleteEmail tel InputAttributeAutocompleteTel url InputAttributeAutocompleteUrl current-password InputAttributeAutocompleteCurrentPassword new-password InputAttributeAutocompleteNewPassword one-time-code InputAttributeAutocompleteOneTimeCode
	Autocomplete *string `json:"autocomplete,omitempty"`
	// Sets the input's disabled field to true or false.
	Disabled bool `json:"disabled"`
	Label *UiText `json:"label,omitempty"`
	// The input's element name.
	Name string `json:"name"`
	// NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0.  In this struct it technically always is \"input\".
	NodeType string `json:"node_type"`
	// OnClick may contain javascript which should be executed on click. This is primarily used for WebAuthn.
	Onclick *string `json:"onclick,omitempty"`
	// The input's pattern.
	Pattern *string `json:"pattern,omitempty"`
	// Mark this input field as required.
	Required *bool `json:"required,omitempty"`
	// The input's element type. text InputAttributeTypeText password InputAttributeTypePassword number InputAttributeTypeNumber checkbox InputAttributeTypeCheckbox hidden InputAttributeTypeHidden email InputAttributeTypeEmail tel InputAttributeTypeTel submit InputAttributeTypeSubmit button InputAttributeTypeButton datetime-local InputAttributeTypeDateTimeLocal date InputAttributeTypeDate url InputAttributeTypeURI
	Type string `json:"type"`
	// The input's value.
	Value interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UiNodeInputAttributes UiNodeInputAttributes

// NewUiNodeInputAttributes instantiates a new UiNodeInputAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeInputAttributes(disabled bool, name string, nodeType string, type_ string) *UiNodeInputAttributes {
	this := UiNodeInputAttributes{}
	this.Disabled = disabled
	this.Name = name
	this.NodeType = nodeType
	this.Type = type_
	return &this
}

// NewUiNodeInputAttributesWithDefaults instantiates a new UiNodeInputAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeInputAttributesWithDefaults() *UiNodeInputAttributes {
	this := UiNodeInputAttributes{}
	return &this
}

// GetAutocomplete returns the Autocomplete field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetAutocomplete() string {
	if o == nil || IsNil(o.Autocomplete) {
		var ret string
		return ret
	}
	return *o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetAutocompleteOk() (*string, bool) {
	if o == nil || IsNil(o.Autocomplete) {
		return nil, false
	}
	return o.Autocomplete, true
}

// HasAutocomplete returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasAutocomplete() bool {
	if o != nil && !IsNil(o.Autocomplete) {
		return true
	}

	return false
}

// SetAutocomplete gets a reference to the given string and assigns it to the Autocomplete field.
func (o *UiNodeInputAttributes) SetAutocomplete(v string) {
	o.Autocomplete = &v
}

// GetDisabled returns the Disabled field value
func (o *UiNodeInputAttributes) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *UiNodeInputAttributes) SetDisabled(v bool) {
	o.Disabled = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetLabel() UiText {
	if o == nil || IsNil(o.Label) {
		var ret UiText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetLabelOk() (*UiText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given UiText and assigns it to the Label field.
func (o *UiNodeInputAttributes) SetLabel(v UiText) {
	o.Label = &v
}

// GetName returns the Name field value
func (o *UiNodeInputAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UiNodeInputAttributes) SetName(v string) {
	o.Name = v
}

// GetNodeType returns the NodeType field value
func (o *UiNodeInputAttributes) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *UiNodeInputAttributes) SetNodeType(v string) {
	o.NodeType = v
}

// GetOnclick returns the Onclick field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetOnclick() string {
	if o == nil || IsNil(o.Onclick) {
		var ret string
		return ret
	}
	return *o.Onclick
}

// GetOnclickOk returns a tuple with the Onclick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetOnclickOk() (*string, bool) {
	if o == nil || IsNil(o.Onclick) {
		return nil, false
	}
	return o.Onclick, true
}

// HasOnclick returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasOnclick() bool {
	if o != nil && !IsNil(o.Onclick) {
		return true
	}

	return false
}

// SetOnclick gets a reference to the given string and assigns it to the Onclick field.
func (o *UiNodeInputAttributes) SetOnclick(v string) {
	o.Onclick = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *UiNodeInputAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *UiNodeInputAttributes) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *UiNodeInputAttributes) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value
func (o *UiNodeInputAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UiNodeInputAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UiNodeInputAttributes) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UiNodeInputAttributes) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UiNodeInputAttributes) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UiNodeInputAttributes) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *UiNodeInputAttributes) SetValue(v interface{}) {
	o.Value = v
}

func (o UiNodeInputAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiNodeInputAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Autocomplete) {
		toSerialize["autocomplete"] = o.Autocomplete
	}
	toSerialize["disabled"] = o.Disabled
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["name"] = o.Name
	toSerialize["node_type"] = o.NodeType
	if !IsNil(o.Onclick) {
		toSerialize["onclick"] = o.Onclick
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	toSerialize["type"] = o.Type
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UiNodeInputAttributes) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disabled",
		"name",
		"node_type",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUiNodeInputAttributes := _UiNodeInputAttributes{}

	err = json.Unmarshal(bytes, &varUiNodeInputAttributes)

	if err != nil {
		return err
	}

	*o = UiNodeInputAttributes(varUiNodeInputAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "autocomplete")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "node_type")
		delete(additionalProperties, "onclick")
		delete(additionalProperties, "pattern")
		delete(additionalProperties, "required")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUiNodeInputAttributes struct {
	value *UiNodeInputAttributes
	isSet bool
}

func (v NullableUiNodeInputAttributes) Get() *UiNodeInputAttributes {
	return v.value
}

func (v *NullableUiNodeInputAttributes) Set(val *UiNodeInputAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeInputAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeInputAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeInputAttributes(val *UiNodeInputAttributes) *NullableUiNodeInputAttributes {
	return &NullableUiNodeInputAttributes{value: val, isSet: true}
}

func (v NullableUiNodeInputAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeInputAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


