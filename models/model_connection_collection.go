package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ConnectionCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionCollection{}

// ConnectionCollection A collection of connections.
type ConnectionCollection struct {
	Connections []Connection `json:"connections"`
	Pagination  *Pagination  `json:"pagination,omitempty"`
}

type _ConnectionCollection ConnectionCollection

// NewConnectionCollection instantiates a new ConnectionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionCollection(connections []Connection) *ConnectionCollection {
	this := ConnectionCollection{}
	this.Connections = connections
	return &this
}

// NewConnectionCollectionWithDefaults instantiates a new ConnectionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionCollectionWithDefaults() *ConnectionCollection {
	this := ConnectionCollection{}
	return &this
}

// GetConnections returns the Connections field value
func (o *ConnectionCollection) GetConnections() []Connection {
	if o == nil {
		var ret []Connection
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *ConnectionCollection) GetConnectionsOk() ([]Connection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connections, true
}

// SetConnections sets field value
func (o *ConnectionCollection) SetConnections(v []Connection) {
	o.Connections = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ConnectionCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ConnectionCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ConnectionCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ConnectionCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

func (o *ConnectionCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connections",
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

	varConnectionCollection := _ConnectionCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectionCollection)

	if err != nil {
		return err
	}

	*o = ConnectionCollection(varConnectionCollection)

	return err
}

type NullableConnectionCollection struct {
	value *ConnectionCollection
	isSet bool
}

func (v NullableConnectionCollection) Get() *ConnectionCollection {
	return v.value
}

func (v *NullableConnectionCollection) Set(val *ConnectionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCollection(val *ConnectionCollection) *NullableConnectionCollection {
	return &NullableConnectionCollection{value: val, isSet: true}
}

func (v NullableConnectionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
