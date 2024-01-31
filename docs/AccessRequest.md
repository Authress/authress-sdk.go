# AccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique identifier for the request. | [readonly] 
**LastUpdated** | Pointer to **time.Time** | The expected last time the request was updated | [optional] [readonly] 
**Status** | Pointer to **string** | Current status of the access request. | [optional] [readonly] 
**Access** | [**AccessTemplate**](AccessTemplate.md) |  | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewAccessRequest

`func NewAccessRequest(requestId string, access AccessTemplate, ) *AccessRequest`

NewAccessRequest instantiates a new AccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestWithDefaults

`func NewAccessRequestWithDefaults() *AccessRequest`

NewAccessRequestWithDefaults instantiates a new AccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *AccessRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AccessRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AccessRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetLastUpdated

`func (o *AccessRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AccessRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AccessRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AccessRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *AccessRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccess

`func (o *AccessRequest) GetAccess() AccessTemplate`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccessRequest) GetAccessOk() (*AccessTemplate, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccessRequest) SetAccess(v AccessTemplate)`

SetAccess sets Access field to given value.


### GetLinks

`func (o *AccessRequest) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessRequest) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessRequest) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccessRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *AccessRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccessRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccessRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


