# UserOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreChangePasswordUponFirstUse** | Pointer to **bool** |  | [optional] 
**IgnorePasswordExpiry** | Pointer to **bool** |  | [optional] 
**IgnoreLockoutFailureAttempts** | Pointer to **bool** |  | [optional] 
**LockPassword** | Pointer to **bool** |  | [optional] 
**MultiFactorAuthEnabled** | Pointer to **bool** |  | [optional] 
**MultiFactorAuthRules** | Pointer to **bool** |  | [optional] 
**IgnoreUserInactivity** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserOptions

`func NewUserOptions() *UserOptions`

NewUserOptions instantiates a new UserOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOptionsWithDefaults

`func NewUserOptionsWithDefaults() *UserOptions`

NewUserOptionsWithDefaults instantiates a new UserOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreChangePasswordUponFirstUse

`func (o *UserOptions) GetIgnoreChangePasswordUponFirstUse() bool`

GetIgnoreChangePasswordUponFirstUse returns the IgnoreChangePasswordUponFirstUse field if non-nil, zero value otherwise.

### GetIgnoreChangePasswordUponFirstUseOk

`func (o *UserOptions) GetIgnoreChangePasswordUponFirstUseOk() (*bool, bool)`

GetIgnoreChangePasswordUponFirstUseOk returns a tuple with the IgnoreChangePasswordUponFirstUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChangePasswordUponFirstUse

`func (o *UserOptions) SetIgnoreChangePasswordUponFirstUse(v bool)`

SetIgnoreChangePasswordUponFirstUse sets IgnoreChangePasswordUponFirstUse field to given value.

### HasIgnoreChangePasswordUponFirstUse

`func (o *UserOptions) HasIgnoreChangePasswordUponFirstUse() bool`

HasIgnoreChangePasswordUponFirstUse returns a boolean if a field has been set.

### GetIgnorePasswordExpiry

`func (o *UserOptions) GetIgnorePasswordExpiry() bool`

GetIgnorePasswordExpiry returns the IgnorePasswordExpiry field if non-nil, zero value otherwise.

### GetIgnorePasswordExpiryOk

`func (o *UserOptions) GetIgnorePasswordExpiryOk() (*bool, bool)`

GetIgnorePasswordExpiryOk returns a tuple with the IgnorePasswordExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePasswordExpiry

`func (o *UserOptions) SetIgnorePasswordExpiry(v bool)`

SetIgnorePasswordExpiry sets IgnorePasswordExpiry field to given value.

### HasIgnorePasswordExpiry

`func (o *UserOptions) HasIgnorePasswordExpiry() bool`

HasIgnorePasswordExpiry returns a boolean if a field has been set.

### GetIgnoreLockoutFailureAttempts

`func (o *UserOptions) GetIgnoreLockoutFailureAttempts() bool`

GetIgnoreLockoutFailureAttempts returns the IgnoreLockoutFailureAttempts field if non-nil, zero value otherwise.

### GetIgnoreLockoutFailureAttemptsOk

`func (o *UserOptions) GetIgnoreLockoutFailureAttemptsOk() (*bool, bool)`

GetIgnoreLockoutFailureAttemptsOk returns a tuple with the IgnoreLockoutFailureAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreLockoutFailureAttempts

`func (o *UserOptions) SetIgnoreLockoutFailureAttempts(v bool)`

SetIgnoreLockoutFailureAttempts sets IgnoreLockoutFailureAttempts field to given value.

### HasIgnoreLockoutFailureAttempts

`func (o *UserOptions) HasIgnoreLockoutFailureAttempts() bool`

HasIgnoreLockoutFailureAttempts returns a boolean if a field has been set.

### GetLockPassword

`func (o *UserOptions) GetLockPassword() bool`

GetLockPassword returns the LockPassword field if non-nil, zero value otherwise.

### GetLockPasswordOk

`func (o *UserOptions) GetLockPasswordOk() (*bool, bool)`

GetLockPasswordOk returns a tuple with the LockPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockPassword

`func (o *UserOptions) SetLockPassword(v bool)`

SetLockPassword sets LockPassword field to given value.

### HasLockPassword

`func (o *UserOptions) HasLockPassword() bool`

HasLockPassword returns a boolean if a field has been set.

### GetMultiFactorAuthEnabled

`func (o *UserOptions) GetMultiFactorAuthEnabled() bool`

GetMultiFactorAuthEnabled returns the MultiFactorAuthEnabled field if non-nil, zero value otherwise.

### GetMultiFactorAuthEnabledOk

`func (o *UserOptions) GetMultiFactorAuthEnabledOk() (*bool, bool)`

GetMultiFactorAuthEnabledOk returns a tuple with the MultiFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthEnabled

`func (o *UserOptions) SetMultiFactorAuthEnabled(v bool)`

SetMultiFactorAuthEnabled sets MultiFactorAuthEnabled field to given value.

### HasMultiFactorAuthEnabled

`func (o *UserOptions) HasMultiFactorAuthEnabled() bool`

HasMultiFactorAuthEnabled returns a boolean if a field has been set.

### GetMultiFactorAuthRules

`func (o *UserOptions) GetMultiFactorAuthRules() bool`

GetMultiFactorAuthRules returns the MultiFactorAuthRules field if non-nil, zero value otherwise.

### GetMultiFactorAuthRulesOk

`func (o *UserOptions) GetMultiFactorAuthRulesOk() (*bool, bool)`

GetMultiFactorAuthRulesOk returns a tuple with the MultiFactorAuthRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiFactorAuthRules

`func (o *UserOptions) SetMultiFactorAuthRules(v bool)`

SetMultiFactorAuthRules sets MultiFactorAuthRules field to given value.

### HasMultiFactorAuthRules

`func (o *UserOptions) HasMultiFactorAuthRules() bool`

HasMultiFactorAuthRules returns a boolean if a field has been set.

### GetIgnoreUserInactivity

`func (o *UserOptions) GetIgnoreUserInactivity() bool`

GetIgnoreUserInactivity returns the IgnoreUserInactivity field if non-nil, zero value otherwise.

### GetIgnoreUserInactivityOk

`func (o *UserOptions) GetIgnoreUserInactivityOk() (*bool, bool)`

GetIgnoreUserInactivityOk returns a tuple with the IgnoreUserInactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserInactivity

`func (o *UserOptions) SetIgnoreUserInactivity(v bool)`

SetIgnoreUserInactivity sets IgnoreUserInactivity field to given value.

### HasIgnoreUserInactivity

`func (o *UserOptions) HasIgnoreUserInactivity() bool`

HasIgnoreUserInactivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


