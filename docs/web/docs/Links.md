# Links

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | Pointer to **string** | Link to the previous page | [optional] 
**Self** | Pointer to **string** | Link to the current page | [optional] 
**Next** | Pointer to **string** | Link to the next page | [optional] 

## Methods

### NewLinks

`func NewLinks() *Links`

NewLinks instantiates a new Links object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksWithDefaults

`func NewLinksWithDefaults() *Links`

NewLinksWithDefaults instantiates a new Links object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *Links) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Links) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Links) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Links) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetSelf

`func (o *Links) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Links) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Links) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Links) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *Links) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Links) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Links) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Links) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


