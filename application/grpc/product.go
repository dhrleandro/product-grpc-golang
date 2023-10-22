package productgrpc

import (
	"context"

	"github.com/dhrleandro/product-grpc-golang/application/grpc/pb"
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
		Product: &pb.Product{Id: product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price.ToBrazilianReal(),
		},
	}, nil
}

func (svc *productServiceServer) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	prs, err := svc.productUseCase.FindProduct(in.Term)
	if err != nil {
		return nil, err
	}

	products := []*pb.Product{}
	for i := range prs {
		newProduct := &pb.Product{
			Id:          prs[i].ID,
			Name:        prs[i].Name,
			Description: prs[i].Description,
			Price:       prs[i].Price.ToBrazilianReal(),
		}
		products = append(products, newProduct)
	}

	return &pb.FindProductsResponse{Products: products}, nil
}

func NewService(productUseCase *usecase.ProductUseCase) *productServiceServer {
	return &productServiceServer{productUseCase: productUseCase}
}
