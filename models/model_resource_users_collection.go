package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ResourceUsersCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceUsersCollection{}

// ResourceUsersCollection A collection of users with explicit permission to a resource.
type ResourceUsersCollection struct {
	// A list of users
	Users      []UserRoleCollection `json:"users"`
	Pagination *Pagination          `json:"pagination,omitempty"`
	Links      CollectionLinks      `json:"links"`
}

type _ResourceUsersCollection ResourceUsersCollection

// NewResourceUsersCollection instantiates a new ResourceUsersCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceUsersCollection(users []UserRoleCollection, links CollectionLinks) *ResourceUsersCollection {
	this := ResourceUsersCollection{}
	this.Users = users
	this.Links = links
	return &this
}

// NewResourceUsersCollectionWithDefaults instantiates a new ResourceUsersCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUsersCollectionWithDefaults() *ResourceUsersCollection {
	this := ResourceUsersCollection{}
	return &this
}

// GetUsers returns the Users field value
func (o *ResourceUsersCollection) GetUsers() []UserRoleCollection {
	if o == nil {
		var ret []UserRoleCollection
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ResourceUsersCollection) GetUsersOk() ([]UserRoleCollection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ResourceUsersCollection) SetUsers(v []UserRoleCollection) {
	o.Users = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ResourceUsersCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceUsersCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ResourceUsersCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ResourceUsersCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *ResourceUsersCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceUsersCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceUsersCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o ResourceUsersCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceUsersCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ResourceUsersCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varResourceUsersCollection := _ResourceUsersCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varResourceUsersCollection)

	if err != nil {
		return err
	}

	*o = ResourceUsersCollection(varResourceUsersCollection)

	return err
}

type NullableResourceUsersCollection struct {
	value *ResourceUsersCollection
	isSet bool
}

func (v NullableResourceUsersCollection) Get() *ResourceUsersCollection {
	return v.value
}

func (v *NullableResourceUsersCollection) Set(val *ResourceUsersCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceUsersCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUsersCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUsersCollection(val *ResourceUsersCollection) *NullableResourceUsersCollection {
	return &NullableResourceUsersCollection{value: val, isSet: true}
}

func (v NullableResourceUsersCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUsersCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
