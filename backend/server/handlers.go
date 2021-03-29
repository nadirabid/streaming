package server

import (
	"net/http"
	"strconv"
	"streaming/models"

	"github.com/labstack/echo"
)

func handleGetPaginatedContentList(c echo.Context) error {
	ctx := c.(*Context)
	db := ctx.db

	contentList := []models.Content{}
	if err := db.Preload("MiniSeries").Scopes(paginate(ctx)).Find(contentList).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusOK, contentList)
}

func handleGetContent(c echo.Context) error {
	ctx := c.(*Context)
	db := ctx.db

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorMessage(err.Error()))
	}

	content := &models.Content{}
	if err := db.Scopes(paginate(ctx)).Where("id = ?", id).First(content).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusOK, content)
}
