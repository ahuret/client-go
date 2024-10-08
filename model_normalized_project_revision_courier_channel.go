/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.6
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the NormalizedProjectRevisionCourierChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NormalizedProjectRevisionCourierChannel{}

// NormalizedProjectRevisionCourierChannel struct for NormalizedProjectRevisionCourierChannel
type NormalizedProjectRevisionCourierChannel struct {
	// The Channel's public ID
	ChannelId string `json:"channel_id"`
	// The creation date
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// API key location  Can either be \"header\" or \"query\"
	RequestConfigAuthConfigApiKeyIn *string `json:"request_config_auth_config_api_key_in,omitempty"`
	// API key name  Only used if the auth type is api_key
	RequestConfigAuthConfigApiKeyName *string `json:"request_config_auth_config_api_key_name,omitempty"`
	// API key value  Only used if the auth type is api_key
	RequestConfigAuthConfigApiKeyValue *string `json:"request_config_auth_config_api_key_value,omitempty"`
	// Basic Auth Password  Only used if the auth type is basic_auth
	RequestConfigAuthConfigBasicAuthPassword *string `json:"request_config_auth_config_basic_auth_password,omitempty"`
	// Basic Auth Username  Only used if the auth type is basic_auth
	RequestConfigAuthConfigBasicAuthUser *string `json:"request_config_auth_config_basic_auth_user,omitempty"`
	// HTTP Auth Method to use for the HTTP call  Can either be basic_auth or api_key basic_auth CourierChannelAuthTypeBasicAuth api_key CourierChannelAuthTypeApiKey
	RequestConfigAuthType *string `json:"request_config_auth_type,omitempty"`
	// URI pointing to the JsonNet template used for HTTP body payload generation.
	RequestConfigBody string `json:"request_config_body"`
	// NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
	RequestConfigHeaders map[string]interface{} `json:"request_config_headers,omitempty"`
	// The HTTP method to use (GET, POST, etc) for the HTTP call
	RequestConfigMethod string `json:"request_config_method"`
	RequestConfigUrl *string `json:"request_config_url,omitempty"`
	// Last upate time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NormalizedProjectRevisionCourierChannel NormalizedProjectRevisionCourierChannel

// NewNormalizedProjectRevisionCourierChannel instantiates a new NormalizedProjectRevisionCourierChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNormalizedProjectRevisionCourierChannel(channelId string, requestConfigBody string, requestConfigMethod string) *NormalizedProjectRevisionCourierChannel {
	this := NormalizedProjectRevisionCourierChannel{}
	this.ChannelId = channelId
	this.RequestConfigBody = requestConfigBody
	this.RequestConfigMethod = requestConfigMethod
	return &this
}

// NewNormalizedProjectRevisionCourierChannelWithDefaults instantiates a new NormalizedProjectRevisionCourierChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNormalizedProjectRevisionCourierChannelWithDefaults() *NormalizedProjectRevisionCourierChannel {
	this := NormalizedProjectRevisionCourierChannel{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *NormalizedProjectRevisionCourierChannel) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *NormalizedProjectRevisionCourierChannel) SetChannelId(v string) {
	o.ChannelId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NormalizedProjectRevisionCourierChannel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetRequestConfigAuthConfigApiKeyIn returns the RequestConfigAuthConfigApiKeyIn field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyIn() string {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyIn) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthConfigApiKeyIn
}

// GetRequestConfigAuthConfigApiKeyInOk returns a tuple with the RequestConfigAuthConfigApiKeyIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyInOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyIn) {
		return nil, false
	}
	return o.RequestConfigAuthConfigApiKeyIn, true
}

// HasRequestConfigAuthConfigApiKeyIn returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthConfigApiKeyIn() bool {
	if o != nil && !IsNil(o.RequestConfigAuthConfigApiKeyIn) {
		return true
	}

	return false
}

// SetRequestConfigAuthConfigApiKeyIn gets a reference to the given string and assigns it to the RequestConfigAuthConfigApiKeyIn field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthConfigApiKeyIn(v string) {
	o.RequestConfigAuthConfigApiKeyIn = &v
}

// GetRequestConfigAuthConfigApiKeyName returns the RequestConfigAuthConfigApiKeyName field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyName() string {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyName) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthConfigApiKeyName
}

// GetRequestConfigAuthConfigApiKeyNameOk returns a tuple with the RequestConfigAuthConfigApiKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyName) {
		return nil, false
	}
	return o.RequestConfigAuthConfigApiKeyName, true
}

// HasRequestConfigAuthConfigApiKeyName returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthConfigApiKeyName() bool {
	if o != nil && !IsNil(o.RequestConfigAuthConfigApiKeyName) {
		return true
	}

	return false
}

// SetRequestConfigAuthConfigApiKeyName gets a reference to the given string and assigns it to the RequestConfigAuthConfigApiKeyName field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthConfigApiKeyName(v string) {
	o.RequestConfigAuthConfigApiKeyName = &v
}

