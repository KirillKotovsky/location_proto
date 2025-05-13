PROTO_DIR=proto
OUT_DIR=.

generate:
	mkdir -p $(OUT_DIR)
	protoc \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/location.proto