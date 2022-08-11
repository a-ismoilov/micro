package postgres

import (
	"context"
	"database/sql"
	"github.com/akbarshoh/microOLX/proto"
)

type PostgresRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) PostgresRepository {
	return PostgresRepository{
		DB: db,
	}
}

func (p PostgresRepository) GetOrder(ctx context.Context, request proto.Request) (proto.Order, error) {
	return proto.Order{}, nil
}

func (p PostgresRepository) OrderList(ctx context.Context, request proto.Request) (proto.Orders, error) {
	return proto.Orders{}, nil
}

func (p PostgresRepository) Payment(ctx context.Context, request proto.Request) error {
	return nil
}

func (p PostgresRepository) Choose(ctx context.Context, request proto.Request) error {
	return nil
}

func (p PostgresRepository) MealList(ctx context.Context, request proto.Request) (proto.Meals, error) {
	return proto.Meals{}, nil
}

//GetOrder(ctx context.Context, request proto.Request) (proto.Order, error)
//Payment(ctx context.Context, request proto.Request) error
//Choose(ctx context.Context, request proto.Request) error
//MealList(ctx context.Context, request proto.Request) (proto.Meals, error)
//OrderList(ctx context.Context, request proto.Request) (proto.Orders, error)
