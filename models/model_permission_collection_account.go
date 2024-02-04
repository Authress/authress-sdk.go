package models

import (
	"encoding/json"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the PermissionCollectionAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCollectionAccount{}

// PermissionCollectionAccount struct for PermissionCollectionAccount
type PermissionCollectionAccount struct {
	AccountId *string `json:"accountId,omitempty"`
}

// NewPermissionCollectionAccount instantiates a new PermissionCollectionAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCollectionAccount() *PermissionCollectionAccount {
	this := PermissionCollectionAccount{}
	return &this
}

// NewPermissionCollectionAccountWithDefaults instantiates a new PermissionCollectionAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCollectionAccountWithDefaults() *PermissionCollectionAccount {
	this := PermissionCollectionAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PermissionCollectionAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionCollectionAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PermissionCollectionAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PermissionCollectionAccount) SetAccountId(v string) {
	o.AccountId = &v
}

func (o PermissionCollectionAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCollectionAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	return toSerialize, nil
}

type NullablePermissionCollectionAccount struct {
	value *PermissionCollectionAccount
	isSet bool
}

func (v NullablePermissionCollectionAccount) Get() *PermissionCollectionAccount {
	return v.value
}

func (v *NullablePermissionCollectionAccount) Set(val *PermissionCollectionAccount) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCollectionAccount) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCollectionAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCollectionAccount(val *PermissionCollectionAccount) *NullablePermissionCollectionAccount {
	return &NullablePermissionCollectionAccount{value: val, isSet: true}
}

func (v NullablePermissionCollectionAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCollectionAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
