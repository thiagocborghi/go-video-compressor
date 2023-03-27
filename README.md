
# MP4 Compressor

This is a command-line tool that compresses all MP4 files in a specified folder to reduce their size without losing too much quality. The tool uses the FFmpeg library to compress the videos by reducing their bitrate.


## Usage

To compress all MP4 files in a folder, run the following command:

```bash
./mp4-compressor -i /tmp/videos -o /tmp/videos/output
```

Replace /tmp/videos with the path to the folder containing the MP4 files that you want to compress, and replace /tmp/videos/output with the path to the folder where you want to save the compressed files.

You can also set the video and audio bitrate for the compressed files using the -vb and -ab options. For example:

```bash
./mp4-compressor -i /tmp/videos -o /tmp/videos/output -vb 1000k -ab 128k
```

This will set the video bitrate to 1000k and the audio bitrate to 128k.

