/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ContinueWithVerificationUi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueWithVerificationUi{}

// ContinueWithVerificationUi Indicates, that the UI flow could be continued by showing a verification ui
type ContinueWithVerificationUi struct {
	// Action will always be `show_verification_ui` show_verification_ui ContinueWithActionShowVerificationUIString
	Action string `json:"action"`
	Flow ContinueWithVerificationUiFlow `json:"flow"`
	AdditionalProperties map[string]interface{}
}

type _ContinueWithVerificationUi ContinueWithVerificationUi

// NewContinueWithVerificationUi instantiates a new ContinueWithVerificationUi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithVerificationUi(action string, flow ContinueWithVerificationUiFlow) *ContinueWithVerificationUi {
	this := ContinueWithVerificationUi{}
	this.Action = action
	this.Flow = flow
	return &this
}

// NewContinueWithVerificationUiWithDefaults instantiates a new ContinueWithVerificationUi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithVerificationUiWithDefaults() *ContinueWithVerificationUi {
	this := ContinueWithVerificationUi{}
	return &this
}

// GetAction returns the Action field value
func (o *ContinueWithVerificationUi) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContinueWithVerificationUi) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContinueWithVerificationUi) SetAction(v string) {
	o.Action = v
}

// GetFlow returns the Flow field value
func (o *ContinueWithVerificationUi) GetFlow() ContinueWithVerificationUiFlow {
	if o == nil {
		var ret ContinueWithVerificationUiFlow
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *ContinueWithVerificationUi) GetFlowOk() (*ContinueWithVerificationUiFlow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *ContinueWithVerificationUi) SetFlow(v ContinueWithVerificationUiFlow) {
	o.Flow = v
}

func (o ContinueWithVerificationUi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueWithVerificationUi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["flow"] = o.Flow

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContinueWithVerificationUi) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"flow",
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

	varContinueWithVerificationUi := _ContinueWithVerificationUi{}

	err = json.Unmarshal(bytes, &varContinueWithVerificationUi)

	if err != nil {
		return err
	}

	*o = ContinueWithVerificationUi(varContinueWithVerificationUi)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "flow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContinueWithVerificationUi struct {
	value *ContinueWithVerificationUi
	isSet bool
}

func (v NullableContinueWithVerificationUi) Get() *ContinueWithVerificationUi {
	return v.value
}

func (v *NullableContinueWithVerificationUi) Set(val *ContinueWithVerificationUi) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithVerificationUi) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithVerificationUi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithVerificationUi(val *ContinueWithVerificationUi) *NullableContinueWithVerificationUi {
	return &NullableContinueWithVerificationUi{value: val, isSet: true}
}

func (v NullableContinueWithVerificationUi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithVerificationUi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


