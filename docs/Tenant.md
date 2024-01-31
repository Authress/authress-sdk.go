# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** |  | [optional] 
**TenantLookupIdentifier** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to [**TenantData**](TenantData.md) |  | [optional] 
**Connection** | Pointer to [**NullableTenantConnection**](TenantConnection.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *Tenant) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Tenant) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Tenant) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Tenant) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Tenant) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Tenant) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantLookupIdentifier

`func (o *Tenant) GetTenantLookupIdentifier() string`

GetTenantLookupIdentifier returns the TenantLookupIdentifier field if non-nil, zero value otherwise.

### GetTenantLookupIdentifierOk

`func (o *Tenant) GetTenantLookupIdentifierOk() (*string, bool)`

GetTenantLookupIdentifierOk returns a tuple with the TenantLookupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantLookupIdentifier

`func (o *Tenant) SetTenantLookupIdentifier(v string)`

SetTenantLookupIdentifier sets TenantLookupIdentifier field to given value.

### HasTenantLookupIdentifier

`func (o *Tenant) HasTenantLookupIdentifier() bool`

HasTenantLookupIdentifier returns a boolean if a field has been set.

### SetTenantLookupIdentifierNil

`func (o *Tenant) SetTenantLookupIdentifierNil(b bool)`

 SetTenantLookupIdentifierNil sets the value for TenantLookupIdentifier to be an explicit nil

### UnsetTenantLookupIdentifier
`func (o *Tenant) UnsetTenantLookupIdentifier()`

UnsetTenantLookupIdentifier ensures that no value is present for TenantLookupIdentifier, not even an explicit nil
### GetData

`func (o *Tenant) GetData() TenantData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Tenant) GetDataOk() (*TenantData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Tenant) SetData(v TenantData)`

SetData sets Data field to given value.

### HasData

`func (o *Tenant) HasData() bool`

HasData returns a boolean if a field has been set.

### GetConnection

`func (o *Tenant) GetConnection() TenantConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Tenant) GetConnectionOk() (*TenantConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Tenant) SetConnection(v TenantConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *Tenant) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### SetConnectionNil

`func (o *Tenant) SetConnectionNil(b bool)`

 SetConnectionNil sets the value for Connection to be an explicit nil

### UnsetConnection
`func (o *Tenant) UnsetConnection()`

UnsetConnection ensures that no value is present for Connection, not even an explicit nil
### GetCreatedTime

`func (o *Tenant) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Tenant) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Tenant) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Tenant) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


