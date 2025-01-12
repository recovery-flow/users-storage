# TeamsDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Team name | 
**Description** | **string** | Team description | 
**Members** | [**[]Object**](Object.md) |  | 
**CreatedAt** | **time.Time** | Team created at | 

## Methods

### NewTeamsDataAttributes

`func NewTeamsDataAttributes(name string, description string, members []Object, createdAt time.Time, ) *TeamsDataAttributes`

NewTeamsDataAttributes instantiates a new TeamsDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsDataAttributesWithDefaults

`func NewTeamsDataAttributesWithDefaults() *TeamsDataAttributes`

NewTeamsDataAttributesWithDefaults instantiates a new TeamsDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamsDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamsDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamsDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamsDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamsDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamsDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMembers

`func (o *TeamsDataAttributes) GetMembers() []Object`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamsDataAttributes) GetMembersOk() (*[]Object, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamsDataAttributes) SetMembers(v []Object)`

SetMembers sets Members field to given value.


### GetCreatedAt

`func (o *TeamsDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamsDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamsDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


