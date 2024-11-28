package ep

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware для проверки авторизации
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer my-secret-token" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func NewEntrypoint() error {
	// Создаем роутер Gin
	router := gin.Default()

	// Применяем middleware к защищенным маршрутам
	protected := router.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/data", handleData)
		protected.POST("/data", handlePostData)
	}

	// Открытый маршрут
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API",
		})
	})

	// Запуск сервера
	return router.Run(":8080")
}

// Обработчик для GET-запроса на /api/data
func handleData(c *gin.Context) {
	// Чтение кастомного заголовка
	customHeader := c.GetHeader("X-Custom-Header")
	if customHeader == "" {
		customHeader = "default-value"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Here is your protected data",
		"customHeader": customHeader,
	})
}

// Обработчик для POST-запроса на /api/data
func handlePostData(c *gin.Context) {
	// Чтение заголовков
	contentType := c.GetHeader("Content-Type")
	c.JSON(http.StatusOK, gin.H{
		"message":     "POST request successful",
		"contentType": contentType,
	})
}
