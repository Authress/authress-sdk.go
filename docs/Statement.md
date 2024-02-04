# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** |  | 
**Resources** | [**[]Resource**](Resource.md) |  | 
**Users** | Pointer to [**[]User**](User.md) | The list of users this statement applies to. Users and groups can be set at either the statement level or the record level, but not both. | [optional] 
**Groups** | Pointer to [**[]LinkedGroup**](LinkedGroup.md) | The list of groups this statement applies to. Users in these groups will be receive access to the resources listed. Users and groups can be set at either the statement level or the record level, but not both. | [optional] 

## Methods

### NewStatement

`func NewStatement(roles []string, resources []Resource, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *Statement) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Statement) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Statement) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetResources

`func (o *Statement) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Statement) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Statement) SetResources(v []Resource)`

SetResources sets Resources field to given value.


### GetUsers

`func (o *Statement) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Statement) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Statement) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Statement) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *Statement) GetGroups() []LinkedGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Statement) GetGroupsOk() (*[]LinkedGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Statement) SetGroups(v []LinkedGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Statement) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


