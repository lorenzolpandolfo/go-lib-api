package dto

type CreateBookRequest struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

type BookResponse struct {
	Title string `json:"title"`
	Author string `json:"author"`
}