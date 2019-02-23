# JobsCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCluster** | [***NewCluster**](NewCluster.md) |  | [optional] [default to null]
**ExistingClusterId** | **string** |  | [optional] [default to null]
**NotebookTask** | [***NotebookTask**](NotebookTask.md) |  | [optional] [default to null]
**SparkJarTask** | [***SparkJarTask**](SparkJarTask.md) |  | [optional] [default to null]
**SparkPythonTask** | [***SparkPythonTask**](SparkPythonTask.md) |  | [optional] [default to null]
**SparkSubmitTask** | [***SparkSubmitTask**](SparkSubmitTask.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Libraries** | [**[]Library**](Library.md) |  | [optional] [default to null]
**EmailNotifications** | [***JobEmailNotifications**](JobEmailNotifications.md) |  | [optional] [default to null]
**TimeoutSeconds** | **int32** |  | [optional] [default to null]
**MaxRetries** | **int32** |  | [optional] [default to null]
**MinRetryIntervalMillis** | **int32** |  | [optional] [default to null]
**RetryOnTimeout** | **bool** |  | [optional] [default to null]
**Schedule** | [***CronSchedule**](CronSchedule.md) |  | [optional] [default to null]
**MaxConcurrentRuns** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


