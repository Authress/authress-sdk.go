# ResourcePermissions API

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissionedResource**](#GetPermissionedResource) | **Get** /v1/resources/{resourceUri} | Retrieve resource configuration
[**GetPermissionedResources**](#GetPermissionedResources) | **Get** /v1/resources | List all resource configurations
[**GetResourceUsers**](#GetResourceUsers) | **Get** /v1/resources/{resourceUri}/users | List users with resource access
[**UpdatePermissionedResource**](#UpdatePermissionedResource) | **Put** /v1/resources/{resourceUri} | Update resource configuration



## GetPermissionedResource

> PermissionedResource GetPermissionedResource(ctx, resourceUri).Execute()

Retrieve resource configuration



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
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to validate, must be URL encoded, uri segments are allowed.

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := apiClient.ResourcePermissions.GetPermissionedResource(context.Background(), resourceUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePermissions.GetPermissionedResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionedResource`: PermissionedResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcePermissions.GetPermissionedResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceUri** | **string** | The uri path of a resource to validate, must be URL encoded, uri segments are allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionedResource**](PermissionedResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetPermissionedResources

> PermissionedResourceCollection GetPermissionedResources(ctx).Execute()

List all resource configurations



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

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := apiClient.ResourcePermissions.GetPermissionedResources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePermissions.GetPermissionedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionedResources`: PermissionedResourceCollection
	fmt.Fprintf(os.Stdout, "Response from `ResourcePermissions.GetPermissionedResources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionedResourcesRequest struct via the builder pattern


### Return type

[**PermissionedResourceCollection**](PermissionedResourceCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetResourceUsers

> ResourceUsersCollection GetResourceUsers(ctx, resourceUri).Limit(limit).Cursor(cursor).Execute()

List users with resource access



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
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := apiClient.ResourcePermissions.GetResourceUsers(context.Background(), resourceUri).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePermissions.GetResourceUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceUsers`: ResourceUsersCollection
	fmt.Fprintf(os.Stdout, "Response from `ResourcePermissions.GetResourceUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceUri** | **string** | The uri path of a resource to validate, must be URL encoded, uri segments are allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 

### Return type

[**ResourceUsersCollection**](ResourceUsersCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## UpdatePermissionedResource

> UpdatePermissionedResource(ctx, resourceUri).PermissionedResource(permissionedResource).Execute()

Update resource configuration



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
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to validate, must be URL encoded, uri segments are allowed.
	permissionedResource := *authress.NewPermissionedResource([]authress.ResourcePermission{*authress.NewResourcePermission("Action_example", false)}) // PermissionedResource | The contents of the permission to set on the resource. Overwrites existing data.

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	r, err := apiClient.ResourcePermissions.UpdatePermissionedResource(context.Background(), resourceUri).PermissionedResource(permissionedResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcePermissions.UpdatePermissionedResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceUri** | **string** | The uri path of a resource to validate, must be URL encoded, uri segments are allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionedResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionedResource** | [**PermissionedResource**](PermissionedResource.md) | The contents of the permission to set on the resource. Overwrites existing data. | 

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

