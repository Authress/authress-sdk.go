# ClaimRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionResource** | **string** | The parent resource to add a sub-resource to. The resource must have a resource configuration that add the permission CLAIM for this to work. | 
**ResourceId** | **string** | The sub-resource the user is requesting Admin ownership over. | 

## Methods

### NewClaimRequest

`func NewClaimRequest(collectionResource string, resourceId string, ) *ClaimRequest`

NewClaimRequest instantiates a new ClaimRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimRequestWithDefaults

`func NewClaimRequestWithDefaults() *ClaimRequest`

NewClaimRequestWithDefaults instantiates a new ClaimRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionResource

`func (o *ClaimRequest) GetCollectionResource() string`

GetCollectionResource returns the CollectionResource field if non-nil, zero value otherwise.

### GetCollectionResourceOk

`func (o *ClaimRequest) GetCollectionResourceOk() (*string, bool)`

GetCollectionResourceOk returns a tuple with the CollectionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionResource

`func (o *ClaimRequest) SetCollectionResource(v string)`

SetCollectionResource sets CollectionResource field to given value.


### GetResourceId

`func (o *ClaimRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ClaimRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ClaimRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


