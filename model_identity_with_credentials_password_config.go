/*
Ory APIs

# Introduction Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.  ## SDKs This document describes the APIs available in the Ory Network. The APIs are available as SDKs for the following languages:  | Language       | Download SDK                                                     | Documentation                                                                        | | -------------- | ---------------------------------------------------------------- | ------------------------------------------------------------------------------------ | | Dart           | [pub.dev](https://pub.dev/packages/ory_client)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/dart/README.md)       | | .NET           | [nuget.org](https://www.nuget.org/packages/Ory.Client/)          | [README](https://github.com/ory/sdk/blob/master/clients/client/dotnet/README.md)     | | Elixir         | [hex.pm](https://hex.pm/packages/ory_client)                     | [README](https://github.com/ory/sdk/blob/master/clients/client/elixir/README.md)     | | Go             | [github.com](https://github.com/ory/client-go)                   | [README](https://github.com/ory/sdk/blob/master/clients/client/go/README.md)         | | Java           | [maven.org](https://search.maven.org/artifact/sh.ory/ory-client) | [README](https://github.com/ory/sdk/blob/master/clients/client/java/README.md)       | | JavaScript     | [npmjs.com](https://www.npmjs.com/package/@ory/client)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript/README.md) | | JavaScript (With fetch) | [npmjs.com](https://www.npmjs.com/package/@ory/client-fetch)           | [README](https://github.com/ory/sdk/blob/master/clients/client/typescript-fetch/README.md) |  | PHP            | [packagist.org](https://packagist.org/packages/ory/client)       | [README](https://github.com/ory/sdk/blob/master/clients/client/php/README.md)        | | Python         | [pypi.org](https://pypi.org/project/ory-client/)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/python/README.md)     | | Ruby           | [rubygems.org](https://rubygems.org/gems/ory-client)             | [README](https://github.com/ory/sdk/blob/master/clients/client/ruby/README.md)       | | Rust           | [crates.io](https://crates.io/crates/ory-client)                 | [README](https://github.com/ory/sdk/blob/master/clients/client/rust/README.md)       | 

API version: v1.15.12
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IdentityWithCredentialsPasswordConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityWithCredentialsPasswordConfig{}

// IdentityWithCredentialsPasswordConfig Create Identity and Import Password Credentials Configuration
type IdentityWithCredentialsPasswordConfig struct {
	// The hashed password in [PHC format](https://www.ory.sh/docs/kratos/manage-identities/import-user-accounts-identities#hashed-passwords)
	HashedPassword *string `json:"hashed_password,omitempty"`
	// The password in plain text if no hash is available.
	Password *string `json:"password,omitempty"`
	// If set to true, the password will be migrated using the password migration hook.
	UsePasswordMigrationHook *bool `json:"use_password_migration_hook,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityWithCredentialsPasswordConfig IdentityWithCredentialsPasswordConfig

// NewIdentityWithCredentialsPasswordConfig instantiates a new IdentityWithCredentialsPasswordConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityWithCredentialsPasswordConfig() *IdentityWithCredentialsPasswordConfig {
	this := IdentityWithCredentialsPasswordConfig{}
	return &this
}

// NewIdentityWithCredentialsPasswordConfigWithDefaults instantiates a new IdentityWithCredentialsPasswordConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithCredentialsPasswordConfigWithDefaults() *IdentityWithCredentialsPasswordConfig {
	this := IdentityWithCredentialsPasswordConfig{}
	return &this
}

// GetHashedPassword returns the HashedPassword field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPasswordConfig) GetHashedPassword() string {
	if o == nil || IsNil(o.HashedPassword) {
		var ret string
		return ret
	}
	return *o.HashedPassword
}

// GetHashedPasswordOk returns a tuple with the HashedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPasswordConfig) GetHashedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.HashedPassword) {
		return nil, false
	}
	return o.HashedPassword, true
}

// HasHashedPassword returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPasswordConfig) HasHashedPassword() bool {
	if o != nil && !IsNil(o.HashedPassword) {
		return true
	}

	return false
}

// SetHashedPassword gets a reference to the given string and assigns it to the HashedPassword field.
func (o *IdentityWithCredentialsPasswordConfig) SetHashedPassword(v string) {
	o.HashedPassword = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPasswordConfig) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPasswordConfig) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPasswordConfig) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IdentityWithCredentialsPasswordConfig) SetPassword(v string) {
	o.Password = &v
}

// GetUsePasswordMigrationHook returns the UsePasswordMigrationHook field value if set, zero value otherwise.
func (o *IdentityWithCredentialsPasswordConfig) GetUsePasswordMigrationHook() bool {
	if o == nil || IsNil(o.UsePasswordMigrationHook) {
		var ret bool
		return ret
	}
	return *o.UsePasswordMigrationHook
}

// GetUsePasswordMigrationHookOk returns a tuple with the UsePasswordMigrationHook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityWithCredentialsPasswordConfig) GetUsePasswordMigrationHookOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePasswordMigrationHook) {
		return nil, false
	}
	return o.UsePasswordMigrationHook, true
}

// HasUsePasswordMigrationHook returns a boolean if a field has been set.
func (o *IdentityWithCredentialsPasswordConfig) HasUsePasswordMigrationHook() bool {
	if o != nil && !IsNil(o.UsePasswordMigrationHook) {
		return true
	}

	return false
}

// SetUsePasswordMigrationHook gets a reference to the given bool and assigns it to the UsePasswordMigrationHook field.
func (o *IdentityWithCredentialsPasswordConfig) SetUsePasswordMigrationHook(v bool) {
	o.UsePasswordMigrationHook = &v
}

func (o IdentityWithCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityWithCredentialsPasswordConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HashedPassword) {
		toSerialize["hashed_password"] = o.HashedPassword
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UsePasswordMigrationHook) {
		toSerialize["use_password_migration_hook"] = o.UsePasswordMigrationHook
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityWithCredentialsPasswordConfig) UnmarshalJSON(data []byte) (err error) {
	varIdentityWithCredentialsPasswordConfig := _IdentityWithCredentialsPasswordConfig{}

	err = json.Unmarshal(data, &varIdentityWithCredentialsPasswordConfig)

	if err != nil {
		return err
	}

	*o = IdentityWithCredentialsPasswordConfig(varIdentityWithCredentialsPasswordConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hashed_password")
		delete(additionalProperties, "password")
		delete(additionalProperties, "use_password_migration_hook")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityWithCredentialsPasswordConfig struct {
	value *IdentityWithCredentialsPasswordConfig
	isSet bool
}

func (v NullableIdentityWithCredentialsPasswordConfig) Get() *IdentityWithCredentialsPasswordConfig {
	return v.value
}

func (v *NullableIdentityWithCredentialsPasswordConfig) Set(val *IdentityWithCredentialsPasswordConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityWithCredentialsPasswordConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityWithCredentialsPasswordConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityWithCredentialsPasswordConfig(val *IdentityWithCredentialsPasswordConfig) *NullableIdentityWithCredentialsPasswordConfig {
	return &NullableIdentityWithCredentialsPasswordConfig{value: val, isSet: true}
}

func (v NullableIdentityWithCredentialsPasswordConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityWithCredentialsPasswordConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


