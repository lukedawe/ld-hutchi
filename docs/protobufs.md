# Protobufs

For this API, Protobufs is used to communicate between the website and the API.

The common messages can be found in `/common/*.proto`.

The protobuffer files can be generated for the api with the command:

```bash
protoc --go_out=./ common/dog_operations.proto common/dogs.proto
```
