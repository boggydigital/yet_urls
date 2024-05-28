package youtube_urls

import (
	"encoding/json"
	"github.com/boggydigital/match_node"
	"net/http"
	"strings"
)

func GetSearchResultsPage(client *http.Client, terms ...string) (*SearchInitialData, error) {
	searchResultsUrl := SearchResultsUrl(terms...)

	scriptMatches := make(map[string]match_node.Matcher)
	scriptMatches[ytInitialData] = &initialDataScriptMatcher{}
	scriptMatches[ytCfg] = &ytCfgScriptMatcher{}

	resp, err := client.Get(searchResultsUrl.String())
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

	var sid SearchInitialData
	if err := json.NewDecoder(idReader).Decode(&sid); err != nil {
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

	sid.Context = &itc

	return &sid, nil
}
