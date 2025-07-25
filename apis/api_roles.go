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

// RolesApi Roles service
type RolesApi struct {
	Client *HttpClient
}

type ApiCreateRoleRequest struct {
	ctx            context.Context
	ThisApiHandler *RolesApi
	role           *Role
}

func (r ApiCreateRoleRequest) Role(role Role) ApiCreateRoleRequest {
	r.role = &role
	return r
}

func (r ApiCreateRoleRequest) Execute() (*Role, *http.Response, error) {
	return r.ThisApiHandler.CreateRoleExecute(r)
}

/*
CreateRole Create role

Creates a role with permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRoleRequest
*/
func (a *RolesApi) CreateRole(ctx context.Context) ApiCreateRoleRequest {
	return ApiCreateRoleRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Role
func (a *RolesApi) CreateRoleExecute(r ApiCreateRoleRequest) (*Role, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Role
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "RolesApi.CreateRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.role == nil {
		return localVarReturnValue, nil, reportError("role is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.role
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

type ApiDeleteRoleRequest struct {
	ctx            context.Context
	ThisApiHandler *RolesApi
	roleId         string
}

func (r ApiDeleteRoleRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteRoleExecute(r)
}

/*
DeleteRole Deletes role

Remove a role. If a record references the role, that record will not be modified.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleId The identifier of the role.
	@return ApiDeleteRoleRequest
*/
func (a *RolesApi) DeleteRole(ctx context.Context, roleId string) ApiDeleteRoleRequest {
	return ApiDeleteRoleRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		roleId:         roleId,
	}
}

// Execute executes the request
func (a *RolesApi) DeleteRoleExecute(r ApiDeleteRoleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "RolesApi.DeleteRole")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.roleId) < 1 {
		return nil, reportError("roleId must have at least 1 elements")
	}
	if strlen(r.roleId) > 64 {
		return nil, reportError("roleId must have less than 64 elements")
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

type ApiGetRoleRequest struct {
	ctx            context.Context
	ThisApiHandler *RolesApi
	roleId         string
}

func (r ApiGetRoleRequest) Execute() (*Role, *http.Response, error) {
	return r.ThisApiHandler.GetRoleExecute(r)
}

/*
GetRole Retrieve role

Roles contain a list of permissions that will be applied to any user or resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleId The identifier of the role.
	@return ApiGetRoleRequest
*/
func (a *RolesApi) GetRole(ctx context.Context, roleId string) ApiGetRoleRequest {
	return ApiGetRoleRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		roleId:         roleId,
	}
}

// Execute executes the request
//
//	@return Role
func (a *RolesApi) GetRoleExecute(r ApiGetRoleRequest) (*Role, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Role
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "RolesApi.GetRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.roleId) < 1 {
		return localVarReturnValue, nil, reportError("roleId must have at least 1 elements")
	}
	if strlen(r.roleId) > 64 {
		return localVarReturnValue, nil, reportError("roleId must have less than 64 elements")
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

type ApiGetRolesRequest struct {
	ctx            context.Context
	ThisApiHandler *RolesApi
	limit          *uint8
	cursor         *string
	filter         *string
}

// Max number of results to return.
func (r ApiGetRolesRequest) Limit(limit uint8) ApiGetRolesRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging.
func (r ApiGetRolesRequest) Cursor(cursor string) ApiGetRolesRequest {
	r.cursor = &cursor
	return r
}

// Filter to search roles by. This is a case insensitive search.
func (r ApiGetRolesRequest) Filter(filter string) ApiGetRolesRequest {
	r.filter = &filter
	return r
}

func (r ApiGetRolesRequest) Execute() (*RoleCollection, *http.Response, error) {
	return r.ThisApiHandler.GetRolesExecute(r)
}

/*
GetRoles List roles

Get all the account roles. Roles contain a list of permissions that will be applied to any user or resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRolesRequest
*/
func (a *RolesApi) GetRoles(ctx context.Context) ApiGetRolesRequest {
	return ApiGetRolesRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return RoleCollection
func (a *RolesApi) GetRolesExecute(r ApiGetRolesRequest) (*RoleCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *RoleCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "RolesApi.GetRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue uint8 = 100
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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

type ApiUpdateRoleRequest struct {
	ctx            context.Context
	ThisApiHandler *RolesApi
	roleId         string
	role           *Role
}

func (r ApiUpdateRoleRequest) Role(role Role) ApiUpdateRoleRequest {
	r.role = &role
	return r
}

func (r ApiUpdateRoleRequest) Execute() (*Role, *http.Response, error) {
	return r.ThisApiHandler.UpdateRoleExecute(r)
}

/*
UpdateRole Update role

Updates a role adding or removing permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param roleId The identifier of the role.
	@return ApiUpdateRoleRequest
*/
func (a *RolesApi) UpdateRole(ctx context.Context, roleId string) ApiUpdateRoleRequest {
	return ApiUpdateRoleRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		roleId:         roleId,
	}
}

// Execute executes the request
//
//	@return Role
func (a *RolesApi) UpdateRoleExecute(r ApiUpdateRoleRequest) (*Role, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Role
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "RolesApi.UpdateRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/roles/{roleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.roleId) < 1 {
		return localVarReturnValue, nil, reportError("roleId must have at least 1 elements")
	}
	if strlen(r.roleId) > 64 {
		return localVarReturnValue, nil, reportError("roleId must have less than 64 elements")
	}
	if r.role == nil {
		return localVarReturnValue, nil, reportError("role is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.role
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
