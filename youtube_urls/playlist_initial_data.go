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
							} `json:"text,omitempty"`
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
		} `json:"playlistSidebarPrimaryInfoRenderer,omitempty"`
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
		} `json:"playlistSidebarSecondaryInfoRenderer,omitempty"`
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
									Contents []struct {
										PlaylistVideoListRenderer struct {
											PlaylistId string                             `json:"playlistId"`
											Contents   []PlaylistVideoListRendererContent `json:"contents"`
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
	VideoId       string   `json:"videoId"`
	Title         TextRuns `json:"title"`
	LengthSeconds string   `json:"lengthSeconds"`
	// normally contains video channel title
	ShortBylineText TextRuns `json:"shortBylineText"`
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
	LengthSeconds string
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
					pvlc = append(pvlc, itemSection.PlaylistVideoListRenderer.Contents...)
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
		title, titleRuns := "", vlc.PlaylistVideoRenderer.Title.Runs
		for _, r := range titleRuns {
			title += r.Text
		}
		sbTitle, sbTitleRuns := "", vlc.PlaylistVideoRenderer.ShortBylineText.Runs
		for _, r := range sbTitleRuns {
			sbTitle += r.Text
		}
		vits = append(vits, VideoIdTitleLengthChannel{
			VideoId:       videoId,
			Title:         title,
			LengthSeconds: vlc.PlaylistVideoRenderer.LengthSeconds,
			Channel:       sbTitle,
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
