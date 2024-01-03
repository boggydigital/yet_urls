package yt_urls

type SearchInitialData struct {
	EstimatedResults string `json:"estimatedResults"`
	Contents         struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []struct {
								ChannelRenderer  ChannelRenderer  `json:"channelRenderer,omitempty"`
								PlaylistRenderer PlaylistRenderer `json:"playlistRenderer,omitempty"`
								VideoRenderer    VideoRenderer    `json:"videoRenderer,omitempty"`
							} `json:"contents"`
						} `json:"itemSectionRenderer,omitempty"`
						ContinuationItemRenderer ContinuationItemRenderer `json:"continuationItemRenderer,omitempty"`
					} `json:"contents"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
	Refinements []string `json:"refinements"`
	Context     *ytCfgInnerTubeContext
}

type ChannelRenderer struct {
	ChannelId           string             `json:"channelId"`
	Title               SimpleText         `json:"title"`
	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
	DescriptionSnippet  TextRuns           `json:"descriptionSnippet"`
	ShortBylineText     TextRuns           `json:"shortBylineText"`
	VideoCountText      SimpleText         `json:"videoCountText"`
	SubscriberCountText SimpleText         `json:"subscriberCountText"`
	LongBylineText      TextRuns           `json:"longBylineText"`
}

type PlaylistRenderer struct {
	PlaylistId         string             `json:"playlistId"`
	Title              SimpleText         `json:"title"`
	VideoCount         string             `json:"videoCount"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	ViewPlaylistText   TextRuns           `json:"viewPlaylistText"`
	ShortBylineText    TextRuns           `json:"shortBylineText"`
	Videos             []struct {
		ChildVideoRenderer struct {
			Title              SimpleText         `json:"title"`
			NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
			LengthText         SimpleText         `json:"lengthText"`
			VideoId            string             `json:"videoId"`
		} `json:"childVideoRenderer"`
	} `json:"videos"`
	VideoCountText TextRuns `json:"videoCountText"`
	TrackingParams string   `json:"trackingParams"`
	ThumbnailText  TextRuns `json:"thumbnailText"`
	LongBylineText TextRuns `json:"longBylineText"`
}

type VideoRenderer struct {
	VideoId            string             `json:"videoId"`
	Title              TextRuns           `json:"title"`
	LongBylineText     TextRuns           `json:"longBylineText"`
	PublishedTimeText  SimpleText         `json:"publishedTimeText"`
	LengthText         SimpleText         `json:"lengthText"`
	ViewCountText      SimpleText         `json:"viewCountText"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	OwnerText          TextRuns           `json:"ownerText"`
	ShortBylineText    TextRuns           `json:"shortBylineText"`
	ShortViewCountText SimpleText         `json:"shortViewCountText"`
	IsWatched          bool               `json:"isWatched,omitempty"`
}

type NavigationEndpoint struct {
	ClickTrackingParams string `json:"clickTrackingParams"`
	CommandMetadata     struct {
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
	WatchEndpoint struct {
		VideoId    string `json:"videoId"`
		PlaylistId string `json:"playlistId"`
	}
}

func (sid *SearchInitialData) ChannelRenderers() []ChannelRenderer {
	chrs := make([]ChannelRenderer, 0)

	for _, content := range sid.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, isrc := range content.ItemSectionRenderer.Contents {
			if isrc.ChannelRenderer.ChannelId == "" {
				continue
			}
			chrs = append(chrs, isrc.ChannelRenderer)
		}
	}
	return chrs
}

func (sid *SearchInitialData) PlaylistRenderers() []PlaylistRenderer {
	plrs := make([]PlaylistRenderer, 0)

	for _, content := range sid.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, isrc := range content.ItemSectionRenderer.Contents {
			if isrc.PlaylistRenderer.PlaylistId == "" {
				continue
			}
			plrs = append(plrs, isrc.PlaylistRenderer)
		}
	}
	return plrs
}

func (sid *SearchInitialData) VideoRenderers() []VideoRenderer {
	vrs := make([]VideoRenderer, 0)

	for _, content := range sid.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, isrc := range content.ItemSectionRenderer.Contents {
			if isrc.VideoRenderer.VideoId == "" {
				continue
			}
			vrs = append(vrs, isrc.VideoRenderer)
		}
	}

	return vrs
}
