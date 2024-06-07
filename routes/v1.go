package routes

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func serveV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	photoGroup := v1.Group("photos")
	photoController := controller.Photo{}
	{
		photoGroup.GET("", photoController.FindAll)
		photoGroup.GET(":id", photoController.FindOne)
		photoGroup.POST("", photoController.Create)
		photoGroup.PATCH(":id", photoController.Update)
		photoGroup.DELETE(":id", photoController.Delete)
	}

}
