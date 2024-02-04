package apis

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "github.com/authress/authress-sdk.go/models"
	. "github.com/authress/authress-sdk.go/utilities"
)

// AccessRecordsApi AccessRecords service
type AccessRecordsApi struct {
	Client *HttpClient
}

type ApiCreateClaimRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	claimRequest   *ClaimRequest
}

func (r ApiCreateClaimRequest) ClaimRequest(claimRequest ClaimRequest) ApiCreateClaimRequest {
	r.claimRequest = &claimRequest
	return r
}

func (r ApiCreateClaimRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ThisApiHandler.CreateClaimExecute(r)
}

/*
CreateClaim Create resource Claim

Claim a resource by allowing a user to pick an identifier and receive admin access to that resource if it hasn't already been claimed. This only works for resources specifically marked as <strong>CLAIM</strong>. The result will be a new access record listing that user as the admin for this resource. The resourceUri will be appended to the collection resource uri using a `/` (forward slash) automatically.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateClaimRequest
*/
func (a *AccessRecordsApi) CreateClaim(ctx context.Context) ApiCreateClaimRequest {
	return ApiCreateClaimRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AccessRecordsApi) CreateClaimExecute(r ApiCreateClaimRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.CreateClaim")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/claims"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.claimRequest == nil {
		return localVarReturnValue, nil, reportError("claimRequest is required and must be specified")
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
	localVarPostBody = r.claimRequest
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

type ApiCreateRecordRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	accessRecord   *AccessRecord
}

func (r ApiCreateRecordRequest) AccessRecord(accessRecord AccessRecord) ApiCreateRecordRequest {
	r.accessRecord = &accessRecord
	return r
}

func (r ApiCreateRecordRequest) Execute() (*AccessRecord, *http.Response, error) {
	return r.ThisApiHandler.CreateRecordExecute(r)
}

/*
CreateRecord Create access record

Specify user roles for specific resources. (Records have a maximum size of ~100KB)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRecordRequest
*/
func (a *AccessRecordsApi) CreateRecord(ctx context.Context) ApiCreateRecordRequest {
	return ApiCreateRecordRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return AccessRecord
func (a *AccessRecordsApi) CreateRecordExecute(r ApiCreateRecordRequest) (*AccessRecord, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRecord
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.CreateRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/records"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessRecord == nil {
		return localVarReturnValue, nil, reportError("accessRecord is required and must be specified")
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
	localVarPostBody = r.accessRecord
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

type ApiCreateRequestRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	accessRequest  *AccessRequest
}

func (r ApiCreateRequestRequest) AccessRequest(accessRequest AccessRequest) ApiCreateRequestRequest {
	r.accessRequest = &accessRequest
	return r
}

func (r ApiCreateRequestRequest) Execute() (*AccessRequest, *http.Response, error) {
	return r.ThisApiHandler.CreateRequestExecute(r)
}

/*
CreateRequest Create access request

Specify a request in the form of an access record that an admin can approve.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRequestRequest
*/
func (a *AccessRecordsApi) CreateRequest(ctx context.Context) ApiCreateRequestRequest {
	return ApiCreateRequestRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return AccessRequest
func (a *AccessRecordsApi) CreateRequestExecute(r ApiCreateRequestRequest) (*AccessRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRequest
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.CreateRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessRequest == nil {
		return localVarReturnValue, nil, reportError("accessRequest is required and must be specified")
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
	localVarPostBody = r.accessRequest
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

type ApiDeleteRecordRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	recordId       string
}

func (r ApiDeleteRecordRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteRecordExecute(r)
}

/*
DeleteRecord Deletes access record

Remove an access record, removing associated permissions from all users in record. If a user has a permission from another record, that permission will not be removed. (Note: This disables the record by changing the status to <strong>DELETED</strong> but not completely remove the record for tracking purposes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordId The identifier of the access record.
	@return ApiDeleteRecordRequest
*/
func (a *AccessRecordsApi) DeleteRecord(ctx context.Context, recordId string) ApiDeleteRecordRequest {
	return ApiDeleteRecordRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		recordId:       recordId,
	}
}

// Execute executes the request
func (a *AccessRecordsApi) DeleteRecordExecute(r ApiDeleteRecordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.DeleteRecord")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/records/{recordId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterValueToString(r.recordId, "recordId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.recordId) < 1 {
		return nil, reportError("recordId must have at least 1 elements")
	}
	if strlen(r.recordId) > 100 {
		return nil, reportError("recordId must have less than 100 elements")
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

type ApiDeleteRequestRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	requestId      string
}

func (r ApiDeleteRequestRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteRequestExecute(r)
}

/*
DeleteRequest Deletes access request

Remove an access request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestId The identifier of the access request.
	@return ApiDeleteRequestRequest
*/
func (a *AccessRecordsApi) DeleteRequest(ctx context.Context, requestId string) ApiDeleteRequestRequest {
	return ApiDeleteRequestRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		requestId:      requestId,
	}
}

// Execute executes the request
func (a *AccessRecordsApi) DeleteRequestExecute(r ApiDeleteRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.DeleteRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.requestId) < 1 {
		return nil, reportError("requestId must have at least 1 elements")
	}
	if strlen(r.requestId) > 100 {
		return nil, reportError("requestId must have less than 100 elements")
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

type ApiGetRecordRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	recordId       string
}

func (r ApiGetRecordRequest) Execute() (*AccessRecord, *http.Response, error) {
	return r.ThisApiHandler.GetRecordExecute(r)
}

/*
GetRecord Retrieve access record

Access records contain information assigning permissions to users for resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordId The identifier of the access record.
	@return ApiGetRecordRequest
*/
func (a *AccessRecordsApi) GetRecord(ctx context.Context, recordId string) ApiGetRecordRequest {
	return ApiGetRecordRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		recordId:       recordId,
	}
}

// Execute executes the request
//
//	@return AccessRecord
func (a *AccessRecordsApi) GetRecordExecute(r ApiGetRecordRequest) (*AccessRecord, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRecord
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.GetRecord")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/records/{recordId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterValueToString(r.recordId, "recordId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.recordId) < 1 {
		return localVarReturnValue, nil, reportError("recordId must have at least 1 elements")
	}
	if strlen(r.recordId) > 100 {
		return localVarReturnValue, nil, reportError("recordId must have less than 100 elements")
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

type ApiGetRecordsRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	limit          *int32
	cursor         *string
	filter         *string
	status         *string
}

// Max number of results to return
func (r ApiGetRecordsRequest) Limit(limit int32) ApiGetRecordsRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetRecordsRequest) Cursor(cursor string) ApiGetRecordsRequest {
	r.cursor = &cursor
	return r
}

