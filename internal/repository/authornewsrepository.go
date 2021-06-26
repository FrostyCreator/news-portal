package repository

import (
	"context"
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/pkg/database"
	"github.com/FrostyCreator/news-portal/pkg/logger"

	"github.com/google/uuid"
)

type AuthorNewsRepo struct {
	db *database.DB
}

func NewAuthorNewsRepo(db *database.DB) *AuthorNewsRepo {
	return &AuthorNewsRepo{
		db: db,
	}
}

func (repo AuthorNewsRepo) GetAuthorNews(ctx context.Context, authorId uuid.UUID) (*[]domain.News, error) {
	news := make([]domain.News, 0)

	query := "SELECT news_id, title, text, created_at FROM news.view_authors_with_news WHERE author_id = $1;"
	rows, err := repo.db.Query(ctx, query, authorId)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", authorId)
		return nil, err
	}
	for rows.Next() {
		tempNews := domain.News{}
		if err = rows.Scan(&tempNews.ID, &tempNews.Title, &tempNews.Text, &tempNews.CreatedAt); err != nil {
			logger.LogError(err)
			return nil, err
		}
		news = append(news, tempNews)
	}

	return &news, nil
}

func (repo AuthorNewsRepo) GetNewsAuthors(ctx context.Context, newsId uuid.UUID) (*[]domain.Author, error) {
	authors := make([]domain.Author, 0)

	query := "SELECT author_id, first_name, last_name, father_name FROM news.view_authors_with_news WHERE news_id = $1;"
	rows, err := repo.db.Query(ctx, query, newsId)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", newsId)
		return nil, err
	}
	for rows.Next() {
		tempAuth := domain.Author{}
		if err = rows.Scan(&tempAuth.ID, &tempAuth.FirstName, &tempAuth.LastName, &tempAuth.FatherName); err != nil {
			logger.LogError(err)
			return nil, err
		}
		authors = append(authors, tempAuth)
	}

	return &authors, nil
}

func (repo AuthorNewsRepo) SetNewsAuthors(ctx context.Context, newsId, authorId uuid.UUID) error {
	query := "INSERT INTO news.tbl_author_news (id, author_id, news_id) VALUES (DEFAULT, $1, $2);"
	_, err := repo.db.Exec(ctx, query, authorId, newsId)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", authorId, newsId)
		return err
	}

	return nil
}
