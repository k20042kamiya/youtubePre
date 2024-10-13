package inputport

import "youtube/domain"

type YoutubeVideoList interface {
	FindYoutubeVideoList(apikey string) (YoutubeVideoList domain.Youtube, err error)
}
