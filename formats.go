package yet_urls

import "strings"

const (
	videoMIMETypePrefix          = "video"
	preferredVideoMIMETypePrefix = "video/mp4; codecs=\"avc1"
	audioMIMETypePrefix          = "audio"
	preferredAudioMIMETypePrefix = "audio/mp4; codecs=\"mp4a"
)

type Range struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Format captures stream data provided by YouTube
type Format struct {
	iTag             int     `json:"itag"`
	Url              string  `json:"url"`
	MIMEType         string  `json:"mimeType"`
	Bitrate          int     `json:"bitrate"`
	Width            int     `json:"width"`
	Height           int     `json:"height"`
	InitRange        Range   `json:"initRange"`
	IndexRange       Range   `json:"indexRange"`
	LastModified     string  `json:"lastModified"`
	ContentLength    string  `json:"contentLength"`
	Quality          string  `json:"quality"`
	FPS              int     `json:"fps"`
	QualityLabel     string  `json:"qualityLabel"`
	ProjectionType   string  `json:"projectionType"`
	AverageBitrate   int     `json:"averageBitrate"`
	HighReplication  bool    `json:"highReplication"`
	AudioQuality     string  `json:"audioQuality"`
	ApproxDurationMs string  `json:"approxDurationMs"`
	AudioSampleRate  string  `json:"audioSampleRate"`
	AudioChannels    int     `json:"audioChannels"`
	LoudnessDb       float64 `json:"loudnessDb"`
	SignatureCipher  string  `json:"signatureCipher"`
}

type Formats []*Format

func (fs Formats) Len() int {
	return len(fs)
}

func (fs Formats) Less(i, j int) bool {
	return fs[i].AverageBitrate < fs[j].AverageBitrate
}

func (fs Formats) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs Formats) filterByMIMETypePrefix(pfx string) Formats {
	formats := make(Formats, 0, len(fs))
	for _, f := range fs {
		if !strings.HasPrefix(f.MIMEType, pfx) {
			continue
		}
		formats = append(formats, f)
	}
	return formats
}

func (fs Formats) Video() Formats {
	return fs.filterByMIMETypePrefix(videoMIMETypePrefix)
}

func (fs Formats) PreferredVideoFormats() Formats {
	return fs.filterByMIMETypePrefix(preferredVideoMIMETypePrefix)
}

func (fs Formats) Audio() Formats {
	return fs.filterByMIMETypePrefix(audioMIMETypePrefix)
}

func (fs Formats) PreferredAudioFormats() Formats {
	return fs.filterByMIMETypePrefix(preferredAudioMIMETypePrefix)
}

var mimeExt = map[string]string{
	"video/mp4":  mp4Ext,
	"video/webm": webmExt,
	"audio/mp4":  mp4Ext,
	"audio/webm": webmExt,
}

func (f Format) Ext() string {
	mt := f.MIMEType
	mime, _, ok := strings.Cut(mt, ";")
	if !ok {
		return DefaultVideoExt
	}
	if ext, ok := mimeExt[mime]; ok {
		return ext
	}
	return DefaultVideoExt
}
