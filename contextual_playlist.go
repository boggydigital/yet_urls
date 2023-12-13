package yt_urls

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ContextualPlaylist struct {
	Playlist *PlaylistInitialData
	Context  *ytCfgInnerTubeContext
}

func (cp *ContextualPlaylist) Videos() []VideoIdTitleChannel {
	var vits []VideoIdTitleChannel
	pc := cp.Playlist.PlaylistContent()
	vits = make([]VideoIdTitleChannel, 0, len(pc))
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
		vits = append(vits, VideoIdTitleChannel{
			VideoId: videoId,
			Title:   title,
			Channel: sbTitle,
		})
	}
	return vits
}

func (cp *ContextualPlaylist) HasContinuation() bool {
	pc := cp.Playlist.PlaylistContent()
	for i := len(pc) - 1; i >= 0; i-- {
		if pc[i].ContinuationItemRenderer.Trigger != "" {
			return true
		}
	}
	return false
}

func (cp *ContextualPlaylist) continuationEndpoint() *ContinuationEndpoint {
	pc := cp.Playlist.PlaylistContent()
	for i := len(pc) - 1; i >= 0; i-- {
		if pc[i].ContinuationItemRenderer.Trigger != "" {
			return &pc[i].ContinuationItemRenderer.ContinuationEndpoint
		}
	}
	return nil
}

func (cp *ContextualPlaylist) Continue(client *http.Client) error {

	if !cp.HasContinuation() {
		return nil
	}

	contEndpoint := cp.continuationEndpoint()

	data := browseRequest{
		Context:      cp.Context.InnerTubeContext,
		Continuation: contEndpoint.ContinuationCommand.Token,
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
		return err
	}

	browseUrl := BrowseUrl(
		contEndpoint.CommandMetadata.WebCommandMetadata.ApiUrl,
		cp.Context.APIKey)

	resp, err := client.Post(browseUrl.String(), contentType, b)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	var br browseResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return err
	}

	// update contents internals
	cp.Playlist.SetContent(br.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems)

	return nil
}
