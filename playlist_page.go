package yt_urls

import "encoding/json"

const (
	ytInitialData = "var ytInitialData ="
)

func GetPlaylistPage(playlistId string) (*InitialData, error) {
	playlistUrl := PlaylistUrl(playlistId)

	idReader, err := getScriptJsonReader(playlistUrl, initialDataScript)
	if err != nil {
		return nil, err
	}

	var id InitialData
	if err := json.NewDecoder(idReader).Decode(&id); err != nil {
		return nil, err
	}

	return &id, nil
}
