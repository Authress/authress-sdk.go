# Applications API

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelegateUserLogin**](#DelegateUserLogin) | **Post** /v1/applications/{applicationId}/users/{userId}/delegation | Log user into third-party application



## DelegateUserLogin

> ApplicationDelegation DelegateUserLogin(ctx, applicationId, userId).Execute()

Log user into third-party application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	applicationId := "applicationId_example" // string | The application to have the user log into.
	userId := TODO // UserId | The user.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.Applications.DelegateUserLogin(context.Background(), applicationId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Applications.DelegateUserLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelegateUserLogin`: ApplicationDelegation
	fmt.Fprintf(os.Stdout, "Response from `Applications.DelegateUserLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The application to have the user log into. | 
**userId** | **string** | The user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegateUserLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationDelegation**](ApplicationDelegation.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

