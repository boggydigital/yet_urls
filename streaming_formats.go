package yt_urls

//StreamingFormat is a minimal set of data required to perform format selection
type StreamingFormat struct {
	Bitrate         int    `json:"bitrate"`
	Url             string `json:"url"`
	SignatureCipher string `json:"signatureCipher"`
}

type StreamingFormats []StreamingFormat

func (sf StreamingFormats) Len() int {
	return len(sf)
}

func (sf StreamingFormats) Less(i, j int) bool {
	return sf[i].Bitrate < sf[j].Bitrate
}

func (sf StreamingFormats) Swap(i, j int) {
	sf[i], sf[j] = sf[j], sf[i]
}
