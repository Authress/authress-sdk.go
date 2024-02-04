# Users API

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](#DeleteUser) | **Delete** /v1/users/{userId} | Delete a user
[**GetUser**](#GetUser) | **Get** /v1/users/{userId} | Retrieve a user
[**GetUsers**](#GetUsers) | **Get** /v1/users | List users



## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete a user



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
	userId := TODO // UserId | The user identifier.

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	r, err := apiClient.Users.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Users.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetUser

> UserIdentity GetUser(ctx, userId).Execute()

Retrieve a user



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
	userId := TODO // UserId | The user identifier.

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := apiClient.Users.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Users.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: UserIdentity
	fmt.Fprintf(os.Stdout, "Response from `Users.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserIdentity**](UserIdentity.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetUsers

> UserIdentityCollection GetUsers(ctx).Limit(limit).Cursor(cursor).Filter(filter).TenantId(tenantId).Execute()

List users



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
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 100)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)
	filter := "filter_example" // string | Filter to search users by. This is a case insensitive search through every text field. (optional)
	tenantId := TODO // TenantId | Return only users that are part of the specified tenant. Users can only be part of one tenant, using this parameter will limit returned users that have logged into this tenant. (optional)

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := apiClient.Users.GetUsers(context.Background()).Limit(limit).Cursor(cursor).Filter(filter).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Users.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: UserIdentityCollection
	fmt.Fprintf(os.Stdout, "Response from `Users.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 100]
 **cursor** | **string** | Continuation cursor for paging | 
 **filter** | **string** | Filter to search users by. This is a case insensitive search through every text field. | 
 **tenantId** | [**TenantId**](TenantId.md) | Return only users that are part of the specified tenant. Users can only be part of one tenant, using this parameter will limit returned users that have logged into this tenant. | 

### Return type

[**UserIdentityCollection**](UserIdentityCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

