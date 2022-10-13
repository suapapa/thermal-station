package receipt

import (
	"bufio"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/lestrrat-go/dither"
	"github.com/pkg/errors"
	"github.com/tarm/serial"
)

const (
	MaxWidth = 576
)

type Printer struct {
	c io.Closer
	w *bufio.Writer
}

// NewUSBPrinter make a new Printer from USB device path
// example args:
//
//	pavPath:
func NewUSBPrinter(devPath string) *Printer {
	dev, err := os.OpenFile(devPath, os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	bufW := bufio.NewWriter(dev)

	return &Printer{
		c: dev,
		w: bufW,
	}
}

// NewSerialPrinter make a new Printer from Serial device path
// example args:
//
//	pavPath: /dev/ttyACM0, speed: 9600
func NewSerialPrinter(devPath string, speed int) *Printer {
	c := &serial.Config{Name: devPath, Baud: speed}
	dev, err := serial.OpenPort(c)
	if err != nil {
		panic(err)
	}
	bufW := bufio.NewWriter(dev)

	return &Printer{
		c: dev,
		w: bufW,
	}
}

func (p *Printer) Close() error {
	return p.c.Close()
}

func (p *Printer) PrintImage8bitDouble(img image.Image) error {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()

	ditheredImg := dither.Monochrome(dither.Burkes, img, 1.18)
	dataBuf := make([]byte, (w*h+7)/8)

	// 가로방향 점의 개수: nL + nH x 256
	nH := byte(w / 256)
	nL := byte(w % 256)
	mode := byte(1)
	cmdBuf := []byte{0x1B, 0x2A, mode, nL, nH}

	dataBufIdx := 0
	for y := 0; y < h; y += 8 {
		for x := 0; x < w; x++ {
			var dataByte byte
			for yi := 0; yi < 8; yi++ {
				currY := y + yi
				if currY > h {
					continue
				}
				var bit byte
				if r, g, b, _ := ditheredImg.At(x, currY).RGBA(); r == 0 && g == 0 && b == 0 {
					bit = 1 << (7 - yi)
				}
				dataByte |= bit
			}
			if dataBufIdx >= len(dataBuf) {
				// log.Printf("dbIdx=%d, dbLen=%d", dataBufIdx, len(dataBuf))
				break
			}
			dataBuf[dataBufIdx] = dataByte
			dataBufIdx += 1
		}
	}
	return p.printBuf(cmdBuf, dataBuf, w)
}

func (p *Printer) PrintImage24bitDouble(img image.Image) error {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()

	ditheredImg := dither.Monochrome(dither.Burkes, img, 1.18)
	dataBuf := make([]byte, (w*h+7)/8)

	// 가로방향 점의 개수: nL + nH x 256
	nH := byte(w / 256)
	nL := byte(w % 256)
	mode := byte(33)
	// log.Println(nL, nH, mode)
	cmdBuf := []byte{0x1B, 0x2A, mode, nL, nH}

	dataBufIdx := 0
	for y := 0; y < h; y += (8 * 3) {
		for x := 0; x < w; x++ {
			for yCnt := 0; yCnt < 3; yCnt++ {
				var dataByte byte
				for yi := 0; yi < 8; yi++ {
					currY := y + (8 * yCnt) + yi
					if currY > h {
						continue
					}
					var bit byte
					if r, g, b, _ := ditheredImg.At(x, currY).RGBA(); r == 0 && g == 0 && b == 0 {
						bit = 1 << (7 - yi)
					}
					dataByte |= bit
				}
				if dataBufIdx >= len(dataBuf) {
					// log.Printf("dbIdx=%d, dbLen=%d", dataBufIdx, len(dataBuf))
					break
				}
				dataBuf[dataBufIdx] = dataByte
				dataBufIdx += 1
			}
		}
	}
	return p.printBuf(cmdBuf, dataBuf, w*3)
}

func (p *Printer) printBuf(cmdBuf, dataBuf []byte, widthDataLen int) error {
	// Standard mode
	p.w.Write([]byte{0x1B, 0x0C})

	// 가운데 정렬
	p.w.Write([]byte{0x1B, 0x61, 1})

	// Line spacing
	p.w.Write([]byte{0x1B, 0x33, 0})

	for i := 0; i < len(dataBuf); i += widthDataLen {
		var end int
		if i+widthDataLen >= len(dataBuf) {
			end = len(dataBuf)
		} else {
			end = i + widthDataLen
		}

		printBuf := append(cmdBuf, dataBuf[i:end]...)
		if _, err := p.w.Write(printBuf); err != nil {
			return errors.Wrap(err, "fail to print buf")
		}
		if err := p.w.Flush(); err != nil {
			return errors.Wrap(err, "fail to print buf")
		}
	}
	return nil
}

func (p *Printer) CutPaper() error {
	if _, err := p.w.Write([]byte("\x1B@\x1DVA0")); err != nil {
		return errors.Wrap(err, "fail to cut paper")
	}
	if err := p.w.Flush(); err != nil {
		return errors.Wrap(err, "fail to cut paper")
	}
	return nil
}

func (p *Printer) WriteString(txt string) error {
	if _, err := p.w.Write([]byte(txt)); err != nil {
		return errors.Wrap(err, "fail to cut paper")
	}
	if err := p.w.Flush(); err != nil {
		return errors.Wrap(err, "fail to cut paper")
	}
	return nil
}
