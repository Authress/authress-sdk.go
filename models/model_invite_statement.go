package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Statement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteStatement{}

// InviteStatement struct for InviteStatement
type InviteStatement struct {
	Roles     []string   `json:"roles"`
	Resources []Resource `json:"resources"`
}

type _InviteStatement InviteStatement

// NewInviteStatement instantiates a new InviteStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteStatement(roles []string, resources []Resource) *InviteStatement {
	this := InviteStatement{}
	this.Roles = roles
	this.Resources = resources
	return &this
}

// NewInviteStatementWithDefaults instantiates a new InviteStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteStatementWithDefaults() *InviteStatement {
	this := InviteStatement{}
	return &this
}

// GetRoles returns the Roles field value
func (o *InviteStatement) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *InviteStatement) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *InviteStatement) SetRoles(v []string) {
	o.Roles = v
}

// GetResources returns the Resources field value
func (o *InviteStatement) GetResources() []Resource {
	if o == nil {
		var ret []Resource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *InviteStatement) GetResourcesOk() ([]Resource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *InviteStatement) SetResources(v []Resource) {
	o.Resources = v
}

func (o InviteStatement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteStatement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

func (o *InviteStatement) UnmarshalJSON(data []byte) (err error) {
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

	varInviteStatement := _InviteStatement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInviteStatement)

	if err != nil {
		return err
	}

	*o = InviteStatement(varInviteStatement)

	return err
}

type NullableInviteStatement struct {
	value *InviteStatement
	isSet bool
}

func (v NullableInviteStatement) Get() *InviteStatement {
	return v.value
}

func (v *NullableInviteStatement) Set(val *InviteStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteStatement(val *InviteStatement) *NullableInviteStatement {
	return &NullableInviteStatement{value: val, isSet: true}
}

func (v NullableInviteStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
