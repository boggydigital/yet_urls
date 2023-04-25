package yt_urls

import (
	"golang.org/x/net/html"
	"sort"
	"strings"
	"time"
)

const (
	ytInitialPlayerResponse = "var ytInitialPlayerResponse"
)

type initialPlayerResponseMatcher struct {
}

// iprScriptTextContent is an HTML node filter for YouTube <script> text content
// that contains ytInitialPlayerResponse data
func (iprm *initialPlayerResponseMatcher) Match(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytInitialPlayerResponse)
}

// InitialPlayerResponse is a minimal set of data structures required to decode and
// extract streaming data formats for video URL ytInitialPlayerResponse
type InitialPlayerResponse struct {
	PlayabilityStatus struct {
		Status      string `json:"status"`
		Reason      string `json:"reason"`
		ErrorScreen struct {
			PlayerErrorMessageRenderer struct {
				SubReason struct {
					SimpleText string `json:"simpleText"`
				} `json:"subreason"`
			} `json:"playerErrorMessageRenderer"`
		} `json:"errorScreen"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		ExpiresInSeconds string  `json:"expiresInSeconds"`
		Formats          Formats `json:"formats"`
		AdaptiveFormats  Formats `json:"adaptiveFormats"`
	} `json:"streamingData"`
	VideoDetails struct {
		VideoId   string `json:"videoId"`
		Title     string `json:"title"`
		ChannelId string `json:"channelId"`
	} `json:"videoDetails"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			PublishDate string `json:"publishDate"`
			UploadDate  string `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
}

func (ipr *InitialPlayerResponse) Title() string {
	return ipr.VideoDetails.Title
}

func (ipr *InitialPlayerResponse) Formats() Formats {
	sort.Sort(sort.Reverse(ipr.StreamingData.Formats))
	return ipr.StreamingData.Formats
}

func (ipr *InitialPlayerResponse) AdaptiveVideoFormats() Formats {
	vfs := ipr.StreamingData.AdaptiveFormats.PreferredVideo()
	sort.Sort(sort.Reverse(vfs))
	return vfs
}

func (ipr *InitialPlayerResponse) AdaptiveAudioFormats() Formats {
	afs := ipr.StreamingData.AdaptiveFormats.PreferredAudio()
	sort.Sort(sort.Reverse(afs))
	return afs
}

func (ipr *InitialPlayerResponse) ChannelId() string {
	return ipr.VideoDetails.ChannelId
}

func parseDateOrDefault(date string) time.Time {
	tm, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Now()
	}
	return tm
}

func (ipr *InitialPlayerResponse) PublishDate() time.Time {
	return parseDateOrDefault(ipr.Microformat.PlayerMicroformatRenderer.PublishDate)
}

func (ipr *InitialPlayerResponse) UploadDate() time.Time {
	return parseDateOrDefault(ipr.Microformat.PlayerMicroformatRenderer.UploadDate)
}
