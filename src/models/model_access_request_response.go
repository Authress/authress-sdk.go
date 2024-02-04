package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRequestResponse{}

// AccessRequestResponse A dynamic body to support request PATCH operations
type AccessRequestResponse struct {
	// New result, either approve or deny the request
	Status string `json:"status"`
}

type _AccessRequestResponse AccessRequestResponse

// NewAccessRequestResponse instantiates a new AccessRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequestResponse(status string) *AccessRequestResponse {
	this := AccessRequestResponse{}
	this.Status = status
	return &this
}

// NewAccessRequestResponseWithDefaults instantiates a new AccessRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestResponseWithDefaults() *AccessRequestResponse {
	this := AccessRequestResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *AccessRequestResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccessRequestResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccessRequestResponse) SetStatus(v string) {
	o.Status = v
}

func (o AccessRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *AccessRequestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varAccessRequestResponse := _AccessRequestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRequestResponse)

	if err != nil {
		return err
	}

	*o = AccessRequestResponse(varAccessRequestResponse)

	return err
}

type NullableAccessRequestResponse struct {
	value *AccessRequestResponse
	isSet bool
}

func (v NullableAccessRequestResponse) Get() *AccessRequestResponse {
	return v.value
}

func (v *NullableAccessRequestResponse) Set(val *AccessRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequestResponse(val *AccessRequestResponse) *NullableAccessRequestResponse {
	return &NullableAccessRequestResponse{value: val, isSet: true}
}

func (v NullableAccessRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


