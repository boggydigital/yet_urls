package youtube_urls

type ChannelPlaylistsInitialData struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Title          string `json:"title"`
					TrackingParams string `json:"trackingParams"`
					Selected       bool   `json:"selected,omitempty"`
					Content        struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										GridRenderer struct {
											Items []struct {
												GridPlaylistRenderer GridPlaylistRenderer `json:"gridPlaylistRenderer"`
											} `json:"items"`
										} `json:"gridRenderer"`
									} `json:"contents"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
						} `json:"sectionListRenderer"`
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
	Context *ytCfgInnerTubeContext
}

type GridPlaylistRenderer struct {
	PlaylistId          string             `json:"playlistId"`
	Title               SimpleText         `json:"title"`
	VideoCountText      TextRuns           `json:"videoCountText"`
	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
	VideoCountShortText SimpleText         `json:"videoCountShortText"`
	ViewPlaylistText    TextRuns           `json:"viewPlaylistText"`
	PublishedTimeText   SimpleText         `json:"publishedTimeText,omitempty"`
}

func (cpid *ChannelPlaylistsInitialData) Playlists() []*GridPlaylistRenderer {
	playlists := make([]*GridPlaylistRenderer, 0)
	for _, tab := range cpid.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		for _, tabContent := range tab.TabRenderer.Content.SectionListRenderer.Contents {
			for _, itemContent := range tabContent.ItemSectionRenderer.Contents {
				for _, item := range itemContent.GridRenderer.Items {
					playlists = append(playlists, &item.GridPlaylistRenderer)
				}
			}
		}
	}
	return playlists
}
