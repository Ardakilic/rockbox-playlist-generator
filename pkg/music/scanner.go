package music

import (
	"os"
	"path/filepath"
	"strings"
)

// Track represents a music track with its metadata
type Track struct {
	Path        string
	Artist      string
	AlbumArtist string
	Title       string
	Album       string
	Year        int
	Duration    int
}

// Scanner handles scanning and processing music files
type Scanner struct {
	rootPath string
}

// NewScanner creates a new music scanner
func NewScanner(rootPath string) *Scanner {
	return &Scanner{
		rootPath: rootPath,
	}
}

// ScanDirectory scans a directory for music files
func (s *Scanner) ScanDirectory() ([]Track, error) {
	var tracks []Track
	err := filepath.Walk(s.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isMusicFile(path) {
			track, err := s.processFile(path)
			if err != nil {
				return err
			}
			tracks = append(tracks, track)
		}

		return nil
	})

	return tracks, err
}

// isMusicFile checks if the file is a supported music format
func isMusicFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	supportedFormats := []string{".mp3", ".flac", ".aac", ".alac", ".ogg", ".wav", ".m4a"}
	for _, format := range supportedFormats {
		if ext == format {
			return true
		}
	}
	return false
}

// processFile reads metadata from a music file
func (s *Scanner) processFile(path string) (Track, error) {
	// TODO: Implement file processing using tag package
	return Track{}, nil
}
