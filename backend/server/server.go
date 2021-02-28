package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// TODO: https://qvault.io/2020/09/04/golang-video-stream-server/

func New(v *viper.Viper) error {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.Static("/hls", "assets/hls")

	return e.Start(fmt.Sprintf("0.0.0.0:%s", v.GetString("server.port")))
}
