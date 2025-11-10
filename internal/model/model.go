package model

type Poster struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"` // binding:"required" - проверка на пустоту, если "", -> ошибка
	Body  string `json:"body" binding:"required"`
}

var Posts = []Poster{
	{ID: 1, Title: "Первый пост", Body: "Привет, это мой блог!"},
}
