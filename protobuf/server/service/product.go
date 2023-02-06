package service

import "context"

type productService struct {
}

var ProductService = &productService{}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := request.GetProductId()
	return &ProductResponse{ProductStock: stock}, nil
}
