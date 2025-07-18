# Simple IDM Manager

[![CI/CD](https://github.com/sw3do/simple-idm-manager/actions/workflows/ci.yml/badge.svg)](https://github.com/sw3do/simple-idm-manager/actions/workflows/ci.yml)
[![Release](https://github.com/sw3do/simple-idm-manager/actions/workflows/release.yml/badge.svg)](https://github.com/sw3do/simple-idm-manager/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sw3do/simple-idm-manager)](https://goreportcard.com/report/github.com/sw3do/simple-idm-manager)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A CLI-based Internet Download Manager written in Go that supports concurrent downloads, progress tracking, and resume functionality.

## Features

- **Concurrent Downloads**: Download files using multiple connections for faster speeds
- **Progress Tracking**: Real-time progress bar with download speed and ETA
- **Resume Support**: Resume interrupted downloads
- **Range Request Support**: Automatically detects if server supports partial downloads
- **Cross-platform**: Works on Windows, macOS, and Linux

## Installation

```bash
# Clone the repository
git clone https://github.com/sw3do/simple-idm-manager.git
cd simple-idm-manager

# Build the application
go build -o simple-idm
```

### Using Go Install
```bash
go install github.com/sw3do/simple-idm-manager@latest
```

### Download Pre-built Binaries
Download the latest release from the [releases page](https://github.com/sw3do/simple-idm-manager/releases).

## Usage

### Basic Download
```bash
./simple-idm -url https://example.com/file.zip
```

### Specify Output File
```bash
./simple-idm -url https://example.com/file.zip -output myfile.zip
```

### Use Multiple Connections
```bash
./simple-idm -url https://example.com/file.zip -connections 8
```

### Resume Download
```bash
./simple-idm -url https://example.com/file.zip -resume
```

### Complete Example
```bash
./simple-idm -url https://releases.ubuntu.com/22.04/ubuntu-22.04.3-desktop-amd64.iso -output ubuntu.iso -connections 8 -resume
```

## Command Line Options

- `-url`: URL of the file to download (required)
- `-output`: Output file path (optional, defaults to filename from URL)
- `-connections`: Number of concurrent connections (default: 8)
- `-resume`: Resume incomplete download (default: false)
- `-version`: Show version information

## How It Works

1. **File Info Detection**: The application first sends a HEAD request to get file size and check if the server supports range requests
2. **Single vs Multi-part**: If range requests are supported, the file is split into chunks for concurrent download
3. **Progress Tracking**: Real-time progress is displayed with download speed and completion percentage
4. **Chunk Merging**: Downloaded chunks are automatically merged into the final file
5. **Resume Capability**: Incomplete downloads can be resumed from where they left off

## Examples

### Download a large file with 8 connections
```bash
./simple-idm -url https://example.com/largefile.zip -connections 8
```

### Resume a previously interrupted download
```bash
./simple-idm -url https://example.com/largefile.zip -resume
```

### Download to a specific location
```bash
./simple-idm -url https://example.com/file.pdf -output /path/to/downloads/document.pdf
```

### Check version
```bash
./simple-idm -version
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.