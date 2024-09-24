/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the GetMetricsEventTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMetricsEventTypesResponse{}

// GetMetricsEventTypesResponse Response of the getMetricsEventTypes endpoint
type GetMetricsEventTypesResponse struct {
	// The list of data points.
	Events []string `json:"events"`
	AdditionalProperties map[string]interface{}
}

type _GetMetricsEventTypesResponse GetMetricsEventTypesResponse

// NewGetMetricsEventTypesResponse instantiates a new GetMetricsEventTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMetricsEventTypesResponse(events []string) *GetMetricsEventTypesResponse {
	this := GetMetricsEventTypesResponse{}
	this.Events = events
	return &this
}

// NewGetMetricsEventTypesResponseWithDefaults instantiates a new GetMetricsEventTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMetricsEventTypesResponseWithDefaults() *GetMetricsEventTypesResponse {
	this := GetMetricsEventTypesResponse{}
	return &this
}

// GetEvents returns the Events field value
func (o *GetMetricsEventTypesResponse) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *GetMetricsEventTypesResponse) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *GetMetricsEventTypesResponse) SetEvents(v []string) {
	o.Events = v
}

func (o GetMetricsEventTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMetricsEventTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMetricsEventTypesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varGetMetricsEventTypesResponse := _GetMetricsEventTypesResponse{}

	err = json.Unmarshal(data, &varGetMetricsEventTypesResponse)

	if err != nil {
		return err
	}

	*o = GetMetricsEventTypesResponse(varGetMetricsEventTypesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMetricsEventTypesResponse struct {
	value *GetMetricsEventTypesResponse
	isSet bool
}

func (v NullableGetMetricsEventTypesResponse) Get() *GetMetricsEventTypesResponse {
	return v.value
}

func (v *NullableGetMetricsEventTypesResponse) Set(val *GetMetricsEventTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMetricsEventTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMetricsEventTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMetricsEventTypesResponse(val *GetMetricsEventTypesResponse) *NullableGetMetricsEventTypesResponse {
	return &NullableGetMetricsEventTypesResponse{value: val, isSet: true}
}

func (v NullableGetMetricsEventTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMetricsEventTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


