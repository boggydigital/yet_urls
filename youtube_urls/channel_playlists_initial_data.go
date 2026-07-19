package youtube_urls

type GridPlaylistRenderer struct {
	PlaylistId string `json:"playlistId"`
	Title      string `json:"title"`
}

func (cpid *ChannelPlaylistsInitialData) Playlists() []*GridPlaylistRenderer {
	playlists := make([]*GridPlaylistRenderer, 0)
	for _, tab := range cpid.Contents.TwoColumnBrowseResultsRenderer.Tabs {
		for _, tabContent := range tab.TabRenderer.Content.SectionListRenderer.Contents {
			for _, itemContent := range tabContent.ItemSectionRenderer.Contents {
				for _, item := range itemContent.GridRenderer.Items {
					playlists = append(playlists, new(GridPlaylistRenderer{
						PlaylistId: item.LockupViewModel.ContentId,
						Title:      item.LockupViewModel.Metadata.LockupMetadataViewModel.Title.Content,
					}))
				}
			}
		}
	}
	return playlists
}

type ChannelPlaylistsInitialData struct {
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
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		ResponseId                      string `json:"responseId"`
		WebResponseContextExtensionData struct {
			WebResponseContextPreloadData struct {
				PreloadMessageNames []string `json:"preloadMessageNames"`
			} `json:"webResponseContextPreloadData"`
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Endpoint struct {
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
							Params           string `json:"params"`
							CanonicalBaseUrl string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
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
												LockupViewModel struct {
													ContentImage struct {
														CollectionThumbnailViewModel struct {
															PrimaryThumbnail struct {
																ThumbnailViewModel struct {
																	Image struct {
																		Sources []struct {
																			Url    string `json:"url"`
																			Width  int    `json:"width"`
																			Height int    `json:"height"`
																		} `json:"sources"`
																	} `json:"image"`
																	Overlays []struct {
																		ThumbnailOverlayBadgeViewModel struct {
																			ThumbnailBadges []struct {
																				ThumbnailBadgeViewModel struct {
																					Icon struct {
																						Sources []struct {
																							ClientResource struct {
																								ImageName string `json:"imageName"`
																							} `json:"clientResource"`
																						} `json:"sources"`
																					} `json:"icon"`
																					Text            string `json:"text"`
																					BadgeStyle      string `json:"badgeStyle"`
																					BackgroundColor struct {
																						LightTheme int `json:"lightTheme"`
																						DarkTheme  int `json:"darkTheme"`
																					} `json:"backgroundColor"`
																				} `json:"thumbnailBadgeViewModel"`
																			} `json:"thumbnailBadges"`
																			Position string `json:"position"`
																		} `json:"thumbnailOverlayBadgeViewModel,omitempty"`
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
																		} `json:"thumbnailHoverOverlayViewModel,omitempty"`
																	} `json:"overlays"`
																	BackgroundColor struct {
																		LightTheme int `json:"lightTheme"`
																		DarkTheme  int `json:"darkTheme"`
																	} `json:"backgroundColor"`
																} `json:"thumbnailViewModel"`
															} `json:"primaryThumbnail"`
															StackColor struct {
																LightTheme int `json:"lightTheme"`
																DarkTheme  int `json:"darkTheme"`
															} `json:"stackColor"`
														} `json:"collectionThumbnailViewModel"`
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
																								BrowseId string `json:"browseId"`
																							} `json:"browseEndpoint"`
																						} `json:"innertubeCommand"`
																					} `json:"onTap"`
																				} `json:"commandRuns,omitempty"`
																				StyleRuns []struct {
																					StartIndex  int    `json:"startIndex"`
																					Length      int    `json:"length"`
																					WeightLabel string `json:"weightLabel"`
																				} `json:"styleRuns,omitempty"`
																			} `json:"text"`
																		} `json:"metadataParts"`
																	} `json:"metadataRows"`
																	Delimiter string `json:"delimiter"`
																} `json:"contentMetadataViewModel"`
															} `json:"metadata"`
														} `json:"lockupMetadataViewModel"`
													} `json:"metadata"`
													ContentId    string `json:"contentId"`
													ContentType  string `json:"contentType"`
													ItemPlayback struct {
														InlinePlayerData struct {
															OnSelect struct {
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
																		Params         string `json:"params"`
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
																		PlayerParams string `json:"playerParams,omitempty"`
																	} `json:"watchEndpoint"`
																} `json:"innertubeCommand"`
															} `json:"onSelect"`
															OnVisible struct {
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
															} `json:"onVisible"`
														} `json:"inlinePlayerData"`
													} `json:"itemPlayback"`
													RendererContext struct {
														LoggingContext struct {
															LoggingDirectives struct {
																TrackingParams string `json:"trackingParams"`
																Visibility     struct {
																	Types string `json:"types"`
																} `json:"visibility"`
															} `json:"loggingDirectives"`
														} `json:"loggingContext"`
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
																		Params         string `json:"params"`
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
																		PlayerParams string `json:"playerParams,omitempty"`
																	} `json:"watchEndpoint"`
																} `json:"innertubeCommand"`
															} `json:"onTap"`
														} `json:"commandContext"`
													} `json:"rendererContext"`
												} `json:"lockupViewModel"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
											TargetId       string `json:"targetId"`
										} `json:"gridRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							SubMenu        struct {
								ChannelSubMenuRenderer struct {
									SortSetting struct {
										SortFilterSubMenuRenderer struct {
											SubMenuItems []struct {
												Title              string `json:"title"`
												Selected           bool   `json:"selected"`
												NavigationEndpoint struct {
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
														Params           string `json:"params"`
														CanonicalBaseUrl string `json:"canonicalBaseUrl"`
													} `json:"browseEndpoint"`
												} `json:"navigationEndpoint"`
												TrackingParams string `json:"trackingParams"`
											} `json:"subMenuItems"`
											Title string `json:"title"`
											Icon  struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Accessibility struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibility"`
											TrackingParams string `json:"trackingParams"`
										} `json:"sortFilterSubMenuRenderer"`
									} `json:"sortSetting"`
								} `json:"channelSubMenuRenderer"`
							} `json:"subMenu"`
							TargetId             string `json:"targetId"`
							DisablePullToRefresh bool   `json:"disablePullToRefresh"`
						} `json:"sectionListRenderer"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
				ExpandableTabRenderer struct {
					Endpoint struct {
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
							Params           string `json:"params"`
							CanonicalBaseUrl string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
				} `json:"expandableTabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Header struct {
		PageHeaderRenderer struct {
			PageTitle string `json:"pageTitle"`
			Content   struct {
				PageHeaderViewModel struct {
					Title struct {
						DynamicTextViewModel struct {
							Text struct {
								Content string `json:"content"`
							} `json:"text"`
							MaxLines        int `json:"maxLines"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"dynamicTextViewModel"`
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
										Processor struct {
											BorderImageProcessor struct {
												Circular bool `json:"circular"`
											} `json:"borderImageProcessor"`
										} `json:"processor"`
									} `json:"image"`
									AvatarImageSize   string `json:"avatarImageSize"`
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
									} `json:"loggingDirectives"`
								} `json:"avatarViewModel"`
							} `json:"avatar"`
						} `json:"decoratedAvatarViewModel"`
					} `json:"image"`
					Metadata struct {
						ContentMetadataViewModel struct {
							MetadataRows []struct {
								MetadataParts []struct {
									Text struct {
										Content   string `json:"content"`
										StyleRuns []struct {
											WeightLabel        string `json:"weightLabel,omitempty"`
											StyleRunExtensions struct {
												StyleRunColorMapExtension struct {
													ColorMap []struct {
														Key   string `json:"key"`
														Value int64  `json:"value"`
													} `json:"colorMap"`
												} `json:"styleRunColorMapExtension"`
											} `json:"styleRunExtensions,omitempty"`
											StartIndex int `json:"startIndex,omitempty"`
											Length     int `json:"length,omitempty"`
										} `json:"styleRuns,omitempty"`
									} `json:"text"`
									EnableTruncation   bool   `json:"enableTruncation,omitempty"`
									AccessibilityLabel string `json:"accessibilityLabel,omitempty"`
								} `json:"metadataParts"`
							} `json:"metadataRows"`
							Delimiter       string `json:"delimiter"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"contentMetadataViewModel"`
					} `json:"metadata"`
					Actions struct {
						FlexibleActionsViewModel struct {
							ActionsRows []struct {
								Actions []struct {
									ButtonViewModel struct {
										Title string `json:"title"`
										OnTap struct {
											InnertubeCommand struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														IgnoreNavigation bool `json:"ignoreNavigation"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ModalEndpoint struct {
													Modal struct {
														ModalWithTitleAndButtonRenderer struct {
															Title struct {
																SimpleText string `json:"simpleText"`
															} `json:"title"`
															Content struct {
																SimpleText string `json:"simpleText"`
															} `json:"content"`
															Button struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		SimpleText string `json:"simpleText"`
																	} `json:"text"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignInEndpoint struct {
																			NextEndpoint struct {
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
																					Params           string `json:"params"`
																					CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																				} `json:"browseEndpoint"`
																			} `json:"nextEndpoint"`
																			ContinueAction string `json:"continueAction"`
																			IdamTag        string `json:"idamTag"`
																		} `json:"signInEndpoint"`
																	} `json:"navigationEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"button"`
														} `json:"modalWithTitleAndButtonRenderer"`
													} `json:"modal"`
												} `json:"modalEndpoint"`
											} `json:"innertubeCommand"`
										} `json:"onTap"`
										AccessibilityText string `json:"accessibilityText"`
										Style             string `json:"style"`
										TrackingParams    string `json:"trackingParams"`
										IsFullWidth       bool   `json:"isFullWidth"`
										Type              string `json:"type"`
										ButtonSize        string `json:"buttonSize"`
										State             string `json:"state"`
									} `json:"buttonViewModel"`
								} `json:"actions"`
							} `json:"actionsRows"`
							JustifyContent   string `json:"justifyContent"`
							MinimumRowHeight int    `json:"minimumRowHeight"`
							RendererContext  struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
										ClientVeSpec struct {
											UiType    int `json:"uiType"`
											VeCounter int `json:"veCounter"`
										} `json:"clientVeSpec"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"flexibleActionsViewModel"`
					} `json:"actions"`
					Description struct {
						DescriptionPreviewViewModel struct {
							Description struct {
								Content string `json:"content"`
							} `json:"description"`
							MaxLines       int `json:"maxLines"`
							TruncationText struct {
								Content   string `json:"content"`
								StyleRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									Weight     int `json:"weight"`
								} `json:"styleRuns"`
							} `json:"truncationText"`
							AlwaysShowTruncationText bool `json:"alwaysShowTruncationText"`
							RendererContext          struct {
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
											ClickTrackingParams         string `json:"clickTrackingParams"`
											ShowEngagementPanelEndpoint struct {
												EngagementPanel struct {
													EngagementPanelSectionListRenderer struct {
														Header struct {
															EngagementPanelTitleHeaderRenderer struct {
																Title struct {
																	SimpleText string `json:"simpleText"`
																} `json:"title"`
																VisibilityButton struct {
																	ButtonRenderer struct {
																		Style      string `json:"style"`
																		Size       string `json:"size"`
																		IsDisabled bool   `json:"isDisabled"`
																		Icon       struct {
																			IconType string `json:"iconType"`
																		} `json:"icon"`
																		Accessibility struct {
																			Label string `json:"label"`
																		} `json:"accessibility"`
																		TrackingParams    string `json:"trackingParams"`
																		AccessibilityData struct {
																			AccessibilityData struct {
																				Label string `json:"label"`
																			} `json:"accessibilityData"`
																		} `json:"accessibilityData"`
																		Command struct {
																			ClickTrackingParams                   string `json:"clickTrackingParams"`
																			ChangeEngagementPanelVisibilityAction struct {
																				TargetId   string `json:"targetId"`
																				Visibility string `json:"visibility"`
																			} `json:"changeEngagementPanelVisibilityAction"`
																		} `json:"command"`
																	} `json:"buttonRenderer"`
																} `json:"visibilityButton"`
																TrackingParams string `json:"trackingParams"`
															} `json:"engagementPanelTitleHeaderRenderer"`
														} `json:"header"`
														Content struct {
															SectionListRenderer struct {
																Contents []struct {
																	ItemSectionRenderer struct {
																		Contents []struct {
																			ContinuationItemRenderer struct {
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
																			} `json:"continuationItemRenderer"`
																		} `json:"contents"`
																		TrackingParams    string `json:"trackingParams"`
																		SectionIdentifier string `json:"sectionIdentifier"`
																		TargetId          string `json:"targetId"`
																	} `json:"itemSectionRenderer"`
																} `json:"contents"`
																TrackingParams  string `json:"trackingParams"`
																ScrollPaneStyle struct {
																	Scrollable bool `json:"scrollable"`
																} `json:"scrollPaneStyle"`
															} `json:"sectionListRenderer"`
														} `json:"content"`
														TargetId   string `json:"targetId"`
														Identifier struct {
															Surface string `json:"surface"`
															Tag     string `json:"tag"`
														} `json:"identifier"`
													} `json:"engagementPanelSectionListRenderer"`
												} `json:"engagementPanel"`
												Identifier struct {
													Surface string `json:"surface"`
													Tag     string `json:"tag"`
												} `json:"identifier"`
												EngagementPanelPresentationConfigs struct {
													EngagementPanelPopupPresentationConfig struct {
														PopupType string `json:"popupType"`
													} `json:"engagementPanelPopupPresentationConfig"`
												} `json:"engagementPanelPresentationConfigs"`
											} `json:"showEngagementPanelEndpoint"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
								} `json:"commandContext"`
							} `json:"rendererContext"`
						} `json:"descriptionPreviewViewModel"`
					} `json:"description"`
					Banner struct {
						ImageBannerViewModel struct {
							Image struct {
								Sources []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"sources"`
							} `json:"image"`
							Style           string `json:"style"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
							} `json:"rendererContext"`
						} `json:"imageBannerViewModel"`
					} `json:"banner"`
					RendererContext struct {
						LoggingContext struct {
							LoggingDirectives struct {
								TrackingParams string `json:"trackingParams"`
								Visibility     struct {
									Types string `json:"types"`
								} `json:"visibility"`
							} `json:"loggingDirectives"`
						} `json:"loggingContext"`
					} `json:"rendererContext"`
				} `json:"pageHeaderViewModel"`
			} `json:"content"`
		} `json:"pageHeaderRenderer"`
	} `json:"header"`
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
			ChannelUrl             string   `json:"channelUrl"`
			IsFamilySafe           bool     `json:"isFamilySafe"`
			AvailableCountryCodes  []string `json:"availableCountryCodes"`
			AndroidDeepLink        string   `json:"androidDeepLink"`
			AndroidAppindexingLink string   `json:"androidAppindexingLink"`
			IosAppindexingLink     string   `json:"iosAppindexingLink"`
			VanityChannelUrl       string   `json:"vanityChannelUrl"`
		} `json:"channelMetadataRenderer"`
	} `json:"metadata"`
	TrackingParams string `json:"trackingParams"`
	Topbar         struct {
		DesktopTopbarRenderer struct {
			Logo struct {
				TopbarLogoRenderer struct {
					IconImage struct {
						IconType string `json:"iconType"`
					} `json:"iconImage"`
					TooltipText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"tooltipText"`
					Endpoint struct {
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
							BrowseId string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					TrackingParams    string `json:"trackingParams"`
					OverrideEntityKey string `json:"overrideEntityKey"`
				} `json:"topbarLogoRenderer"`
			} `json:"logo"`
			Searchbox struct {
				FusionSearchboxRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					PlaceholderText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"placeholderText"`
					Config struct {
						WebSearchboxConfig struct {
							RequestLanguage     string `json:"requestLanguage"`
							RequestDomain       string `json:"requestDomain"`
							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
							FocusSearchbox      bool   `json:"focusSearchbox"`
						} `json:"webSearchboxConfig"`
					} `json:"config"`
					TrackingParams string `json:"trackingParams"`
					SearchEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SearchEndpoint struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"searchEndpoint"`
					ClearButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"clearButton"`
					ShowImageSourceDialog struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ShowDialogCommand   struct {
							PanelLoadingStrategy struct {
								InlineContent struct {
									DialogViewModel struct {
										Header struct {
											DialogHeaderViewModel struct {
												Headline struct {
													Content string `json:"content"`
												} `json:"headline"`
											} `json:"dialogHeaderViewModel"`
										} `json:"header"`
										Footer struct {
											PanelFooterViewModel struct {
												PrimaryButton struct {
													ButtonViewModel struct {
														Title          string `json:"title"`
														Style          string `json:"style"`
														TrackingParams string `json:"trackingParams"`
														IsFullWidth    bool   `json:"isFullWidth"`
														Type           string `json:"type"`
													} `json:"buttonViewModel"`
												} `json:"primaryButton"`
												SecondaryButton struct {
													ButtonViewModel struct {
														Title          string `json:"title"`
														Style          string `json:"style"`
														TrackingParams string `json:"trackingParams"`
														IsFullWidth    bool   `json:"isFullWidth"`
														Type           string `json:"type"`
													} `json:"buttonViewModel"`
												} `json:"secondaryButton"`
												ShouldHideDivider bool `json:"shouldHideDivider"`
											} `json:"panelFooterViewModel"`
										} `json:"footer"`
										Content struct {
											BasicContentViewModel struct {
												Paragraphs []struct {
													Text struct {
														Content string `json:"content"`
													} `json:"text"`
												} `json:"paragraphs"`
											} `json:"basicContentViewModel"`
										} `json:"content"`
									} `json:"dialogViewModel"`
								} `json:"inlineContent"`
							} `json:"panelLoadingStrategy"`
						} `json:"showDialogCommand"`
					} `json:"showImageSourceDialog"`
					DisableAiAppearance bool `json:"disableAiAppearance"`
				} `json:"fusionSearchboxRenderer"`
			} `json:"searchbox"`
			TrackingParams string `json:"trackingParams"`
			TopbarButtons  []struct {
				TopbarMenuButtonRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					MenuRequest struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										MultiPageMenuRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											Style              string `json:"style"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"multiPageMenuRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"menuRequest"`
					TrackingParams string `json:"trackingParams"`
					Accessibility  struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					Tooltip string `json:"tooltip"`
					Style   string `json:"style"`
				} `json:"topbarMenuButtonRenderer,omitempty"`
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignInEndpoint struct {
							IdamTag string `json:"idamTag"`
						} `json:"signInEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					TargetId       string `json:"targetId"`
				} `json:"buttonRenderer,omitempty"`
			} `json:"topbarButtons"`
			HotkeyDialog struct {
				HotkeyDialogRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					Sections []struct {
						HotkeyDialogSectionRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Options []struct {
								HotkeyDialogSectionOptionRenderer struct {
									Label struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"label"`
									Hotkey                   string `json:"hotkey"`
									HotkeyAccessibilityLabel struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"hotkeyAccessibilityLabel,omitempty"`
								} `json:"hotkeyDialogSectionOptionRenderer"`
							} `json:"options"`
						} `json:"hotkeyDialogSectionRenderer"`
					} `json:"sections"`
					DismissButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"dismissButton"`
					TrackingParams string `json:"trackingParams"`
				} `json:"hotkeyDialogRenderer"`
			} `json:"hotkeyDialog"`
			BackButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"backButton"`
			ForwardButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"forwardButton"`
			A11YSkipNavigationButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"a11ySkipNavigationButton"`
			VoiceSearchButton struct {
				ButtonRenderer struct {
					Style           string `json:"style"`
					Size            string `json:"size"`
					IsDisabled      bool   `json:"isDisabled"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										VoiceSearchDialogRenderer struct {
											PlaceholderHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"placeholderHeader"`
											PromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptHeader"`
											ExampleQuery1 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery1"`
											ExampleQuery2 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery2"`
											PromptMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptMicrophoneLabel"`
											LoadingHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"loadingHeader"`
											ConnectionErrorHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorHeader"`
											ConnectionErrorMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorMicrophoneLabel"`
											PermissionsHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsHeader"`
											PermissionsSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsSubtext"`
											DisabledHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledHeader"`
											DisabledSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledSubtext"`
											MicrophoneButtonAriaLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneButtonAriaLabel"`
											ExitButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Icon       struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
												} `json:"buttonRenderer"`
											} `json:"exitButton"`
											TrackingParams            string `json:"trackingParams"`
											MicrophoneOffPromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneOffPromptHeader"`
										} `json:"voiceSearchDialogRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"serviceEndpoint"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					Tooltip           string `json:"tooltip"`
					TrackingParams    string `json:"trackingParams"`
					AccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityData"`
				} `json:"buttonRenderer"`
			} `json:"voiceSearchButton"`
		} `json:"desktopTopbarRenderer"`
	} `json:"topbar"`
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
			ChannelProfileMicroformatDetails struct {
				ProfilePage struct {
					Context     string `json:"context"`
					Type        string `json:"type"`
					Url         string `json:"url"`
					Name        string `json:"name"`
					Description string `json:"description"`
					MainEntity  struct {
						Type                 string `json:"type"`
						Name                 string `json:"name"`
						Url                  string `json:"url"`
						AlternateName        string `json:"alternateName"`
						Image                string `json:"image"`
						Description          string `json:"description"`
						InteractionStatistic []struct {
							Type            string `json:"type"`
							InteractionType struct {
								Type string `json:"type"`
							} `json:"interactionType"`
							UserInteractionCount string `json:"userInteractionCount"`
						} `json:"interactionStatistic"`
					} `json:"mainEntity"`
				} `json:"profilePage"`
			} `json:"channelProfileMicroformatDetails"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
	Context *ytCfgInnerTubeContext
}
