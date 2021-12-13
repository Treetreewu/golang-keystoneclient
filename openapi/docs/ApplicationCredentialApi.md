# \ApplicationCredentialApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationCredential**](ApplicationCredentialApi.md#CreateApplicationCredential) | **Post** /v3/users/{user_id}/application_credentials | 
[**DeleteApplicationCredential**](ApplicationCredentialApi.md#DeleteApplicationCredential) | **Delete** /v3/users/{user_id}/application_credentials/{application_credential_id} | 
[**ListUserApplicationCredentials**](ApplicationCredentialApi.md#ListUserApplicationCredentials) | **Get** /v3/users/{user_id}/application_credentials | 



## CreateApplicationCredential

> ApplicationCredentialResponse CreateApplicationCredential(ctx, userId).ApplicationCredentialCreateRequest(applicationCredentialCreateRequest).Execute()





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
    userId := "userId_example" // string | The ID of the user who owns the application credential.
    applicationCredentialCreateRequest := *openapiclient.NewApplicationCredentialCreateRequest(*openapiclient.NewApplicationCredentialCreate("Name_example")) // ApplicationCredentialCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationCredentialApi.CreateApplicationCredential(context.Background(), userId).ApplicationCredentialCreateRequest(applicationCredentialCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialApi.CreateApplicationCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationCredential`: ApplicationCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialApi.CreateApplicationCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user who owns the application credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationCredentialCreateRequest** | [**ApplicationCredentialCreateRequest**](ApplicationCredentialCreateRequest.md) |  | 

### Return type

[**ApplicationCredentialResponse**](ApplicationCredentialResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationCredential

> DeleteApplicationCredential(ctx, userId, applicationCredentialId).Execute()





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
    userId := "userId_example" // string | The ID of the user who owns the application credential.
    applicationCredentialId := "applicationCredentialId_example" // string | The ID of the application credential.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationCredentialApi.DeleteApplicationCredential(context.Background(), userId, applicationCredentialId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialApi.DeleteApplicationCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user who owns the application credential. | 
**applicationCredentialId** | **string** | The ID of the application credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationCredentialRequest struct via the builder pattern


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


## ListUserApplicationCredentials

> ApplicationCredentialListResponse ListUserApplicationCredentials(ctx, userId).Name(name).Execute()



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
    userId := "userId_example" // string | The ID of the user who owns the application credential.
    name := "name_example" // string | The name of the application credential. Must be unique to a user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationCredentialApi.ListUserApplicationCredentials(context.Background(), userId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCredentialApi.ListUserApplicationCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserApplicationCredentials`: ApplicationCredentialListResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCredentialApi.ListUserApplicationCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The ID of the user who owns the application credential. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserApplicationCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the application credential. Must be unique to a user. | 

### Return type

[**ApplicationCredentialListResponse**](ApplicationCredentialListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

