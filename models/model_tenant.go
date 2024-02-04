package models

import (
	"encoding/json"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant struct for Tenant
type Tenant struct {
	TenantId               NullableString           `json:"tenantId,omitempty"`
	TenantLookupIdentifier NullableString           `json:"tenantLookupIdentifier,omitempty"`
	Data                   *TenantData              `json:"data,omitempty"`
	Connection             NullableTenantConnection `json:"connection,omitempty"`
	CreatedTime            *time.Time               `json:"createdTime,omitempty"`
}

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant() *Tenant {
	this := Tenant{}
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetTenantId() string {
	if o == nil || IsNil(o.TenantId.Get()) {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *Tenant) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *Tenant) SetTenantId(v string) {
	o.TenantId.Set(&v)
}

// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *Tenant) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *Tenant) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetTenantLookupIdentifier returns the TenantLookupIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetTenantLookupIdentifier() string {
	if o == nil || IsNil(o.TenantLookupIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.TenantLookupIdentifier.Get()
}

// GetTenantLookupIdentifierOk returns a tuple with the TenantLookupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetTenantLookupIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantLookupIdentifier.Get(), o.TenantLookupIdentifier.IsSet()
}

// HasTenantLookupIdentifier returns a boolean if a field has been set.
func (o *Tenant) HasTenantLookupIdentifier() bool {
	if o != nil && o.TenantLookupIdentifier.IsSet() {
		return true
	}

	return false
}

// SetTenantLookupIdentifier gets a reference to the given NullableString and assigns it to the TenantLookupIdentifier field.
func (o *Tenant) SetTenantLookupIdentifier(v string) {
	o.TenantLookupIdentifier.Set(&v)
}

// SetTenantLookupIdentifierNil sets the value for TenantLookupIdentifier to be an explicit nil
func (o *Tenant) SetTenantLookupIdentifierNil() {
	o.TenantLookupIdentifier.Set(nil)
}

// UnsetTenantLookupIdentifier ensures that no value is present for TenantLookupIdentifier, not even an explicit nil
func (o *Tenant) UnsetTenantLookupIdentifier() {
	o.TenantLookupIdentifier.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Tenant) GetData() TenantData {
	if o == nil || IsNil(o.Data) {
		var ret TenantData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDataOk() (*TenantData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Tenant) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TenantData and assigns it to the Data field.
func (o *Tenant) SetData(v TenantData) {
	o.Data = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetConnection() TenantConnection {
	if o == nil || IsNil(o.Connection.Get()) {
		var ret TenantConnection
		return ret
	}
	return *o.Connection.Get()
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetConnectionOk() (*TenantConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connection.Get(), o.Connection.IsSet()
}

// HasConnection returns a boolean if a field has been set.
func (o *Tenant) HasConnection() bool {
	if o != nil && o.Connection.IsSet() {
		return true
	}

	return false
}

// SetConnection gets a reference to the given NullableTenantConnection and assigns it to the Connection field.
func (o *Tenant) SetConnection(v TenantConnection) {
	o.Connection.Set(&v)
}

// SetConnectionNil sets the value for Connection to be an explicit nil
func (o *Tenant) SetConnectionNil() {
	o.Connection.Set(nil)
}

// UnsetConnection ensures that no value is present for Connection, not even an explicit nil
func (o *Tenant) UnsetConnection() {
	o.Connection.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Tenant) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Tenant) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *Tenant) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.TenantLookupIdentifier.IsSet() {
		toSerialize["tenantLookupIdentifier"] = o.TenantLookupIdentifier.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Connection.IsSet() {
		toSerialize["connection"] = o.Connection.Get()
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	return toSerialize, nil
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
