package youtube_urls

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
