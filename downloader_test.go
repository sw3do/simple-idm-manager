package main

import (
	"testing"
)

func TestNewDownloader(t *testing.T) {
	connections := 4
	downloader := NewDownloader(connections)

	if downloader == nil {
		t.Fatal("NewDownloader returned nil")
	}

	if downloader.connections != connections {
		t.Errorf("Expected %d connections, got %d", connections, downloader.connections)
	}

	if downloader.client == nil {
		t.Fatal("HTTP client is nil")
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
	}

	for _, test := range tests {
		result := formatBytes(test.input)
		if result != test.expected {
			t.Errorf("formatBytes(%d) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestProgressTracker(t *testing.T) {
	tracker := &ProgressTracker{
		totalSize: 1000,
	}

	tracker.AddProgress(100)
	if tracker.downloaded != 100 {
		t.Errorf("Expected downloaded to be 100, got %d", tracker.downloaded)
	}

	tracker.AddProgress(200)
	if tracker.downloaded != 300 {
		t.Errorf("Expected downloaded to be 300, got %d", tracker.downloaded)
	}
}

func TestProgressTrackerWrite(t *testing.T) {
	tracker := &ProgressTracker{}
	data := []byte("test data")

	n, err := tracker.Write(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if n != len(data) {
		t.Errorf("Expected %d bytes written, got %d", len(data), n)
	}

	if tracker.downloaded != int64(len(data)) {
		t.Errorf("Expected downloaded to be %d, got %d", len(data), tracker.downloaded)
	}
}
