/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.6.2
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type JwkAPI interface {

	/*
	CreateJsonWebKeySet Create JSON Web Key

	This endpoint is capable of generating JSON Web Key Sets for you. There a different strategies available, such as symmetric cryptographic keys (HS256, HS512) and asymetric cryptographic keys (RS256, ECDSA). If the specified JSON Web Key Set does not exist, it will be created.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@return JwkAPICreateJsonWebKeySetRequest
	*/
	CreateJsonWebKeySet(ctx context.Context, set string) JwkAPICreateJsonWebKeySetRequest

	// CreateJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	CreateJsonWebKeySetExecute(r JwkAPICreateJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	DeleteJsonWebKey Delete JSON Web Key

	Use this endpoint to delete a single JSON Web Key.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A
JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses
this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens),
and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set
	@param kid The JSON Web Key ID (kid)
	@return JwkAPIDeleteJsonWebKeyRequest
	*/
	DeleteJsonWebKey(ctx context.Context, set string, kid string) JwkAPIDeleteJsonWebKeyRequest

	// DeleteJsonWebKeyExecute executes the request
	DeleteJsonWebKeyExecute(r JwkAPIDeleteJsonWebKeyRequest) (*http.Response, error)

	/*
	DeleteJsonWebKeySet Delete JSON Web Key Set

	Use this endpoint to delete a complete JSON Web Key Set and all the keys in that set.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set
	@return JwkAPIDeleteJsonWebKeySetRequest
	*/
	DeleteJsonWebKeySet(ctx context.Context, set string) JwkAPIDeleteJsonWebKeySetRequest

	// DeleteJsonWebKeySetExecute executes the request
	DeleteJsonWebKeySetExecute(r JwkAPIDeleteJsonWebKeySetRequest) (*http.Response, error)

	/*
	GetJsonWebKey Get JSON Web Key

	This endpoint returns a singular JSON Web Key contained in a set. It is identified by the set and the specific key ID (kid).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set JSON Web Key Set ID
	@param kid JSON Web Key ID
	@return JwkAPIGetJsonWebKeyRequest
	*/
	GetJsonWebKey(ctx context.Context, set string, kid string) JwkAPIGetJsonWebKeyRequest

	// GetJsonWebKeyExecute executes the request
	//  @return JsonWebKeySet
	GetJsonWebKeyExecute(r JwkAPIGetJsonWebKeyRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	GetJsonWebKeySet Retrieve a JSON Web Key Set

	This endpoint can be used to retrieve JWK Sets stored in ORY Hydra.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set JSON Web Key Set ID
	@return JwkAPIGetJsonWebKeySetRequest
	*/
	GetJsonWebKeySet(ctx context.Context, set string) JwkAPIGetJsonWebKeySetRequest

	// GetJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	GetJsonWebKeySetExecute(r JwkAPIGetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	SetJsonWebKey Set JSON Web Key

	Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@param kid JSON Web Key ID
	@return JwkAPISetJsonWebKeyRequest
	*/
	SetJsonWebKey(ctx context.Context, set string, kid string) JwkAPISetJsonWebKeyRequest

	// SetJsonWebKeyExecute executes the request
	//  @return JsonWebKey
	SetJsonWebKeyExecute(r JwkAPISetJsonWebKeyRequest) (*JsonWebKey, *http.Response, error)

	/*
	SetJsonWebKeySet Update a JSON Web Key Set

	Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@return JwkAPISetJsonWebKeySetRequest
	*/
	SetJsonWebKeySet(ctx context.Context, set string) JwkAPISetJsonWebKeySetRequest

	// SetJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	SetJsonWebKeySetExecute(r JwkAPISetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)
}

// JwkAPIService JwkAPI service
type JwkAPIService service

type JwkAPICreateJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
	createJsonWebKeySet *CreateJsonWebKeySet
}

func (r JwkAPICreateJsonWebKeySetRequest) CreateJsonWebKeySet(createJsonWebKeySet CreateJsonWebKeySet) JwkAPICreateJsonWebKeySetRequest {
	r.createJsonWebKeySet = &createJsonWebKeySet
	return r
}

func (r JwkAPICreateJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.CreateJsonWebKeySetExecute(r)
}

