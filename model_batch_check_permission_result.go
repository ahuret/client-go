/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.15.10
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the BatchCheckPermissionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchCheckPermissionResult{}

// BatchCheckPermissionResult Batch Check Permission Result
type BatchCheckPermissionResult struct {
	// An array of check results. The order aligns with the input order.
	Results []CheckPermissionResultWithError `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _BatchCheckPermissionResult BatchCheckPermissionResult

// NewBatchCheckPermissionResult instantiates a new BatchCheckPermissionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCheckPermissionResult(results []CheckPermissionResultWithError) *BatchCheckPermissionResult {
	this := BatchCheckPermissionResult{}
	this.Results = results
	return &this
}

// NewBatchCheckPermissionResultWithDefaults instantiates a new BatchCheckPermissionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCheckPermissionResultWithDefaults() *BatchCheckPermissionResult {
	this := BatchCheckPermissionResult{}
	return &this
}

// GetResults returns the Results field value
func (o *BatchCheckPermissionResult) GetResults() []CheckPermissionResultWithError {
	if o == nil {
		var ret []CheckPermissionResultWithError
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchCheckPermissionResult) GetResultsOk() ([]CheckPermissionResultWithError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchCheckPermissionResult) SetResults(v []CheckPermissionResultWithError) {
	o.Results = v
}

func (o BatchCheckPermissionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchCheckPermissionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BatchCheckPermissionResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varBatchCheckPermissionResult := _BatchCheckPermissionResult{}

	err = json.Unmarshal(data, &varBatchCheckPermissionResult)

	if err != nil {
		return err
	}

	*o = BatchCheckPermissionResult(varBatchCheckPermissionResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBatchCheckPermissionResult struct {
	value *BatchCheckPermissionResult
	isSet bool
}

func (v NullableBatchCheckPermissionResult) Get() *BatchCheckPermissionResult {
	return v.value
}

func (v *NullableBatchCheckPermissionResult) Set(val *BatchCheckPermissionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCheckPermissionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCheckPermissionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCheckPermissionResult(val *BatchCheckPermissionResult) *NullableBatchCheckPermissionResult {
	return &NullableBatchCheckPermissionResult{value: val, isSet: true}
}

func (v NullableBatchCheckPermissionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCheckPermissionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


