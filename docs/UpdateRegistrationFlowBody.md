# UpdateRegistrationFlowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRFToken is the anti-CSRF token | [optional] 
**Method** | **string** | Method  Should be set to \&quot;webauthn\&quot; when trying to add, update, or remove a webAuthn pairing. | 
**Password** | **string** | Password to sign the user up with | 
**Traits** | **map[string]interface{}** | The identity&#39;s traits | 
**Provider** | **string** | The provider to register with | 
**WebauthnRegister** | Pointer to **string** | Register a WebAuthn Security Key  It is expected that the JSON returned by the WebAuthn registration process is included here. | [optional] 
**WebauthnRegisterDisplayname** | Pointer to **string** | Name of the WebAuthn Security Key to be Added  A human-readable name for the security key which will be added. | [optional] 

## Methods

### NewUpdateRegistrationFlowBody

`func NewUpdateRegistrationFlowBody(method string, password string, traits map[string]interface{}, provider string, ) *UpdateRegistrationFlowBody`

NewUpdateRegistrationFlowBody instantiates a new UpdateRegistrationFlowBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRegistrationFlowBodyWithDefaults

`func NewUpdateRegistrationFlowBodyWithDefaults() *UpdateRegistrationFlowBody`

NewUpdateRegistrationFlowBodyWithDefaults instantiates a new UpdateRegistrationFlowBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *UpdateRegistrationFlowBody) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *UpdateRegistrationFlowBody) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *UpdateRegistrationFlowBody) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *UpdateRegistrationFlowBody) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMethod

`func (o *UpdateRegistrationFlowBody) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *UpdateRegistrationFlowBody) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *UpdateRegistrationFlowBody) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPassword

`func (o *UpdateRegistrationFlowBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRegistrationFlowBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRegistrationFlowBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTraits

`func (o *UpdateRegistrationFlowBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *UpdateRegistrationFlowBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *UpdateRegistrationFlowBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.


### GetProvider

`func (o *UpdateRegistrationFlowBody) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UpdateRegistrationFlowBody) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UpdateRegistrationFlowBody) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetWebauthnRegister

`func (o *UpdateRegistrationFlowBody) GetWebauthnRegister() string`

GetWebauthnRegister returns the WebauthnRegister field if non-nil, zero value otherwise.

### GetWebauthnRegisterOk

`func (o *UpdateRegistrationFlowBody) GetWebauthnRegisterOk() (*string, bool)`

GetWebauthnRegisterOk returns a tuple with the WebauthnRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnRegister

`func (o *UpdateRegistrationFlowBody) SetWebauthnRegister(v string)`

SetWebauthnRegister sets WebauthnRegister field to given value.

### HasWebauthnRegister

`func (o *UpdateRegistrationFlowBody) HasWebauthnRegister() bool`

HasWebauthnRegister returns a boolean if a field has been set.

### GetWebauthnRegisterDisplayname

`func (o *UpdateRegistrationFlowBody) GetWebauthnRegisterDisplayname() string`

GetWebauthnRegisterDisplayname returns the WebauthnRegisterDisplayname field if non-nil, zero value otherwise.

### GetWebauthnRegisterDisplaynameOk

`func (o *UpdateRegistrationFlowBody) GetWebauthnRegisterDisplaynameOk() (*string, bool)`

GetWebauthnRegisterDisplaynameOk returns a tuple with the WebauthnRegisterDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthnRegisterDisplayname

`func (o *UpdateRegistrationFlowBody) SetWebauthnRegisterDisplayname(v string)`

SetWebauthnRegisterDisplayname sets WebauthnRegisterDisplayname field to given value.

### HasWebauthnRegisterDisplayname

`func (o *UpdateRegistrationFlowBody) HasWebauthnRegisterDisplayname() bool`

HasWebauthnRegisterDisplayname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

