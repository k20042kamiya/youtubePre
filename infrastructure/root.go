package infrastructure

import (
	"youtube/interface/controller"

	"youtube/usecase/inputport"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	youtubeContololler := controller.NewYoutubeContololler(inputport.YoutubeVideoList)
	e.GET("/video", func(c echo.Context) error {
		return youtubeContololler.getVideos(c)
	})
	// e.GET("/videos", func(c echo.Context) error {
	// 	return youtubeContololler.getVideos(c)
	// })

}
