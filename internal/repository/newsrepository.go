package repository

import (
	"context"
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/pkg/database"
	"github.com/FrostyCreator/news-portal/pkg/logger"

	"github.com/google/uuid"
)

const countMaxNews = 100

type NewsRepo struct {
	db *database.DB
}

func NewNewsRepo(db *database.DB) *NewsRepo {
	return &NewsRepo{
		db: db,
	}
}

func (repo *NewsRepo) Create(ctx context.Context, news *domain.News) (*uuid.UUID, error) {
	query := "INSERT INTO news.tbl_news (id, title, text, created_at) VALUES ($1, $2, $3, DEFAULT)"

	id := uuid.New()

	_, err := repo.db.Exec(ctx, query, id, news.Title, news.Text)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s, %s", news.Title, news.Text)
		return nil, err
	}

	return &id, nil
}

func (repo *NewsRepo) Get(ctx context.Context, count int) (*[]*domain.News, error) {
	if count > countMaxNews {
		count = countMaxNews
	}

	news := make([]*domain.News, 0)

	query := "SELECT * FROM news.tbl_news LIMIT $1"
	rows, err := repo.db.Query(ctx, query, count)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", count)
		return nil, err
	}
	for rows.Next() {
		tempNews := new(domain.News)
		if err = rows.Scan(&tempNews.ID, &tempNews.Title, &tempNews.Text, &tempNews.CreatedAt); err != nil {
			return nil, err
		}
		news = append(news, tempNews)
	}

	return &news, nil
}

func (repo *NewsRepo) GetById(ctx context.Context, id uuid.UUID) (*domain.News, error) {
	news := new(domain.News)

	query := "SELECT * FROM news.tbl_news WHERE id = $1"
	err := repo.db.QueryRow(ctx, query, id).
		Scan(&news.ID, &news.Title, &news.Text, &news.CreatedAt)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", id)
		return nil, err
	}

	return news, nil
}

func (repo *NewsRepo) Update(ctx context.Context, news *domain.News) error {
	query := "UPDATE news.tbl_news SET title = $1, text = $2 WHERE id = $3"

	_, err := repo.db.Exec(ctx, query, news.Title, news.Text, news.ID)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s, %s", news.Title, news.Text, news.ID)
		return err
	}

	return nil
}

func (repo *NewsRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM news.tbl_news WHERE id = $1"

	_, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", id.String())
		return err
	}

	return nil
}
