package controller

import "PicNest/internal/services"

type UploadController struct {
	UploadService *services.UploadService
}

// func NewUploadController(UploadService *services.UploadService) *UploadController {
// 	return &UploadController{UploadService: UploadService}
// }
