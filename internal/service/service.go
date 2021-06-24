package service

import (
	"github.com/FrostyCreator/news-portal/internal/repository"
)

type Service struct {
	News repository.News
}

func InitService(repo *repository.Repository) *Service {
	return &Service{
		News: repo.News,
	}
}
