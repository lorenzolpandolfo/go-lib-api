package repository

import (
	"context"
	"go-lib-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	Create(ctx context.Context, book model.Book) (int, error)
	GetByID(ctx context.Context, id int) (*model.Book, error)
}

type bookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) BookRepository {
	return &bookRepository{db: db}
}

func (b *bookRepository) Create(ctx context.Context, book model.Book) (int, error) {
	query := `INSERT INTO books (author, title, created_at) VALUES ($1, $2, NOW()) RETURNING id`

	var id int

	err := b.db.QueryRowContext(ctx, query, book.Author, book.Title).Scan(&id)
	return id, err
}

func (b *bookRepository) GetByID(ctx context.Context, id int) (*model.Book, error) {
	var book model.Book

	query := `SELECT id, author, title, created_at FROM books WHERE id = $1`

	err := b.db.GetContext(ctx, &book, query, id)

	if err != nil {
		return nil, err
	}

	return &book, nil
}


