package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the PermissionCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCollection{}

// PermissionCollection A collect of permissions that the user has to a resource.
type PermissionCollection struct {
	Account *PermissionCollectionAccount `json:"account,omitempty"`
	UserId  UserId                       `json:"userId"`
	// A list of the permissions
	Permissions []PermissionObject `json:"permissions"`
}

type _PermissionCollection PermissionCollection

// NewPermissionCollection instantiates a new PermissionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCollection(userId UserId, permissions []PermissionObject) *PermissionCollection {
	this := PermissionCollection{}
	this.UserId = userId
	this.Permissions = permissions
	return &this
}

// NewPermissionCollectionWithDefaults instantiates a new PermissionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCollectionWithDefaults() *PermissionCollection {
	this := PermissionCollection{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *PermissionCollection) GetAccount() PermissionCollectionAccount {
	if o == nil || IsNil(o.Account) {
		var ret PermissionCollectionAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionCollection) GetAccountOk() (*PermissionCollectionAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *PermissionCollection) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given PermissionCollectionAccount and assigns it to the Account field.
func (o *PermissionCollection) SetAccount(v PermissionCollectionAccount) {
	o.Account = &v
}

// GetUserId returns the UserId field value
func (o *PermissionCollection) GetUserId() UserId {
	if o == nil {
		var ret UserId
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PermissionCollection) GetUserIdOk() (*UserId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PermissionCollection) SetUserId(v UserId) {
	o.UserId = v
}

// GetPermissions returns the Permissions field value
func (o *PermissionCollection) GetPermissions() []PermissionObject {
	if o == nil {
		var ret []PermissionObject
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionCollection) GetPermissionsOk() ([]PermissionObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionCollection) SetPermissions(v []PermissionObject) {
	o.Permissions = v
}

func (o PermissionCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["userId"] = o.UserId
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

func (o *PermissionCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"permissions",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPermissionCollection := _PermissionCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionCollection)

	if err != nil {
		return err
	}

	*o = PermissionCollection(varPermissionCollection)

	return err
}

type NullablePermissionCollection struct {
	value *PermissionCollection
	isSet bool
}

func (v NullablePermissionCollection) Get() *PermissionCollection {
	return v.value
}

func (v *NullablePermissionCollection) Set(val *PermissionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCollection(val *PermissionCollection) *NullablePermissionCollection {
	return &NullablePermissionCollection{value: val, isSet: true}
}

func (v NullablePermissionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
