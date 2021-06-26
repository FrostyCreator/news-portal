package pg

import (
	"context"
	"fmt"
	"github.com/FrostyCreator/news-portal/internal/config"
	"github.com/FrostyCreator/news-portal/pkg/database"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(config config.PostgreSQLConfig) (*database.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &database.DB{Pool: db}, nil
}
