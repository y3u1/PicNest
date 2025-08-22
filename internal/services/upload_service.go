package services

import (
	"PicNest/internal/app/config"
	"PicNest/internal/utils"
	"mime/multipart"
	"path/filepath"
)

type UploadService struct {
}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (up *UploadService) UploadImage(file *multipart.FileHeader) (string, error) {

	// 获取文件名
	filename := file.Filename
	// 获取文件后缀
	ext := filepath.Ext(filename)
	// 生成新的文件名
	newFilename := utils.RandName() + ext
	filepath := filepath.Join(config.Conf.App.FileSavePath, newFilename)
	// 保存文件到指定目录
	if err := utils.SaveFile(file, filepath); err != nil {
		return "", err
	}

	return newFilename, nil
}
