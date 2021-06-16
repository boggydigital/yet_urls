package yt_urls

import (
	"encoding/json"
	"fmt"
	"github.com/boggydigital/gost"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
)

const (
	ytInitialPlayerResponse = "var ytInitialPlayerResponse ="
	opCuBrace               = "{"
	clCuBrace               = "}"
)

//iprScriptTextContent is an HTML node filter for YouTube <script> text content
//that contains ytInitialPlayerResponse data
func iprScriptTextContent(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytInitialPlayerResponse)
}

func extractJsonObject(data string) string {
	fi, li := strings.Index(data, opCuBrace), strings.LastIndex(data, clCuBrace)
	return data[fi : li+1]
}

func getBitrateSortedStreamingFormats(videoId string) ([]string, error) {
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

		iprReader := strings.NewReader(extractJsonObject(iprNode.Data))

		var ipr initialPlayerResponse
		if err := json.NewDecoder(iprReader).Decode(&ipr); err != nil {
			return nil, err
		}

		if ipr.PlayabilityStatus.Status != StatusOK {
			return nil, fmt.Errorf(ipr.PlayabilityStatus.Status)
		}

		signatureCipher := false

		formats := make(map[string]int, len(ipr.StreamingData.Formats))
		for _, f := range ipr.StreamingData.Formats {
			if f.Url == "" && f.SignatureCipher != "" {
				signatureCipher = true
				continue
			}
			formats[f.Url] = f.Bitrate
		}

		if len(formats) == 0 && signatureCipher {
			//TODO: support signature cipher YouTube URLs
			//https://stackoverflow.com/questions/21510857/best-approach-to-decode-youtube-cipher-signature-using-php-or-js
			return nil, fmt.Errorf(ErrorSignatureCipher)
		}

		_, sortedFormats := gost.NewIntSortedStrSetWith(formats, true)

		return sortedFormats, nil
	}

	return nil, nil
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

	streamingFormats, err := getBitrateSortedStreamingFormats(videoId)

	if err != nil {
		return nil, err
	}

	if len(streamingFormats) == 0 {
		return nil, fmt.Errorf("yt_url: no streaming formats detected")
	}

	return url.Parse(streamingFormats[0])
}

func StreamingUrls(videoId string) ([]*url.URL, error) {
	streamingUrls := make([]*url.URL, 0)

	streamingFormats, err := getBitrateSortedStreamingFormats(videoId)
	if err != nil {
		return streamingUrls, err
	}

	for _, sf := range streamingFormats {
		streamingUrl, err := url.Parse(sf)
		if err != nil {
			return streamingUrls, err
		}
		streamingUrls = append(streamingUrls, streamingUrl)
	}

	return streamingUrls, nil
}
