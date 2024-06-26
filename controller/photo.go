package controller

import (
	"fmt"
	"go-api/dto"
	"go-api/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Photo struct {
}

var photos []models.Photo

func (p *Photo) FindAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if limit == 0 {
		ctx.JSON(200, photos)
	}
	ctx.JSON(200, photos[:limit])

}

func (p *Photo) FindOne(ctx *gin.Context) {
	id, _ := getID(ctx)

	for _, p := range photos {
		if p.ID == id {
			ctx.JSON(200, p)
			return
		}
	}
	ctx.JSON(404, gin.H{"error": "photo not found"})
}

func (p *Photo) Create(ctx *gin.Context) {
	//json
	// var form dto.PhotoForm
	// if err := ctx.ShouldBindJSON(&form); err != nil {
	// 	ctx.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	id := uint(len(photos)) + 1
	var form dto.PhotoForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	file, _ := ctx.FormFile("image")
	//path := "/upload/photos/" + strconv.Itoa((int(id))) + "/" + file.Filename
	path := fmt.Sprintf("/upload/photos/%d/%s", id, file.Filename)

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	photo := models.Photo{
		ID:    id,
		Name:  form.Name,
		Price: form.Price,
		Image: path,
	}
	photos := append(photos, photo)
	ctx.JSON(200, photos)
}

func (p *Photo) Update(ctx *gin.Context) {
	var form dto.UpdatePhotoForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := getID(ctx)

	for index, p := range photos {
		if p.ID == id {
			target := photos[index]

			if form.Name != "" {
				target.Name = form.Name
			}
			if form.Price != 0.0 {
				target.Price = form.Price
			}

			file, _ := ctx.FormFile("image")
			if file != nil {
				oldFile := target.Image
				path := fmt.Sprintf("/upload/photos/%d/%s", id, file.Filename)

				if err := ctx.SaveUploadedFile(file, path); err != nil {
					ctx.JSON(400, gin.H{"error": err.Error()})
					return
				}
				pwd, _ := os.Getwd()
				os.Remove(pwd + "/" + oldFile)
				target.Image = path
			}

			ctx.JSON(200, target)
			return
		}
	}
	ctx.JSON(404, gin.H{"error": "photo not found"})
}

func (p *Photo) Delete(ctx *gin.Context) {
	id, _ := getID(ctx)

	var i int
	for index, p := range photos {
		if p.ID == id {
			//target := photos[index]
			// pwd, _ := os.Getwd()
			// os.Remove(pwd + "/" + target.Image)
			i = index
			photos = append(photos[:i], photos[i+1:]...)
			ctx.Status(204)
			return
		}

	}
	ctx.JSON(404, gin.H{"error": "photo not found"})
}
