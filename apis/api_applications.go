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

// ApplicationsApi Applications service
type ApplicationsApi struct {
	Client *HttpClient
}

type ApiDelegateUserLoginRequest struct {
	ctx            context.Context
	ThisApiHandler *ApplicationsApi
	applicationId  string
	userId         UserId
}

func (r ApiDelegateUserLoginRequest) Execute() (*ApplicationDelegation, *http.Response, error) {
	return r.ThisApiHandler.DelegateUserLoginExecute(r)
}

/*
DelegateUserLogin Log user into third-party application

Redirect the user to an external application to login them in. Authress uses OpenID and SAML configuration to securely pass the user's login session in your platform to an external platform. The user will then be logged into that platform.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param applicationId The application to have the user log into.
	@param userId The user.
	@return ApiDelegateUserLoginRequest
*/
func (a *ApplicationsApi) DelegateUserLogin(ctx context.Context, applicationId string, userId UserId) ApiDelegateUserLoginRequest {
	return ApiDelegateUserLoginRequest{
		ThisApiHandler: a,
		ctx:            ctx,
		applicationId:  applicationId,
		userId:         userId,
	}
}

// Execute executes the request
//
//	@return ApplicationDelegation
func (a *ApplicationsApi) DelegateUserLoginExecute(r ApiDelegateUserLoginRequest) (*ApplicationDelegation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *ApplicationDelegation
	)

	localBasePath, err := a.Client.ClientConfiguration.ServerURLWithContext(r.ctx, "ApplicationsApi.DelegateUserLogin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/applications/{applicationId}/users/{userId}/delegation"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.applicationId) < 1 {
		return localVarReturnValue, nil, reportError("applicationId must have at least 1 elements")
	}
	if strlen(r.applicationId) > 64 {
		return localVarReturnValue, nil, reportError("applicationId must have less than 64 elements")
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
