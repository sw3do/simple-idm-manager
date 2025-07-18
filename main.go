package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev"

func main() {
	var url string
	var output string
	var connections int
	var resume bool
	var showVersion bool

	flag.StringVar(&url, "url", "", "URL to download")
	flag.StringVar(&output, "output", "", "Output file path")
	flag.IntVar(&connections, "connections", 8, "Number of concurrent connections")
	flag.BoolVar(&resume, "resume", false, "Resume incomplete download")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.Parse()

	if showVersion {
		fmt.Printf("Simple IDM Manager %s\n", version)
		return
	}

	if url == "" {
		fmt.Println("Usage: simple-idm -url <URL> [-output <file>] [-connections <num>] [-resume]")
		fmt.Println("Example: simple-idm -url https://example.com/file.zip -output file.zip -connections 16")
		os.Exit(1)
	}

	downloader := NewDownloader(connections)
	err := downloader.Download(url, output, resume)
	if err != nil {
		fmt.Printf("Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download completed successfully!")
}
