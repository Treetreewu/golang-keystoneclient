# Keystone client

## Introduction
The `openapi` folder is generated from `keystone.yaml`(i.e. the OAS) using `openapi-generator-cli`.

Another OAS `token.json` is not used for now because the `openapi-generator-cli v5.3.0` does not support object combinations(`oneOf`, `anyOf` and `allOf`), in which case brings too much complexity.

The `options` parameter in some APIs is not supported. 
## Build
```shell
docker run --rm -v $(pwd)/keystone.yaml:/tmp/keystone.yaml -v $(pwd):/tmp/keystone openapitools/openapi-generator-cli generate -i /tmp/keystone.yaml -g go -o /tmp/keystone/openapi
rm -f ./openapi/go.mod ./openapi/go.sum
```
### Note
For openapi-generator-cli does not support `oneOf` schema, define `OneOfstringprojectResponse` in `model_project_get.go`
```go
type OneOfstringprojectResponse interface{}
````

## Usage
```go
// Build client
config := openapi.NewConfiguration()
client, err := keystone.NewClient(config)
if err != nil {
    panic(err)
}

// Authenticate
domain := "default"
scope := model.NewDomainScope(&domain, nil)
user := "drone"
cred := model.NewPasswordCredential(nil, &user, "password", &domain, nil)
response, err := client.Auth.Authenticate(scope, cred)

if err != nil {
    fmt.Printf("%v\n", response)
    fmt.Printf("%v\n", err)
    panic(err)
}
fmt.Println(client.Token)

// API Call
resp, r, err := client.DomainApi.ListDomains(context.TODO()).Execute()
```

### How to change the server configuration?
#### Default
The client uses the first server (index=0) defined in OAS by default.

#### Change server index
By specifying the server index, client will use corresponding server configuration defined in OAS.
```go
// Build the client
serverIndex := 1
config := openapi.NewConfiguration()
client, err := keystone.NewClientWithServerIndex(config, serverIndex)

// Authentication is omitted.

// API Call
ctx := context.WithValue(context.TODO(), openapi.ContextServerIndex, serverIndex)
resp, r, err := client.DomainApi.ListDomains(ctx).Execute()

```

#### Change manually in code
Just edit the `config` object.
```go
config := openapi.NewConfiguration()
config.Servers = openapi.ServerConfigurations{{
    URL:         "localhost:8000",
    Description: "local test",
    Variables:   nil,
}}
client, err := keystone.NewClient(config)
```