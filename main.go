package main

import (
    "log"
    "os"
    "os/exec"
    "path/filepath"
)

func main() {
    inputFolderPath := "tmp/videos"
    outputFolderPath := "tmp/videos/output"
    videoBitrate := "1000k"
    audioBitrate := "128k"

    err := filepath.Walk(inputFolderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() && filepath.Ext(path) == ".mp4" {
            compressedFilePath := filepath.Join(outputFolderPath, filepath.Base(path))
            
            cmd := exec.Command("ffmpeg", "-i", path, "-b:v", videoBitrate, "-b:a", audioBitrate, "-y", compressedFilePath)
            if err := cmd.Run(); err != nil {
                return err
            }
            log.Printf("File %s compressed successfully to %s", path, compressedFilePath)
        }
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking through input folder: %v", err)
    }
}
