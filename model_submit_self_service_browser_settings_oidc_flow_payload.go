/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.12
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceBrowserSettingsOIDCFlowPayload struct for SubmitSelfServiceBrowserSettingsOIDCFlowPayload
type SubmitSelfServiceBrowserSettingsOIDCFlowPayload struct {
	// Flow ID is the flow's ID.  in: query
	Flow *string `json:"flow,omitempty"`
	// Link this provider  Either this or `unlink` must be set.  type: string in: body
	Link *string `json:"link,omitempty"`
	// Unlink this provider  Either this or `link` must be set.  type: string in: body
	Unlink *string `json:"unlink,omitempty"`
}

// NewSubmitSelfServiceBrowserSettingsOIDCFlowPayload instantiates a new SubmitSelfServiceBrowserSettingsOIDCFlowPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceBrowserSettingsOIDCFlowPayload() *SubmitSelfServiceBrowserSettingsOIDCFlowPayload {
	this := SubmitSelfServiceBrowserSettingsOIDCFlowPayload{}
	return &this
}

// NewSubmitSelfServiceBrowserSettingsOIDCFlowPayloadWithDefaults instantiates a new SubmitSelfServiceBrowserSettingsOIDCFlowPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceBrowserSettingsOIDCFlowPayloadWithDefaults() *SubmitSelfServiceBrowserSettingsOIDCFlowPayload {
	this := SubmitSelfServiceBrowserSettingsOIDCFlowPayload{}
	return &this
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetFlow() string {
	if o == nil || o.Flow == nil {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetFlowOk() (*string, bool) {
	if o == nil || o.Flow == nil {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) HasFlow() bool {
	if o != nil && o.Flow != nil {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) SetFlow(v string) {
	o.Flow = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) SetLink(v string) {
	o.Link = &v
}

// GetUnlink returns the Unlink field value if set, zero value otherwise.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetUnlink() string {
	if o == nil || o.Unlink == nil {
		var ret string
		return ret
	}
	return *o.Unlink
}

// GetUnlinkOk returns a tuple with the Unlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) GetUnlinkOk() (*string, bool) {
	if o == nil || o.Unlink == nil {
		return nil, false
	}
	return o.Unlink, true
}

// HasUnlink returns a boolean if a field has been set.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) HasUnlink() bool {
	if o != nil && o.Unlink != nil {
		return true
	}

	return false
}

// SetUnlink gets a reference to the given string and assigns it to the Unlink field.
func (o *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) SetUnlink(v string) {
	o.Unlink = &v
}

func (o SubmitSelfServiceBrowserSettingsOIDCFlowPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flow != nil {
		toSerialize["flow"] = o.Flow
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Unlink != nil {
		toSerialize["unlink"] = o.Unlink
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload struct {
	value *SubmitSelfServiceBrowserSettingsOIDCFlowPayload
	isSet bool
}

func (v NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) Get() *SubmitSelfServiceBrowserSettingsOIDCFlowPayload {
	return v.value
}

func (v *NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) Set(val *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload(val *SubmitSelfServiceBrowserSettingsOIDCFlowPayload) *NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload {
	return &NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceBrowserSettingsOIDCFlowPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


