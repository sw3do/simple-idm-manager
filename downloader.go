package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Downloader struct {
	connections int
	client      *http.Client
}

type ProgressTracker struct {
	totalSize  int64
	downloaded int64
	startTime  time.Time
	mu         sync.Mutex
}

func NewDownloader(connections int) *Downloader {
	return &Downloader{
		connections: connections,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (d *Downloader) Download(url, output string, resume bool) error {
	if output == "" {
		output = filepath.Base(url)
		if output == "/" || output == "" {
			output = "download"
		}
	}

	resp, err := d.client.Head(url)
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}
	resp.Body.Close()

	fileSize := resp.ContentLength
	supportsRange := resp.Header.Get("Accept-Ranges") == "bytes"

	if fileSize <= 0 {
		return d.downloadSingle(url, output)
	}

	if !supportsRange || d.connections <= 1 {
		return d.downloadSingle(url, output)
	}

	return d.downloadMultipart(url, output, fileSize, resume)
}

func (d *Downloader) downloadSingle(url, output string) error {
	resp, err := d.client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	tracker := &ProgressTracker{
		totalSize: resp.ContentLength,
		startTime: time.Now(),
	}

	go d.showProgress(tracker)

	_, err = io.Copy(io.MultiWriter(file, tracker), resp.Body)
	return err
}

func (d *Downloader) downloadMultipart(url, output string, fileSize int64, resume bool) error {
	chunkSize := fileSize / int64(d.connections)
	if chunkSize == 0 {
		return d.downloadSingle(url, output)
	}

	var existingSize int64
	if resume {
		if stat, err := os.Stat(output); err == nil {
			existingSize = stat.Size()
		}
	}

	if existingSize >= fileSize {
		fmt.Println("File already downloaded")
		return nil
	}

	tracker := &ProgressTracker{
		totalSize:  fileSize,
		downloaded: existingSize,
		startTime:  time.Now(),
	}

	go d.showProgress(tracker)

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	errorChan := make(chan error, d.connections)

	for i := 0; i < d.connections; i++ {
		start := int64(i) * chunkSize
		end := start + chunkSize - 1
		if i == d.connections-1 {
			end = fileSize - 1
		}

		if start < existingSize {
			if end < existingSize {
				continue
			}
			start = existingSize
		}

		wg.Add(1)
		go func(start, end int64) {
			defer wg.Done()
			err := d.downloadChunk(url, file, start, end, tracker)
			if err != nil {
				errorChan <- err
			}
		}(start, end)
	}

	wg.Wait()
	close(errorChan)

	if err := <-errorChan; err != nil {
		return err
	}

	fmt.Println("\nDownload completed successfully!")
	return nil
}

func (d *Downloader) downloadChunk(url string, file *os.File, start, end int64, tracker *ProgressTracker) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	resp, err := d.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buffer := make([]byte, 32*1024)
	offset := start

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			_, writeErr := file.WriteAt(buffer[:n], offset)
			if writeErr != nil {
				return writeErr
			}
			offset += int64(n)
			tracker.AddProgress(int64(n))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProgressTracker) Write(data []byte) (int, error) {
	n := len(data)
	p.AddProgress(int64(n))
	return n, nil
}

func (p *ProgressTracker) AddProgress(bytes int64) {
	p.mu.Lock()
	p.downloaded += bytes
	p.mu.Unlock()
}

func (d *Downloader) showProgress(tracker *ProgressTracker) {
	for {
		tracker.mu.Lock()
		downloaded := tracker.downloaded
		total := tracker.totalSize
		tracker.mu.Unlock()

		if total > 0 {
			percent := float64(downloaded) / float64(total) * 100
			elapsed := time.Since(tracker.startTime)
			speed := float64(downloaded) / elapsed.Seconds()

			var eta string
			if speed > 0 {
				remaining := float64(total-downloaded) / speed
				eta = time.Duration(remaining * float64(time.Second)).String()
			} else {
				eta = "--"
			}

			fmt.Printf("\rProgress: %.1f%% [%s/%s] Speed: %s/s ETA: %s",
				percent,
				formatBytes(downloaded),
				formatBytes(total),
				formatBytes(int64(speed)),
				eta)
		} else {
			fmt.Printf("\rDownloaded: %s", formatBytes(downloaded))
		}

		if downloaded >= total && total > 0 {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
