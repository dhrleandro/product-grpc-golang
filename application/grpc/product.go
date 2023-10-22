package productgrpc

import (
	"context"

	pb "github.com/dhrleandro/product-grpc-golang/application/grpc/protofiles"
	"github.com/dhrleandro/product-grpc-golang/application/usecase"
)

type productServiceServer struct {
	pb.UnimplementedProductServiceServer
	productUseCase *usecase.ProductUseCase
}

func (svc *productServiceServer) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product, err := svc.productUseCase.CreateProduct(in.Name, in.Description, float32(in.Price))
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{Id: product.ID, Name: product.Name, Description: product.Description, Price: float32(product.Price.ToBrazilianReal())},
	}, nil
}

func (svc *productServiceServer) FindProducts(ctx context.Context, c *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products := []*pb.Product{
		{Id: 1, Name: "Carros", Description: "", Price: 0},
		{Id: 2, Name: "Barcos", Description: "", Price: 0},
		{Id: 3, Name: "Motos", Description: "", Price: 0},
	}
	return &pb.FindProductsResponse{Products: products}, nil
}

func NewService(productUseCase *usecase.ProductUseCase) *productServiceServer {
	return &productServiceServer{productUseCase: productUseCase}
}