// GetRequestConfigAuthConfigApiKeyValue returns the RequestConfigAuthConfigApiKeyValue field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyValue() string {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyValue) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthConfigApiKeyValue
}

// GetRequestConfigAuthConfigApiKeyValueOk returns a tuple with the RequestConfigAuthConfigApiKeyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigApiKeyValueOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthConfigApiKeyValue) {
		return nil, false
	}
	return o.RequestConfigAuthConfigApiKeyValue, true
}

// HasRequestConfigAuthConfigApiKeyValue returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthConfigApiKeyValue() bool {
	if o != nil && !IsNil(o.RequestConfigAuthConfigApiKeyValue) {
		return true
	}

	return false
}

// SetRequestConfigAuthConfigApiKeyValue gets a reference to the given string and assigns it to the RequestConfigAuthConfigApiKeyValue field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthConfigApiKeyValue(v string) {
	o.RequestConfigAuthConfigApiKeyValue = &v
}

// GetRequestConfigAuthConfigBasicAuthPassword returns the RequestConfigAuthConfigBasicAuthPassword field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigBasicAuthPassword() string {
	if o == nil || IsNil(o.RequestConfigAuthConfigBasicAuthPassword) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthConfigBasicAuthPassword
}

// GetRequestConfigAuthConfigBasicAuthPasswordOk returns a tuple with the RequestConfigAuthConfigBasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigBasicAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthConfigBasicAuthPassword) {
		return nil, false
	}
	return o.RequestConfigAuthConfigBasicAuthPassword, true
}

// HasRequestConfigAuthConfigBasicAuthPassword returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthConfigBasicAuthPassword() bool {
	if o != nil && !IsNil(o.RequestConfigAuthConfigBasicAuthPassword) {
		return true
	}

	return false
}

// SetRequestConfigAuthConfigBasicAuthPassword gets a reference to the given string and assigns it to the RequestConfigAuthConfigBasicAuthPassword field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthConfigBasicAuthPassword(v string) {
	o.RequestConfigAuthConfigBasicAuthPassword = &v
}

// GetRequestConfigAuthConfigBasicAuthUser returns the RequestConfigAuthConfigBasicAuthUser field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigBasicAuthUser() string {
	if o == nil || IsNil(o.RequestConfigAuthConfigBasicAuthUser) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthConfigBasicAuthUser
}

// GetRequestConfigAuthConfigBasicAuthUserOk returns a tuple with the RequestConfigAuthConfigBasicAuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthConfigBasicAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthConfigBasicAuthUser) {
		return nil, false
	}
	return o.RequestConfigAuthConfigBasicAuthUser, true
}

// HasRequestConfigAuthConfigBasicAuthUser returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthConfigBasicAuthUser() bool {
	if o != nil && !IsNil(o.RequestConfigAuthConfigBasicAuthUser) {
		return true
	}

	return false
}

// SetRequestConfigAuthConfigBasicAuthUser gets a reference to the given string and assigns it to the RequestConfigAuthConfigBasicAuthUser field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthConfigBasicAuthUser(v string) {
	o.RequestConfigAuthConfigBasicAuthUser = &v
}

// GetRequestConfigAuthType returns the RequestConfigAuthType field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthType() string {
	if o == nil || IsNil(o.RequestConfigAuthType) {
		var ret string
		return ret
	}
	return *o.RequestConfigAuthType
}

// GetRequestConfigAuthTypeOk returns a tuple with the RequestConfigAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigAuthType) {
		return nil, false
	}
	return o.RequestConfigAuthType, true
}

// HasRequestConfigAuthType returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigAuthType() bool {
	if o != nil && !IsNil(o.RequestConfigAuthType) {
		return true
	}

	return false
}

// SetRequestConfigAuthType gets a reference to the given string and assigns it to the RequestConfigAuthType field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigAuthType(v string) {
	o.RequestConfigAuthType = &v
}

// GetRequestConfigBody returns the RequestConfigBody field value
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestConfigBody
}

// GetRequestConfigBodyOk returns a tuple with the RequestConfigBody field value
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestConfigBody, true
}

// SetRequestConfigBody sets field value
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigBody(v string) {
	o.RequestConfigBody = v
}

// GetRequestConfigHeaders returns the RequestConfigHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigHeaders() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RequestConfigHeaders
}

// GetRequestConfigHeadersOk returns a tuple with the RequestConfigHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigHeadersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RequestConfigHeaders) {
		return map[string]interface{}{}, false
	}
	return o.RequestConfigHeaders, true
}

// HasRequestConfigHeaders returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigHeaders() bool {
	if o != nil && !IsNil(o.RequestConfigHeaders) {
		return true
	}

	return false
}

// SetRequestConfigHeaders gets a reference to the given map[string]interface{} and assigns it to the RequestConfigHeaders field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigHeaders(v map[string]interface{}) {
	o.RequestConfigHeaders = v
}

