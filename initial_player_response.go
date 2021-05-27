package yt_urls

//streamingFormat is a minimal set of data required to make format selection and
//return a URL
type streamingFormat struct {
	Url     string `json:"url"`
	Bitrate int    `json:"bitrate"`
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
		Formats []streamingFormat `json:"formats"`
	} `json:"streamingData"`
}
