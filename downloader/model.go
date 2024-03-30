package downloader

const (
	// ResponseStatusSuccess is the status of the response when the request is successful.
	ResponseStatusSuccess = "success"
	// ResponseStatusFailed is the status of the response when the request is failed.
	ResponseStatusFailed = "failed"
)

// DefaultResponse is the default response from the API.
type DefaultResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// VideoResponse is the response struct for the video information that contains the video information and default response.
type VideoResponse struct {
	DefaultResponse
	Video
}

// Video is the struct that contains the information about the video.
type Video struct {
	ACodec               string                `json:"acodec"`
	AgeLimit             int                   `json:"age_limit"`
	AspectRatio          float64               `json:"aspect_ratio"`
	AudioChannels        int                   `json:"audio_channels"`
	Availability         string                `json:"availability"`
	AverageRating        string                `json:"average_rating"`
	Categories           []string              `json:"categories"`
	Channel              string                `json:"channel"`
	ChannelFollowerCount int                   `json:"channel_follower_count"`
	ChannelID            string                `json:"channel_id"`
	ChannelURL           string                `json:"channel_url"`
	Chapters             []Chapter             `json:"chapters"`
	CommentCount         int                   `json:"comment_count"`
	Description          string                `json:"description"`
	Duration             int                   `json:"duration"`
	DurationString       string                `json:"duration_string"`
	DynamicRange         string                `json:"dynamic_range"`
	Epoch                int                   `json:"epoch"`
	Ext                  string                `json:"ext"`
	Extractor            string                `json:"extractor"`
	ExtractorKey         string                `json:"extractor_key"`
	FilesizeApprox       int                   `json:"filesize_approx"`
	Format               string                `json:"format"`
	FormatID             string                `json:"format_id"`
	FormatNote           string                `json:"format_note"`
	Formats              []Format              `json:"formats"`
	Resolution           string                `json:"resolution"`
	Subtitles            map[string][]Subtitle `json:"subtitles"`
	Tags                 []string              `json:"tags"`
	Thumbnail            string                `json:"thumbnail"`
	Thumbnails           []Thumbnail           `json:"thumbnails"`
	Title                string                `json:"title"`
	UploadDate           string                `json:"upload_date"`
	Uploader             string                `json:"uploader"`
	UploaderID           string                `json:"uploader_id"`
	UploaderURL          string                `json:"uploader_url"`
	ViewCount            int                   `json:"view_count"`
	WebpageURL           string                `json:"webpage_url"`
	WebpageURLBasename   string                `json:"webpage_url_basename"`
	WebpageURLDomain     string                `json:"webpage_url_domain"`
}

// Chapter is the struct that contains the information about the chapter for the video.
type Chapter struct {
	Title     string  `json:"title"`
	EndTime   float64 `json:"end_time"`
	StartTime float64 `json:"start_time"`
}

// Thumbnail is the struct that contains the information about the thumbnail for the video.
type Thumbnail struct {
	ID         string  `json:"id"`
	URL        string  `json:"url"`
	Height     *int    `json:"height"`
	Width      *int    `json:"width"`
	Preference int     `json:"preference"`
	Resolution *string `json:"resolution"`
}

// Subtitle is the struct that contains the information about the subtitle for the video.
type Subtitle struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Extension string `json:"ext"`
}

// Format is the struct that contains the information about the format and download URL for the video.
type Format struct {
	URL         string  `json:"url"`
	AspectRatio float64 `json:"aspect_ratio"`
	Extension   string  `json:"ext"`
	Resolution  string  `json:"resolution"`
	Height      int     `json:"height"`
	Width       int     `json:"width"`
	Protocol    string  `json:"protocol"`
	Format      string  `json:"format"`
	FormatID    string  `json:"format_id"`
	FormatNote  string  `json:"format_note"`
	Fps         float64 `json:"fps"`
}
