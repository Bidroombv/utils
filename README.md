# Utilities

## gRPC logging interceptor for Zerolog

Automagically logs gRPC calls via global instance of Zerolog. In case of error response, context or error is also logged.

### Usage

Interceptor is implemented as `ServerOption`. Find the place where gRPC server is created and simply add call to
`logging.UnaryInterceptor` imported from this repository.

```go
server := grpc.NewServer(
			logging.UnaryInterceptor(),
		)
```

### Example output

```
2021-03-01T11:10:59Z ERR pkg/logger/grpc.go:48 > gRPC error="rpc error: code = NotFound desc = Id: " code=NotFound details=[] duration=73.929187 ip=127.0.0.1:51444 method=GetPropertyDetails service=inventory.v1.Inventory
2021-03-01T11:11:15Z INF pkg/logger/grpc.go:52 > gRPC duration=0.038067 ip=127.0.0.1:51444 method=ListCodes service=inventory.v1.InventorySearch
```

## Contributing

Install **golangci-lint** to run build locally:

```
brew install golangci/tap/golangci-lint
```

## Errors

Go 1.13 style error wrappers. They help trace where the error originated.
