# Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteId** | **string** | The unique identifier for the invite. Use this ID to accept the invite. This parameter is ignored during invite creation. | [readonly] 
**TenantId** | Pointer to [**TenantId**](TenantId.md) |  | [optional] 
**Statements** | [**[]Statement**](Statement.md) | A list of statements which match roles to resources. The invited user will all statements apply to them when the invite is accepted. | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 

## Methods

### NewInvite

`func NewInvite(inviteId string, statements []Statement, ) *Invite`

NewInvite instantiates a new Invite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteWithDefaults

`func NewInviteWithDefaults() *Invite`

NewInviteWithDefaults instantiates a new Invite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteId

`func (o *Invite) GetInviteId() string`

GetInviteId returns the InviteId field if non-nil, zero value otherwise.

### GetInviteIdOk

`func (o *Invite) GetInviteIdOk() (*string, bool)`

GetInviteIdOk returns a tuple with the InviteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteId

`func (o *Invite) SetInviteId(v string)`

SetInviteId sets InviteId field to given value.


### GetTenantId

`func (o *Invite) GetTenantId() TenantId`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Invite) GetTenantIdOk() (*TenantId, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Invite) SetTenantId(v TenantId)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Invite) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetStatements

`func (o *Invite) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *Invite) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *Invite) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.


### GetLinks

`func (o *Invite) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Invite) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Invite) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Invite) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


