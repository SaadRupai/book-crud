package routes

import (
	"book-crud/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	//grouping route endpoints
	book := e.Group("/bookstore")

	//initializing http methods - routing endpoints and their handlers
	book.POST("/book", controllers.CreateBook)
	book.GET("/book", controllers.GetBook)
	book.PUT("/book/:bookID", controllers.UpdateBook)
	book.DELETE("/book/:bookID", controllers.DeleteBook)
}
