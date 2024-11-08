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
)

// checks if the VerifiableCredentialProof type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifiableCredentialProof{}

// VerifiableCredentialProof struct for VerifiableCredentialProof
type VerifiableCredentialProof struct {
	Jwt *string `json:"jwt,omitempty"`
	ProofType *string `json:"proof_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifiableCredentialProof VerifiableCredentialProof

// NewVerifiableCredentialProof instantiates a new VerifiableCredentialProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiableCredentialProof() *VerifiableCredentialProof {
	this := VerifiableCredentialProof{}
	return &this
}

// NewVerifiableCredentialProofWithDefaults instantiates a new VerifiableCredentialProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiableCredentialProofWithDefaults() *VerifiableCredentialProof {
	this := VerifiableCredentialProof{}
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *VerifiableCredentialProof) GetJwt() string {
	if o == nil || IsNil(o.Jwt) {
		var ret string
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialProof) GetJwtOk() (*string, bool) {
	if o == nil || IsNil(o.Jwt) {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *VerifiableCredentialProof) HasJwt() bool {
	if o != nil && !IsNil(o.Jwt) {
		return true
	}

	return false
}

// SetJwt gets a reference to the given string and assigns it to the Jwt field.
func (o *VerifiableCredentialProof) SetJwt(v string) {
	o.Jwt = &v
}

// GetProofType returns the ProofType field value if set, zero value otherwise.
func (o *VerifiableCredentialProof) GetProofType() string {
	if o == nil || IsNil(o.ProofType) {
		var ret string
		return ret
	}
	return *o.ProofType
}

// GetProofTypeOk returns a tuple with the ProofType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialProof) GetProofTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProofType) {
		return nil, false
	}
	return o.ProofType, true
}

// HasProofType returns a boolean if a field has been set.
func (o *VerifiableCredentialProof) HasProofType() bool {
	if o != nil && !IsNil(o.ProofType) {
		return true
	}

	return false
}

// SetProofType gets a reference to the given string and assigns it to the ProofType field.
func (o *VerifiableCredentialProof) SetProofType(v string) {
	o.ProofType = &v
}

func (o VerifiableCredentialProof) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifiableCredentialProof) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jwt) {
		toSerialize["jwt"] = o.Jwt
	}
	if !IsNil(o.ProofType) {
		toSerialize["proof_type"] = o.ProofType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifiableCredentialProof) UnmarshalJSON(data []byte) (err error) {
	varVerifiableCredentialProof := _VerifiableCredentialProof{}

	err = json.Unmarshal(data, &varVerifiableCredentialProof)

	if err != nil {
		return err
	}

	*o = VerifiableCredentialProof(varVerifiableCredentialProof)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jwt")
		delete(additionalProperties, "proof_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifiableCredentialProof struct {
	value *VerifiableCredentialProof
	isSet bool
}

func (v NullableVerifiableCredentialProof) Get() *VerifiableCredentialProof {
	return v.value
}

func (v *NullableVerifiableCredentialProof) Set(val *VerifiableCredentialProof) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiableCredentialProof) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiableCredentialProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiableCredentialProof(val *VerifiableCredentialProof) *NullableVerifiableCredentialProof {
	return &NullableVerifiableCredentialProof{value: val, isSet: true}
}

func (v NullableVerifiableCredentialProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiableCredentialProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


