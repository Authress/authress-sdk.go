# ExtensionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | [**[]Extension**](Extension.md) |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewExtensionCollection

`func NewExtensionCollection(extensions []Extension, ) *ExtensionCollection`

NewExtensionCollection instantiates a new ExtensionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionCollectionWithDefaults

`func NewExtensionCollectionWithDefaults() *ExtensionCollection`

NewExtensionCollectionWithDefaults instantiates a new ExtensionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *ExtensionCollection) GetExtensions() []Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ExtensionCollection) GetExtensionsOk() (*[]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ExtensionCollection) SetExtensions(v []Extension)`

SetExtensions sets Extensions field to given value.


### GetPagination

`func (o *ExtensionCollection) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ExtensionCollection) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ExtensionCollection) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ExtensionCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


