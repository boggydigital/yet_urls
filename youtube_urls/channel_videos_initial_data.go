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
					} `json:"content"`
				} `json:"tabRenderer"`
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

//type RichGridRendererContents struct {
//	RichItemRenderer struct {
//		Content struct {
//			VideoRenderer struct {
//				VideoId            string   `json:"videoId"`
//				Title              TextRuns `json:"title"`
//				DescriptionSnippet struct {
//					Runs []struct {
//						Text string `json:"text"`
//					} `json:"runs"`
//				} `json:"descriptionSnippet"`
//				PublishedTimeText struct {
//					SimpleText string `json:"simpleText"`
//				} `json:"publishedTimeText"`
//				LengthText struct {
//					Accessibility struct {
//						AccessibilityData struct {
//							Label string `json:"label"`
//						} `json:"accessibilityData"`
//					} `json:"accessibility"`
//					SimpleText string `json:"simpleText"`
//				} `json:"lengthText"`
//				ViewCountText struct {
//					SimpleText string `json:"simpleText"`
//				} `json:"viewCountText"`
//				IsWatched bool `json:"isWatched,omitempty"`
//			} `json:"videoRenderer"`
//		} `json:"content"`
//	} `json:"richItemRenderer"`
//	ContinuationItemRenderer ContinuationItemRenderer `json:"continuationItemRenderer"`
//}

