package service

import "errors"

var (
	ErrInvalidURL     = errors.New("invalid url")
	ErrInvalidHost    = errors.New("unsupported URL host")
	ErrInvalidVideoID = errors.New("missing YouTube video ID")
)
