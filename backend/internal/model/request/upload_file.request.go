package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UploadFileReq struct {
	UploadType string                `form:"uploadType" binding:"required,upload-type"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadFileToS3CompanyReq struct {
	UploadType string                `form:"uploadType" binding:"required,upload-type"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

func (u *UploadFileToS3CompanyReq) ToUploadS3Req() UploadFileReq {
	return UploadFileReq{
		UploadType: u.UploadType,
		File:       u.File,
	}
}

type UploadFileToS3CrewReq struct {
	UploadType string                `form:"uploadType" binding:"required,oneof=PROFILE_PICTURE"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

func (u *UploadFileToS3CrewReq) ToUploadS3Req() UploadFileReq {
	return UploadFileReq{
		UploadType: u.UploadType,
		File:       u.File,
	}
}

type GetSignedURLFileReq struct {
	ID         uuid.UUID
	UploadType string
}

type RedirectToS3SignedURLReq struct {
	ID string `uri:"id" binding:"required,uuid"`
}
