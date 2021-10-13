package yt_urls

import "net/url"

const (
	keyParam    = "key"
	contentType = "application/json"
)

func BrowseUrl(path, apiKey string) *url.URL {
	browseUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   youtubeHost,
		Path:   path,
	}

	q := browseUrl.Query()
	q.Add(keyParam, apiKey)
	browseUrl.RawQuery = q.Encode()

	return browseUrl
}

type browseRequest struct {
	Context      ytCfgContext `json:"context"`
	Continuation string       `json:"continuation"`
}

type browseResponse struct {
	OnResponseReceivedActions []struct {
		AppendContinuationItemsAction struct {
			ContinuationItems []PlaylistVideoListRendererContent `json:"continuationItems"`
			TargetId          string                             `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedActions"`
}
