package youtube_urls

type ThumbnailQuality int

const (
	ThumbnailQualityUnknown ThumbnailQuality = iota
	ThumbnailQualitySD
	ThumbnailQualityMQ
	ThumbnailQualityHQ
	ThumbnailQualityMaxRes
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
	ThumbnailQualityUnknown: "unknown",
	ThumbnailQualityMaxRes:  "maxresdefault",
	ThumbnailQualityHQ:      "hqdefault",
	ThumbnailQualityMQ:      "mqdefault",
	ThumbnailQualitySD:      "sddefault",
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
	return ThumbnailQualityUnknown
}

func LowerQuality(q ThumbnailQuality) ThumbnailQuality {
	tqi := int(q)
	if tqi <= int(ThumbnailQualityMaxRes) && tqi > int(ThumbnailQualitySD) {
		return ThumbnailQuality(tqi - 1)
	}
	return ThumbnailQualityUnknown
}
