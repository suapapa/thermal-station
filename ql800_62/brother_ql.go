package ql800_62

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
)

const (
	MaxWidth = 696 // 62mm endless
)

var (
	tmpPngPath = path.Join(os.TempDir(), "ql-800_temp.png")
)

func PrintLabel(dev string, img image.Image, rotate int) error {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	if w > MaxWidth {
		h = (h * MaxWidth) / w
		w = MaxWidth
		img = resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
	}

	// log.Printf("save label image to %s", tmpPngPath)
	if err := saveImg2Png(img, tmpPngPath); err != nil {
		return errors.Wrap(err, "fail to print from")
	}
	defer os.RemoveAll(tmpPngPath)

	err := exec.Command("sh", "-c",
		fmt.Sprintf("brother_ql -b linux_kernel -p %s -m QL-800 print -r %d -l 62 %s", dev, rotate, tmpPngPath)).Run()
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}
	return nil
}

func saveImg2Png(img image.Image, pngFN string) error {
	f, err := os.Create(pngFN)
	if err != nil {
		return errors.Wrap(err, "fail to savePNG")
	}
	defer f.Close()

	if err = png.Encode(f, img); err != nil {
		return errors.Wrap(err, "fail to savePNG")
	}

	return nil
}
