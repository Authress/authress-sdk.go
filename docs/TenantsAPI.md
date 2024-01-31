# \TenantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantsAPI.md#CreateTenant) | **Post** /v1/tenants | Create tenant
[**DeleteTenant**](TenantsAPI.md#DeleteTenant) | **Delete** /v1/tenants/{tenantId} | Delete tenant
[**GetTenant**](TenantsAPI.md#GetTenant) | **Get** /v1/tenants/{tenantId} | Retrieve tenant
[**GetTenants**](TenantsAPI.md#GetTenants) | **Get** /v1/tenants | List tenants
[**UpdateTenant**](TenantsAPI.md#UpdateTenant) | **Put** /v1/tenants/{tenantId} | Update tenant



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
	openapiclient "//"
)

func main() {
	tenant := *openapiclient.NewTenant() // Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.CreateTenant(context.Background()).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenant`: %v\n", resp)
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

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "//"
)

func main() {
	tenantId := TODO // TenantId | The tenantId.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantsAPI.DeleteTenant(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | [**TenantId**](.md) | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


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
	openapiclient "//"
)

func main() {
	tenantId := TODO // TenantId | The tenantId.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.GetTenant(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | [**TenantId**](.md) | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "//"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.GetTenants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenants`: TenantCollection
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


### Return type

[**TenantCollection**](TenantCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "//"
)

func main() {
	tenantId := TODO // TenantId | The tenantId.
	tenant := *openapiclient.NewTenant() // Tenant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.UpdateTenant(context.Background(), tenantId).Tenant(tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | [**TenantId**](.md) | The tenantId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenant** | [**Tenant**](Tenant.md) |  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

