package youtube_urls

import "net/url"

const (
	playlistPath = "playlist"
	listParam    = "list"
)

// PlaylistUrl provides a URL for a playlist-id,
// e.g. http://www.youtube.com/playlist?list=playlist-id1 for "playlist-id1"
func PlaylistUrl(playlistId string) *url.URL {
	playlistUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   youtubeWwwHost,
		Path:   playlistPath,
	}

	q := playlistUrl.Query()
	q.Add(listParam, playlistId)
	playlistUrl.RawQuery = q.Encode()

	return playlistUrl
}

// PlaylistId extracts playlist-id from a PlaylistUrl conforming URL
func PlaylistId(ytUrlStr string) (string, error) {
	ytUrl, err := url.Parse(ytUrlStr)
	if err != nil {
		return ytUrlStr, err
	}

	q := ytUrl.Query()
	if q.Has(listParam) {
		return q.Get(listParam), nil
	} else {
		return ytUrlStr, nil
	}
}