// Filter to search records by. This is a case insensitive search through every text field.
func (r ApiGetRecordsRequest) Filter(filter string) ApiGetRecordsRequest {
	r.filter = &filter
	return r
}

// Filter records by their current status.
func (r ApiGetRecordsRequest) Status(status string) ApiGetRecordsRequest {
	r.status = &status
	return r
}

func (r ApiGetRecordsRequest) Execute() (*AccessRecordCollection, *http.Response, error) {
	return r.ThisApiHandler.GetRecordsExecute(r)
}

/*
GetRecords List access records

Returns a paginated records list for the account. Only records the user has access to are returned. The results sort order is not stable, on subsequent requests different results may be returned. This query resource is meant for administrative actions only, therefore has a lower rate limit tier than user permissions related resources. Additionally, the results from a query to Access Records may be delayed up to 5 minutes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRecordsRequest
*/
func (a *AccessRecordsApi) GetRecords(ctx context.Context) ApiGetRecordsRequest {
	return ApiGetRecordsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return AccessRecordCollection
func (a *AccessRecordsApi) GetRecordsExecute(r ApiGetRecordsRequest) (*AccessRecordCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRecordCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.GetRecords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/records"

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
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
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

type ApiGetRequestRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	requestId      string
}

func (r ApiGetRequestRequest) Execute() (*AccessRequest, *http.Response, error) {
	return r.ThisApiHandler.GetRequestExecute(r)
}

/*
GetRequest Retrieve access request

Access request contain information requesting permissions for users to resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestId The identifier of the access request.
	@return ApiGetRequestRequest
*/
func (a *AccessRecordsApi) GetRequest(ctx context.Context, requestId string) ApiGetRequestRequest {
	return ApiGetRequestRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		requestId:      requestId,
	}
}

// Execute executes the request
//
//	@return AccessRequest
func (a *AccessRecordsApi) GetRequestExecute(r ApiGetRequestRequest) (*AccessRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRequest
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.GetRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.requestId) < 1 {
		return localVarReturnValue, nil, reportError("requestId must have at least 1 elements")
	}
	if strlen(r.requestId) > 100 {
		return localVarReturnValue, nil, reportError("requestId must have less than 100 elements")
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

type ApiGetRequestsRequest struct {
	ctx            context.Context
	ThisApiHandler *AccessRecordsApi
	limit          *int32
	cursor         *string
	status         *string
}

// Max number of results to return
func (r ApiGetRequestsRequest) Limit(limit int32) ApiGetRequestsRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetRequestsRequest) Cursor(cursor string) ApiGetRequestsRequest {
	r.cursor = &cursor
	return r
}

