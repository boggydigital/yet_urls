package youtube_urls

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/html"
	"net/http"
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

//type PlaylistHeaderRenderer struct {
//	PlaylistId      string     `json:"playlistId"`
//	Title           SimpleText `json:"title"`
//	DescriptionText SimpleText `json:"descriptionText"`
//	OwnerText       TextRuns   `json:"ownerText"`
//	ViewCountText   SimpleText `json:"viewCountText"`
//	Privacy         string     `json:"privacy"`
//}

type PageHeaderRenderer struct {
	PageTitle string `json:"pageTitle"`
	Content   struct {
		PageHeaderViewModel struct {
			Title struct {
				DynamicTextViewModel struct {
					Text struct {
						Content string `json:"content"`
					} `json:"text"`
				} `json:"dynamicTextViewModel"`
			} `json:"title"`
			Metadata struct {
				ContentMetadataViewModel struct {
					MetadataRows []struct {
						MetadataParts []struct {
							Text struct {
								Content string `json:"content"`
							} `json:"text"`
						} `json:"metadataParts"`
					} `json:"metadataRows"`
					Delimiter string `json:"delimiter"`
				} `json:"contentMetadataViewModel"`
			} `json:"metadata"`
			Description struct {
				DescriptionPreviewViewModel struct {
					Description struct {
						Content string `json:"content"`
					} `json:"description"`
					TruncationText struct {
						Content   string `json:"content"`
						StyleRuns []struct {
							StartIndex int `json:"startIndex"`
							Length     int `json:"length"`
							Weight     int `json:"weight"`
						} `json:"styleRuns"`
					} `json:"truncationText"`
				} `json:"descriptionPreviewViewModel"`
			} `json:"description"`
			HeroImage struct {
				ContentPreviewImageViewModel struct {
					Image struct {
						Sources []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"sources"`
					} `json:"image"`
					Style      string `json:"style"`
					LayoutMode string `json:"layoutMode"`
					Overlays   []struct {
						ThumbnailHoverOverlayViewModel struct {
							Icon struct {
								Sources []struct {
									ClientResource struct {
										ImageName string `json:"imageName"`
									} `json:"clientResource"`
								} `json:"sources"`
							} `json:"icon"`
							Text struct {
								Content   string `json:"content"`
								StyleRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
								} `json:"styleRuns"`
							} `json:"text"`
							Style string `json:"style"`
						} `json:"thumbnailHoverOverlayViewModel"`
					} `json:"overlays"`
				} `json:"contentPreviewImageViewModel"`
			} `json:"heroImage"`
			Background struct {
				CinematicContainerViewModel struct {
					BackgroundImageConfig struct {
						Image struct {
							Sources []struct {
								Url    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"sources"`
						} `json:"image"`
					} `json:"backgroundImageConfig"`
					GradientColorConfig []struct {
						LightThemeColor int64   `json:"lightThemeColor"`
						DarkThemeColor  int64   `json:"darkThemeColor"`
						StartLocation   float64 `json:"startLocation"`
					} `json:"gradientColorConfig"`
					Config struct {
						LightThemeBackgroundColor int64 `json:"lightThemeBackgroundColor"`
						DarkThemeBackgroundColor  int64 `json:"darkThemeBackgroundColor"`
						ColorSourceSizeMultiplier int   `json:"colorSourceSizeMultiplier"`
						ApplyClientImageBlur      bool  `json:"applyClientImageBlur"`
					} `json:"config"`
				} `json:"cinematicContainerViewModel"`
			} `json:"background"`
		} `json:"pageHeaderViewModel"`
	} `json:"content"`
}

