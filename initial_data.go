package yt_urls

import (
	"golang.org/x/net/html"
	"strings"
)

const (
	ytInitialData = "var ytInitialData"
)

type initialDataScriptMatcher struct{}

//initialDataScript is an HTML node filter for YouTube <script> text content
//that contains ytInitialData
func (idsm *initialDataScriptMatcher) Match(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytInitialData)
}

//InitialData is a minimal set of data structures required to decode and
//extract videoIds for playlist URL ytInitialData
type InitialData struct {
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

type VideoIdTitle struct {
	VideoId string
	Title   string
}

func (id *InitialData) playlistVideoListContent() []PlaylistVideoListRendererContent {
	pvlc := make([]PlaylistVideoListRendererContent, 0)
	for _, tab := range id.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		for _, sectionList := range tab.TabRenderer.Content.SectionListRenderer.Contents {
			for _, itemSection := range sectionList.ItemSectionRenderer.Contents {
				pvlc = append(pvlc, itemSection.PlaylistVideoListRenderer.Contents...)
			}
		}
	}
	return pvlc
}
