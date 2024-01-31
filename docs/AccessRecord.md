# AccessRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | Pointer to **string** | Unique identifier for the record, can be specified on record creation. | [optional] 
**Name** | **string** | A helpful name for this record | 
**Description** | Pointer to **NullableString** | More details about this record | [optional] 
**Capacity** | Pointer to **float32** | Percentage capacity of record that is filled. | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The expected last time the record was updated | [optional] [readonly] 
**Status** | Pointer to **string** | Current status of the access record. | [optional] [readonly] 
**Account** | Pointer to [**AccessRecordAccount**](AccessRecordAccount.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | The list of users this record applies to | [optional] 
**Admins** | Pointer to [**[]User**](User.md) | The list of admin that can edit this record even if they do not have global record edit permissions. | [optional] 
**Groups** | Pointer to [**[]LinkedGroup**](LinkedGroup.md) | The list of groups this record applies to. Users in these groups will be receive access to the resources listed. | [optional] 
**Statements** | [**[]Statement**](Statement.md) | A list of statements which match roles to resources. | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewAccessRecord

`func NewAccessRecord(name string, statements []Statement, ) *AccessRecord`

NewAccessRecord instantiates a new AccessRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRecordWithDefaults

`func NewAccessRecordWithDefaults() *AccessRecord`

NewAccessRecordWithDefaults instantiates a new AccessRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *AccessRecord) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *AccessRecord) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *AccessRecord) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *AccessRecord) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetName

`func (o *AccessRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessRecord) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AccessRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessRecord) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessRecord) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapacity

`func (o *AccessRecord) GetCapacity() float32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AccessRecord) GetCapacityOk() (*float32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AccessRecord) SetCapacity(v float32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AccessRecord) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AccessRecord) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AccessRecord) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AccessRecord) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AccessRecord) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *AccessRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccount

`func (o *AccessRecord) GetAccount() AccessRecordAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccessRecord) GetAccountOk() (*AccessRecordAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccessRecord) SetAccount(v AccessRecordAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccessRecord) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUsers

`func (o *AccessRecord) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccessRecord) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccessRecord) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AccessRecord) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *AccessRecord) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *AccessRecord) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetAdmins

`func (o *AccessRecord) GetAdmins() []User`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *AccessRecord) GetAdminsOk() (*[]User, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *AccessRecord) SetAdmins(v []User)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *AccessRecord) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### SetAdminsNil

`func (o *AccessRecord) SetAdminsNil(b bool)`

 SetAdminsNil sets the value for Admins to be an explicit nil

### UnsetAdmins
`func (o *AccessRecord) UnsetAdmins()`

UnsetAdmins ensures that no value is present for Admins, not even an explicit nil
### GetGroups

`func (o *AccessRecord) GetGroups() []LinkedGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccessRecord) GetGroupsOk() (*[]LinkedGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccessRecord) SetGroups(v []LinkedGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AccessRecord) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *AccessRecord) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *AccessRecord) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetStatements

`func (o *AccessRecord) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccessRecord) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccessRecord) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.


### GetLinks

`func (o *AccessRecord) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccessRecord) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccessRecord) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccessRecord) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTags

`func (o *AccessRecord) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessRecord) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessRecord) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccessRecord) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccessRecord) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


