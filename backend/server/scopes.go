package server

import (
	"strconv"
	"streaming/utils"

	"gorm.io/gorm"
)

func paginate(ctx *Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, err := strconv.Atoi(utils.GetStringOrDefault(ctx.QueryParam("page"), "0"))
		if err != nil {
			db.AddError(err)
		}

		if page < 0 {
			page = 0
		}

		size, err := strconv.Atoi(utils.GetStringOrDefault(ctx.QueryParam("size"), "20"))
		if err != nil {
			db.AddError(err)
		}

		if size < 0 {
			size = 20
		}

		offset := page * size

		return db.Limit(size).Offset(offset)
	}
}
