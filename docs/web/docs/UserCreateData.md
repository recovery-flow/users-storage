# UserCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**UserCreateDataAttributes**](UserCreateDataAttributes.md) |  | 

## Methods

### NewUserCreateData

`func NewUserCreateData(type_ string, attributes UserCreateDataAttributes, ) *UserCreateData`

NewUserCreateData instantiates a new UserCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateDataWithDefaults

`func NewUserCreateDataWithDefaults() *UserCreateData`

NewUserCreateDataWithDefaults instantiates a new UserCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UserCreateData) GetAttributes() UserCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserCreateData) GetAttributesOk() (*UserCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserCreateData) SetAttributes(v UserCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


