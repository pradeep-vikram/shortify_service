package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/middlewares"
	"gitlab.com/xpresslane1/xintegrations/internal/controllers"
)

func mapUrls(appconfig *xconf.XAppConfig) {
	urlshotnerController := controllers.NewUrlShortnerController(appconfig)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	urlApi := router.Group("/v1")
	{
		
		urlApi.Use(middlewares.BasicAuthMiddleware(appconfig.Server.XPRESSLANE_ADMIN_CLIENT_ID, appconfig.Server.XPRESSLANE_ADMIN_CLIENT_SECRET, appconfig.Server.BASIC_AUTHURLMAP))
		urlApi.POST("/short-url", urlshotnerController.GetShortUrl)
		urlApi.POST("/long-url", urlshotnerController.GetLongUrl)
	}

}
