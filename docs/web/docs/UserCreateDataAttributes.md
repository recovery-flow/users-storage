# UserCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username | 
**Title** | Pointer to **string** | User title | [optional] 
**Status** | Pointer to **string** | User status | [optional] 
**Avatar** | Pointer to **string** | User avatar | [optional] 
**Bio** | Pointer to **string** | User bio | [optional] 

## Methods

### NewUserCreateDataAttributes

`func NewUserCreateDataAttributes(username string, ) *UserCreateDataAttributes`

NewUserCreateDataAttributes instantiates a new UserCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateDataAttributesWithDefaults

`func NewUserCreateDataAttributesWithDefaults() *UserCreateDataAttributes`

NewUserCreateDataAttributesWithDefaults instantiates a new UserCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserCreateDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreateDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreateDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetTitle

`func (o *UserCreateDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserCreateDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserCreateDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserCreateDataAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *UserCreateDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCreateDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCreateDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCreateDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvatar

`func (o *UserCreateDataAttributes) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserCreateDataAttributes) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserCreateDataAttributes) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserCreateDataAttributes) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBio

`func (o *UserCreateDataAttributes) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserCreateDataAttributes) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserCreateDataAttributes) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UserCreateDataAttributes) HasBio() bool`

HasBio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


