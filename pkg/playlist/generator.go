package playlist

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ardakilic/rockbox-playlist-generator/pkg/music"
)

// Generator handles playlist generation
type Generator struct {
	outputDir string
}

// NewGenerator creates a new playlist generator
func NewGenerator(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
	}
}

// GeneratePlaylist creates a Rockbox-compatible M3U playlist for an artist
func (g *Generator) GeneratePlaylist(artist string, tracks []music.Track) error {
	// Create a sanitized filename for the artist
	filename := sanitizeFilename(artist)
	playlistPath := filepath.Join(g.outputDir, fmt.Sprintf("rockbox_%s.m3u", filename))

	// Create the playlist file
	file, err := os.Create(playlistPath)
	if err != nil {
		return fmt.Errorf("failed to create playlist file: %w", err)
	}
	defer file.Close()

	// Write the M3U header
	if _, err := file.WriteString("#EXTM3U\n"); err != nil {
		return fmt.Errorf("failed to write M3U header: %w", err)
	}

	// Write each track
	for _, track := range tracks {
		// Write the EXTINF line with metadata
		extinf := fmt.Sprintf("#EXTINF:%d,%s - %s (%d - %s)\n",
			track.Duration,
			track.Artist,
			track.Title,
			track.Year,
			track.Album)
		if _, err := file.WriteString(extinf); err != nil {
			return fmt.Errorf("failed to write EXTINF line: %w", err)
		}

		// Write the file path
		if _, err := file.WriteString(track.Path + "\n"); err != nil {
			return fmt.Errorf("failed to write file path: %w", err)
		}
	}

	return nil
}

// sanitizeFilename creates a safe filename from an artist name
func sanitizeFilename(name string) string {
	// Replace invalid filename characters
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name
	for _, char := range invalid {
		result = strings.ReplaceAll(result, char, "_")
	}
	return strings.ToLower(result)
} 