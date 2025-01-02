// Adapted from https://github.com/code-heim/go_21_goroutines_pipeline/blob/main/main.go
package main

import (
	"fmt"
	"image"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/lloydlobo/ezimg/pkg/ezimg"
	"github.com/lloydlobo/ezimg/pkg/utils"
)

var (
	version = "0.1.0"

	projectDir = utils.Must(os.Getwd())
	imageDir   = filepath.Join(projectDir, "images")
	imagePaths = []string{
		filepath.Join(imageDir, "image1.jpg"),
		filepath.Join(imageDir, "image2.jpg"),
		filepath.Join(imageDir, "image3.jpg"),
		filepath.Join(imageDir, "image4.jpg"),
	}
)

const banner = `
 ___ ___ _ __ __  __  
| __|_  | |  V  |/ _] 
| _| / /| | \_/ | [/\ 
|___|___|_|_| |_|\__/ `

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	pink   = "\033[35m"
	cyan   = "\033[36m"
)

func main() {
	startTime := time.Now()

	fmt.Println(banner, "("+version+"),", "built with Go", runtime.Version()+"\n")
	fmt.Printf("%swatching %v/%s\n", cyan, filepath.Base(imageDir), reset)

	chan1 := loadImage(imagePaths)

	chan2 := resizeImage(chan1)
	chan3 := convertToGrayscale(chan2)

	results := saveImage(chan3)

	for success := range results {
		if !success {
			fmt.Printf("%sfailed%s\n", yellow, reset)
			continue
		}
		fmt.Printf("%ssuccess%s\n", cyan, reset)
	}

	fmt.Printf("%scleaning...%s\n", pink, reset)
	fmt.Printf("%ssee you again~%s\n", pink, reset)

	slog.Info(fmt.Sprintf("%stook %v%s\n", green, time.Since(startTime), reset))

}

type Job struct {
	Image      image.Image
	InputPath  string
	OutputPath string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, path := range paths {
			fmt.Printf("%sloading %s%s\n", yellow, filepath.Base(path), reset)

			outpath := strings.Replace(path, "images/", "images/output/", 1)
			out <- Job{
				Image:      ezimg.Read(path),
				InputPath:  path,
				OutputPath: outpath,
			}
		}
		close(out)
	}()
	return out
}

func resizeImage(c <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		w := uint(500)
		h := uint(500)
		for job := range c {
			fmt.Printf("%sresizing %s%s\n", green, filepath.Base(job.InputPath), reset)
			slog.Debug("Resizing...", "image", filepath.Base(job.InputPath),
				"from", job.Image.Bounds(),
				"to", image.Rect(0, 0, int(w), int(h)))
			job.Image = ezimg.Resize(job.Image, w, h)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(c <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range c {
			fmt.Printf("%sconverting grayscale %s%s\n", green, filepath.Base(job.InputPath), reset)
			job.Image = ezimg.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(c <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range c {
			fmt.Printf("%ssaving %s%s\n", yellow, filepath.Base(job.OutputPath), reset)
			ezimg.Write(job.OutputPath, job.Image)
			out <- true // Success!
		}
		close(out)
	}()
	return out
}
