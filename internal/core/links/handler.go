package links

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	repo Repo
}

func NewLinkHandler(r Repo) Handler {
	return Handler{repo: r}
}

type CreateLink struct {
	Url string `json:"url" validate:"required"`
}

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.POST("", func(c *echo.Context) error {

		input := new(CreateLink)

		if err := c.Bind(input); err != nil {
			return err
		}

		if input.Url == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "url is required")
		}

		link, err := h.repo.Create(c.Request().Context(), input.Url)

		if err != nil {
			log.Panicln(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Error while link creation")
		}

		return c.JSON(http.StatusOK, link)

	})
}
