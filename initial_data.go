package yt_urls

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

//initialDataScript is an HTML node filter for YouTube <script> text content
//that contains ytInitialData
func initialDataScript(node *html.Node) bool {
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

type ContinuationItemRenderer struct {
	Trigger              string `json:"trigger"`
	ContinuationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				SendPost bool   `json:"sendPost"`
				ApiUrl   string `json:"apiUrl"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
		ContinuationCommand struct {
			Token   string `json:"token"`
			Request string `json:"request"`
		} `json:"continuationCommand"`
	} `json:"continuationEndpoint"`
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

func (id *InitialData) Videos() []VideoIdTitle {
	var vits []VideoIdTitle
	pvlc := id.playlistVideoListContent()
	vits = make([]VideoIdTitle, 0, len(pvlc))
	for _, vlc := range pvlc {
		videoId := vlc.PlaylistVideoRenderer.VideoId
		if videoId == "" {
			continue
		}
		title := vlc.PlaylistVideoRenderer.Title.Runs[0].Text
		for ii := 1; ii < len(vlc.PlaylistVideoRenderer.Title.Runs); ii++ {
			title += vlc.PlaylistVideoRenderer.Title.Runs[ii].Text
		}
		vits = append(vits, VideoIdTitle{
			VideoId: videoId,
			Title:   title,
		})
	}
	return vits
}

func (id *InitialData) NextPage() (*InitialData, error) {
	pvlc := id.playlistVideoListContent()
	for i := len(pvlc) - 1; i >= 0; i-- {
		vlc := pvlc[i]
		if vlc.ContinuationItemRenderer.Trigger == "" {
			continue
		}
		fmt.Println(vlc.ContinuationItemRenderer)
	}
	return nil, nil
}
