# \ClusterApi

All URIs are relative to *https://petstore.swagger.io:80/api/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClusterApi.md#CreateCluster) | **Post** /clusters/create | 
[**EditCluster**](ClusterApi.md#EditCluster) | **Post** /clusters/edit | 
[**GetCluster**](ClusterApi.md#GetCluster) | **Get** /clusters/get | 
[**PermanentDeleteCluster**](ClusterApi.md#PermanentDeleteCluster) | **Post** /clusters/permanent-delete | 


# **CreateCluster**
> ClustersCreateResponse CreateCluster(ctx, body)


Creates a new Spark cluster. This method acquires new instances from the cloud provider if necessary. This method is asynchronous; the returned cluster_id can be used to poll the cluster state. When this method returns, the cluster is in a PENDING state. The cluster is usable once it enters a RUNNING state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**NewCluster**](NewCluster.md)|  | 

### Return type

[**ClustersCreateResponse**](ClustersCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditCluster**
> EditCluster(ctx, body)


Edits the configuration of a cluster to match the provided attributes and size.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ClustersEditRequest**](ClustersEditRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCluster**
> ClustersGetResponse GetCluster(ctx, clusterId)


Retrieves the information for a cluster given its identifier. Clusters can be described while they are running, or up to 30 days after they are terminated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterId** | **string**|  | 

### Return type

[**ClustersGetResponse**](ClustersGetResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PermanentDeleteCluster**
> PermanentDeleteCluster(ctx, body)


Terminates a Spark cluster given its id. The cluster is removed asynchronously. Once the termination has completed, the cluster will be in a TERMINATED state. If the cluster is already in a TERMINATING or TERMINATED state, nothing will happen.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ClustersPermanentDeleteRequest**](ClustersPermanentDeleteRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

