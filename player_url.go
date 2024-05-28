package yet_urls

import "net/url"

func PlayerUrl(path string) *url.URL {
	playerUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   youtubeWwwHost,
		Path:   path,
	}

	return playerUrl
}
