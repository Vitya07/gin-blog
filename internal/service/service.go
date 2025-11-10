package service

import "my-gin-lesson/internal/model"

// internal/model/model.go
//Поиск poster по id
func FindPostByID(id int) (*model.Poster, bool) {
	for _, poster := range model.Posts {
		if poster.ID == id {
			return &poster, true
		}
	}
	return nil, false
}

func DeletePoster(id int) bool {
	for i, post := range model.Posts {
		if post.ID == id {
			// Удаляем элемент по индексу i
			model.Posts = append(model.Posts[:i], model.Posts[i+1:]...)
			return true
		}
	}
	return false
}
