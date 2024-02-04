package authress

import (
	"encoding/json"
)

// checks if the Pagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pagination{}

// Pagination Details containing how to pagination through the collection. Consists of an optional *next* property that may contain a cursor. Pagination is mutable and the list can change between requests.
type Pagination struct {
	Next NullablePaginationNext `json:"next,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pagination) GetNext() PaginationNext {
	if o == nil || IsNil(o.Next.Get()) {
		var ret PaginationNext
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pagination) GetNextOk() (*PaginationNext, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *Pagination) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullablePaginationNext and assigns it to the Next field.
func (o *Pagination) SetNext(v PaginationNext) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *Pagination) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *Pagination) UnsetNext() {
	o.Next.Unset()
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	return toSerialize, nil
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


