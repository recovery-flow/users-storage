# BanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | user id | 
**Type** | **string** |  | 
**Attributes** | [**BanDataAttributes**](BanDataAttributes.md) |  | 

## Methods

### NewBanData

`func NewBanData(id string, type_ string, attributes BanDataAttributes, ) *BanData`

NewBanData instantiates a new BanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanDataWithDefaults

`func NewBanDataWithDefaults() *BanData`

NewBanDataWithDefaults instantiates a new BanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BanData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BanData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BanData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BanData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BanData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BanData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BanData) GetAttributes() BanDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BanData) GetAttributesOk() (*BanDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BanData) SetAttributes(v BanDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


