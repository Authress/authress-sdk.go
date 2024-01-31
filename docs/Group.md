# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique identifier for the groupId, can be specified on record creation. Must begin with grp_. | [optional] 
**Name** | **string** | A helpful name for this record | 
**LastUpdated** | Pointer to **time.Time** | The expected last time the group was updated | [optional] [readonly] 
**Users** | [**[]User**](User.md) | The list of users in this group | 
**Admins** | [**[]User**](User.md) | The list of admins that can edit this record even if they do not have global record edit permissions. | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewGroup

`func NewGroup(name string, users []User, admins []User, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Group) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Group) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Group) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Group) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetLastUpdated

`func (o *Group) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Group) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Group) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Group) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetUsers

`func (o *Group) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []User)`

SetUsers sets Users field to given value.


### GetAdmins

`func (o *Group) GetAdmins() []User`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *Group) GetAdminsOk() (*[]User, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *Group) SetAdmins(v []User)`

SetAdmins sets Admins field to given value.


### GetLinks

`func (o *Group) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Group) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Group) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Group) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *Group) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Group) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Group) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Group) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Group) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


