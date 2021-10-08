package yt_urls

import "net/url"

//VideoId extracts video-id from a WatchUrl conforming URL
func VideoId(ytUrlStr string) (string, error) {
	ytUrl, err := url.Parse(ytUrlStr)
	if err != nil {
		return ytUrlStr, err
	}

	q := ytUrl.Query()
	if q.Has("v") {
		return q.Get("v"), nil
	} else {
		return ytUrlStr, nil
	}
}
