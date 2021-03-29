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
	q := db.Preload("MiniSeries")

	listResponse, err := paging(ctx, q, &contentList)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusOK, listResponse)
}

func handleGetContent(c echo.Context) error {
	ctx := c.(*Context)
	db := ctx.db

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorMessage(err.Error()))
	}

	content := &models.Content{}
	if err := db.Where("id = ?", id).First(content).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorMessage(err.Error()))
	}

	return ctx.JSON(http.StatusOK, content)
}
