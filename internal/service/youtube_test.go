package service

import "testing"

func TestValidateYouTubeURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid youtube URL",
			url:     "https://youtube.com/watch?v=dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "valid short URL",
			url:     "https://youtu.be/dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "missing video ID",
			url:     "https://youtube.com/watch",
			wantErr: true,
		},
		{
			name:    "invalid host",
			url:     "https://google.com/watch?v=qQw4w9WgXcQ",
			wantErr: true,
		},
		{
			name:    "missing short vidio ID",
			url:     "https://youtu.be/",
			wantErr: true,
		},
		{
			name:    "invalid URL format",
			url:     "hello worl",
			wantErr: true,
		},
		{
			name:    "invalid short video ID",
			url:     "https://youtu.be/abc/def",
			wantErr: true,
		},
		{
			name:    "invalid regular video ID",
			url:     "https://youtube.com/watch?v=abc",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateYouTubeURL(tt.url)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateYouTubeURL(%q) error = %v, wantErr = %v", tt.url, err, tt.wantErr)
			}
		})
	}

}
