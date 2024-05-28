package youtube_urls

import (
	"net/url"
	"strings"
)

const (
	searchQueryParam = "search_query"
)

// SearchResultsUrl provides a URL for a set of search terms,
// e.g. http://www.youtube.com/results?search_query=some+terms
func SearchResultsUrl(terms ...string) *url.URL {
	resultsUrl := &url.URL{
		Scheme: httpsScheme,
		Host:   youtubeWwwHost,
		Path:   resultsPath,
	}

	q := resultsUrl.Query()
	q.Add(searchQueryParam, strings.Join(terms, "+"))
	resultsUrl.RawQuery = q.Encode()

	return resultsUrl
}
