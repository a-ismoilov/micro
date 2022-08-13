package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/akbarshoh/microOLX/protos/userproto"
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

func (p PostgresRepository) Log(ctx context.Context, request *userproto.User) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	var check string
	log.Println(request)
	if err := p.DB.QueryRowContext(c, "select username from users where id=$1", request.Id).Scan(&check); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	log.Println(check)
	defer cancel()
	if check == "" {
		if _, err := p.DB.ExecContext(c, "insert into users (username, budget) values ($1,$2)", request.Username, request.Budget); err != nil {
			return err
		}
	}
	return nil
}
func (p PostgresRepository) Payment(ctx context.Context, request *userproto.PayRequest) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(c, `with b as (select price from orders where id=$1)
		update users set budget=budget-b where id=$2`, request.OrderId, request.Id); err != nil {
		tx.Rollback()
		return err
	}
	if _, err := tx.ExecContext(c, "update orders set closed_at=current_timestamp"); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (p PostgresRepository) UserList(ctx context.Context, request *userproto.Admin) (*userproto.Users, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	u := userproto.User{}
	us := userproto.Users{}
	rows, err := p.DB.QueryContext(c, "select * from users;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&u.Id, &u.Username, &u.Budget); err != nil {
			return nil, err
		}
		us.List = append(us.List, &u)
	}
	return &us, nil
}
func (p PostgresRepository) LogAdmin(ctx context.Context, request *userproto.Admin) error {
	c, cancel := context.WithTimeout(ctx, time.Second*3)
	var pass string
	defer cancel()
	row := p.DB.QueryRowContext(c, "select password from admins where id=$1", request.Id)
	if err := row.Scan(&pass); err != nil {
		return err
	}
	if pass != request.Password {
		return fmt.Errorf("password mismatch")
	}
	return nil
}
