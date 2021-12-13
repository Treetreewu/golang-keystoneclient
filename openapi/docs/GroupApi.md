# \GroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](GroupApi.md#AddUser) | **Put** /v3/groups/{group_id}/users/{user_id} | 
[**CheckUser**](GroupApi.md#CheckUser) | **Head** /v3/groups/{group_id}/users/{user_id} | 
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /v3/groups | 
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /v3/groups/{group_id} | 
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /v3/groups/{group_id} | 
[**ListGroupUsers**](GroupApi.md#ListGroupUsers) | **Get** /v3/groups/{group_id}/users | 
[**ListGroups**](GroupApi.md#ListGroups) | **Get** /v3/groups | 
[**RemoveUser**](GroupApi.md#RemoveUser) | **Delete** /v3/groups/{group_id}/users/{user_id} | 
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Patch** /v3/groups/{group_id} | 



## AddUser

> AddUser(ctx, groupId, userId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.AddUser(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.AddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


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


## CheckUser

> CheckUser(ctx, groupId, userId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.CheckUser(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CheckUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserRequest struct via the builder pattern


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


## CreateGroup

> GroupResponse CreateGroup(ctx).GroupCreateRequest(groupCreateRequest).Execute()





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
    groupCreateRequest := *openapiclient.NewGroupCreateRequest(*openapiclient.NewGroupCreate("Name_example")) // GroupCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.CreateGroup(context.Background()).GroupCreateRequest(groupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateRequest** | [**GroupCreateRequest**](GroupCreateRequest.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()





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
    groupId := "groupId_example" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> GroupResponse GetGroup(ctx, groupId).Execute()





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
    groupId := "groupId_example" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.GetGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupUsers

> UserListResponse ListGroupUsers(ctx, groupId).PasswordExpiresAt(passwordExpiresAt).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    passwordExpiresAt := "lt:2016-12-08T22:02:00Z" // string | Filter results based on which user passwords have expired. The query should include an `operator` and a `timestamp` with a colon (`:`) separating the two, for example:  ```password_expires_at={operator}:{timestamp}```  * Valid operators are: `lt`, `lte`, `gt`, `gte`, `eq`, and `neq`   * lt: expiration time lower than the timestamp   * lte: expiration time lower than or equal to the timestamp   * gt: expiration time higher than the timestamp   * gte: expiration time higher than or equal to the timestamp   * eq: expiration time equal to the timestamp   * neq: expiration time not equal to the timestamp * Valid timestamps are of the form: `YYYY-MM-DDTHH:mm:ssZ`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.ListGroupUsers(context.Background(), groupId).PasswordExpiresAt(passwordExpiresAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupUsers`: UserListResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordExpiresAt** | **string** | Filter results based on which user passwords have expired. The query should include an &#x60;operator&#x60; and a &#x60;timestamp&#x60; with a colon (&#x60;:&#x60;) separating the two, for example:  &#x60;&#x60;&#x60;password_expires_at&#x3D;{operator}:{timestamp}&#x60;&#x60;&#x60;  * Valid operators are: &#x60;lt&#x60;, &#x60;lte&#x60;, &#x60;gt&#x60;, &#x60;gte&#x60;, &#x60;eq&#x60;, and &#x60;neq&#x60;   * lt: expiration time lower than the timestamp   * lte: expiration time lower than or equal to the timestamp   * gt: expiration time higher than the timestamp   * gte: expiration time higher than or equal to the timestamp   * eq: expiration time equal to the timestamp   * neq: expiration time not equal to the timestamp * Valid timestamps are of the form: &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;  | 

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> GroupListResponse ListGroups(ctx).Name(name).DomainId(domainId).Execute()





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
    name := "name_example" // string | Filters the response by name. (optional)
    domainId := "domainId_example" // string | Filters the response by a domain ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.ListGroups(context.Background()).Name(name).DomainId(domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: GroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filters the response by name. | 
 **domainId** | **string** | Filters the response by a domain ID. | 

### Return type

[**GroupListResponse**](GroupListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUser

> RemoveUser(ctx, groupId, userId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.RemoveUser(context.Background(), groupId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.RemoveUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserRequest struct via the builder pattern


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


## UpdateGroup

> GroupResponse UpdateGroup(ctx, groupId).GroupUpdateRequest(groupUpdateRequest).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    groupUpdateRequest := *openapiclient.NewGroupUpdateRequest(*openapiclient.NewGroupUpdate()) // GroupUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupApi.UpdateGroup(context.Background(), groupId).GroupUpdateRequest(groupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupUpdateRequest** | [**GroupUpdateRequest**](GroupUpdateRequest.md) |  | 

### Return type

[**GroupResponse**](GroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

