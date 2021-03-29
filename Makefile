up:
	docker-compose up -d
down:
	docker-compose down
build:
	docker-compose build --no-cache --force-rm
protoc:
	protoc -I=./protos accounts.proto \
 		--go_out=plugins="grpc:./accounts" \
 		--go_out=plugins="grpc:./reviews" \
 		--go_out=plugins="grpc:./database"
	protoc -I=./protos products.proto \
 		--go_out=plugins="grpc:./products" \
 		--go_out=plugins="grpc:./reviews" \
 		--go_out=plugins="grpc:./database"
	protoc -I=./protos reviews.proto \
    	--go_out=plugins="grpc:./reviews" \
        --go_out=plugins="grpc:./database"
