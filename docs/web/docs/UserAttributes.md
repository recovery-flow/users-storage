# UserAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username | 
**Avatar** | Pointer to **string** | Avatar | [optional] 
**Role** | **string** | Role | 
**Type** | **string** | Type | 
**Verified** | **bool** | Verified | 
**BanStatus** | **string** | Ban status | 
**TitleName** | Pointer to **string** | Title | [optional] 
**Speciality** | Pointer to **string** | Speciality | [optional] 
**City** | Pointer to **string** | City | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**Level** | **int32** | Level | 
**Points** | **int32** | Points | 
**UpdatedAt** | Pointer to **time.Time** | Updated at | [optional] 
**CreatedAt** | **time.Time** | Created at | 

## Methods

### NewUserAttributes

`func NewUserAttributes(username string, role string, type_ string, verified bool, banStatus string, level int32, points int32, createdAt time.Time, ) *UserAttributes`

NewUserAttributes instantiates a new UserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesWithDefaults

`func NewUserAttributesWithDefaults() *UserAttributes`

NewUserAttributesWithDefaults instantiates a new UserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAvatar

`func (o *UserAttributes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserAttributes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserAttributes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserAttributes) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetRole

`func (o *UserAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetType

`func (o *UserAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetVerified

`func (o *UserAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetBanStatus

`func (o *UserAttributes) GetBanStatus() string`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *UserAttributes) GetBanStatusOk() (*string, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *UserAttributes) SetBanStatus(v string)`

SetBanStatus sets BanStatus field to given value.


### GetTitleName

`func (o *UserAttributes) GetTitleName() string`

GetTitleName returns the TitleName field if non-nil, zero value otherwise.

### GetTitleNameOk

`func (o *UserAttributes) GetTitleNameOk() (*string, bool)`

GetTitleNameOk returns a tuple with the TitleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleName

`func (o *UserAttributes) SetTitleName(v string)`

SetTitleName sets TitleName field to given value.

### HasTitleName

`func (o *UserAttributes) HasTitleName() bool`

HasTitleName returns a boolean if a field has been set.

### GetSpeciality

`func (o *UserAttributes) GetSpeciality() string`

GetSpeciality returns the Speciality field if non-nil, zero value otherwise.

### GetSpecialityOk

`func (o *UserAttributes) GetSpecialityOk() (*string, bool)`

GetSpecialityOk returns a tuple with the Speciality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeciality

`func (o *UserAttributes) SetSpeciality(v string)`

SetSpeciality sets Speciality field to given value.

### HasSpeciality

`func (o *UserAttributes) HasSpeciality() bool`

HasSpeciality returns a boolean if a field has been set.

### GetCity

`func (o *UserAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *UserAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLevel

`func (o *UserAttributes) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UserAttributes) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UserAttributes) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetPoints

`func (o *UserAttributes) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *UserAttributes) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *UserAttributes) SetPoints(v int32)`

SetPoints sets Points field to given value.


### GetUpdatedAt

`func (o *UserAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


