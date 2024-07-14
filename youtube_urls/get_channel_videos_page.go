package youtube_urls

import (
	"bytes"
	"encoding/json"
	"github.com/boggydigital/match_node"
	"net/http"
	"strings"
)

func GetChannelVideosPage(client *http.Client, channelId string) (*ChannelVideosInitialData, error) {
	channelUrl := ChannelVideosUrl(channelId)

	scriptMatches := make(map[string]match_node.Matcher)
	scriptMatches[ytInitialData] = &initialDataScriptMatcher{}
	scriptMatches[ytCfg] = &ytCfgScriptMatcher{}

	resp, err := client.Get(channelUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	scriptNodes, err := getMatchingNodes(resp.Body, scriptMatches)
	if err != nil {
		return nil, err
	}

	if _, ok := scriptNodes[ytInitialData]; !ok {
		return nil, ErrMissingRequiredNode
	}
	idReader := strings.NewReader(extractJsonObject(scriptNodes[ytInitialData].Data))

	var cvid ChannelVideosInitialData
	if err := json.NewDecoder(idReader).Decode(&cvid); err != nil {
		return nil, err
	}

	if _, ok := scriptNodes[ytCfg]; !ok {
		return nil, ErrMissingRequiredNode
	}
	itcReader := strings.NewReader(extractYtCfgJsonObject(scriptNodes[ytCfg].Data))

	var itc ytCfgInnerTubeContext

	if err := json.NewDecoder(itcReader).Decode(&itc); err != nil {
		return nil, err
	}

	cvid.Context = &itc

	return &cvid, nil
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
