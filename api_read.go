/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.1.0-alpha.10
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type ReadApi interface {

	/*
	 * GetCheck Check a relation tuple
	 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiGetCheckRequest
	 */
	GetCheck(ctx context.Context) ReadApiApiGetCheckRequest

	/*
	 * GetCheckExecute executes the request
	 * @return GetCheckResponse
	 */
	GetCheckExecute(r ReadApiApiGetCheckRequest) (*GetCheckResponse, *http.Response, error)

	/*
	 * GetCheckMirrorStatus Check a relation tuple
	 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiGetCheckMirrorStatusRequest
	 */
	GetCheckMirrorStatus(ctx context.Context) ReadApiApiGetCheckMirrorStatusRequest

	/*
	 * GetCheckMirrorStatusExecute executes the request
	 * @return GetCheckResponse
	 */
	GetCheckMirrorStatusExecute(r ReadApiApiGetCheckMirrorStatusRequest) (*GetCheckResponse, *http.Response, error)

	/*
	 * GetExpand Expand a Relation Tuple
	 * Use this endpoint to expand a relation tuple.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiGetExpandRequest
	 */
	GetExpand(ctx context.Context) ReadApiApiGetExpandRequest

	/*
	 * GetExpandExecute executes the request
	 * @return ExpandTree
	 */
	GetExpandExecute(r ReadApiApiGetExpandRequest) (*ExpandTree, *http.Response, error)

	/*
	 * GetRelationTuples Query relation tuples
	 * Get all relation tuples that match the query. Only the namespace field is required.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiGetRelationTuplesRequest
	 */
	GetRelationTuples(ctx context.Context) ReadApiApiGetRelationTuplesRequest

	/*
	 * GetRelationTuplesExecute executes the request
	 * @return GetRelationTuplesResponse
	 */
	GetRelationTuplesExecute(r ReadApiApiGetRelationTuplesRequest) (*GetRelationTuplesResponse, *http.Response, error)

	/*
	 * PostCheck Check a relation tuple
	 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiPostCheckRequest
	 */
	PostCheck(ctx context.Context) ReadApiApiPostCheckRequest

	/*
	 * PostCheckExecute executes the request
	 * @return GetCheckResponse
	 */
	PostCheckExecute(r ReadApiApiPostCheckRequest) (*GetCheckResponse, *http.Response, error)

	/*
	 * PostCheckMirrorStatus Check a relation tuple
	 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ReadApiApiPostCheckMirrorStatusRequest
	 */
	PostCheckMirrorStatus(ctx context.Context) ReadApiApiPostCheckMirrorStatusRequest

	/*
	 * PostCheckMirrorStatusExecute executes the request
	 * @return GetCheckResponse
	 */
	PostCheckMirrorStatusExecute(r ReadApiApiPostCheckMirrorStatusRequest) (*GetCheckResponse, *http.Response, error)
}

// ReadApiService ReadApi service
type ReadApiService service

type ReadApiApiGetCheckRequest struct {
	ctx context.Context
	ApiService ReadApi
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
	maxDepth *int64
}

func (r ReadApiApiGetCheckRequest) Namespace(namespace string) ReadApiApiGetCheckRequest {
	r.namespace = &namespace
	return r
}
func (r ReadApiApiGetCheckRequest) Object(object string) ReadApiApiGetCheckRequest {
	r.object = &object
	return r
}
func (r ReadApiApiGetCheckRequest) Relation(relation string) ReadApiApiGetCheckRequest {
	r.relation = &relation
	return r
}
func (r ReadApiApiGetCheckRequest) SubjectId(subjectId string) ReadApiApiGetCheckRequest {
	r.subjectId = &subjectId
	return r
}
func (r ReadApiApiGetCheckRequest) SubjectSetNamespace(subjectSetNamespace string) ReadApiApiGetCheckRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r ReadApiApiGetCheckRequest) SubjectSetObject(subjectSetObject string) ReadApiApiGetCheckRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r ReadApiApiGetCheckRequest) SubjectSetRelation(subjectSetRelation string) ReadApiApiGetCheckRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}
func (r ReadApiApiGetCheckRequest) MaxDepth(maxDepth int64) ReadApiApiGetCheckRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r ReadApiApiGetCheckRequest) Execute() (*GetCheckResponse, *http.Response, error) {
	return r.ApiService.GetCheckExecute(r)
}

/*
 * GetCheck Check a relation tuple
 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiGetCheckRequest
 */
