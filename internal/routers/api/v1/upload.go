package v1

import (
	"common_service/global"
	"common_service/internal/service"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"common_service/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		resp.ToError(errcode.InvalidParams.WithDetails(errors.Wrap(err, "wrong file").Error()))
		return
	}
	fileType, err := strconv.Atoi(c.PostForm("file_type"))
	if err != nil || fileType <= 0 {
		resp.ToError(errcode.InvalidParams.WithDetails(errors.Wrap(err, "wrong file type").Error()))
		return
	}

	svc := service.New(c)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), fileHeader)
	switch errors.Cause(err) {
	case nil:
		resp.Success(fileInfo)
	case errcode.UploadExtNotSupported:
		fallthrough
	case errcode.UploadExcessMaxSize:
		fallthrough
	case errcode.UploadFailed:
		resp.ToError(err.(*errcode.ApiError))
	default:
		global.Logger.Error(errors.Wrap(err, "svc.UploadFile"))
		resp.ToError(errcode.InternalError)
	}
}
