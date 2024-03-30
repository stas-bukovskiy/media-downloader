// Package model provides the data model for the application.
package model

// DownloadLink represents a downloadable link for a video.
type DownloadLink struct {
	URL        string `json:"url"`
	Resolution string `json:"resolution"`
	Format     string `json:"format"`
}

// VideoInfo represents the information about the video.
type VideoInfo struct {
	Title       string        `json:"title"`
	Thumbnail   string        `json:"thumbnail"`
	UploadDate  string        `json:"upload_date"`
	Uploader    string        `json:"uploader"`
	AspectRatio float64       `json:"aspect_ratio"`
	Description string        `json:"description"`
	Duration    int           `json:"duration"`
	Tags        []string      `json:"tags"`
	Channel     ChannelInfo   `json:"channel"`
	Statistic   StatisticInfo `json:"statistic"`
}

// ChannelInfo represents the information about the channel.
type ChannelInfo struct {
	Title         string `json:"title"`
	FollowerCount int    `json:"follower_count"`
	URL           string `json:"url"`
}

// StatisticInfo represents the statistics about the video.
type StatisticInfo struct {
	ViewCount     int    `json:"view_count"`
	AverageRating string `json:"average_rating"`
	CommentCount  int    `json:"comment_count"`
}
