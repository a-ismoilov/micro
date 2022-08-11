package postgres

import (
	"context"
	"database/sql"
	"github.com/akbarshoh/microOLX/protos/userproto"
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

func (p PostgresRepository) Log(context.Context, *userproto.User) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (p PostgresRepository) Payment(context.Context, *userproto.PayRequest) error {
	return status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (p PostgresRepository) UserList(context.Context, *userproto.Admin) (*userproto.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (p PostgresRepository) LogAdmin(context.Context, *userproto.Admin) error {
	return status.Errorf(codes.Unimplemented, "method LogAdmin not implemented")
}
