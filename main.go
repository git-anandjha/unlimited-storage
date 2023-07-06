package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	videoWidth   = 1080
	videoHeight  = 1080
	pixelSize    = 1
	frameWidth   = videoWidth / pixelSize
	frameHeight  = videoHeight / pixelSize
	bytesPerLine = frameWidth / 8
	fps          = 24
)

// ConvertToBinaryPixels reads the input file and converts its content into binary pixels, saving them as frames in an MP4 video file.
func ConvertToBinaryPixels(inputFile string) error {
	// Read file content
	fileData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Calculate the number of frames needed
	numFrames := len(fileData) * 8 / (frameWidth * frameHeight)

	// Create the frames slice
	frames := make([][]byte, numFrames)

	// Generate frames with binary pixels
	for i := 0; i < numFrames; i++ {
		frame := make([]byte, frameWidth*frameHeight)

		for y := 0; y < frameHeight; y++ {
			for x := 0; x < frameWidth; x++ {
				index := y*frameWidth + x
				bitIndex := i*frameWidth*frameHeight + index
				byteIndex := bitIndex / 8
				bitOffset := bitIndex % 8

				if byteIndex < len(fileData) {
					bit := (fileData[byteIndex] >> uint(7-bitOffset)) & 1
					frame[index] = bit
				}
			}
		}

		frames[i] = frame
	}

	// Create the video file
	outputFile := "output.mp4"
	cmd := exec.Command(
		"ffmpeg",
		"-f", "rawvideo",
		"-pixel_format", "monob",
		"-video_size", fmt.Sprintf("%dx%d", videoWidth, videoHeight),
		"-framerate", fmt.Sprintf("%d", fps),
		"-i", "-",
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		outputFile,
	)
	cmd.Stdin = createFramesStream(frames)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create video: %w", err)
	}

	return nil
}

// createFramesStream creates a stream of frames from the given frames slice.
func createFramesStream(frames [][]byte) *os.File {
	r, w, _ := os.Pipe()

	go func() {
		defer w.Close()

		for _, frame := range frames {
			w.Write(frame)
		}
	}()

	return r
}

// DecodeVideoToBinaryPixels decodes the MP4 video file and outputs the original file.
func DecodeVideoToBinaryPixels(inputFile string, outputFile string) error {
	// Create temporary frames directory
	framesDir := "temp_frames"
	if err := os.Mkdir(framesDir, 0755); err != nil {
		return fmt.Errorf("failed to create temporary directory: %w", err)
	}
	defer os.RemoveAll(framesDir)

	// Extract frames from the video
	cmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-vf", fmt.Sprintf("fps=%d", fps),
		filepath.Join(framesDir, "frame%03d.png"),
	)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract frames from video: %w", err)
	}

	// Read frames and convert to binary data
	var binaryData []byte
	frameFiles, err := ioutil.ReadDir(framesDir)
	if err != nil {
		return fmt.Errorf("failed to read frame files: %w", err)
	}

	for _, frameFile := range frameFiles {
		framePath := filepath.Join(framesDir, frameFile.Name())
		frameData, err := os.ReadFile(framePath)
		if err != nil {
			return fmt.Errorf("failed to read frame file: %w", err)
		}

		// Process the frame data and convert to binary
		for _, pixelValue := range frameData {
			// Convert the pixel value to binary (0 or 1)
			bit := byte(0)
			if pixelValue != 0 {
				bit = 1
			}

			// Append the binary value to the binary data
			binaryData = append(binaryData, bit)
		}
	}

	// Write binary data to the output file
	if err := os.WriteFile(outputFile, binaryData, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

func main() {
	// ConvertToBinaryPixels("/Users/anandjha/Documents/SELF/VIDEO-CODEC/abc.txt")
	err := DecodeVideoToBinaryPixels("output.mp4", "output_file.txt")
	if err != nil {
		fmt.Println("Failed to decode video:", err)
	} else {
		fmt.Println("Video decoded successfully.")
	}
}
