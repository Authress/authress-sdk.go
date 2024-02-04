package authress

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Group{}

// Group A group of users, which can be added to access records.
type Group struct {
	// Unique identifier for the groupId, can be specified on record creation. Must begin with grp_.
	GroupId *string `json:"groupId,omitempty"`
	// A helpful name for this record
	Name string `json:"name"`
	// The expected last time the group was updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The list of users in this group
	Users []User `json:"users"`
	// The list of admins that can edit this record even if they do not have global record edit permissions.
	Admins []User `json:"admins"`
	Links *AccountLinks `json:"links,omitempty"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(name string, users []User, admins []User) *Group {
	this := Group{}
	this.Name = name
	this.Users = users
	this.Admins = admins
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *Group) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *Group) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *Group) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value
func (o *Group) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Group) SetName(v string) {
	o.Name = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Group) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Group) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Group) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetUsers returns the Users field value
func (o *Group) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *Group) GetUsersOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *Group) SetUsers(v []User) {
	o.Users = v
}

// GetAdmins returns the Admins field value
func (o *Group) GetAdmins() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value
// and a boolean to check if the value has been set.
func (o *Group) GetAdminsOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Admins, true
}

// SetAdmins sets field value
func (o *Group) SetAdmins(v []User) {
	o.Admins = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() AccountLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *Group) SetLinks(v AccountLinks) {
	o.Links = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Group) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Group) SetTags(v map[string]string) {
	o.Tags = v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["users"] = o.Users
	toSerialize["admins"] = o.Admins
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *Group) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"users",
		"admins",
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

	varGroup := _Group{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroup)

	if err != nil {
		return err
	}

	*o = Group(varGroup)

	return err
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


