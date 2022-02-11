# \DomainApi

All URIs are relative to *http://keystone-api.openstack.svc.cluster.local*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainApi.md#CreateDomain) | **Post** /v3/domains | 
[**CreateDomainConfig**](DomainApi.md#CreateDomainConfig) | **Put** /v3/domains/{domain_id}/config | 
[**DeleteDomain**](DomainApi.md#DeleteDomain) | **Delete** /v3/domains/{domain_id} | 
[**DeleteDomainConfig**](DomainApi.md#DeleteDomainConfig) | **Delete** /v3/domains/{domain_id}/config | 
[**GetDomain**](DomainApi.md#GetDomain) | **Get** /v3/domains/{domain_id} | 
[**GetDomainConfig**](DomainApi.md#GetDomainConfig) | **Get** /v3/domains/{domain_id}/config | 
[**ListDomains**](DomainApi.md#ListDomains) | **Get** /v3/domains | 
[**UpdateDomain**](DomainApi.md#UpdateDomain) | **Patch** /v3/domains/{domain_id} | 
[**UpdateDomainConfig**](DomainApi.md#UpdateDomainConfig) | **Patch** /v3/domains/{domain_id}/config | 
[**VerifyDomainConfig**](DomainApi.md#VerifyDomainConfig) | **Post** /v3/domains/{domain_id}/config/verify | 



## CreateDomain

> DomainResponse CreateDomain(ctx).DomainCreateRequest(domainCreateRequest).Execute()





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
    domainCreateRequest := *openapiclient.NewDomainCreateRequest(*openapiclient.NewDomainCreate("Name_example")) // DomainCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.CreateDomain(context.Background()).DomainCreateRequest(domainCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainCreateRequest** | [**DomainCreateRequest**](DomainCreateRequest.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDomainConfig

> DomainConfig CreateDomainConfig(ctx, domainId).DomainConfig(domainConfig).Execute()





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
    domainId := "domainId_example" // string | The domain ID.
    domainConfig := *openapiclient.NewDomainConfig(*openapiclient.NewDomainConfigConfig()) // DomainConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.CreateDomainConfig(context.Background(), domainId).DomainConfig(domainConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.CreateDomainConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomainConfig`: DomainConfig
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.CreateDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainConfig** | [**DomainConfig**](DomainConfig.md) |  | 

### Return type

[**DomainConfig**](DomainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> DeleteDomain(ctx, domainId).Execute()





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
    domainId := "domainId_example" // string | The domain ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.DeleteDomain(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomainConfig

> DeleteDomainConfig(ctx, domainId).Execute()





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
    domainId := "domainId_example" // string | The domain ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.DeleteDomainConfig(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.DeleteDomainConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomain

> DomainResponse GetDomain(ctx, domainId).Execute()





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
    domainId := "domainId_example" // string | The domain ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.GetDomain(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainConfig

> DomainConfig GetDomainConfig(ctx, domainId).Execute()





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
    domainId := "domainId_example" // string | The domain ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.GetDomainConfig(context.Background(), domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.GetDomainConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomainConfig`: DomainConfig
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.GetDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainConfig**](DomainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> DomainListResponse ListDomains(ctx).Enabled(enabled).Name(name).RootOnly(rootOnly).Execute()





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
    enabled := true // bool | If set to true, then only enabled resources will be returned. Any value other than 0 (including no value) will be interpreted as true. (optional)
    name := "name_example" // string | Filters the response by name. (optional)
    rootOnly := true // bool | Return root domains only. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.ListDomains(context.Background()).Enabled(enabled).Name(name).RootOnly(rootOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.ListDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomains`: DomainListResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.ListDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | If set to true, then only enabled resources will be returned. Any value other than 0 (including no value) will be interpreted as true. | 
 **name** | **string** | Filters the response by name. | 
 **rootOnly** | **bool** | Return root domains only. | [default to false]

### Return type

[**DomainListResponse**](DomainListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomain

> DomainResponse UpdateDomain(ctx, domainId).DomainUpdateRequest(domainUpdateRequest).Execute()





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
    domainId := "domainId_example" // string | The domain ID.
    domainUpdateRequest := *openapiclient.NewDomainUpdateRequest(*openapiclient.NewDomainUpdate()) // DomainUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.UpdateDomain(context.Background(), domainId).DomainUpdateRequest(domainUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.UpdateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomain`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainUpdateRequest** | [**DomainUpdateRequest**](DomainUpdateRequest.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainConfig

> DomainConfig UpdateDomainConfig(ctx, domainId).DomainConfig(domainConfig).Execute()





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
    domainId := "domainId_example" // string | The domain ID.
    domainConfig := *openapiclient.NewDomainConfig(*openapiclient.NewDomainConfigConfig()) // DomainConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.UpdateDomainConfig(context.Background(), domainId).DomainConfig(domainConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.UpdateDomainConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomainConfig`: DomainConfig
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.UpdateDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainConfig** | [**DomainConfig**](DomainConfig.md) |  | 

### Return type

[**DomainConfig**](DomainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyDomainConfig

> DomainConfigVerifyResponse VerifyDomainConfig(ctx, domainId).DomainConfigVerifyRequest(domainConfigVerifyRequest).Execute()





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
    domainId := "domainId_example" // string | The domain ID.
    domainConfigVerifyRequest := *openapiclient.NewDomainConfigVerifyRequest() // DomainConfigVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainApi.VerifyDomainConfig(context.Background(), domainId).DomainConfigVerifyRequest(domainConfigVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainApi.VerifyDomainConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyDomainConfig`: DomainConfigVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainApi.VerifyDomainConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyDomainConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainConfigVerifyRequest** | [**DomainConfigVerifyRequest**](DomainConfigVerifyRequest.md) |  | 

### Return type

[**DomainConfigVerifyResponse**](DomainConfigVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

