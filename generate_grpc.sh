#! /bin/bash
# protoc --go_out=. --go-grpc_out=. api/proto/ping.proto

protoc \
  --go_opt=module=github.com/anchamber/genetics-system \
  --go_out=../genetics-system \
  --go-grpc_opt=module=github.com/anchamber/genetics-system \
  --go-grpc_out=../genetics-system \
  --proto_path=proto \
  system.proto

protoc \
  --go_opt=module=github.com/anchamber/genetics-api \
  --go_out=. \
  --go-grpc_opt=module=github.com/anchamber/genetics-api \
  --go-grpc_out=. \
  --proto_path=proto \
  api.proto