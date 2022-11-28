/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.0.0
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type JwkApi interface {

	/*
	CreateJsonWebKeySet Create JSON Web Key

	This endpoint is capable of generating JSON Web Key Sets for you. There a different strategies available, such as symmetric cryptographic keys (HS256, HS512) and asymetric cryptographic keys (RS256, ECDSA). If the specified JSON Web Key Set does not exist, it will be created.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@return JwkApiCreateJsonWebKeySetRequest
	*/
	CreateJsonWebKeySet(ctx context.Context, set string) JwkApiCreateJsonWebKeySetRequest

	// CreateJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	CreateJsonWebKeySetExecute(r JwkApiCreateJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)

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
	@return JwkApiDeleteJsonWebKeyRequest
	*/
	DeleteJsonWebKey(ctx context.Context, set string, kid string) JwkApiDeleteJsonWebKeyRequest

	// DeleteJsonWebKeyExecute executes the request
	DeleteJsonWebKeyExecute(r JwkApiDeleteJsonWebKeyRequest) (*http.Response, error)

	/*
	DeleteJsonWebKeySet Delete JSON Web Key Set

	Use this endpoint to delete a complete JSON Web Key Set and all the keys in that set.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set
	@return JwkApiDeleteJsonWebKeySetRequest
	*/
	DeleteJsonWebKeySet(ctx context.Context, set string) JwkApiDeleteJsonWebKeySetRequest

	// DeleteJsonWebKeySetExecute executes the request
	DeleteJsonWebKeySetExecute(r JwkApiDeleteJsonWebKeySetRequest) (*http.Response, error)

	/*
	GetJsonWebKey Get JSON Web Key

	This endpoint returns a singular JSON Web Key contained in a set. It is identified by the set and the specific key ID (kid).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set JSON Web Key Set ID
	@param kid JSON Web Key ID
	@return JwkApiGetJsonWebKeyRequest
	*/
	GetJsonWebKey(ctx context.Context, set string, kid string) JwkApiGetJsonWebKeyRequest

	// GetJsonWebKeyExecute executes the request
	//  @return JsonWebKeySet
	GetJsonWebKeyExecute(r JwkApiGetJsonWebKeyRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	GetJsonWebKeySet Retrieve a JSON Web Key Set

	This endpoint can be used to retrieve JWK Sets stored in ORY Hydra.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set JSON Web Key Set ID
	@return JwkApiGetJsonWebKeySetRequest
	*/
	GetJsonWebKeySet(ctx context.Context, set string) JwkApiGetJsonWebKeySetRequest

	// GetJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	GetJsonWebKeySetExecute(r JwkApiGetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	SetJsonWebKey Set JSON Web Key

	Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@param kid JSON Web Key ID
	@return JwkApiSetJsonWebKeyRequest
	*/
	SetJsonWebKey(ctx context.Context, set string, kid string) JwkApiSetJsonWebKeyRequest

	// SetJsonWebKeyExecute executes the request
	//  @return JsonWebKey
	SetJsonWebKeyExecute(r JwkApiSetJsonWebKeyRequest) (*JsonWebKey, *http.Response, error)

	/*
	SetJsonWebKeySet Update a JSON Web Key Set

	Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param set The JSON Web Key Set ID
	@return JwkApiSetJsonWebKeySetRequest
	*/
	SetJsonWebKeySet(ctx context.Context, set string) JwkApiSetJsonWebKeySetRequest

	// SetJsonWebKeySetExecute executes the request
	//  @return JsonWebKeySet
	SetJsonWebKeySetExecute(r JwkApiSetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error)
}

// JwkApiService JwkApi service
type JwkApiService service

type JwkApiCreateJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
	createJsonWebKeySet *CreateJsonWebKeySet
}

func (r JwkApiCreateJsonWebKeySetRequest) CreateJsonWebKeySet(createJsonWebKeySet CreateJsonWebKeySet) JwkApiCreateJsonWebKeySetRequest {
	r.createJsonWebKeySet = &createJsonWebKeySet
	return r
}

func (r JwkApiCreateJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.CreateJsonWebKeySetExecute(r)
}

/*
CreateJsonWebKeySet Create JSON Web Key

This endpoint is capable of generating JSON Web Key Sets for you. There a different strategies available, such as symmetric cryptographic keys (HS256, HS512) and asymetric cryptographic keys (RS256, ECDSA). If the specified JSON Web Key Set does not exist, it will be created.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @return JwkApiCreateJsonWebKeySetRequest
*/
func (a *JwkApiService) CreateJsonWebKeySet(ctx context.Context, set string) JwkApiCreateJsonWebKeySetRequest {
	return JwkApiCreateJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkApiService) CreateJsonWebKeySetExecute(r JwkApiCreateJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.CreateJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type JwkApiDeleteJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
	kid string
}

