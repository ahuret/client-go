/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.8.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Identity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Identity{}

// Identity An [identity](https://www.ory.sh/docs/kratos/concepts/identity-user-model) represents a (human) user in Ory.
type Identity struct {
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Credentials represents all credentials that can be used for authenticating this identity.
	Credentials *map[string]IdentityCredentials `json:"credentials,omitempty"`
	// ID is the identity's unique identifier.  The Identity ID can not be changed and can not be chosen. This ensures future compatibility and optimization for distributed stores such as CockroachDB.
	Id string `json:"id"`
	// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
	MetadataAdmin map[string]interface{} `json:"metadata_admin,omitempty"`
	// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
	MetadataPublic map[string]interface{} `json:"metadata_public,omitempty"`
	OrganizationId NullableString `json:"organization_id,omitempty"`
	// RecoveryAddresses contains all the addresses that can be used to recover an identity.
	RecoveryAddresses []RecoveryIdentityAddress `json:"recovery_addresses,omitempty"`
	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	SchemaId string `json:"schema_id"`
	// SchemaURL is the URL of the endpoint where the identity's traits schema can be fetched from.  format: url
	SchemaUrl string `json:"schema_url"`
	// State is the identity's state.  This value has currently no effect. active StateActive inactive StateInactive
	State *string `json:"state,omitempty"`
	StateChangedAt *time.Time `json:"state_changed_at,omitempty"`
	// Traits represent an identity's traits. The identity is able to create, modify, and delete traits in a self-service manner. The input will always be validated against the JSON Schema defined in `schema_url`.
	Traits interface{} `json:"traits"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// VerifiableAddresses contains all the addresses that can be verified by the user.
	VerifiableAddresses []VerifiableIdentityAddress `json:"verifiable_addresses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Identity Identity

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(id string, schemaId string, schemaUrl string, traits interface{}) *Identity {
	this := Identity{}
	this.Id = id
	this.SchemaId = schemaId
	this.SchemaUrl = schemaUrl
	this.Traits = traits
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Identity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Identity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Identity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Identity) GetCredentials() map[string]IdentityCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret map[string]IdentityCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetCredentialsOk() (*map[string]IdentityCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Identity) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]IdentityCredentials and assigns it to the Credentials field.
func (o *Identity) SetCredentials(v map[string]IdentityCredentials) {
	o.Credentials = &v
}

// GetId returns the Id field value
func (o *Identity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Identity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Identity) SetId(v string) {
	o.Id = v
}

// GetMetadataAdmin returns the MetadataAdmin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetMetadataAdmin() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetadataAdmin
}

// GetMetadataAdminOk returns a tuple with the MetadataAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetMetadataAdminOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetadataAdmin) {
		return map[string]interface{}{}, false
	}
	return o.MetadataAdmin, true
}

// HasMetadataAdmin returns a boolean if a field has been set.
func (o *Identity) HasMetadataAdmin() bool {
	if o != nil && IsNil(o.MetadataAdmin) {
		return true
	}

	return false
}

// SetMetadataAdmin gets a reference to the given map[string]interface{} and assigns it to the MetadataAdmin field.
func (o *Identity) SetMetadataAdmin(v map[string]interface{}) {
	o.MetadataAdmin = v
}

// GetMetadataPublic returns the MetadataPublic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetMetadataPublic() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MetadataPublic
}

// GetMetadataPublicOk returns a tuple with the MetadataPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetMetadataPublicOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetadataPublic) {
		return map[string]interface{}{}, false
	}
	return o.MetadataPublic, true
}

// HasMetadataPublic returns a boolean if a field has been set.
func (o *Identity) HasMetadataPublic() bool {
	if o != nil && IsNil(o.MetadataPublic) {
		return true
	}

	return false
}

// SetMetadataPublic gets a reference to the given map[string]interface{} and assigns it to the MetadataPublic field.
func (o *Identity) SetMetadataPublic(v map[string]interface{}) {
	o.MetadataPublic = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *Identity) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableString and assigns it to the OrganizationId field.
func (o *Identity) SetOrganizationId(v string) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *Identity) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *Identity) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetRecoveryAddresses returns the RecoveryAddresses field value if set, zero value otherwise.
func (o *Identity) GetRecoveryAddresses() []RecoveryIdentityAddress {
	if o == nil || IsNil(o.RecoveryAddresses) {
		var ret []RecoveryIdentityAddress
		return ret
	}
	return o.RecoveryAddresses
}

// GetRecoveryAddressesOk returns a tuple with the RecoveryAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetRecoveryAddressesOk() ([]RecoveryIdentityAddress, bool) {
	if o == nil || IsNil(o.RecoveryAddresses) {
		return nil, false
	}
	return o.RecoveryAddresses, true
}

