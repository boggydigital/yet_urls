package rutube_urls

import (
	"net/url"
	"path"
)

func PlayOptionsUrl(videoId, p string) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   RutubeHost,
		Path:   path.Join(apiPlayOptionsPath, videoId),
	}

	q := u.Query()
	q.Set("no_404", "true")
	if p != "" {
		q.Set("p", p)
	}
	q.Set("referer", "https://rutube.ru")
	q.Set("pver", "v2")

	u.RawQuery = q.Encode()

	return u
}
