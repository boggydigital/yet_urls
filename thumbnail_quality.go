package yt_urls

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
