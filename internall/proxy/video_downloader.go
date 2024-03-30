package proxy

import (
	"log"
	"sync"
	"time"

	"github.com/stas-bukovskiy/media-downloader/downloader"
)

// VideoDownloaderProxy is a proxy for the downloader.VideoDownloader that caches the responses and logs the time consumed for each download
type VideoDownloaderProxy struct {
	// cache is a map that holds the downloaded video responses
	cache map[string]*downloader.VideoResponse
	// realDownloader is the real downloader that the proxy delegates the download requests to
	realDownloader *downloader.VideoDownloader

	// mu is a mutex to synchronize the access to the cache
	mu sync.Mutex
}

// NewVideoDownloaderProxy creates a new VideoDownloaderProxy
func NewVideoDownloaderProxy(realDownloader *downloader.VideoDownloader) *VideoDownloaderProxy {
	return &VideoDownloaderProxy{
		realDownloader: realDownloader,
		cache:          make(map[string]*downloader.VideoResponse),
	}
}

// Download wraps the real downloader's Download method and caches the responses
func (p *VideoDownloaderProxy) Download(url string) (*downloader.VideoResponse, error) {
	// check if the response is already cached
	p.mu.Lock()
	if response, found := p.cache[url]; found {
		log.Printf("Cache hit for URL: %s\n", url)
		return response, nil
	}
	p.mu.Unlock()

	log.Printf("Cache miss for URL: %s. Downloading...\n", url)
	start := time.Now()

	// retrieve the response from the real downloader
	response, err := p.realDownloader.Download(url)
	if err != nil {
		return nil, err
	}

	// cache the response
	p.mu.Lock()
	p.cache[url] = response
	p.mu.Unlock()

	// log the time consumed for the download
	elapsed := time.Since(start)
	log.Printf("Download completed in %s\n", elapsed)

	return response, nil
}
