#Connections

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](Connections.md#CreateConnection) | **Post** /v1/connections | Create SSO connection
[**DeleteConnection**](Connections.md#DeleteConnection) | **Delete** /v1/connections/{connectionId} | Delete SSO connection
[**GetConnection**](Connections.md#GetConnection) | **Get** /v1/connections/{connectionId} | Retrieve SSO connection
[**GetConnectionCredentials**](Connections.md#GetConnectionCredentials) | **Get** /v1/connections/{connectionId}/users/{userId}/credentials | Retrieve user connection credentials
[**GetConnections**](Connections.md#GetConnections) | **Get** /v1/connections | List SSO connections
[**UpdateConnection**](Connections.md#UpdateConnection) | **Put** /v1/connections/{connectionId} | Update SSO connection



## CreateConnection

> Connection CreateConnection(ctx).Connection(connection).Execute()

Create SSO connection



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
	connection := *openapiclient.NewConnection() // Connection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Connections.CreateConnection(context.Background()).Connection(connection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.CreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnection`: Connection
	fmt.Fprintf(os.Stdout, "Response from `Connections.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connection** | [**Connection**](Connection.md) |  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, connectionId).Execute()

Delete SSO connection



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
	connectionId := "connectionId_example" // string | The connection identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Connections.DeleteConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.DeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


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


## GetConnection

> Connection GetConnection(ctx, connectionId).Execute()

Retrieve SSO connection



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
	connectionId := "connectionId_example" // string | The connection identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Connections.GetConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.GetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnection`: Connection
	fmt.Fprintf(os.Stdout, "Response from `Connections.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Connection**](Connection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionCredentials

> UserConnectionCredentials GetConnectionCredentials(ctx, connectionId, userId).Execute()

Retrieve user connection credentials



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
	connectionId := "connectionId_example" // string | The connection identifier.
	userId := TODO // UserId | The connection user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Connections.GetConnectionCredentials(context.Background(), connectionId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.GetConnectionCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionCredentials`: UserConnectionCredentials
	fmt.Fprintf(os.Stdout, "Response from `Connections.GetConnectionCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection identifier. | 
**userId** | [**UserId**](.md) | The connection user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserConnectionCredentials**](UserConnectionCredentials.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> ConnectionCollection GetConnections(ctx).Execute()

List SSO connections



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Connections.GetConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.GetConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnections`: ConnectionCollection
	fmt.Fprintf(os.Stdout, "Response from `Connections.GetConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


### Return type

[**ConnectionCollection**](ConnectionCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> Connection UpdateConnection(ctx, connectionId).Connection(connection).Execute()

Update SSO connection



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
	connectionId := "connectionId_example" // string | The connection identifier.
	connection := *openapiclient.NewConnection() // Connection | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Connections.UpdateConnection(context.Background(), connectionId).Connection(connection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Connections.UpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnection`: Connection
	fmt.Fprintf(os.Stdout, "Response from `Connections.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**Connection**](Connection.md) |  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

