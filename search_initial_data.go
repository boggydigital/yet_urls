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
	ChannelId          string             `json:"channelId"`
	Title              SimpleText         `json:"title"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	DescriptionSnippet struct {
		Runs []struct {
			Text string `json:"text"`
			Bold bool   `json:"bold,omitempty"`
		} `json:"runs"`
	} `json:"descriptionSnippet"`
	ShortBylineText BylineText `json:"shortBylineText"`
	VideoCountText  struct {
		SimpleText string `json:"simpleText"`
	} `json:"videoCountText"`
	SubscriberCountText SimpleText `json:"subscriberCountText"`
	LongBylineText      BylineText `json:"longBylineText"`
}

type PlaylistRenderer struct {
	PlaylistId         string             `json:"playlistId"`
	Title              SimpleText         `json:"title"`
	VideoCount         string             `json:"videoCount"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	ViewPlaylistText   struct {
		Runs []struct {
			Text               string             `json:"text"`
			NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"viewPlaylistText"`
	ShortBylineText BylineText `json:"shortBylineText"`
	Videos          []struct {
		ChildVideoRenderer struct {
			Title              SimpleText         `json:"title"`
			NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
			LengthText         SimpleText         `json:"lengthText"`
			VideoId            string             `json:"videoId"`
		} `json:"childVideoRenderer"`
	} `json:"videos"`
	VideoCountText struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"videoCountText"`
	TrackingParams string `json:"trackingParams"`
	ThumbnailText  struct {
		Runs []struct {
			Text string `json:"text"`
			Bold bool   `json:"bold,omitempty"`
		} `json:"runs"`
	} `json:"thumbnailText"`
	LongBylineText BylineText `json:"longBylineText"`
}

type VideoRenderer struct {
	VideoId string `json:"videoId"`
	Title   struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
	} `json:"title"`
	LongBylineText     BylineText         `json:"longBylineText"`
	PublishedTimeText  SimpleText         `json:"publishedTimeText"`
	LengthText         SimpleText         `json:"lengthText"`
	ViewCountText      SimpleText         `json:"viewCountText"`
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	OwnerText          struct {
		Runs []struct {
			Text               string             `json:"text"`
			NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"ownerText"`
	ShortBylineText    BylineText `json:"shortBylineText"`
	ShowActionMenu     bool       `json:"showActionMenu"`
	ShortViewCountText SimpleText `json:"shortViewCountText"`
	IsWatched          bool       `json:"isWatched,omitempty"`
}

type BylineText struct {
	Runs []struct {
		Text               string             `json:"text"`
		NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
	} `json:"runs"`
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
