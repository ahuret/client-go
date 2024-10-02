/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.15.5
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
)


type PermissionAPI interface {

	/*
	BatchCheckPermission Batch check permissions

	To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPIBatchCheckPermissionRequest
	*/
	BatchCheckPermission(ctx context.Context) PermissionAPIBatchCheckPermissionRequest

	// BatchCheckPermissionExecute executes the request
	//  @return BatchCheckPermissionResult
	BatchCheckPermissionExecute(r PermissionAPIBatchCheckPermissionRequest) (*BatchCheckPermissionResult, *http.Response, error)

	/*
	CheckPermission Check a permission

	To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPICheckPermissionRequest
	*/
	CheckPermission(ctx context.Context) PermissionAPICheckPermissionRequest

	// CheckPermissionExecute executes the request
	//  @return CheckPermissionResult
	CheckPermissionExecute(r PermissionAPICheckPermissionRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	CheckPermissionOrError Check a permission

	To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPICheckPermissionOrErrorRequest
	*/
	CheckPermissionOrError(ctx context.Context) PermissionAPICheckPermissionOrErrorRequest

	// CheckPermissionOrErrorExecute executes the request
	//  @return CheckPermissionResult
	CheckPermissionOrErrorExecute(r PermissionAPICheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	ExpandPermissions Expand a Relationship into permissions.

	Use this endpoint to expand a relationship tuple into permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPIExpandPermissionsRequest
	*/
	ExpandPermissions(ctx context.Context) PermissionAPIExpandPermissionsRequest

	// ExpandPermissionsExecute executes the request
	//  @return ExpandedPermissionTree
	ExpandPermissionsExecute(r PermissionAPIExpandPermissionsRequest) (*ExpandedPermissionTree, *http.Response, error)

	/*
	PostCheckPermission Check a permission

	To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPIPostCheckPermissionRequest
	*/
	PostCheckPermission(ctx context.Context) PermissionAPIPostCheckPermissionRequest

	// PostCheckPermissionExecute executes the request
	//  @return CheckPermissionResult
	PostCheckPermissionExecute(r PermissionAPIPostCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error)

	/*
	PostCheckPermissionOrError Check a permission

	To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PermissionAPIPostCheckPermissionOrErrorRequest
	*/
	PostCheckPermissionOrError(ctx context.Context) PermissionAPIPostCheckPermissionOrErrorRequest

	// PostCheckPermissionOrErrorExecute executes the request
	//  @return CheckPermissionResult
	PostCheckPermissionOrErrorExecute(r PermissionAPIPostCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error)
}

// PermissionAPIService PermissionAPI service
type PermissionAPIService service

type PermissionAPIBatchCheckPermissionRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	maxDepth *int64
	batchCheckPermissionBody *BatchCheckPermissionBody
}

func (r PermissionAPIBatchCheckPermissionRequest) MaxDepth(maxDepth int64) PermissionAPIBatchCheckPermissionRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPIBatchCheckPermissionRequest) BatchCheckPermissionBody(batchCheckPermissionBody BatchCheckPermissionBody) PermissionAPIBatchCheckPermissionRequest {
	r.batchCheckPermissionBody = &batchCheckPermissionBody
	return r
}

func (r PermissionAPIBatchCheckPermissionRequest) Execute() (*BatchCheckPermissionResult, *http.Response, error) {
	return r.ApiService.BatchCheckPermissionExecute(r)
}

