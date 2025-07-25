package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the TenantCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantCollection{}

// TenantCollection A collection of tenants.
type TenantCollection struct {
	Tenants    []Tenant    `json:"tenants"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _TenantCollection TenantCollection

// NewTenantCollection instantiates a new TenantCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCollection(tenants []Tenant) *TenantCollection {
	this := TenantCollection{}
	this.Tenants = tenants
	return &this
}

// NewTenantCollectionWithDefaults instantiates a new TenantCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCollectionWithDefaults() *TenantCollection {
	this := TenantCollection{}
	return &this
}

// GetTenants returns the Tenants field value
func (o *TenantCollection) GetTenants() []Tenant {
	if o == nil {
		var ret []Tenant
		return ret
	}

	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value
// and a boolean to check if the value has been set.
func (o *TenantCollection) GetTenantsOk() ([]Tenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenants, true
}

// SetTenants sets field value
func (o *TenantCollection) SetTenants(v []Tenant) {
	o.Tenants = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TenantCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TenantCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *TenantCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o TenantCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenants"] = o.Tenants
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

func (o *TenantCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenants",
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

	varTenantCollection := _TenantCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTenantCollection)

	if err != nil {
		return err
	}

	*o = TenantCollection(varTenantCollection)

	return err
}

type NullableTenantCollection struct {
	value *TenantCollection
	isSet bool
}

func (v NullableTenantCollection) Get() *TenantCollection {
	return v.value
}

func (v *NullableTenantCollection) Set(val *TenantCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCollection(val *TenantCollection) *NullableTenantCollection {
	return &NullableTenantCollection{value: val, isSet: true}
}

func (v NullableTenantCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
