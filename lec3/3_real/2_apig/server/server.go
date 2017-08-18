package server

import (
	"github.com/jirfag/gointensive/lec3/3_production/2_apig/middleware"
	"github.com/jirfag/gointensive/lec3/3_production/2_apig/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
