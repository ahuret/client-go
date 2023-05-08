/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.27
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateSettingsFlowWithOidcMethod Update Settings Flow with OpenID Connect Method
type UpdateSettingsFlowWithOidcMethod struct {
	// Flow ID is the flow's ID.  in: query
	Flow *string `json:"flow,omitempty"`
	// Link this provider  Either this or `unlink` must be set.  type: string in: body
	Link *string `json:"link,omitempty"`
	// Method  Should be set to profile when trying to update a profile.
	Method string `json:"method"`
	// The identity's traits  in: body
	Traits map[string]interface{} `json:"traits,omitempty"`
	// Unlink this provider  Either this or `link` must be set.  type: string in: body
	Unlink *string `json:"unlink,omitempty"`
	// UpstreamParameters are the parameters that are passed to the upstream identity provider.  These parameters are optional and depend on what the upstream identity provider supports. Supported parameters are: `login_hint` (string): The `login_hint` parameter suppresses the account chooser and either pre-fills the email box on the sign-in form, or selects the proper session. `hd` (string): The `hd` parameter limits the login/registration process to a Google Organization, e.g. `mycollege.edu`.
	UpstreamParameters map[string]interface{} `json:"upstream_parameters,omitempty"`
}

// NewUpdateSettingsFlowWithOidcMethod instantiates a new UpdateSettingsFlowWithOidcMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingsFlowWithOidcMethod(method string) *UpdateSettingsFlowWithOidcMethod {
	this := UpdateSettingsFlowWithOidcMethod{}
	this.Method = method
	return &this
}

// NewUpdateSettingsFlowWithOidcMethodWithDefaults instantiates a new UpdateSettingsFlowWithOidcMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingsFlowWithOidcMethodWithDefaults() *UpdateSettingsFlowWithOidcMethod {
	this := UpdateSettingsFlowWithOidcMethod{}
	return &this
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithOidcMethod) GetFlow() string {
	if o == nil || o.Flow == nil {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetFlowOk() (*string, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithOidcMethod) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *UpdateSettingsFlowWithOidcMethod) SetFlow(v string) {
	o.Flow = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithOidcMethod) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithOidcMethod) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UpdateSettingsFlowWithOidcMethod) SetLink(v string) {
	o.Link = &v
}

// GetMethod returns the Method field value
func (o *UpdateSettingsFlowWithOidcMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateSettingsFlowWithOidcMethod) SetMethod(v string) {
	o.Method = v
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithOidcMethod) GetTraits() map[string]interface{} {
	if o == nil || o.Traits == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || o.Traits == nil {
		return nil, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithOidcMethod) HasTraits() bool {
	if o != nil && o.Traits != nil {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *UpdateSettingsFlowWithOidcMethod) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetUnlink returns the Unlink field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithOidcMethod) GetUnlink() string {
	if o == nil || o.Unlink == nil {
		var ret string
		return ret
	}
	return *o.Unlink
}

// GetUnlinkOk returns a tuple with the Unlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetUnlinkOk() (*string, bool) {
	if o == nil || o.Unlink == nil {
		return nil, false
	}
	return o.Unlink, true
}

// HasUnlink returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithOidcMethod) HasUnlink() bool {
	if o != nil && o.Unlink != nil {
		return true
	}

	return false
}

// SetUnlink gets a reference to the given string and assigns it to the Unlink field.
func (o *UpdateSettingsFlowWithOidcMethod) SetUnlink(v string) {
	o.Unlink = &v
}

// GetUpstreamParameters returns the UpstreamParameters field value if set, zero value otherwise.
func (o *UpdateSettingsFlowWithOidcMethod) GetUpstreamParameters() map[string]interface{} {
	if o == nil || o.UpstreamParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UpstreamParameters
}

// GetUpstreamParametersOk returns a tuple with the UpstreamParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingsFlowWithOidcMethod) GetUpstreamParametersOk() (map[string]interface{}, bool) {
	if o == nil || o.UpstreamParameters == nil {
		return nil, false
	}
	return o.UpstreamParameters, true
}

// HasUpstreamParameters returns a boolean if a field has been set.
func (o *UpdateSettingsFlowWithOidcMethod) HasUpstreamParameters() bool {
	if o != nil && o.UpstreamParameters != nil {
		return true
	}

	return false
}

// SetUpstreamParameters gets a reference to the given map[string]interface{} and assigns it to the UpstreamParameters field.
func (o *UpdateSettingsFlowWithOidcMethod) SetUpstreamParameters(v map[string]interface{}) {
	o.UpstreamParameters = v
}

func (o UpdateSettingsFlowWithOidcMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Traits != nil {
		toSerialize["traits"] = o.Traits
	}
	if o.Unlink != nil {
		toSerialize["unlink"] = o.Unlink
	}
	if o.UpstreamParameters != nil {
		toSerialize["upstream_parameters"] = o.UpstreamParameters
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSettingsFlowWithOidcMethod struct {
	value *UpdateSettingsFlowWithOidcMethod
	isSet bool
}

func (v NullableUpdateSettingsFlowWithOidcMethod) Get() *UpdateSettingsFlowWithOidcMethod {
	return v.value
}

func (v *NullableUpdateSettingsFlowWithOidcMethod) Set(val *UpdateSettingsFlowWithOidcMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingsFlowWithOidcMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingsFlowWithOidcMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingsFlowWithOidcMethod(val *UpdateSettingsFlowWithOidcMethod) *NullableUpdateSettingsFlowWithOidcMethod {
	return &NullableUpdateSettingsFlowWithOidcMethod{value: val, isSet: true}
}

func (v NullableUpdateSettingsFlowWithOidcMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingsFlowWithOidcMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


