package postgres

import (
	"context"
	"database/sql"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"time"
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
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	o := orderproto.Order{}
	if err := p.DB.QueryRowContext(c,
		"select * from orders where id=$1", request.Id,
	).Scan(&o.Id, &o.UserId, &o.OpenedAt, &o.ClosedAt, &o.Price, &o.MealList); err != nil {
		return orderproto.Order{}, err
	}
	return o, nil
}

func (p PostgresRepository) OrderList(ctx context.Context, request orderproto.Request) (orderproto.Orders, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	os := orderproto.Orders{}
	o := orderproto.Order{}
	rows, err := p.DB.QueryContext(c,
		"select * from orders where closed_at != opened_at",
	)
	if err != nil {
		return orderproto.Orders{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&o.Id, &o.UserId, &o.OpenedAt, &o.ClosedAt, &o.Price, &o.MealList); err != nil {
			return orderproto.Orders{}, err
		}
		os.List = append(os.List, &o)
	}
	return os, nil
}

func (p PostgresRepository) Payment(ctx context.Context, request orderproto.Request) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if _, err := p.DB.ExecContext(c, "update orders set closed_at=current_timestamp where id=$1;", request.Id); err != nil {
		return err
	}
	return nil
}

func (p PostgresRepository) Choose(ctx context.Context, request orderproto.MealChoice) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	if request.OrderId != 0 {
		p.DB.ExecContext(c, "insert into users values($1)", request.Id)
		p.DB.ExecContext(c, "insert into orders (user_id) values ($1)", request.Id)
	}
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	defer cancel()
	if _, err := tx.ExecContext(c, "update orders set meals = array_append(meals, $1) where id=$2", request.Id, request.OrderId); err != nil {
		return err
	}
	if _, err := tx.ExecContext(c, "with p as ( select price from meal where id=$1) update orders set price=price+p where id=$2", request.Id, request.OrderId); err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (p PostgresRepository) MealList(ctx context.Context, request orderproto.Request) (orderproto.Meals, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	m := orderproto.Meal{}
	ms := orderproto.Meals{}
	rows, err := p.DB.QueryContext(c, "select * from meals")
	if err != nil {
		return orderproto.Meals{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&m.Id, &m.Name, &m.Price); err != nil {
			return orderproto.Meals{}, err
		}
		ms.List = append(ms.List, &m)
	}
	return ms, nil
}

func (p PostgresRepository) Cancel(ctx context.Context, request *orderproto.Request) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if _, err := p.DB.ExecContext(c, "delete from orders where id=$1", request.Id); err != nil {
		return err
	}
	return nil
}
func (p PostgresRepository) NewMeal(ctx context.Context, request *orderproto.Meal) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if _, err := p.DB.ExecContext(c, "insert into meals (name, price) values ($1,$2)", request.Name, request.Price); err != nil {
		return err
	}
	return nil
}
func (p PostgresRepository) UpdateMeal(ctx context.Context, request *orderproto.Meal) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if _, err := p.DB.ExecContext(c, "update meals set name=$1 and price=$2 where id=$3", request.Name, request.Price, request.Id); err != nil {
		return err
	}
	return nil
}
