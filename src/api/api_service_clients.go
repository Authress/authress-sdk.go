package authress

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ServiceClientsApi ServiceClients service
type ServiceClientsApi service

type ApiCreateClientRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	client         *Client
}

func (r ApiCreateClientRequest) Client(client Client) ApiCreateClientRequest {
	r.client = &client
	return r
}

func (r ApiCreateClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ThisApiHandler.CreateClientExecute(r)
}

/*
CreateClient Create service client

Creates a service client to interact with Authress or any other service on behalf of users. Each client has secret private keys used to authenticate with Authress. To use service clients created through other mechanisms, skip creating a client and create access records with the client identifier.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateClientRequest
*/
func (a *ServiceClientsApi) CreateClient(ctx context.Context) ApiCreateClientRequest {
	return ApiCreateClientRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ServiceClientsApi) CreateClientExecute(r ApiCreateClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.CreateClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.client == nil {
		return localVarReturnValue, nil, reportError("client is required and must be specified")
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
	localVarPostBody = r.client
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

type ApiDeleteAccessKeyRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	clientId       string
	keyId          string
}

func (r ApiDeleteAccessKeyRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteAccessKeyExecute(r)
}

/*
DeleteAccessKey Delete service client access key

Deletes an access key for a client prevent it from being used to authenticate with Authress.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The unique identifier of the client.
	@param keyId The ID of the access key to remove from the client.
	@return ApiDeleteAccessKeyRequest
*/
func (a *ServiceClientsApi) DeleteAccessKey(ctx context.Context, clientId string, keyId string) ApiDeleteAccessKeyRequest {
	return ApiDeleteAccessKeyRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		clientId:       clientId,
		keyId:          keyId,
	}
}

// Execute executes the request
func (a *ServiceClientsApi) DeleteAccessKeyExecute(r ApiDeleteAccessKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.DeleteAccessKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients/{clientId}/access-keys/{keyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyId"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return nil, reportError("clientId must have at least 1 elements")
	}
	if strlen(r.clientId) > 64 {
		return nil, reportError("clientId must have less than 64 elements")
	}
	if strlen(r.keyId) < 1 {
		return nil, reportError("keyId must have at least 1 elements")
	}
	if strlen(r.keyId) > 64 {
		return nil, reportError("keyId must have less than 64 elements")
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

type ApiDeleteClientRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	clientId       string
}

func (r ApiDeleteClientRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteClientExecute(r)
}

/*
DeleteClient Delete service client

This deletes the service client.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The unique identifier for the client.
	@return ApiDeleteClientRequest
*/
func (a *ServiceClientsApi) DeleteClient(ctx context.Context, clientId string) ApiDeleteClientRequest {
	return ApiDeleteClientRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		clientId:       clientId,
	}
}

// Execute executes the request
func (a *ServiceClientsApi) DeleteClientExecute(r ApiDeleteClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.DeleteClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return nil, reportError("clientId must have at least 1 elements")
	}
	if strlen(r.clientId) > 64 {
		return nil, reportError("clientId must have less than 64 elements")
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

type ApiGetClientRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	clientId       string
}

func (r ApiGetClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ThisApiHandler.GetClientExecute(r)
}

/*
GetClient Retrieve service client

Returns all information related to client except for the private access keys.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The unique identifier for the client.
	@return ApiGetClientRequest
*/
func (a *ServiceClientsApi) GetClient(ctx context.Context, clientId string) ApiGetClientRequest {
	return ApiGetClientRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		clientId:       clientId,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ServiceClientsApi) GetClientExecute(r ApiGetClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.GetClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
	}
	if strlen(r.clientId) > 64 {
		return localVarReturnValue, nil, reportError("clientId must have less than 64 elements")
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

type ApiGetClientsRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	limit          *int32
	cursor         *string
}

// Max number of results to return
func (r ApiGetClientsRequest) Limit(limit int32) ApiGetClientsRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging.
func (r ApiGetClientsRequest) Cursor(cursor string) ApiGetClientsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetClientsRequest) Execute() (*ClientCollection, *http.Response, error) {
	return r.ThisApiHandler.GetClientsExecute(r)
}

/*
GetClients List service clients

Returns all clients that the user has access to in the account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClientsRequest
*/
func (a *ServiceClientsApi) GetClients(ctx context.Context) ApiGetClientsRequest {
	return ApiGetClientsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return ClientCollection
func (a *ServiceClientsApi) GetClientsExecute(r ApiGetClientsRequest) (*ClientCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.GetClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients"

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

type ApiRequestAccessKeyRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	clientId       string
}

func (r ApiRequestAccessKeyRequest) Execute() (*ClientAccessKey, *http.Response, error) {
	return r.ThisApiHandler.RequestAccessKeyExecute(r)
}

/*
RequestAccessKey Generate service client access key

Create a new access key for the client or upload a compliant public key, so that a service can authenticate with Authress as that client. Using the client will allow delegation of permission checking of users. Enables machine-to-machine authentication between your services and with your end users. (There is a limit of 5 Active keys per client)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The unique identifier of the client.
	@return ApiRequestAccessKeyRequest
*/
func (a *ServiceClientsApi) RequestAccessKey(ctx context.Context, clientId string) ApiRequestAccessKeyRequest {
	return ApiRequestAccessKeyRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		clientId:       clientId,
	}
}

// Execute executes the request
//
//	@return ClientAccessKey
func (a *ServiceClientsApi) RequestAccessKeyExecute(r ApiRequestAccessKeyRequest) (*ClientAccessKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ClientAccessKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.RequestAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients/{clientId}/access-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
	}
	if strlen(r.clientId) > 64 {
		return localVarReturnValue, nil, reportError("clientId must have less than 64 elements")
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

type ApiUpdateClientRequest struct {
	ctx            context.Context
	ThisApiHandler *ServiceClientsApi
	clientId       string
	client         *Client
}

func (r ApiUpdateClientRequest) Client(client Client) ApiUpdateClientRequest {
	r.client = &client
	return r
}

func (r ApiUpdateClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ThisApiHandler.UpdateClientExecute(r)
}

/*
UpdateClient Update service client

Updates a client information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clientId The unique identifier for the client.
	@return ApiUpdateClientRequest
*/
func (a *ServiceClientsApi) UpdateClient(ctx context.Context, clientId string) ApiUpdateClientRequest {
	return ApiUpdateClientRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		clientId:       clientId,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ServiceClientsApi) UpdateClientExecute(r ApiUpdateClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceClientsApi.UpdateClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
	}
	if strlen(r.clientId) > 64 {
		return localVarReturnValue, nil, reportError("clientId must have less than 64 elements")
	}
	if r.client == nil {
		return localVarReturnValue, nil, reportError("client is required and must be specified")
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
	localVarPostBody = r.client
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
