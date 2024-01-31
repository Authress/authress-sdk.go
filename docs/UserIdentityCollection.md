# UserIdentityCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserIdentity**](UserIdentity.md) | A list of users | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewUserIdentityCollection

`func NewUserIdentityCollection(users []UserIdentity, links CollectionLinks, ) *UserIdentityCollection`

NewUserIdentityCollection instantiates a new UserIdentityCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityCollectionWithDefaults

`func NewUserIdentityCollectionWithDefaults() *UserIdentityCollection`

NewUserIdentityCollectionWithDefaults instantiates a new UserIdentityCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UserIdentityCollection) GetUsers() []UserIdentity`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserIdentityCollection) GetUsersOk() (*[]UserIdentity, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserIdentityCollection) SetUsers(v []UserIdentity)`

SetUsers sets Users field to given value.


### GetPagination

`func (o *UserIdentityCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UserIdentityCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UserIdentityCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UserIdentityCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *UserIdentityCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserIdentityCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserIdentityCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


