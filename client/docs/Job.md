# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Type** | **int64** | 0 - local, 1 - remote | [optional] 
**Command** | **string** |  | [optional] 
**Owner** | **string** |  | [optional] 
**Disabled** | **bool** |  | [optional] 
**Schedule** | **string** |  | [optional] 
**Retries** | **int64** |  | [optional] 
**Epsilon** | **string** | Duration in which it is safe to retry the Job. | [optional] 
**NextRunAt** | [**time.Time**](time.Time.md) |  | [optional] 
**RemoteProperties** | [**RemoteProperties**](RemoteProperties.md) |  | [optional] 
**IsDone** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


