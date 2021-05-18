### Prereqs

https://grpc.io/docs/languages/go/quickstart/

### Compile

```shell
protoc proto/product_info.proto --go_out=. --go-grpc_out=.
```