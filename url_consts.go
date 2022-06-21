package yt_urls

const (
	videoParam          = "v"
	httpsScheme         = "https"
	YoutubeHost         = "youtube.com"
	youtubeWwwHost      = "www." + YoutubeHost
	youtubeImgHost      = "img." + YoutubeHost
	watchPath           = "watch"
	viPath              = "vi"
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

var thumbnailQualityFilenames = map[ThumbnailQuality]string{
	ThumbnailQualityMaxRes: "maxresdefault",
	ThumbnailQualityHQ:     "hqdefault",
	ThumbnailQualityMQ:     "mqdefault",
	ThumbnailQualitySD:     "sddefault",
}
