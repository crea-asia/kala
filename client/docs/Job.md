# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int64** | 0 - local, 1 - remote | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Retries** | Pointer to **int64** |  | [optional] 
**Epsilon** | Pointer to **string** |  | [optional] 
**NextRunAt** | Pointer to **string** |  | [optional] 
**RemoteProperties** | Pointer to [**RemoteProperties**](RemoteProperties.md) |  | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCommand

`func (o *Job) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Job) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Job) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Job) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetOwner

`func (o *Job) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Job) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Job) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Job) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDisabled

`func (o *Job) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Job) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Job) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Job) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetSchedule

`func (o *Job) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Job) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Job) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Job) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRetries

`func (o *Job) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Job) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Job) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *Job) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetEpsilon

`func (o *Job) GetEpsilon() string`

GetEpsilon returns the Epsilon field if non-nil, zero value otherwise.

### GetEpsilonOk

`func (o *Job) GetEpsilonOk() (*string, bool)`

GetEpsilonOk returns a tuple with the Epsilon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsilon

`func (o *Job) SetEpsilon(v string)`

SetEpsilon sets Epsilon field to given value.

### HasEpsilon

`func (o *Job) HasEpsilon() bool`

HasEpsilon returns a boolean if a field has been set.

### GetNextRunAt

`func (o *Job) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *Job) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *Job) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *Job) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetRemoteProperties

`func (o *Job) GetRemoteProperties() RemoteProperties`

GetRemoteProperties returns the RemoteProperties field if non-nil, zero value otherwise.

### GetRemotePropertiesOk

`func (o *Job) GetRemotePropertiesOk() (*RemoteProperties, bool)`

GetRemotePropertiesOk returns a tuple with the RemoteProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProperties

`func (o *Job) SetRemoteProperties(v RemoteProperties)`

SetRemoteProperties sets RemoteProperties field to given value.

### HasRemoteProperties

`func (o *Job) HasRemoteProperties() bool`

HasRemoteProperties returns a boolean if a field has been set.

### GetIsDone

`func (o *Job) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *Job) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *Job) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *Job) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


