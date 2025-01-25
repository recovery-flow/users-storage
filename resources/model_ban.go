/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Ban type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ban{}

// Ban struct for Ban
type Ban struct {
	Data BanData `json:"data"`
}

type _Ban Ban

// NewBan instantiates a new Ban object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBan(data BanData) *Ban {
	this := Ban{}
	this.Data = data
	return &this
}

// NewBanWithDefaults instantiates a new Ban object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanWithDefaults() *Ban {
	this := Ban{}
	return &this
}

// GetData returns the Data field value
func (o *Ban) GetData() BanData {
	if o == nil {
		var ret BanData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Ban) GetDataOk() (*BanData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Ban) SetData(v BanData) {
	o.Data = v
}

func (o Ban) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ban) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *Ban) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varBan := _Ban{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBan)

	if err != nil {
		return err
	}

	*o = Ban(varBan)

	return err
}

type NullableBan struct {
	value *Ban
	isSet bool
}

func (v NullableBan) Get() *Ban {
	return v.value
}

func (v *NullableBan) Set(val *Ban) {
	v.value = val
	v.isSet = true
}

func (v NullableBan) IsSet() bool {
	return v.isSet
}

func (v *NullableBan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBan(val *Ban) *NullableBan {
	return &NullableBan{value: val, isSet: true}
}

func (v NullableBan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


