package service

import (
	"context"
	"grpc/cmd/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProduct(ctx context.Context, id *pb.Id) (*pb.ResponseProduct, error) {

	row := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).
		Where(`products.id = ?`, id.GetId()).Row()

	var product pb.Product
	var category pb.Category

	if err := row.Scan(&product.Id, &product.Name,
		&product.Price, &product.Stock, &category.Id,
		&category.Name); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	product.Category = &category

	response := &pb.ResponseProduct{
		Status:       "000",
		Description:  "Success",
		Detail:       "Get Data Product Successfully",
		ResponseData: &product,
	}

	return response, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, req *pb.RequestProduct) (*pb.ResponseProduct, error) {

	var product pb.Product
	var category pb.Category

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		category := pb.Category{
			Id:   1,
			Name: req.GetCategory().GetName(),
		}

		product := struct {
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Name:        req.GetName(),
			Price:       req.GetPrice(),
			Stock:       req.GetStock(),
			Category_id: category.GetId(),
		}

		if err := tx.Table(`public.products`).Create(&product).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(
			codes.Internal, err.Error(),
		)
	}

	row := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).
		Where(`products.name = ?`, req.GetName()).Row()

	if err := row.Scan(&product.Id, &product.Name,
		&product.Price, &product.Stock, &category.Id,
		&category.Name); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	product.Category = &category

	response := &pb.ResponseProduct{
		Status:       "000",
		Description:  "Success",
		Detail:       "Create Data Product Successfully",
		ResponseData: &product,
	}

	return response, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *pb.RequestProduct) (*pb.ResponseProduct, error) {

	var product pb.Product
	var category pb.Category

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		category := pb.Category{
			Id:   1,
			Name: req.GetCategory().GetName(),
		}

		product := struct {
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Name:        req.GetName(),
			Price:       req.GetPrice(),
			Stock:       req.GetStock(),
			Category_id: category.GetId(),
		}

		if err := tx.Table(`public.products`).Where(`products.name = ?`, product.Name).Updates(&product).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(
			codes.Internal, err.Error(),
		)
	}

	row := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).
		Where(`products.name = ?`, req.GetName()).Row()

	if err := row.Scan(&product.Id, &product.Name,
		&product.Price, &product.Stock, &category.Id,
		&category.Name); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	product.Category = &category

	response := &pb.ResponseProduct{
		Status:       "000",
		Description:  "Success",
		Detail:       "Update Data Product Successfully",
		ResponseData: &product,
	}

	return response, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id *pb.Id) (*pb.ResponseProduct, error) {

	var product pb.Product
	var category pb.Category

	row := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).
		Where(`products.id = ?`, id.GetId()).Row()

	if err := row.Scan(&product.Id, &product.Name,
		&product.Price, &product.Stock, &category.Id,
		&category.Name); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	product.Category = &category

	response := &pb.ResponseProduct{
		Status:       "000",
		Description:  "Success",
		Detail:       "Delete Data Product Successfully",
		ResponseData: &product,
	}

	if err := p.DB.Table(`public.products`).Where(`products.id = ?`, id.GetId()).Delete(nil).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}
