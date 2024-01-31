# \ServiceClientsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClient**](ServiceClientsAPI.md#CreateClient) | **Post** /v1/clients | Create service client
[**DeleteAccessKey**](ServiceClientsAPI.md#DeleteAccessKey) | **Delete** /v1/clients/{clientId}/access-keys/{keyId} | Delete service client access key
[**DeleteClient**](ServiceClientsAPI.md#DeleteClient) | **Delete** /v1/clients/{clientId} | Delete service client
[**GetClient**](ServiceClientsAPI.md#GetClient) | **Get** /v1/clients/{clientId} | Retrieve service client
[**GetClients**](ServiceClientsAPI.md#GetClients) | **Get** /v1/clients | List service clients
[**RequestAccessKey**](ServiceClientsAPI.md#RequestAccessKey) | **Post** /v1/clients/{clientId}/access-keys | Generate service client access key
[**UpdateClient**](ServiceClientsAPI.md#UpdateClient) | **Put** /v1/clients/{clientId} | Update service client



## CreateClient

> Client CreateClient(ctx).Client(client).Execute()

Create service client



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
	client := *openapiclient.NewClient("ClientId_example", time.Now()) // Client | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceClientsAPI.CreateClient(context.Background()).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.CreateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClientsAPI.CreateClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKey

> DeleteAccessKey(ctx, clientId, keyId).Execute()

Delete service client access key



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
	clientId := "clientId_example" // string | The unique identifier of the client.
	keyId := "keyId_example" // string | The ID of the access key to remove from the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceClientsAPI.DeleteAccessKey(context.Background(), clientId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.DeleteAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique identifier of the client. | 
**keyId** | **string** | The ID of the access key to remove from the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessKeyRequest struct via the builder pattern


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


## DeleteClient

> DeleteClient(ctx, clientId).Execute()

Delete service client



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
	clientId := "clientId_example" // string | The unique identifier for the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceClientsAPI.DeleteClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique identifier for the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


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


## GetClient

> Client GetClient(ctx, clientId).Execute()

Retrieve service client



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
	clientId := "clientId_example" // string | The unique identifier for the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceClientsAPI.GetClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClientsAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique identifier for the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Client**](Client.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> ClientCollection GetClients(ctx).Limit(limit).Cursor(cursor).Execute()

List service clients



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
	cursor := "cursor_example" // string | Continuation cursor for paging. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceClientsAPI.GetClients(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: ClientCollection
	fmt.Fprintf(os.Stdout, "Response from `ServiceClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging. | 

### Return type

[**ClientCollection**](ClientCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestAccessKey

> ClientAccessKey RequestAccessKey(ctx, clientId).Execute()

Generate service client access key



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
	clientId := "clientId_example" // string | The unique identifier of the client.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceClientsAPI.RequestAccessKey(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.RequestAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestAccessKey`: ClientAccessKey
	fmt.Fprintf(os.Stdout, "Response from `ServiceClientsAPI.RequestAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique identifier of the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientAccessKey**](ClientAccessKey.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClient

> Client UpdateClient(ctx, clientId).Client(client).Execute()

Update service client



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
	clientId := "clientId_example" // string | The unique identifier for the client.
	client := *openapiclient.NewClient("ClientId_example", time.Now()) // Client | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceClientsAPI.UpdateClient(context.Background(), clientId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClientsAPI.UpdateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClientsAPI.UpdateClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The unique identifier for the client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

