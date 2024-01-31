# \InvitesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvite**](InvitesAPI.md#CreateInvite) | **Post** /v1/invites | Create user invite
[**DeleteInvite**](InvitesAPI.md#DeleteInvite) | **Delete** /v1/invites/{inviteId} | Delete invite
[**GetInvite**](InvitesAPI.md#GetInvite) | **Get** /v1/invites/{inviteId} | Retrieve invite
[**RespondToInvite**](InvitesAPI.md#RespondToInvite) | **Patch** /v1/invites/{inviteId} | Accept invite



## CreateInvite

> Invite CreateInvite(ctx).Invite(invite).Execute()

Create user invite



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
	invite := *openapiclient.NewInvite("InviteId_example", []openapiclient.Statement{*openapiclient.NewStatement([]string{"Roles_example"}, []openapiclient.Resource{*openapiclient.NewResource("/organizations/org_a/documents/doc_1")})}) // Invite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.CreateInvite(context.Background()).Invite(invite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.CreateInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvite`: Invite
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.CreateInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invite** | [**Invite**](Invite.md) |  | 

### Return type

[**Invite**](Invite.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvite

> DeleteInvite(ctx, inviteId).Execute()

Delete invite



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
	inviteId := "inviteId_example" // string | The identifier of the invite.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvitesAPI.DeleteInvite(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.DeleteInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** | The identifier of the invite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInviteRequest struct via the builder pattern


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


## GetInvite

> Invite GetInvite(ctx, inviteId).Execute()

Retrieve invite



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
	inviteId := "inviteId_example" // string | The identifier of the invite.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.GetInvite(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.GetInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvite`: Invite
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.GetInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** | The identifier of the invite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invite**](Invite.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondToInvite

> Account RespondToInvite(ctx, inviteId).Execute()

Accept invite



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
	inviteId := "inviteId_example" // string | The identifier of the invite.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitesAPI.RespondToInvite(context.Background(), inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitesAPI.RespondToInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondToInvite`: Account
	fmt.Fprintf(os.Stdout, "Response from `InvitesAPI.RespondToInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** | The identifier of the invite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondToInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

