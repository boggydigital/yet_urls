package youtube_urls

import (
	"net/url"
	"path"
)

func ChannelVideosUrl(channelId string) *url.URL {
	channelUrl := ChannelUrl(channelId)
	channelUrl.Path = path.Join(channelUrl.Path, channelVideosPath)
	return channelUrl
}
