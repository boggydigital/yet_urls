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

func titleTextContent(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "title" {
		return false
	}

	return true
}

func extractJsonObject(data string) string {
	fi, li := strings.Index(data, opCuBrace), strings.LastIndex(data, clCuBrace)
	return data[fi : li+1]
}

func getDocument(videoId string) (*html.Node, error) {

	watchUrl := WatchUrl(videoId)

	resp, err := http.Get(watchUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func getBitrateSortedStreamingFormats(doc *html.Node) ([]string, error) {

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

//TitleStreamingUrls returns page title and streaming URLs sorted by bitrate for a given videoId
func TitleStreamingUrls(videoId string) (string, []*url.URL, error) {

	page, err := getDocument(videoId)
	if err != nil {
		return "", nil, err
	}

	title := ""
	titleText := match_node.Match(page, titleTextContent)
	if titleText != nil {
		title = titleText.Data
	}

	streamingUrls := make([]*url.URL, 0)

	streamingFormats, err := getBitrateSortedStreamingFormats(page)
	if err != nil {
		return title, streamingUrls, err
	}

	for _, sf := range streamingFormats {
		streamingUrl, err := url.Parse(sf)
		if err != nil {
			return title, streamingUrls, err
		}
		streamingUrls = append(streamingUrls, streamingUrl)
	}

	return title, streamingUrls, nil
}
