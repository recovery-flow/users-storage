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

// checks if the UserCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCollection{}

// UserCollection struct for UserCollection
type UserCollection struct {
	Data UserCollectionData `json:"data"`
}

type _UserCollection UserCollection

// NewUserCollection instantiates a new UserCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCollection(data UserCollectionData) *UserCollection {
	this := UserCollection{}
	this.Data = data
	return &this
}

// NewUserCollectionWithDefaults instantiates a new UserCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCollectionWithDefaults() *UserCollection {
	this := UserCollection{}
	return &this
}

// GetData returns the Data field value
func (o *UserCollection) GetData() UserCollectionData {
	if o == nil {
		var ret UserCollectionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserCollection) GetDataOk() (*UserCollectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserCollection) SetData(v UserCollectionData) {
	o.Data = v
}

func (o UserCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UserCollection) UnmarshalJSON(data []byte) (err error) {
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

	varUserCollection := _UserCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCollection)

	if err != nil {
		return err
	}

	*o = UserCollection(varUserCollection)

	return err
}

type NullableUserCollection struct {
	value *UserCollection
	isSet bool
}

func (v NullableUserCollection) Get() *UserCollection {
	return v.value
}

func (v *NullableUserCollection) Set(val *UserCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCollection(val *UserCollection) *NullableUserCollection {
	return &NullableUserCollection{value: val, isSet: true}
}

func (v NullableUserCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


