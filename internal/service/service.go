package service

import (
	"github.com/go-redis/redis"
	"github.com/onemgvv/linkshot.git/internal/repository"
)

type Links interface {
	ShortLink(link string) (string, error)
	FindLink(link string) (string, error)
}

type Services struct {
	Links Links
}

type Deps struct {
	Repos *repository.Repositories
	Redis *redis.Client
}

func NewServices(deps *Deps) *Services {
	linksService := NewLinksService(deps.Redis)
	return &Services{
		Links: linksService,
	}
}
