package transport

import (
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const defaultCountNews = 50

func (h *Handler) getNews(c echo.Context) error {
	countInQuery := c.QueryParam("count")

	var err error
	count := 0

	if countInQuery == "" {
		count = defaultCountNews
	} else {
		if count, err = strconv.Atoi(countInQuery); err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid count param"))
		}
	}

	news, err := h.Service.News.Get(c.Request().Context(), count)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get objects from repo: %s", err))
	}

	return c.JSON(http.StatusBadRequest, news)
}

func (h *Handler) getNewsById(c echo.Context) error {
	idInQuery := c.Param("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	news, err := h.Service.News.GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get object from repo: %s", err))
	}

	return c.JSON(http.StatusBadRequest, news)
}

func (h *Handler) createNews(c echo.Context) error {
	news := new(domain.News)

	if err := c.Bind(news); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	if !news.IsValid() {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	if err := h.Service.News.Create(c.Request().Context(), news); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed create news: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateNews(c echo.Context) error {
	news := new(domain.News)

	if err := c.Bind(news); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	if !news.IsValid() {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	if err := h.Service.News.Update(c.Request().Context(), news); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed create news: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteNews(c echo.Context) error {
	idInQuery := c.FormValue("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	if err = h.Service.News.DeleteById(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed delete news: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}
