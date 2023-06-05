package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/errors"
	"gitlab.com/xpresslane1/xcommon/models"
)

type UrlShortnerRepo struct {
	AppConfig *xconf.XAppConfig
	DB        *gorm.DB
}

func NewUrlShortnerRepo(appconf *xconf.XAppConfig) *UrlShortnerRepo {
	return &UrlShortnerRepo{
		AppConfig: appconf,
		DB:        appconf.DB,
	}
}

func (usr *UrlShortnerRepo) SaveUrl(urlShortner *models.UrlShortner) errors.RestAPIError {
	db := usr.AppConfig.DB
	if err := db.Table(models.Table_URL_SHORTNER).Save(urlShortner).Error; err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error in saving url - %s", err.Error()))
	}
	return errors.NO_ERROR()
}

func (usr *UrlShortnerRepo) GetLongUrlByShortUrl(urlShortner *models.UrlShortner) errors.RestAPIError {
	db := usr.AppConfig.DB
	if err := db.Table(models.Table_URL_SHORTNER).Where(&models.UrlShortner{ShortUrl: urlShortner.ShortUrl}).Find(&urlShortner).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewBadRequestError("Invalid Url")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error in fetching url - %s", err.Error()))
	}
	return errors.NO_ERROR()
}

func (usr *UrlShortnerRepo) GetLongUrlByShortCode(urlShortner *models.UrlShortner, shortCode string) errors.RestAPIError {
	db := usr.AppConfig.DB
	if err := db.Table(models.Table_URL_SHORTNER).Where("short_code = ?", shortCode).Find(&urlShortner).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewBadRequestError("Invalid Url")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error in fetching url - %s", err.Error()))
	}
	return errors.NO_ERROR()
}
