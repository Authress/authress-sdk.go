package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the UserRoleCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRoleCollection{}

// UserRoleCollection A collect of roles that the user has to a resource.
type UserRoleCollection struct {
	UserId UserId `json:"userId"`
	// A list of the roles
	Roles []UserRole `json:"roles"`
}

type _UserRoleCollection UserRoleCollection

// NewUserRoleCollection instantiates a new UserRoleCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRoleCollection(userId UserId, roles []UserRole) *UserRoleCollection {
	this := UserRoleCollection{}
	this.UserId = userId
	this.Roles = roles
	return &this
}

// NewUserRoleCollectionWithDefaults instantiates a new UserRoleCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRoleCollectionWithDefaults() *UserRoleCollection {
	this := UserRoleCollection{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserRoleCollection) GetUserId() UserId {
	if o == nil {
		var ret UserId
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserRoleCollection) GetUserIdOk() (*UserId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserRoleCollection) SetUserId(v UserId) {
	o.UserId = v
}

// GetRoles returns the Roles field value
func (o *UserRoleCollection) GetRoles() []UserRole {
	if o == nil {
		var ret []UserRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserRoleCollection) GetRolesOk() ([]UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UserRoleCollection) SetRoles(v []UserRole) {
	o.Roles = v
}

func (o UserRoleCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRoleCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *UserRoleCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"roles",
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

	varUserRoleCollection := _UserRoleCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRoleCollection)

	if err != nil {
		return err
	}

	*o = UserRoleCollection(varUserRoleCollection)

	return err
}

type NullableUserRoleCollection struct {
	value *UserRoleCollection
	isSet bool
}

func (v NullableUserRoleCollection) Get() *UserRoleCollection {
	return v.value
}

func (v *NullableUserRoleCollection) Set(val *UserRoleCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoleCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoleCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoleCollection(val *UserRoleCollection) *NullableUserRoleCollection {
	return &NullableUserRoleCollection{value: val, isSet: true}
}

func (v NullableUserRoleCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoleCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
