# TeamsUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Team ID | 
**Type** | **string** |  | 
**Attributes** | [**TeamsCreateDataAttributes**](TeamsCreateDataAttributes.md) |  | 

## Methods

### NewTeamsUpdateData

`func NewTeamsUpdateData(id string, type_ string, attributes TeamsCreateDataAttributes, ) *TeamsUpdateData`

NewTeamsUpdateData instantiates a new TeamsUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsUpdateDataWithDefaults

`func NewTeamsUpdateDataWithDefaults() *TeamsUpdateData`

NewTeamsUpdateDataWithDefaults instantiates a new TeamsUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamsUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamsUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamsUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TeamsUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamsUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamsUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TeamsUpdateData) GetAttributes() TeamsCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamsUpdateData) GetAttributesOk() (*TeamsCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamsUpdateData) SetAttributes(v TeamsCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


