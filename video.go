// package main

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"os"
// 	"sync"
// )

// const (
// 	imageWidth    = 1080
// 	imageHeight   = 1080
// 	pixelSize     = 10
// 	framesPerFile = 1 // Number of frames per temporary PNG file
// 	fps           = 24  // Frames per second of the final video
// )

// func main() {
// 	inputFilePath := "/Users/anandjha/Documents/SELF/VIDEO-CODEC/sample-pdf-with-images.pdf"
// 	outputFolderPath := "frames"

// 	err := convertToFrames(inputFilePath, outputFolderPath)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// func convertToFrames(inputFilePath, outputFolderPath string) error {
// 	// Read the input file as binary data
// 	inputData, err := readFile(inputFilePath)
// 	if err != nil {
// 		return err
// 	}

// 	// Create the output folder if it doesn't exist
// 	err = os.MkdirAll(outputFolderPath, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	// Calculate the number of frames needed
// 	numFrames := len(inputData) / (imageWidth * imageHeight / 8)

// 	// Generate and save frames as individual PNG files
// 	var wg sync.WaitGroup
// 	for i := 0; i < numFrames; i += framesPerFile {
// 		endIndex := i + framesPerFile
// 		if endIndex > numFrames {
// 			endIndex = numFrames
// 		}

// 		wg.Add(1)
// 		go func(startIndex, endIndex int) {
// 			defer wg.Done()

// 			for j := startIndex; j < endIndex; j++ {
// 				frameIndex := j

// 				// Calculate the range of data for the current frame
// 				start := frameIndex * (imageWidth * imageHeight / 8)
// 				end := start + (imageWidth * imageHeight / 8)

// 				// Get the frame data for the current frame
// 				frameData := inputData[start:end]

// 				// Create the frame as a grayscale image
// 				img := createFrame(frameData)

// 				// Save the frame as a PNG file
// 				frameFilePath := fmt.Sprintf("%s/frame_%04d.png", outputFolderPath, frameIndex)
// 				err := saveImage(img, frameFilePath)
// 				if err != nil {
// 					fmt.Println("Error saving frame:", err)
// 				}
// 			}
// 		}(i, endIndex)
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	return nil
// }

// func readFile(filePath string) ([]byte, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fileSize := fileInfo.Size()
// 	data := make([]byte, fileSize)

// 	_, err = file.Read(data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func createFrame(frameData []byte) image.Image {
// 	img := image.NewGray(image.Rect(0, 0, imageWidth*pixelSize, imageHeight*pixelSize))

// 	for i, b := range frameData {
// 		x := (i * pixelSize) % (imageWidth * pixelSize)
// 		y := (i * pixelSize) / (imageWidth * pixelSize) * pixelSize
// 		g := color.Gray{Y: b}

// 		for py := 0; py < pixelSize; py++ {
// 			for px := 0; px < pixelSize; px++ {
// 				img.SetGray(x+px, y+py, g)
// 			}
// 		}
// 	}

// 	return img
// }

// func saveImage(img image.Image, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	err = png.Encode(file, img)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
