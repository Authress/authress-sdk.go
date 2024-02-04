package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the AccessRequestCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestCollection{}

// AccessRequestCollection A collection of access requests
type AccessRequestCollection struct {
	// A list of access requests
	Records    []AccessRequest `json:"records,omitempty"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Links      CollectionLinks `json:"links"`
}

type _AccessRequestCollection AccessRequestCollection

// NewAccessRequestCollection instantiates a new AccessRequestCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestCollection(links CollectionLinks) *AccessRequestCollection {
	this := AccessRequestCollection{}
	this.Links = links
	return &this
}

// NewAccessRequestCollectionWithDefaults instantiates a new AccessRequestCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestCollectionWithDefaults() *AccessRequestCollection {
	this := AccessRequestCollection{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *AccessRequestCollection) GetRecords() []AccessRequest {
	if o == nil || IsNil(o.Records) {
		var ret []AccessRequest
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestCollection) GetRecordsOk() ([]AccessRequest, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *AccessRequestCollection) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []AccessRequest and assigns it to the Records field.
func (o *AccessRequestCollection) SetRecords(v []AccessRequest) {
	o.Records = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AccessRequestCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequestCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AccessRequestCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AccessRequestCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *AccessRequestCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AccessRequestCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AccessRequestCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o AccessRequestCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AccessRequestCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAccessRequestCollection := _AccessRequestCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRequestCollection)

	if err != nil {
		return err
	}

	*o = AccessRequestCollection(varAccessRequestCollection)

	return err
}

type NullableAccessRequestCollection struct {
	value *AccessRequestCollection
	isSet bool
}

func (v NullableAccessRequestCollection) Get() *AccessRequestCollection {
	return v.value
}

func (v *NullableAccessRequestCollection) Set(val *AccessRequestCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestCollection(val *AccessRequestCollection) *NullableAccessRequestCollection {
	return &NullableAccessRequestCollection{value: val, isSet: true}
}

func (v NullableAccessRequestCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
