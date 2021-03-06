package server

import (
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/server/controllers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// init router
func InitRouter(engine *gin.Engine) {
	log.Infof("初始化路由...")
	engine.Use(CorsHandler())
	engine.Use(static.Serve("/",static.LocalFile(conf.Conf.Server.Static,false)))
	engine.NoRoute(func(c *gin.Context) {
		c.File(conf.Conf.Server.Static+"/index.html")
	})
	InitApiRouter(engine)
}

// init api router
func InitApiRouter(engine *gin.Engine) {
	v2:=engine.Group("/api")
	{
		v2.GET("/info",controllers.Info)
		v2.POST("/get",controllers.Get)
		v2.POST("/list",controllers.List)
		v2.POST("/search",controllers.Search)
		v2.POST("/office_preview",controllers.OfficePreview)
	}
	engine.GET("/d/*file_id",controllers.Down)
	engine.GET("/cache/:password",controllers.RefreshCache)
}