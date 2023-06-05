package app

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/models"
)

var (
	router = gin.New()
)

func StartApplication(appConf *xconf.XAppConfig) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// ExposeHeaders:    []string{"X-Total-Count"},
	}))
	
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/ping"},
	}))
	router.Use(gin.Recovery())
	pprof.Register(router)

	basicAuthUrlsToMap(appConf)
	mapUrls(appConf)
	runDbMigrate(appConf)
	log.Printf("Starting service: %v on port %v\n", appConf.Server.APINAME, appConf.Server.APIPORT)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", appConf.Server.APIPORT),
		Handler:      router,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  1 * time.Minute,
	}
	s.ListenAndServe()

	//router.Run(fmt.Sprintf(":%v", appConf.Server.APIPORT))
}

func runDbMigrate(appConf *xconf.XAppConfig) {
	db := appConf.DB
	gormDB := db.Table(models.Table_URL_SHORTNER).AutoMigrate(&models.UrlShortner{})
	if gormDB.Error != nil {
		log.Println("Error in migrating ",models.Table_URL_SHORTNER," - ",gormDB.Error)
	}
}

func basicAuthUrlsToMap(appConf *xconf.XAppConfig) {
	urlmap := make(map[int]string)
	for i, url := range strings.Split(appConf.Server.BASIC_AUTHURL, "|") {
		urlmap[i] = url
	}

	appConf.Server.BASIC_AUTHURLMAP = urlmap
}
