// package main

// import (
// 	"fmt"
// 	"image"
// 	"image/color"
// 	"image/png"
// 	"io/ioutil"
// 	"math"
// 	"os"

// 	"github.com/disintegration/imaging"
// )

// const (
// 	imageWidth   = 1080
// 	pixelSize    = 10
// 	maxImageSize = math.MaxInt32
// )

// func convertFileToImage(filePath string) error {
// 	// Read the file
// 	data, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	// Calculate the image height based on the file size
// 	fileSize := len(data)
// 	imageHeight := int(math.Ceil(float64(fileSize) / (imageWidth * pixelSize)))

// 	// Create a new blank image
// 	img := image.NewGray(image.Rect(0, 0, imageWidth*pixelSize, imageHeight*pixelSize))

// 	// Iterate over the data and set the pixels in the image
// 	for i, b := range data {
// 		// Calculate the pixel coordinates
// 		x := (i * pixelSize) % (imageWidth * pixelSize)
// 		y := (i * pixelSize) / (imageWidth * pixelSize) * pixelSize

// 		// Set the grayscale color based on the byte value
// 		g := color.Gray{Y: b}

// 		// Set the color for each pixel in the pixelSize x pixelSize block
// 		for py := 0; py < pixelSize; py++ {
// 			for px := 0; px < pixelSize; px++ {
// 				img.SetGray(x+px, y+py, g)
// 			}
// 		}
// 	}

// 	// Resize the image to the desired dimensions
// 	resizedImg := imaging.Resize(img, imageWidth, imageHeight, imaging.NearestNeighbor)

// 	// Create the output file
// 	outputPath := "output.png"
// 	outputFile, err := os.Create(outputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Encode the image as PNG and save it to the file
// 	err = png.Encode(outputFile, resizedImg)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Image saved as", outputPath)
// 	return nil
// }

// func main() {
// 	// Example usage
// 	filePath := "/Users/anandjha/Documents/SELF/VIDEO-CODEC/sample-pdf-with-images.pdf"

// 	err := convertFileToImage(filePath)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
