package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the PaginationNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationNext{}

// PaginationNext A reference to the next page in the collection, pass the cursor as a query parameter in the subsequent request to get the next page.
type PaginationNext struct {
	Cursor string `json:"cursor"`
}

type _PaginationNext PaginationNext

// NewPaginationNext instantiates a new PaginationNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationNext(cursor string) *PaginationNext {
	this := PaginationNext{}
	this.Cursor = cursor
	return &this
}

// NewPaginationNextWithDefaults instantiates a new PaginationNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationNextWithDefaults() *PaginationNext {
	this := PaginationNext{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *PaginationNext) GetCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *PaginationNext) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *PaginationNext) SetCursor(v string) {
	o.Cursor = v
}

func (o PaginationNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

func (o *PaginationNext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cursor",
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

	varPaginationNext := _PaginationNext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPaginationNext)

	if err != nil {
		return err
	}

	*o = PaginationNext(varPaginationNext)

	return err
}

type NullablePaginationNext struct {
	value *PaginationNext
	isSet bool
}

func (v NullablePaginationNext) Get() *PaginationNext {
	return v.value
}

func (v *NullablePaginationNext) Set(val *PaginationNext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationNext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationNext(val *PaginationNext) *NullablePaginationNext {
	return &NullablePaginationNext{value: val, isSet: true}
}

func (v NullablePaginationNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
