# \GroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /v1/groups | Create group
[**DeleteGroup**](GroupsAPI.md#DeleteGroup) | **Delete** /v1/groups/{groupId} | Deletes group
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /v1/groups/{groupId} | Retrieve group
[**GetGroups**](GroupsAPI.md#GetGroups) | **Get** /v1/groups | List groups
[**UpdateGroup**](GroupsAPI.md#UpdateGroup) | **Put** /v1/groups/{groupId} | Update a group



## CreateGroup

> Group CreateGroup(ctx).Group(group).Execute()

Create group



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
	group := *openapiclient.NewGroup("Name_example", []openapiclient.User{*openapiclient.NewUser("oauth|userId")}, []openapiclient.User{*openapiclient.NewUser("oauth|userId")}) // Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**Group**](Group.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()

Deletes group



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
	groupId := TODO // GroupId | The identifier of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**GroupId**](.md) | The identifier of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> Group GetGroup(ctx, groupId).Execute()

Retrieve group



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
	groupId := TODO // GroupId | The identifier of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**GroupId**](.md) | The identifier of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GroupCollection GetGroups(ctx).Limit(limit).Cursor(cursor).Filter(filter).Execute()

List groups



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
	filter := "filter_example" // string | Filter to search groups by. This is a case insensitive search through every text field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroups(context.Background()).Limit(limit).Cursor(cursor).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: GroupCollection
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 
 **filter** | **string** | Filter to search groups by. This is a case insensitive search through every text field. | 

### Return type

[**GroupCollection**](GroupCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupId).Group(group).IfUnmodifiedSince(ifUnmodifiedSince).Execute()

Update a group



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
	groupId := TODO // GroupId | The identifier of the group.
	group := *openapiclient.NewGroup("Name_example", []openapiclient.User{*openapiclient.NewUser("oauth|userId")}, []openapiclient.User{*openapiclient.NewUser("oauth|userId")}) // Group | 
	ifUnmodifiedSince := time.Now() // time.Time | The expected last time the group was modified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), groupId).Group(group).IfUnmodifiedSince(ifUnmodifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**GroupId**](.md) | The identifier of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) |  | 
 **ifUnmodifiedSince** | **time.Time** | The expected last time the group was modified. | 

### Return type

[**Group**](Group.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

