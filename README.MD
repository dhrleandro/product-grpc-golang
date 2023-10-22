# Observação

Este era um dos desafios da Imersão Fullcyle. Infelizmente não pude terminar a tempo de entregá-lo no dia 20/10/2023 mas para aprimorar meus conhecimentos em Go decidi conclui-lo mesmo assim.

Eu usei um arquivo protobuf pronto e não alterei, por esse motivo o servidor gRPC está com o pacote github.com.codeedu.codepix

No entando todo o resto do código foi desenvolvido por min (dhrleandro). Tentei implementar um pouco da ideia do Domain Driven Design e Clean Architecture. Minhas principais inspirações:

https://github.com/stemmlerjs/ddd-forum
https://github.com/devfullcycle/imersao15

# Testar gRPC

Protobuf de exemplo obtido de:
https://gist.github.com/argentinaluiz/8a2e1a7a32da3d1b107d88ec9f1007b2

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

-------------------

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