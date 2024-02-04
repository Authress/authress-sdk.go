package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the RoleCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleCollection{}

// RoleCollection A collection of roles
type RoleCollection struct {
	// A list of roles
	Roles      []Role          `json:"roles"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Links      CollectionLinks `json:"links"`
}

type _RoleCollection RoleCollection

// NewRoleCollection instantiates a new RoleCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCollection(roles []Role, links CollectionLinks) *RoleCollection {
	this := RoleCollection{}
	this.Roles = roles
	this.Links = links
	return &this
}

// NewRoleCollectionWithDefaults instantiates a new RoleCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCollectionWithDefaults() *RoleCollection {
	this := RoleCollection{}
	return &this
}

// GetRoles returns the Roles field value
func (o *RoleCollection) GetRoles() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RoleCollection) GetRolesOk() ([]Role, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RoleCollection) SetRoles(v []Role) {
	o.Roles = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *RoleCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *RoleCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *RoleCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *RoleCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RoleCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RoleCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o RoleCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *RoleCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roles",
		"links",
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

	varRoleCollection := _RoleCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoleCollection)

	if err != nil {
		return err
	}

	*o = RoleCollection(varRoleCollection)

	return err
}

type NullableRoleCollection struct {
	value *RoleCollection
	isSet bool
}

func (v NullableRoleCollection) Get() *RoleCollection {
	return v.value
}

func (v *NullableRoleCollection) Set(val *RoleCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCollection(val *RoleCollection) *NullableRoleCollection {
	return &NullableRoleCollection{value: val, isSet: true}
}

func (v NullableRoleCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
