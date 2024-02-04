package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Extension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Extension{}

// Extension struct for Extension
type Extension struct {
	ExtensionId string `json:"extensionId"`
	// The name of the extension. This name is visible in the Authress management portal
	Name        NullableString        `json:"name,omitempty"`
	CreatedTime time.Time             `json:"createdTime"`
	Application *ExtensionApplication `json:"application,omitempty"`
	Client      ExtensionClient       `json:"client"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

type _Extension Extension

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension(extensionId string, createdTime time.Time, client ExtensionClient) *Extension {
	this := Extension{}
	this.ExtensionId = extensionId
	this.CreatedTime = createdTime
	this.Client = client
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetExtensionId returns the ExtensionId field value
func (o *Extension) GetExtensionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionId
}

// GetExtensionIdOk returns a tuple with the ExtensionId field value
// and a boolean to check if the value has been set.
func (o *Extension) GetExtensionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionId, true
}

// SetExtensionId sets field value
func (o *Extension) SetExtensionId(v string) {
	o.ExtensionId = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Extension) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Extension) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Extension) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Extension) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Extension) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Extension) UnsetName() {
	o.Name.Unset()
}

// GetCreatedTime returns the CreatedTime field value
func (o *Extension) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Extension) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Extension) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *Extension) GetApplication() ExtensionApplication {
	if o == nil || IsNil(o.Application) {
		var ret ExtensionApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetApplicationOk() (*ExtensionApplication, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *Extension) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ExtensionApplication and assigns it to the Application field.
func (o *Extension) SetApplication(v ExtensionApplication) {
	o.Application = &v
}

// GetClient returns the Client field value
func (o *Extension) GetClient() ExtensionClient {
	if o == nil {
		var ret ExtensionClient
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *Extension) GetClientOk() (*ExtensionClient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *Extension) SetClient(v ExtensionClient) {
	o.Client = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Extension) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Extension) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Extension) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Extension) SetTags(v map[string]string) {
	o.Tags = v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Extension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extensionId"] = o.ExtensionId
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["createdTime"] = o.CreatedTime
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	toSerialize["client"] = o.Client
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *Extension) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"extensionId",
		"createdTime",
		"client",
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

	varExtension := _Extension{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtension)

	if err != nil {
		return err
	}

	*o = Extension(varExtension)

	return err
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
