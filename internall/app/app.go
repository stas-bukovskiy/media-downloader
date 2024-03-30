package app

import (
	"log"

	"github.com/stas-bukovskiy/media-downloader/downloader"
	"github.com/stas-bukovskiy/media-downloader/internall/config"
	"github.com/stas-bukovskiy/media-downloader/internall/proxy"
	"github.com/stas-bukovskiy/media-downloader/internall/server"
	"github.com/stas-bukovskiy/media-downloader/internall/service"
)

// Run is the entry point of the application that initializes and starts the HTTP server.
func Run() {
	// Load the configuration
	cong, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize the video downloader
	videoDownloader := downloader.NewVideDownloader(cong.APIKey)
	videoDownloaderProxy := proxy.NewVideoDownloaderProxy(videoDownloader)

	// Initialize the video service
	videoService := service.NewVideoService(videoDownloaderProxy)

	// Initialize and start the HTTP server
	srv := server.NewServer(cong.AppPort, videoService)
	srv.Start()
}
