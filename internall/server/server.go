package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Server is a struct that represents an HTTP server.
type Server struct {
	port    string
	service VideoService
}

// NewServer creates a new instance of the Server struct with provided port and service.
func NewServer(port string, service VideoService) *Server {
	return &Server{port: port, service: service}
}

// Start starts the HTTP server.
func (s *Server) Start() {
	// Define the API endpoints
	http.HandleFunc("/meta-info", s.handleMetaInfo)
	http.HandleFunc("/download-links", s.handleDownloadLinks)

	// Start the server
	log.Printf("Starting server on port %s\n", s.port)
	if err := http.ListenAndServe(":"+s.port, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}

// handleMetaInfo is an HTTP handler that returns the meta-information about the video from the given URL.
func (s *Server) handleMetaInfo(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameter from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	// Get the meta-information about the video
	info, err := s.service.MetaInformation(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the meta-information into JSON and write it to the response
	if err = json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleDownloadLinks is an HTTP handler that returns the download links for the video from the given URL.
func (s *Server) handleDownloadLinks(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameter from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	// Get the download links for the video
	links, err := s.service.DownloadLinks(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the download links into JSON and write it to the response
	if err = json.NewEncoder(w).Encode(links); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
