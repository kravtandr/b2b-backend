protogen-api-with-validator:
	protoc 								\
    		-I. 								\
    		-I./third_party 								\
    		-I./third_party/googleapis 								\
    		--go_out=. --go_opt=paths=source_relative 				\
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    		--validate_out=lang=go,paths=source_relative:. \
    		$(path)/api.proto && mockgen -source=$(path)/api_grpc.pb.go -destination=$(path)/mock/api_mock.go

protogen-api-auth-service:
	make protogen-api-with-validator path=pkg/services/auth

protogen-api-chat-service:
	make protogen-api-with-validator path=pkg/services/chat

protogen-api-fastOrder-service:
	make protogen-api-with-validator path=pkg/services/fastOrder

protogen-api-company-service:
	make protogen-api-with-validator path=pkg/services/company

protogen-api-productsCategories-service:
	make protogen-api-with-validator path=pkg/services/productsCategories

all:
	make protogen-api-with-validator path=pkg/services/auth && \
	make protogen-api-with-validator path=pkg/services/chat && \
	make protogen-api-with-validator path=pkg/services/fastOrder && \
	make protogen-api-with-validator path=pkg/services/company && \
	make protogen-api-with-validator path=pkg/services/productsCategories 


prepare-auth_service-env:
	export USER_DB_URL="postgres://tripadvisor:12345@localhost:5432/tripadvisor" && \
			export USER_GRPC_PORT="10123" && export USER_PREFIX_LEN="0"

prepare-websocket_service-env:
	export WEBSOCKET_DB_URL="postgres://tripadvisor:12345@localhost:5432/tripadvisor" && \
			export WEBSOCKET_PORT="5050" && export WEBSOCKET_URL="http://localhost"

prepare-gateway-env:
	export GATEWAY_HTTP_PORT=":8080" && export GATEWAY_AUTH_ENDPOINT="localhost:10123" && export GATEWAY_TRIP_ENDPOINT="localhost:6666"

run-auth:
	make prepare-auth_service-env && go run cmd/auth_service/main.go
run-gateway:
	make prepare-gateway-env && go run cmd/gateway/main.go
