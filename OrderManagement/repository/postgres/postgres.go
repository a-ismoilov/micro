package postgres

import (
	"context"
	"database/sql"
	"github.com/akbarshoh/microOLX/protos/orderproto"
	"log"
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
	var (
		cl, ol time.Time
		ids    []uint8
	)
	if err := p.DB.QueryRowContext(c,
		"select * from orders where id=$1", request.Id,
	).Scan(&o.Id, &o.UserId, &ol, &cl, &o.Price, &ids); err != nil {
		return orderproto.Order{}, err
	}
	o.OpenedAt = ol.String()
	o.ClosedAt = cl.String()
	for i := range ids {
		o.MealList = append(o.MealList, int32(ids[i]))
	}
	log.Println(o)
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
	var (
		ca, oa time.Time
		ids    []uint8
	)
	ors := make([]orderproto.Order, 0)
	for rows.Next() {
		if err := rows.Scan(&o.Id, &o.UserId, &oa, &ca, &o.Price, &ids); err != nil {
			return orderproto.Orders{}, err
		}
		o.OpenedAt = oa.String()
		o.ClosedAt = ca.String()
		for i := range ids {
			o.MealList = append(o.MealList, int32(ids[i]))
		}
		ors = append(ors, o)
		log.Println(ors)
	}
	for i := range ors {
		os.List = append(os.List, &ors[i])
	}
	log.Println(os)
	return os, nil
}

func (p PostgresRepository) Payment(ctx context.Context, request orderproto.Request) (int, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	if _, err := p.DB.ExecContext(c, "update orders set closed_at=current_timestamp where id=$1;", request.Id); err != nil {
		return 0, err
	}
	var price int
	if err := p.DB.QueryRowContext(c, "select price from orders where id=$1", request.Id).Scan(&price); err != nil {
		return 0, err
	}
	return price, nil
}

func (p PostgresRepository) Choose(ctx context.Context, request orderproto.MealChoice) error {
	c, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	log.Println(request)
	tx, err := p.DB.Begin()
	if request.OrderId == 0 {
		if _, err := tx.ExecContext(c, "insert into users values($1) on conflict do nothing", request.Id); err != nil {
			tx.Rollback()
			return err
		}
		if _, err := tx.ExecContext(c, "insert into orders (user_id) values ($1)", request.Id); err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.QueryRowContext(c, "select last_value from orders_id_seq").Scan(&request.OrderId); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(c, "update orders set meals = array_append(meals, $1) where id=$2", request.MealId, request.OrderId); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.ExecContext(c, "update orders set price=price+(select price from meals where id=$1) where id=$2", request.MealId, request.OrderId); err != nil {
		tx.Rollback()
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
	mls := make([]orderproto.Meal, 0)
	for rows.Next() {
		if err := rows.Scan(&m.Id, &m.Name, &m.Price); err != nil {
			return orderproto.Meals{}, err
		}
		mls = append(mls, m)
	}
	for i := range mls {
		ms.List = append(ms.List, &mls[i])
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
