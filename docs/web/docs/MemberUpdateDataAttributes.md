# MemberUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **string** | Team id | 
**UserId** | **string** | User id | 
**Role** | Pointer to **string** | User role | [optional] 
**Description** | Pointer to **string** | Description | [optional] 

## Methods

### NewMemberUpdateDataAttributes

`func NewMemberUpdateDataAttributes(teamId string, userId string, ) *MemberUpdateDataAttributes`

NewMemberUpdateDataAttributes instantiates a new MemberUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberUpdateDataAttributesWithDefaults

`func NewMemberUpdateDataAttributesWithDefaults() *MemberUpdateDataAttributes`

NewMemberUpdateDataAttributesWithDefaults instantiates a new MemberUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *MemberUpdateDataAttributes) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MemberUpdateDataAttributes) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MemberUpdateDataAttributes) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetUserId

`func (o *MemberUpdateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberUpdateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberUpdateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRole

`func (o *MemberUpdateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberUpdateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberUpdateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MemberUpdateDataAttributes) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDescription

`func (o *MemberUpdateDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MemberUpdateDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MemberUpdateDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MemberUpdateDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


