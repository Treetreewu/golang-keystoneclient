# Keystone client

## Introduction
The `openapi` folder is generated from `keystone.yaml`(i.e. the OAS) using `openapi-generator-cli`.

Another OAS `token.json` is not used for now because the `openapi-generator-cli v5.3.0` does not support object combinations(`oneOf`, `anyOf` and `allOf`), in which case brings too much complexity.

The `options` parameter in some APIs is not supported. 
## Build
```shell
./build.sh
```

### Note
For openapi-generator-cli does not support `oneOf` schema, define `OneOfstringprojectResponse` in `model_project_get.go`
```go
type OneOfstringprojectResponse interface{}
````

## Usage
### Build client by token
```go
// Build client
client, err := keystone.NewDefaultClientByToken("a_keystone_token")
if err != nil {
    panic(err)
}

// API Call
resp, r, err := client.DomainApi.ListDomains(context.TODO()).Execute()
```
### Build client by password
```go
// Build client
client, err := keystone.NewDefaultClient()
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
```yaml
- description: Service domain
  url: http://keystone-api.openstack.svc.cluster.local
```

#### Change manually in code
Just edit the `config` object.
```go
config := openapi.NewConfiguration()
config.Servers = openapi.ServerConfigurations{{
    URL:         "localhost:8000",
}}
client, err := keystone.NewClient(config)
```