/*
CreateJsonWebKeySet Create JSON Web Key

This endpoint is capable of generating JSON Web Key Sets for you. There a different strategies available, such as symmetric cryptographic keys (HS256, HS512) and asymetric cryptographic keys (RS256, ECDSA). If the specified JSON Web Key Set does not exist, it will be created.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @return JwkAPICreateJsonWebKeySetRequest
*/
func (a *JwkAPIService) CreateJsonWebKeySet(ctx context.Context, set string) JwkAPICreateJsonWebKeySetRequest {
	return JwkAPICreateJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkAPIService) CreateJsonWebKeySetExecute(r JwkAPICreateJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.CreateJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createJsonWebKeySet == nil {
		return localVarReturnValue, nil, reportError("createJsonWebKeySet is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createJsonWebKeySet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JwkAPIDeleteJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
	kid string
}

func (r JwkAPIDeleteJsonWebKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJsonWebKeyExecute(r)
}

/*
DeleteJsonWebKey Delete JSON Web Key

Use this endpoint to delete a single JSON Web Key.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A
JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses
this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens),
and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set
 @param kid The JSON Web Key ID (kid)
 @return JwkAPIDeleteJsonWebKeyRequest
*/
func (a *JwkAPIService) DeleteJsonWebKey(ctx context.Context, set string, kid string) JwkAPIDeleteJsonWebKeyRequest {
	return JwkAPIDeleteJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
func (a *JwkAPIService) DeleteJsonWebKeyExecute(r JwkAPIDeleteJsonWebKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.DeleteJsonWebKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterValueToString(r.kid, "kid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JwkAPIDeleteJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
}

func (r JwkAPIDeleteJsonWebKeySetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJsonWebKeySetExecute(r)
}

/*
DeleteJsonWebKeySet Delete JSON Web Key Set

Use this endpoint to delete a complete JSON Web Key Set and all the keys in that set.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set
 @return JwkAPIDeleteJsonWebKeySetRequest
*/
func (a *JwkAPIService) DeleteJsonWebKeySet(ctx context.Context, set string) JwkAPIDeleteJsonWebKeySetRequest {
	return JwkAPIDeleteJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
func (a *JwkAPIService) DeleteJsonWebKeySetExecute(r JwkAPIDeleteJsonWebKeySetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.DeleteJsonWebKeySet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JwkAPIGetJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
	kid string
}

func (r JwkAPIGetJsonWebKeyRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.GetJsonWebKeyExecute(r)
}

/*
GetJsonWebKey Get JSON Web Key

This endpoint returns a singular JSON Web Key contained in a set. It is identified by the set and the specific key ID (kid).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set JSON Web Key Set ID
 @param kid JSON Web Key ID
 @return JwkAPIGetJsonWebKeyRequest
*/
func (a *JwkAPIService) GetJsonWebKey(ctx context.Context, set string, kid string) JwkAPIGetJsonWebKeyRequest {
	return JwkAPIGetJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkAPIService) GetJsonWebKeyExecute(r JwkAPIGetJsonWebKeyRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.GetJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterValueToString(r.kid, "kid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JwkAPIGetJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
}

func (r JwkAPIGetJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.GetJsonWebKeySetExecute(r)
}

/*
GetJsonWebKeySet Retrieve a JSON Web Key Set

This endpoint can be used to retrieve JWK Sets stored in ORY Hydra.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set JSON Web Key Set ID
 @return JwkAPIGetJsonWebKeySetRequest
*/
func (a *JwkAPIService) GetJsonWebKeySet(ctx context.Context, set string) JwkAPIGetJsonWebKeySetRequest {
	return JwkAPIGetJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkAPIService) GetJsonWebKeySetExecute(r JwkAPIGetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.GetJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JwkAPISetJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
	kid string
	jsonWebKey *JsonWebKey
}

func (r JwkAPISetJsonWebKeyRequest) JsonWebKey(jsonWebKey JsonWebKey) JwkAPISetJsonWebKeyRequest {
	r.jsonWebKey = &jsonWebKey
	return r
}

func (r JwkAPISetJsonWebKeyRequest) Execute() (*JsonWebKey, *http.Response, error) {
	return r.ApiService.SetJsonWebKeyExecute(r)
}

/*
SetJsonWebKey Set JSON Web Key

Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @param kid JSON Web Key ID
 @return JwkAPISetJsonWebKeyRequest
*/
func (a *JwkAPIService) SetJsonWebKey(ctx context.Context, set string, kid string) JwkAPISetJsonWebKeyRequest {
	return JwkAPISetJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
//  @return JsonWebKey
func (a *JwkAPIService) SetJsonWebKeyExecute(r JwkAPISetJsonWebKeyRequest) (*JsonWebKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.SetJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterValueToString(r.kid, "kid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.jsonWebKey
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JwkAPISetJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkAPI
	set string
	jsonWebKeySet *JsonWebKeySet
}

func (r JwkAPISetJsonWebKeySetRequest) JsonWebKeySet(jsonWebKeySet JsonWebKeySet) JwkAPISetJsonWebKeySetRequest {
	r.jsonWebKeySet = &jsonWebKeySet
	return r
}

func (r JwkAPISetJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.SetJsonWebKeySetExecute(r)
}

/*
SetJsonWebKeySet Update a JSON Web Key Set

Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @return JwkAPISetJsonWebKeySetRequest
*/
func (a *JwkAPIService) SetJsonWebKeySet(ctx context.Context, set string) JwkAPISetJsonWebKeySetRequest {
	return JwkAPISetJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkAPIService) SetJsonWebKeySetExecute(r JwkAPISetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkAPIService.SetJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterValueToString(r.set, "set")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.jsonWebKeySet
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ErrorOAuth2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
