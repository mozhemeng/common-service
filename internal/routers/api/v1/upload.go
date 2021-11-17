package v1

import (
	"common_service/internal/service"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"common_service/pkg/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Description 上传文件
// @Tags upload
// @Accept mpfd
// @Produce json
// @Param file formData file true "文件"
// @Param file_type formData string true "文件类型"
// @Success 200 {object} app.Result{data=service.FileInfo}
// @Security JWT
// @Router /api/v1/upload [post]
func (u Upload) UploadFile(c *gin.Context) {
	resp := app.NewResponse(c)
	fileHeader, err := c.FormFile("file")
	if err != nil || fileHeader == nil {
		resp.ToError(errcode.InvalidParams.WithDetails(fmt.Errorf("wrong file: %w", err).Error()))
		return
	}
	fileType, err := strconv.Atoi(c.PostForm("file_type"))
	if err != nil || fileType <= 0 {
		resp.ToError(errcode.InvalidParams.WithDetails(fmt.Errorf("wrong file type: %w", err).Error()))
		return
	}

	svc := service.New(c)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), fileHeader)
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.Success(fileInfo)
}
