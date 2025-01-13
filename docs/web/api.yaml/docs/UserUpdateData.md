# UserUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Type** | **string** |  | 
**Attributes** | [**UserUpdateDataAttributes**](UserUpdateDataAttributes.md) |  | 

## Methods

### NewUserUpdateData

`func NewUserUpdateData(id string, type_ string, attributes UserUpdateDataAttributes, ) *UserUpdateData`

NewUserUpdateData instantiates a new UserUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateDataWithDefaults

`func NewUserUpdateDataWithDefaults() *UserUpdateData`

NewUserUpdateDataWithDefaults instantiates a new UserUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UserUpdateData) GetAttributes() UserUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserUpdateData) GetAttributesOk() (*UserUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserUpdateData) SetAttributes(v UserUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


