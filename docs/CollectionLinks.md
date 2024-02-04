# CollectionLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**NullableLink**](Link.md) |  | 
**Next** | Pointer to [**NullableLink**](Link.md) |  | [optional] 

## Methods

### NewCollectionLinks

`func NewCollectionLinks(self NullableLink, ) *CollectionLinks`

NewCollectionLinks instantiates a new CollectionLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionLinksWithDefaults

`func NewCollectionLinksWithDefaults() *CollectionLinks`

NewCollectionLinksWithDefaults instantiates a new CollectionLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *CollectionLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CollectionLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CollectionLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.


### SetSelfNil

`func (o *CollectionLinks) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *CollectionLinks) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetNext

`func (o *CollectionLinks) GetNext() Link`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CollectionLinks) GetNextOk() (*Link, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CollectionLinks) SetNext(v Link)`

SetNext sets Next field to given value.

### HasNext

`func (o *CollectionLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *CollectionLinks) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *CollectionLinks) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


