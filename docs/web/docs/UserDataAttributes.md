# UserDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username | 
**Role** | **string** | User role | 
**Avatar** | **string** | User avatar | 
**UpdatedAt** | Pointer to **time.Time** | User updated | [optional] 
**CreatedAt** | **time.Time** | User created at | 
**Organizations** | Pointer to **[]string** | User organization | [optional] 
**Projects** | Pointer to **[]string** | User projects | [optional] 
**Ideas** | Pointer to **[]string** | User ideas | [optional] 
**ReportsSent** | Pointer to **[]string** | User reports sent | [optional] 
**ReportsReceived** | Pointer to **[]string** | User reports received | [optional] 
**Ban** | Pointer to [**UserDataAttributesBan**](UserDataAttributesBan.md) |  | [optional] 

## Methods

### NewUserDataAttributes

`func NewUserDataAttributes(username string, role string, avatar string, createdAt time.Time, ) *UserDataAttributes`

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


### GetUpdatedAt

`func (o *UserDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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


### GetOrganizations

`func (o *UserDataAttributes) GetOrganizations() []string`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserDataAttributes) GetOrganizationsOk() (*[]string, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserDataAttributes) SetOrganizations(v []string)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *UserDataAttributes) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

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

### GetReportsSent

`func (o *UserDataAttributes) GetReportsSent() []string`

GetReportsSent returns the ReportsSent field if non-nil, zero value otherwise.

### GetReportsSentOk

`func (o *UserDataAttributes) GetReportsSentOk() (*[]string, bool)`

GetReportsSentOk returns a tuple with the ReportsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsSent

`func (o *UserDataAttributes) SetReportsSent(v []string)`

SetReportsSent sets ReportsSent field to given value.

### HasReportsSent

`func (o *UserDataAttributes) HasReportsSent() bool`

HasReportsSent returns a boolean if a field has been set.

### GetReportsReceived

`func (o *UserDataAttributes) GetReportsReceived() []string`

GetReportsReceived returns the ReportsReceived field if non-nil, zero value otherwise.

### GetReportsReceivedOk

`func (o *UserDataAttributes) GetReportsReceivedOk() (*[]string, bool)`

GetReportsReceivedOk returns a tuple with the ReportsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsReceived

`func (o *UserDataAttributes) SetReportsReceived(v []string)`

SetReportsReceived sets ReportsReceived field to given value.

### HasReportsReceived

`func (o *UserDataAttributes) HasReportsReceived() bool`

HasReportsReceived returns a boolean if a field has been set.

### GetBan

`func (o *UserDataAttributes) GetBan() UserDataAttributesBan`

GetBan returns the Ban field if non-nil, zero value otherwise.

### GetBanOk

`func (o *UserDataAttributes) GetBanOk() (*UserDataAttributesBan, bool)`

GetBanOk returns a tuple with the Ban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBan

`func (o *UserDataAttributes) SetBan(v UserDataAttributesBan)`

SetBan sets Ban field to given value.

### HasBan

`func (o *UserDataAttributes) HasBan() bool`

HasBan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


