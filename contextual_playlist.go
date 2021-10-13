package yt_urls

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ContextualPlaylist struct {
	Playlist []PlaylistVideoListRendererContent
	Context  *ytCfgInnerTubeContext
}

func (cp *ContextualPlaylist) Videos() []VideoIdTitle {
	var vits []VideoIdTitle
	vits = make([]VideoIdTitle, 0, len(cp.Playlist))
	for _, vlc := range cp.Playlist {
		videoId := vlc.PlaylistVideoRenderer.VideoId
		if videoId == "" {
			continue
		}
		title := vlc.PlaylistVideoRenderer.Title.Runs[0].Text
		for ii := 1; ii < len(vlc.PlaylistVideoRenderer.Title.Runs); ii++ {
			title += vlc.PlaylistVideoRenderer.Title.Runs[ii].Text
		}
		vits = append(vits, VideoIdTitle{
			VideoId: videoId,
			Title:   title,
		})
	}
	return vits
}

func (cp *ContextualPlaylist) HasContinuation() bool {
	for i := len(cp.Playlist) - 1; i >= 0; i-- {
		if cp.Playlist[i].ContinuationItemRenderer.Trigger != "" {
			return true
		}
	}
	return false
}

func (cp *ContextualPlaylist) continuationEndpoint() *ContinuationEndpoint {
	for i := len(cp.Playlist) - 1; i >= 0; i-- {
		if cp.Playlist[i].ContinuationItemRenderer.Trigger != "" {
			return &cp.Playlist[i].ContinuationItemRenderer.ContinuationEndpoint
		}
	}
	return nil
}

func (cp *ContextualPlaylist) Continue() (*ContextualPlaylist, error) {

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

	resp, err := http.Post(browseUrl.String(), contentType, b)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var br browseResponse
	if err := json.NewDecoder(resp.Body).Decode(&br); err != nil {
		return nil, err
	}

	return &ContextualPlaylist{
		Playlist: br.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems,
		Context:  cp.Context,
	}, nil
}
