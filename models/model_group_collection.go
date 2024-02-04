package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the GroupCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCollection{}

// GroupCollection A collection of groups
type GroupCollection struct {
	// A list of groups
	Groups     []Group         `json:"groups"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Links      CollectionLinks `json:"links"`
}

type _GroupCollection GroupCollection

// NewGroupCollection instantiates a new GroupCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCollection(groups []Group, links CollectionLinks) *GroupCollection {
	this := GroupCollection{}
	this.Groups = groups
	this.Links = links
	return &this
}

// NewGroupCollectionWithDefaults instantiates a new GroupCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCollectionWithDefaults() *GroupCollection {
	this := GroupCollection{}
	return &this
}

// GetGroups returns the Groups field value
func (o *GroupCollection) GetGroups() []Group {
	if o == nil {
		var ret []Group
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetGroupsOk() ([]Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *GroupCollection) SetGroups(v []Group) {
	o.Groups = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GroupCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GroupCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GroupCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *GroupCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GroupCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GroupCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GroupCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groups",
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

	varGroupCollection := _GroupCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupCollection)

	if err != nil {
		return err
	}

	*o = GroupCollection(varGroupCollection)

	return err
}

type NullableGroupCollection struct {
	value *GroupCollection
	isSet bool
}

func (v NullableGroupCollection) Get() *GroupCollection {
	return v.value
}

func (v *NullableGroupCollection) Set(val *GroupCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCollection(val *GroupCollection) *NullableGroupCollection {
	return &NullableGroupCollection{value: val, isSet: true}
}

func (v NullableGroupCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
