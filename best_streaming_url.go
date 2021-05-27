package yt_urls

import (
	"encoding/json"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
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

//bestFormatByBitrate selects "the best" available format for a video, based
//only on the max bitrate
func bestFormatByBitrate(formats []streamingFormat) *streamingFormat {

	if len(formats) == 0 {
		return nil
	}

	bitrateUrls := make(map[int]*streamingFormat, len(formats))

	maxBitrate := 0
	for _, format := range formats {
		bitrateUrls[format.Bitrate] = &format
		if format.Bitrate > maxBitrate {
			maxBitrate = format.Bitrate
		}
	}

	return bitrateUrls[maxBitrate]
}

//BestStreamingUrl extracts the URL for "the best" streaming format for a given
//YouTube video-id. Here are the key steps to make that happen:
//1) convert video-id to a full YouTube.com/watch URL
//2) request page content at that URL
//3) parse response as HTML document and find required node
//(iprScriptTextContent contains selection criteria)
//4) decode ytInitialPlayerResponse object (to a minimal data struct)
//5) select "the best" streaming format available
//(bestFormatByBitrate contains selection criteria)
//6) return URL for that format
func BestStreamingUrl(videoId string) (*url.URL, error) {

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

		bstFmt := bestFormatByBitrate(ipr.StreamingData.Formats)
		if bstFmt != nil {
			return url.Parse(bstFmt.Url)
		}
	}

	return nil, nil
}
