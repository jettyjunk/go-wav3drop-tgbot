package service

import (
	"net/url"
	"strings"
)

func ValidateYouTubeURL(rawURL string) error {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	if parsed.Scheme != "https" && parsed.Scheme != "http" {
		return ErrInvalidURL
	}

	if parsed.Hostname() != "youtube.com" && parsed.Hostname() != "www.youtube.com" && parsed.Hostname() != "youtu.be" {
		return ErrInvalidHost
	}

	switch parsed.Hostname() {
	case "youtube.com", "www.youtube.com":
		if parsed.Path != "/watch" {
			return ErrInvalidURL
		}

		videoID := parsed.Query().Get("v")
		if !ValidateVideoID(videoID) {
			return ErrInvalidVideoID
		}

	case "youtu.be":
		videoID := strings.TrimPrefix(parsed.Path, "/")
		if !ValidateVideoID(videoID) {
			return ErrInvalidVideoID
		}
	}

	return nil
}

func ValidateVideoID(videoID string) bool {
	if len(videoID) != 11 {
		return false
	}

	for _, ch := range videoID {
		isLower := ch >= 'a' && ch <= 'z'
		isUpper := ch >= 'A' && ch <= 'Z'
		isDigit := ch >= '0' && ch <= '9'
		isSpecial := ch == '_' || ch == '-'

		if isLower || isUpper || isDigit || isSpecial {
			continue
		}

		return false
	}

	return true
}
