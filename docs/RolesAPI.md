# Roles API

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](#CreateRole) | **Post** /v1/roles | Create role
[**DeleteRole**](#DeleteRole) | **Delete** /v1/roles/{roleId} | Deletes role
[**GetRole**](#GetRole) | **Get** /v1/roles/{roleId} | Retrieve role
[**GetRoles**](#GetRoles) | **Get** /v1/roles | List roles
[**UpdateRole**](#UpdateRole) | **Put** /v1/roles/{roleId} | Update role



## CreateRole

> Role CreateRole(ctx).Role(role).Execute()

Create role



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
	role := *authress.NewRole("Name_example", []authress.PermissionObject{*authress.NewPermissionObject("documents:read", false, false, false)}) // Role | 

	authressClient := authress.AuthressClient.New(authress.AuthressSettings {})
	resp, r, err := apiClient.Roles.CreateRole(context.Background()).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Roles.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `Roles.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## DeleteRole

> DeleteRole(ctx, roleId).Execute()

Deletes role



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
	roleId := "roleId_example" // string | The identifier of the role.

	authressClient := authress.AuthressClient.New(authress.AuthressSettings {})
	r, err := apiClient.Roles.DeleteRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Roles.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


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


## GetRole

> Role GetRole(ctx, roleId).Execute()

Retrieve role



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
	roleId := "roleId_example" // string | The identifier of the role.

	authressClient := authress.AuthressClient.New(authress.AuthressSettings {})
	resp, r, err := apiClient.Roles.GetRole(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Roles.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `Roles.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetRoles

> RoleCollection GetRoles(ctx).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List roles



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
	limit := int32(56) // int32 | Max number of results to return. (optional) (default to 100)
	cursor := "cursor_example" // string | Continuation cursor for paging. (optional)
	filter := "filter_example" // string | Filter to search roles by. This is a case insensitive search. (optional)

	authressClient := authress.AuthressClient.New(authress.AuthressSettings {})
	resp, r, err := apiClient.Roles.GetRoles(context.Background()).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Roles.GetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoles`: RoleCollection
	fmt.Fprintf(os.Stdout, "Response from `Roles.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. | [default to 100]
 **cursor** | **string** | Continuation cursor for paging. | 
 **filter** | **string** | Filter to search roles by. This is a case insensitive search. | 

### Return type

[**RoleCollection**](RoleCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## UpdateRole

> Role UpdateRole(ctx, roleId).Role(role).Execute()

Update role



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
	roleId := "roleId_example" // string | The identifier of the role.
	role := *authress.NewRole("Name_example", []authress.PermissionObject{*authress.NewPermissionObject("documents:read", false, false, false)}) // Role | 

	authressClient := authress.AuthressClient.New(authress.AuthressSettings {})
	resp, r, err := apiClient.Roles.UpdateRole(context.Background(), roleId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Roles.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `Roles.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The identifier of the role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

