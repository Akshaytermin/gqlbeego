package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Akshaytermin/gqlbeego/graph/generated"
	"github.com/Akshaytermin/gqlbeego/graph/model"
	"github.com/Akshaytermin/gqlbeego/repository"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	var resp model.Product

	product := model.Product{
		Name:  input.Name,
		Price: input.Price,
	}

	//Create by passing the pointer to the product
	resp, err := repository.Create(ctx, product)
	if err != nil {
		log.Fatal(err)
		return &resp, err
	}

	return &resp, nil

}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id *int, input *model.NewProduct) (*model.Product, error) {
	var product model.Product

	query := model.Product{
		Name:  input.Name,
		Price: input.Price,
	}

	product, err := repository.Update(ctx, *id, query)
	if err != nil {
		log.Fatal(err)
		return &product, err
	}

	return &product, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id *int) ([]*model.Product, error) {

	var resp []*model.Product

	_, err := repository.Delete(ctx, *id)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil

}

//Finding all product details
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {

	var products []*model.Product

	products, err := repository.Find(ctx)
	if err != nil {
		log.Fatal(err)
		return products, err
	}

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
