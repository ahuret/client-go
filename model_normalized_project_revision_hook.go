/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.3
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the NormalizedProjectRevisionHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NormalizedProjectRevisionHook{}

// NormalizedProjectRevisionHook struct for NormalizedProjectRevisionHook
type NormalizedProjectRevisionHook struct {
	// The Hooks Config Key
	ConfigKey string `json:"config_key"`
	// The Project's Revision Creation Date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The Hook Type
	Hook string `json:"hook"`
	// ID of the entry
	Id *string `json:"id,omitempty"`
	// The Revision's ID this schema belongs to
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
	// If enabled allows the web hook to interrupt / abort the self-service flow. It only applies to certain flows (registration/verification/login/settings) and requires a valid response format.
	WebHookConfigCanInterrupt *bool `json:"web_hook_config_can_interrupt,omitempty"`
	// The HTTP method to use (GET, POST, etc) for the Web-Hook
	WebHookConfigMethod *string `json:"web_hook_config_method,omitempty"`
	// Whether to ignore the Web Hook response
	WebHookConfigResponseIgnore *bool `json:"web_hook_config_response_ignore,omitempty"`
	// Whether to parse the Web Hook response
	WebHookConfigResponseParse *bool `json:"web_hook_config_response_parse,omitempty"`
	// The URL the Web-Hook should call
	WebHookConfigUrl *string `json:"web_hook_config_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NormalizedProjectRevisionHook NormalizedProjectRevisionHook

// NewNormalizedProjectRevisionHook instantiates a new NormalizedProjectRevisionHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionHook(configKey string, hook string) *NormalizedProjectRevisionHook {
	this := NormalizedProjectRevisionHook{}
	this.ConfigKey = configKey
	this.Hook = hook
	return &this
}

// NewNormalizedProjectRevisionHookWithDefaults instantiates a new NormalizedProjectRevisionHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionHookWithDefaults() *NormalizedProjectRevisionHook {
	this := NormalizedProjectRevisionHook{}
	return &this
}

// GetConfigKey returns the ConfigKey field value
func (o *NormalizedProjectRevisionHook) GetConfigKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigKey
}

// GetConfigKeyOk returns a tuple with the ConfigKey field value
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetConfigKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigKey, true
}

// SetConfigKey sets field value
func (o *NormalizedProjectRevisionHook) SetConfigKey(v string) {
	o.ConfigKey = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionHook) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHook returns the Hook field value
func (o *NormalizedProjectRevisionHook) GetHook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hook
}

// GetHookOk returns a tuple with the Hook field value
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetHookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hook, true
}

// SetHook sets field value
func (o *NormalizedProjectRevisionHook) SetHook(v string) {
	o.Hook = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NormalizedProjectRevisionHook) SetId(v string) {
	o.Id = &v
}

// GetProjectRevisionId returns the ProjectRevisionId field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetProjectRevisionId() string {
	if o == nil || IsNil(o.ProjectRevisionId) {
		var ret string
		return ret
	}
	return *o.ProjectRevisionId
}

// GetProjectRevisionIdOk returns a tuple with the ProjectRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetProjectRevisionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectRevisionId) {
		return nil, false
	}
	return o.ProjectRevisionId, true
}

// HasProjectRevisionId returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasProjectRevisionId() bool {
	if o != nil && !IsNil(o.ProjectRevisionId) {
		return true
	}

	return false
}

// SetProjectRevisionId gets a reference to the given string and assigns it to the ProjectRevisionId field.
func (o *NormalizedProjectRevisionHook) SetProjectRevisionId(v string) {
	o.ProjectRevisionId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionHook) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWebHookConfigAuthApiKeyIn returns the WebHookConfigAuthApiKeyIn field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyIn() string {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyIn) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyIn
}

// GetWebHookConfigAuthApiKeyInOk returns a tuple with the WebHookConfigAuthApiKeyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyInOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyIn) {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyIn, true
}

// HasWebHookConfigAuthApiKeyIn returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthApiKeyIn() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthApiKeyIn) {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyIn gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyIn field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthApiKeyIn(v string) {
	o.WebHookConfigAuthApiKeyIn = &v
}

// GetWebHookConfigAuthApiKeyName returns the WebHookConfigAuthApiKeyName field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyName() string {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyName) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyName
}

// GetWebHookConfigAuthApiKeyNameOk returns a tuple with the WebHookConfigAuthApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyName) {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyName, true
}

// HasWebHookConfigAuthApiKeyName returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthApiKeyName() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthApiKeyName) {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyName gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyName field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthApiKeyName(v string) {
	o.WebHookConfigAuthApiKeyName = &v
}

// GetWebHookConfigAuthApiKeyValue returns the WebHookConfigAuthApiKeyValue field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyValue() string {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyValue) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthApiKeyValue
}

// GetWebHookConfigAuthApiKeyValueOk returns a tuple with the WebHookConfigAuthApiKeyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthApiKeyValueOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthApiKeyValue) {
		return nil, false
	}
	return o.WebHookConfigAuthApiKeyValue, true
}

// HasWebHookConfigAuthApiKeyValue returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthApiKeyValue() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthApiKeyValue) {
		return true
	}

	return false
}

// SetWebHookConfigAuthApiKeyValue gets a reference to the given string and assigns it to the WebHookConfigAuthApiKeyValue field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthApiKeyValue(v string) {
	o.WebHookConfigAuthApiKeyValue = &v
}

// GetWebHookConfigAuthBasicAuthPassword returns the WebHookConfigAuthBasicAuthPassword field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthBasicAuthPassword() string {
	if o == nil || IsNil(o.WebHookConfigAuthBasicAuthPassword) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthBasicAuthPassword
}

// GetWebHookConfigAuthBasicAuthPasswordOk returns a tuple with the WebHookConfigAuthBasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthBasicAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthBasicAuthPassword) {
		return nil, false
	}
	return o.WebHookConfigAuthBasicAuthPassword, true
}

// HasWebHookConfigAuthBasicAuthPassword returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthBasicAuthPassword() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthBasicAuthPassword) {
		return true
	}

	return false
}

// SetWebHookConfigAuthBasicAuthPassword gets a reference to the given string and assigns it to the WebHookConfigAuthBasicAuthPassword field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthBasicAuthPassword(v string) {
	o.WebHookConfigAuthBasicAuthPassword = &v
}

// GetWebHookConfigAuthBasicAuthUser returns the WebHookConfigAuthBasicAuthUser field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthBasicAuthUser() string {
	if o == nil || IsNil(o.WebHookConfigAuthBasicAuthUser) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthBasicAuthUser
}

// GetWebHookConfigAuthBasicAuthUserOk returns a tuple with the WebHookConfigAuthBasicAuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthBasicAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthBasicAuthUser) {
		return nil, false
	}
	return o.WebHookConfigAuthBasicAuthUser, true
}

// HasWebHookConfigAuthBasicAuthUser returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthBasicAuthUser() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthBasicAuthUser) {
		return true
	}

	return false
}

// SetWebHookConfigAuthBasicAuthUser gets a reference to the given string and assigns it to the WebHookConfigAuthBasicAuthUser field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthBasicAuthUser(v string) {
	o.WebHookConfigAuthBasicAuthUser = &v
}

// GetWebHookConfigAuthType returns the WebHookConfigAuthType field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthType() string {
	if o == nil || IsNil(o.WebHookConfigAuthType) {
		var ret string
		return ret
	}
	return *o.WebHookConfigAuthType
}

// GetWebHookConfigAuthTypeOk returns a tuple with the WebHookConfigAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigAuthType) {
		return nil, false
	}
	return o.WebHookConfigAuthType, true
}

// HasWebHookConfigAuthType returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigAuthType() bool {
	if o != nil && !IsNil(o.WebHookConfigAuthType) {
		return true
	}

	return false
}

// SetWebHookConfigAuthType gets a reference to the given string and assigns it to the WebHookConfigAuthType field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigAuthType(v string) {
	o.WebHookConfigAuthType = &v
}

// GetWebHookConfigBody returns the WebHookConfigBody field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigBody() string {
	if o == nil || IsNil(o.WebHookConfigBody) {
		var ret string
		return ret
	}
	return *o.WebHookConfigBody
}

// GetWebHookConfigBodyOk returns a tuple with the WebHookConfigBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigBodyOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigBody) {
		return nil, false
	}
	return o.WebHookConfigBody, true
}

// HasWebHookConfigBody returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigBody() bool {
	if o != nil && !IsNil(o.WebHookConfigBody) {
		return true
	}

	return false
}

// SetWebHookConfigBody gets a reference to the given string and assigns it to the WebHookConfigBody field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigBody(v string) {
	o.WebHookConfigBody = &v
}

// GetWebHookConfigCanInterrupt returns the WebHookConfigCanInterrupt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigCanInterrupt() bool {
	if o == nil || IsNil(o.WebHookConfigCanInterrupt) {
		var ret bool
		return ret
	}
	return *o.WebHookConfigCanInterrupt
}

// GetWebHookConfigCanInterruptOk returns a tuple with the WebHookConfigCanInterrupt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigCanInterruptOk() (*bool, bool) {
	if o == nil || IsNil(o.WebHookConfigCanInterrupt) {
		return nil, false
	}
	return o.WebHookConfigCanInterrupt, true
}

// HasWebHookConfigCanInterrupt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigCanInterrupt() bool {
	if o != nil && !IsNil(o.WebHookConfigCanInterrupt) {
		return true
	}

	return false
}

// SetWebHookConfigCanInterrupt gets a reference to the given bool and assigns it to the WebHookConfigCanInterrupt field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigCanInterrupt(v bool) {
	o.WebHookConfigCanInterrupt = &v
}

// GetWebHookConfigMethod returns the WebHookConfigMethod field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigMethod() string {
	if o == nil || IsNil(o.WebHookConfigMethod) {
		var ret string
		return ret
	}
	return *o.WebHookConfigMethod
}

// GetWebHookConfigMethodOk returns a tuple with the WebHookConfigMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigMethodOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigMethod) {
		return nil, false
	}
	return o.WebHookConfigMethod, true
}

// HasWebHookConfigMethod returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigMethod() bool {
	if o != nil && !IsNil(o.WebHookConfigMethod) {
		return true
	}

	return false
}

// SetWebHookConfigMethod gets a reference to the given string and assigns it to the WebHookConfigMethod field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigMethod(v string) {
	o.WebHookConfigMethod = &v
}

// GetWebHookConfigResponseIgnore returns the WebHookConfigResponseIgnore field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigResponseIgnore() bool {
	if o == nil || IsNil(o.WebHookConfigResponseIgnore) {
		var ret bool
		return ret
	}
	return *o.WebHookConfigResponseIgnore
}

// GetWebHookConfigResponseIgnoreOk returns a tuple with the WebHookConfigResponseIgnore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigResponseIgnoreOk() (*bool, bool) {
	if o == nil || IsNil(o.WebHookConfigResponseIgnore) {
		return nil, false
	}
	return o.WebHookConfigResponseIgnore, true
}

// HasWebHookConfigResponseIgnore returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigResponseIgnore() bool {
	if o != nil && !IsNil(o.WebHookConfigResponseIgnore) {
		return true
	}

	return false
}

// SetWebHookConfigResponseIgnore gets a reference to the given bool and assigns it to the WebHookConfigResponseIgnore field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigResponseIgnore(v bool) {
	o.WebHookConfigResponseIgnore = &v
}

// GetWebHookConfigResponseParse returns the WebHookConfigResponseParse field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigResponseParse() bool {
	if o == nil || IsNil(o.WebHookConfigResponseParse) {
		var ret bool
		return ret
	}
	return *o.WebHookConfigResponseParse
}

// GetWebHookConfigResponseParseOk returns a tuple with the WebHookConfigResponseParse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigResponseParseOk() (*bool, bool) {
	if o == nil || IsNil(o.WebHookConfigResponseParse) {
		return nil, false
	}
	return o.WebHookConfigResponseParse, true
}

// HasWebHookConfigResponseParse returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigResponseParse() bool {
	if o != nil && !IsNil(o.WebHookConfigResponseParse) {
		return true
	}

	return false
}

// SetWebHookConfigResponseParse gets a reference to the given bool and assigns it to the WebHookConfigResponseParse field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigResponseParse(v bool) {
	o.WebHookConfigResponseParse = &v
}

// GetWebHookConfigUrl returns the WebHookConfigUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigUrl() string {
	if o == nil || IsNil(o.WebHookConfigUrl) {
		var ret string
		return ret
	}
	return *o.WebHookConfigUrl
}

// GetWebHookConfigUrlOk returns a tuple with the WebHookConfigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionHook) GetWebHookConfigUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebHookConfigUrl) {
		return nil, false
	}
	return o.WebHookConfigUrl, true
}

// HasWebHookConfigUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionHook) HasWebHookConfigUrl() bool {
	if o != nil && !IsNil(o.WebHookConfigUrl) {
		return true
	}

	return false
}

// SetWebHookConfigUrl gets a reference to the given string and assigns it to the WebHookConfigUrl field.
func (o *NormalizedProjectRevisionHook) SetWebHookConfigUrl(v string) {
	o.WebHookConfigUrl = &v
}

func (o NormalizedProjectRevisionHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NormalizedProjectRevisionHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config_key"] = o.ConfigKey
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["hook"] = o.Hook
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectRevisionId) {
		toSerialize["project_revision_id"] = o.ProjectRevisionId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.WebHookConfigAuthApiKeyIn) {
		toSerialize["web_hook_config_auth_api_key_in"] = o.WebHookConfigAuthApiKeyIn
	}
	if !IsNil(o.WebHookConfigAuthApiKeyName) {
		toSerialize["web_hook_config_auth_api_key_name"] = o.WebHookConfigAuthApiKeyName
	}
	if !IsNil(o.WebHookConfigAuthApiKeyValue) {
		toSerialize["web_hook_config_auth_api_key_value"] = o.WebHookConfigAuthApiKeyValue
	}
	if !IsNil(o.WebHookConfigAuthBasicAuthPassword) {
		toSerialize["web_hook_config_auth_basic_auth_password"] = o.WebHookConfigAuthBasicAuthPassword
	}
	if !IsNil(o.WebHookConfigAuthBasicAuthUser) {
		toSerialize["web_hook_config_auth_basic_auth_user"] = o.WebHookConfigAuthBasicAuthUser
	}
	if !IsNil(o.WebHookConfigAuthType) {
		toSerialize["web_hook_config_auth_type"] = o.WebHookConfigAuthType
	}
	if !IsNil(o.WebHookConfigBody) {
		toSerialize["web_hook_config_body"] = o.WebHookConfigBody
	}
	if !IsNil(o.WebHookConfigCanInterrupt) {
		toSerialize["web_hook_config_can_interrupt"] = o.WebHookConfigCanInterrupt
	}
	if !IsNil(o.WebHookConfigMethod) {
		toSerialize["web_hook_config_method"] = o.WebHookConfigMethod
	}
	if !IsNil(o.WebHookConfigResponseIgnore) {
		toSerialize["web_hook_config_response_ignore"] = o.WebHookConfigResponseIgnore
	}
	if !IsNil(o.WebHookConfigResponseParse) {
		toSerialize["web_hook_config_response_parse"] = o.WebHookConfigResponseParse
	}
	if !IsNil(o.WebHookConfigUrl) {
		toSerialize["web_hook_config_url"] = o.WebHookConfigUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NormalizedProjectRevisionHook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config_key",
		"hook",
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

	varNormalizedProjectRevisionHook := _NormalizedProjectRevisionHook{}

	err = json.Unmarshal(data, &varNormalizedProjectRevisionHook)

	if err != nil {
		return err
	}

	*o = NormalizedProjectRevisionHook(varNormalizedProjectRevisionHook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config_key")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "hook")
		delete(additionalProperties, "id")
		delete(additionalProperties, "project_revision_id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "web_hook_config_auth_api_key_in")
		delete(additionalProperties, "web_hook_config_auth_api_key_name")
		delete(additionalProperties, "web_hook_config_auth_api_key_value")
		delete(additionalProperties, "web_hook_config_auth_basic_auth_password")
		delete(additionalProperties, "web_hook_config_auth_basic_auth_user")
		delete(additionalProperties, "web_hook_config_auth_type")
		delete(additionalProperties, "web_hook_config_body")
		delete(additionalProperties, "web_hook_config_can_interrupt")
		delete(additionalProperties, "web_hook_config_method")
		delete(additionalProperties, "web_hook_config_response_ignore")
		delete(additionalProperties, "web_hook_config_response_parse")
		delete(additionalProperties, "web_hook_config_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNormalizedProjectRevisionHook struct {
	value *NormalizedProjectRevisionHook
	isSet bool
}

func (v NullableNormalizedProjectRevisionHook) Get() *NormalizedProjectRevisionHook {
	return v.value
}

func (v *NullableNormalizedProjectRevisionHook) Set(val *NormalizedProjectRevisionHook) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionHook) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionHook(val *NormalizedProjectRevisionHook) *NullableNormalizedProjectRevisionHook {
	return &NullableNormalizedProjectRevisionHook{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


