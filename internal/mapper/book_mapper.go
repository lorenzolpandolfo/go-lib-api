package mapper

import (
	"go-lib-api/internal/dto"
	"go-lib-api/internal/model"
)


func ToBookModel(req dto.CreateBookRequest) model.Book {

	return model.Book{
		Title: req.Title,
		Author: req.Author,
	}
}

func ToBookResponse(book model.Book) dto.BookResponse {

	return dto.BookResponse{
		Title: book.Title,
		Author: book.Author,
	}
}