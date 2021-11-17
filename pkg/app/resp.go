package app

import (
	"common_service/global"
	"common_service/pkg/errcode"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
}

type Result struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data"`
}

type PagedResult struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Pager Pager       `json:"pager"`
	Data  interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Success(data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	resp := Result{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
	r.Ctx.JSON(http.StatusOK, resp)
}

func (r *Response) SuccessList(data interface{}, totalCount int) {
	if data == nil {
		data = []string{}
	}

	resp := PagedResult{
		Code: 200,
		Msg:  "success",
		Data: data,
		Pager: Pager{
			Page:       GetPage(r.Ctx),
			PageSize:   GetPageSize(r.Ctx),
			TotalCount: totalCount,
		},
	}
	r.Ctx.JSON(http.StatusOK, resp)
}

func (r *Response) toApiError(err *errcode.ApiError) {
	r.Ctx.AbortWithStatusJSON(err.StatusCode(), err.JSON())
}

func (r *Response) ToError(err error) {
	var apiError *errcode.ApiError
	if errors.As(err, &apiError) {
		r.toApiError(apiError)
		return
	}
	var valError validator.ValidationErrors
	if errors.As(err, &valError) {
		r.toApiError(errcode.InvalidParams.WithValidErrs(valError.Translate(global.Trans)))
		return
	}
	global.Logger.Error(err)
	if global.ServerSetting.RunMode == "debug" {
		r.toApiError(errcode.InternalError.WithDetails(err.Error()))
	} else {
		r.toApiError(errcode.InternalError)
	}
}
