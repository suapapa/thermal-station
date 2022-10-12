package draw

import (
	"fmt"
	"image"
	"image/color"
	"sort"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

func FitToLines(ff font.Face, maxWidth int, origTxt string) []string {
	origTxt = strings.ReplaceAll(origTxt, "\r\n", "\n")
	lines := strings.Split(origTxt, "\n")
	var outLines []string

	for _, line := range lines {
		rl := []rune(line)
		rlLen := len(rl)
		rlStart := 0
		for rlStart < rlLen {
			rlSub := rl[rlStart:]
			rlSubLen := len(rlSub)

			i := sort.Search(rlSubLen, func(i int) bool {
				w, _ := MeasureTxt(ff, string(rlSub[:i]))
				return w > maxWidth
			})

			if i > rlSubLen {
				i = rlSubLen
			}

			w, _ := MeasureTxt(ff, string(rlSub[:i]))
			for w > maxWidth {
				i -= 1
				w, _ = MeasureTxt(ff, string(rlSub[:i]))
			}

			// wordWrap
			origI := i
			var lastSpaceIdx int
			for ; i > 1; i-- {
				if rl[rlStart+i-1] == ' ' {
					lastSpaceIdx = i
					break
				}
			}
			if i <= 1 {
				i = origI
			} else if w, _ := MeasureTxt(ff, string(rl[rlStart:])); w < maxWidth {
				i = origI
			} else {
				i = lastSpaceIdx
			}

			sl := string(rl[rlStart : rlStart+i])
			sl = strings.TrimSpace(sl)
			// log.Printf("sl += '%s'", sl)
			outLines = append(outLines, sl)

			rlStart += i
		}
	}
	return outLines
}

func MeasureTxt(ff font.Face, txt string) (w, h int) {
	// log.Printf("measuring '%s'...", txt)
	d := &font.Drawer{
		Face: ff,
	}
	w = int(d.MeasureString(txt) >> 6)
	h = int(ff.Metrics().Height >> 6)
	w = ((w + 7) / 8) * 8 // set w to multiple 8
	h = ((h + 7) / 8) * 8 // set h to multiple 8
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}

	return
}

func Txt2Img(ff font.Face, w int, txt string) (image.Image, error) {
	if ff == nil {
		return nil, fmt.Errorf("nil font")
	}

	_, h := MeasureTxt(ff, txt)
	dc := gg.NewContext(w, h)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(ff)
	dc.DrawStringAnchored(txt, float64(w)/2, float64(h)/2, 0.5, 0.3)

	return dc.Image(), nil
}
