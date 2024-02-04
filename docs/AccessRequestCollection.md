# AccessRequestCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]AccessRequest**](AccessRequest.md) | A list of access requests | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewAccessRequestCollection

`func NewAccessRequestCollection(links CollectionLinks, ) *AccessRequestCollection`

NewAccessRequestCollection instantiates a new AccessRequestCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestCollectionWithDefaults

`func NewAccessRequestCollectionWithDefaults() *AccessRequestCollection`

NewAccessRequestCollectionWithDefaults instantiates a new AccessRequestCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *AccessRequestCollection) GetRecords() []AccessRequest`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *AccessRequestCollection) GetRecordsOk() (*[]AccessRequest, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *AccessRequestCollection) SetRecords(v []AccessRequest)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *AccessRequestCollection) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetPagination

`func (o *AccessRequestCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AccessRequestCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AccessRequestCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AccessRequestCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLinks

`func (o *AccessRequestCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessRequestCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessRequestCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


