package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the UserRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRole{}

// UserRole A role with associated role data.
type UserRole struct {
	RoleId RoleId `json:"roleId"`
}

type _UserRole UserRole

// NewUserRole instantiates a new UserRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRole(roleId RoleId) *UserRole {
	this := UserRole{}
	this.RoleId = roleId
	return &this
}

// NewUserRoleWithDefaults instantiates a new UserRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRoleWithDefaults() *UserRole {
	this := UserRole{}
	return &this
}

// GetRoleId returns the RoleId field value
func (o *UserRole) GetRoleId() RoleId {
	if o == nil {
		var ret RoleId
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *UserRole) GetRoleIdOk() (*RoleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *UserRole) SetRoleId(v RoleId) {
	o.RoleId = v
}

func (o UserRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	return toSerialize, nil
}

func (o *UserRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roleId",
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

	varUserRole := _UserRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRole)

	if err != nil {
		return err
	}

	*o = UserRole(varUserRole)

	return err
}

type NullableUserRole struct {
	value *UserRole
	isSet bool
}

func (v NullableUserRole) Get() *UserRole {
	return v.value
}

func (v *NullableUserRole) Set(val *UserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRole(val *UserRole) *NullableUserRole {
	return &NullableUserRole{value: val, isSet: true}
}

func (v NullableUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
