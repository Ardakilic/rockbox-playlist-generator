package lastfm

import (
	"fmt"
	"strconv"

	"github.com/shkh/lastfm-go/lastfm"
)

// Client represents a Last.fm API client
type Client struct {
	api *lastfm.Api
}

// NewClient creates a new Last.fm client with the provided API credentials
func NewClient(apiKey, apiSecret string) (*Client, error) {
	if apiKey == "" || apiSecret == "" {
		return nil, fmt.Errorf("API key and secret are required")
	}

	api := lastfm.New(apiKey, apiSecret)
	return &Client{api: api}, nil
}

// Track represents a Last.fm track
type Track struct {
	Name   string `json:"name"`
	Artist struct {
		Name string `json:"name"`
	} `json:"artist"`
	Duration int `json:"duration"`
}

// GetTopTracks fetches top tracks for an artist
func (c *Client) GetTopTracks(artist string, limit int) ([]Track, error) {
	params := lastfm.P{
		"artist": artist,
		"limit":  limit,
	}

	result, err := c.api.Artist.GetTopTracks(params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top tracks: %w", err)
	}

	// Convert the library's response to our Track format
	tracks := make([]Track, 0, len(result.Tracks))
	for _, t := range result.Tracks {
		// Convert Duration from string to int (Last.fm returns it as string)
		duration := 0
		if t.Duration != "" {
			if durationInt, err := strconv.Atoi(t.Duration); err == nil {
				duration = durationInt
			}
		}

		track := Track{
			Name:     t.Name,
			Duration: duration,
		}
		// Set the artist name from the top-level of the result
		track.Artist.Name = result.Artist
		tracks = append(tracks, track)
	}

	return tracks, nil
}
