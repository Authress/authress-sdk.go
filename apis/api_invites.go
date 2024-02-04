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

// InvitesApi Invites service
type InvitesApi struct {
	Client *HttpClient
}

type ApiCreateInviteRequest struct {
	ctx            context.Context
	ThisApiHandler *InvitesApi
	invite         *Invite
}

func (r ApiCreateInviteRequest) Invite(invite Invite) ApiCreateInviteRequest {
	r.invite = &invite
	return r
}

func (r ApiCreateInviteRequest) Execute() (*Invite, *http.Response, error) {
	return r.ThisApiHandler.CreateInviteExecute(r)
}

/*
CreateInvite Create user invite

Invites are used to easily assign permissions to users that have not been created in your identity provider yet.

 1. This generates an invite record.

 2. Send the invite ID to the user.

 3. Log the user in.

 4. As the user PATCH the /invite/{inviteId} endpoint

 5. This accepts the invite.
    When the user accepts the invite they will automatically receive the permissions assigned in the Invite. Invites automatically expire after 7 days.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return ApiCreateInviteRequest
*/
func (a *InvitesApi) CreateInvite(ctx context.Context) ApiCreateInviteRequest {
	return ApiCreateInviteRequest{
		ThisApiHandler: a,
		ctx:            ctx,
	}
}

// Execute executes the request
//
//	@return Invite
func (a *InvitesApi) CreateInviteExecute(r ApiCreateInviteRequest) (*Invite, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Invite
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "InvitesApi.CreateInvite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/invites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invite == nil {
		return localVarReturnValue, nil, reportError("invite is required and must be specified")
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
	localVarPostBody = r.invite
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

type ApiDeleteInviteRequest struct {
	ctx            context.Context
	ThisApiHandler *InvitesApi
	inviteId       string
}

func (r ApiDeleteInviteRequest) Execute() (*http.Response, error) {
	return r.ThisApiHandler.DeleteInviteExecute(r)
}

/*
DeleteInvite Delete invite

Deletes an invite.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The identifier of the invite.
	@return ApiDeleteInviteRequest
*/
func (a *InvitesApi) DeleteInvite(ctx context.Context, inviteId string) ApiDeleteInviteRequest {
	return ApiDeleteInviteRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		inviteId:       inviteId,
	}
}

// Execute executes the request
func (a *InvitesApi) DeleteInviteExecute(r ApiDeleteInviteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []FormFile
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "InvitesApi.DeleteInvite")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/invites/{inviteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inviteId"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.inviteId) < 1 {
		return nil, reportError("inviteId must have at least 1 elements")
	}
	if strlen(r.inviteId) > 256 {
		return nil, reportError("inviteId must have less than 256 elements")
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

type ApiGetInviteRequest struct {
	ctx            context.Context
	ThisApiHandler *InvitesApi
	inviteId       string
}

func (r ApiGetInviteRequest) Execute() (*Invite, *http.Response, error) {
	return r.ThisApiHandler.GetInviteExecute(r)
}

/*
GetInvite Retrieve invite

Gets the invite along with the relevant invite data. Invites enable users inviting other users to their tenant, organization, or account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The identifier of the invite.
	@return ApiGetInviteRequest
*/
func (a *InvitesApi) GetInvite(ctx context.Context, inviteId string) ApiGetInviteRequest {
	return ApiGetInviteRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		inviteId:       inviteId,
	}
}

// Execute executes the request
//
//	@return Invite
func (a *InvitesApi) GetInviteExecute(r ApiGetInviteRequest) (*Invite, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Invite
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "InvitesApi.GetInvite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/invites/{inviteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inviteId"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.inviteId) < 1 {
		return localVarReturnValue, nil, reportError("inviteId must have at least 1 elements")
	}
	if strlen(r.inviteId) > 256 {
		return localVarReturnValue, nil, reportError("inviteId must have less than 256 elements")
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

type ApiRespondToInviteRequest struct {
	ctx            context.Context
	ThisApiHandler *InvitesApi
	inviteId       string
}

func (r ApiRespondToInviteRequest) Execute() (*Account, *http.Response, error) {
	return r.ThisApiHandler.RespondToInviteExecute(r)
}

/*
RespondToInvite Accept invite

Accepts an invite by claiming this invite by this user. The user access token used for this request will gain the permissions associated with the invite.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inviteId The identifier of the invite.
	@return ApiRespondToInviteRequest
*/
func (a *InvitesApi) RespondToInvite(ctx context.Context, inviteId string) ApiRespondToInviteRequest {
	return ApiRespondToInviteRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		inviteId:       inviteId,
	}
}

// Execute executes the request
//
//	@return Account
func (a *InvitesApi) RespondToInviteExecute(r ApiRespondToInviteRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *Account
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "InvitesApi.RespondToInvite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/invites/{inviteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inviteId"+"}", url.PathEscape(parameterValueToString(r.inviteId, "inviteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.inviteId) < 1 {
		return localVarReturnValue, nil, reportError("inviteId must have at least 1 elements")
	}
	if strlen(r.inviteId) > 256 {
		return localVarReturnValue, nil, reportError("inviteId must have less than 256 elements")
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
