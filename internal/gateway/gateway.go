package gateway

import "github.com/jackc/pgx/v5/pgxpool"

type Gateway struct {
}

func NewGateway(db *pgxpool.Pool) *Gateway {
	return &Gateway{}
}
