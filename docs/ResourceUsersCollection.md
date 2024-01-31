# ResourceUsersCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserRoleCollection**](UserRoleCollection.md) | A list of users | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewResourceUsersCollection

`func NewResourceUsersCollection(users []UserRoleCollection, links CollectionLinks, ) *ResourceUsersCollection`

NewResourceUsersCollection instantiates a new ResourceUsersCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsersCollectionWithDefaults

`func NewResourceUsersCollectionWithDefaults() *ResourceUsersCollection`

NewResourceUsersCollectionWithDefaults instantiates a new ResourceUsersCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ResourceUsersCollection) GetUsers() []UserRoleCollection`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResourceUsersCollection) GetUsersOk() (*[]UserRoleCollection, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResourceUsersCollection) SetUsers(v []UserRoleCollection)`

SetUsers sets Users field to given value.


### GetPagination

`func (o *ResourceUsersCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ResourceUsersCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ResourceUsersCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ResourceUsersCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceUsersCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceUsersCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceUsersCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


