package repository

import (
	"context"
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/pkg/database"

	"github.com/google/uuid"
)

type News interface {
	Create(ctx context.Context, news *domain.News) (*uuid.UUID, error)
	Get(ctx context.Context, count int) (*[]*domain.News, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.News, error)
	Update(ctx context.Context, news *domain.News) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type Author interface {
	Create(ctx context.Context, a *domain.Author) error
	Get(ctx context.Context, count int) (*[]*domain.Author, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.Author, error)
	Update(ctx context.Context, a *domain.Author) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}

type AuthorWithNews interface {
	GetAuthorNews(ctx context.Context, authorId uuid.UUID) (*[]domain.News, error)
	GetNewsAuthors(ctx context.Context, newsId uuid.UUID) (*[]domain.Author, error)
	SetNewsAuthors(ctx context.Context, newsId, authorId uuid.UUID) error
}

type Repository struct {
	News       News
	Author     Author
	AuthorNews AuthorWithNews
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		News:       NewNewsRepo(db),
		Author:     NewAuthorRepo(db),
		AuthorNews: NewAuthorNewsRepo(db),
	}
}
