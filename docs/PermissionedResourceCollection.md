# PermissionedResourceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]PermissionedResource**](PermissionedResource.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewPermissionedResourceCollection

`func NewPermissionedResourceCollection(resources []PermissionedResource, links CollectionLinks, ) *PermissionedResourceCollection`

NewPermissionedResourceCollection instantiates a new PermissionedResourceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionedResourceCollectionWithDefaults

`func NewPermissionedResourceCollectionWithDefaults() *PermissionedResourceCollection`

NewPermissionedResourceCollectionWithDefaults instantiates a new PermissionedResourceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *PermissionedResourceCollection) GetResources() []PermissionedResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionedResourceCollection) GetResourcesOk() (*[]PermissionedResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionedResourceCollection) SetResources(v []PermissionedResource)`

SetResources sets Resources field to given value.


### GetPagination

`func (o *PermissionedResourceCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PermissionedResourceCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PermissionedResourceCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PermissionedResourceCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *PermissionedResourceCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PermissionedResourceCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PermissionedResourceCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