/*
BatchCheckPermission Batch check permissions

To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPIBatchCheckPermissionRequest
*/
func (a *PermissionAPIService) BatchCheckPermission(ctx context.Context) PermissionAPIBatchCheckPermissionRequest {
	return PermissionAPIBatchCheckPermissionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BatchCheckPermissionResult
func (a *PermissionAPIService) BatchCheckPermissionExecute(r PermissionAPIBatchCheckPermissionRequest) (*BatchCheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BatchCheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.BatchCheckPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/batch/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
	localVarPostBody = r.batchCheckPermissionBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type PermissionAPICheckPermissionRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
	maxDepth *int64
}

// Namespace of the Relationship
func (r PermissionAPICheckPermissionRequest) Namespace(namespace string) PermissionAPICheckPermissionRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r PermissionAPICheckPermissionRequest) Object(object string) PermissionAPICheckPermissionRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r PermissionAPICheckPermissionRequest) Relation(relation string) PermissionAPICheckPermissionRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r PermissionAPICheckPermissionRequest) SubjectId(subjectId string) PermissionAPICheckPermissionRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r PermissionAPICheckPermissionRequest) SubjectSetNamespace(subjectSetNamespace string) PermissionAPICheckPermissionRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r PermissionAPICheckPermissionRequest) SubjectSetObject(subjectSetObject string) PermissionAPICheckPermissionRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r PermissionAPICheckPermissionRequest) SubjectSetRelation(subjectSetRelation string) PermissionAPICheckPermissionRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r PermissionAPICheckPermissionRequest) MaxDepth(maxDepth int64) PermissionAPICheckPermissionRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPICheckPermissionRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.CheckPermissionExecute(r)
}

/*
CheckPermission Check a permission

To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPICheckPermissionRequest
*/
func (a *PermissionAPIService) CheckPermission(ctx context.Context) PermissionAPICheckPermissionRequest {
	return PermissionAPICheckPermissionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckPermissionResult
func (a *PermissionAPIService) CheckPermissionExecute(r PermissionAPICheckPermissionRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.CheckPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "")
	}
	if r.relation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relation", r.relation, "")
	}
	if r.subjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_id", r.subjectId, "")
	}
	if r.subjectSetNamespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.namespace", r.subjectSetNamespace, "")
	}
	if r.subjectSetObject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.object", r.subjectSetObject, "")
	}
	if r.subjectSetRelation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.relation", r.subjectSetRelation, "")
	}
	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type PermissionAPICheckPermissionOrErrorRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
	maxDepth *int64
}

// Namespace of the Relationship
func (r PermissionAPICheckPermissionOrErrorRequest) Namespace(namespace string) PermissionAPICheckPermissionOrErrorRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r PermissionAPICheckPermissionOrErrorRequest) Object(object string) PermissionAPICheckPermissionOrErrorRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r PermissionAPICheckPermissionOrErrorRequest) Relation(relation string) PermissionAPICheckPermissionOrErrorRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r PermissionAPICheckPermissionOrErrorRequest) SubjectId(subjectId string) PermissionAPICheckPermissionOrErrorRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r PermissionAPICheckPermissionOrErrorRequest) SubjectSetNamespace(subjectSetNamespace string) PermissionAPICheckPermissionOrErrorRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r PermissionAPICheckPermissionOrErrorRequest) SubjectSetObject(subjectSetObject string) PermissionAPICheckPermissionOrErrorRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r PermissionAPICheckPermissionOrErrorRequest) SubjectSetRelation(subjectSetRelation string) PermissionAPICheckPermissionOrErrorRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r PermissionAPICheckPermissionOrErrorRequest) MaxDepth(maxDepth int64) PermissionAPICheckPermissionOrErrorRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPICheckPermissionOrErrorRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.CheckPermissionOrErrorExecute(r)
}

/*
CheckPermissionOrError Check a permission

To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPICheckPermissionOrErrorRequest
*/
func (a *PermissionAPIService) CheckPermissionOrError(ctx context.Context) PermissionAPICheckPermissionOrErrorRequest {
	return PermissionAPICheckPermissionOrErrorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckPermissionResult
func (a *PermissionAPIService) CheckPermissionOrErrorExecute(r PermissionAPICheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.CheckPermissionOrError")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "")
	}
	if r.relation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relation", r.relation, "")
	}
	if r.subjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_id", r.subjectId, "")
	}
	if r.subjectSetNamespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.namespace", r.subjectSetNamespace, "")
	}
	if r.subjectSetObject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.object", r.subjectSetObject, "")
	}
	if r.subjectSetRelation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subject_set.relation", r.subjectSetRelation, "")
	}
	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CheckPermissionResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type PermissionAPIExpandPermissionsRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	namespace *string
	object *string
	relation *string
	maxDepth *int64
}

// Namespace of the Subject Set
func (r PermissionAPIExpandPermissionsRequest) Namespace(namespace string) PermissionAPIExpandPermissionsRequest {
	r.namespace = &namespace
	return r
}

