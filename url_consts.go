package yt_urls

const (
	videoParam          = "v"
	httpsScheme         = "https"
	YoutubeHost         = "youtube.com"
	ytImgHost           = "i.ytimg.com"
	youtubeWwwHost      = "www." + YoutubeHost
	watchPath           = "watch"
	viPath              = "/vi"
	mp4Ext              = ".mp4"
	webmExt             = ".webm"
	jpegExt             = ".jpg"
	DefaultVideoExt     = mp4Ext
	DefaultThumbnailExt = jpegExt
)

type ThumbnailQuality int

const (
	ThumbnailQualityMaxRes ThumbnailQuality = iota
	ThumbnailQualityHQ
	ThumbnailQualityMQ
	ThumbnailQualitySD
)

func AllThumbnailQualities() []ThumbnailQuality {
	return []ThumbnailQuality{
		ThumbnailQualityMaxRes,
		ThumbnailQualityHQ,
		ThumbnailQualityMQ,
		ThumbnailQualitySD,
	}
}

var thumbnailQualityFilenames = map[ThumbnailQuality]string{
	ThumbnailQualityMaxRes: "maxresdefault",
	ThumbnailQualityHQ:     "hqdefault",
	ThumbnailQualityMQ:     "mqdefault",
	ThumbnailQualitySD:     "sddefault",
}

func (tq ThumbnailQuality) String() string {
	return thumbnailQualityFilenames[tq]
}

func ParseThumbnailQuality(tqs string) ThumbnailQuality {
	for k, v := range thumbnailQualityFilenames {
		if v == tqs {
			return k
		}
	}
	return ThumbnailQualityHQ
}
