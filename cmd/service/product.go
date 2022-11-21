package service

import (
	"context"
	"grpc/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProducts(context.Context, *pb.Empty) (*pb.Products, error) {

	var products []*pb.Product

	rows, err := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product pb.Product
		var category pb.Category

		if err := rows.Scan(&product.Id, &product.Name,
			&product.Price, &product.Stock, &category.Id,
			&category.Name); err != nil {
			log.Fatalf("Gagal mengambil row data %v", err.Error())
		}

		product.Category = &category
		products = append(products, &product)
	}

	res := &pb.Products{
		Pagination: &pb.Pagination{
			Total:       2,
			PerPage:     1,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: products,
	}

	return res, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *pb.Id) (*pb.Product, error) {

	rows := p.DB.Table(`public.products`).
		Joins(`LEFT JOIN public.categories ON categories.id = products.category_id`).
		Select(`products.id, products."name", products.price, products.stock, categories.id, categories.name`).
		Where(`products.id = ?`, id.GetId()).Row()

	var product pb.Product
	var category pb.Category

	if err := rows.Scan(&product.Id, &product.Name,
		&product.Price, &product.Stock, &category.Id,
		&category.Name); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	product.Category = &category

	return &product, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Id, error) {

	var Response pb.Id

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		category := pb.Category{
			Id:   1,
			Name: req.GetCategory().GetName(),
		}

		product := struct {
			Id          uint64
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Id:          req.GetId(),
			Name:        req.GetName(),
			Price:       req.GetPrice(),
			Stock:       req.GetStock(),
			Category_id: category.GetId(),
		}

		if err := tx.Table(`public.products`).Create(&product).Error; err != nil {
			return err
		}

		Response.Id = uint32(product.Id)

		return nil
	})

	if err != nil {
		return nil, status.Error(
			codes.Internal, err.Error(),
		)
	}

	return &Response, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Status, error) {
	
	var Response pb.Status

	err := p.DB.Transaction(func(tx *gorm.DB) error {

		category := pb.Category{
			Id:   1,
			Name: req.GetCategory().GetName(),
		}

		product := struct {
			Id          uint64
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Id:          req.GetId(),
			Name:        req.GetName(),
			Price:       req.GetPrice(),
			Stock:       req.GetStock(),
			Category_id: category.GetId(),
		}

		if err := tx.Table(`public.products`).Where(`products.id = ?`, product.Id).Updates(&product).Error; err != nil {
			return err
		}

		Response.Status = 1

		return nil
	})

	if err != nil {
		return nil, status.Error(
			codes.Internal, err.Error(),
		)
	}

	return &Response, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id *pb.Id) (*pb.Status, error) {
	
	var Response pb.Status

	if err := p.DB.Table(`public.products`).Where(`products.id = ?`, id.GetId()).Delete(nil).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	Response.Status = 1

	return &Response, nil
}