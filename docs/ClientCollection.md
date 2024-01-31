# ClientCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | [**[]Client**](Client.md) | A list of clients | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewClientCollection

`func NewClientCollection(clients []Client, links CollectionLinks, ) *ClientCollection`

NewClientCollection instantiates a new ClientCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCollectionWithDefaults

`func NewClientCollectionWithDefaults() *ClientCollection`

NewClientCollectionWithDefaults instantiates a new ClientCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ClientCollection) GetClients() []Client`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientCollection) GetClientsOk() (*[]Client, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientCollection) SetClients(v []Client)`

SetClients sets Clients field to given value.


### GetPagination

`func (o *ClientCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ClientCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ClientCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ClientCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *ClientCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClientCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClientCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


