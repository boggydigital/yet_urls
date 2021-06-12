package yt_urls

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	iprPfx = "var ytInitialPlayerResponse = "
	iprSfx = ";"
)

//iprScriptTextContent is an HTML node filter for YouTube <script> text content
//that contains ytInitialPlayerResponse data
func iprScriptTextContent(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, iprPfx)
}

type streamingFormats []StreamingFormat

func (sf streamingFormats) Len() int {
	return len(sf)
}

func (sf streamingFormats) Less(i, j int) bool {
	return sf[i].Bitrate < sf[j].Bitrate
}

func (sf streamingFormats) Swap(i, j int) {
	sf[i], sf[j] = sf[j], sf[i]
}

func getBitrateSortedStreamingFormats(videoId string) (streamingFormats, error) {
	watchUrl := WatchUrl(videoId)

	resp, err := http.Get(watchUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	if iprNode := match_node.Match(doc, iprScriptTextContent); iprNode != nil {

		iprReader := strings.NewReader(
			strings.Trim(iprNode.Data, iprPfx+iprSfx))

		var ipr initialPlayerResponse
		if err := json.NewDecoder(iprReader).Decode(&ipr); err != nil {
			return nil, err
		}

		sfs := streamingFormats(ipr.StreamingData.Formats)
		sort.Sort(sort.Reverse(sfs))

		return sfs, nil
	}

	return nil, nil
}

func streamingFormatToUrl(sf *StreamingFormat) (*url.URL, error) {
	if sf.Url != "" {
		return url.Parse(sf.Url)
	} else if sf.SignatureCipher != "" {
		//TODO: support signature cipher YouTube URLs
		//https://stackoverflow.com/questions/21510857/best-approach-to-decode-youtube-cipher-signature-using-php-or-js
	}

	return nil, nil
}

//HighestBitrateStreamingUrl extracts the URL for "the best" streaming format for a given
//YouTube video-id. Here are the key steps to make that happen:
//1) convert video-id to a full YouTube.com/watch URL
//2) request page content at that URL
//3) parse response as HTML document and find required node
//(iprScriptTextContent contains selection criteria)
//4) decode ytInitialPlayerResponse object (to a minimal data struct)
//5) select "the best" streaming format available
//(bestFormatByBitrate contains selection criteria)
//6) return URL for that format
func HighestBitrateStreamingUrl(videoId string) (*url.URL, error) {

	streamingFormats, err := getBitrateSortedStreamingFormats(videoId)

	if err != nil {
		return nil, err
	}

	if len(streamingFormats) == 0 {
		return nil, fmt.Errorf("yt_url: no streaming formats detected")
	}

	return streamingFormatToUrl(&streamingFormats[0])
}

func BitrateSortedStreamingUrls(videoId string) ([]*url.URL, error) {
	streamingUrls := make([]*url.URL, 0)

	streamingFormats, err := getBitrateSortedStreamingFormats(videoId)
	if err != nil {
		return streamingUrls, nil
	}

	for _, sf := range streamingFormats {
		url, err := streamingFormatToUrl(&sf)
		if err != nil {
			return streamingUrls, err
		}
		streamingUrls = append(streamingUrls, url)
	}

	return streamingUrls, nil
}
