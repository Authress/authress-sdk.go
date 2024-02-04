package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExtensionClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionClient{}

// ExtensionClient The extension's client configuration.
type ExtensionClient struct {
	// The unique ID of the client.
	ClientId string `json:"clientId"`
	Links *Links `json:"links,omitempty"`
}

type _ExtensionClient ExtensionClient

// NewExtensionClient instantiates a new ExtensionClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionClient(clientId string) *ExtensionClient {
	this := ExtensionClient{}
	this.ClientId = clientId
	return &this
}

// NewExtensionClientWithDefaults instantiates a new ExtensionClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionClientWithDefaults() *ExtensionClient {
	this := ExtensionClient{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ExtensionClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ExtensionClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ExtensionClient) SetClientId(v string) {
	o.ClientId = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExtensionClient) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionClient) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExtensionClient) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ExtensionClient) SetLinks(v Links) {
	o.Links = &v
}

func (o ExtensionClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *ExtensionClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
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

	varExtensionClient := _ExtensionClient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtensionClient)

	if err != nil {
		return err
	}

	*o = ExtensionClient(varExtensionClient)

	return err
}

type NullableExtensionClient struct {
	value *ExtensionClient
	isSet bool
}

func (v NullableExtensionClient) Get() *ExtensionClient {
	return v.value
}

func (v *NullableExtensionClient) Set(val *ExtensionClient) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionClient) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionClient(val *ExtensionClient) *NullableExtensionClient {
	return &NullableExtensionClient{value: val, isSet: true}
}

func (v NullableExtensionClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


