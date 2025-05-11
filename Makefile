protoc:
	cd proto && protoc -I . \
	-I ../thirdparty \
	-I $(shell go list -f '{{ .Dir }}' -m google.golang.org/genproto) \
	-I $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2) \
	--go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../protogen/golang --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=../protogen/golang/swagger \
	--openapiv2_opt=logtostderr=true \
	--openapiv2_opt=allow_merge=true \
	--openapiv2_opt=merge_file_name=swagger \
	**/*.proto