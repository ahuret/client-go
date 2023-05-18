/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.31
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
)


type RelationshipApi interface {

	/*
	CheckOplSyntax Check the syntax of an OPL file

	The OPL file is expected in the body of the request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiCheckOplSyntaxRequest
	*/
	CheckOplSyntax(ctx context.Context) RelationshipApiCheckOplSyntaxRequest

	// CheckOplSyntaxExecute executes the request
	//  @return CheckOplSyntaxResult
	CheckOplSyntaxExecute(r RelationshipApiCheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error)

	/*
	CreateRelationship Create a Relationship

	Use this endpoint to create a relationship.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiCreateRelationshipRequest
	*/
	CreateRelationship(ctx context.Context) RelationshipApiCreateRelationshipRequest

	// CreateRelationshipExecute executes the request
	//  @return Relationship
	CreateRelationshipExecute(r RelationshipApiCreateRelationshipRequest) (*Relationship, *http.Response, error)

	/*
	DeleteRelationships Delete Relationships

	Use this endpoint to delete relationships

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiDeleteRelationshipsRequest
	*/
	DeleteRelationships(ctx context.Context) RelationshipApiDeleteRelationshipsRequest

	// DeleteRelationshipsExecute executes the request
	DeleteRelationshipsExecute(r RelationshipApiDeleteRelationshipsRequest) (*http.Response, error)

	/*
	GetRelationships Query relationships

	Get all relationships that match the query. Only the namespace field is required.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiGetRelationshipsRequest
	*/
	GetRelationships(ctx context.Context) RelationshipApiGetRelationshipsRequest

	// GetRelationshipsExecute executes the request
	//  @return Relationships
	GetRelationshipsExecute(r RelationshipApiGetRelationshipsRequest) (*Relationships, *http.Response, error)

	/*
	ListRelationshipNamespaces Query namespaces

	Get all namespaces

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiListRelationshipNamespacesRequest
	*/
	ListRelationshipNamespaces(ctx context.Context) RelationshipApiListRelationshipNamespacesRequest

	// ListRelationshipNamespacesExecute executes the request
	//  @return RelationshipNamespaces
	ListRelationshipNamespacesExecute(r RelationshipApiListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error)

	/*
	PatchRelationships Patch Multiple Relationships

	Use this endpoint to patch one or more relationships.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RelationshipApiPatchRelationshipsRequest
	*/
	PatchRelationships(ctx context.Context) RelationshipApiPatchRelationshipsRequest

	// PatchRelationshipsExecute executes the request
	PatchRelationshipsExecute(r RelationshipApiPatchRelationshipsRequest) (*http.Response, error)
}

// RelationshipApiService RelationshipApi service
type RelationshipApiService service

type RelationshipApiCheckOplSyntaxRequest struct {
	ctx context.Context
	ApiService RelationshipApi
	body *string
}

func (r RelationshipApiCheckOplSyntaxRequest) Body(body string) RelationshipApiCheckOplSyntaxRequest {
	r.body = &body
	return r
}

func (r RelationshipApiCheckOplSyntaxRequest) Execute() (*CheckOplSyntaxResult, *http.Response, error) {
	return r.ApiService.CheckOplSyntaxExecute(r)
}

/*
CheckOplSyntax Check the syntax of an OPL file

The OPL file is expected in the body of the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiCheckOplSyntaxRequest
*/
func (a *RelationshipApiService) CheckOplSyntax(ctx context.Context) RelationshipApiCheckOplSyntaxRequest {
	return RelationshipApiCheckOplSyntaxRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckOplSyntaxResult
func (a *RelationshipApiService) CheckOplSyntaxExecute(r RelationshipApiCheckOplSyntaxRequest) (*CheckOplSyntaxResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckOplSyntaxResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.CheckOplSyntax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/opl/syntax/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type RelationshipApiCreateRelationshipRequest struct {
	ctx context.Context
	ApiService RelationshipApi
	createRelationshipBody *CreateRelationshipBody
}

func (r RelationshipApiCreateRelationshipRequest) CreateRelationshipBody(createRelationshipBody CreateRelationshipBody) RelationshipApiCreateRelationshipRequest {
	r.createRelationshipBody = &createRelationshipBody
	return r
}

func (r RelationshipApiCreateRelationshipRequest) Execute() (*Relationship, *http.Response, error) {
	return r.ApiService.CreateRelationshipExecute(r)
}

/*
CreateRelationship Create a Relationship

Use this endpoint to create a relationship.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiCreateRelationshipRequest
*/
func (a *RelationshipApiService) CreateRelationship(ctx context.Context) RelationshipApiCreateRelationshipRequest {
	return RelationshipApiCreateRelationshipRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Relationship
func (a *RelationshipApiService) CreateRelationshipExecute(r RelationshipApiCreateRelationshipRequest) (*Relationship, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Relationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.CreateRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

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
	localVarPostBody = r.createRelationshipBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type RelationshipApiDeleteRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipApi
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

// Namespace of the Relationship
func (r RelationshipApiDeleteRelationshipsRequest) Namespace(namespace string) RelationshipApiDeleteRelationshipsRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r RelationshipApiDeleteRelationshipsRequest) Object(object string) RelationshipApiDeleteRelationshipsRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r RelationshipApiDeleteRelationshipsRequest) Relation(relation string) RelationshipApiDeleteRelationshipsRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r RelationshipApiDeleteRelationshipsRequest) SubjectId(subjectId string) RelationshipApiDeleteRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r RelationshipApiDeleteRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipApiDeleteRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r RelationshipApiDeleteRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipApiDeleteRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r RelationshipApiDeleteRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipApiDeleteRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipApiDeleteRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRelationshipsExecute(r)
}

