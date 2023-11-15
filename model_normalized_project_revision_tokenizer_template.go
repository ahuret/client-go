/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.4.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the NormalizedProjectRevisionTokenizerTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NormalizedProjectRevisionTokenizerTemplate{}

// NormalizedProjectRevisionTokenizerTemplate struct for NormalizedProjectRevisionTokenizerTemplate
type NormalizedProjectRevisionTokenizerTemplate struct {
	// Claims mapper URL
	ClaimsMapperUrl *string `json:"claims_mapper_url,omitempty"`
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The revision ID.
	Id *string `json:"id,omitempty"`
	// JSON Web Key URL
	JwksUrl *string `json:"jwks_url,omitempty"`
	// The unique key of the template
	Key *string `json:"key,omitempty"`
	// The Revision's ID this schema belongs to
	ProjectRevisionId *string `json:"project_revision_id,omitempty"`
	// Token time to live
	Ttl *string `json:"ttl,omitempty"`
	// Last Time Project's Revision was Updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NormalizedProjectRevisionTokenizerTemplate NormalizedProjectRevisionTokenizerTemplate

// NewNormalizedProjectRevisionTokenizerTemplate instantiates a new NormalizedProjectRevisionTokenizerTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionTokenizerTemplate() *NormalizedProjectRevisionTokenizerTemplate {
	this := NormalizedProjectRevisionTokenizerTemplate{}
	var ttl string = "1m"
	this.Ttl = &ttl
	return &this
}

// NewNormalizedProjectRevisionTokenizerTemplateWithDefaults instantiates a new NormalizedProjectRevisionTokenizerTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionTokenizerTemplateWithDefaults() *NormalizedProjectRevisionTokenizerTemplate {
	this := NormalizedProjectRevisionTokenizerTemplate{}
	var ttl string = "1m"
	this.Ttl = &ttl
	return &this
}

// GetClaimsMapperUrl returns the ClaimsMapperUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetClaimsMapperUrl() string {
	if o == nil || IsNil(o.ClaimsMapperUrl) {
		var ret string
		return ret
	}
	return *o.ClaimsMapperUrl
}

// GetClaimsMapperUrlOk returns a tuple with the ClaimsMapperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetClaimsMapperUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimsMapperUrl) {
		return nil, false
	}
	return o.ClaimsMapperUrl, true
}

// HasClaimsMapperUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasClaimsMapperUrl() bool {
	if o != nil && !IsNil(o.ClaimsMapperUrl) {
		return true
	}

	return false
}

// SetClaimsMapperUrl gets a reference to the given string and assigns it to the ClaimsMapperUrl field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetClaimsMapperUrl(v string) {
	o.ClaimsMapperUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetId(v string) {
	o.Id = &v
}

// GetJwksUrl returns the JwksUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetJwksUrl() string {
	if o == nil || IsNil(o.JwksUrl) {
		var ret string
		return ret
	}
	return *o.JwksUrl
}

// GetJwksUrlOk returns a tuple with the JwksUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetJwksUrlOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUrl) {
		return nil, false
	}
	return o.JwksUrl, true
}

// HasJwksUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasJwksUrl() bool {
	if o != nil && !IsNil(o.JwksUrl) {
		return true
	}

	return false
}

// SetJwksUrl gets a reference to the given string and assigns it to the JwksUrl field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetJwksUrl(v string) {
	o.JwksUrl = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetKey(v string) {
	o.Key = &v
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetProjectRevisionId() string {
	if o == nil || IsNil(o.ProjectRevisionId) {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectRevisionId) {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasProjectRevisionId() bool {
	if o != nil && !IsNil(o.ProjectRevisionId) {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetTtl(v string) {
	o.Ttl = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionTokenizerTemplate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionTokenizerTemplate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NormalizedProjectRevisionTokenizerTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NormalizedProjectRevisionTokenizerTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClaimsMapperUrl) {
		toSerialize["claims_mapper_url"] = o.ClaimsMapperUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JwksUrl) {
		toSerialize["jwks_url"] = o.JwksUrl
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.ProjectRevisionId) {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NormalizedProjectRevisionTokenizerTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varNormalizedProjectRevisionTokenizerTemplate := _NormalizedProjectRevisionTokenizerTemplate{}

	err = json.Unmarshal(bytes, &varNormalizedProjectRevisionTokenizerTemplate)

	if err != nil {
		return err
	}

	*o = NormalizedProjectRevisionTokenizerTemplate(varNormalizedProjectRevisionTokenizerTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "claims_mapper_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "jwks_url")
		delete(additionalProperties, "key")
		delete(additionalProperties, "project_revision_id")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNormalizedProjectRevisionTokenizerTemplate struct {
	value *NormalizedProjectRevisionTokenizerTemplate
	isSet bool
}

func (v NullableNormalizedProjectRevisionTokenizerTemplate) Get() *NormalizedProjectRevisionTokenizerTemplate {
	return v.value
}

func (v *NullableNormalizedProjectRevisionTokenizerTemplate) Set(val *NormalizedProjectRevisionTokenizerTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionTokenizerTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionTokenizerTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionTokenizerTemplate(val *NormalizedProjectRevisionTokenizerTemplate) *NullableNormalizedProjectRevisionTokenizerTemplate {
	return &NullableNormalizedProjectRevisionTokenizerTemplate{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionTokenizerTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionTokenizerTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


