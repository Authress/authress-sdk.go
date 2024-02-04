# AccessTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]User**](User.md) | The list of users the access applies to | 
**Statements** | [**[]Statement**](Statement.md) | A list of statements which match roles to resources. | 

## Methods

### NewAccessTemplate

`func NewAccessTemplate(users []User, statements []Statement, ) *AccessTemplate`

NewAccessTemplate instantiates a new AccessTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTemplateWithDefaults

`func NewAccessTemplateWithDefaults() *AccessTemplate`

NewAccessTemplateWithDefaults instantiates a new AccessTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *AccessTemplate) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccessTemplate) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccessTemplate) SetUsers(v []User)`

SetUsers sets Users field to given value.


### GetStatements

`func (o *AccessTemplate) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccessTemplate) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccessTemplate) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


