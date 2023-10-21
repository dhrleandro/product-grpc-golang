package main

import (
	"fmt"
	"os"

	pb "github.com/dhrleandro/product-grpc-golang/application/grpc/protofiles"
	"github.com/dhrleandro/product-grpc-golang/application/usecase"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database/repository"
)

func main() {
	db := database.ConnectDB(os.Getenv("env"))
	r := &repository.ProductRepositoryGORM{Db: db}

	fmt.Println("Executando UseCase CreateProduct")
	usecase := &usecase.ProductUseCase{r}
	res, err := usecase.CreateProduct("Lápis de Cor", "feito de madeira", 258)
	if err != nil {
		fmt.Printf("%w", err)
	}
	fmt.Println("Produto inserido no banco:")
	fmt.Printf("%v", res)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("Fetch")
	plist, _ := usecase.FindProduct("lápis")
	for i := range plist {
		fmt.Println(plist[i])
	}

	a := pb.CreateProductRequest{}
	fmt.Println(a)

	fmt.Println("Ok")
}
