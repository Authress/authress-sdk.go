package apis

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	. "github.com/authress/authress-sdk.go/models"
	. "github.com/authress/authress-sdk.go/utilities"
)

// UserPermissionsApi UserPermissions service
type UserPermissionsApi struct {
	Client *HttpClient
}

type ApiAuthorizeUserRequest struct {
	ctx            context.Context
	ThisApiHandler *UserPermissionsApi
	userId         UserId
	resourceUri    string
	permission     Action
}

func (r ApiAuthorizeUserRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.AuthorizeUserExecute(r)
}

/*
AuthorizeUser Verify user authorization

Performs the user authorization check. Does the user have the specified permission to the resource?

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The user identifier to check and verify the permissions of.
	@param resourceUri The uri path of a resource to validate, must be URL encoded, uri segments are allowed, the resource must be a full path.
	@param permission Permission to check, '*' and scoped permissions can also be checked here.
	@return ApiAuthorizeUserRequest
*/
func (a *UserPermissionsApi) AuthorizeUser(ctx context.Context, userId UserId, resourceUri string, permission Action) ApiAuthorizeUserRequest {
	return ApiAuthorizeUserRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		userId:         userId,
		resourceUri:    resourceUri,
		permission:     permission,
	}
}

// Execute executes the request
func (a *UserPermissionsApi) AuthorizeUserExecute(r ApiAuthorizeUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "UserPermissionsApi.AuthorizeUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/resources/{resourceUri}/permissions/{permission}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUri"+"}", url.PathEscape(parameterValueToString(r.resourceUri, "resourceUri")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"permission"+"}", url.PathEscape(parameterValueToString(r.permission, "permission")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.resourceUri) < 1 {
		return nil, reportError("resourceUri must have at least 1 elements")
	}
	if strlen(r.resourceUri) > 512 {
		return nil, reportError("resourceUri must have less than 512 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(req)
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetUserPermissionsForResourceRequest struct {
	ctx            context.Context
	ThisApiHandler *UserPermissionsApi
	userId         UserId
	resourceUri    string
}

func (r ApiGetUserPermissionsForResourceRequest) Execute() (*PermissionCollection, *http.Response, error) {
	return r.ThisApiHandler.GetUserPermissionsForResourceExecute(r)
}

/*
GetUserPermissionsForResource Get user permissions for resource

Get a summary of the permissions a user has to a particular resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The user identifier for the user to check permissions.
	@param resourceUri The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	@return ApiGetUserPermissionsForResourceRequest
*/
func (a *UserPermissionsApi) GetUserPermissionsForResource(ctx context.Context, userId UserId, resourceUri string) ApiGetUserPermissionsForResourceRequest {
	return ApiGetUserPermissionsForResourceRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		userId:         userId,
		resourceUri:    resourceUri,
	}
}

// Execute executes the request
//
//	@return PermissionCollection
func (a *UserPermissionsApi) GetUserPermissionsForResourceExecute(r ApiGetUserPermissionsForResourceRequest) (*PermissionCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *PermissionCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "UserPermissionsApi.GetUserPermissionsForResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/resources/{resourceUri}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUri"+"}", url.PathEscape(parameterValueToString(r.resourceUri, "resourceUri")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.resourceUri) < 1 {
		return localVarReturnValue, nil, reportError("resourceUri must have at least 1 elements")
	}
	if strlen(r.resourceUri) > 512 {
		return localVarReturnValue, nil, reportError("resourceUri must have less than 512 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/links+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(req)
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
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetUserResourcesRequest struct {
	ctx                     context.Context
	ThisApiHandler          *UserPermissionsApi
	userId                  UserId
	resourceUri             *string
	collectionConfiguration *string
	permissions             *Action
	limit                   *uint8
	cursor                  *string
}

// The top level uri path of a resource to query for. Will only match explicit or nested sub-resources. Will not partial match resource names.
func (r ApiGetUserResourcesRequest) ResourceUri(resourceUri string) ApiGetUserResourcesRequest {
	r.resourceUri = &resourceUri
	return r
}

// &#x60;TOP_LEVEL_ONLY&#x60; - returns only directly nested resources under the resourceUri. A query to &#x60;resourceUri&#x3D;Collection&#x60; will return &#x60;Collection/resource_1&#x60;.&lt;br&gt;&#x60;INCLUDE_NESTED&#x60; - will return all sub-resources as well as deeply nested resources that the user has the specified permission to. A query to &#x60;resourceUri&#x3D;Collection&#x60; will return &#x60;Collection/namespaces/ns/resources/resource_1&#x60;.&lt;br&gt;&lt;br&gt;To return matching resources for nested resources, set this parameter to &#x60;INCLUDE_NESTED&#x60;.
func (r ApiGetUserResourcesRequest) CollectionConfiguration(collectionConfiguration string) ApiGetUserResourcesRequest {
	r.collectionConfiguration = &collectionConfiguration
	return r
}

// Permission to check, &#39;*&#39; and scoped permissions can also be checked here. By default if the user has any permission explicitly to a resource, it will be included in the list.
func (r ApiGetUserResourcesRequest) Permissions(permissions Action) ApiGetUserResourcesRequest {
	r.permissions = &permissions
	return r
}

// Max number of results to return
func (r ApiGetUserResourcesRequest) Limit(limit uint8) ApiGetUserResourcesRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetUserResourcesRequest) Cursor(cursor string) ApiGetUserResourcesRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetUserResourcesRequest) Execute() (*UserResourcesCollection, *http.Response, error) {
	return r.ThisApiHandler.GetUserResourcesExecute(r)
}

/*
GetUserResources List user resources

Get the users resources. This result is a list of resource uris that a user has an permission to. By default only the top level matching resources are returned. To get a user's list of deeply nested resources, set the `collectionConfiguration` to be `INCLUDE_NESTED`. This collection is paginated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The user identifier for which to list all accessible resources.
	@return ApiGetUserResourcesRequest
*/
func (a *UserPermissionsApi) GetUserResources(ctx context.Context, userId UserId) ApiGetUserResourcesRequest {
	return ApiGetUserResourcesRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		userId:         userId,
	}
}

// Execute executes the request
//
//	@return UserResourcesCollection
func (a *UserPermissionsApi) GetUserResourcesExecute(r ApiGetUserResourcesRequest) (*UserResourcesCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *UserResourcesCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "UserPermissionsApi.GetUserResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/resources"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resourceUri != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resourceUri", r.resourceUri, "")
	}
	if r.collectionConfiguration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "collectionConfiguration", r.collectionConfiguration, "")
	} else {
		var defaultValue string = "TOP_LEVEL_ONLY"
		r.collectionConfiguration = &defaultValue
	}
	if r.permissions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "permissions", r.permissions, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue uint8 = 20
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/links+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(req)
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
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetUserRolesForResourceRequest struct {
	ctx            context.Context
	ThisApiHandler *UserPermissionsApi
	userId         UserId
	resourceUri    string
}

func (r ApiGetUserRolesForResourceRequest) Execute() (*UserRoleCollection, *http.Response, error) {
	return r.ThisApiHandler.GetUserRolesForResourceExecute(r)
}

/*
GetUserRolesForResource Get user roles for resource

Get a summary of the roles a user has to a particular resource. Users can be assigned roles from multiple access records, this may cause the same role to appear in the list more than once.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId The user to get roles for.
	@param resourceUri The uri path of a resource to get roles for, must be URL encoded. Checks for explicit resource roles, roles attached to parent resources are not returned.
	@return ApiGetUserRolesForResourceRequest
*/
func (a *UserPermissionsApi) GetUserRolesForResource(ctx context.Context, userId UserId, resourceUri string) ApiGetUserRolesForResourceRequest {
	return ApiGetUserRolesForResourceRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		userId:         userId,
		resourceUri:    resourceUri,
	}
}

// Execute executes the request
//
//	@return UserRoleCollection
func (a *UserPermissionsApi) GetUserRolesForResourceExecute(r ApiGetUserRolesForResourceRequest) (*UserRoleCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *UserRoleCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "UserPermissionsApi.GetUserRolesForResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/users/{userId}/resources/{resourceUri}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUri"+"}", url.PathEscape(parameterValueToString(r.resourceUri, "resourceUri")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.resourceUri) < 1 {
		return localVarReturnValue, nil, reportError("resourceUri must have at least 1 elements")
	}
	if strlen(r.resourceUri) > 512 {
		return localVarReturnValue, nil, reportError("resourceUri must have less than 512 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/links+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.callAPI(req)
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
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
