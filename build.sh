# Copyright 2022 EasyStack, Inc.

rm -rf openapi

docker run --rm \
  -v $(pwd)/keystone.yaml:/tmp/keystone.yaml \
  -v $(pwd):/tmp/keystone openapitools/openapi-generator-cli \
  generate -i /tmp/keystone.yaml -g go -o /tmp/keystone/openapi

rm -rf ./openapi/go.mod ./openapi/go.sum ./openapi/api