package yt_urls

//StreamingFormat is a minimal set of data required to perform format selection
type StreamingFormat struct {
	Bitrate         int    `json:"bitrate"`
	Url             string `json:"url"`
	SignatureCipher string `json:"signatureCipher"`
}

//initialPlayerResponse is a minimal set of data structures required to decode and
//extract streaming data formats
type initialPlayerResponse struct {
	StreamingData struct {
		//use of Formats and not AdaptiveFormats is intentional, even though the former seemed
		//to be capped at 720p. AdaptiveFormats come as a separate video and audio tracks and
		//would require merging those two together.
		//Formats on the other hand contain URLs to files that contain both video and audio.
		//If you have a need for something more complex or flexible - you should consider
		//youtube-dl or any of the alternatives available
		Formats []StreamingFormat `json:"formats"`
	} `json:"streamingData"`
}
