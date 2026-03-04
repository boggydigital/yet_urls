package youtube_urls

import "net/url"

func HostUrl() *url.URL {
	return new(url.URL{
		Scheme: httpsScheme,
		Host:   YoutubeHost,
	})
}
