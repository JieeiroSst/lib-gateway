gen_proxy:
	protoc -I./googleapis -I. \
	--go_out=./${folder}/gateway --go_opt=paths=source_relative \
	--go-grpc_out=./${folder}/gateway --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./${folder}/gateway --grpc-gateway_opt=paths=source_relative \
	./${folder}/*.proto