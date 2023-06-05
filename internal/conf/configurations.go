package conf

import (
	xconf "gitlab.com/xpresslane1/xcommon/conf"
)

func NewConfig() *xconf.XAppConfig {
	return &xconf.XAppConfig{
		Server: xconf.XServerConfig{
			APIPORT:                        xconf.GetEnvAsInt("APIPORT", 8085),
			APINAME:                        xconf.GetEnv("APINAME", "xintegrations"),
			XPRESSLANE_ADMIN_CLIENT_ID:     xconf.GetEnv("XPRESSLANE_ADMIN_CLIENT_ID", "xintegrations"),
			XPRESSLANE_ADMIN_CLIENT_SECRET: xconf.GetEnv("XPRESSLANE_ADMIN_CLIENT_SECRET", "Ify}7KWpY_[T"),
			BASIC_AUTHURL:                  "/v1/short-url|/v1/long-url",
		},
		DBConfig: xconf.XDBConfig{
			DBHOST:     xconf.GetEnv("DBHOST", "localhost"),
			DBPASSWORD: xconf.GetEnv("DBPASSWORD", "root"),
			DBUSER:     xconf.GetEnv("DBUSER", "root"),
			DBNAME:     xconf.GetEnv("DBNAME", "xintegration"),
			DBPORT:     xconf.GetEnv("DBPORT", "3306"),
			DBLOGMODE:  xconf.GetEnvAsBool("DBLOGMODE", false),
		},
	}
}
