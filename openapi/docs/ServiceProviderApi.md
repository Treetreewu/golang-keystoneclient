# \ServiceProviderApi

All URIs are relative to *http://keystone-api.openstack.svc.cluster.local*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceProvider**](ServiceProviderApi.md#GetServiceProvider) | **Get** /v3/OS-FEDERATION/service_providers/{service_provider_id} | 
[**ListServiceProviders**](ServiceProviderApi.md#ListServiceProviders) | **Get** /v3/OS-FEDERATION/service_providers | 



## GetServiceProvider

> ServiceProviderResponse GetServiceProvider(ctx, serviceProviderId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceProviderId := "serviceProviderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceProviderApi.GetServiceProvider(context.Background(), serviceProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProviderApi.GetServiceProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceProvider`: ServiceProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceProviderApi.GetServiceProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceProviderResponse**](ServiceProviderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceProviders

> ServiceProviderListResponse ListServiceProviders(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceProviderApi.ListServiceProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProviderApi.ListServiceProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceProviders`: ServiceProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceProviderApi.ListServiceProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceProvidersRequest struct via the builder pattern


### Return type

[**ServiceProviderListResponse**](ServiceProviderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

