package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	// A resource path which can be top level, fully qualified, or end with a *. Parent resources imply permissions to sub-resources.
	ResourceUri string `json:"resourceUri"`
}

type _Resource Resource

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(resourceUri string) *Resource {
	this := Resource{}
	this.ResourceUri = resourceUri
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *Resource) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *Resource) SetResourceUri(v string) {
	o.ResourceUri = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceUri"] = o.ResourceUri
	return toSerialize, nil
}

func (o *Resource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceUri",
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

	varResource := _Resource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResource)

	if err != nil {
		return err
	}

	*o = Resource(varResource)

	return err
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


