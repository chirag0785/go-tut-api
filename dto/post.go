package dto

type PostCreateDTO struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type PostUpdateDTO struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}