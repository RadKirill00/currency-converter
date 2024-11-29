package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) 
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row 
	Begin(ctx context.Context) (pgx.Tx, error)
}

func newClient(ctx context.Context, maAttempts int, username, password, host, port, database string) (*Client, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	
	for maAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err := pgxpool.Connect(ctx, dsn) 
		if err != nil{
			fmt.Println("failed to connect")
			return nil, err
		}
 	}
}