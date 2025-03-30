package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Ardakilic/rockbox-playlist-generator/config"
	"github.com/Ardakilic/rockbox-playlist-generator/pkg/lastfm"
	"github.com/Ardakilic/rockbox-playlist-generator/pkg/music"
	"github.com/Ardakilic/rockbox-playlist-generator/pkg/playlist"
)

var (
	musicPath = flag.String("path", "", "Path to your local music library (required)")
	limit     = flag.Int("limit", 0, "Number of top tracks to fetch per artist (default: from config)")
)

func main() {
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if *musicPath == "" {
		fmt.Println("Error: -path flag is required")
		flag.Usage()
		os.Exit(1)
	}

	// Use command line limit if provided, otherwise use config default
	trackLimit := cfg.DefaultLimit
	if *limit > 0 {
		trackLimit = *limit
	}

	// Initialize Last.fm client
	lastfmClient, err := lastfm.NewClient(cfg.LastFMAPIKey, cfg.LastFMAPISecret)
	if err != nil {
		log.Fatalf("Failed to initialize Last.fm client: %v", err)
	}

	// Initialize music scanner
	scanner := music.NewScanner(*musicPath)

	// Scan for music files
	tracks, err := scanner.ScanDirectory()
	if err != nil {
		log.Fatalf("Failed to scan music directory: %v", err)
	}

	// Group tracks by artist
	artistTracks := make(map[string][]music.Track)
	for _, track := range tracks {
		artist := track.AlbumArtist
		if artist == "" {
			artist = track.Artist
		}
		artistTracks[artist] = append(artistTracks[artist], track)
	}

	// Create output directory for playlists
	outputDir := filepath.Join(*musicPath, "playlists")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Initialize playlist generator
	generator := playlist.NewGenerator(outputDir)

	// Process each artist
	for artist, tracks := range artistTracks {
		log.Printf("Processing artist: %s", artist)

		// Get top tracks from Last.fm
		topTracks, err := lastfmClient.GetTopTracks(artist, trackLimit)
		if err != nil {
			log.Printf("Warning: Failed to fetch top tracks for %s: %v", artist, err)
			continue
		}

		// Match Last.fm tracks with local tracks
		var matchedTracks []music.Track
		for _, topTrack := range topTracks {
			for _, localTrack := range tracks {
				if strings.EqualFold(localTrack.Title, topTrack.Name) {
					matchedTracks = append(matchedTracks, localTrack)
					break
				}
			}
		}

		// Generate playlist if we have matches
		if len(matchedTracks) > 0 {
			if err := generator.GeneratePlaylist(artist, matchedTracks); err != nil {
				log.Printf("Warning: Failed to generate playlist for %s: %v", artist, err)
				continue
			}
			log.Printf("Generated playlist for %s with %d tracks", artist, len(matchedTracks))
		}
	}

	log.Println("Playlist generation completed!")
}
