/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.103
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// ProjectRevisionHook struct for ProjectRevisionHook
type ProjectRevisionHook struct {
	// The Hooks Config Key
	ConfigKey string `json:"config_key"`
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The Hook Type
	Hook string `json:"hook"`
	Id *string `json:"id,omitempty"`
	ProjectRevisionId *string `json:"project_revision_id,omitempty"`
	// Last Time Project's Revision was Updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether to send the API Key in the HTTP Header or as a HTTP Cookie
	WebHookConfigAuthApiKeyIn *string `json:"web_hook_config_auth_api_key_in,omitempty"`
	// The name of the api key
	WebHookConfigAuthApiKeyName *string `json:"web_hook_config_auth_api_key_name,omitempty"`
	// The value of the api key
	WebHookConfigAuthApiKeyValue *string `json:"web_hook_config_auth_api_key_value,omitempty"`
	// The password to be sent in the HTTP Basic Auth Header
	WebHookConfigAuthBasicAuthPassword *string `json:"web_hook_config_auth_basic_auth_password,omitempty"`
	// The username to be sent in the HTTP Basic Auth Header
	WebHookConfigAuthBasicAuthUser *string `json:"web_hook_config_auth_basic_auth_user,omitempty"`
	// HTTP Auth Method to use for the Web-Hook
	WebHookConfigAuthType *string `json:"web_hook_config_auth_type,omitempty"`
	// URI pointing to the JsonNet template used for Web-Hook payload generation. Only used for those HTTP methods, which support HTTP body payloads.
	WebHookConfigBody *string `json:"web_hook_config_body,omitempty"`
	// The HTTP method to use (GET, POST, etc) for the Web-Hook
	WebHookConfigMethod *string `json:"web_hook_config_method,omitempty"`
	// The URL the Web-Hook should call
	WebHookConfigUrl *string `json:"web_hook_config_url,omitempty"`
}

// NewProjectRevisionHook instantiates a new ProjectRevisionHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectRevisionHook(configKey string, hook string) *ProjectRevisionHook {
	this := ProjectRevisionHook{}
	this.ConfigKey = configKey
	this.Hook = hook
	return &this
}

// NewProjectRevisionHookWithDefaults instantiates a new ProjectRevisionHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectRevisionHookWithDefaults() *ProjectRevisionHook {
	this := ProjectRevisionHook{}
	return &this
}

// GetConfigKey returns the ConfigKey field value
func (o *ProjectRevisionHook) GetConfigKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigKey
}

// GetConfigKeyOk returns a tuple with the ConfigKey field value
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetConfigKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigKey, true
}

// SetConfigKey sets field value
func (o *ProjectRevisionHook) SetConfigKey(v string) {
	o.ConfigKey = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectRevisionHook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHook returns the Hook field value
func (o *ProjectRevisionHook) GetHook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hook
}

// GetHookOk returns a tuple with the Hook field value
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetHookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hook, true
}

// SetHook sets field value
func (o *ProjectRevisionHook) SetHook(v string) {
	o.Hook = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProjectRevisionHook) SetId(v string) {
	o.Id = &v
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetProjectRevisionId() string {
	if o == nil || o.ProjectRevisionId == nil {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || o.ProjectRevisionId == nil {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasProjectRevisionId() bool {
	if o != nil && o.ProjectRevisionId != nil {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *ProjectRevisionHook) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectRevisionHook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWebHookConfigAuthApiKeyIn returns the WebHookConfigAuthApiKeyIn field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyIn() string {
	if o == nil || o.WebHookConfigAuthApiKeyIn == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyIn
}

// GetWebHookConfigAuthApiKeyInOk returns a tuple with the WebHookConfigAuthApiKeyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyInOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthApiKeyIn == nil {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyIn, true
}

// HasWebHookConfigAuthApiKeyIn returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthApiKeyIn() bool {
	if o != nil && o.WebHookConfigAuthApiKeyIn != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyIn gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyIn field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthApiKeyIn(v string) {
	o.WebHookConfigAuthApiKeyIn = &v
}

// GetWebHookConfigAuthApiKeyName returns the WebHookConfigAuthApiKeyName field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyName() string {
	if o == nil || o.WebHookConfigAuthApiKeyName == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyName
}

// GetWebHookConfigAuthApiKeyNameOk returns a tuple with the WebHookConfigAuthApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyNameOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthApiKeyName == nil {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyName, true
}

// HasWebHookConfigAuthApiKeyName returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthApiKeyName() bool {
	if o != nil && o.WebHookConfigAuthApiKeyName != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyName gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyName field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthApiKeyName(v string) {
	o.WebHookConfigAuthApiKeyName = &v
}

// GetWebHookConfigAuthApiKeyValue returns the WebHookConfigAuthApiKeyValue field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyValue() string {
	if o == nil || o.WebHookConfigAuthApiKeyValue == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyValue
}

// GetWebHookConfigAuthApiKeyValueOk returns a tuple with the WebHookConfigAuthApiKeyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthApiKeyValueOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthApiKeyValue == nil {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyValue, true
}

