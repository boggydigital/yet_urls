package yt_urls

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ContextualPlaylist struct {
	Playlist *PlaylistContent
	Context  *ytCfgInnerTubeContext
}

func (cp *ContextualPlaylist) Videos() []VideoIdTitleChannel {
	var vits []VideoIdTitleChannel
	vits = make([]VideoIdTitleChannel, 0, len(cp.Playlist.Content))
	for _, vlc := range cp.Playlist.Content {
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
	for i := len(cp.Playlist.Content) - 1; i >= 0; i-- {
		if cp.Playlist.Content[i].ContinuationItemRenderer.Trigger != "" {
			return true
		}
	}
	return false
}

func (cp *ContextualPlaylist) continuationEndpoint() *ContinuationEndpoint {
	for i := len(cp.Playlist.Content) - 1; i >= 0; i-- {
		if cp.Playlist.Content[i].ContinuationItemRenderer.Trigger != "" {
			return &cp.Playlist.Content[i].ContinuationItemRenderer.ContinuationEndpoint
		}
	}
	return nil
}

func (cp *ContextualPlaylist) Continue(client *http.Client) (*ContextualPlaylist, error) {

	if !cp.HasContinuation() {
		return nil, nil
	}

	contEndpoint := cp.continuationEndpoint()

	data := browseRequest{
		Context:      cp.Context.InnerTubeContext,
		Continuation: contEndpoint.ContinuationCommand.Token,
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
		return nil, err
	}

	browseUrl := BrowseUrl(
		contEndpoint.CommandMetadata.WebCommandMetadata.ApiUrl,
		cp.Context.APIKey)

	resp, err := client.Post(browseUrl.String(), contentType, b)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var br browseResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return nil, err
	}

	cp.Playlist.Content = br.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems

	return &ContextualPlaylist{
		Playlist: cp.Playlist,
		Context:  cp.Context,
	}, nil
}
