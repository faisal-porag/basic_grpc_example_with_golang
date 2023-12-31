generate_details_protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        details/details.proto


generate_tprotos_protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        tprotos/tprotos.proto


generate_person_protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        person_proto/person_proto.proto


run_server:
	go run main.go


run_client:
	go run client/main.go



private_pem_gen:
	openssl genpkey -algorithm RSA -out private_key.pem

public_pem_gen:
	openssl rsa -pubout -in private_key.pem -out public_key.pem

