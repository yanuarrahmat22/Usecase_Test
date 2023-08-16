package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")

	NewOrderRouter(v1, db)

	return r
}
