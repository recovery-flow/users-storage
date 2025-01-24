# UserDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username | 
**Description** | Pointer to **string** | User description | [optional] 
**Role** | **string** | User role | 
**Avatar** | **string** | User avatar | 
**CreatedAt** | **time.Time** | User created at | 
**Projects** | Pointer to **[]string** | User projects | [optional] 
**Ideas** | Pointer to **[]string** | User ideas | [optional] 
**ReportSent** | Pointer to **[]string** | User reports sent | [optional] 
**ReportReceived** | Pointer to **[]string** | User reports received | [optional] 
**BanStatus** | **bool** | User ban status | 

## Methods

### NewUserDataAttributes

`func NewUserDataAttributes(username string, role string, avatar string, createdAt time.Time, banStatus bool, ) *UserDataAttributes`

NewUserDataAttributes instantiates a new UserDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataAttributesWithDefaults

`func NewUserDataAttributesWithDefaults() *UserDataAttributes`

NewUserDataAttributesWithDefaults instantiates a new UserDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *UserDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *UserDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetAvatar

`func (o *UserDataAttributes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserDataAttributes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserDataAttributes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetCreatedAt

`func (o *UserDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProjects

`func (o *UserDataAttributes) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UserDataAttributes) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UserDataAttributes) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UserDataAttributes) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetIdeas

`func (o *UserDataAttributes) GetIdeas() []string`

GetIdeas returns the Ideas field if non-nil, zero value otherwise.

### GetIdeasOk

`func (o *UserDataAttributes) GetIdeasOk() (*[]string, bool)`

GetIdeasOk returns a tuple with the Ideas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeas

`func (o *UserDataAttributes) SetIdeas(v []string)`

SetIdeas sets Ideas field to given value.

### HasIdeas

`func (o *UserDataAttributes) HasIdeas() bool`

HasIdeas returns a boolean if a field has been set.

### GetReportSent

`func (o *UserDataAttributes) GetReportSent() []string`

GetReportSent returns the ReportSent field if non-nil, zero value otherwise.

### GetReportSentOk

`func (o *UserDataAttributes) GetReportSentOk() (*[]string, bool)`

GetReportSentOk returns a tuple with the ReportSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportSent

`func (o *UserDataAttributes) SetReportSent(v []string)`

SetReportSent sets ReportSent field to given value.

### HasReportSent

`func (o *UserDataAttributes) HasReportSent() bool`

HasReportSent returns a boolean if a field has been set.

### GetReportReceived

`func (o *UserDataAttributes) GetReportReceived() []string`

GetReportReceived returns the ReportReceived field if non-nil, zero value otherwise.

### GetReportReceivedOk

`func (o *UserDataAttributes) GetReportReceivedOk() (*[]string, bool)`

GetReportReceivedOk returns a tuple with the ReportReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportReceived

`func (o *UserDataAttributes) SetReportReceived(v []string)`

SetReportReceived sets ReportReceived field to given value.

### HasReportReceived

`func (o *UserDataAttributes) HasReportReceived() bool`

HasReportReceived returns a boolean if a field has been set.

### GetBanStatus

`func (o *UserDataAttributes) GetBanStatus() bool`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *UserDataAttributes) GetBanStatusOk() (*bool, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *UserDataAttributes) SetBanStatus(v bool)`

SetBanStatus sets BanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


