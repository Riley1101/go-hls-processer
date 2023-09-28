package utils

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strconv"
	src "vid/src"
)

func CreateHLS(file *src.File, segmentDuration int) error {
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

	input := file.Path
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", string(input),
		"-profile:v", "baseline", // baseline profile is compatible with most devices
		"-level", "3.0",
		"-start_number", "0", // start numbering segments from 0
		"-hls_time", strconv.Itoa(segmentDuration), // duration of each segment in seconds
		"-hls_list_size", "0", // keep all segments in the playlist
		"-f", "hls",
		fmt.Sprintf("%s/playlist.m3u8", outDir),
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
