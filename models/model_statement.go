package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Statement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Statement{}

// Statement struct for Statement
type Statement struct {
	Roles     []string   `json:"roles"`
	Resources []Resource `json:"resources"`
	// The list of users this statement applies to. Users and groups can be set at either the statement level or the record level, but not both.
	Users []User `json:"users,omitempty"`
	// The list of groups this statement applies to. Users in these groups will be receive access to the resources listed. Users and groups can be set at either the statement level or the record level, but not both.
	Groups []LinkedGroup `json:"groups,omitempty"`
}

type _Statement Statement

// NewStatement instantiates a new Statement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatement(roles []string, resources []Resource) *Statement {
	this := Statement{}
	this.Roles = roles
	this.Resources = resources
	return &this
}

// NewStatementWithDefaults instantiates a new Statement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementWithDefaults() *Statement {
	this := Statement{}
	return &this
}

// GetRoles returns the Roles field value
func (o *Statement) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *Statement) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *Statement) SetRoles(v []string) {
	o.Roles = v
}

// GetResources returns the Resources field value
func (o *Statement) GetResources() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *Statement) GetResourcesOk() ([]Resource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *Statement) SetResources(v []Resource) {
	o.Resources = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Statement) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statement) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Statement) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *Statement) SetUsers(v []User) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Statement) GetGroups() []LinkedGroup {
	if o == nil || IsNil(o.Groups) {
		var ret []LinkedGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statement) GetGroupsOk() ([]LinkedGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Statement) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []LinkedGroup and assigns it to the Groups field.
func (o *Statement) SetGroups(v []LinkedGroup) {
	o.Groups = v
}

func (o Statement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Statement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	toSerialize["resources"] = o.Resources
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

func (o *Statement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roles",
		"resources",
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

	varStatement := _Statement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStatement)

	if err != nil {
		return err
	}

	*o = Statement(varStatement)

	return err
}

type NullableStatement struct {
	value *Statement
	isSet bool
}

func (v NullableStatement) Get() *Statement {
	return v.value
}

func (v *NullableStatement) Set(val *Statement) {
	v.value = val
	v.isSet = true
}

func (v NullableStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatement(val *Statement) *NullableStatement {
	return &NullableStatement{value: val, isSet: true}
}

func (v NullableStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
