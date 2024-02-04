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

// TenantsApi Tenants service
type TenantsApi struct {
	Client *HttpClient
}

type ApiCreateTenantRequest struct {
	ctx            context.Context
	ThisApiHandler *TenantsApi
	tenant         *Tenant
}

func (r ApiCreateTenantRequest) Tenant(tenant Tenant) ApiCreateTenantRequest {
	r.tenant = &tenant
	return r
}

func (r ApiCreateTenantRequest) Execute() (*Tenant, *http.Response, error) {
	return r.ThisApiHandler.CreateTenantExecute(r)
}

/*
CreateTenant Create tenant

Specify tenant identity details for tenant tracking, separation, and user management. Tenant identifiers are persisted to access tokens created by Authress.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTenantRequest
*/
func (a *TenantsApi) CreateTenant(ctx context.Context) ApiCreateTenantRequest {
	return ApiCreateTenantRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Tenant
func (a *TenantsApi) CreateTenantExecute(r ApiCreateTenantRequest) (*Tenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Tenant
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "TenantsApi.CreateTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenant == nil {
		return localVarReturnValue, nil, reportError("tenant is required and must be specified")
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
	localVarPostBody = r.tenant
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

type ApiDeleteTenantRequest struct {
	ctx            context.Context
	ThisApiHandler *TenantsApi
	tenantId       TenantId
}

func (r ApiDeleteTenantRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteTenantExecute(r)
}

/*
DeleteTenant Delete tenant

Delete a tenant. If a connection was created for the tenant then it is deleted as well.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId The tenantId.
	@return ApiDeleteTenantRequest
*/
func (a *TenantsApi) DeleteTenant(ctx context.Context, tenantId TenantId) ApiDeleteTenantRequest {
	return ApiDeleteTenantRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		tenantId:       tenantId,
	}
}

// Execute executes the request
func (a *TenantsApi) DeleteTenantExecute(r ApiDeleteTenantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "TenantsApi.DeleteTenant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenantId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

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

type ApiGetTenantRequest struct {
	ctx            context.Context
	ThisApiHandler *TenantsApi
	tenantId       TenantId
}

func (r ApiGetTenantRequest) Execute() (*Tenant, *http.Response, error) {
	return r.ThisApiHandler.GetTenantExecute(r)
}

/*
GetTenant Retrieve tenant

Get the tenant details for an Authress tenant.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId The tenantId.
	@return ApiGetTenantRequest
*/
func (a *TenantsApi) GetTenant(ctx context.Context, tenantId TenantId) ApiGetTenantRequest {
	return ApiGetTenantRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		tenantId:       tenantId,
	}
}

// Execute executes the request
//
//	@return Tenant
func (a *TenantsApi) GetTenantExecute(r ApiGetTenantRequest) (*Tenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Tenant
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "TenantsApi.GetTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenantId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

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

type ApiGetTenantsRequest struct {
	ctx            context.Context
	ThisApiHandler *TenantsApi
}

func (r ApiGetTenantsRequest) Execute() (*TenantCollection, *http.Response, error) {
	return r.ThisApiHandler.GetTenantsExecute(r)
}

/*
GetTenants List tenants

Returns a paginated tenants list for the account. Only tenants the user has access to are returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTenantsRequest
*/
func (a *TenantsApi) GetTenants(ctx context.Context) ApiGetTenantsRequest {
	return ApiGetTenantsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return TenantCollection
func (a *TenantsApi) GetTenantsExecute(r ApiGetTenantsRequest) (*TenantCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *TenantCollection
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "TenantsApi.GetTenants")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants"

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

type ApiUpdateTenantRequest struct {
	ctx            context.Context
	ThisApiHandler *TenantsApi
	tenantId       TenantId
	tenant         *Tenant
}

func (r ApiUpdateTenantRequest) Tenant(tenant Tenant) ApiUpdateTenantRequest {
	r.tenant = &tenant
	return r
}

func (r ApiUpdateTenantRequest) Execute() (*Tenant, *http.Response, error) {
	return r.ThisApiHandler.UpdateTenantExecute(r)
}

/*
UpdateTenant Update tenant

Updates the tenants information and linked tenant if it exists.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId The tenantId.
	@return ApiUpdateTenantRequest
*/
func (a *TenantsApi) UpdateTenant(ctx context.Context, tenantId TenantId) ApiUpdateTenantRequest {
	return ApiUpdateTenantRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		tenantId:       tenantId,
	}
}

// Execute executes the request
//
//	@return Tenant
func (a *TenantsApi) UpdateTenantExecute(r ApiUpdateTenantRequest) (*Tenant, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Tenant
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "TenantsApi.UpdateTenant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/tenants/{tenantId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenant == nil {
		return localVarReturnValue, nil, reportError("tenant is required and must be specified")
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
	localVarPostBody = r.tenant
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
