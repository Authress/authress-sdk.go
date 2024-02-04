package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserIdentityCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentityCollection{}

// UserIdentityCollection A collection of user identities
type UserIdentityCollection struct {
	// A list of users
	Users []UserIdentity `json:"users"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Links CollectionLinks `json:"links"`
}

type _UserIdentityCollection UserIdentityCollection

// NewUserIdentityCollection instantiates a new UserIdentityCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentityCollection(users []UserIdentity, links CollectionLinks) *UserIdentityCollection {
	this := UserIdentityCollection{}
	this.Users = users
	this.Links = links
	return &this
}

// NewUserIdentityCollectionWithDefaults instantiates a new UserIdentityCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityCollectionWithDefaults() *UserIdentityCollection {
	this := UserIdentityCollection{}
	return &this
}

// GetUsers returns the Users field value
func (o *UserIdentityCollection) GetUsers() []UserIdentity {
	if o == nil {
		var ret []UserIdentity
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UserIdentityCollection) GetUsersOk() ([]UserIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UserIdentityCollection) SetUsers(v []UserIdentity) {
	o.Users = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UserIdentityCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentityCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UserIdentityCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *UserIdentityCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *UserIdentityCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserIdentityCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserIdentityCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o UserIdentityCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentityCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *UserIdentityCollection) UnmarshalJSON(data []byte) (err error) {
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
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserIdentityCollection := _UserIdentityCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserIdentityCollection)

	if err != nil {
		return err
	}

	*o = UserIdentityCollection(varUserIdentityCollection)

	return err
}

type NullableUserIdentityCollection struct {
	value *UserIdentityCollection
	isSet bool
}

func (v NullableUserIdentityCollection) Get() *UserIdentityCollection {
	return v.value
}

func (v *NullableUserIdentityCollection) Set(val *UserIdentityCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentityCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentityCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentityCollection(val *UserIdentityCollection) *NullableUserIdentityCollection {
	return &NullableUserIdentityCollection{value: val, isSet: true}
}

func (v NullableUserIdentityCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentityCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