// Filter requests by their current status.
func (r ApiGetRequestsRequest) Status(status string) ApiGetRequestsRequest {
	r.status = &status
	return r
}

func (r ApiGetRequestsRequest) Execute() (*AccessRequestCollection, *http.Response, error) {
	return r.ThisApiHandler.GetRequestsExecute(r)
}

/*
GetRequests List access requests

Returns a paginated request list. Only requests the user has access to are returned. This query resource is meant for administrative actions only, therefore has a lower rate limit tier than user permissions related resources.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRequestsRequest
*/
func (a *AccessRecordsApi) GetRequests(ctx context.Context) ApiGetRequestsRequest {
	return ApiGetRequestsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return AccessRequestCollection
func (a *AccessRecordsApi) GetRequestsExecute(r ApiGetRequestsRequest) (*AccessRequestCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRequestCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.GetRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/requests"

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
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
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

type ApiRespondToAccessRequestRequest struct {
	ctx                   context.Context
	ThisApiHandler        *AccessRecordsApi
	requestId             string
	accessRequestResponse *AccessRequestResponse
}

func (r ApiRespondToAccessRequestRequest) AccessRequestResponse(accessRequestResponse AccessRequestResponse) ApiRespondToAccessRequestRequest {
	r.accessRequestResponse = &accessRequestResponse
	return r
}

func (r ApiRespondToAccessRequestRequest) Execute() (*AccessRequest, *http.Response, error) {
	return r.ThisApiHandler.RespondToAccessRequestExecute(r)
}

/*
RespondToAccessRequest Approve or deny access request

Updates an access request, approving it or denying it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param requestId The identifier of the access request.
	@return ApiRespondToAccessRequestRequest
*/
func (a *AccessRecordsApi) RespondToAccessRequest(ctx context.Context, requestId string) ApiRespondToAccessRequestRequest {
	return ApiRespondToAccessRequestRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		requestId:      requestId,
	}
}

// Execute executes the request
//
//	@return AccessRequest
func (a *AccessRecordsApi) RespondToAccessRequestExecute(r ApiRespondToAccessRequestRequest) (*AccessRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *AccessRequest
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.RespondToAccessRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", url.PathEscape(parameterValueToString(r.requestId, "requestId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.requestId) < 1 {
		return localVarReturnValue, nil, reportError("requestId must have at least 1 elements")
	}
	if strlen(r.requestId) > 100 {
		return localVarReturnValue, nil, reportError("requestId must have less than 100 elements")
	}
	if r.accessRequestResponse == nil {
		return localVarReturnValue, nil, reportError("accessRequestResponse is required and must be specified")
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
	localVarPostBody = r.accessRequestResponse
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

type ApiUpdateRecordRequest struct {
	ctx               context.Context
	ThisApiHandler    *AccessRecordsApi
	recordId          string
	accessRecord      *AccessRecord
	ifUnmodifiedSince *time.Time
}

func (r ApiUpdateRecordRequest) AccessRecord(accessRecord AccessRecord) ApiUpdateRecordRequest {
	r.accessRecord = &accessRecord
	return r
}

// The expected last time the record was modified.
func (r ApiUpdateRecordRequest) IfUnmodifiedSince(ifUnmodifiedSince time.Time) ApiUpdateRecordRequest {
	r.ifUnmodifiedSince = &ifUnmodifiedSince
	return r
}

func (r ApiUpdateRecordRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.UpdateRecordExecute(r)
}

/*
UpdateRecord Update access record

Updates an access record adding or removing user permissions to resources. (Records have a maximum size of ~100KB)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param recordId The identifier of the access record.
	@return ApiUpdateRecordRequest
*/
func (a *AccessRecordsApi) UpdateRecord(ctx context.Context, recordId string) ApiUpdateRecordRequest {
	return ApiUpdateRecordRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		recordId:       recordId,
	}
}

// Execute executes the request
func (a *AccessRecordsApi) UpdateRecordExecute(r ApiUpdateRecordRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "AccessRecordsApi.UpdateRecord")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/records/{recordId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recordId"+"}", url.PathEscape(parameterValueToString(r.recordId, "recordId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.recordId) < 1 {
		return nil, reportError("recordId must have at least 1 elements")
	}
	if strlen(r.recordId) > 100 {
		return nil, reportError("recordId must have less than 100 elements")
	}
	if r.accessRecord == nil {
		return nil, reportError("accessRecord is required and must be specified")
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
	if r.ifUnmodifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Unmodified-Since", r.ifUnmodifiedSince, "")
	}
	// body params
	localVarPostBody = r.accessRecord
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
