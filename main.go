package main

import (
	"fmt"
	"os"

	"github.com/dhrleandro/product-grpc-golang/domain/model"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database"
	"github.com/dhrleandro/product-grpc-golang/infrastructure/database/repository"
)

func main() {
	db := database.ConnectDB(os.Getenv("env"))
	r := repository.ProductRepositoryGORM{Db: db}

	p, _ := model.NewProduct(
		0,
		"Table",
		"Wood",
		584799,
	)
	pr, _ := r.Save(p)
	fmt.Println(p, pr)
	fmt.Println("Ok")
}
