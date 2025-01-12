# TeamsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Team ID | 
**Type** | **string** |  | 
**Attributes** | [**TeamsDataAttributes**](TeamsDataAttributes.md) |  | 

## Methods

### NewTeamsData

`func NewTeamsData(id string, type_ string, attributes TeamsDataAttributes, ) *TeamsData`

NewTeamsData instantiates a new TeamsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsDataWithDefaults

`func NewTeamsDataWithDefaults() *TeamsData`

NewTeamsDataWithDefaults instantiates a new TeamsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamsData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TeamsData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamsData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamsData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TeamsData) GetAttributes() TeamsDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamsData) GetAttributesOk() (*TeamsDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamsData) SetAttributes(v TeamsDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


