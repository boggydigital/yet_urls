package yt_urls

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/http"
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
	body io.Reader,
	matches map[string]match_node.Matcher) (map[string]*html.Node, error) {

	doc, err := html.Parse(body)
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

const (
	playerUrlKey = "player_ias.vflset"
	playerUrlSfx = "base.js"
)

type playerScriptMatcher struct{}

func (psm *playerScriptMatcher) Match(node *html.Node) bool {
	if (node.DataAtom != atom.Script && node.DataAtom != atom.Link) ||
		len(node.Attr) < 1 {
		return false
	}

	for _, attr := range node.Attr {
		if (node.DataAtom == atom.Script && attr.Key == "src") ||
			(node.DataAtom == atom.Link && attr.Key == "href") {
			if strings.HasSuffix(attr.Val, playerUrlSfx) && strings.Contains(attr.Val, playerUrlKey) {
				return true
			}
		}
	}

	return false
}

func getPlayerUrl(body io.Reader) (string, error) {

	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	node := match_node.Match(doc, &playerScriptMatcher{})

	src := match_node.AttrVal(node, "src")

	return src, nil
}

func GetVideoPage(client *http.Client, videoId string) (*InitialPlayerResponse, error) {

	videoUrl := VideoUrl(videoId)

	scriptMatch := map[string]match_node.Matcher{
		ytInitialPlayerResponse: &initialPlayerResponseMatcher{},
	}

	resp, err := client.Get(videoUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	sb := &strings.Builder{}
	tr := io.TeeReader(resp.Body, sb)

	playerUrl, err := getPlayerUrl(tr)
	if err != nil {
		return nil, err
	}

	scriptNodes, err := getMatchingNodes(strings.NewReader(sb.String()), scriptMatch)
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
		return nil, fmt.Errorf("%s: %s",
			ipr.PlayabilityStatus.Reason,
			ipr.PlayabilityStatus.ErrorScreen.PlayerErrorMessageRenderer.SubReason.SimpleText)
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

	// set player URL before returning
	ipr.PlayerUrl = playerUrl

	return &ipr, nil
}
