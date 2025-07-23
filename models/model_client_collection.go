package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ClientCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientCollection{}

// ClientCollection The collection of a list of clients
type ClientCollection struct {
	// A list of clients
	Clients    []Client        `json:"clients"`
	Pagination *Pagination     `json:"pagination,omitempty"`
	Links      CollectionLinks `json:"links"`
}

type _ClientCollection ClientCollection

// NewClientCollection instantiates a new ClientCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCollection(clients []Client, links CollectionLinks) *ClientCollection {
	this := ClientCollection{}
	this.Clients = clients
	this.Links = links
	return &this
}

// NewClientCollectionWithDefaults instantiates a new ClientCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientCollectionWithDefaults() *ClientCollection {
	this := ClientCollection{}
	return &this
}

// GetClients returns the Clients field value
func (o *ClientCollection) GetClients() []Client {
	if o == nil {
		var ret []Client
		return ret
	}

	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value
// and a boolean to check if the value has been set.
func (o *ClientCollection) GetClientsOk() ([]Client, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clients, true
}

// SetClients sets field value
func (o *ClientCollection) SetClients(v []Client) {
	o.Clients = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ClientCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ClientCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ClientCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *ClientCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ClientCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ClientCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o ClientCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clients"] = o.Clients
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *ClientCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clients",
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

	varClientCollection := _ClientCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varClientCollection)

	if err != nil {
		return err
	}

	*o = ClientCollection(varClientCollection)

	return err
}

type NullableClientCollection struct {
	value *ClientCollection
	isSet bool
}

func (v NullableClientCollection) Get() *ClientCollection {
	return v.value
}

func (v *NullableClientCollection) Set(val *ClientCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCollection(val *ClientCollection) *NullableClientCollection {
	return &NullableClientCollection{value: val, isSet: true}
}

func (v NullableClientCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
