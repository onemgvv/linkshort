package service

import (
	"github.com/go-redis/redis"
	"github.com/onemgvv/linkshot.git/pkg/generator"
	"time"
)

type LinksService struct {
	redis *redis.Client
}

func NewLinksService(rds *redis.Client) *LinksService {
	return &LinksService{
		redis: rds,
	}
}

func (s *LinksService) ShortLink(link string) (string, error) {
	shortLink := generator.ShortLinkGenerate(7)
	if err := s.redis.Set(shortLink, link, 180*time.Hour).Err(); err != nil {
		return "", err
	}

	return shortLink, nil
}

func (s *LinksService) FindLink(link string) (string, error) {
	return s.redis.Get(link).Result()
}
