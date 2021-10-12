package yt_urls

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	StatusOK = "OK"
)

var (
	ErrorSignatureCipher    = errors.New("signatureCipher")
	ErrorScriptNodeNotFound = errors.New("script node with JSON data not found")
)

func getScriptJsonReader(u *url.URL, matchScript func(node *html.Node) bool) (io.Reader, error) {

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	if node := match_node.Match(doc, matchScript); node != nil {
		return strings.NewReader(extractJsonObject(node.Data)), nil
	}

	return nil, ErrorScriptNodeNotFound
}

func GetVideoPage(videoId string) (*InitialPlayerResponse, error) {

	videoUrl := VideoUrl(videoId)

	iprReader, err := getScriptJsonReader(videoUrl, initialPlayerResponseScript)
	if err != nil {
		return nil, err
	}

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
		return nil, ErrorSignatureCipher
	}

	return &ipr, nil
}