type RichGridRendererContents struct {
	RichItemRenderer struct {
		Content struct {
			LockupViewModel struct {
				ContentImage struct {
					ThumbnailViewModel struct {
						Image struct {
							Sources []struct {
								Url    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"image"`
						Overlays []struct {
							ThumbnailBottomOverlayViewModel struct {
								Badges []struct {
									ThumbnailBadgeViewModel struct {
										Text                         string `json:"text"`
										BadgeStyle                   string `json:"badgeStyle"`
										AnimationActivationTargetId  string `json:"animationActivationTargetId"`
										AnimationActivationEntityKey string `json:"animationActivationEntityKey"`
										LottieData                   struct {
											Url      string `json:"url"`
											Settings struct {
												Loop     bool `json:"loop"`
												Autoplay bool `json:"autoplay"`
											} `json:"settings"`
										} `json:"lottieData"`
										AnimatedText                          string `json:"animatedText"`
										AnimationActivationEntitySelectorType string `json:"animationActivationEntitySelectorType"`
										RendererContext                       struct {
											AccessibilityContext struct {
												Label string `json:"label"`
											} `json:"accessibilityContext"`
										} `json:"rendererContext"`
									} `json:"thumbnailBadgeViewModel"`
								} `json:"badges"`
							} `json:"thumbnailBottomOverlayViewModel,omitempty"`
							ThumbnailHoverOverlayToggleActionsViewModel struct {
								Buttons []struct {
									ToggleButtonViewModel struct {
										DefaultButtonViewModel struct {
											ButtonViewModel struct {
												IconName string `json:"iconName"`
												OnTap    struct {
													InnertubeCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl,omitempty"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistId string `json:"playlistId"`
															Actions    []struct {
																AddedVideoId string `json:"addedVideoId"`
																Action       string `json:"action"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint,omitempty"`
														SignalServiceEndpoint struct {
															Signal  string `json:"signal"`
															Actions []struct {
																ClickTrackingParams  string `json:"clickTrackingParams"`
																AddToPlaylistCommand struct {
																	OpenMiniplayer      bool   `json:"openMiniplayer"`
																	VideoId             string `json:"videoId"`
																	ListType            string `json:"listType"`
																	OnCreateListCommand struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				ApiUrl   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		CreatePlaylistServiceEndpoint struct {
																			VideoIds []string `json:"videoIds"`
																			Params   string   `json:"params"`
																		} `json:"createPlaylistServiceEndpoint"`
																	} `json:"onCreateListCommand"`
																	VideoIds     []string `json:"videoIds"`
																	VideoCommand struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		WatchEndpoint struct {
																			VideoId                            string `json:"videoId"`
																			WatchEndpointSupportedOnesieConfig struct {
																				Html5PlaybackOnesieConfig struct {
																					CommonConfig struct {
																						Url string `json:"url"`
																					} `json:"commonConfig"`
																				} `json:"html5PlaybackOnesieConfig"`
																			} `json:"watchEndpointSupportedOnesieConfig"`
																			PlayerParams string `json:"playerParams,omitempty"`
																		} `json:"watchEndpoint"`
																	} `json:"videoCommand"`
																} `json:"addToPlaylistCommand"`
															} `json:"actions"`
														} `json:"signalServiceEndpoint,omitempty"`
													} `json:"innertubeCommand"`
												} `json:"onTap"`
												AccessibilityText string `json:"accessibilityText"`
												Style             string `json:"style"`
												TrackingParams    string `json:"trackingParams"`
												Type              string `json:"type"`
												ButtonSize        string `json:"buttonSize"`
												State             string `json:"state"`
											} `json:"buttonViewModel"`
										} `json:"defaultButtonViewModel"`
										ToggledButtonViewModel struct {
											ButtonViewModel struct {
												IconName string `json:"iconName"`
												OnTap    struct {
													InnertubeCommand struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														PlaylistEditEndpoint struct {
															PlaylistId string `json:"playlistId"`
															Actions    []struct {
																Action         string `json:"action"`
																RemovedVideoId string `json:"removedVideoId"`
															} `json:"actions"`
														} `json:"playlistEditEndpoint"`
													} `json:"innertubeCommand"`
												} `json:"onTap,omitempty"`
												AccessibilityText string `json:"accessibilityText"`
												Style             string `json:"style"`
												TrackingParams    string `json:"trackingParams"`
												Type              string `json:"type"`
												ButtonSize        string `json:"buttonSize"`
												State             string `json:"state"`
											} `json:"buttonViewModel"`
										} `json:"toggledButtonViewModel"`
										IsToggled      bool   `json:"isToggled"`
										TrackingParams string `json:"trackingParams"`
									} `json:"toggleButtonViewModel"`
								} `json:"buttons"`
							} `json:"thumbnailHoverOverlayToggleActionsViewModel,omitempty"`
						} `json:"overlays"`
					} `json:"thumbnailViewModel"`
				} `json:"contentImage"`
				Metadata struct {
					LockupMetadataViewModel struct {
						Title struct {
							Content string `json:"content"`
						} `json:"title"`
						Metadata struct {
							ContentMetadataViewModel struct {
								MetadataRows []struct {
									MetadataParts []struct {
										Text struct {
											Content string `json:"content"`
										} `json:"text"`
										AccessibilityLabel string `json:"accessibilityLabel,omitempty"`
										LeadingIcon        struct {
											Name   string `json:"name"`
											Height int    `json:"height"`
											Width  int    `json:"width"`
										} `json:"leadingIcon,omitempty"`
									} `json:"metadataParts"`
								} `json:"metadataRows"`
								Delimiter string `json:"delimiter"`
							} `json:"contentMetadataViewModel"`
						} `json:"metadata"`
						MenuButton struct {
							ButtonViewModel struct {
								IconName string `json:"iconName"`
								OnTap    struct {
									InnertubeCommand struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										ShowSheetCommand    struct {
											PanelLoadingStrategy struct {
												InlineContent struct {
													SheetViewModel struct {
														Content struct {
															ListViewModel struct {
																ListItems []struct {
																	ListItemViewModel struct {
																		Title struct {
																			Content string `json:"content"`
																		} `json:"title"`
																		LeadingImage struct {
																			Sources []struct {
																				ClientResource struct {
																					ImageName string `json:"imageName"`
																				} `json:"clientResource"`
																			} `json:"sources"`
																		} `json:"leadingImage"`
																		RendererContext struct {
																			LoggingContext struct {
																				LoggingDirectives struct {
																					TrackingParams string `json:"trackingParams"`
																					Visibility     struct {
																						Types string `json:"types"`
																					} `json:"visibility"`
																				} `json:"loggingDirectives"`
																			} `json:"loggingContext,omitempty"`
																			CommandContext struct {
																				OnTap struct {
																					InnertubeCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost    bool   `json:"sendPost,omitempty"`
																								Url         string `json:"url,omitempty"`
																								WebPageType string `json:"webPageType,omitempty"`
																								RootVe      int    `json:"rootVe,omitempty"`
																								ApiUrl      string `json:"apiUrl,omitempty"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						SignalServiceEndpoint struct {
																							Signal  string `json:"signal"`
																							Actions []struct {
																								ClickTrackingParams  string `json:"clickTrackingParams"`
																								AddToPlaylistCommand struct {
																									OpenMiniplayer      bool   `json:"openMiniplayer"`
																									VideoId             string `json:"videoId"`
																									ListType            string `json:"listType"`
																									OnCreateListCommand struct {
																										ClickTrackingParams string `json:"clickTrackingParams"`
																										CommandMetadata     struct {
																											WebCommandMetadata struct {
																												SendPost bool   `json:"sendPost"`
																												ApiUrl   string `json:"apiUrl"`
																											} `json:"webCommandMetadata"`
																										} `json:"commandMetadata"`
																										CreatePlaylistServiceEndpoint struct {
																											VideoIds []string `json:"videoIds"`
																											Params   string   `json:"params"`
																										} `json:"createPlaylistServiceEndpoint"`
																									} `json:"onCreateListCommand"`
																									VideoIds     []string `json:"videoIds"`
																									VideoCommand struct {
																										ClickTrackingParams string `json:"clickTrackingParams"`
																										CommandMetadata     struct {
																											WebCommandMetadata struct {
																												Url         string `json:"url"`
																												WebPageType string `json:"webPageType"`
																												RootVe      int    `json:"rootVe"`
																											} `json:"webCommandMetadata"`
																										} `json:"commandMetadata"`
																										WatchEndpoint struct {
																											VideoId                            string `json:"videoId"`
																											WatchEndpointSupportedOnesieConfig struct {
																												Html5PlaybackOnesieConfig struct {
																													CommonConfig struct {
																														Url string `json:"url"`
																													} `json:"commonConfig"`
																												} `json:"html5PlaybackOnesieConfig"`
																											} `json:"watchEndpointSupportedOnesieConfig"`
																											PlayerParams string `json:"playerParams,omitempty"`
																										} `json:"watchEndpoint"`
																									} `json:"videoCommand"`
																								} `json:"addToPlaylistCommand"`
																							} `json:"actions"`
																						} `json:"signalServiceEndpoint,omitempty"`
																						SignInEndpoint struct {
																							NextEndpoint struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								ShowSheetCommand    struct {
																									PanelLoadingStrategy struct {
																										RequestTemplate struct {
																											PanelId string `json:"panelId"`
																											Params  string `json:"params"`
																										} `json:"requestTemplate"`
																									} `json:"panelLoadingStrategy"`
																								} `json:"showSheetCommand"`
																							} `json:"nextEndpoint"`
																						} `json:"signInEndpoint,omitempty"`
																						ShareEntityServiceEndpoint struct {
																							SerializedShareEntity string `json:"serializedShareEntity"`
																							Commands              []struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								OpenPopupAction     struct {
																									Popup struct {
																										UnifiedSharePanelRenderer struct {
																											TrackingParams     string `json:"trackingParams"`
																											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																										} `json:"unifiedSharePanelRenderer"`
																									} `json:"popup"`
																									PopupType string `json:"popupType"`
																									BeReused  bool   `json:"beReused"`
																								} `json:"openPopupAction"`
																							} `json:"commands"`
																						} `json:"shareEntityServiceEndpoint,omitempty"`
																					} `json:"innertubeCommand"`
																				} `json:"onTap"`
																			} `json:"commandContext"`
																		} `json:"rendererContext"`
																	} `json:"listItemViewModel"`
																} `json:"listItems"`
															} `json:"listViewModel"`
														} `json:"content"`
													} `json:"sheetViewModel"`
												} `json:"inlineContent"`
											} `json:"panelLoadingStrategy"`
										} `json:"showSheetCommand"`
									} `json:"innertubeCommand"`
								} `json:"onTap"`
								AccessibilityText string `json:"accessibilityText"`
								Style             string `json:"style"`
								TrackingParams    string `json:"trackingParams"`
								Type              string `json:"type"`
								ButtonSize        string `json:"buttonSize"`
								State             string `json:"state"`
							} `json:"buttonViewModel"`
						} `json:"menuButton"`
					} `json:"lockupMetadataViewModel"`
				} `json:"metadata"`
				ContentId       string `json:"contentId"`
				ContentType     string `json:"contentType"`
				RendererContext struct {
					LoggingContext struct {
						LoggingDirectives struct {
							TrackingParams string `json:"trackingParams"`
							Visibility     struct {
								Types string `json:"types"`
							} `json:"visibility"`
						} `json:"loggingDirectives"`
					} `json:"loggingContext"`
					AccessibilityContext struct {
						Label string `json:"label"`
					} `json:"accessibilityContext"`
					CommandContext struct {
						OnTap struct {
							InnertubeCommand struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
									PlayerParams string `json:"playerParams,omitempty"`
								} `json:"watchEndpoint"`
							} `json:"innertubeCommand"`
						} `json:"onTap"`
					} `json:"commandContext"`
				} `json:"rendererContext"`
			} `json:"lockupViewModel"`
		} `json:"content"`
		TrackingParams string `json:"trackingParams"`
	} `json:"richItemRenderer,omitempty"`
	ContinuationItemRenderer struct {
		Trigger              string               `json:"trigger"`
		ContinuationEndpoint ContinuationEndpoint `json:"continuationEndpoint"`
	} `json:"continuationItemRenderer,omitempty"`
	TrackingParams string `json:"trackingParams"`
	Header         struct {
		ChipBarViewModel struct {
			Chips []struct {
				ChipViewModel struct {
					Text        string `json:"text"`
					Selected    bool   `json:"selected"`
					DisplayType string `json:"displayType"`
					TapCommand  struct {
						InnertubeCommand struct {
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
								Command struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									ShowReloadUiCommand struct {
										TargetId string `json:"targetId"`
									} `json:"showReloadUiCommand"`
								} `json:"command"`
							} `json:"continuationCommand"`
						} `json:"innertubeCommand"`
					} `json:"tapCommand"`
					AccessibilityLabel string `json:"accessibilityLabel"`
					LoggingDirectives  struct {
						TrackingParams string `json:"trackingParams"`
						Visibility     struct {
							Types string `json:"types"`
						} `json:"visibility"`
					} `json:"loggingDirectives"`
				} `json:"chipViewModel"`
			} `json:"chips"`
		} `json:"chipBarViewModel"`
	} `json:"header"`
	TargetId string `json:"targetId"`
	Style    string `json:"style"`
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

		lvm := vlc.RichItemRenderer.Content.LockupViewModel

		var seconds int64
		for _, overlay := range lvm.ContentImage.ThumbnailViewModel.Overlays {
			for _, badge := range overlay.ThumbnailBottomOverlayViewModel.Badges {
				if seconds = lengthTextToSeconds(badge.ThumbnailBadgeViewModel.Text); seconds > 0 {
					break
				}
			}
		}

		videoId := lvm.ContentId
		if videoId == "" {
			continue
		}

		title := lvm.Metadata.LockupMetadataViewModel.Title.Content

		vit := VideoIdTitleLengthChannel{
			VideoId: videoId,
			Title:   title,
		}

		if seconds > 0 {
			vit.LengthSeconds = seconds
		}

		vits = append(vits, vit)
	}
	return vits
}

func lengthTextToSeconds(lt string) int64 {
	var seconds int64
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
	return seconds
}
