package yt_urls

import (
	"golang.org/x/net/html"
	"strings"
)

const (
	ytInitialData = "var ytInitialData"
)

type initialDataScriptMatcher struct{}

// initialDataScript is an HTML node filter for YouTube <script> text content
// that contains ytInitialData
func (idsm *initialDataScriptMatcher) Match(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytInitialData)
}

type PlaylistHeaderRenderer struct {
	PlaylistId string `json:"playlistId"`
	Title      struct {
		SimpleText string `json:"simpleText"`
	} `json:"title"`
	PlaylistHeaderBanner struct {
		HeroPlaylistThumbnailRenderer struct {
			Thumbnail struct {
				Thumbnails []Thumbnail `json:"thumbnails"`
			} `json:"thumbnail"`
		} `json:"heroPlaylistThumbnailRenderer"`
	} `json:"playlistHeaderBanner"`
	DescriptionText SimpleText `json:"descriptionText"`
	OwnerText       struct {
		Runs []OwnerTextRun `json:"runs"`
	} `json:"ownerText"`
	ViewCountText SimpleText `json:"viewCountText"`
	Privacy       string     `json:"privacy"`
}

// PlaylistInitialData is a minimal set of data structures required to decode and
// extract videoIds for playlist URL ytInitialData
type PlaylistInitialData struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										PlaylistVideoListRenderer struct {
											Contents   []PlaylistVideoListRendererContent `json:"contents"`
											PlaylistId string                             `json:"playlistId"`
										} `json:"playlistVideoListRenderer"`
									} `json:"contents"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		PlaylistHeaderRenderer PlaylistHeaderRenderer `json:"playlistHeaderRenderer"`
	} `json:"header"`

	videoListContent []PlaylistVideoListRendererContent
}

type OwnerTextRun struct {
	Text               string `json:"text"`
	NavigationEndpoint struct {
		CommandMetadata struct {
			WebCommandMetadata struct {
				Url         string `json:"url"`
				WebPageType string `json:"webPageType"`
				RootVe      int    `json:"rootVe"`
				ApiUrl      string `json:"apiUrl"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
		BrowseEndpoint struct {
			BrowseId         string `json:"browseId"`
			CanonicalBaseUrl string `json:"canonicalBaseUrl"`
		} `json:"browseEndpoint"`
	} `json:"navigationEndpoint"`
}

type PlaylistVideoListRendererContent struct {
	PlaylistVideoRenderer    PlaylistVideoRenderer
	ContinuationItemRenderer ContinuationItemRenderer
}

type PlaylistVideoRenderer struct {
	VideoId string `json:"videoId"`
	Title   struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	// normally contains video channel title
	ShortBylineText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"shortBylineText"`
}

type ContinuationEndpoint struct {
	CommandMetadata struct {
		WebCommandMetadata struct {
			SendPost bool   `json:"sendPost"`
			ApiUrl   string `json:"apiUrl"`
		} `json:"webCommandMetadata"`
	} `json:"commandMetadata"`
	ContinuationCommand struct {
		Token   string `json:"token"`
		Request string `json:"request"`
	} `json:"continuationCommand"`
}

type ContinuationItemRenderer struct {
	Trigger              string               `json:"trigger"`
	ContinuationEndpoint ContinuationEndpoint `json:"continuationEndpoint"`
}

type VideoIdTitleChannel struct {
	VideoId string
	Title   string
	Channel string
}

func (id *PlaylistInitialData) PlaylistHeader() PlaylistHeaderRenderer {
	return id.Header.PlaylistHeaderRenderer
}

func (id *PlaylistInitialData) PlaylistContent() []PlaylistVideoListRendererContent {

	if id.videoListContent == nil {
		pvlc := make([]PlaylistVideoListRendererContent, 0)

		for _, tab := range id.Contents.TwoColumnBrowseResultsRenderer.Tabs {
			for _, sectionList := range tab.TabRenderer.Content.SectionListRenderer.Contents {
				for _, itemSection := range sectionList.ItemSectionRenderer.Contents {
					pvlc = append(pvlc, itemSection.PlaylistVideoListRenderer.Contents...)
				}
			}
		}

		id.videoListContent = pvlc
	}

	return id.videoListContent
}

func (id *PlaylistInitialData) SetContent(ct []PlaylistVideoListRendererContent) {
	id.videoListContent = ct
}

func (id *PlaylistInitialData) PlaylistOwner() string {
	ownerTextRuns := make([]string, 0, len(id.Header.PlaylistHeaderRenderer.OwnerText.Runs))
	for _, r := range id.Header.PlaylistHeaderRenderer.OwnerText.Runs {
		ownerTextRuns = append(ownerTextRuns, r.Text)
	}
	return strings.Join(ownerTextRuns, "")
}
