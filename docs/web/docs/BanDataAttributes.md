# BanDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Term** | **int32** | ban period at days | 
**Sort** | **string** | type of block | 
**Desc** | **string** | reason for ban and explanation | 

## Methods

### NewBanDataAttributes

`func NewBanDataAttributes(term int32, sort string, desc string, ) *BanDataAttributes`

NewBanDataAttributes instantiates a new BanDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanDataAttributesWithDefaults

`func NewBanDataAttributesWithDefaults() *BanDataAttributes`

NewBanDataAttributesWithDefaults instantiates a new BanDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerm

`func (o *BanDataAttributes) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *BanDataAttributes) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *BanDataAttributes) SetTerm(v int32)`

SetTerm sets Term field to given value.


### GetSort

`func (o *BanDataAttributes) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BanDataAttributes) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BanDataAttributes) SetSort(v string)`

SetSort sets Sort field to given value.


### GetDesc

`func (o *BanDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *BanDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *BanDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