type PlaylistMetadataRenderer struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PlaylistSidebarRenderer struct {
	Items []struct {
		PlaylistSidebarPrimaryInfoRenderer struct {
			ThumbnailRenderer struct {
				PlaylistVideoThumbnailRenderer struct {
					Thumbnail struct {
						Thumbnails []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"thumbnail"`
					TrackingParams string `json:"trackingParams"`
				} `json:"playlistVideoThumbnailRenderer"`
			} `json:"thumbnailRenderer"`
			Title TextRuns   `json:"title"`
			Stats []TextRuns `json:"stats"`
			Menu  struct {
				MenuRenderer struct {
					Items []struct {
						MenuNavigationItemRenderer struct {
							Text struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
						} `json:"menuNavigationItemRenderer"`
					} `json:"items"`
					TargetId string `json:"targetId"`
				} `json:"menuRenderer"`
			} `json:"menu"`
			ThumbnailOverlays []struct {
				ThumbnailOverlaySidePanelRenderer struct {
					Text struct {
						SimpleText string `json:"simpleText"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
				} `json:"thumbnailOverlaySidePanelRenderer"`
			} `json:"thumbnailOverlays"`
			Description struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			ShowMoreText TextRuns `json:"showMoreText"`
		} `json:"playlistSidebarPrimaryInfoRenderer"`
		PlaylistSidebarSecondaryInfoRenderer struct {
			VideoOwner struct {
				VideoOwnerRenderer struct {
					Thumbnail struct {
						Thumbnails []struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"thumbnail"`
					Title TextRuns `json:"title"`
				} `json:"videoOwnerRenderer"`
			} `json:"videoOwner"`
		} `json:"playlistSidebarSecondaryInfoRenderer"`
	} `json:"items"`
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
									Contents []LockupViewModelContent `json:"contents"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		PageHeaderRenderer PageHeaderRenderer `json:"pageHeaderRenderer"`
	} `json:"header"`
	Metadata struct {
		PlaylistMetadataRenderer PlaylistMetadataRenderer `json:"playlistMetadataRenderer"`
	} `json:"metadata"`
	Sidebar struct {
		PlaylistSidebarRenderer PlaylistSidebarRenderer `json:"playlistSidebarRenderer"`
	} `json:"sidebar"`
	videoListContent []PlaylistVideoListRendererContent
	Context          *ytCfgInnerTubeContext
}

type LockupViewModelContent struct {
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
				Image struct {
					DecoratedAvatarViewModel struct {
						Avatar struct {
							AvatarViewModel struct {
								Image struct {
									Sources []struct {
										Url    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"sources"`
								} `json:"image"`
								AvatarImageSize string `json:"avatarImageSize"`
							} `json:"avatarViewModel"`
						} `json:"avatar"`
						A11YLabel       string `json:"a11yLabel"`
						RendererContext struct {
							CommandContext struct {
								OnTap struct {
									InnertubeCommand struct {
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
									} `json:"innertubeCommand"`
								} `json:"onTap"`
							} `json:"commandContext"`
						} `json:"rendererContext"`
					} `json:"decoratedAvatarViewModel"`
				} `json:"image"`
				Metadata struct {
					ContentMetadataViewModel struct {
						MetadataRows []struct {
							MetadataParts []struct {
								Text struct {
									Content     string `json:"content"`
									CommandRuns []struct {
										StartIndex int `json:"startIndex"`
										Length     int `json:"length"`
										OnTap      struct {
											InnertubeCommand struct {
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
											} `json:"innertubeCommand"`
										} `json:"onTap"`
									} `json:"commandRuns,omitempty"`
									StyleRuns []struct {
										StartIndex         int    `json:"startIndex"`
										Length             int    `json:"length,omitempty"`
										WeightLabel        string `json:"weightLabel,omitempty"`
										StyleRunExtensions struct {
											StyleRunColorMapExtension struct {
												ColorMap []struct {
													Key   string `json:"key"`
													Value int64  `json:"value"`
												} `json:"colorMap"`
											} `json:"styleRunColorMapExtension"`
										} `json:"styleRunExtensions,omitempty"`
									} `json:"styleRuns,omitempty"`
									AttachmentRuns []struct {
										StartIndex int `json:"startIndex"`
										Length     int `json:"length"`
										Element    struct {
											Type struct {
												ImageType struct {
													Image struct {
														Sources []struct {
															ClientResource struct {
																ImageName string `json:"imageName"`
															} `json:"clientResource"`
															Width  int `json:"width"`
															Height int `json:"height"`
														} `json:"sources"`
													} `json:"image"`
												} `json:"imageType"`
											} `json:"type"`
											Properties struct {
												LayoutProperties struct {
													Height struct {
														Value int    `json:"value"`
														Unit  string `json:"unit"`
													} `json:"height"`
													Width struct {
														Value int    `json:"value"`
														Unit  string `json:"unit"`
													} `json:"width"`
													Margin struct {
														Left struct {
															Value int    `json:"value"`
															Unit  string `json:"unit"`
														} `json:"left"`
													} `json:"margin"`
												} `json:"layoutProperties"`
											} `json:"properties"`
										} `json:"element"`
										Alignment string `json:"alignment"`
									} `json:"attachmentRuns,omitempty"`
								} `json:"text"`
								AccessibilityLabel string `json:"accessibilityLabel,omitempty"`
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
							VideoId        string `json:"videoId"`
							PlaylistId     string `json:"playlistId"`
							Index          int    `json:"index"`
							Params         string `json:"params"`
							PlayerParams   string `json:"playerParams"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointSupportedOnesieConfig struct {
								Html5PlaybackOnesieConfig struct {
									CommonConfig struct {
										Url string `json:"url"`
									} `json:"commonConfig"`
								} `json:"html5PlaybackOnesieConfig"`
							} `json:"watchEndpointSupportedOnesieConfig"`
						} `json:"watchEndpoint"`
					} `json:"innertubeCommand"`
				} `json:"onTap"`
			} `json:"commandContext"`
		} `json:"rendererContext"`
	} `json:"lockupViewModel"`
}

type Text struct {
	Text string `json:"text"`
}

type TextRuns struct {
	Runs []struct {
		Text string `json:"text"`
	} `json:"runs"`
	Accessibility struct {
		AccessibilityData struct {
			Label string `json:"label"`
		} `json:"accessibilityData"`
	} `json:"accessibility"`
}

func (tr *TextRuns) String() string {
	textRuns := make([]string, 0, len(tr.Runs))
	for _, r := range tr.Runs {
		textRuns = append(textRuns, r.Text)
	}
	return strings.Join(textRuns, "")
}

type PlaylistVideoListRendererContent struct {
	PlaylistVideoRenderer    PlaylistVideoRenderer
	ContinuationItemRenderer ContinuationItemRenderer
}

type PlaylistVideoRenderer struct {
	VideoId       string `json:"videoId"`
	Title         string `json:"title"`
	LengthSeconds int64  `json:"lengthSeconds"`
	Channel       string `json:"channel"`
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

type VideoIdTitleLengthChannel struct {
	VideoId       string
	Title         string
	LengthSeconds int64
	Channel       string
}

func (id *PlaylistInitialData) PageHeaderRenderer() *PageHeaderRenderer {
	return &id.Header.PageHeaderRenderer
}

func (id *PlaylistInitialData) PlaylistTitle() string {
	return id.Metadata.PlaylistMetadataRenderer.Title
}

func (id *PlaylistInitialData) PlaylistContent() []PlaylistVideoListRendererContent {

	if id.videoListContent == nil {
		pvlc := make([]PlaylistVideoListRendererContent, 0)

		for _, tab := range id.Contents.TwoColumnBrowseResultsRenderer.Tabs {
			for _, sectionList := range tab.TabRenderer.Content.SectionListRenderer.Contents {
				for _, itemSection := range sectionList.ItemSectionRenderer.Contents {

					pvlrc := PlaylistVideoListRendererContent{
						PlaylistVideoRenderer: PlaylistVideoRenderer{
							VideoId: itemSection.LockupViewModel.ContentId,
							Title:   itemSection.LockupViewModel.Metadata.LockupMetadataViewModel.Title.Content,
						},
					}

					lengthSet := false
					for _, overlay := range itemSection.LockupViewModel.ContentImage.ThumbnailViewModel.Overlays {
						if lengthSet {
							break
						}
						for _, badge := range overlay.ThumbnailBottomOverlayViewModel.Badges {
							pvlrc.PlaylistVideoRenderer.LengthSeconds = lengthTextToSeconds(badge.ThumbnailBadgeViewModel.Text)
							lengthSet = true
							break
						}
					}

					channelSet := false
					for _, row := range itemSection.LockupViewModel.Metadata.LockupMetadataViewModel.Metadata.ContentMetadataViewModel.MetadataRows {
						if channelSet {
							break
						}
						for _, part := range row.MetadataParts {
							pvlrc.PlaylistVideoRenderer.Channel = part.Text.Content
							channelSet = true
							break
						}
					}

					//itemSection.LockupViewModel.Metadata.LockupMetadataViewModel.
					pvlc = append(pvlc, pvlrc)
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
	for _, item := range id.Sidebar.PlaylistSidebarRenderer.Items {
		if videoOwnerTitle := item.PlaylistSidebarSecondaryInfoRenderer.VideoOwner.VideoOwnerRenderer.Title.String(); videoOwnerTitle != "" {
			return videoOwnerTitle
		}
	}
	return ""
}

func (pid *PlaylistInitialData) Videos() []VideoIdTitleLengthChannel {
	var vits []VideoIdTitleLengthChannel
	pc := pid.PlaylistContent()
	vits = make([]VideoIdTitleLengthChannel, 0, len(pc))
	for _, vlc := range pc {
		videoId := vlc.PlaylistVideoRenderer.VideoId
		if videoId == "" {
			continue
		}

		title := vlc.PlaylistVideoRenderer.Title

		vits = append(vits, VideoIdTitleLengthChannel{
			VideoId:       videoId,
			Title:         title,
			LengthSeconds: vlc.PlaylistVideoRenderer.LengthSeconds,
			Channel:       vlc.PlaylistVideoRenderer.Channel,
		})
	}
	return vits
}

func (pid *PlaylistInitialData) HasContinuation() bool {
	pc := pid.PlaylistContent()
	for i := len(pc) - 1; i >= 0; i-- {
		if pc[i].ContinuationItemRenderer.Trigger != "" {
			return true
		}
	}
	return false
}

func (pid *PlaylistInitialData) continuationEndpoint() *ContinuationEndpoint {
	pc := pid.PlaylistContent()
	for i := len(pc) - 1; i >= 0; i-- {
		if pc[i].ContinuationItemRenderer.Trigger != "" {
			return &pc[i].ContinuationItemRenderer.ContinuationEndpoint
		}
	}
	return nil
}

func (pid *PlaylistInitialData) Continue(client *http.Client) error {

	if !pid.HasContinuation() {
		return nil
	}

	contEndpoint := pid.continuationEndpoint()

	data := browseRequest{
		Context:      pid.Context.InnerTubeContext,
		Continuation: contEndpoint.ContinuationCommand.Token,
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
		return err
	}

	browseUrl := BrowseUrl(
		contEndpoint.CommandMetadata.WebCommandMetadata.ApiUrl,
		pid.Context.APIKey)

	resp, err := client.Post(browseUrl.String(), contentType, b)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	var br playlistBrowseResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return err
	}

	// update contents internals
	pid.SetContent(br.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems)

	return nil
}
