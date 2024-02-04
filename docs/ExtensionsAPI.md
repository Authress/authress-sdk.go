#Extensions

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtension**](Extensions.md#CreateExtension) | **Post** /v1/extensions | Create extension
[**DeleteExtension**](Extensions.md#DeleteExtension) | **Delete** /v1/extensions/{extensionId} | Delete extension
[**GetExtension**](Extensions.md#GetExtension) | **Get** /v1/extensions/{extensionId} | Retrieve extension
[**GetExtensions**](Extensions.md#GetExtensions) | **Get** /v1/extensions | List extensions
[**Login**](Extensions.md#Login) | **Get** / | OAuth Authorize
[**RequestToken**](Extensions.md#RequestToken) | **Post** /api/authentication/oauth/tokens | OAuth Token
[**UpdateExtension**](Extensions.md#UpdateExtension) | **Put** /v1/extensions/{extensionId} | Update extension



## CreateExtension

> Extension CreateExtension(ctx).Extension(extension).Execute()

Create extension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "//"
)

func main() {
	extension := *openapiclient.NewExtension("ExtensionId_example", time.Now(), *openapiclient.NewExtensionClient("ClientId_example")) // Extension | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.CreateExtension(context.Background()).Extension(extension).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.CreateExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `Extensions.CreateExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extension** | [**Extension**](Extension.md) |  | 

### Return type

[**Extension**](Extension.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtension

> DeleteExtension(ctx, extensionId).Execute()

Delete extension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	extensionId := "extensionId_example" // string | The extension identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Extensions.DeleteExtension(context.Background(), extensionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.DeleteExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | The extension identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtension

> Extension GetExtension(ctx, extensionId).Execute()

Retrieve extension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	extensionId := "extensionId_example" // string | The extension identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.GetExtension(context.Background(), extensionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.GetExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `Extensions.GetExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | The extension identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Extension**](Extension.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensions

> ExtensionCollection GetExtensions(ctx).Limit(limit).Cursor(cursor).Execute()

List extensions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.GetExtensions(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.GetExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtensions`: ExtensionCollection
	fmt.Fprintf(os.Stdout, "Response from `Extensions.GetExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 

### Return type

[**ExtensionCollection**](ExtensionCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> OAuthAuthorizeResponse Login(ctx).ClientId(clientId).CodeChallenge(codeChallenge).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).Execute()

OAuth Authorize



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	clientId := "ext_00AA" // string | The client identifier to constrain the token to.
	codeChallenge := "6fdkQaPm51l13DSukcAH3Mdx7_ntecHYd1vi3n0hMZY" // string | The PKCE Code challenge generated by the extension UI to secure the code exchange from [RFC 7636](https://datatracker.ietf.org/doc/html/rfc7636).
	redirectUri := "https://extension.application.com/login-redirect" // string | The location to redirect the user back to after login. This redirect_uri must be a URL that matches one of the preconfigured urls in the Authress Application.
	codeChallengeMethod := "codeChallengeMethod_example" // string | The method used to generate the code_challenge from the code_verifier. `code_challenge_method(code_verifier) = code_challenge` (optional) (default to "S256")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.Login(context.Background()).ClientId(clientId).CodeChallenge(codeChallenge).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: OAuthAuthorizeResponse
	fmt.Fprintf(os.Stdout, "Response from `Extensions.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The client identifier to constrain the token to. | 
 **codeChallenge** | **string** | The PKCE Code challenge generated by the extension UI to secure the code exchange from [RFC 7636](https://datatracker.ietf.org/doc/html/rfc7636). | 
 **redirectUri** | **string** | The location to redirect the user back to after login. This redirect_uri must be a URL that matches one of the preconfigured urls in the Authress Application. | 
 **codeChallengeMethod** | **string** | The method used to generate the code_challenge from the code_verifier. &#x60;code_challenge_method(code_verifier) &#x3D; code_challenge&#x60; | [default to &quot;S256&quot;]

### Return type

[**OAuthAuthorizeResponse**](OAuthAuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestToken

> OAuthTokenResponse RequestToken(ctx).OAuthTokenRequest(oAuthTokenRequest).Execute()

OAuth Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	oAuthTokenRequest := *openapiclient.NewOAuthTokenRequest("ClientId_example") // OAuthTokenRequest | The contents of an OAuth token request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.RequestToken(context.Background()).OAuthTokenRequest(oAuthTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.RequestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestToken`: OAuthTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `Extensions.RequestToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuthTokenRequest** | [**OAuthTokenRequest**](OAuthTokenRequest.md) | The contents of an OAuth token request. | 

### Return type

[**OAuthTokenResponse**](OAuthTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtension

> Extension UpdateExtension(ctx, extensionId).Extension(extension).Execute()

Update extension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "//"
)

func main() {
	extensionId := "extensionId_example" // string | The extension identifier.
	extension := *openapiclient.NewExtension("ExtensionId_example", time.Now(), *openapiclient.NewExtensionClient("ClientId_example")) // Extension | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Extensions.UpdateExtension(context.Background(), extensionId).Extension(extension).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Extensions.UpdateExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExtension`: Extension
	fmt.Fprintf(os.Stdout, "Response from `Extensions.UpdateExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string** | The extension identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extension** | [**Extension**](Extension.md) |  | 

### Return type

[**Extension**](Extension.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

