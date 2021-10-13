package yt_urls

import (
	"golang.org/x/net/html"
	"sort"
	"strings"
)

const (
	ytInitialPlayerResponse = "var ytInitialPlayerResponse"
)

//iprScriptTextContent is an HTML node filter for YouTube <script> text content
//that contains ytInitialPlayerResponse data
func initialPlayerResponseScript(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytInitialPlayerResponse)
}

//InitialPlayerResponse is a minimal set of data structures required to decode and
//extract streaming data formats for video URL ytInitialPlayerResponse
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
