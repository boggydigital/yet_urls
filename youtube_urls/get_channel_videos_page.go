package youtube_urls

import (
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