/*
DeleteRelationships Delete Relationships

Use this endpoint to delete relationships

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiDeleteRelationshipsRequest
*/
func (a *RelationshipApiService) DeleteRelationships(ctx context.Context) RelationshipApiDeleteRelationshipsRequest {
	return RelationshipApiDeleteRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RelationshipApiService) DeleteRelationshipsExecute(r RelationshipApiDeleteRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.DeleteRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type RelationshipApiGetRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipApi
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

func (r RelationshipApiGetRelationshipsRequest) PageToken(pageToken string) RelationshipApiGetRelationshipsRequest {
	r.pageToken = &pageToken
	return r
}

func (r RelationshipApiGetRelationshipsRequest) PageSize(pageSize int64) RelationshipApiGetRelationshipsRequest {
	r.pageSize = &pageSize
	return r
}

// Namespace of the Relationship
func (r RelationshipApiGetRelationshipsRequest) Namespace(namespace string) RelationshipApiGetRelationshipsRequest {
	r.namespace = &namespace
	return r
}

// Object of the Relationship
func (r RelationshipApiGetRelationshipsRequest) Object(object string) RelationshipApiGetRelationshipsRequest {
	r.object = &object
	return r
}

// Relation of the Relationship
func (r RelationshipApiGetRelationshipsRequest) Relation(relation string) RelationshipApiGetRelationshipsRequest {
	r.relation = &relation
	return r
}

// SubjectID of the Relationship
func (r RelationshipApiGetRelationshipsRequest) SubjectId(subjectId string) RelationshipApiGetRelationshipsRequest {
	r.subjectId = &subjectId
	return r
}

// Namespace of the Subject Set
func (r RelationshipApiGetRelationshipsRequest) SubjectSetNamespace(subjectSetNamespace string) RelationshipApiGetRelationshipsRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}

// Object of the Subject Set
func (r RelationshipApiGetRelationshipsRequest) SubjectSetObject(subjectSetObject string) RelationshipApiGetRelationshipsRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}

// Relation of the Subject Set
func (r RelationshipApiGetRelationshipsRequest) SubjectSetRelation(subjectSetRelation string) RelationshipApiGetRelationshipsRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r RelationshipApiGetRelationshipsRequest) Execute() (*Relationships, *http.Response, error) {
	return r.ApiService.GetRelationshipsExecute(r)
}

/*
GetRelationships Query relationships

Get all relationships that match the query. Only the namespace field is required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiGetRelationshipsRequest
*/
func (a *RelationshipApiService) GetRelationships(ctx context.Context) RelationshipApiGetRelationshipsRequest {
	return RelationshipApiGetRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Relationships
func (a *RelationshipApiService) GetRelationshipsExecute(r RelationshipApiGetRelationshipsRequest) (*Relationships, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Relationships
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.GetRelationships")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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

type RelationshipApiListRelationshipNamespacesRequest struct {
	ctx context.Context
	ApiService RelationshipApi
}

func (r RelationshipApiListRelationshipNamespacesRequest) Execute() (*RelationshipNamespaces, *http.Response, error) {
	return r.ApiService.ListRelationshipNamespacesExecute(r)
}

/*
ListRelationshipNamespaces Query namespaces

Get all namespaces

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiListRelationshipNamespacesRequest
*/
func (a *RelationshipApiService) ListRelationshipNamespaces(ctx context.Context) RelationshipApiListRelationshipNamespacesRequest {
	return RelationshipApiListRelationshipNamespacesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RelationshipNamespaces
func (a *RelationshipApiService) ListRelationshipNamespacesExecute(r RelationshipApiListRelationshipNamespacesRequest) (*RelationshipNamespaces, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RelationshipNamespaces
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.ListRelationshipNamespaces")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/namespaces"

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
			var v ErrorGeneric
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

type RelationshipApiPatchRelationshipsRequest struct {
	ctx context.Context
	ApiService RelationshipApi
	relationshipPatch *[]RelationshipPatch
}

func (r RelationshipApiPatchRelationshipsRequest) RelationshipPatch(relationshipPatch []RelationshipPatch) RelationshipApiPatchRelationshipsRequest {
	r.relationshipPatch = &relationshipPatch
	return r
}

func (r RelationshipApiPatchRelationshipsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchRelationshipsExecute(r)
}

/*
PatchRelationships Patch Multiple Relationships

Use this endpoint to patch one or more relationships.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RelationshipApiPatchRelationshipsRequest
*/
func (a *RelationshipApiService) PatchRelationships(ctx context.Context) RelationshipApiPatchRelationshipsRequest {
	return RelationshipApiPatchRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RelationshipApiService) PatchRelationshipsExecute(r RelationshipApiPatchRelationshipsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipApiService.PatchRelationships")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

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
	localVarPostBody = r.relationshipPatch
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v ErrorGeneric
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
