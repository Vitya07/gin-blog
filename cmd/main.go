package main

import (
	"my-gin-lesson/internal/config"
	"my-gin-lesson/internal/handler"

	//"my-gin-lesson/internal/middleware" для своего logger (являеться middleware)

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load() // подгружаем config (.env)

	r := gin.Default()
	//r.Use(middleware.Logger())    - logger

	r.LoadHTMLGlob("templates/*") // регистрируем папку с шаблонами

	r.Static("/static", "./static") // 💡 Подключаем статику (CSS, изображения)

	r.GET("/", handler.ShowPostsPage)

	// Создание поста (из формы)
	r.POST("/posts", handler.CreatePostFromForm)

	// Удаление поста (через форму)
	r.POST("/posts/:id/delete", handler.DeletePostFromForm)

	// HTML отображение
	r.GET("/posts/:id", handler.ShowPostPage)

	// //Регистрация маршрутов
	// r.GET("/", handler.ShowPostsPage)
	// r.GET("/posts", handler.GetPosters)
	// //r.GET("/posts/:id", handler.GetPosterById) использую как api для postman
	// r.POST("/posts", handler.CreatePoster)
	// r.DELETE("/posts/:id", handler.DeletePosterById)

	addr := ":" + cfg.ServerPort
	r.Run(addr)
}
