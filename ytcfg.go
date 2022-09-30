package yt_urls

import (
	"golang.org/x/net/html"
	"strings"
)

const (
	opYtCfgMarker = "ytcfg.set({"
	clYtCfgMarker = "});"
	ytCfg         = "(function() {window.ytplayer={};\n" + opYtCfgMarker
)

func extractYtCfgJsonObject(data string) string {
	fi, li := strings.Index(data, opYtCfgMarker), strings.Index(data, clYtCfgMarker)
	return data[fi+len(opYtCfgMarker)-1 : li+1]
}

type ytCfgScriptMatcher struct{}

//ytCfgScript is an HTML node filter for YouTube <script> text content
//that contains ytcfg initialization data
func (ytcsm *ytCfgScriptMatcher) Match(node *html.Node) bool {
	if node.Type != html.TextNode ||
		node.Parent == nil ||
		node.Parent.Data != "script" {
		return false
	}

	return strings.HasPrefix(node.Data, ytCfg)
}

type ytCfgInnerTubeContext struct {
	APIKey           string       `json:"INNERTUBE_API_KEY"`
	InnerTubeContext ytCfgContext `json:"INNERTUBE_CONTEXT"`
}

type ytCfgContext struct {
	Client struct {
		Hl               string `json:"hl"`
		Gl               string `json:"gl"`
		RemoteHost       string `json:"remoteHost"`
		DeviceMake       string `json:"deviceMake"`
		DeviceModel      string `json:"deviceModel"`
		VisitorData      string `json:"visitorData"`
		UserAgent        string `json:"userAgent"`
		ClientName       string `json:"clientName"`
		ClientVersion    string `json:"clientVersion"`
		OsName           string `json:"osName"`
		OsVersion        string `json:"osVersion"`
		OriginalUrl      string `json:"originalUrl"`
		Platform         string `json:"platform"`
		ClientFormFactor string `json:"clientFormFactor"`
		ConfigInfo       struct {
			AppInstallData string `json:"appInstallData"`
		} `json:"configInfo"`
	} `json:"client"`
	User struct {
		LockedSafetyMode bool `json:"lockedSafetyMode"`
	} `json:"user"`
	Request struct {
		UseSsl bool `json:"useSsl"`
	} `json:"request"`
	ClickTracking struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
	} `json:"clickTracking"`
}
