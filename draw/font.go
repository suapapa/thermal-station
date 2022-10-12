package draw

import (
	"embed"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	//go:embed font/MaruBuri-Regular.ttf
	efs embed.FS

	fonts map[float64]font.Face
)

func init() {
	fonts = make(map[float64]font.Face)
}

func GetFont(points float64) (font.Face, error) {
	if ff, ok := fonts[points]; ok {
		return ff, nil
	}

	fontName := "font/MaruBuri-Regular.ttf"
	data, err := efs.ReadFile(fontName)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(data)
	if err != nil {
		return nil, err
	}

	nface := truetype.NewFace(f, &truetype.Options{
		Size:    points,
		Hinting: font.HintingFull,
		// Hinting: font.HintingNone,
	})
	fonts[points] = nface
	return nface, nil
}
