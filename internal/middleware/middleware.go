package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Logger — middleware для логирования запросов
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Выполняем следующий middleware или хендлер
		c.Next()

		// После выполнения — логируем
		duration := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		println(
			"[LOG] ", method, path,
			"→", statusCode,
			"(", duration.String(), ")",
		)
	}
}