// Object of the Subject Set
func (r PermissionAPIExpandPermissionsRequest) Object(object string) PermissionAPIExpandPermissionsRequest {
	r.object = &object
	return r
}

// Relation of the Subject Set
func (r PermissionAPIExpandPermissionsRequest) Relation(relation string) PermissionAPIExpandPermissionsRequest {
	r.relation = &relation
	return r
}

func (r PermissionAPIExpandPermissionsRequest) MaxDepth(maxDepth int64) PermissionAPIExpandPermissionsRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPIExpandPermissionsRequest) Execute() (*ExpandedPermissionTree, *http.Response, error) {
	return r.ApiService.ExpandPermissionsExecute(r)
}

/*
ExpandPermissions Expand a Relationship into permissions.

Use this endpoint to expand a relationship tuple into permissions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPIExpandPermissionsRequest
*/
func (a *PermissionAPIService) ExpandPermissions(ctx context.Context) PermissionAPIExpandPermissionsRequest {
	return PermissionAPIExpandPermissionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExpandedPermissionTree
func (a *PermissionAPIService) ExpandPermissionsExecute(r PermissionAPIExpandPermissionsRequest) (*ExpandedPermissionTree, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExpandedPermissionTree
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.ExpandPermissions")
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

	parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "relation", r.relation, "")
	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type PermissionAPIPostCheckPermissionRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	maxDepth *int64
	postCheckPermissionBody *PostCheckPermissionBody
}

func (r PermissionAPIPostCheckPermissionRequest) MaxDepth(maxDepth int64) PermissionAPIPostCheckPermissionRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPIPostCheckPermissionRequest) PostCheckPermissionBody(postCheckPermissionBody PostCheckPermissionBody) PermissionAPIPostCheckPermissionRequest {
	r.postCheckPermissionBody = &postCheckPermissionBody
	return r
}

func (r PermissionAPIPostCheckPermissionRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.PostCheckPermissionExecute(r)
}

/*
PostCheckPermission Check a permission

To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPIPostCheckPermissionRequest
*/
func (a *PermissionAPIService) PostCheckPermission(ctx context.Context) PermissionAPIPostCheckPermissionRequest {
	return PermissionAPIPostCheckPermissionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckPermissionResult
func (a *PermissionAPIService) PostCheckPermissionExecute(r PermissionAPIPostCheckPermissionRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.PostCheckPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check/openapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
	localVarPostBody = r.postCheckPermissionBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type PermissionAPIPostCheckPermissionOrErrorRequest struct {
	ctx context.Context
	ApiService PermissionAPI
	maxDepth *int64
	postCheckPermissionOrErrorBody *PostCheckPermissionOrErrorBody
}

func (r PermissionAPIPostCheckPermissionOrErrorRequest) MaxDepth(maxDepth int64) PermissionAPIPostCheckPermissionOrErrorRequest {
	r.maxDepth = &maxDepth
	return r
}

func (r PermissionAPIPostCheckPermissionOrErrorRequest) PostCheckPermissionOrErrorBody(postCheckPermissionOrErrorBody PostCheckPermissionOrErrorBody) PermissionAPIPostCheckPermissionOrErrorRequest {
	r.postCheckPermissionOrErrorBody = &postCheckPermissionOrErrorBody
	return r
}

func (r PermissionAPIPostCheckPermissionOrErrorRequest) Execute() (*CheckPermissionResult, *http.Response, error) {
	return r.ApiService.PostCheckPermissionOrErrorExecute(r)
}

/*
PostCheckPermissionOrError Check a permission

To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PermissionAPIPostCheckPermissionOrErrorRequest
*/
func (a *PermissionAPIService) PostCheckPermissionOrError(ctx context.Context) PermissionAPIPostCheckPermissionOrErrorRequest {
	return PermissionAPIPostCheckPermissionOrErrorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckPermissionResult
func (a *PermissionAPIService) PostCheckPermissionOrErrorExecute(r PermissionAPIPostCheckPermissionOrErrorRequest) (*CheckPermissionResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckPermissionResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionAPIService.PostCheckPermissionOrError")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/relation-tuples/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max-depth", r.maxDepth, "")
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
	localVarPostBody = r.postCheckPermissionOrErrorBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CheckPermissionResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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
