# MemberDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**MemberDataRelationshipsTeams**](MemberDataRelationshipsTeams.md) |  | [optional] 
**User** | Pointer to [**MemberDataRelationshipsUser**](MemberDataRelationshipsUser.md) |  | [optional] 

## Methods

### NewMemberDataRelationships

`func NewMemberDataRelationships() *MemberDataRelationships`

NewMemberDataRelationships instantiates a new MemberDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDataRelationshipsWithDefaults

`func NewMemberDataRelationshipsWithDefaults() *MemberDataRelationships`

NewMemberDataRelationshipsWithDefaults instantiates a new MemberDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *MemberDataRelationships) GetTeams() MemberDataRelationshipsTeams`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *MemberDataRelationships) GetTeamsOk() (*MemberDataRelationshipsTeams, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *MemberDataRelationships) SetTeams(v MemberDataRelationshipsTeams)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *MemberDataRelationships) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetUser

`func (o *MemberDataRelationships) GetUser() MemberDataRelationshipsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MemberDataRelationships) GetUserOk() (*MemberDataRelationshipsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MemberDataRelationships) SetUser(v MemberDataRelationshipsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *MemberDataRelationships) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


