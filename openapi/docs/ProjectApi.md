# \ProjectApi

All URIs are relative to *http://keystone-api.openstack.svc.cluster.local*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /v3/projects | 
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /v3/projects/{project_id} | 
[**GetProject**](ProjectApi.md#GetProject) | **Get** /v3/projects/{project_id} | 
[**ListProjects**](ProjectApi.md#ListProjects) | **Get** /v3/projects | 
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Patch** /v3/projects/{project_id} | 



## CreateProject

> ProjectResponse CreateProject(ctx).ProjectCreateRequest(projectCreateRequest).Execute()





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
    projectCreateRequest := *openapiclient.NewProjectCreateRequest(*openapiclient.NewProjectCreate("Name_example")) // ProjectCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.CreateProject(context.Background()).ProjectCreateRequest(projectCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateRequest** | [**ProjectCreateRequest**](ProjectCreateRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectId).Execute()





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
    projectId := "projectId_example" // string | The project ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## GetProject

> ProjectGetResponse GetProject(ctx, projectId).ParentsAsList(parentsAsList).SubtreeAsList(subtreeAsList).ParentsAsIds(parentsAsIds).SubtreeAsIds(subtreeAsIds).Execute()





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
    projectId := "projectId_example" // string | The project ID.
    parentsAsList := int32(56) // int32 | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The parent hierarchy will be included as a list in the response. This list will contain the projects found by traversing up the hierarchy to the top-level project. The returned list will be filtered against the projects the user has an effective role assignment on.  (optional) (default to 0)
    subtreeAsList := int32(56) // int32 | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The child hierarchy will be included as a list in the response. This list will contain the projects found by traversing down the hierarchy. The returned list will be filtered against the projects the user has an effective role assignment on.  (optional) (default to 0)
    parentsAsIds := int32(56) // int32 | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The entire parent hierarchy will be included as nested dictionaries in the response. It will contain all projects ids found by traversing up the hierarchy to the top-level project.  (optional) (default to 0)
    subtreeAsIds := int32(56) // int32 | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The entire child hierarchy will be included as nested dictionaries in the response. It will contain all the projects ids found by traversing down the hierarchy.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.GetProject(context.Background(), projectId).ParentsAsList(parentsAsList).SubtreeAsList(subtreeAsList).ParentsAsIds(parentsAsIds).SubtreeAsIds(subtreeAsIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentsAsList** | **int32** | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The parent hierarchy will be included as a list in the response. This list will contain the projects found by traversing up the hierarchy to the top-level project. The returned list will be filtered against the projects the user has an effective role assignment on.  | [default to 0]
 **subtreeAsList** | **int32** | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The child hierarchy will be included as a list in the response. This list will contain the projects found by traversing down the hierarchy. The returned list will be filtered against the projects the user has an effective role assignment on.  | [default to 0]
 **parentsAsIds** | **int32** | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The entire parent hierarchy will be included as nested dictionaries in the response. It will contain all projects ids found by traversing up the hierarchy to the top-level project.  | [default to 0]
 **subtreeAsIds** | **int32** | 0 is treated as False. Any other value is considered to be equivalent to True, including the absence of a value.  The entire child hierarchy will be included as nested dictionaries in the response. It will contain all the projects ids found by traversing down the hierarchy.  | [default to 0]

### Return type

[**ProjectGetResponse**](ProjectGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> ProjectListResponse ListProjects(ctx).Enabled(enabled).Name(name).DomainId(domainId).ParentId(parentId).IsDomain(isDomain).Execute()





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
    domainId := "domainId_example" // string | Filters the response by a domain ID. (optional)
    parentId := "parentId_example" // string | Filters the response by a parent ID. (optional) (default to "false")
    isDomain := true // bool | If this is specified as `true`, then only projects acting as a domain are included. Otherwise, only projects that are not acting as a domain are included. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.ListProjects(context.Background()).Enabled(enabled).Name(name).DomainId(domainId).ParentId(parentId).IsDomain(isDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: ProjectListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | If set to true, then only enabled resources will be returned. Any value other than 0 (including no value) will be interpreted as true. | 
 **name** | **string** | Filters the response by name. | 
 **domainId** | **string** | Filters the response by a domain ID. | 
 **parentId** | **string** | Filters the response by a parent ID. | [default to &quot;false&quot;]
 **isDomain** | **bool** | If this is specified as &#x60;true&#x60;, then only projects acting as a domain are included. Otherwise, only projects that are not acting as a domain are included. | 

### Return type

[**ProjectListResponse**](ProjectListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> ProjectResponse UpdateProject(ctx, projectId).ProjectUpdateRequest(projectUpdateRequest).Execute()





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
    projectId := "projectId_example" // string | The project ID.
    projectUpdateRequest := *openapiclient.NewProjectUpdateRequest(*openapiclient.NewProjectUpdate()) // ProjectUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.UpdateProject(context.Background(), projectId).ProjectUpdateRequest(projectUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectUpdateRequest** | [**ProjectUpdateRequest**](ProjectUpdateRequest.md) |  | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

