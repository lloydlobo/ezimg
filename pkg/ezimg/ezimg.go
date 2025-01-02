package ezimg

import (
	"image"
	"image/color"
	"image/jpeg"
	"log/slog"
	"os"

	"github.com/nfnt/resize"

	"github.com/lloydlobo/ezimg/pkg/utils"
)

func Read(path string) image.Image {
	file := utils.Must(os.Open(path))
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		slog.Error("Read: while decoding", "error", err.Error())
		os.Exit(1)
	}
	return img
}

func Write(path string, img image.Image) {
	file := utils.Must(os.Create(path))
	defer file.Close()

	if err := jpeg.Encode(file, img, nil); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func Grayscale(img image.Image) image.Image {
	rect := img.Bounds()
	gray := image.NewGray(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			prev := img.At(x, y)
			next := color.GrayModel.Convert(prev)
			gray.Set(x, y, next)
		}
	}
	return gray
}

func Resize(img image.Image, w uint, h uint) image.Image {
	return resize.Resize(w, h, img, resize.Lanczos3)
}
