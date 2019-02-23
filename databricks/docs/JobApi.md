# \JobApi

All URIs are relative to *https://petstore.swagger.io:80/api/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJob**](JobApi.md#CreateJob) | **Post** /jobs/create | 
[**DeleteJob**](JobApi.md#DeleteJob) | **Post** /jobs/delete | 
[**GetJob**](JobApi.md#GetJob) | **Get** /jobs/get | 
[**ResetJob**](JobApi.md#ResetJob) | **Post** /jobs/reset | 


# **CreateJob**
> JobsCreateResponse CreateJob(ctx, body)


Creates a new job with the provided settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsCreateRequest**](JobsCreateRequest.md)|  | 

### Return type

[**JobsCreateResponse**](JobsCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJob**
> DeleteJob(ctx, body)


Deletes the job and sends an email to the addresses specified in JobSettings.email_notifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsDeleteRequest**](JobsDeleteRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJob**
> JobsGetResponse GetJob(ctx, jobId)


Retrieves information about a single job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **int64**|  | 

### Return type

[**JobsGetResponse**](JobsGetResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetJob**
> ResetJob(ctx, body)


Overwrites the settings of a job with the provided settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsResetRequest**](JobsResetRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

