package services

import (
	"strings"

	"github.com/google/uuid"
	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/errors"
	"gitlab.com/xpresslane1/xcommon/logging"
	"gitlab.com/xpresslane1/xcommon/models"
	"gitlab.com/xpresslane1/xintegrations/internal/repo"
)

type UrlShortnerService struct {
	appConfig *xconf.XAppConfig
	urlShortnerrepo *repo.UrlShortnerRepo
}

func NewUrlShortnerService(appconf *xconf.XAppConfig) *UrlShortnerService {
	return &UrlShortnerService{
		appConfig: appconf,
		urlShortnerrepo: repo.NewUrlShortnerRepo(appconf),
	}
}

func (uss *UrlShortnerService) GetShortUrl(requestId string,urlShortener *models.UrlShortner) errors.RestAPIError {
	m:="GetShortUrl"
	uss.logInfo(requestId,m,"Start")

	uss.GenerateShortUrls(requestId,urlShortener)
	
	urlShortener.CreatedTime = uss.appConfig.CurrentTime()
	urlShortener.Status = "ACTIVE"
	urlShortener.ModifiedTime = uss.appConfig.CurrentTime()

	err := uss.urlShortnerrepo.SaveUrl(urlShortener)
	if err.IsNotNull() {
		uss.logError(requestId,m,"Error in saving url - ",err)
		return err
	}

	uss.logInfo(requestId,m,"Complete")
	return errors.NO_ERROR()
}

func (uss *UrlShortnerService) GetLongUrlByShortUrl(requestId string,urlShortener *models.UrlShortner) errors.RestAPIError {
	m:="GetLongUrl"
	uss.logInfo(requestId,m,"Start")


	err := uss.urlShortnerrepo.GetLongUrlByShortUrl(urlShortener)
	if err.IsNotNull() {
		uss.logError(requestId,m,"Error in fetching longurl - ",err)
		return err
	}

	uss.logInfo(requestId,m,"Complete")
	return errors.NO_ERROR()
}

func (uss *UrlShortnerService) GetLongUrlByShortCode(requestId string,urlShortener *models.UrlShortner) errors.RestAPIError {
	m:="GetLongUrl"
	uss.logInfo(requestId,m,"Start")


	err := uss.urlShortnerrepo.GetLongUrlByShortCode(urlShortener,urlShortener.ShortCode)
	if err.IsNotNull() {
		uss.logError(requestId,m,"Error in fetching long url - ",err)
		return err
	}

	uss.logInfo(requestId,m,"Complete")
	return errors.NO_ERROR()
}


func (uss *UrlShortnerService) GenerateShortUrls(requestId string,urlShortener *models.UrlShortner) errors.RestAPIError{
	m:="GenerateShortUrls"
	uss.logInfo(requestId,m,"Start")

	// MD5

	// UUID
	uuid,_ := uuid.NewUUID()
	shortCode := strings.Replace(uuid.String(), "-", "", -1)[:7]
	urlShortener.ShortCode = shortCode

	uss.logInfo(requestId,m,"Complete")
	return errors.NO_ERROR()
}


func (uss *UrlShortnerService) logError(requestId, method string, message string, a ...interface{}) {
	logging.LogError(uss.appConfig,requestId, "UrlShortnerService", method, message, a)
}

func (uss *UrlShortnerService) logInfo(requestId, method string, message string, a ...interface{}) {
	logging.LogInfo(uss.appConfig,requestId, "UrlShortnerService", method, message, a)
}

