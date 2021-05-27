# yt_urls

yt_urls extracts a single video stream URL ("the best" quality one) from a YouTube video page. This module has the following exported methods:

## WatchUrl

```golang
func WatchUrl(videoId string) *url.URL
```

`WatchUrl` provides a URL for a video-id, e.g. http://www.youtube.com/watch?v=video-id1 for "video-id1".

## BestStreamingUrl

```golang
func BestStreamingUrl(videoId string) (*url.URL, error)
```

`BestStreamingUrl` extracts the URL for "the best" streaming format for a given YouTube video-id. 

Here are the key steps to make that happen:

1) convert video-id to a full YouTube.com/watch URL
2) request page content at that URL
3) parse response as HTML document and find required node (`iprScriptTextContent` contains selection criteria)
4) decode `ytInitialPlayerResponse` object (to a minimal data struct)
5) select "the best" streaming format available (`bestFormatByBitrate` contains selection criteria)
6) return URL for that format
