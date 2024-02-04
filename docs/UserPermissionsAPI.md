#UserPermissions

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeUser**](UserPermissions.md#AuthorizeUser) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions/{permission} | Verify user authorization
[**GetUserPermissionsForResource**](UserPermissions.md#GetUserPermissionsForResource) | **Get** /v1/users/{userId}/resources/{resourceUri}/permissions | Get user permissions for resource
[**GetUserResources**](UserPermissions.md#GetUserResources) | **Get** /v1/users/{userId}/resources | List user resources
[**GetUserRolesForResource**](UserPermissions.md#GetUserRolesForResource) | **Get** /v1/users/{userId}/resources/{resourceUri}/roles | Get user roles for resource



## AuthorizeUser

> AuthorizeUser(ctx, userId, resourceUri, permission).Execute()

Verify user authorization



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
	userId := TODO // UserId | The user identifier to check and verify the permissions of.
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to validate, must be URL encoded, uri segments are allowed, the resource must be a full path.
	permission := TODO // Action | Permission to check, '*' and scoped permissions can also be checked here.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserPermissions.AuthorizeUser(context.Background(), userId, resourceUri, permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPermissions.AuthorizeUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user identifier to check and verify the permissions of. | 
**resourceUri** | **string** | The uri path of a resource to validate, must be URL encoded, uri segments are allowed, the resource must be a full path. | 
**permission** | [**Action**](.md) | Permission to check, &#39;*&#39; and scoped permissions can also be checked here. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeUserRequest struct via the builder pattern


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


## GetUserPermissionsForResource

> PermissionCollection GetUserPermissionsForResource(ctx, userId, resourceUri).Execute()

Get user permissions for resource



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
	userId := TODO // UserId | The user identifier for the user to check permissions.
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to validate, must be URL encoded, uri segments are allowed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPermissions.GetUserPermissionsForResource(context.Background(), userId, resourceUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPermissions.GetUserPermissionsForResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPermissionsForResource`: PermissionCollection
	fmt.Fprintf(os.Stdout, "Response from `UserPermissions.GetUserPermissionsForResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user identifier for the user to check permissions. | 
**resourceUri** | **string** | The uri path of a resource to validate, must be URL encoded, uri segments are allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionsForResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PermissionCollection**](PermissionCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserResources

> UserResourcesCollection GetUserResources(ctx, userId).ResourceUri(resourceUri).CollectionConfiguration(collectionConfiguration).Permissions(permissions).Limit(limit).Cursor(cursor).Execute()

List user resources



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
	userId := TODO // UserId | The user identifier for which to list all accessible resources.
	resourceUri := "/organizations" // string | The top level uri path of a resource to query for. Will only match explicit or nested sub-resources. Will not partial match resource names. (optional)
	collectionConfiguration := "collectionConfiguration_example" // string | `TOP_LEVEL_ONLY` - returns only directly nested resources under the resourceUri. A query to `resourceUri=Collection` will return `Collection/resource_1`.<br>`INCLUDE_NESTED` - will return all sub-resources as well as deeply nested resources that the user has the specified permission to. A query to `resourceUri=Collection` will return `Collection/namespaces/ns/resources/resource_1`.<br><br>To return matching resources for nested resources, set this parameter to `INCLUDE_NESTED`. (optional) (default to "TOP_LEVEL_ONLY")
	permissions := TODO // Action | Permission to check, '*' and scoped permissions can also be checked here. By default if the user has any permission explicitly to a resource, it will be included in the list. (optional)
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPermissions.GetUserResources(context.Background(), userId).ResourceUri(resourceUri).CollectionConfiguration(collectionConfiguration).Permissions(permissions).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPermissions.GetUserResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserResources`: UserResourcesCollection
	fmt.Fprintf(os.Stdout, "Response from `UserPermissions.GetUserResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user identifier for which to list all accessible resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceUri** | **string** | The top level uri path of a resource to query for. Will only match explicit or nested sub-resources. Will not partial match resource names. | 
 **collectionConfiguration** | **string** | &#x60;TOP_LEVEL_ONLY&#x60; - returns only directly nested resources under the resourceUri. A query to &#x60;resourceUri&#x3D;Collection&#x60; will return &#x60;Collection/resource_1&#x60;.&lt;br&gt;&#x60;INCLUDE_NESTED&#x60; - will return all sub-resources as well as deeply nested resources that the user has the specified permission to. A query to &#x60;resourceUri&#x3D;Collection&#x60; will return &#x60;Collection/namespaces/ns/resources/resource_1&#x60;.&lt;br&gt;&lt;br&gt;To return matching resources for nested resources, set this parameter to &#x60;INCLUDE_NESTED&#x60;. | [default to &quot;TOP_LEVEL_ONLY&quot;]
 **permissions** | [**Action**](Action.md) | Permission to check, &#39;*&#39; and scoped permissions can also be checked here. By default if the user has any permission explicitly to a resource, it will be included in the list. | 
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 

### Return type

[**UserResourcesCollection**](UserResourcesCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRolesForResource

> UserRoleCollection GetUserRolesForResource(ctx, userId, resourceUri).Execute()

Get user roles for resource



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
	userId := TODO // UserId | The user to get roles for.
	resourceUri := "/organizations/org_a/documents/doc_1" // string | The uri path of a resource to get roles for, must be URL encoded. Checks for explicit resource roles, roles attached to parent resources are not returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserPermissions.GetUserRolesForResource(context.Background(), userId, resourceUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPermissions.GetUserRolesForResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRolesForResource`: UserRoleCollection
	fmt.Fprintf(os.Stdout, "Response from `UserPermissions.GetUserRolesForResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**UserId**](.md) | The user to get roles for. | 
**resourceUri** | **string** | The uri path of a resource to get roles for, must be URL encoded. Checks for explicit resource roles, roles attached to parent resources are not returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRolesForResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserRoleCollection**](UserRoleCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

