package yt_urls

import (
	"encoding/json"
	"github.com/boggydigital/match_node"
	"net/http"
	"strings"
)

func GetPlaylistPage(client *http.Client, playlistId string) (*ContextualPlaylist, error) {
	playlistUrl := PlaylistUrl(playlistId)

	scriptMatches := make(map[string]match_node.MatchDelegate)
	scriptMatches[ytInitialData] = initialDataScript
	scriptMatches[ytCfg] = ytCfgScript

	scriptNodes, err := getMatchingNodes(client, playlistUrl, scriptMatches)
	if err != nil {
		return nil, err
	}

	if _, ok := scriptNodes[ytInitialData]; !ok {
		return nil, ErrMissingRequiredNode
	}
	idReader := strings.NewReader(extractJsonObject(scriptNodes[ytInitialData].Data))

	var id InitialData
	if err := json.NewDecoder(idReader).Decode(&id); err != nil {
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

	return &ContextualPlaylist{
		Playlist: id.playlistVideoListContent(),
		Context:  &itc,
	}, nil
}
