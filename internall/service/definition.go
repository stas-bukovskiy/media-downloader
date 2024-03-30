package service

import "github.com/stas-bukovskiy/media-downloader/downloader"

// VideoDownloader is the interface that wraps the Download method.
type VideoDownloader interface {
	Download(url string) (*downloader.VideoResponse, error)
}
