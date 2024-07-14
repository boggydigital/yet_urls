package youtube_urls

import (
	"net/url"
	"path"
)

func ChannelPlaylistsUrl(channelId string) *url.URL {
	channelUrl := ChannelUrl(channelId)
	channelUrl.Path = path.Join(channelUrl.Path, channelPlaylistsPath)
	return channelUrl
}
