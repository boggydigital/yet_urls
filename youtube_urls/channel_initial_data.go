package youtube_urls

//	type ChannelInitialData struct {
//		Contents struct {
//			TwoColumnBrowseResultsRenderer struct {
//				Tabs []Tab `json:"tabs"`
//			} `json:"twoColumnBrowseResultsRenderer"`
//		} `json:"contents"`
//		Metadata struct {
//			ChannelMetadataRenderer ChannelMetadataRenderer `json:"channelMetadataRenderer"`
//		} `json:"metadata"`
//		Context *ytCfgInnerTubeContext
//	}
//
//	type ChannelMetadataRenderer struct {
//		Title                 string   `json:"title"`
//		Description           string   `json:"description"`
//		RSSUrl                string   `json:"rssUrl"`
//		ExternalId            string   `json:"externalId"`
//		Keywords              string   `json:"keywords"`
//		OwnerUrls             []string `json:"ownerUrls"`
//		ChannelUrl            string   `json:"channelUrl"`
//		IsFamilySafe          bool     `json:"isFamilySafe"`
//		AvailableCountryCodes []string `json:"availableCountryCodes"`
//		VanityChannelUrl      string   `json:"vanityChannelUrl"`
//	}
//
//	type ItemSectionRenderer struct {
//		Contents []struct {
//			ShelfRenderer ShelfRenderer `json:"shelfRenderer,omitempty"`
//		} `json:"contents"`
//	}
//
//	type ShelfRenderer struct {
//		Title   TextRuns `json:"title"`
//		Content struct {
//			HorizontalListRenderer struct {
//				Items []struct {
//					GridVideoRenderer    GridVideoRenderer    `json:"gridVideoRenderer,omitempty"`
//					GridPlaylistRenderer GridPlaylistRenderer `json:"gridPlaylistRenderer,omitempty"`
//					GridChannelRenderer  GridChannelRenderer  `json:"gridChannelRenderer,omitempty"`
//				} `json:"items"`
//			} `json:"horizontalListRenderer"`
//		} `json:"content"`
//		PlayAllButton struct {
//			ButtonRenderer struct {
//				Style      string `json:"style"`
//				Size       string `json:"size"`
//				IsDisabled bool   `json:"isDisabled"`
//				Text       struct {
//					Runs []struct {
//						Text string `json:"text"`
//					} `json:"runs"`
//				} `json:"text"`
//				Icon struct {
//					IconType string `json:"iconType"`
//				} `json:"icon"`
//				NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
//			} `json:"buttonRenderer"`
//		} `json:"playAllButton"`
//	}
//
//	type GridVideoRenderer struct {
//		VideoId            string             `json:"videoId"`
//		Title              SimpleText         `json:"title"`
//		PublishedTimeText  SimpleText         `json:"publishedTimeText"`
//		ViewCountText      SimpleText         `json:"viewCountText"`
//		NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
//		ShortViewCountText SimpleText         `json:"shortViewCountText"`
//		ShortBylineText    TextRuns           `json:"shortBylineText,omitempty"`
//	}
//
//type GridChannelRenderer struct {
//	ChannelId           string             `json:"channelId"`
//	VideoCountText      TextRuns           `json:"videoCountText"`
//	SubscriberCountText SimpleText         `json:"subscriberCountText"`
//	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
//	Title               SimpleText         `json:"title"`
//}
//
//type Tab struct {
//	TabRenderer struct {
//		Content struct {
//			SectionListRenderer struct {
//				Contents []struct {
//					ItemSectionRenderer ItemSectionRenderer `json:"itemSectionRenderer"`
//				} `json:"contents"`
//			} `json:"sectionListRenderer"`
//		} `json:"content,omitempty"`
//	} `json:"tabRenderer,omitempty"`
//}
//
//func (cid *ChannelInitialData) ChannelMetadataRenderer() *ChannelMetadataRenderer {
//	return &cid.Metadata.ChannelMetadataRenderer
//}
//
//func (cid *ChannelInitialData) Tabs() []Tab {
//	return cid.Contents.TwoColumnBrowseResultsRenderer.Tabs
//}
//
//func (tab *Tab) Sections() []ShelfRenderer {
//
//	srs := make([]ShelfRenderer, 0)
//
//	for _, oc := range tab.TabRenderer.Content.SectionListRenderer.Contents {
//		for _, ic := range oc.ItemSectionRenderer.Contents {
//			if ic.ShelfRenderer.Title.String() == "" {
//				continue
//			}
//			srs = append(srs, ic.ShelfRenderer)
//		}
//	}
//
//	return srs
//}
//
//func (sr *ShelfRenderer) PlaylistId() string {
//	return sr.PlayAllButton.ButtonRenderer.NavigationEndpoint.WatchEndpoint.PlaylistId
//}
//
//func (sr *ShelfRenderer) GridVideoRenderers() []GridVideoRenderer {
//	gvrs := make([]GridVideoRenderer, 0)
//	for _, item := range sr.Content.HorizontalListRenderer.Items {
//		gvrs = append(gvrs, item.GridVideoRenderer)
//	}
//	return gvrs
//}
