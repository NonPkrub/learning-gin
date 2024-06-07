package dto

import "mime/multipart"

type PhotoForm struct {
	Name  string                `form:"name" binding:"required"`
	Price float64               `form:"price" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

type UpdatePhotoForm struct {
	Name  string                `form:"name"`
	Price float64               `form:"price" binding:"omitempty,min=0"`
	Image *multipart.FileHeader `form:"image"`
}
