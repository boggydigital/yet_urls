package yet_urls

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

type CaptionTrack struct {
	BaseUrl        string     `json:"baseUrl"`
	Name           SimpleText `json:"name"`
	VSSId          string     `json:"vssId"`
	LanguageCode   string     `json:"languageCode"`
	Kind           string     `json:"kind"`
	IsTranslatable bool       `json:"isTranslatable"`
	TrackName      string     `json:"trackName"`
}

// InitialPlayerResponse is a minimal set of data structures required to decode and
// extract streaming data formats for video URL ytInitialPlayerResponse
type InitialPlayerResponse struct {
	PlayerUrl         string
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
		LengthSeconds    string   `json:"lengthSeconds"`
		Keywords         []string `json:"keywords"`
		ChannelId        string   `json:"channelId"`
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
			Thumbnail struct {
				Thumbnails []Thumbnail `json:"thumbnails"`
			} `json:"thumbnail"`
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
			CaptionTracks []CaptionTrack `json:"captionTracks"`
			AudioTracks   []struct {
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

func (ipr *InitialPlayerResponse) BestFormat() *Format {

	formats := ipr.StreamingData.Formats

	if len(formats) == 0 {
		return nil
	}

	qualityIndex := make(map[string]int)

	for ii, ff := range formats {
		qualityIndex[ff.QualityLabel] = ii
	}

	qualityOrder := []string{"2160p", "1440p", "1080p", "720p"}
	bestIndex := -1
	for _, q := range qualityOrder {
		if ii, ok := qualityIndex[q]; ok {
			bestIndex = ii
		}
	}

	if bestIndex == -1 && len(formats) > 0 {
		// use the first available if none of the best quality formats are present
		bestIndex = 0
	}

	return formats[bestIndex]
}

func (ipr *InitialPlayerResponse) BestAdaptiveVideoFormat() *Format {

	if len(ipr.StreamingData.AdaptiveFormats) == 0 {
		return nil
	}

	vfs := ipr.StreamingData.AdaptiveFormats.PreferredVideoFormats()
	sort.Sort(sort.Reverse(vfs))

	return vfs[0]
}

func (ipr *InitialPlayerResponse) BestAdaptiveAudioFormat() *Format {

	if len(ipr.StreamingData.AdaptiveFormats) == 0 {
		return nil
	}

	afs := ipr.StreamingData.AdaptiveFormats.PreferredAudioFormats()
	sort.Sort(sort.Reverse(afs))

	return afs[0]
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

func (ipr *InitialPlayerResponse) SignatureCipher() bool {
	for _, f := range ipr.StreamingData.Formats {
		if f.SignatureCipher != "" {
			return true
		}
	}
	for _, f := range ipr.StreamingData.AdaptiveFormats {
		if f.SignatureCipher != "" {
			return true
		}
	}
	return false
}
