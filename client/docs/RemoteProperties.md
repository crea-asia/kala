# RemoteProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**Timeout** | Pointer to **int64** |  | [optional] 
**ExpectedResponseCodes** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewRemoteProperties

`func NewRemoteProperties() *RemoteProperties`

NewRemoteProperties instantiates a new RemoteProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePropertiesWithDefaults

`func NewRemotePropertiesWithDefaults() *RemoteProperties`

NewRemotePropertiesWithDefaults instantiates a new RemoteProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RemoteProperties) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteProperties) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteProperties) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RemoteProperties) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMethod

`func (o *RemoteProperties) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RemoteProperties) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RemoteProperties) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RemoteProperties) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetBody

`func (o *RemoteProperties) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RemoteProperties) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RemoteProperties) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *RemoteProperties) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *RemoteProperties) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RemoteProperties) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RemoteProperties) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RemoteProperties) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTimeout

`func (o *RemoteProperties) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RemoteProperties) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RemoteProperties) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RemoteProperties) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetExpectedResponseCodes

`func (o *RemoteProperties) GetExpectedResponseCodes() []int64`

GetExpectedResponseCodes returns the ExpectedResponseCodes field if non-nil, zero value otherwise.

### GetExpectedResponseCodesOk

`func (o *RemoteProperties) GetExpectedResponseCodesOk() (*[]int64, bool)`

GetExpectedResponseCodesOk returns a tuple with the ExpectedResponseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponseCodes

`func (o *RemoteProperties) SetExpectedResponseCodes(v []int64)`

SetExpectedResponseCodes sets ExpectedResponseCodes field to given value.

### HasExpectedResponseCodes

`func (o *RemoteProperties) HasExpectedResponseCodes() bool`

HasExpectedResponseCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


