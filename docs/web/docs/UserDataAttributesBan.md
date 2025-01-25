# UserDataAttributesBan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **bool** | Ban status | 
**Start** | **time.Time** | Ban start date | 
**End** | **time.Time** | Ban end date | 
**Sort** | **string** | Type of block | 
**Desc** | **string** | Reason for ban and explanation | 
**Initiator** | **string** | Ban initiator id | 

## Methods

### NewUserDataAttributesBan

`func NewUserDataAttributesBan(status bool, start time.Time, end time.Time, sort string, desc string, initiator string, ) *UserDataAttributesBan`

NewUserDataAttributesBan instantiates a new UserDataAttributesBan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataAttributesBanWithDefaults

`func NewUserDataAttributesBanWithDefaults() *UserDataAttributesBan`

NewUserDataAttributesBanWithDefaults instantiates a new UserDataAttributesBan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserDataAttributesBan) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDataAttributesBan) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDataAttributesBan) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetStart

`func (o *UserDataAttributesBan) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UserDataAttributesBan) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UserDataAttributesBan) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *UserDataAttributesBan) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UserDataAttributesBan) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UserDataAttributesBan) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetSort

`func (o *UserDataAttributesBan) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UserDataAttributesBan) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UserDataAttributesBan) SetSort(v string)`

SetSort sets Sort field to given value.


### GetDesc

`func (o *UserDataAttributesBan) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *UserDataAttributesBan) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *UserDataAttributesBan) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetInitiator

`func (o *UserDataAttributesBan) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *UserDataAttributesBan) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *UserDataAttributesBan) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


