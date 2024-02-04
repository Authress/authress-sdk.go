# TenantCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | [**[]Tenant**](Tenant.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewTenantCollection

`func NewTenantCollection(tenants []Tenant, ) *TenantCollection`

NewTenantCollection instantiates a new TenantCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCollectionWithDefaults

`func NewTenantCollectionWithDefaults() *TenantCollection`

NewTenantCollectionWithDefaults instantiates a new TenantCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *TenantCollection) GetTenants() []Tenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *TenantCollection) GetTenantsOk() (*[]Tenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *TenantCollection) SetTenants(v []Tenant)`

SetTenants sets Tenants field to given value.


### GetPagination

`func (o *TenantCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TenantCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TenantCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TenantCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


