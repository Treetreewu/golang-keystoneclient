# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserApi.md#CreateUser) | **Post** /v3/users | 
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /v3/users/{user_id} | 
[**GetUser**](UserApi.md#GetUser) | **Get** /v3/users/{user_id} | 
[**ListUserGroups**](UserApi.md#ListUserGroups) | **Get** /v3/users/{user_id}/groups | 
[**ListUserProjects**](UserApi.md#ListUserProjects) | **Get** /v3/users/{user_id}/projects | 
[**ListUsers**](UserApi.md#ListUsers) | **Get** /v3/users | 
[**UpdateUser**](UserApi.md#UpdateUser) | **Patch** /v3/users/{user_id} | 
[**UpdateUserPassword**](UserApi.md#UpdateUserPassword) | **Post** /v3/users/{user_id}/password | 



## CreateUser

> UserResponse CreateUser(ctx).UserCreateRequest(userCreateRequest).Execute()





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
    userCreateRequest := *openapiclient.NewUserCreateRequest(*openapiclient.NewUserCreate("Name_example")) // UserCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.CreateUser(context.Background()).UserCreateRequest(userCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateRequest** | [**UserCreateRequest**](UserCreateRequest.md) |  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()





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
    userId := "userId_example" // string | The user id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetUser

> UserGetResponse GetUser(ctx, userId).Execute()





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
    userId := "userId_example" // string | The user id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserGetResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserGetResponse**](UserGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserGroups

> GroupListResponse ListUserGroups(ctx, userId).Execute()





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
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.ListUserGroups(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserGroups`: GroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListUserProjects

> ProjectListResponse ListUserProjects(ctx, userId).Execute()





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
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.ListUserProjects(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUserProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserProjects`: ProjectListResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUserProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListUsers

> UserListResponse ListUsers(ctx).DomainId(domainId).Name(name).Enabled(enabled).PasswordExpiresAt(passwordExpiresAt).IdpId(idpId).ProtocolId(protocolId).UniqueId(uniqueId).Execute()





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
    domainId := "domainId_example" // string | Filters the response by a domain ID. (optional)
    name := "name_example" // string | Filters the response by name. (optional)
    enabled := true // bool | If set to true, then only enabled resources will be returned. Any value other than 0 (including no value) will be interpreted as true. (optional)
    passwordExpiresAt := "lt:2016-12-08T22:02:00Z" // string | Filter results based on which user passwords have expired. The query should include an `operator` and a `timestamp` with a colon (`:`) separating the two, for example:  ```password_expires_at={operator}:{timestamp}```  * Valid operators are: `lt`, `lte`, `gt`, `gte`, `eq`, and `neq`   * lt: expiration time lower than the timestamp   * lte: expiration time lower than or equal to the timestamp   * gt: expiration time higher than the timestamp   * gte: expiration time higher than or equal to the timestamp   * eq: expiration time equal to the timestamp   * neq: expiration time not equal to the timestamp * Valid timestamps are of the form: `YYYY-MM-DDTHH:mm:ssZ`  (optional)
    idpId := "idpId_example" // string | Filters the response by an identity provider ID. (optional)
    protocolId := "protocolId_example" // string | Filters the response by a protocol id. (optional)
    uniqueId := "uniqueId_example" // string | Filters the response by a unique id. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.ListUsers(context.Background()).DomainId(domainId).Name(name).Enabled(enabled).PasswordExpiresAt(passwordExpiresAt).IdpId(idpId).ProtocolId(protocolId).UniqueId(uniqueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UserListResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **string** | Filters the response by a domain ID. | 
 **name** | **string** | Filters the response by name. | 
 **enabled** | **bool** | If set to true, then only enabled resources will be returned. Any value other than 0 (including no value) will be interpreted as true. | 
 **passwordExpiresAt** | **string** | Filter results based on which user passwords have expired. The query should include an &#x60;operator&#x60; and a &#x60;timestamp&#x60; with a colon (&#x60;:&#x60;) separating the two, for example:  &#x60;&#x60;&#x60;password_expires_at&#x3D;{operator}:{timestamp}&#x60;&#x60;&#x60;  * Valid operators are: &#x60;lt&#x60;, &#x60;lte&#x60;, &#x60;gt&#x60;, &#x60;gte&#x60;, &#x60;eq&#x60;, and &#x60;neq&#x60;   * lt: expiration time lower than the timestamp   * lte: expiration time lower than or equal to the timestamp   * gt: expiration time higher than the timestamp   * gte: expiration time higher than or equal to the timestamp   * eq: expiration time equal to the timestamp   * neq: expiration time not equal to the timestamp * Valid timestamps are of the form: &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;  | 
 **idpId** | **string** | Filters the response by an identity provider ID. | 
 **protocolId** | **string** | Filters the response by a protocol id. | 
 **uniqueId** | **string** | Filters the response by a unique id. | 

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


## UpdateUser

> UserResponse UpdateUser(ctx, userId).UserUpdateRequest(userUpdateRequest).Execute()





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
    userId := "userId_example" // string | The user id.
    userUpdateRequest := *openapiclient.NewUserUpdateRequest(*openapiclient.NewUserUpdate()) // UserUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UpdateUser(context.Background(), userId).UserUpdateRequest(userUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateRequest** | [**UserUpdateRequest**](UserUpdateRequest.md) |  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUserPassword(ctx, userId).UserPasswordUpdateRequest(userPasswordUpdateRequest).Execute()





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
    userId := "userId_example" // string | The user ID.
    userPasswordUpdateRequest := *openapiclient.NewUserPasswordUpdateRequest(*openapiclient.NewUserPasswordUpdate("OriginalPassword_example", "Password_example")) // UserPasswordUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UpdateUserPassword(context.Background(), userId).UserPasswordUpdateRequest(userPasswordUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPasswordUpdateRequest** | [**UserPasswordUpdateRequest**](UserPasswordUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

