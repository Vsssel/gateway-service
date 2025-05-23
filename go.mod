module gateway

go 1.23.0

toolchain go1.23.4

require google.golang.org/grpc v1.72.0 // для gRPC

require (
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250428153025-10db94c68c34 // indirect
	google.golang.org/protobuf v1.36.6 // indirect; для protobuf
)

require github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3

require (
	google.golang.org/genproto v0.0.0-20250505200425-f936aa4a68b2 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250428153025-10db94c68c34 // indirect
)
