# TeamUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Team ID | 
**Type** | **string** |  | 
**Attributes** | [**TeamCreateDataAttributes**](TeamCreateDataAttributes.md) |  | 

## Methods

### NewTeamUpdateData

`func NewTeamUpdateData(id string, type_ string, attributes TeamCreateDataAttributes, ) *TeamUpdateData`

NewTeamUpdateData instantiates a new TeamUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdateDataWithDefaults

`func NewTeamUpdateDataWithDefaults() *TeamUpdateData`

NewTeamUpdateDataWithDefaults instantiates a new TeamUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TeamUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TeamUpdateData) GetAttributes() TeamCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TeamUpdateData) GetAttributesOk() (*TeamCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TeamUpdateData) SetAttributes(v TeamCreateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