func (r JwkApiDeleteJsonWebKeyRequest) Execute() (*http.Response, error) {
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
 @return JwkApiDeleteJsonWebKeyRequest
*/
func (a *JwkApiService) DeleteJsonWebKey(ctx context.Context, set string, kid string) JwkApiDeleteJsonWebKeyRequest {
	return JwkApiDeleteJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
func (a *JwkApiService) DeleteJsonWebKeyExecute(r JwkApiDeleteJsonWebKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.DeleteJsonWebKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterToString(r.kid, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JwkApiDeleteJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
}

func (r JwkApiDeleteJsonWebKeySetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJsonWebKeySetExecute(r)
}

/*
DeleteJsonWebKeySet Delete JSON Web Key Set

Use this endpoint to delete a complete JSON Web Key Set and all the keys in that set.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set
 @return JwkApiDeleteJsonWebKeySetRequest
*/
func (a *JwkApiService) DeleteJsonWebKeySet(ctx context.Context, set string) JwkApiDeleteJsonWebKeySetRequest {
	return JwkApiDeleteJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
func (a *JwkApiService) DeleteJsonWebKeySetExecute(r JwkApiDeleteJsonWebKeySetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.DeleteJsonWebKeySet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type JwkApiGetJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
	kid string
}

func (r JwkApiGetJsonWebKeyRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.GetJsonWebKeyExecute(r)
}

/*
GetJsonWebKey Get JSON Web Key

This endpoint returns a singular JSON Web Key contained in a set. It is identified by the set and the specific key ID (kid).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set JSON Web Key Set ID
 @param kid JSON Web Key ID
 @return JwkApiGetJsonWebKeyRequest
*/
func (a *JwkApiService) GetJsonWebKey(ctx context.Context, set string, kid string) JwkApiGetJsonWebKeyRequest {
	return JwkApiGetJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkApiService) GetJsonWebKeyExecute(r JwkApiGetJsonWebKeyRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.GetJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterToString(r.kid, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type JwkApiGetJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
}

func (r JwkApiGetJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.GetJsonWebKeySetExecute(r)
}

/*
GetJsonWebKeySet Retrieve a JSON Web Key Set

This endpoint can be used to retrieve JWK Sets stored in ORY Hydra.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set JSON Web Key Set ID
 @return JwkApiGetJsonWebKeySetRequest
*/
func (a *JwkApiService) GetJsonWebKeySet(ctx context.Context, set string) JwkApiGetJsonWebKeySetRequest {
	return JwkApiGetJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkApiService) GetJsonWebKeySetExecute(r JwkApiGetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.GetJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type JwkApiSetJsonWebKeyRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
	kid string
	jsonWebKey *JsonWebKey
}

func (r JwkApiSetJsonWebKeyRequest) JsonWebKey(jsonWebKey JsonWebKey) JwkApiSetJsonWebKeyRequest {
	r.jsonWebKey = &jsonWebKey
	return r
}

func (r JwkApiSetJsonWebKeyRequest) Execute() (*JsonWebKey, *http.Response, error) {
	return r.ApiService.SetJsonWebKeyExecute(r)
}

/*
SetJsonWebKey Set JSON Web Key

Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @param kid JSON Web Key ID
 @return JwkApiSetJsonWebKeyRequest
*/
func (a *JwkApiService) SetJsonWebKey(ctx context.Context, set string, kid string) JwkApiSetJsonWebKeyRequest {
	return JwkApiSetJsonWebKeyRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
		kid: kid,
	}
}

// Execute executes the request
//  @return JsonWebKey
func (a *JwkApiService) SetJsonWebKeyExecute(r JwkApiSetJsonWebKeyRequest) (*JsonWebKey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.SetJsonWebKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}/{kid}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kid"+"}", url.PathEscape(parameterToString(r.kid, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type JwkApiSetJsonWebKeySetRequest struct {
	ctx context.Context
	ApiService JwkApi
	set string
	jsonWebKeySet *JsonWebKeySet
}

func (r JwkApiSetJsonWebKeySetRequest) JsonWebKeySet(jsonWebKeySet JsonWebKeySet) JwkApiSetJsonWebKeySetRequest {
	r.jsonWebKeySet = &jsonWebKeySet
	return r
}

func (r JwkApiSetJsonWebKeySetRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.SetJsonWebKeySetExecute(r)
}

/*
SetJsonWebKeySet Update a JSON Web Key Set

Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.

A JSON Web Key (JWK) is a JavaScript Object Notation (JSON) data structure that represents a cryptographic key. A JWK Set is a JSON data structure that represents a set of JWKs. A JSON Web Key is identified by its set and key id. ORY Hydra uses this functionality to store cryptographic keys used for TLS and JSON Web Tokens (such as OpenID Connect ID tokens), and allows storing user-defined keys as well.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param set The JSON Web Key Set ID
 @return JwkApiSetJsonWebKeySetRequest
*/
func (a *JwkApiService) SetJsonWebKeySet(ctx context.Context, set string) JwkApiSetJsonWebKeySetRequest {
	return JwkApiSetJsonWebKeySetRequest{
		ApiService: a,
		ctx: ctx,
		set: set,
	}
}

// Execute executes the request
//  @return JsonWebKeySet
func (a *JwkApiService) SetJsonWebKeySetExecute(r JwkApiSetJsonWebKeySetRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JwkApiService.SetJsonWebKeySet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/keys/{set}"
	localVarPath = strings.Replace(localVarPath, "{"+"set"+"}", url.PathEscape(parameterToString(r.set, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
