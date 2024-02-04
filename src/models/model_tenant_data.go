package authress

import (
	"encoding/json"
)

// checks if the TenantData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantData{}

// TenantData struct for TenantData
type TenantData struct {
	Name NullableString `json:"name,omitempty"`
}

// NewTenantData instantiates a new TenantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantData() *TenantData {
	this := TenantData{}
	return &this
}

// NewTenantDataWithDefaults instantiates a new TenantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantDataWithDefaults() *TenantData {
	this := TenantData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TenantData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TenantData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TenantData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TenantData) UnsetName() {
	o.Name.Unset()
}

func (o TenantData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableTenantData struct {
	value *TenantData
	isSet bool
}

func (v NullableTenantData) Get() *TenantData {
	return v.value
}

func (v *NullableTenantData) Set(val *TenantData) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantData) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantData(val *TenantData) *NullableTenantData {
	return &NullableTenantData{value: val, isSet: true}
}

func (v NullableTenantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


