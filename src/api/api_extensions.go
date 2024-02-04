package authress

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ExtensionsApi Extensions service
type ExtensionsApi service

type ApiCreateExtensionRequest struct {
	ctx            context.Context
	ThisApiHandler *ExtensionsApi
	extension      *Extension
}

func (r ApiCreateExtensionRequest) Extension(extension Extension) ApiCreateExtensionRequest {
	r.extension = &extension
	return r
}

func (r ApiCreateExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ThisApiHandler.CreateExtensionExecute(r)
}

/*
CreateExtension Create extension

Specify the extension details for a new developer extension. Creating the extension enables developers to build applications that can log in to your platform and interact with your users' data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateExtensionRequest
*/
func (a *ExtensionsApi) CreateExtension(ctx context.Context) ApiCreateExtensionRequest {
	return ApiCreateExtensionRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Extension
func (a *ExtensionsApi) CreateExtensionExecute(r ApiCreateExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.CreateExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/extensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.extension == nil {
		return localVarReturnValue, nil, reportError("extension is required and must be specified")
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
	localVarPostBody = r.extension
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

type ApiDeleteExtensionRequest struct {
	ctx            context.Context
	ThisApiHandler *ExtensionsApi
	extensionId    string
}

func (r ApiDeleteExtensionRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteExtensionExecute(r)
}

/*
DeleteExtension Delete extension

Deletes the specified extension. When deleted an extension can no longer be accessed. Additionally users cannot use that extension to log in, nor can the service client associated with the extension be used to access data secured by Authress. The related Access Records will automatically be deleted.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extensionId The extension identifier.
	@return ApiDeleteExtensionRequest
*/
func (a *ExtensionsApi) DeleteExtension(ctx context.Context, extensionId string) ApiDeleteExtensionRequest {
	return ApiDeleteExtensionRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		extensionId:    extensionId,
	}
}

// Execute executes the request
func (a *ExtensionsApi) DeleteExtensionExecute(r ApiDeleteExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.DeleteExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/extensions/{extensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.extensionId) < 1 {
		return nil, reportError("extensionId must have at least 1 elements")
	}
	if strlen(r.extensionId) > 64 {
		return nil, reportError("extensionId must have less than 64 elements")
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

type ApiGetExtensionRequest struct {
	ctx            context.Context
	ThisApiHandler *ExtensionsApi
	extensionId    string
}

func (r ApiGetExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ThisApiHandler.GetExtensionExecute(r)
}

/*
GetExtension Retrieve extension

Gets the platform extension details for the existing extension.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extensionId The extension identifier.
	@return ApiGetExtensionRequest
*/
func (a *ExtensionsApi) GetExtension(ctx context.Context, extensionId string) ApiGetExtensionRequest {
	return ApiGetExtensionRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		extensionId:    extensionId,
	}
}

// Execute executes the request
//
//	@return Extension
func (a *ExtensionsApi) GetExtensionExecute(r ApiGetExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.GetExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/extensions/{extensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.extensionId) < 1 {
		return localVarReturnValue, nil, reportError("extensionId must have at least 1 elements")
	}
	if strlen(r.extensionId) > 64 {
		return localVarReturnValue, nil, reportError("extensionId must have less than 64 elements")
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

type ApiGetExtensionsRequest struct {
	ctx            context.Context
	ThisApiHandler *ExtensionsApi
	limit          *int32
	cursor         *string
}

// Max number of results to return
func (r ApiGetExtensionsRequest) Limit(limit int32) ApiGetExtensionsRequest {
	r.limit = &limit
	return r
}

// Continuation cursor for paging
func (r ApiGetExtensionsRequest) Cursor(cursor string) ApiGetExtensionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetExtensionsRequest) Execute() (*ExtensionCollection, *http.Response, error) {
	return r.ThisApiHandler.GetExtensionsExecute(r)
}

/*
GetExtensions List extensions

Lists the platform extensions. Extensions are the applications that developers of your platform have created for your users to interact with. Returns a paginated extension list for the account. Only extensions the user has access to are returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetExtensionsRequest
*/
func (a *ExtensionsApi) GetExtensions(ctx context.Context) ApiGetExtensionsRequest {
	return ApiGetExtensionsRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return ExtensionCollection
func (a *ExtensionsApi) GetExtensionsExecute(r ApiGetExtensionsRequest) (*ExtensionCollection, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExtensionCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.GetExtensions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/extensions"

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

type ApiLoginRequest struct {
	ctx                 context.Context
	ThisApiHandler      *ExtensionsApi
	clientId            *string
	codeChallenge       *string
	redirectUri         *string
	codeChallengeMethod *string
}

// The client identifier to constrain the token to.
func (r ApiLoginRequest) ClientId(clientId string) ApiLoginRequest {
	r.clientId = &clientId
	return r
}

// The PKCE Code challenge generated by the extension UI to secure the code exchange from [RFC 7636](https://datatracker.ietf.org/doc/html/rfc7636).
func (r ApiLoginRequest) CodeChallenge(codeChallenge string) ApiLoginRequest {
	r.codeChallenge = &codeChallenge
	return r
}

// The location to redirect the user back to after login. This redirect_uri must be a URL that matches one of the preconfigured urls in the Authress Application.
func (r ApiLoginRequest) RedirectUri(redirectUri string) ApiLoginRequest {
	r.redirectUri = &redirectUri
	return r
}

// The method used to generate the code_challenge from the code_verifier. &#x60;code_challenge_method(code_verifier) &#x3D; code_challenge&#x60;
func (r ApiLoginRequest) CodeChallengeMethod(codeChallengeMethod string) ApiLoginRequest {
	r.codeChallengeMethod = &codeChallengeMethod
	return r
}

func (r ApiLoginRequest) Execute() (*OAuthAuthorizeResponse, *http.Response, error) {
	return r.ThisApiHandler.LoginExecute(r)
}

/*
Login OAuth Authorize

*Note*: This endpoint is only to be used for [Authress Platform Extensions](https://authress.io/knowledge-base/docs/extensions/). If you are not building an app marketplace, then tokens can be directly requested for Service Clients, using the relevant [SDK](https://authress.io/app/#/api).<br><br>Start the OAuth login by redirecting the user to the OAuth Authorize endpoint. This generates a JWT for the user using the configured application, client ID, and connection.<br><br>The OAuth 2.1 authorization request that follows [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749). Enables users to request a JWT signed by Authress and Returns an OAuth JWT containing the relevant user claims. Tokens generated must be verified before usage by validating the `sub`, `iss`, and `aud` properties in the JWT. Please note, that the properties in the request and response use snake_case to explicitly follow the standard.<br><br>The ExtensionClient in the [@authress/login](https://github.com/Authress/authress-login.js#platform-extension-login) npm package provides all the necessary logic to make this easy.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLoginRequest
*/
func (a *ExtensionsApi) Login(ctx context.Context) ApiLoginRequest {
	return ApiLoginRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return OAuthAuthorizeResponse
func (a *ExtensionsApi) LoginExecute(r ApiLoginRequest) (*OAuthAuthorizeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OAuthAuthorizeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.Login")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.codeChallenge == nil {
		return localVarReturnValue, nil, reportError("codeChallenge is required and must be specified")
	}
	if r.redirectUri == nil {
		return localVarReturnValue, nil, reportError("redirectUri is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "client_id", r.clientId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "code_challenge", r.codeChallenge, "")
	if r.codeChallengeMethod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "code_challenge_method", r.codeChallengeMethod, "")
	} else {
		var defaultValue string = "S256"
		r.codeChallengeMethod = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "redirect_uri", r.redirectUri, "")
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

type ApiRequestTokenRequest struct {
	ctx               context.Context
	ThisApiHandler    *ExtensionsApi
	oAuthTokenRequest *OAuthTokenRequest
}

// The contents of an OAuth token request.
func (r ApiRequestTokenRequest) OAuthTokenRequest(oAuthTokenRequest OAuthTokenRequest) ApiRequestTokenRequest {
	r.oAuthTokenRequest = &oAuthTokenRequest
	return r
}

func (r ApiRequestTokenRequest) Execute() (*OAuthTokenResponse, *http.Response, error) {
	return r.ThisApiHandler.RequestTokenExecute(r)
}

/*
RequestToken OAuth Token

*Note*: This endpoint is only to be used for [Authress Platform Extensions](https://authress.io/knowledge-base/docs/extensions/). If you are not building an app marketplace, then tokens can be directly requested for Service Clients, using the relevant [SDK](https://authress.io/app/#/api).<br><br>Request an OAuth JWT. Can either be called with service client credentials or as the second part of the user authorize login flow.<br>When using the `password` grant_type, service client authentication must be used via the Authress SDKs, and requires the `Authress:AuthenticateUser` role.<br><br>The OAuth 2.1 token request that follows [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749). Enables users to request a JWT signed by Authress, and returns an OAuth JWT containing the relevant user claims. Tokens generated must be verified before usage by validating the `sub`, `iss`, and `aud` properties in the JWT. Please note, that the properties in the request and response use snake_case to explicitly follow the standard.<br><br>The ExtensionClient in the [@authress/login](https://github.com/Authress/authress-login.js#platform-extension-login) npm package provides all the necessary logic to make this easy.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRequestTokenRequest
*/
func (a *ExtensionsApi) RequestToken(ctx context.Context) ApiRequestTokenRequest {
	return ApiRequestTokenRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return OAuthTokenResponse
func (a *ExtensionsApi) RequestTokenExecute(r ApiRequestTokenRequest) (*OAuthTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OAuthTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.RequestToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/authentication/oauth/tokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oAuthTokenRequest == nil {
		return localVarReturnValue, nil, reportError("oAuthTokenRequest is required and must be specified")
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
	localVarPostBody = r.oAuthTokenRequest
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

type ApiUpdateExtensionRequest struct {
	ctx            context.Context
	ThisApiHandler *ExtensionsApi
	extensionId    string
	extension      *Extension
}

func (r ApiUpdateExtensionRequest) Extension(extension Extension) ApiUpdateExtensionRequest {
	r.extension = &extension
	return r
}

func (r ApiUpdateExtensionRequest) Execute() (*Extension, *http.Response, error) {
	return r.ThisApiHandler.UpdateExtensionExecute(r)
}

/*
UpdateExtension Update extension

Specify the updated extension. The extension will be updated and these changes will be reflected to the access control and user authentication associated with the extension.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param extensionId The extension identifier.
	@return ApiUpdateExtensionRequest
*/
func (a *ExtensionsApi) UpdateExtension(ctx context.Context, extensionId string) ApiUpdateExtensionRequest {
	return ApiUpdateExtensionRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		extensionId:    extensionId,
	}
}

// Execute executes the request
//
//	@return Extension
func (a *ExtensionsApi) UpdateExtensionExecute(r ApiUpdateExtensionRequest) (*Extension, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Extension
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExtensionsApi.UpdateExtension")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/extensions/{extensionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.extensionId) < 1 {
		return localVarReturnValue, nil, reportError("extensionId must have at least 1 elements")
	}
	if strlen(r.extensionId) > 64 {
		return localVarReturnValue, nil, reportError("extensionId must have less than 64 elements")
	}
	if r.extension == nil {
		return localVarReturnValue, nil, reportError("extension is required and must be specified")
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
	localVarPostBody = r.extension
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
