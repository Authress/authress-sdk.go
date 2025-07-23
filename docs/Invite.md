# Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteId** | **string** | The unique identifier for the invite. Use this ID to accept the invite. This parameter is ignored during invite creation. | [readonly] 
**DefaultLoginTenantId** | **string** |  | [optional] 
**LinkedUserId** | **string** | Specify a User ID that logging in user should receive when login completes. This ID is used to automatically assign a user ID to the user rather than a dynamically generated Authress User ID when using the Authress Login UI SDK. This parameter is ignored when accepting invites directly. Note: If the user logging in has already signed up, then this parameter is ignored. | [optional] 
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


### GetDefaultLoginTenantId

`func (o *Invite) GetDefaultLoginTenantId() string`

GetDefaultLoginTenantId returns the DefaultLoginTenantId field if non-nil, zero value otherwise.

### GetDefaultLoginTenantIdOk

`func (o *Invite) GetDefaultLoginTenantIdOk() (*string, bool)`

GetDefaultLoginTenantIdOk returns a tuple with the DefaultLoginTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLoginTenantId

`func (o *Invite) SetDefaultLoginTenantId(v string)`

SetDefaultLoginTenantId sets DefaultLoginTenantId field to given value.

### HasDefaultLoginTenantId

`func (o *Invite) HasDefaultLoginTenantId() bool`

HasDefaultLoginTenantId returns a boolean if a field has been set.

### GetLinkedUserId

`func (o *Invite) GetLinkedUserId() string`

GetLinkedUserId returns the LinkedUserId field if non-nil, zero value otherwise.

### GetLinkedUserIdOk

`func (o *Invite) GetLinkedUserIdOk() (*string, bool)`

GetLinkedUserIdOk returns a tuple with the LinkedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedUserId

`func (o *Invite) SetLinkedUserId(v string)`

SetLinkedUserId sets LinkedUserId field to given value.

### HasLinkedUserId

`func (o *Invite) HasLinkedUserId() bool`

HasLinkedUserId returns a boolean if a field has been set.

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


