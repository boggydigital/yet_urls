package youtube_urls

import (
	"net/url"
	"path"
)

func ChannelUrl(channelId string) *url.URL {
	channelUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   youtubeWwwHost,
		Path:   path.Join(channelPath, channelId),
	}

	return channelUrl
}
