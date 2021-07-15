package service

import (
	"common_service/global"
	"common_service/pkg/errcode"
	"common_service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string `json:"name"`
	AccessUrl string `json:"access_url"`
}

func (svc *Service) UploadFile(fileType upload.FileType, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	savePath := upload.GetSavePath()
	if upload.CheckSavePath(savePath) {
		if err := upload.CreateSavePath(savePath, os.ModePerm); err != nil {
			return nil, errcode.UploadFailed.WithDetails(err.Error())
		}
	}
	if upload.CheckPermission(savePath) {
		return nil, errcode.UploadFailed
	}
	if !upload.CheckExt(fileType, fileName) {
		return nil, errcode.UploadExtNotSupported
	}
	if !upload.CheckSize(fileType, fileHeader) {
		return nil, errcode.UploadExcessMaxSize
	}

	dst := savePath + "/" + fileName
	if err := svc.ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		return nil, errcode.UploadFailed.WithDetails(err.Error())
	}

	fi := FileInfo{
		Name:      fileName,
		AccessUrl: global.AppSetting.UploadServerUrl + "/" + fileName,
	}

	return &fi, nil
}
