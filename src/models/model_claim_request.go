package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ClaimRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimRequest{}

// ClaimRequest struct for ClaimRequest
type ClaimRequest struct {
	// The parent resource to add a sub-resource to. The resource must have a resource configuration that add the permission CLAIM for this to work.
	CollectionResource string `json:"collectionResource"`
	// The sub-resource the user is requesting Admin ownership over.
	ResourceId string `json:"resourceId"`
}

type _ClaimRequest ClaimRequest

// NewClaimRequest instantiates a new ClaimRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimRequest(collectionResource string, resourceId string) *ClaimRequest {
	this := ClaimRequest{}
	this.CollectionResource = collectionResource
	this.ResourceId = resourceId
	return &this
}

// NewClaimRequestWithDefaults instantiates a new ClaimRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimRequestWithDefaults() *ClaimRequest {
	this := ClaimRequest{}
	return &this
}

// GetCollectionResource returns the CollectionResource field value
func (o *ClaimRequest) GetCollectionResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionResource
}

// GetCollectionResourceOk returns a tuple with the CollectionResource field value
// and a boolean to check if the value has been set.
func (o *ClaimRequest) GetCollectionResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionResource, true
}

// SetCollectionResource sets field value
func (o *ClaimRequest) SetCollectionResource(v string) {
	o.CollectionResource = v
}

// GetResourceId returns the ResourceId field value
func (o *ClaimRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ClaimRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ClaimRequest) SetResourceId(v string) {
	o.ResourceId = v
}

func (o ClaimRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collectionResource"] = o.CollectionResource
	toSerialize["resourceId"] = o.ResourceId
	return toSerialize, nil
}

func (o *ClaimRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"collectionResource",
		"resourceId",
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

	varClaimRequest := _ClaimRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClaimRequest)

	if err != nil {
		return err
	}

	*o = ClaimRequest(varClaimRequest)

	return err
}

type NullableClaimRequest struct {
	value *ClaimRequest
	isSet bool
}

func (v NullableClaimRequest) Get() *ClaimRequest {
	return v.value
}

func (v *NullableClaimRequest) Set(val *ClaimRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimRequest(val *ClaimRequest) *NullableClaimRequest {
	return &NullableClaimRequest{value: val, isSet: true}
}

func (v NullableClaimRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