// HasRecoveryAddresses returns a boolean if a field has been set.
func (o *Identity) HasRecoveryAddresses() bool {
	if o != nil && !IsNil(o.RecoveryAddresses) {
		return true
	}

	return false
}

// SetRecoveryAddresses gets a reference to the given []RecoveryIdentityAddress and assigns it to the RecoveryAddresses field.
func (o *Identity) SetRecoveryAddresses(v []RecoveryIdentityAddress) {
	o.RecoveryAddresses = v
}

// GetSchemaId returns the SchemaId field value
func (o *Identity) GetSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *Identity) GetSchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *Identity) SetSchemaId(v string) {
	o.SchemaId = v
}

// GetSchemaUrl returns the SchemaUrl field value
func (o *Identity) GetSchemaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaUrl
}

// GetSchemaUrlOk returns a tuple with the SchemaUrl field value
// and a boolean to check if the value has been set.
func (o *Identity) GetSchemaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaUrl, true
}

// SetSchemaUrl sets field value
func (o *Identity) SetSchemaUrl(v string) {
	o.SchemaUrl = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Identity) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Identity) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Identity) SetState(v string) {
	o.State = &v
}

// GetStateChangedAt returns the StateChangedAt field value if set, zero value otherwise.
func (o *Identity) GetStateChangedAt() time.Time {
	if o == nil || IsNil(o.StateChangedAt) {
		var ret time.Time
		return ret
	}
	return *o.StateChangedAt
}

// GetStateChangedAtOk returns a tuple with the StateChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetStateChangedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StateChangedAt) {
		return nil, false
	}
	return o.StateChangedAt, true
}

// HasStateChangedAt returns a boolean if a field has been set.
func (o *Identity) HasStateChangedAt() bool {
	if o != nil && !IsNil(o.StateChangedAt) {
		return true
	}

	return false
}

// SetStateChangedAt gets a reference to the given time.Time and assigns it to the StateChangedAt field.
func (o *Identity) SetStateChangedAt(v time.Time) {
	o.StateChangedAt = &v
}

// GetTraits returns the Traits field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Identity) GetTraits() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetTraitsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Traits) {
		return nil, false
	}
	return &o.Traits, true
}

// SetTraits sets field value
func (o *Identity) SetTraits(v interface{}) {
	o.Traits = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Identity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Identity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Identity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerifiableAddresses returns the VerifiableAddresses field value if set, zero value otherwise.
func (o *Identity) GetVerifiableAddresses() []VerifiableIdentityAddress {
	if o == nil || IsNil(o.VerifiableAddresses) {
		var ret []VerifiableIdentityAddress
		return ret
	}
	return o.VerifiableAddresses
}

// GetVerifiableAddressesOk returns a tuple with the VerifiableAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetVerifiableAddressesOk() ([]VerifiableIdentityAddress, bool) {
	if o == nil || IsNil(o.VerifiableAddresses) {
		return nil, false
	}
	return o.VerifiableAddresses, true
}

// HasVerifiableAddresses returns a boolean if a field has been set.
func (o *Identity) HasVerifiableAddresses() bool {
	if o != nil && !IsNil(o.VerifiableAddresses) {
		return true
	}

	return false
}

// SetVerifiableAddresses gets a reference to the given []VerifiableIdentityAddress and assigns it to the VerifiableAddresses field.
func (o *Identity) SetVerifiableAddresses(v []VerifiableIdentityAddress) {
	o.VerifiableAddresses = v
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Identity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["id"] = o.Id
	if o.MetadataAdmin != nil {
		toSerialize["metadata_admin"] = o.MetadataAdmin
	}
	if o.MetadataPublic != nil {
		toSerialize["metadata_public"] = o.MetadataPublic
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organization_id"] = o.OrganizationId.Get()
	}
	if !IsNil(o.RecoveryAddresses) {
		toSerialize["recovery_addresses"] = o.RecoveryAddresses
	}
	toSerialize["schema_id"] = o.SchemaId
	toSerialize["schema_url"] = o.SchemaUrl
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateChangedAt) {
		toSerialize["state_changed_at"] = o.StateChangedAt
	}
	if o.Traits != nil {
		toSerialize["traits"] = o.Traits
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VerifiableAddresses) {
		toSerialize["verifiable_addresses"] = o.VerifiableAddresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Identity) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"schema_id",
		"schema_url",
		"traits",
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

	varIdentity := _Identity{}

	err = json.Unmarshal(bytes, &varIdentity)

	if err != nil {
		return err
	}

	*o = Identity(varIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadata_admin")
		delete(additionalProperties, "metadata_public")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "recovery_addresses")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "schema_url")
		delete(additionalProperties, "state")
		delete(additionalProperties, "state_changed_at")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "verifiable_addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


