package postgres

import (
	"context"
	"database/sql"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostgresRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) PostgresRepository {
	return PostgresRepository{
		DB: db,
	}
}

func (p PostgresRepository) GetOrder(ctx context.Context, request orderproto.Request) (orderproto.Order, error) {
	return orderproto.Order{}, nil
}

func (p PostgresRepository) OrderList(ctx context.Context, request orderproto.Request) (orderproto.Orders, error) {
	return orderproto.Orders{}, nil
}

func (p PostgresRepository) Payment(ctx context.Context, request orderproto.Request) error {
	return nil
}

func (p PostgresRepository) Choose(ctx context.Context, request orderproto.Request) error {
	return nil
}

func (p PostgresRepository) MealList(ctx context.Context, request orderproto.Request) (orderproto.Meals, error) {
	return orderproto.Meals{}, nil
}

func (p PostgresRepository) Cancel(context.Context, *orderproto.Request) error {
	return status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (p PostgresRepository) NewMeal(context.Context, *orderproto.Meal) error {
	return status.Errorf(codes.Unimplemented, "method NewMeal not implemented")
}
func (p PostgresRepository) NewOrder(context.Context, *orderproto.Order) error {
	return status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}

//GetOrder(ctx context.Context, request userproto.Request) (userproto.Order, error)
//Payment(ctx context.Context, request userproto.Request) error
//Choose(ctx context.Context, request userproto.Request) error
//MealList(ctx context.Context, request userproto.Request) (userproto.Meals, error)
//OrderList(ctx context.Context, request userproto.Request) (userproto.Orders, error)
