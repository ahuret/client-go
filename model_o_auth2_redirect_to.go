/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.2.6
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OAuth2RedirectTo Contains a redirect URL used to complete a login, consent, or logout request.
type OAuth2RedirectTo struct {
	// RedirectURL is the URL which you should redirect the user's browser to once the authentication process is completed.
	RedirectTo string `json:"redirect_to"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RedirectTo OAuth2RedirectTo

// NewOAuth2RedirectTo instantiates a new OAuth2RedirectTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RedirectTo(redirectTo string) *OAuth2RedirectTo {
	this := OAuth2RedirectTo{}
	this.RedirectTo = redirectTo
	return &this
}

// NewOAuth2RedirectToWithDefaults instantiates a new OAuth2RedirectTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RedirectToWithDefaults() *OAuth2RedirectTo {
	this := OAuth2RedirectTo{}
	return &this
}

// GetRedirectTo returns the RedirectTo field value
func (o *OAuth2RedirectTo) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *OAuth2RedirectTo) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *OAuth2RedirectTo) SetRedirectTo(v string) {
	o.RedirectTo = v
}

func (o OAuth2RedirectTo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["redirect_to"] = o.RedirectTo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2RedirectTo) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2RedirectTo := _OAuth2RedirectTo{}

	if err = json.Unmarshal(bytes, &varOAuth2RedirectTo); err == nil {
		*o = OAuth2RedirectTo(varOAuth2RedirectTo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "redirect_to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2RedirectTo struct {
	value *OAuth2RedirectTo
	isSet bool
}

func (v NullableOAuth2RedirectTo) Get() *OAuth2RedirectTo {
	return v.value
}

func (v *NullableOAuth2RedirectTo) Set(val *OAuth2RedirectTo) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RedirectTo) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RedirectTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RedirectTo(val *OAuth2RedirectTo) *NullableOAuth2RedirectTo {
	return &NullableOAuth2RedirectTo{value: val, isSet: true}
}

func (v NullableOAuth2RedirectTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RedirectTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


