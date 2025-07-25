# Tenants API

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](#CreateTenant) | **Post** /v1/tenants | Create tenant
[**DeleteTenant**](#DeleteTenant) | **Delete** /v1/tenants/{tenantId} | Delete tenant
[**GetTenant**](#GetTenant) | **Get** /v1/tenants/{tenantId} | Retrieve tenant
[**GetTenants**](#GetTenants) | **Get** /v1/tenants | List tenants
[**UpdateTenant**](#UpdateTenant) | **Put** /v1/tenants/{tenantId} | Update tenant



## CreateTenant

> Tenant CreateTenant(ctx).Tenant(tenant).Execute()

Create tenant



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
	tenant := *models.NewTenant() // Tenant | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.Tenants.CreateTenant(context.Background()).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tenants.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `Tenants.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## DeleteTenant

> DeleteTenant(ctx, tenantId).Execute()

Delete tenant



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
	tenantId := TODO // TenantId | The tenantId.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.Tenants.DeleteTenant(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tenants.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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


## GetTenant

> Tenant GetTenant(ctx, tenantId).Execute()

Retrieve tenant



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
	tenantId := TODO // TenantId | The tenantId.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.Tenants.GetTenant(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tenants.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `Tenants.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetTenants

> TenantCollection GetTenants(ctx).Execute()

List tenants



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

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.Tenants.GetTenants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tenants.GetTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenants`: TenantCollection
	fmt.Fprintf(os.Stdout, "Response from `Tenants.GetTenants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


### Return type

[**TenantCollection**](TenantCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## UpdateTenant

> Tenant UpdateTenant(ctx, tenantId).Tenant(tenant).Execute()

Update tenant



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
	tenantId := TODO // TenantId | The tenantId.
	tenant := *models.NewTenant() // Tenant | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.Tenants.UpdateTenant(context.Background(), tenantId).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tenants.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `Tenants.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

