package domain

type OneNewsWithAuthors struct {
	News    News
	Authors []Author
}

type NewsWithAuthors []OneNewsWithAuthors
