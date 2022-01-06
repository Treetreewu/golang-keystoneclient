# \RoleAssignmentApi

All URIs are relative to *http://keystone-api.openstack.svc.cluster.local*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignGroupDomainRole**](RoleAssignmentApi.md#AssignGroupDomainRole) | **Put** /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id} | 
[**AssignGroupProjectRole**](RoleAssignmentApi.md#AssignGroupProjectRole) | **Put** /v3/projects/{project_id}/groups/{group_id}/roles/{role_id} | 
[**AssignUserDomainRole**](RoleAssignmentApi.md#AssignUserDomainRole) | **Put** /v3/domains/{domain_id}/users/{user_id}/roles/{role_id} | 
[**AssignUserProjectRole**](RoleAssignmentApi.md#AssignUserProjectRole) | **Put** /v3/projects/{project_id}/users/{user_id}/roles/{role_id} | 
[**CheckGroupDomainRole**](RoleAssignmentApi.md#CheckGroupDomainRole) | **Head** /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id} | 
[**CheckGroupProjectRole**](RoleAssignmentApi.md#CheckGroupProjectRole) | **Head** /v3/projects/{project_id}/groups/{group_id}/roles/{role_id} | 
[**CheckUserDomainRole**](RoleAssignmentApi.md#CheckUserDomainRole) | **Head** /v3/domains/{domain_id}/users/{user_id}/roles/{role_id} | 
[**CheckUserProjectRole**](RoleAssignmentApi.md#CheckUserProjectRole) | **Head** /v3/projects/{project_id}/users/{user_id}/roles/{role_id} | 
[**ListGroupDomainRoles**](RoleAssignmentApi.md#ListGroupDomainRoles) | **Get** /v3/domains/{domain_id}/groups/{group_id}/roles | 
[**ListGroupProjectRoles**](RoleAssignmentApi.md#ListGroupProjectRoles) | **Get** /v3/projects/{project_id}/groups/{group_id}/roles | 
[**ListRoleAssignments**](RoleAssignmentApi.md#ListRoleAssignments) | **Get** /v3/role_assignments | 
[**ListRoleAssignmentsIncludingNames**](RoleAssignmentApi.md#ListRoleAssignmentsIncludingNames) | **Get** /v3/role_assignments#including_names | 
[**ListUserDomainRoles**](RoleAssignmentApi.md#ListUserDomainRoles) | **Get** /v3/domains/{domain_id}/users/{user_id}/roles | 
[**ListUserProjectRoles**](RoleAssignmentApi.md#ListUserProjectRoles) | **Get** /v3/projects/{project_id}/users/{user_id}/roles | 
[**UnassignGroupDomainRole**](RoleAssignmentApi.md#UnassignGroupDomainRole) | **Delete** /v3/domains/{domain_id}/groups/{group_id}/roles/{role_id} | 
[**UnassignGroupProjectRole**](RoleAssignmentApi.md#UnassignGroupProjectRole) | **Delete** /v3/projects/{project_id}/groups/{group_id}/roles/{role_id} | 
[**UnassignUserDomainRole**](RoleAssignmentApi.md#UnassignUserDomainRole) | **Delete** /v3/domains/{domain_id}/users/{user_id}/roles/{role_id} | 
[**UnassignUserProjectRole**](RoleAssignmentApi.md#UnassignUserProjectRole) | **Delete** /v3/projects/{project_id}/users/{user_id}/roles/{role_id} | 



## AssignGroupDomainRole

> AssignGroupDomainRole(ctx, domainId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.AssignGroupDomainRole(context.Background(), domainId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignGroupDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupDomainRoleRequest struct via the builder pattern


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


## AssignGroupProjectRole

> AssignGroupProjectRole(ctx, projectId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.AssignGroupProjectRole(context.Background(), projectId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignGroupProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignGroupProjectRoleRequest struct via the builder pattern


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


## AssignUserDomainRole

> AssignUserDomainRole(ctx, domainId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.AssignUserDomainRole(context.Background(), domainId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignUserDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserDomainRoleRequest struct via the builder pattern


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


## AssignUserProjectRole

> AssignUserProjectRole(ctx, projectId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.AssignUserProjectRole(context.Background(), projectId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.AssignUserProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUserProjectRoleRequest struct via the builder pattern


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


## CheckGroupDomainRole

> CheckGroupDomainRole(ctx, domainId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.CheckGroupDomainRole(context.Background(), domainId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.CheckGroupDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckGroupDomainRoleRequest struct via the builder pattern


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


## CheckGroupProjectRole

> CheckGroupProjectRole(ctx, projectId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.CheckGroupProjectRole(context.Background(), projectId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.CheckGroupProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckGroupProjectRoleRequest struct via the builder pattern


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


## CheckUserDomainRole

> CheckUserDomainRole(ctx, domainId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.CheckUserDomainRole(context.Background(), domainId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.CheckUserDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserDomainRoleRequest struct via the builder pattern


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


## CheckUserProjectRole

> CheckUserProjectRole(ctx, projectId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.CheckUserProjectRole(context.Background(), projectId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.CheckUserProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserProjectRoleRequest struct via the builder pattern


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


## ListGroupDomainRoles

> RoleListResponse ListGroupDomainRoles(ctx, domainId, groupId).Execute()





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
    groupId := "groupId_example" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListGroupDomainRoles(context.Background(), domainId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListGroupDomainRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupDomainRoles`: RoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListGroupDomainRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupDomainRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupProjectRoles

> RoleListResponse ListGroupProjectRoles(ctx, projectId, groupId).Execute()





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
    groupId := "groupId_example" // string | The group ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListGroupProjectRoles(context.Background(), projectId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListGroupProjectRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupProjectRoles`: RoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListGroupProjectRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupProjectRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleAssignments

> RoleAssignmentListResponse ListRoleAssignments(ctx).Effective(effective).IncludeSubtree(includeSubtree).GroupId(groupId).RoleId(roleId).ScopeDomainId(scopeDomainId).ScopeProjectId(scopeProjectId).UserId(userId).Execute()





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
    effective := "effective_example" // string | Returns the effective assignments, including any assignments gained by virtue of group membership. key-only (no value required) (optional)
    includeSubtree := int32(56) // int32 | If set to true, then relevant assignments in the project hierarchy below the project specified in the `scope.project_id` query parameter are also included in the response. Any value other than `0` (including no value) for `include_subtree` will be interpreted as true.  (optional) (default to 0)
    groupId := "groupId_example" // string | Filters the response by a group ID. (optional)
    roleId := "roleId_example" // string | Filters the response by a role ID. (optional)
    scopeDomainId := "scopeDomainId_example" // string | Filters the response by a domain ID. (optional)
    scopeProjectId := "scopeProjectId_example" // string | Filters the response by a project ID. (optional)
    userId := "userId_example" // string | Filters the response by a user ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListRoleAssignments(context.Background()).Effective(effective).IncludeSubtree(includeSubtree).GroupId(groupId).RoleId(roleId).ScopeDomainId(scopeDomainId).ScopeProjectId(scopeProjectId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleAssignments`: RoleAssignmentListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **effective** | **string** | Returns the effective assignments, including any assignments gained by virtue of group membership. key-only (no value required) | 
 **includeSubtree** | **int32** | If set to true, then relevant assignments in the project hierarchy below the project specified in the &#x60;scope.project_id&#x60; query parameter are also included in the response. Any value other than &#x60;0&#x60; (including no value) for &#x60;include_subtree&#x60; will be interpreted as true.  | [default to 0]
 **groupId** | **string** | Filters the response by a group ID. | 
 **roleId** | **string** | Filters the response by a role ID. | 
 **scopeDomainId** | **string** | Filters the response by a domain ID. | 
 **scopeProjectId** | **string** | Filters the response by a project ID. | 
 **userId** | **string** | Filters the response by a user ID. | 

### Return type

[**RoleAssignmentListResponse**](RoleAssignmentListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleAssignmentsIncludingNames

> RoleAssignmentListResponseWithNames ListRoleAssignmentsIncludingNames(ctx).Effective(effective).IncludeNames(includeNames).IncludeSubtree(includeSubtree).GroupId(groupId).RoleId(roleId).ScopeDomainId(scopeDomainId).ScopeProjectId(scopeProjectId).UserId(userId).Execute()





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
    effective := true // bool | Returns the effective assignments, including any assignments gained by virtue of group membership. key-only (no value required) (optional)
    includeNames := int32(56) // int32 | If set to true, then the names of any entities returned will be include as well as their IDs. Any value other than 0 (including no value) will be interpreted as true.  New in version 3.6  (optional) (default to 1)
    includeSubtree := int32(56) // int32 | If set to true, then relevant assignments in the project hierarchy below the project specified in the `scope.project_id` query parameter are also included in the response. If set to `0`, . Any value other than `0` (including no value) for `include_subtree` will be interpreted as true.  New in version 3.6  (optional) (default to 0)
    groupId := "groupId_example" // string | Filters the response by a group ID. (optional)
    roleId := "roleId_example" // string | Filters the response by a role ID. (optional)
    scopeDomainId := "scopeDomainId_example" // string | Filters the response by a domain ID. (optional)
    scopeProjectId := "scopeProjectId_example" // string | Filters the response by a project ID. (optional)
    userId := "userId_example" // string | Filters the response by a user ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListRoleAssignmentsIncludingNames(context.Background()).Effective(effective).IncludeNames(includeNames).IncludeSubtree(includeSubtree).GroupId(groupId).RoleId(roleId).ScopeDomainId(scopeDomainId).ScopeProjectId(scopeProjectId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListRoleAssignmentsIncludingNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleAssignmentsIncludingNames`: RoleAssignmentListResponseWithNames
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListRoleAssignmentsIncludingNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleAssignmentsIncludingNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **effective** | **bool** | Returns the effective assignments, including any assignments gained by virtue of group membership. key-only (no value required) | 
 **includeNames** | **int32** | If set to true, then the names of any entities returned will be include as well as their IDs. Any value other than 0 (including no value) will be interpreted as true.  New in version 3.6  | [default to 1]
 **includeSubtree** | **int32** | If set to true, then relevant assignments in the project hierarchy below the project specified in the &#x60;scope.project_id&#x60; query parameter are also included in the response. If set to &#x60;0&#x60;, . Any value other than &#x60;0&#x60; (including no value) for &#x60;include_subtree&#x60; will be interpreted as true.  New in version 3.6  | [default to 0]
 **groupId** | **string** | Filters the response by a group ID. | 
 **roleId** | **string** | Filters the response by a role ID. | 
 **scopeDomainId** | **string** | Filters the response by a domain ID. | 
 **scopeProjectId** | **string** | Filters the response by a project ID. | 
 **userId** | **string** | Filters the response by a user ID. | 

### Return type

[**RoleAssignmentListResponseWithNames**](RoleAssignmentListResponseWithNames.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserDomainRoles

> RoleListResponse ListUserDomainRoles(ctx, domainId, userId).Execute()





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
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListUserDomainRoles(context.Background(), domainId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListUserDomainRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserDomainRoles`: RoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListUserDomainRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserDomainRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserProjectRoles

> RoleListResponse ListUserProjectRoles(ctx, projectId, userId).Execute()





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
    userId := "userId_example" // string | The user ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.ListUserProjectRoles(context.Background(), projectId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.ListUserProjectRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserProjectRoles`: RoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleAssignmentApi.ListUserProjectRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**userId** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserProjectRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleListResponse**](RoleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignGroupDomainRole

> UnassignGroupDomainRole(ctx, domainId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.UnassignGroupDomainRole(context.Background(), domainId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignGroupDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignGroupDomainRoleRequest struct via the builder pattern


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


## UnassignGroupProjectRole

> UnassignGroupProjectRole(ctx, projectId, groupId, roleId).Execute()





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
    groupId := "groupId_example" // string | The group ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.UnassignGroupProjectRole(context.Background(), projectId, groupId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignGroupProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**groupId** | **string** | The group ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignGroupProjectRoleRequest struct via the builder pattern


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


## UnassignUserDomainRole

> UnassignUserDomainRole(ctx, domainId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.UnassignUserDomainRole(context.Background(), domainId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignUserDomainRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | The domain ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUserDomainRoleRequest struct via the builder pattern


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


## UnassignUserProjectRole

> UnassignUserProjectRole(ctx, projectId, userId, roleId).Execute()





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
    userId := "userId_example" // string | The user ID.
    roleId := "roleId_example" // string | The role ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoleAssignmentApi.UnassignUserProjectRole(context.Background(), projectId, userId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleAssignmentApi.UnassignUserProjectRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | The project ID. | 
**userId** | **string** | The user ID. | 
**roleId** | **string** | The role ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUserProjectRoleRequest struct via the builder pattern


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

