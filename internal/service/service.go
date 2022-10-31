package service

import "github.com/jahngeor/avito-tech/internal/gateway"

type Service struct {
}

func NewServices(gtw *gateway.Gateway) *Service {
	return &Service{}
}
