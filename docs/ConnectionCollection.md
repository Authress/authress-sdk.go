# ConnectionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]Connection**](Connection.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewConnectionCollection

`func NewConnectionCollection(connections []Connection, ) *ConnectionCollection`

NewConnectionCollection instantiates a new ConnectionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCollectionWithDefaults

`func NewConnectionCollectionWithDefaults() *ConnectionCollection`

NewConnectionCollectionWithDefaults instantiates a new ConnectionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ConnectionCollection) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ConnectionCollection) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ConnectionCollection) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetPagination

`func (o *ConnectionCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ConnectionCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ConnectionCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ConnectionCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


