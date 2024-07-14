package youtube_urls

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ChannelVideosInitialData struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			DatasyncId    string `json:"datasyncId"`
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				SessionIndex          int    `json:"sessionIndex"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Title    string `json:"title"`
					Selected bool   `json:"selected,omitempty"`
					Content  struct {
						RichGridRenderer struct {
							Contents []RichGridRendererContents `json:"contents"`
						} `json:"richGridRenderer"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Metadata struct {
		ChannelMetadataRenderer struct {
			Title                 string   `json:"title"`
			Description           string   `json:"description"`
			RssUrl                string   `json:"rssUrl"`
			ChannelConversionUrl  string   `json:"channelConversionUrl"`
			ExternalId            string   `json:"externalId"`
			Keywords              string   `json:"keywords"`
			OwnerUrls             []string `json:"ownerUrls"`
			ChannelUrl            string   `json:"channelUrl"`
			IsFamilySafe          bool     `json:"isFamilySafe"`
			AvailableCountryCodes []string `json:"availableCountryCodes"`
			VanityChannelUrl      string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
	Microformat struct {
		MicroformatDataRenderer struct {
			UrlCanonical string `json:"urlCanonical"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			Thumbnail    struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			SiteName           string   `json:"siteName"`
			AppName            string   `json:"appName"`
			AndroidPackage     string   `json:"androidPackage"`
			IosAppStoreId      string   `json:"iosAppStoreId"`
			IosAppArguments    string   `json:"iosAppArguments"`
			OgType             string   `json:"ogType"`
			UrlApplinksWeb     string   `json:"urlApplinksWeb"`
			UrlApplinksIos     string   `json:"urlApplinksIos"`
			UrlApplinksAndroid string   `json:"urlApplinksAndroid"`
			UrlTwitterIos      string   `json:"urlTwitterIos"`
			UrlTwitterAndroid  string   `json:"urlTwitterAndroid"`
			TwitterCardType    string   `json:"twitterCardType"`
			TwitterSiteHandle  string   `json:"twitterSiteHandle"`
			SchemaDotOrgType   string   `json:"schemaDotOrgType"`
			Noindex            bool     `json:"noindex"`
			Unlisted           bool     `json:"unlisted"`
			FamilySafe         bool     `json:"familySafe"`
			Tags               []string `json:"tags"`
			AvailableCountries []string `json:"availableCountries"`
			LinkAlternates     []struct {
				HrefUrl string `json:"hrefUrl"`
			} `json:"linkAlternates"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
	videosContent []RichGridRendererContents
	Context       *ytCfgInnerTubeContext
}

type RichGridRendererContents struct {
	RichItemRenderer struct {
		Content struct {
			VideoRenderer struct {
				VideoId            string   `json:"videoId"`
				Title              TextRuns `json:"title"`
				DescriptionSnippet struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"descriptionSnippet"`
				PublishedTimeText struct {
					SimpleText string `json:"simpleText"`
				} `json:"publishedTimeText"`
				LengthText struct {
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					SimpleText string `json:"simpleText"`
				} `json:"lengthText"`
				ViewCountText struct {
					SimpleText string `json:"simpleText"`
				} `json:"viewCountText"`
				IsWatched bool `json:"isWatched,omitempty"`
			} `json:"videoRenderer"`
		} `json:"content"`
	} `json:"richItemRenderer,omitempty"`
	ContinuationItemRenderer ContinuationItemRenderer `json:"continuationItemRenderer,omitempty"`
}

type channelVideosBrowseResponse struct {
	OnResponseReceivedActions []struct {
		AppendContinuationItemsAction struct {
			ContinuationItems []RichGridRendererContents `json:"continuationItems"`
			TargetId          string                     `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedActions"`
}

func (cvid *ChannelVideosInitialData) VideosContent() []RichGridRendererContents {
	if cvid.videosContent == nil {
		vc := make([]RichGridRendererContents, 0)

		for _, tab := range cvid.Contents.TwoColumnBrowseResultsRenderer.Tabs {
			for _, sectionList := range tab.TabRenderer.Content.RichGridRenderer.Contents {
				vc = append(vc, sectionList)
			}
		}

		cvid.videosContent = vc
	}

	return cvid.videosContent
}

func (cvid *ChannelVideosInitialData) SetContent(ct []RichGridRendererContents) {
	cvid.videosContent = ct
}

func (cvid *ChannelVideosInitialData) HasContinuation() bool {
	vc := cvid.VideosContent()
	for i := len(vc) - 1; i >= 0; i-- {
		if vc[i].ContinuationItemRenderer.Trigger != "" {
			return true
		}
	}
	return false
}

func (cvid *ChannelVideosInitialData) continuationEndpoint() *ContinuationEndpoint {
	pc := cvid.VideosContent()
	for i := len(pc) - 1; i >= 0; i-- {
		if pc[i].ContinuationItemRenderer.Trigger != "" {
			return &pc[i].ContinuationItemRenderer.ContinuationEndpoint
		}
	}
	return nil
}

func (cvid *ChannelVideosInitialData) Continue(client *http.Client) error {

	if !cvid.HasContinuation() {
		return nil
	}

	contEndpoint := cvid.continuationEndpoint()

	data := browseRequest{
		Context:      cvid.Context.InnerTubeContext,
		Continuation: contEndpoint.ContinuationCommand.Token,
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
		return err
	}

	browseUrl := BrowseUrl(
		contEndpoint.CommandMetadata.WebCommandMetadata.ApiUrl,
		cvid.Context.APIKey)

	resp, err := client.Post(browseUrl.String(), contentType, b)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	var br channelVideosBrowseResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return err
	}

	// update contents internals
	cvid.SetContent(br.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems)

	return nil
}

func (cvid *ChannelVideosInitialData) Videos() []VideoIdTitleLengthChannel {
	var vits []VideoIdTitleLengthChannel
	pc := cvid.VideosContent()
	vits = make([]VideoIdTitleLengthChannel, 0, len(pc))
	for _, vlc := range pc {
		videoId := vlc.RichItemRenderer.Content.VideoRenderer.VideoId
		if videoId == "" {
			continue
		}
		title, titleRuns := "", vlc.RichItemRenderer.Content.VideoRenderer.Title.Runs
		for _, r := range titleRuns {
			title += r.Text
		}

		lengthSeconds := lengthTextToSeconds(vlc.RichItemRenderer.Content.VideoRenderer.LengthText.SimpleText)

		vits = append(vits, VideoIdTitleLengthChannel{
			VideoId:       videoId,
			Title:         title,
			LengthSeconds: lengthSeconds,
		})
	}
	return vits
}

func lengthTextToSeconds(lt string) string {
	seconds := int64(0)
	parts := strings.Split(lt, ":")
	if len(parts) > 0 {
		if si, err := strconv.ParseInt(parts[len(parts)-1], 10, 64); err == nil {
			seconds += si
		}
		if len(parts) > 1 {
			if mi, err := strconv.ParseInt(parts[len(parts)-2], 10, 64); err == nil {
				seconds += mi * 60
			}
			if len(parts) > 2 {
				if hi, err := strconv.ParseInt(parts[len(parts)-3], 10, 64); err == nil {
					seconds += hi * 60 * 60
				}
				if len(parts) > 3 {
					if di, err := strconv.ParseInt(parts[len(parts)-4], 10, 64); err == nil {
						seconds += di * 60 * 60 * 24
					}
				}
			}
		}
	}
	return strconv.FormatInt(seconds, 10)
}
