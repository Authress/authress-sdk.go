package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PermissionedResourceCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionedResourceCollection{}

// PermissionedResourceCollection A collection of resource permissions that have been defined.
type PermissionedResourceCollection struct {
	Resources []PermissionedResource `json:"resources"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Links CollectionLinks `json:"links"`
}

type _PermissionedResourceCollection PermissionedResourceCollection

// NewPermissionedResourceCollection instantiates a new PermissionedResourceCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionedResourceCollection(resources []PermissionedResource, links CollectionLinks) *PermissionedResourceCollection {
	this := PermissionedResourceCollection{}
	this.Resources = resources
	this.Links = links
	return &this
}

// NewPermissionedResourceCollectionWithDefaults instantiates a new PermissionedResourceCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionedResourceCollectionWithDefaults() *PermissionedResourceCollection {
	this := PermissionedResourceCollection{}
	return &this
}

// GetResources returns the Resources field value
func (o *PermissionedResourceCollection) GetResources() []PermissionedResource {
	if o == nil {
		var ret []PermissionedResource
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *PermissionedResourceCollection) GetResourcesOk() ([]PermissionedResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *PermissionedResourceCollection) SetResources(v []PermissionedResource) {
	o.Resources = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PermissionedResourceCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionedResourceCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PermissionedResourceCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *PermissionedResourceCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLinks returns the Links field value
func (o *PermissionedResourceCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PermissionedResourceCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PermissionedResourceCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o PermissionedResourceCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionedResourceCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *PermissionedResourceCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resources",
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

	varPermissionedResourceCollection := _PermissionedResourceCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionedResourceCollection)

	if err != nil {
		return err
	}

	*o = PermissionedResourceCollection(varPermissionedResourceCollection)

	return err
}

type NullablePermissionedResourceCollection struct {
	value *PermissionedResourceCollection
	isSet bool
}

func (v NullablePermissionedResourceCollection) Get() *PermissionedResourceCollection {
	return v.value
}

func (v *NullablePermissionedResourceCollection) Set(val *PermissionedResourceCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionedResourceCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionedResourceCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionedResourceCollection(val *PermissionedResourceCollection) *NullablePermissionedResourceCollection {
	return &NullablePermissionedResourceCollection{value: val, isSet: true}
}

func (v NullablePermissionedResourceCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionedResourceCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


