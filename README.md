# mosha-service-common

Common code used by all mosha services

## Protos/gRPC

All service protos definitions are located in `proto` directory. The communication between services is done using gRPC. 
To regenerate the gRPC code, run:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  protos/author_service/author.proto
```