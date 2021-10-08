package yt_urls

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"net/http"
	"sort"
	"strings"
)

const (
	StatusOK                = "OK"
	ytInitialPlayerResponse = "var ytInitialPlayerResponse ="
	opCuBrace               = "{"
	clCuBrace               = "}"
)

var ErrorSignatureCipher = errors.New("signatureCipher")

//InitialPlayerResponse is a minimal set of data structures required to decode and
//extract streaming data formats
type InitialPlayerResponse struct {
	PlayabilityStatus struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		//use of Formats and not AdaptiveFormats is intentional, even though the former seems
		//to be capped at 720p. AdaptiveFormats come as a separate video and audio tracks and
		//would require merging those two together.
		//Formats on the other hand contain URLs to files that contain both video and audio.
		//If you have a need for something more complex or flexible - you should consider
		//youtube-dl or any of the alternatives available
		Formats StreamingFormats `json:"formats"`
	} `json:"streamingData"`
	VideoDetails struct {
		VideoId   string `json:"videoId"`
		Title     string `json:"title"`
		ChannelId string `json:"channelId"`
	} `json:"videoDetails"`
}

func extractJsonObject(data string) string {
	fi, li := strings.Index(data, opCuBrace), strings.LastIndex(data, clCuBrace)
	return data[fi : li+1]
}

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

func GetVideoPage(videoId string) (*InitialPlayerResponse, error) {
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

	var ipr InitialPlayerResponse

	if iprNode := match_node.Match(doc, iprScriptTextContent); iprNode != nil {

		iprReader := strings.NewReader(extractJsonObject(iprNode.Data))

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
			return nil, ErrorSignatureCipher
		}
	}

	return &ipr, nil
}

func (ipr *InitialPlayerResponse) Title() string {
	return ipr.VideoDetails.Title
}

func (ipr *InitialPlayerResponse) StreamingFormats() StreamingFormats {
	sort.Sort(sort.Reverse(ipr.StreamingData.Formats))
	return ipr.StreamingData.Formats
}

func (ipr *InitialPlayerResponse) ChannelId() string {
	return ipr.VideoDetails.ChannelId
}
