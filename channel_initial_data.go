package yt_urls

type ChannelInitialData struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer ItemSectionRenderer `json:"itemSectionRenderer"`
							} `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Metadata struct {
		ChannelMetadataRenderer struct {
			Title       string   `json:"title"`
			Description string   `json:"description"`
			RssUrl      string   `json:"rssUrl"`
			ExternalId  string   `json:"externalId"`
			Keywords    string   `json:"keywords"`
			OwnerUrls   []string `json:"ownerUrls"`
			Avatar      struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"avatar"`
			ChannelUrl            string   `json:"channelUrl"`
			IsFamilySafe          bool     `json:"isFamilySafe"`
			AvailableCountryCodes []string `json:"availableCountryCodes"`
			VanityChannelUrl      string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
	Context *ytCfgInnerTubeContext
}

type ItemSectionRenderer struct {
	Contents []struct {
		ShelfRenderer struct {
			Title   TextRuns `json:"title"`
			Content struct {
				HorizontalListRenderer struct {
					Items []struct {
						GridVideoRenderer    GridVideoRenderer    `json:"gridVideoRenderer,omitempty"`
						GridPlaylistRenderer GridPlaylistRenderer `json:"gridPlaylistRenderer,omitempty"`
						GridChannelRenderer  GridChannelRenderer  `json:"gridChannelRenderer,omitempty"`
					} `json:"items"`
				} `json:"horizontalListRenderer"`
			} `json:"content"`
		} `json:"shelfRenderer,omitempty"`
	} `json:"contents"`
}

type GridVideoRenderer struct {
	VideoId            string             `json:"videoId"`
	Title              TextRuns           `json:"title"`
	PublishedTimeText  SimpleText         `json:"publishedTimeText"`
	ViewCountText      SimpleText         `json:"viewCountText"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	ShortViewCountText SimpleText         `json:"shortViewCountText"`
	ShortBylineText    TextRuns           `json:"shortBylineText,omitempty"`
}

type GridPlaylistRenderer struct {
	PlaylistId          string             `json:"playlistId"`
	Title               TextRuns           `json:"title"`
	VideoCountText      TextRuns           `json:"videoCountText"`
	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
	VideoCountShortText SimpleText         `json:"videoCountShortText"`
	ViewPlaylistText    TextRuns           `json:"viewPlaylistText"`
	PublishedTimeText   SimpleText         `json:"publishedTimeText,omitempty"`
}

type GridChannelRenderer struct {
	ChannelId           string             `json:"channelId"`
	VideoCountText      TextRuns           `json:"videoCountText"`
	SubscriberCountText SimpleText         `json:"subscriberCountText"`
	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
	Title               SimpleText         `json:"title"`
}
