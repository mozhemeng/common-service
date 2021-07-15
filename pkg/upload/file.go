package upload

import (
	"common_service/global"
	"common_service/pkg/utils"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeDoc
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CheckExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
				return true
			}
		}
	case TypeDoc:
		for _, allowExt := range global.AppSetting.UploadDocAllowExts {
			if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
				return true
			}
		}
	}
	return false
}

func CheckSize(t FileType, fh *multipart.FileHeader) bool {
	f, err := fh.Open()
	if err != nil {
		return false
	}
	defer f.Close()
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size <= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	case TypeDoc:
		if size <= global.AppSetting.UploadDocMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}
