package youtube_urls

import (
	"net/url"
	"path"
)

func ThumbnailUrl(videoId string, tq ThumbnailQuality) *url.URL {
	qualityFilename, ok := thumbnailQualityFilenames[tq]
	if !ok {
		return nil
	}
	return &url.URL{
		Scheme: httpsScheme,
		Host:   ytImgHost,
		Path:   path.Join(viPath, videoId, qualityFilename+jpegExt),
	}
}

func ThumbnailUrls(videoId string) []*url.URL {
	return []*url.URL{
		ThumbnailUrl(videoId, ThumbnailQualityMaxRes),
		ThumbnailUrl(videoId, ThumbnailQualityHQ),
		ThumbnailUrl(videoId, ThumbnailQualityMQ),
		ThumbnailUrl(videoId, ThumbnailQualitySD),
	}
}