func (a *ReadApiService) GetCheck(ctx context.Context) ReadApiApiGetCheckRequest {
	return ReadApiApiGetCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return GetCheckResponse
 */
func (a *ReadApiService) GetCheckExecute(r ReadApiApiGetCheckRequest) (*GetCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *GetCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.GetCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ReadApiApiGetCheckMirrorStatusRequest struct {
	ctx context.Context
	ApiService ReadApi
}


func (r ReadApiApiGetCheckMirrorStatusRequest) Execute() (*GetCheckResponse, *http.Response, error) {
	return r.ApiService.GetCheckMirrorStatusExecute(r)
}

/*
 * GetCheckMirrorStatus Check a relation tuple
 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiGetCheckMirrorStatusRequest
 */
func (a *ReadApiService) GetCheckMirrorStatus(ctx context.Context) ReadApiApiGetCheckMirrorStatusRequest {
	return ReadApiApiGetCheckMirrorStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return GetCheckResponse
 */
func (a *ReadApiService) GetCheckMirrorStatusExecute(r ReadApiApiGetCheckMirrorStatusRequest) (*GetCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *GetCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.GetCheckMirrorStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetCheckResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ReadApiApiGetExpandRequest struct {
	ctx context.Context
	ApiService ReadApi
	namespace *string
	object *string
	relation *string
	maxDepth *int64
}

func (r ReadApiApiGetExpandRequest) Namespace(namespace string) ReadApiApiGetExpandRequest {
	r.namespace = &namespace
	return r
}
func (r ReadApiApiGetExpandRequest) Object(object string) ReadApiApiGetExpandRequest {
	r.object = &object
	return r
}
func (r ReadApiApiGetExpandRequest) Relation(relation string) ReadApiApiGetExpandRequest {
	r.relation = &relation
	return r
}
func (r ReadApiApiGetExpandRequest) MaxDepth(maxDepth int64) ReadApiApiGetExpandRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r ReadApiApiGetExpandRequest) Execute() (*ExpandTree, *http.Response, error) {
	return r.ApiService.GetExpandExecute(r)
}

/*
 * GetExpand Expand a Relation Tuple
 * Use this endpoint to expand a relation tuple.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiGetExpandRequest
 */
func (a *ReadApiService) GetExpand(ctx context.Context) ReadApiApiGetExpandRequest {
	return ReadApiApiGetExpandRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ExpandTree
 */
func (a *ReadApiService) GetExpandExecute(r ReadApiApiGetExpandRequest) (*ExpandTree, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *ExpandTree
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.GetExpand")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/expand"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.namespace == nil {
		return localVarReturnValue, nil, reportError("namespace is required and must be specified")
	}
	if r.object == nil {
		return localVarReturnValue, nil, reportError("object is required and must be specified")
	}
	if r.relation == nil {
		return localVarReturnValue, nil, reportError("relation is required and must be specified")
	}

	localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ReadApiApiGetRelationTuplesRequest struct {
	ctx context.Context
	ApiService ReadApi
	pageToken *string
	pageSize *int64
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

func (r ReadApiApiGetRelationTuplesRequest) PageToken(pageToken string) ReadApiApiGetRelationTuplesRequest {
	r.pageToken = &pageToken
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) PageSize(pageSize int64) ReadApiApiGetRelationTuplesRequest {
	r.pageSize = &pageSize
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) Namespace(namespace string) ReadApiApiGetRelationTuplesRequest {
	r.namespace = &namespace
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) Object(object string) ReadApiApiGetRelationTuplesRequest {
	r.object = &object
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) Relation(relation string) ReadApiApiGetRelationTuplesRequest {
	r.relation = &relation
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) SubjectId(subjectId string) ReadApiApiGetRelationTuplesRequest {
	r.subjectId = &subjectId
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) SubjectSetNamespace(subjectSetNamespace string) ReadApiApiGetRelationTuplesRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) SubjectSetObject(subjectSetObject string) ReadApiApiGetRelationTuplesRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r ReadApiApiGetRelationTuplesRequest) SubjectSetRelation(subjectSetRelation string) ReadApiApiGetRelationTuplesRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r ReadApiApiGetRelationTuplesRequest) Execute() (*GetRelationTuplesResponse, *http.Response, error) {
	return r.ApiService.GetRelationTuplesExecute(r)
}

/*
 * GetRelationTuples Query relation tuples
 * Get all relation tuples that match the query. Only the namespace field is required.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiGetRelationTuplesRequest
 */
func (a *ReadApiService) GetRelationTuples(ctx context.Context) ReadApiApiGetRelationTuplesRequest {
	return ReadApiApiGetRelationTuplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return GetRelationTuplesResponse
 */
func (a *ReadApiService) GetRelationTuplesExecute(r ReadApiApiGetRelationTuplesRequest) (*GetRelationTuplesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *GetRelationTuplesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.GetRelationTuples")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ReadApiApiPostCheckRequest struct {
	ctx context.Context
	ApiService ReadApi
	maxDepth *int64
	relationQuery *RelationQuery
}

func (r ReadApiApiPostCheckRequest) MaxDepth(maxDepth int64) ReadApiApiPostCheckRequest {
	r.maxDepth = &maxDepth
	return r
}
func (r ReadApiApiPostCheckRequest) RelationQuery(relationQuery RelationQuery) ReadApiApiPostCheckRequest {
	r.relationQuery = &relationQuery
	return r
}

func (r ReadApiApiPostCheckRequest) Execute() (*GetCheckResponse, *http.Response, error) {
	return r.ApiService.PostCheckExecute(r)
}

/*
 * PostCheck Check a relation tuple
 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiPostCheckRequest
 */
func (a *ReadApiService) PostCheck(ctx context.Context) ReadApiApiPostCheckRequest {
	return ReadApiApiPostCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return GetCheckResponse
 */
func (a *ReadApiService) PostCheckExecute(r ReadApiApiPostCheckRequest) (*GetCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *GetCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.PostCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		localVarQueryParams.Add("max-depth", parameterToString(*r.maxDepth, ""))
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
	localVarPostBody = r.relationQuery
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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

type ReadApiApiPostCheckMirrorStatusRequest struct {
	ctx context.Context
	ApiService ReadApi
}


func (r ReadApiApiPostCheckMirrorStatusRequest) Execute() (*GetCheckResponse, *http.Response, error) {
	return r.ApiService.PostCheckMirrorStatusExecute(r)
}

/*
 * PostCheckMirrorStatus Check a relation tuple
 * To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ReadApiApiPostCheckMirrorStatusRequest
 */
func (a *ReadApiService) PostCheckMirrorStatus(ctx context.Context) ReadApiApiPostCheckMirrorStatusRequest {
	return ReadApiApiPostCheckMirrorStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return GetCheckResponse
 */
func (a *ReadApiService) PostCheckMirrorStatusExecute(r ReadApiApiPostCheckMirrorStatusRequest) (*GetCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *GetCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReadApiService.PostCheckMirrorStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetCheckResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
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
