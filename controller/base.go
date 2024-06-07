package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getID(ctx *gin.Context) (uint, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
