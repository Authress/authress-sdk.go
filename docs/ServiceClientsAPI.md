# ServiceClients API

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClient**](#CreateClient) | **Post** /v1/clients | Create service client
[**DeleteAccessKey**](#DeleteAccessKey) | **Delete** /v1/clients/{clientId}/access-keys/{keyId} | Delete service client access key
[**DeleteClient**](#DeleteClient) | **Delete** /v1/clients/{clientId} | Delete service client
[**GetClient**](#GetClient) | **Get** /v1/clients/{clientId} | Retrieve service client
[**GetClients**](#GetClients) | **Get** /v1/clients | List service clients
[**RequestAccessKey**](#RequestAccessKey) | **Post** /v1/clients/{clientId}/access-keys | Generate service client access key
[**UpdateClient**](#UpdateClient) | **Put** /v1/clients/{clientId} | Update service client



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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	client := *models.NewClient("ClientId_example", time.Now()) // Client | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.ServiceClients.CreateClient(context.Background()).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.CreateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClients.CreateClient`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	clientId := "clientId_example" // string | The unique identifier of the client.
	keyId := "keyId_example" // string | The ID of the access key to remove from the client.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.ServiceClients.DeleteAccessKey(context.Background(), clientId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.DeleteAccessKey``: %v\n", err)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	clientId := "clientId_example" // string | The unique identifier for the client.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.ServiceClients.DeleteClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.DeleteClient``: %v\n", err)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	clientId := "clientId_example" // string | The unique identifier for the client.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.ServiceClients.GetClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClients.GetClient`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging. (optional)

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.ServiceClients.GetClients(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: ClientCollection
	fmt.Fprintf(os.Stdout, "Response from `ServiceClients.GetClients`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	clientId := "clientId_example" // string | The unique identifier of the client.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.ServiceClients.RequestAccessKey(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.RequestAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestAccessKey`: ClientAccessKey
	fmt.Fprintf(os.Stdout, "Response from `ServiceClients.RequestAccessKey`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	clientId := "clientId_example" // string | The unique identifier for the client.
	client := *models.NewClient("ClientId_example", time.Now()) // Client | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.ServiceClients.UpdateClient(context.Background(), clientId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceClients.UpdateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ServiceClients.UpdateClient`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