// HasWebHookConfigAuthApiKeyValue returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthApiKeyValue() bool {
	if o != nil && o.WebHookConfigAuthApiKeyValue != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyValue gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyValue field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthApiKeyValue(v string) {
	o.WebHookConfigAuthApiKeyValue = &v
}

// GetWebHookConfigAuthBasicAuthPassword returns the WebHookConfigAuthBasicAuthPassword field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthBasicAuthPassword() string {
	if o == nil || o.WebHookConfigAuthBasicAuthPassword == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthBasicAuthPassword
}

// GetWebHookConfigAuthBasicAuthPasswordOk returns a tuple with the WebHookConfigAuthBasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthBasicAuthPasswordOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthBasicAuthPassword == nil {
		return nil, false
	}
	return o.WebHookConfigAuthBasicAuthPassword, true
}

// HasWebHookConfigAuthBasicAuthPassword returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthBasicAuthPassword() bool {
	if o != nil && o.WebHookConfigAuthBasicAuthPassword != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthBasicAuthPassword gets a reference to the given string and assigns it to the WebHookConfigAuthBasicAuthPassword field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthBasicAuthPassword(v string) {
	o.WebHookConfigAuthBasicAuthPassword = &v
}

// GetWebHookConfigAuthBasicAuthUser returns the WebHookConfigAuthBasicAuthUser field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthBasicAuthUser() string {
	if o == nil || o.WebHookConfigAuthBasicAuthUser == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthBasicAuthUser
}

// GetWebHookConfigAuthBasicAuthUserOk returns a tuple with the WebHookConfigAuthBasicAuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthBasicAuthUserOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthBasicAuthUser == nil {
		return nil, false
	}
	return o.WebHookConfigAuthBasicAuthUser, true
}

// HasWebHookConfigAuthBasicAuthUser returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthBasicAuthUser() bool {
	if o != nil && o.WebHookConfigAuthBasicAuthUser != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthBasicAuthUser gets a reference to the given string and assigns it to the WebHookConfigAuthBasicAuthUser field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthBasicAuthUser(v string) {
	o.WebHookConfigAuthBasicAuthUser = &v
}

// GetWebHookConfigAuthType returns the WebHookConfigAuthType field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigAuthType() string {
	if o == nil || o.WebHookConfigAuthType == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthType
}

// GetWebHookConfigAuthTypeOk returns a tuple with the WebHookConfigAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigAuthTypeOk() (*string, bool) {
	if o == nil || o.WebHookConfigAuthType == nil {
		return nil, false
	}
	return o.WebHookConfigAuthType, true
}

// HasWebHookConfigAuthType returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigAuthType() bool {
	if o != nil && o.WebHookConfigAuthType != nil {
		return true
	}

	return false
}

// SetWebHookConfigAuthType gets a reference to the given string and assigns it to the WebHookConfigAuthType field.
func (o *ProjectRevisionHook) SetWebHookConfigAuthType(v string) {
	o.WebHookConfigAuthType = &v
}

// GetWebHookConfigBody returns the WebHookConfigBody field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigBody() string {
	if o == nil || o.WebHookConfigBody == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigBody
}

// GetWebHookConfigBodyOk returns a tuple with the WebHookConfigBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigBodyOk() (*string, bool) {
	if o == nil || o.WebHookConfigBody == nil {
		return nil, false
	}
	return o.WebHookConfigBody, true
}

// HasWebHookConfigBody returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigBody() bool {
	if o != nil && o.WebHookConfigBody != nil {
		return true
	}

	return false
}

