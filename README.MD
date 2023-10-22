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

# Testar gRPC

Para testar o servidor gRPC utilizei a ferramenta
https://github.com/fullstorydev/grpcui

```sh
grpcui -plaintext localhost:9000
```

Antes disso é necessário habilitar o reflection no servidor gRPC para permitir que o grpcui obtenha a especificação do protofile.

É bem simples, praticamente instalar um pacote
https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md

e adicionar a linha:
reflection.Register(server)