package app

import (
	"common_service/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	var page int
	pageStr, ok := c.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(pageStr)
	}
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	var pageSize int
	pageSizeStr, ok := c.GetQuery("page_size")
	if ok {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}
	if pageSize <= 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}