// SetWebHookConfigBody gets a reference to the given string and assigns it to the WebHookConfigBody field.
func (o *ProjectRevisionHook) SetWebHookConfigBody(v string) {
	o.WebHookConfigBody = &v
}

// GetWebHookConfigMethod returns the WebHookConfigMethod field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigMethod() string {
	if o == nil || o.WebHookConfigMethod == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigMethod
}

// GetWebHookConfigMethodOk returns a tuple with the WebHookConfigMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigMethodOk() (*string, bool) {
	if o == nil || o.WebHookConfigMethod == nil {
		return nil, false
	}
	return o.WebHookConfigMethod, true
}

// HasWebHookConfigMethod returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigMethod() bool {
	if o != nil && o.WebHookConfigMethod != nil {
		return true
	}

	return false
}

// SetWebHookConfigMethod gets a reference to the given string and assigns it to the WebHookConfigMethod field.
func (o *ProjectRevisionHook) SetWebHookConfigMethod(v string) {
	o.WebHookConfigMethod = &v
}

// GetWebHookConfigUrl returns the WebHookConfigUrl field value if set, zero value otherwise.
func (o *ProjectRevisionHook) GetWebHookConfigUrl() string {
	if o == nil || o.WebHookConfigUrl == nil {
		var ret string
		return ret
	}
	return *o.WebHookConfigUrl
}

// GetWebHookConfigUrlOk returns a tuple with the WebHookConfigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectRevisionHook) GetWebHookConfigUrlOk() (*string, bool) {
	if o == nil || o.WebHookConfigUrl == nil {
		return nil, false
	}
	return o.WebHookConfigUrl, true
}

// HasWebHookConfigUrl returns a boolean if a field has been set.
func (o *ProjectRevisionHook) HasWebHookConfigUrl() bool {
	if o != nil && o.WebHookConfigUrl != nil {
		return true
	}

	return false
}

// SetWebHookConfigUrl gets a reference to the given string and assigns it to the WebHookConfigUrl field.
func (o *ProjectRevisionHook) SetWebHookConfigUrl(v string) {
	o.WebHookConfigUrl = &v
}

func (o ProjectRevisionHook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config_key"] = o.ConfigKey
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["hook"] = o.Hook
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProjectRevisionId != nil {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.WebHookConfigAuthApiKeyIn != nil {
		toSerialize["web_hook_config_auth_api_key_in"] = o.WebHookConfigAuthApiKeyIn
	}
	if o.WebHookConfigAuthApiKeyName != nil {
		toSerialize["web_hook_config_auth_api_key_name"] = o.WebHookConfigAuthApiKeyName
	}
	if o.WebHookConfigAuthApiKeyValue != nil {
		toSerialize["web_hook_config_auth_api_key_value"] = o.WebHookConfigAuthApiKeyValue
	}
	if o.WebHookConfigAuthBasicAuthPassword != nil {
		toSerialize["web_hook_config_auth_basic_auth_password"] = o.WebHookConfigAuthBasicAuthPassword
	}
	if o.WebHookConfigAuthBasicAuthUser != nil {
		toSerialize["web_hook_config_auth_basic_auth_user"] = o.WebHookConfigAuthBasicAuthUser
	}
	if o.WebHookConfigAuthType != nil {
		toSerialize["web_hook_config_auth_type"] = o.WebHookConfigAuthType
	}
	if o.WebHookConfigBody != nil {
		toSerialize["web_hook_config_body"] = o.WebHookConfigBody
	}
	if o.WebHookConfigMethod != nil {
		toSerialize["web_hook_config_method"] = o.WebHookConfigMethod
	}
	if o.WebHookConfigUrl != nil {
		toSerialize["web_hook_config_url"] = o.WebHookConfigUrl
	}
	return json.Marshal(toSerialize)
}

type NullableProjectRevisionHook struct {
	value *ProjectRevisionHook
	isSet bool
}

func (v NullableProjectRevisionHook) Get() *ProjectRevisionHook {
	return v.value
}

func (v *NullableProjectRevisionHook) Set(val *ProjectRevisionHook) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectRevisionHook) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectRevisionHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectRevisionHook(val *ProjectRevisionHook) *NullableProjectRevisionHook {
	return &NullableProjectRevisionHook{value: val, isSet: true}
}

func (v NullableProjectRevisionHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectRevisionHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


