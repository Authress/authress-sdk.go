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

// ResourcePermissionsApi ResourcePermissions service
type ResourcePermissionsApi struct {
	Client *HttpClient
}

type ApiGetPermissionedResourceRequest struct {
	ctx            context.Context
	ThisApiHandler *ResourcePermissionsApi
	resourceUri    string
}

func (r ApiGetPermissionedResourceRequest) Execute() (*PermissionedResource, *http.Response, error) {
	return r.ThisApiHandler.GetPermissionedResourceExecute(r)
}

/*
GetPermissionedResource Retrieve resource configuration

Permissions can be set globally at a resource level. This will apply to all users in an account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceUri The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	@return ApiGetPermissionedResourceRequest
*/
func (a *ResourcePermissionsApi) GetPermissionedResource(ctx context.Context, resourceUri string) ApiGetPermissionedResourceRequest {
	return ApiGetPermissionedResourceRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		resourceUri:    resourceUri,
	}
}

// Execute executes the request
//
//	@return PermissionedResource
func (a *ResourcePermissionsApi) GetPermissionedResourceExecute(r ApiGetPermissionedResourceRequest) (*PermissionedResource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *PermissionedResource
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "ResourcePermissionsApi.GetPermissionedResource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/resources/{resourceUri}"
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

type ApiGetPermissionedResourcesRequest struct {
	ctx            context.Context
	ThisApiHandler *ResourcePermissionsApi
}

func (r ApiGetPermissionedResourcesRequest) Execute() (*PermissionedResourceCollection, *http.Response, error) {
	return r.ThisApiHandler.GetPermissionedResourcesExecute(r)
}

/*
GetPermissionedResources List all resource configurations

Permissions can be set globally at a resource level. This endpoint returns a list of resources with globally set resource permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPermissionedResourcesRequest
*/
func (a *ResourcePermissionsApi) GetPermissionedResources(ctx context.Context) ApiGetPermissionedResourcesRequest {
	return ApiGetPermissionedResourcesRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return PermissionedResourceCollection
func (a *ResourcePermissionsApi) GetPermissionedResourcesExecute(r ApiGetPermissionedResourcesRequest) (*PermissionedResourceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *PermissionedResourceCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "ResourcePermissionsApi.GetPermissionedResources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/resources"

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

type ApiGetResourceUsersRequest struct {
	ctx            context.Context
	ThisApiHandler *ResourcePermissionsApi
	resourceUri    string
	limit          *uint8
	cursor         *string
}

// Max number of results to return
func (r ApiGetResourceUsersRequest) Limit(limit uint8) ApiGetResourceUsersRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetResourceUsersRequest) Cursor(cursor string) ApiGetResourceUsersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetResourceUsersRequest) Execute() (*ResourceUsersCollection, *http.Response, error) {
	return r.ThisApiHandler.GetResourceUsersExecute(r)
}

/*
GetResourceUsers List users with resource access

Get the resource users with explicit access to the resource. This result is a list of users that have some permission to the resource. Users with access to parent resources and users with access only to a sub-resource will not be returned in this result. In the case that the resource has multiple users, the list will be paginated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceUri The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	@return ApiGetResourceUsersRequest
*/
func (a *ResourcePermissionsApi) GetResourceUsers(ctx context.Context, resourceUri string) ApiGetResourceUsersRequest {
	return ApiGetResourceUsersRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		resourceUri:    resourceUri,
	}
}

// Execute executes the request
//
//	@return ResourceUsersCollection
func (a *ResourcePermissionsApi) GetResourceUsersExecute(r ApiGetResourceUsersRequest) (*ResourceUsersCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *ResourceUsersCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "ResourcePermissionsApi.GetResourceUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/resources/{resourceUri}/users"
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

type ApiUpdatePermissionedResourceRequest struct {
	ctx                  context.Context
	ThisApiHandler       *ResourcePermissionsApi
	resourceUri          string
	permissionedResource *PermissionedResource
}

// The contents of the permission to set on the resource. Overwrites existing data.
func (r ApiUpdatePermissionedResourceRequest) PermissionedResource(permissionedResource PermissionedResource) ApiUpdatePermissionedResourceRequest {
	r.permissionedResource = &permissionedResource
	return r
}

func (r ApiUpdatePermissionedResourceRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.UpdatePermissionedResourceExecute(r)
}

/*
UpdatePermissionedResource Update resource configuration

Updates the global permissions on a resource. This applies to all users in an account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceUri The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	@return ApiUpdatePermissionedResourceRequest
*/
func (a *ResourcePermissionsApi) UpdatePermissionedResource(ctx context.Context, resourceUri string) ApiUpdatePermissionedResourceRequest {
	return ApiUpdatePermissionedResourceRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		resourceUri:    resourceUri,
	}
}

// Execute executes the request
func (a *ResourcePermissionsApi) UpdatePermissionedResourceExecute(r ApiUpdatePermissionedResourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "ResourcePermissionsApi.UpdatePermissionedResource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/resources/{resourceUri}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceUri"+"}", url.PathEscape(parameterValueToString(r.resourceUri, "resourceUri")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.resourceUri) < 1 {
		return nil, reportError("resourceUri must have at least 1 elements")
	}
	if strlen(r.resourceUri) > 512 {
		return nil, reportError("resourceUri must have less than 512 elements")
	}
	if r.permissionedResource == nil {
		return nil, reportError("permissionedResource is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.permissionedResource
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
