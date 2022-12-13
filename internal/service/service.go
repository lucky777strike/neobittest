package service

import (
	"camparser/internal/repository"
	"camparser/pkg/shodan"
)

type Service struct {
	repo   *repository.Repository
	shodan *shodan.Shodan
}

func NewService(repo *repository.Repository, sapi string) *Service {

	return &Service{repo: repo, shodan: shodan.New(sapi)}
}
