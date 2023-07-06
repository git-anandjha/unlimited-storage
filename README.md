# File-to-Video Converter

This repository contains a Go program that allows you to convert any file into a video file. The video file can then be uploaded and stored on platforms like YouTube, acting as unlimited storage for your files. You can later retrieve the original file by decoding the video.

## Usage

1. Clone this repository:

   ```bash
   git clone https://github.com/your-username/your-repo.git
   ```

2. Navigate to the cloned directory:

   ```bash
   cd your-repo
   ```

3. Install the required Go dependencies:

   ```bash
   go mod download
   ```

4. Place the file you want to convert in the same directory as the Go program.

5. Modify the `main` function in the `main.go` file to specify the path of your input file and the desired output file names.

6. Build and run the Go program:

   ```bash
   go run main.go
   ```

   This will convert the input file to a video file and save it as `output.mp4`.

7. Upload the generated `output.mp4` file to your preferred video platform (e.g., YouTube).

8. To decode the video file and retrieve the original file, modify the `main` function to call the `DecodeVideoToBinaryPixels` function with the appropriate input and output file paths.

9. Build and run the Go program again:

   ```bash
   go run main.go
   ```

   This will decode the video file and save the original file as `output_file.txt`.

## File vs. Video Size Ratio

The size of the resulting video file depends on the size of the input file. The ratio between the input file size and the video file size can be calculated using the following formula:

```
videoFileSize = (inputFileSize * 8 * frameWidth * frameHeight) / (frameWidth * frameHeight * 8 + metadataSize)
```

In the given code, the `frameWidth` and `frameHeight` values are set to 1080, and the `metadataSize` represents the size of additional metadata in the video file. The `metadataSize` value can vary based on the video encoding parameters and other factors.

For example, let's consider a scenario where the input file size is 10 MB. Assuming a metadata size of 1 MB, we can calculate the approximate size of the resulting video file:

```
inputFileSize = 10 MB
metadataSize = 1 MB
frameWidth = 1080
frameHeight = 1080

videoFileSize = (10 * 8 * 1080 * 1080) / (1080 * 1080 * 8 + 1)
videoFileSize = 10 MB * 8 * 1080 * 1080 / (1080 * 1080 * 8 + 1)
videoFileSize = 10 * 8 MB / 1
videoFileSize = 80 MB
```

In this example, the resulting video file would be approximately 80 MB.

Please note that this calculation is an approximation, and the actual video file size can vary based on various factors, including the compression efficiency of the video encoding algorithm used.

## Contributing

Contributions to this repository are welcome. If you find any issues or want to add new features, please submit a pull request.

## License

This repository is licensed under the [MIT License](LICENSE). Feel free to use and modify the code as per your needs.
