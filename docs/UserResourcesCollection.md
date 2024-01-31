# UserResourcesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**PermissionCollectionAccount**](PermissionCollectionAccount.md) |  | [optional] 
**UserId** | [**UserId**](UserId.md) |  | 
**Resources** | Pointer to [**[]Resource**](Resource.md) | A list of the resources the user has some permission to. | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewUserResourcesCollection

`func NewUserResourcesCollection(userId UserId, links CollectionLinks, ) *UserResourcesCollection`

NewUserResourcesCollection instantiates a new UserResourcesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResourcesCollectionWithDefaults

`func NewUserResourcesCollectionWithDefaults() *UserResourcesCollection`

NewUserResourcesCollectionWithDefaults instantiates a new UserResourcesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserResourcesCollection) GetAccount() PermissionCollectionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserResourcesCollection) GetAccountOk() (*PermissionCollectionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserResourcesCollection) SetAccount(v PermissionCollectionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserResourcesCollection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UserResourcesCollection) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserResourcesCollection) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserResourcesCollection) SetUserId(v UserId)`

SetUserId sets UserId field to given value.


### GetResources

`func (o *UserResourcesCollection) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserResourcesCollection) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserResourcesCollection) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UserResourcesCollection) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetPagination

`func (o *UserResourcesCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UserResourcesCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UserResourcesCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UserResourcesCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *UserResourcesCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserResourcesCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserResourcesCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


