package service

import (
	"fmt"

	"github.com/stas-bukovskiy/media-downloader/internall/model"
)

// VideoService is the service that provides the methods to interact with the video downloader.
type VideoService struct {
	downloader VideoDownloader
}

// NewVideoService creates a new VideoService with the provided VideoDownloader.
func NewVideoService(downloader VideoDownloader) *VideoService {
	return &VideoService{downloader: downloader}
}

// DownloadLinks returns the download links for the video at the provided URL.
func (s *VideoService) DownloadLinks(url string) ([]model.DownloadLink, error) {
	// call the downloader to get the video response
	video, err := s.downloader.Download(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download video: %w", err)
	}

	// map the video formats to the download links struct
	links := make([]model.DownloadLink, len(video.Formats))
	for i, format := range video.Formats {
		links[i] = model.DownloadLink{
			URL:        format.URL,
			Resolution: format.Resolution,
			Format:     format.Format,
		}
	}

	return links, nil
}

// MetaInformation returns the meta-information for the video at the provided URL.
func (s *VideoService) MetaInformation(url string) (*model.VideoInfo, error) {
	// call the downloader to get the video response
	video, err := s.downloader.Download(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download video: %w", err)
	}

	// map and return the video response to the video info struct
	return &model.VideoInfo{
		Title:       video.Title,
		Thumbnail:   video.Thumbnail,
		UploadDate:  video.UploadDate,
		Uploader:    video.Uploader,
		AspectRatio: video.AspectRatio,
		Description: video.Description,
		Duration:    video.Duration,
		Tags:        video.Tags,
		Channel: model.ChannelInfo{
			Title:         video.Channel,
			FollowerCount: video.ChannelFollowerCount,
			URL:           video.ChannelURL,
		},
		Statistic: model.StatisticInfo{
			ViewCount:     video.ViewCount,
			AverageRating: video.AverageRating,
			CommentCount:  video.CommentCount,
		},
	}, nil
}
