gen:
	protoc --twirp_out=. --twirplura_out=. --twirp_swagger_out=./swagger --go_out=. protos/color-service/service.proto
	protoc --twirp_out=. --twirplura_out=. --twirp_swagger_out=./swagger --go_out=. protos/number-service/service.proto