package service

import (
	"context"
	"errors"
	"go-lib-api/internal/dto"
	"go-lib-api/internal/mapper"
	"go-lib-api/internal/repository"
	"strings"
)

type BookService interface {
	CreateBook(ctx context.Context, req dto.CreateBookRequest) (*dto.BookResponse, error)
	GetBook(ctx context.Context, id int) (*dto.BookResponse, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (b *bookService) CreateBook(ctx context.Context, req dto.CreateBookRequest) (*dto.BookResponse, error) {

	if strings.ToLower(req.Author) == "john" {
		return nil, errors.New("author must not be John")
	}

	bookModel := mapper.ToBookModel(req)

	id, err := b.repo.Create(ctx, bookModel)

	if err != nil {
		return nil, err
	}

	bookModel.ID = id
	response := mapper.ToBookResponse(bookModel)

	return &response, nil
}

func (b *bookService) GetBook(ctx context.Context, id int) (*dto.BookResponse, error) {

	bookResponse, err := b.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	response := mapper.ToBookResponse(*bookResponse)

	return &response, nil
}


