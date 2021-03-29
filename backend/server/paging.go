package server

import (
	"math"
	"strconv"
	"streaming/models"
	"streaming/utils"

	"gorm.io/gorm"
)

func paging(ctx *Context, db *gorm.DB, result interface{}) (*models.ListResponse, error) {
	page, err := strconv.Atoi(utils.GetStringOrDefault(ctx.QueryParam("page"), "0"))
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(utils.GetStringOrDefault(ctx.QueryParam("size"), "20"))
	if err != nil {
		return nil, err
	}

	if page < 0 {
		page = 0
	}

	list := models.ListResponse{}
	count := int64(0)
	offset := page * size

	err = db.Model(result).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if size == 0 {
		err = db.Find(result).Error
	} else {
		err = db.Limit(size).Offset(offset).Find(result).Error
	}

	if err != nil {
		return nil, err
	}

	list.Size = size
	list.Results = result
	list.Page = page

	if size <= 0 {
		list.Pages = 1
	} else {
		list.Pages = int(math.Ceil(float64(count) / float64(size)))
	}

	return &list, nil
}
