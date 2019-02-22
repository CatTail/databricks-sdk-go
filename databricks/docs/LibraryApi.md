# \LibraryApi

All URIs are relative to *https://petstore.swagger.io:80/api/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstallLibrary**](LibraryApi.md#InstallLibrary) | **Post** /libraries/install | 
[**UninstallLibrary**](LibraryApi.md#UninstallLibrary) | **Post** /libraries/uninstall | 


# **InstallLibrary**
> InstallLibrary(ctx, body)


Install libraries on a cluster. The installation is asynchronous - it happens in the background after the completion of this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**LibrariesInstallRequest**](LibrariesInstallRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninstallLibrary**
> UninstallLibrary(ctx, body)


Set libraries to be uninstalled on a cluster. The libraries aren’t uninstalled until the cluster is restarted. Uninstalling libraries that are not installed on the cluster has no impact but is not an error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**LibrariesUninstallRequest**](LibrariesUninstallRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

