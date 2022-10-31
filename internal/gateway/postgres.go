package gateway

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jahngeor/avito-tech/pkg/utils"
	"time"

	"github.com/jahngeor/avito-tech/pkg/config"
	"log"
)

const ()

func NewPostgresDB(ctx context.Context, cfg *config.Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBConfig.Port, cfg.DBName)
	fmt.Println(dsn)
	err = utils.TryConnection(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.New(ctx, dsn)

		if err != nil {
			return err
		}
		return nil
	}, cfg.MaxAttempts, 5*time.Second)

	if err != nil {
		return nil, err
	}

	return pool, nil
}

// TestPing /* метод для пинга соединения с базой postgres */
func TestPing(pool *pgxpool.Pool, ctx context.Context) (err error) {
	for at := 0; at < 5; at++ {
		time.Sleep(2 * time.Second)
		if err = pool.Ping(ctx); err != nil {
			log.Printf("Тест пинга номер: %d:, %s", at, err.Error())
			continue
		}
		return nil
	}
	return err
}
