package transport

import (
	"github.com/FrostyCreator/news-portal/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Init() *echo.Echo {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	h.initAPI(e)

	return e
}

func (h *Handler) initAPI(e *echo.Echo) {
	api := e.Group("/api")
	{
		h.initNewsRoutes(api)
	}
}

func (h *Handler) initNewsRoutes(e *echo.Group) {
	news := e.Group("/news")
	{
		news.GET("", h.getNews)
		news.GET("/:id", h.getNewsById)
		news.POST("", h.createNews)
		news.PUT("", h.updateNews)
		news.DELETE("", h.deleteNews)
	}
}
