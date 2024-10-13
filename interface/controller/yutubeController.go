package controller

import (
	"youtube/usecase/inputport"

	"github.com/labstack/echo"
)

type YoutubeContololler struct {
	youtubeVideoList inputport.YoutubeVideoList
}

func NewYoutubeContololler(YoutubeVideoList inputport.YoutubeVideoList) *YoutubeContololler {
	return &YoutubeContololler{
		youtubeVideoList: YoutubeVideoList,
	}
}

func (contololler *YoutubeContololler) getVideos(c echo.Context) {
	query := c.Param("query")
	youtubeVideos, err := contololler.youtubeVideoList.FindYoutubeVideoList(query)
}
