package utils

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strconv"
	src "vid/src"
)

type VideoResolution string

const (
	LOW  VideoResolution = "low"
	MID  VideoResolution = "mid"
	HIGH VideoResolution = "high"
)

type Video struct {
	file       *src.File
	resolution VideoResolution
}

func NewVideo(file *src.File, resolution VideoResolution) *Video {
	return &Video{
		file:       file,
		resolution: resolution,
	}
}

func (v *Video) CreateHLS(resolution VideoResolution) error {
	scale := fmt.Sprintf("scale=%s", "640:360")
	switch resolution {
	case LOW:
		v.resolution = LOW
		scale = fmt.Sprintf("scale=%s", "640:360")
	case MID:
		v.resolution = MID
		scale = fmt.Sprintf("scale=%s", "1280:720")
	case HIGH:
		v.resolution = HIGH
		scale = fmt.Sprintf("scale=%s", "1920:1080")
	default:
		v.resolution = LOW
		scale = fmt.Sprintf("scale=%s", "640:360")
	}
	return CreateHLS(v, scale)
}

// CreateHLS :w
func CreateHLS(video *Video, scale string) error {
	file := video.file
	segmentDuration := 10
	target_dir_err := src.ValidateDir("processed")
	if target_dir_err != nil {
		return target_dir_err
	}
	outDir := fmt.Sprintf("processed/%s", file.Filename)
	file.ProcessedDir = outDir
	validate_dir_err := src.ValidateDir(outDir)
	if validate_dir_err != nil {
		return validate_dir_err
	}
	resDir := fmt.Sprintf("%s/%s", outDir, video.resolution)
	validate_dir_err = src.ValidateDir(resDir)
	if validate_dir_err != nil {
		return validate_dir_err
	}
	input := file.Path
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", string(input),
		"-profile:v", "baseline", // baseline profile is compatible with most devices
		"-level", "3.0",
		"-vf", scale,
		"-start_number", "0", // start numbering segments from 0
		"-hls_time", strconv.Itoa(segmentDuration), // duration of each segment in seconds
		"-hls_list_size", "0", // keep all segments in the playlist
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", resDir),
	)

	output, process_error := ffmpegCmd.CombinedOutput()
	if process_error != nil {
		slog.Error("Error Processing File",
			"error", process_error,
		)
		slog.Error("Output",
			"output", string(output),
		)
		return process_error
	}
	return nil
}
