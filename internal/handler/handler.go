package handler

import (
	"my-gin-lesson/internal/model"
	"my-gin-lesson/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosters(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": model.Posts,
	})
}

func GetPosterById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, "ID - не число!")
		return
	}

	if post, found := service.FindPostByID(id); found {
		c.JSON(200, post)
	} else {
		c.JSON(404, gin.H{"error": "Пост не найден"})
	}

}

func CreatePoster(c *gin.Context) {
	var newPost model.Poster

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(400, gin.H{"error": "Неверный JSON: " + err.Error()})
		return
	}

	// Присваиваем новый ID (временно — просто увеличиваем счётчик)
	newPost.ID = len(model.Posts) + 1

	// Добавляем в "базу" (наш срез)
	model.Posts = append(model.Posts, newPost)

	// Отправляем созданный пост как ответ
	c.JSON(201, newPost) // 201 — статус "Created"

}

func DeletePosterById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID должен быть числом"})
		return
	}

	// Вызываем сервис для удаления
	success := service.DeletePoster(id)
	if !success {
		c.JSON(404, gin.H{"error": "Пост не найден"})
		return
	}

	c.JSON(204, "Poster удалён!") // 204 No Content — стандартный ответ при успешном удалении
}

// internal/handler/handler.go

func ShowPostsPage(c *gin.Context) {
	c.HTML(200, "posts.html", gin.H{
		"Posts":      model.Posts,
		"TotalCount": len(model.Posts),
	})
}

func ShowPostPage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(400, "error.html", gin.H{"Message": "ID должен быть числом"})
		return
	}

	if post, found := service.FindPostByID(id); found {
		c.HTML(200, "post.html", gin.H{"Post": post})
	} else {
		c.HTML(404, "error.html", gin.H{"Message": "Пост не найден"})
	}
}

func CreatePostFromForm(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")

	// Простая валидация
	if title == "" || len(title) < 3 {
		c.HTML(400, "error.html", gin.H{"Message": "Заголовок должен быть не короче 3 символов"})
		return
	}
	if body == "" || len(body) < 10 {
		c.HTML(400, "error.html", gin.H{"Message": "Текст должен быть не короче 10 символов"})
		return
	}

	// Создаём пост
	newPost := model.Poster{
		ID:    len(model.Posts) + 1,
		Title: title,
		Body:  body,
	}
	model.Posts = append(model.Posts, newPost)

	// Перенаправляем на главную (чтобы избежать повторной отправки формы)
	c.Redirect(303, "/")
}

func DeletePostFromForm(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.HTML(400, "error.html", gin.H{"Message": "Неверный ID"})
		return
	}

	if !service.DeletePoster(id) {
		c.HTML(404, "error.html", gin.H{"Message": "Пост не найден"})
		return
	}

	// Перенаправляем на главную
	c.Redirect(303, "/")
}
