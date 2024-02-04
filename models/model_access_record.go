package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the AccessRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRecord{}

// AccessRecord The access record which links users to roles.
type AccessRecord struct {
	// Unique identifier for the record, can be specified on record creation.
	RecordId *string `json:"recordId,omitempty"`
	// A helpful name for this record
	Name string `json:"name"`
	// More details about this record
	Description NullableString `json:"description,omitempty"`
	// Percentage capacity of record that is filled.
	Capacity *float32 `json:"capacity,omitempty"`
	// The expected last time the record was updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Current status of the access record.
	Status  *string              `json:"status,omitempty"`
	Account *AccessRecordAccount `json:"account,omitempty"`
	// The list of users this record applies to
	Users []User `json:"users,omitempty"`
	// The list of admin that can edit this record even if they do not have global record edit permissions.
	Admins []User `json:"admins,omitempty"`
	// The list of groups this record applies to. Users in these groups will be receive access to the resources listed.
	Groups []LinkedGroup `json:"groups,omitempty"`
	// A list of statements which match roles to resources.
	Statements []Statement   `json:"statements"`
	Links      *AccountLinks `json:"links,omitempty"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

type _AccessRecord AccessRecord

// NewAccessRecord instantiates a new AccessRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRecord(name string, statements []Statement) *AccessRecord {
	this := AccessRecord{}
	this.Name = name
	this.Statements = statements
	return &this
}

// NewAccessRecordWithDefaults instantiates a new AccessRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRecordWithDefaults() *AccessRecord {
	this := AccessRecord{}
	return &this
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *AccessRecord) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *AccessRecord) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *AccessRecord) SetRecordId(v string) {
	o.RecordId = &v
}

// GetName returns the Name field value
func (o *AccessRecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessRecord) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRecord) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRecord) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccessRecord) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccessRecord) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccessRecord) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccessRecord) UnsetDescription() {
	o.Description.Unset()
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *AccessRecord) GetCapacity() float32 {
	if o == nil || IsNil(o.Capacity) {
		var ret float32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetCapacityOk() (*float32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *AccessRecord) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given float32 and assigns it to the Capacity field.
func (o *AccessRecord) SetCapacity(v float32) {
	o.Capacity = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AccessRecord) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AccessRecord) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AccessRecord) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccessRecord) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccessRecord) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccessRecord) SetStatus(v string) {
	o.Status = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AccessRecord) GetAccount() AccessRecordAccount {
	if o == nil || IsNil(o.Account) {
		var ret AccessRecordAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetAccountOk() (*AccessRecordAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AccessRecord) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccessRecordAccount and assigns it to the Account field.
func (o *AccessRecord) SetAccount(v AccessRecordAccount) {
	o.Account = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRecord) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRecord) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AccessRecord) HasUsers() bool {
	if o != nil && IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *AccessRecord) SetUsers(v []User) {
	o.Users = v
}

// GetAdmins returns the Admins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRecord) GetAdmins() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRecord) GetAdminsOk() ([]User, bool) {
	if o == nil || IsNil(o.Admins) {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *AccessRecord) HasAdmins() bool {
	if o != nil && IsNil(o.Admins) {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given []User and assigns it to the Admins field.
func (o *AccessRecord) SetAdmins(v []User) {
	o.Admins = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRecord) GetGroups() []LinkedGroup {
	if o == nil {
		var ret []LinkedGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRecord) GetGroupsOk() ([]LinkedGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AccessRecord) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []LinkedGroup and assigns it to the Groups field.
func (o *AccessRecord) SetGroups(v []LinkedGroup) {
	o.Groups = v
}

// GetStatements returns the Statements field value
func (o *AccessRecord) GetStatements() []Statement {
	if o == nil {
		var ret []Statement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetStatementsOk() ([]Statement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *AccessRecord) SetStatements(v []Statement) {
	o.Statements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccessRecord) GetLinks() AccountLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecord) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccessRecord) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *AccessRecord) SetLinks(v AccountLinks) {
	o.Links = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRecord) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRecord) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccessRecord) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AccessRecord) SetTags(v map[string]string) {
	o.Tags = v
}

func (o AccessRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordId) {
		toSerialize["recordId"] = o.RecordId
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Admins != nil {
		toSerialize["admins"] = o.Admins
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["statements"] = o.Statements
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *AccessRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"statements",
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

	varAccessRecord := _AccessRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRecord)

	if err != nil {
		return err
	}

	*o = AccessRecord(varAccessRecord)

	return err
}

type NullableAccessRecord struct {
	value *AccessRecord
	isSet bool
}

func (v NullableAccessRecord) Get() *AccessRecord {
	return v.value
}

func (v *NullableAccessRecord) Set(val *AccessRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRecord(val *AccessRecord) *NullableAccessRecord {
	return &NullableAccessRecord{value: val, isSet: true}
}

func (v NullableAccessRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
