package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessRecordCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRecordCollection{}

// AccessRecordCollection A collection of access records
type AccessRecordCollection struct {
	// A list of access records
	Records []AccessRecord `json:"records"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Links CollectionLinks `json:"links"`
}

type _AccessRecordCollection AccessRecordCollection

// NewAccessRecordCollection instantiates a new AccessRecordCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRecordCollection(records []AccessRecord, links CollectionLinks) *AccessRecordCollection {
	this := AccessRecordCollection{}
	this.Records = records
	this.Links = links
	return &this
}

// NewAccessRecordCollectionWithDefaults instantiates a new AccessRecordCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRecordCollectionWithDefaults() *AccessRecordCollection {
	this := AccessRecordCollection{}
	return &this
}

// GetRecords returns the Records field value
func (o *AccessRecordCollection) GetRecords() []AccessRecord {
	if o == nil {
		var ret []AccessRecord
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *AccessRecordCollection) GetRecordsOk() ([]AccessRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *AccessRecordCollection) SetRecords(v []AccessRecord) {
	o.Records = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AccessRecordCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRecordCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AccessRecordCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AccessRecordCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *AccessRecordCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AccessRecordCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AccessRecordCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o AccessRecordCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRecordCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["records"] = o.Records
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *AccessRecordCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"records",
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

	varAccessRecordCollection := _AccessRecordCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRecordCollection)

	if err != nil {
		return err
	}

	*o = AccessRecordCollection(varAccessRecordCollection)

	return err
}

type NullableAccessRecordCollection struct {
	value *AccessRecordCollection
	isSet bool
}

func (v NullableAccessRecordCollection) Get() *AccessRecordCollection {
	return v.value
}

func (v *NullableAccessRecordCollection) Set(val *AccessRecordCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRecordCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRecordCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRecordCollection(val *AccessRecordCollection) *NullableAccessRecordCollection {
	return &NullableAccessRecordCollection{value: val, isSet: true}
}

func (v NullableAccessRecordCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRecordCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


