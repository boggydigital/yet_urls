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

type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SimpleText struct {
	SimpleText string `json:"simpleText"`
}

// InitialPlayerResponse is a minimal set of data structures required to decode and
// extract streaming data formats for video URL ytInitialPlayerResponse
type InitialPlayerResponse struct {
	PlayabilityStatus struct {
		Status      string `json:"status"`
		Reason      string `json:"reason"`
		ErrorScreen struct {
			PlayerErrorMessageRenderer struct {
				SubReason SimpleText `json:"subreason"`
			} `json:"playerErrorMessageRenderer"`
		} `json:"errorScreen"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		ExpiresInSeconds string  `json:"expiresInSeconds"`
		Formats          Formats `json:"formats"`
		AdaptiveFormats  Formats `json:"adaptiveFormats"`
		HLSManifestUrl   string  `json:"hlsManifestUrl"`
	} `json:"streamingData"`
	VideoDetails struct {
		VideoId          string   `json:"videoId"`
		Title            string   `json:"title"`
		ChannelId        string   `json:"channelId"`
		Keywords         []string `json:"keywords"`
		ShortDescription string   `json:"shortDescription"`
		Thumbnail        struct {
			Thumbnails []Thumbnail `json:"thumbnails"`
		} `json:"thumbnail"`
		ViewCount string `json:"viewCount"`
		Author    string `json:"author"`
		IsPrivate bool   `json:"isPrivate"`
	} `json:"videoDetails"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			Thumbnail          Thumbnail  `json:"thumbnail"`
			Title              SimpleText `json:"title"`
			Description        SimpleText `json:"description"`
			OwnerProfileUrl    string     `json:"ownerProfileUrl"`
			OwnerChannelName   string     `json:"ownerChannelName"`
			ExternalChannelId  string     `json:"externalChannelId"`
			IsFamilySafe       bool       `json:"IsFamilySafe"`
			AvailableCountries []string   `json:"availableCountries"`
			IsUnlisted         bool       `json:"isUnlisted"`
			ViewCount          string     `json:"viewCount"`
			Category           string     `json:"category"`
			PublishDate        string     `json:"publishDate"`
			UploadDate         string     `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	Captions struct {
		PlayerCaptionsTracklistRenderer struct {
			CaptionTracks []struct {
				BaseUrl        string     `json:"baseUrl"`
				Name           SimpleText `json:"name"`
				VSSId          string     `json:"vssId"`
				LanguageCode   string     `json:"languageCode"`
				Kind           string     `json:"kind"`
				IsTranslatable bool       `json:"isTranslatable"`
				TrackName      string     `json:"trackName"`
			} `json:"captionTracks"`
			AudioTracks []struct {
				CaptionTrackIndices []int `json:"captionTrackIndices"`
			} `json:"audioTracks"`
			TranslationLanguages []struct {
				LanguageCode string     `json:"languageCode"`
				LanguageName SimpleText `json:"languageName"`
			} `json:"translationLanguages"`
			DefaultAudioTrackIndex int `json:"defaultAudioTrackIndex"`
		} `json:"playerCaptionsTracklistRenderer"`
	} `json:"captions"`
}

// TODO: Consider deprecating
func (ipr *InitialPlayerResponse) Formats() Formats {
	sort.Sort(sort.Reverse(ipr.StreamingData.Formats))
	return ipr.StreamingData.Formats
}

// TODO: Consider deprecating
func (ipr *InitialPlayerResponse) AdaptiveVideoFormats() Formats {
	vfs := ipr.StreamingData.AdaptiveFormats.PreferredVideo()
	sort.Sort(sort.Reverse(vfs))
	return vfs
}

// TODO: Consider deprecating
func (ipr *InitialPlayerResponse) AdaptiveAudioFormats() Formats {
	afs := ipr.StreamingData.AdaptiveFormats.PreferredAudio()
	sort.Sort(sort.Reverse(afs))
	return afs
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
