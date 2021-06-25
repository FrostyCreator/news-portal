package transport

import (
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/internal/utils"
	"github.com/FrostyCreator/news-portal/pkg/logger"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const defaultCountAuthors = 50

func (h *Handler) getAuthors(c echo.Context) error {
	countInQuery := c.QueryParam("count")

	var err error
	count := 0

	if countInQuery == "" {
		count = defaultCountAuthors
	} else {
		if count, err = strconv.Atoi(countInQuery); err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid count param"))
		}
	}

	authors, err := h.Service.Authors.Get(c.Request().Context(), count)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get objects from repo: %s", err))
	}

	return c.JSON(http.StatusBadRequest, authors)
}

func (h *Handler) getAuthorById(c echo.Context) error {
	idInQuery := c.Param("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	author, err := h.Service.Authors.GetById(c.Request().Context(), id)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get object from repo: %s", err))
	}

	return c.JSON(http.StatusBadRequest, author)
}

func (h *Handler) GetAuthorNews(c echo.Context) error {
	idInQuery := c.Param("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	news, err := h.Service.AuthorsWithNews.GetAuthorNews(c.Request().Context(), id)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get object from repo: %s", err))
	}

	return c.JSON(http.StatusBadRequest, news)
}

func (h *Handler) createAuthor(c echo.Context) error {
	author := new(domain.Author)

	if err := c.Bind(author); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid author param"))
	}

	if !author.IsValid() {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid author param"))
	}

	if err := h.Service.Authors.Create(c.Request().Context(), author); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed create author: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}

func (h *Handler) updateAuthor(c echo.Context) error {
	author := new(domain.Author)

	if err := c.Bind(author); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid author param"))
	}

	if !author.IsValid() {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid author param"))
	}

	if err := h.Service.Authors.Update(c.Request().Context(), author); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed update author: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}

func (h *Handler) deleteAuthor(c echo.Context) error {
	idInQuery := c.FormValue("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	if err = h.Service.Authors.DeleteById(c.Request().Context(), id); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed delete news: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}
