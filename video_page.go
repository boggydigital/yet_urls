package yt_urls

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
)

const (
	StatusOK = "OK"
)

var (
	ErrSignatureCipher     = errors.New("signatureCipher")
	ErrMissingRequiredNode = errors.New("missing required node")
)

func getMatchingNodes(
	client *http.Client,
	u *url.URL,
	matches map[string]match_node.MatchDelegate) (map[string]*html.Node, error) {

	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	nodes := make(map[string]*html.Node)

	for title, match := range matches {
		if node := match_node.Match(doc, match); node != nil {
			nodes[title] = node
		}
	}

	return nodes, nil
}

func GetVideoPage(client *http.Client, videoId string) (*InitialPlayerResponse, error) {

	videoUrl := VideoUrl(videoId)

	scriptMatch := map[string]match_node.MatchDelegate{
		ytInitialPlayerResponse: initialPlayerResponseScript,
	}

	scriptNodes, err := getMatchingNodes(client, videoUrl, scriptMatch)
	if err != nil {
		return nil, err
	}

	if _, ok := scriptNodes[ytInitialPlayerResponse]; !ok {
		return nil, ErrMissingRequiredNode
	}

	iprReader := strings.NewReader(extractJsonObject(scriptNodes[ytInitialPlayerResponse].Data))

	var ipr InitialPlayerResponse
	if err := json.NewDecoder(iprReader).Decode(&ipr); err != nil {
		return nil, err
	}

	if ipr.PlayabilityStatus.Status != StatusOK {
		return nil, fmt.Errorf(ipr.PlayabilityStatus.Status)
	}

	signatureCipher := false

	formats := make(map[string]int, len(ipr.StreamingData.Formats))
	for _, f := range ipr.StreamingData.Formats {
		if f.Url == "" && f.SignatureCipher != "" {
			signatureCipher = true
			continue
		}
		formats[f.Url] = f.Bitrate
	}

	if len(formats) == 0 && signatureCipher {
		//TODO: support signature cipher YouTube URLs
		return nil, ErrSignatureCipher
	}

	return &ipr, nil
}
