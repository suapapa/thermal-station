package draw

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

func DrawLines(ff font.Face, msg string, maxWidth int) (image.Image, error) {
	lines := FitToLines(ff, maxWidth, msg)
	if len(lines) == 0 {
		return nil, fmt.Errorf("no conents")
	}
	fontHeight := int(ff.Metrics().Height) >> 6
	lineHeight := fontHeight
	// lineHeight += 4 // TODO: hard-coded line height
	pagePadding := 16

	w, h := maxWidth, lineHeight*len(lines)+pagePadding
	x, y := w/2, lineHeight/2

	dc := gg.NewContext(w, h)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(ff)
	for _, txt := range lines {
		dc.DrawStringAnchored(txt, float64(x), float64(y), 0.5, 0.5)
		y += lineHeight
	}
	return dc.Image(), nil
}
