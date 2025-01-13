# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TeamData**](TeamData.md) |  | 
**Included** | Pointer to [**[]Object**](Object.md) |  | [optional] 

## Methods

### NewTeam

`func NewTeam(data TeamData, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Team) GetData() TeamData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Team) GetDataOk() (*TeamData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Team) SetData(v TeamData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *Team) GetIncluded() []Object`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *Team) GetIncludedOk() (*[]Object, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *Team) SetIncluded(v []Object)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *Team) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


