package server

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"streaming/models"
)

// Context is an extention of echo.Context
type Context struct {
	echo.Context
	db         *gorm.DB
	viper      *viper.Viper
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func newContext(
	v *viper.Viper,
	c echo.Context,
	db *gorm.DB,
	sess *session.Session,
	uploader *s3manager.Uploader,
	downloader *s3manager.Downloader,

) *Context {
	return &Context{
		c,
		db,
		v,
		sess,
		uploader,
		downloader,
	}
}

// TODO: https://qvault.io/2020/09/04/golang-video-stream-server/

func New(v *viper.Viper) error {
	var err error

	// INIT S3

	s3Region := v.GetString("s3.region")
	var sess *session.Session
	var uploader *s3manager.Uploader
	var downloader *s3manager.Downloader

	if v.GetBool("s3.enabled") {
		sess = session.Must(session.NewSession(
			&aws.Config{
				Region: aws.String(s3Region),
			},
		))

		uploader = s3manager.NewUploader(sess)
		downloader = s3manager.NewDownloader(sess)
	}

	// INIT DATABASE

	db, err := models.NewDatabase(v)
	if err != nil {
		return err
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(newContext(
				v,
				c,
				db,
				sess,
				uploader,
				downloader,
			))
		}
	})

	e.Static("/hls", "assets/hls")

	return e.Start(fmt.Sprintf("0.0.0.0:%s", v.GetString("server.port")))
}
