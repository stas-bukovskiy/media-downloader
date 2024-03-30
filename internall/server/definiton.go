// Package server provides the HTTP server that serves the API endpoints.
package server

import "github.com/stas-bukovskiy/media-downloader/internall/model"

// VideoService is the interface that provides the methods to interact with the video service.
type VideoService interface {
	MetaInformation(url string) (*model.VideoInfo, error)
	DownloadLinks(url string) ([]model.DownloadLink, error)
}
