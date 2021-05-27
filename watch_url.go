package yt_urls

import (
	"net/url"
)

//WatchUrl provides a URL for a video-id,
//e.g. http://www.youtube.com/watch?v=video-id1 for "video-id1"
func WatchUrl(videoId string) *url.URL {
	watchUrl := &url.URL{
		Scheme: "https",
		Host:   "www.youtube.com",
		Path:   "watch",
	}

	q := watchUrl.Query()
	q.Add("v", videoId)
	watchUrl.RawQuery = q.Encode()

	return watchUrl
}
