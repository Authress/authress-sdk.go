package authress

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GroupsApi Groups service
type GroupsApi service

type ApiCreateGroupRequest struct {
	ctx            context.Context
	ThisApiHandler *GroupsApi
	group          *Group
}

func (r ApiCreateGroupRequest) Group(group Group) ApiCreateGroupRequest {
	r.group = &group
	return r
}

func (r ApiCreateGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ThisApiHandler.CreateGroupExecute(r)
}

/*
CreateGroup Create group

Specify users to be included in a new group. (Groups have a maximum size of ~100KB)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateGroupRequest
*/
func (a *GroupsApi) CreateGroup(ctx context.Context) ApiCreateGroupRequest {
	return ApiCreateGroupRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Group
func (a *GroupsApi) CreateGroupExecute(r ApiCreateGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApi.CreateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.group == nil {
		return localVarReturnValue, nil, reportError("group is required and must be specified")
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
	localVarPostBody = r.group
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

type ApiDeleteGroupRequest struct {
	ctx            context.Context
	ThisApiHandler *GroupsApi
	groupId        GroupId
}

func (r ApiDeleteGroupRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteGroupExecute(r)
}

/*
DeleteGroup Deletes group

Remove a group, users will lose any role that was assigned through membership of this group. This action cannot be undone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The identifier of the group.
	@return ApiDeleteGroupRequest
*/
func (a *GroupsApi) DeleteGroup(ctx context.Context, groupId GroupId) ApiDeleteGroupRequest {
	return ApiDeleteGroupRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		groupId:        groupId,
	}
}

// Execute executes the request
func (a *GroupsApi) DeleteGroupExecute(r ApiDeleteGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApi.DeleteGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetGroupRequest struct {
	ctx            context.Context
	ThisApiHandler *GroupsApi
	groupId        GroupId
}

func (r ApiGetGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ThisApiHandler.GetGroupExecute(r)
}

/*
GetGroup Retrieve group

A group contains multiple users which can be added to an access record, and should be assigned the same roles at the same time.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The identifier of the group.
	@return ApiGetGroupRequest
*/
func (a *GroupsApi) GetGroup(ctx context.Context, groupId GroupId) ApiGetGroupRequest {
	return ApiGetGroupRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		groupId:        groupId,
	}
}

// Execute executes the request
//
//	@return Group
func (a *GroupsApi) GetGroupExecute(r ApiGetGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApi.GetGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

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

type ApiGetGroupsRequest struct {
	ctx            context.Context
	ThisApiHandler *GroupsApi
	limit          *int32
	cursor         *string
	filter         *string
}

// Max number of results to return
func (r ApiGetGroupsRequest) Limit(limit int32) ApiGetGroupsRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetGroupsRequest) Cursor(cursor string) ApiGetGroupsRequest {
	r.cursor = &cursor
	return r
}

// Filter to search groups by. This is a case insensitive search through every text field.
func (r ApiGetGroupsRequest) Filter(filter string) ApiGetGroupsRequest {
	r.filter = &filter
	return r
}

func (r ApiGetGroupsRequest) Execute() (*GroupCollection, *http.Response, error) {
	return r.ThisApiHandler.GetGroupsExecute(r)
}

/*
GetGroups List groups

Returns a paginated groups list for the account. Only groups the user has access to are returned. The results sort order is not stable, on subsequent requests different results may be returned. This query resource is meant for administrative actions only, therefore has a lower rate limit tier than user permissions related resources. Additionally, the results from a query to Groups may be delayed up to 5 minutes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetGroupsRequest
*/
func (a *GroupsApi) GetGroups(ctx context.Context) ApiGetGroupsRequest {
	return ApiGetGroupsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return GroupCollection
func (a *GroupsApi) GetGroupsExecute(r ApiGetGroupsRequest) (*GroupCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GroupCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApi.GetGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 20
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

type ApiUpdateGroupRequest struct {
	ctx               context.Context
	ThisApiHandler    *GroupsApi
	groupId           GroupId
	group             *Group
	ifUnmodifiedSince *time.Time
}

func (r ApiUpdateGroupRequest) Group(group Group) ApiUpdateGroupRequest {
	r.group = &group
	return r
}

// The expected last time the group was modified.
func (r ApiUpdateGroupRequest) IfUnmodifiedSince(ifUnmodifiedSince time.Time) ApiUpdateGroupRequest {
	r.ifUnmodifiedSince = &ifUnmodifiedSince
	return r
}

func (r ApiUpdateGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ThisApiHandler.UpdateGroupExecute(r)
}

/*
UpdateGroup Update a group

Updates a group adding or removing user. Change a group updates the permissions and roles the users have access to. (Groups have a maximum size of ~100KB)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId The identifier of the group.
	@return ApiUpdateGroupRequest
*/
func (a *GroupsApi) UpdateGroup(ctx context.Context, groupId GroupId) ApiUpdateGroupRequest {
	return ApiUpdateGroupRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		groupId:        groupId,
	}
}

// Execute executes the request
//
//	@return Group
func (a *GroupsApi) UpdateGroupExecute(r ApiUpdateGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApi.UpdateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/groups/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.group == nil {
		return localVarReturnValue, nil, reportError("group is required and must be specified")
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
	if r.ifUnmodifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Unmodified-Since", r.ifUnmodifiedSince, "")
	}
	// body params
	localVarPostBody = r.group
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
