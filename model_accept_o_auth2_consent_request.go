/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.13.4
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the AcceptOAuth2ConsentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptOAuth2ConsentRequest{}

// AcceptOAuth2ConsentRequest struct for AcceptOAuth2ConsentRequest
type AcceptOAuth2ConsentRequest struct {
	Context map[string]interface{} `json:"context,omitempty"`
	GrantAccessTokenAudience []string `json:"grant_access_token_audience,omitempty"`
	GrantScope []string `json:"grant_scope,omitempty"`
	HandledAt *time.Time `json:"handled_at,omitempty"`
	// Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same client asks the same user for the same, or a subset of, scope.
	Remember *bool `json:"remember,omitempty"`
	// RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the authorization will be remembered indefinitely.
	RememberFor *int64 `json:"remember_for,omitempty"`
	Session *AcceptOAuth2ConsentRequestSession `json:"session,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcceptOAuth2ConsentRequest AcceptOAuth2ConsentRequest

// NewAcceptOAuth2ConsentRequest instantiates a new AcceptOAuth2ConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptOAuth2ConsentRequest() *AcceptOAuth2ConsentRequest {
	this := AcceptOAuth2ConsentRequest{}
	return &this
}

// NewAcceptOAuth2ConsentRequestWithDefaults instantiates a new AcceptOAuth2ConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptOAuth2ConsentRequestWithDefaults() *AcceptOAuth2ConsentRequest {
	this := AcceptOAuth2ConsentRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetContext() map[string]interface{} {
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *AcceptOAuth2ConsentRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetGrantAccessTokenAudience returns the GrantAccessTokenAudience field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetGrantAccessTokenAudience() []string {
	if o == nil || IsNil(o.GrantAccessTokenAudience) {
		var ret []string
		return ret
	}
	return o.GrantAccessTokenAudience
}

// GetGrantAccessTokenAudienceOk returns a tuple with the GrantAccessTokenAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetGrantAccessTokenAudienceOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantAccessTokenAudience) {
		return nil, false
	}
	return o.GrantAccessTokenAudience, true
}

// HasGrantAccessTokenAudience returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasGrantAccessTokenAudience() bool {
	if o != nil && !IsNil(o.GrantAccessTokenAudience) {
		return true
	}

	return false
}

// SetGrantAccessTokenAudience gets a reference to the given []string and assigns it to the GrantAccessTokenAudience field.
func (o *AcceptOAuth2ConsentRequest) SetGrantAccessTokenAudience(v []string) {
	o.GrantAccessTokenAudience = v
}

// GetGrantScope returns the GrantScope field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetGrantScope() []string {
	if o == nil || IsNil(o.GrantScope) {
		var ret []string
		return ret
	}
	return o.GrantScope
}

// GetGrantScopeOk returns a tuple with the GrantScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetGrantScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantScope) {
		return nil, false
	}
	return o.GrantScope, true
}

// HasGrantScope returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasGrantScope() bool {
	if o != nil && !IsNil(o.GrantScope) {
		return true
	}

	return false
}

// SetGrantScope gets a reference to the given []string and assigns it to the GrantScope field.
func (o *AcceptOAuth2ConsentRequest) SetGrantScope(v []string) {
	o.GrantScope = v
}

// GetHandledAt returns the HandledAt field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetHandledAt() time.Time {
	if o == nil || IsNil(o.HandledAt) {
		var ret time.Time
		return ret
	}
	return *o.HandledAt
}

// GetHandledAtOk returns a tuple with the HandledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetHandledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.HandledAt) {
		return nil, false
	}
	return o.HandledAt, true
}

// HasHandledAt returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasHandledAt() bool {
	if o != nil && !IsNil(o.HandledAt) {
		return true
	}

	return false
}

// SetHandledAt gets a reference to the given time.Time and assigns it to the HandledAt field.
func (o *AcceptOAuth2ConsentRequest) SetHandledAt(v time.Time) {
	o.HandledAt = &v
}

// GetRemember returns the Remember field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetRemember() bool {
	if o == nil || IsNil(o.Remember) {
		var ret bool
		return ret
	}
	return *o.Remember
}

// GetRememberOk returns a tuple with the Remember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetRememberOk() (*bool, bool) {
	if o == nil || IsNil(o.Remember) {
		return nil, false
	}
	return o.Remember, true
}

// HasRemember returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasRemember() bool {
	if o != nil && !IsNil(o.Remember) {
		return true
	}

	return false
}

// SetRemember gets a reference to the given bool and assigns it to the Remember field.
func (o *AcceptOAuth2ConsentRequest) SetRemember(v bool) {
	o.Remember = &v
}

// GetRememberFor returns the RememberFor field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetRememberFor() int64 {
	if o == nil || IsNil(o.RememberFor) {
		var ret int64
		return ret
	}
	return *o.RememberFor
}

// GetRememberForOk returns a tuple with the RememberFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetRememberForOk() (*int64, bool) {
	if o == nil || IsNil(o.RememberFor) {
		return nil, false
	}
	return o.RememberFor, true
}

// HasRememberFor returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasRememberFor() bool {
	if o != nil && !IsNil(o.RememberFor) {
		return true
	}

	return false
}

// SetRememberFor gets a reference to the given int64 and assigns it to the RememberFor field.
func (o *AcceptOAuth2ConsentRequest) SetRememberFor(v int64) {
	o.RememberFor = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *AcceptOAuth2ConsentRequest) GetSession() AcceptOAuth2ConsentRequestSession {
	if o == nil || IsNil(o.Session) {
		var ret AcceptOAuth2ConsentRequestSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcceptOAuth2ConsentRequest) GetSessionOk() (*AcceptOAuth2ConsentRequestSession, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *AcceptOAuth2ConsentRequest) HasSession() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given AcceptOAuth2ConsentRequestSession and assigns it to the Session field.
func (o *AcceptOAuth2ConsentRequest) SetSession(v AcceptOAuth2ConsentRequestSession) {
	o.Session = &v
}

func (o AcceptOAuth2ConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptOAuth2ConsentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.GrantAccessTokenAudience) {
		toSerialize["grant_access_token_audience"] = o.GrantAccessTokenAudience
	}
	if !IsNil(o.GrantScope) {
		toSerialize["grant_scope"] = o.GrantScope
	}
	if !IsNil(o.HandledAt) {
		toSerialize["handled_at"] = o.HandledAt
	}
	if !IsNil(o.Remember) {
		toSerialize["remember"] = o.Remember
	}
	if !IsNil(o.RememberFor) {
		toSerialize["remember_for"] = o.RememberFor
	}
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcceptOAuth2ConsentRequest) UnmarshalJSON(data []byte) (err error) {
	varAcceptOAuth2ConsentRequest := _AcceptOAuth2ConsentRequest{}

	err = json.Unmarshal(data, &varAcceptOAuth2ConsentRequest)

	if err != nil {
		return err
	}

	*o = AcceptOAuth2ConsentRequest(varAcceptOAuth2ConsentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "grant_access_token_audience")
		delete(additionalProperties, "grant_scope")
		delete(additionalProperties, "handled_at")
		delete(additionalProperties, "remember")
		delete(additionalProperties, "remember_for")
		delete(additionalProperties, "session")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptOAuth2ConsentRequest struct {
	value *AcceptOAuth2ConsentRequest
	isSet bool
}

func (v NullableAcceptOAuth2ConsentRequest) Get() *AcceptOAuth2ConsentRequest {
	return v.value
}

func (v *NullableAcceptOAuth2ConsentRequest) Set(val *AcceptOAuth2ConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptOAuth2ConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptOAuth2ConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptOAuth2ConsentRequest(val *AcceptOAuth2ConsentRequest) *NullableAcceptOAuth2ConsentRequest {
	return &NullableAcceptOAuth2ConsentRequest{value: val, isSet: true}
}

func (v NullableAcceptOAuth2ConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptOAuth2ConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


