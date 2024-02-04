package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the AccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequest{}

// AccessRequest The access requested by a user.
type AccessRequest struct {
	// Unique identifier for the request.
	RequestId string `json:"requestId"`
	// The expected last time the request was updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Current status of the access request.
	Status *string        `json:"status,omitempty"`
	Access AccessTemplate `json:"access"`
	Links  *AccountLinks  `json:"links,omitempty"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

type _AccessRequest AccessRequest

// NewAccessRequest instantiates a new AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequest(requestId string, access AccessTemplate) *AccessRequest {
	this := AccessRequest{}
	this.RequestId = requestId
	this.Access = access
	return &this
}

// NewAccessRequestWithDefaults instantiates a new AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestWithDefaults() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *AccessRequest) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AccessRequest) SetRequestId(v string) {
	o.RequestId = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AccessRequest) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AccessRequest) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AccessRequest) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccessRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccessRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccessRequest) SetStatus(v string) {
	o.Status = &v
}

// GetAccess returns the Access field value
func (o *AccessRequest) GetAccess() AccessTemplate {
	if o == nil {
		var ret AccessTemplate
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetAccessOk() (*AccessTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *AccessRequest) SetAccess(v AccessTemplate) {
	o.Access = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccessRequest) GetLinks() AccountLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccessRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *AccessRequest) SetLinks(v AccountLinks) {
	o.Links = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessRequest) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccessRequest) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *AccessRequest) SetTags(v map[string]string) {
	o.Tags = v
}

func (o AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestId"] = o.RequestId
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["access"] = o.Access
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *AccessRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestId",
		"access",
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

	varAccessRequest := _AccessRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRequest)

	if err != nil {
		return err
	}

	*o = AccessRequest(varAccessRequest)

	return err
}

type NullableAccessRequest struct {
	value *AccessRequest
	isSet bool
}

func (v NullableAccessRequest) Get() *AccessRequest {
	return v.value
}

func (v *NullableAccessRequest) Set(val *AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequest(val *AccessRequest) *NullableAccessRequest {
	return &NullableAccessRequest{value: val, isSet: true}
}

func (v NullableAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
