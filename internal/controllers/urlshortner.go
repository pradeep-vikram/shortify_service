package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/errors"
	"gitlab.com/xpresslane1/xcommon/helpers"
	"gitlab.com/xpresslane1/xcommon/logging"
	"gitlab.com/xpresslane1/xcommon/models"
	"gitlab.com/xpresslane1/xintegrations/internal/services"
)

type UrlShotnerController struct {
	appConfig *xconf.XAppConfig
	urlShortnerService *services.UrlShortnerService
}

func NewUrlShortnerController(appconfig *xconf.XAppConfig) *UrlShotnerController {
	return &UrlShotnerController{
		appConfig: appconfig,
		urlShortnerService: services.NewUrlShortnerService(appconfig),
	}
}

func (usc *UrlShotnerController)GetShortUrl(c *gin.Context) {
	m:="GetShortUrl"
	requestId := helpers.GetRequestId(c)
	usc.logInfo(requestId,m,"Start")
	
	urlShortener := models.UrlShortner{}
	if err:= c.ShouldBindJSON(&urlShortener);err != nil {
		usc.logError(requestId,m,"Error in binding request body - ",err.Error())
		restErr := errors.NewBadRequestError("Invalid request body")
		c.JSON(restErr.Status,restErr)
		return
	}

	err := usc.urlShortnerService.GetShortUrl(requestId,&urlShortener)
	if err.IsNotNull() {
		c.JSON(err.Status,err)
		return
	}

	usc.logInfo(requestId,m,"Complete")
	c.JSON(http.StatusOK,urlShortener)
}

func (usc *UrlShotnerController) GetLongUrl(c *gin.Context) {
	m:="GetLongUrl"
	requestId := helpers.GetRequestId(c)
	usc.logInfo(requestId,m,"Start")

	urlShortener := models.UrlShortner{}
	if err:= c.ShouldBindJSON(&urlShortener);err != nil {
		usc.logError(requestId,m,"Error in binding request body - ",err.Error())
		restErr := errors.NewBadRequestError("Invalid request body")
		c.JSON(restErr.Status,restErr)
		return
	}

	err := usc.urlShortnerService.GetLongUrlByShortCode(requestId,&urlShortener)
	if err.IsNotNull() {
		c.JSON(err.Status,err)
		return
	}

	usc.logInfo(requestId,m,"Complete")
	c.JSON(http.StatusOK,urlShortener)
}

func (usc *UrlShotnerController) logError(requestId, method string, message string, a ...interface{}) {
	logging.LogError(usc.appConfig,requestId, "UrlShotnerController", method, message, a)
}

func (usc *UrlShotnerController) logInfo(requestId, method string, message string, a ...interface{}) {
	logging.LogInfo(usc.appConfig,requestId, "UrlShotnerController", method, message, a)
}

