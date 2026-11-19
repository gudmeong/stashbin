package page

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mrmissx/stashbin/database/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func RawPageHandler(c echo.Context) error {
	var doc *model.Document = &model.Document{}
	slug := c.Param("slug")

	db := c.Get("db").(*mongo.Database)
	err := doc.GetBySlug(db, slug)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return nil
	}
	c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response().Header().Set("Content-Disposition", "inline")
	return c.String(http.StatusOK, doc.Content)
}
