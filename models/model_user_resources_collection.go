package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the UserResourcesCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResourcesCollection{}

// UserResourcesCollection A collect of permissions that the user has to a resource.
type UserResourcesCollection struct {
	Account *PermissionCollectionAccount `json:"account,omitempty"`
	UserId  UserId                       `json:"userId"`
	// A list of the resources the user has some permission to.
	Resources  []Resource      `json:"resources,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Links      CollectionLinks `json:"links"`
}

type _UserResourcesCollection UserResourcesCollection

// NewUserResourcesCollection instantiates a new UserResourcesCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResourcesCollection(userId UserId, links CollectionLinks) *UserResourcesCollection {
	this := UserResourcesCollection{}
	this.UserId = userId
	this.Links = links
	return &this
}

// NewUserResourcesCollectionWithDefaults instantiates a new UserResourcesCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResourcesCollectionWithDefaults() *UserResourcesCollection {
	this := UserResourcesCollection{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UserResourcesCollection) GetAccount() PermissionCollectionAccount {
	if o == nil || IsNil(o.Account) {
		var ret PermissionCollectionAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourcesCollection) GetAccountOk() (*PermissionCollectionAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UserResourcesCollection) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given PermissionCollectionAccount and assigns it to the Account field.
func (o *UserResourcesCollection) SetAccount(v PermissionCollectionAccount) {
	o.Account = &v
}

// GetUserId returns the UserId field value
func (o *UserResourcesCollection) GetUserId() UserId {
	if o == nil {
		var ret UserId
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserResourcesCollection) GetUserIdOk() (*UserId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserResourcesCollection) SetUserId(v UserId) {
	o.UserId = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UserResourcesCollection) GetResources() []Resource {
	if o == nil || IsNil(o.Resources) {
		var ret []Resource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourcesCollection) GetResourcesOk() ([]Resource, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *UserResourcesCollection) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *UserResourcesCollection) SetResources(v []Resource) {
	o.Resources = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *UserResourcesCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResourcesCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *UserResourcesCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *UserResourcesCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *UserResourcesCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserResourcesCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserResourcesCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o UserResourcesCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResourcesCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *UserResourcesCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varUserResourcesCollection := _UserResourcesCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserResourcesCollection)

	if err != nil {
		return err
	}

	*o = UserResourcesCollection(varUserResourcesCollection)

	return err
}

type NullableUserResourcesCollection struct {
	value *UserResourcesCollection
	isSet bool
}

func (v NullableUserResourcesCollection) Get() *UserResourcesCollection {
	return v.value
}

func (v *NullableUserResourcesCollection) Set(val *UserResourcesCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResourcesCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResourcesCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResourcesCollection(val *UserResourcesCollection) *NullableUserResourcesCollection {
	return &NullableUserResourcesCollection{value: val, isSet: true}
}

func (v NullableUserResourcesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResourcesCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
