# UserUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username | [optional] 
**Role** | Pointer to **string** | User role | [optional] 
**Type** | Pointer to **string** | User type | [optional] 
**Verified** | Pointer to **bool** | User verified status | [optional] 
**BanStatus** | Pointer to **string** | User ban status | [optional] 
**TitleName** | Pointer to **string** | User title name | [optional] 
**Speciality** | Pointer to **string** | User speciality | [optional] 
**City** | Pointer to **string** | User city | [optional] 
**Country** | Pointer to **string** | User country | [optional] 
**Level** | Pointer to **int64** | User level | [optional] 
**Points** | Pointer to **int64** | User points | [optional] 

## Methods

### NewUserUpdateDataAttributes

`func NewUserUpdateDataAttributes() *UserUpdateDataAttributes`

NewUserUpdateDataAttributes instantiates a new UserUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateDataAttributesWithDefaults

`func NewUserUpdateDataAttributesWithDefaults() *UserUpdateDataAttributes`

NewUserUpdateDataAttributesWithDefaults instantiates a new UserUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserUpdateDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserUpdateDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserUpdateDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserUpdateDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRole

`func (o *UserUpdateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserUpdateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserUpdateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserUpdateDataAttributes) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *UserUpdateDataAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserUpdateDataAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserUpdateDataAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserUpdateDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerified

`func (o *UserUpdateDataAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserUpdateDataAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserUpdateDataAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserUpdateDataAttributes) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetBanStatus

`func (o *UserUpdateDataAttributes) GetBanStatus() string`

GetBanStatus returns the BanStatus field if non-nil, zero value otherwise.

### GetBanStatusOk

`func (o *UserUpdateDataAttributes) GetBanStatusOk() (*string, bool)`

GetBanStatusOk returns a tuple with the BanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStatus

`func (o *UserUpdateDataAttributes) SetBanStatus(v string)`

SetBanStatus sets BanStatus field to given value.

### HasBanStatus

`func (o *UserUpdateDataAttributes) HasBanStatus() bool`

HasBanStatus returns a boolean if a field has been set.

### GetTitleName

`func (o *UserUpdateDataAttributes) GetTitleName() string`

GetTitleName returns the TitleName field if non-nil, zero value otherwise.

### GetTitleNameOk

`func (o *UserUpdateDataAttributes) GetTitleNameOk() (*string, bool)`

GetTitleNameOk returns a tuple with the TitleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleName

`func (o *UserUpdateDataAttributes) SetTitleName(v string)`

SetTitleName sets TitleName field to given value.

### HasTitleName

`func (o *UserUpdateDataAttributes) HasTitleName() bool`

HasTitleName returns a boolean if a field has been set.

### GetSpeciality

`func (o *UserUpdateDataAttributes) GetSpeciality() string`

GetSpeciality returns the Speciality field if non-nil, zero value otherwise.

### GetSpecialityOk

`func (o *UserUpdateDataAttributes) GetSpecialityOk() (*string, bool)`

GetSpecialityOk returns a tuple with the Speciality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeciality

`func (o *UserUpdateDataAttributes) SetSpeciality(v string)`

SetSpeciality sets Speciality field to given value.

### HasSpeciality

`func (o *UserUpdateDataAttributes) HasSpeciality() bool`

HasSpeciality returns a boolean if a field has been set.

### GetCity

`func (o *UserUpdateDataAttributes) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserUpdateDataAttributes) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserUpdateDataAttributes) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserUpdateDataAttributes) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *UserUpdateDataAttributes) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserUpdateDataAttributes) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserUpdateDataAttributes) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserUpdateDataAttributes) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLevel

`func (o *UserUpdateDataAttributes) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UserUpdateDataAttributes) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UserUpdateDataAttributes) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *UserUpdateDataAttributes) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPoints

`func (o *UserUpdateDataAttributes) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *UserUpdateDataAttributes) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *UserUpdateDataAttributes) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *UserUpdateDataAttributes) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


