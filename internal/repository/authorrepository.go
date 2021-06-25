package repository

import (
	"context"
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/pkg/database"
	"github.com/FrostyCreator/news-portal/pkg/logger"

	"github.com/google/uuid"
)

const countMaxAuthors = 100

type AuthorRepo struct {
	db *database.DB
}

func NewAuthorRepo(db *database.DB) *AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}

func (repo *AuthorRepo) Create(ctx context.Context, a *domain.Author) error {
	query := "INSERT INTO news.tbl_author (id, first_name, last_name, father_name) VALUES (DEFAULT, $1, $2, $3);"

	_, err := repo.db.Exec(ctx, query, a.FirstName, a.LastName, a.FatherName)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s, %s", a.FirstName, a.LastName, a.FatherName)
		return err
	}

	return nil
}

func (repo *AuthorRepo) Get(ctx context.Context, count int) (*[]*domain.Author, error) {
	if count > countMaxAuthors {
		count = countMaxAuthors
	}

	authors := make([]*domain.Author, 0, count)

	query := "SELECT * FROM news.tbl_author LIMIT $1"
	rows, err := repo.db.Query(ctx, query, count)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", count)
		return nil, err
	}
	for rows.Next() {
		tempAuthor := new(domain.Author)
		if err = rows.Scan(&tempAuthor.ID, &tempAuthor.FirstName, &tempAuthor.LastName, &tempAuthor.FatherName); err != nil {
			return nil, err
		}
		authors = append(authors, tempAuthor)
	}

	return &authors, nil
}

func (repo *AuthorRepo) GetById(ctx context.Context, id uuid.UUID) (*domain.Author, error) {
	author := new(domain.Author)

	query := "SELECT * FROM news.tbl_author WHERE id = $1"
	err := repo.db.QueryRow(ctx, query, id).
		Scan(&author.ID, &author.FirstName, &author.LastName, &author.FatherName)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", id)
		return nil, err
	}

	return author, nil
}

func (repo *AuthorRepo) Update(ctx context.Context, a *domain.Author) error {
	query := "UPDATE news.tbl_author SET first_name = $1, last_name = $2, father_name = $3 WHERE id = $4"

	_, err := repo.db.Exec(ctx, query, a.FirstName, a.LastName, a.FatherName, a.ID)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s, %s", a.FirstName, a.LastName, a.FatherName, a.ID)
		return err
	}

	return nil
}

func (repo *AuthorRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM news.tbl_author WHERE id = $1"

	_, err := repo.db.Exec(ctx, query, id)
	if err != nil {
		logger.LogErrorf("Query: %s", query)
		logger.LogErrorf("Arguments: %s", id.String())
		return err
	}

	return nil
}
