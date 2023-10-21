package main

import (
	"fmt"
	"os"

	"github.com/dhrleandro/product-grpc-golang/infrastructure/database"
)

func main() {
	database.ConnectDB(os.Getenv("env"))
	fmt.Println("Ok")
}
