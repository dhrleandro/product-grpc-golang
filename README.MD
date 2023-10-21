Genrating protoc

```
https://aprendagolang.com.br/2023/07/13/implementando-uma-api-com-protobuf-e-grpc/

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         application/grpc/protofiles/product.proto
```

docker build -t product-grpc-golang .

docker run product-grpc-golang
docker exec -it product-grpc-golang bash

https://github.com/docker-archive/toolbox/issues/673
https://github.com/moby/moby/issues/24029
docker run --rm -i -t product-grpc-golang /bin/sh

protoc --go_out=. --go_opt=paths=./application/grpc/protofiles --go-grpc_out=. --go-grpc_opt=paths=./application/grpc/protofiles ./application/grpc/protofiles/product.proto