package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the CollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionLinks{}

// CollectionLinks struct for CollectionLinks
type CollectionLinks struct {
	Self NullableLink `json:"self"`
	Next NullableLink `json:"next,omitempty"`
}

type _CollectionLinks CollectionLinks

// NewCollectionLinks instantiates a new CollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionLinks(self NullableLink) *CollectionLinks {
	this := CollectionLinks{}
	this.Self = self
	return &this
}

// NewCollectionLinksWithDefaults instantiates a new CollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionLinksWithDefaults() *CollectionLinks {
	this := CollectionLinks{}
	return &this
}

// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for Link will be returned
func (o *CollectionLinks) GetSelf() Link {
	if o == nil || o.Self.Get() == nil {
		var ret Link
		return ret
	}

	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// SetSelf sets field value
func (o *CollectionLinks) SetSelf(v Link) {
	o.Self.Set(&v)
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionLinks) GetNext() Link {
	if o == nil || IsNil(o.Next.Get()) {
		var ret Link
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionLinks) GetNextOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *CollectionLinks) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableLink and assigns it to the Next field.
func (o *CollectionLinks) SetNext(v Link) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *CollectionLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *CollectionLinks) UnsetNext() {
	o.Next.Unset()
}

func (o CollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self.Get()
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	return toSerialize, nil
}

func (o *CollectionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varCollectionLinks := _CollectionLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionLinks)

	if err != nil {
		return err
	}

	*o = CollectionLinks(varCollectionLinks)

	return err
}

type NullableCollectionLinks struct {
	value *CollectionLinks
	isSet bool
}

func (v NullableCollectionLinks) Get() *CollectionLinks {
	return v.value
}

func (v *NullableCollectionLinks) Set(val *CollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionLinks(val *CollectionLinks) *NullableCollectionLinks {
	return &NullableCollectionLinks{value: val, isSet: true}
}

func (v NullableCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
