package transport

import (
	"fmt"
	"github.com/FrostyCreator/news-portal/internal/domain"
	"github.com/FrostyCreator/news-portal/internal/utils"
	"github.com/FrostyCreator/news-portal/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

const defaultCountNews = 50

func (h Handler) getNewsWithAuthors(c echo.Context) error {
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

	newsWithAuthors := domain.NewsWithAuthors{}
	oneNewsWithAuthors := domain.OneNewsWithAuthors{}

	for _, n := range *news {
		oneNewsWithAuthors.News = *n
		authors, err := h.Service.AuthorsWithNews.GetNewsAuthors(c.Request().Context(), n.ID)
		if err != nil {
			logger.LogError(err)
			return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get objects from repo: %s", err))
		}
		oneNewsWithAuthors.Authors = *authors

		newsWithAuthors = append(newsWithAuthors, oneNewsWithAuthors)
	}

	return c.JSON(http.StatusOK, newsWithAuthors)
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

	return c.JSON(http.StatusOK, news)
}

func (h *Handler) GetNewsAuthors(c echo.Context) error {
	idInQuery := c.Param("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	authors, err := h.Service.AuthorsWithNews.GetNewsAuthors(c.Request().Context(), id)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed get object from repo: %s", err))
	}

	return c.JSON(http.StatusOK, authors)
}

func (h *Handler) createNews(c echo.Context) error {
	news := new(domain.News)

	if err := c.Bind(news); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	if !news.IsValid() {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news param"))
	}

	id, err := h.Service.News.Create(c.Request().Context(), news)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed create news: %s", err))
	}

	return c.JSON(http.StatusOK, id)
}

func (h *Handler) setNewsToAuthor(c echo.Context) error {
	dataFromQuery := struct {
		AuthorsIds []string `json:"authors_ids" form:"authors_ids"`
		NewsId     string   `json:"news_id" form:"news_id"`
	}{}

	err := c.Bind(&dataFromQuery)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid authors ids param"))
	}

	fmt.Println("authorIdFromReq -", dataFromQuery)

	newsId, err := uuid.Parse(dataFromQuery.NewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid news id param"))
	}

	for _, authorId := range dataFromQuery.AuthorsIds {
		authId, err := uuid.Parse(authorId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid author id param"))
		}
		if err := h.Service.AuthorsWithNews.SetNewsAuthors(c.Request().Context(), newsId, authId); err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed create news: %s", err))
		}

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
	idInQuery := c.Param("id")

	id, err := uuid.Parse(idInQuery)
	if err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewBadRequest("invalid id param"))
	}

	if err = h.Service.News.DeleteById(c.Request().Context(), id); err != nil {
		logger.LogError(err)
		return c.JSON(http.StatusBadRequest, utils.NewInternalf("failed delete news: %s", err))
	}

	return c.JSON(http.StatusOK, "ok")
}
