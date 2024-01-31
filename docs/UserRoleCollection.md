# UserRoleCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | [**UserId**](UserId.md) |  | 
**Roles** | [**[]UserRole**](UserRole.md) | A list of the roles | 

## Methods

### NewUserRoleCollection

`func NewUserRoleCollection(userId UserId, roles []UserRole, ) *UserRoleCollection`

NewUserRoleCollection instantiates a new UserRoleCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleCollectionWithDefaults

`func NewUserRoleCollectionWithDefaults() *UserRoleCollection`

NewUserRoleCollectionWithDefaults instantiates a new UserRoleCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserRoleCollection) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserRoleCollection) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserRoleCollection) SetUserId(v UserId)`

SetUserId sets UserId field to given value.


### GetRoles

`func (o *UserRoleCollection) GetRoles() []UserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserRoleCollection) GetRolesOk() (*[]UserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserRoleCollection) SetRoles(v []UserRole)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