// GetRequestConfigMethod returns the RequestConfigMethod field value
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestConfigMethod
}

// GetRequestConfigMethodOk returns a tuple with the RequestConfigMethod field value
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestConfigMethod, true
}

// SetRequestConfigMethod sets field value
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigMethod(v string) {
	o.RequestConfigMethod = v
}

// GetRequestConfigUrl returns the RequestConfigUrl field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigUrl() string {
	if o == nil || IsNil(o.RequestConfigUrl) {
		var ret string
		return ret
	}
	return *o.RequestConfigUrl
}

// GetRequestConfigUrlOk returns a tuple with the RequestConfigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetRequestConfigUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RequestConfigUrl) {
		return nil, false
	}
	return o.RequestConfigUrl, true
}

// HasRequestConfigUrl returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasRequestConfigUrl() bool {
	if o != nil && !IsNil(o.RequestConfigUrl) {
		return true
	}

	return false
}

// SetRequestConfigUrl gets a reference to the given string and assigns it to the RequestConfigUrl field.
func (o *NormalizedProjectRevisionCourierChannel) SetRequestConfigUrl(v string) {
	o.RequestConfigUrl = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NormalizedProjectRevisionCourierChannel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NormalizedProjectRevisionCourierChannel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NormalizedProjectRevisionCourierChannel) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NormalizedProjectRevisionCourierChannel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o NormalizedProjectRevisionCourierChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NormalizedProjectRevisionCourierChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_id"] = o.ChannelId
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.RequestConfigAuthConfigApiKeyIn) {
		toSerialize["request_config_auth_config_api_key_in"] = o.RequestConfigAuthConfigApiKeyIn
	}
	if !IsNil(o.RequestConfigAuthConfigApiKeyName) {
		toSerialize["request_config_auth_config_api_key_name"] = o.RequestConfigAuthConfigApiKeyName
	}
	if !IsNil(o.RequestConfigAuthConfigApiKeyValue) {
		toSerialize["request_config_auth_config_api_key_value"] = o.RequestConfigAuthConfigApiKeyValue
	}
	if !IsNil(o.RequestConfigAuthConfigBasicAuthPassword) {
		toSerialize["request_config_auth_config_basic_auth_password"] = o.RequestConfigAuthConfigBasicAuthPassword
	}
	if !IsNil(o.RequestConfigAuthConfigBasicAuthUser) {
		toSerialize["request_config_auth_config_basic_auth_user"] = o.RequestConfigAuthConfigBasicAuthUser
	}
	if !IsNil(o.RequestConfigAuthType) {
		toSerialize["request_config_auth_type"] = o.RequestConfigAuthType
	}
	toSerialize["request_config_body"] = o.RequestConfigBody
	if o.RequestConfigHeaders != nil {
		toSerialize["request_config_headers"] = o.RequestConfigHeaders
	}
	toSerialize["request_config_method"] = o.RequestConfigMethod
	if !IsNil(o.RequestConfigUrl) {
		toSerialize["request_config_url"] = o.RequestConfigUrl
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NormalizedProjectRevisionCourierChannel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
		"request_config_body",
		"request_config_method",
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

	varNormalizedProjectRevisionCourierChannel := _NormalizedProjectRevisionCourierChannel{}

	err = json.Unmarshal(data, &varNormalizedProjectRevisionCourierChannel)

	if err != nil {
		return err
	}

	*o = NormalizedProjectRevisionCourierChannel(varNormalizedProjectRevisionCourierChannel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "request_config_auth_config_api_key_in")
		delete(additionalProperties, "request_config_auth_config_api_key_name")
		delete(additionalProperties, "request_config_auth_config_api_key_value")
		delete(additionalProperties, "request_config_auth_config_basic_auth_password")
		delete(additionalProperties, "request_config_auth_config_basic_auth_user")
		delete(additionalProperties, "request_config_auth_type")
		delete(additionalProperties, "request_config_body")
		delete(additionalProperties, "request_config_headers")
		delete(additionalProperties, "request_config_method")
		delete(additionalProperties, "request_config_url")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNormalizedProjectRevisionCourierChannel struct {
	value *NormalizedProjectRevisionCourierChannel
	isSet bool
}

func (v NullableNormalizedProjectRevisionCourierChannel) Get() *NormalizedProjectRevisionCourierChannel {
	return v.value
}

func (v *NullableNormalizedProjectRevisionCourierChannel) Set(val *NormalizedProjectRevisionCourierChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableNormalizedProjectRevisionCourierChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableNormalizedProjectRevisionCourierChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNormalizedProjectRevisionCourierChannel(val *NormalizedProjectRevisionCourierChannel) *NullableNormalizedProjectRevisionCourierChannel {
	return &NullableNormalizedProjectRevisionCourierChannel{value: val, isSet: true}
}

func (v NullableNormalizedProjectRevisionCourierChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNormalizedProjectRevisionCourierChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


