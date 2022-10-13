package main

import (
	"image"

	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"github.com/suapapa/thermal-station/input"
	"github.com/suapapa/thermal-station/ql800_62"
)

type LabelPrinter struct {
	dev string
}

func NewLabelPrinter(dev string) *LabelPrinter {
	return &LabelPrinter{
		dev: dev,
	}
}

func (lp *LabelPrinter) PrintOrd(ord *input.Ord) error {
	log.Debugf("label-ord: %v", ord)
	img, err := drawItems(ql800_62.MaxWidth, ord.ID, ord.Items)
	if err != nil {
		return errors.Wrap(err, "fail to print label ord")
	}
	// err = ql800_62.PrintLabel(img)
	// if err != nil {
	// 	return errors.Wrap(err, "fail to print label ord")
	// }
	return lp.PrintImg(img)
}

func (lp *LabelPrinter) PrintAddr(addr *input.Addr) error {
	log.Debugf("label-addr: %v", addr)
	var img image.Image
	var err error
	switch addr.Vertical {
	case false: // From
		if img, err = drawAddressFrom(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
		return lp.printImg(img, 0)
	default: // case true: // To
		if img, err = drawAddressTo(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
		return lp.printImg(img, 90)
	}
}

func (lp *LabelPrinter) PrintQR(content string) error {
	log.Debugf("label-qr: %v", content)
	qrc, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}
	img := qrc.Image(ql800_62.MaxWidth)
	return lp.PrintImg(img)
}

func (lp *LabelPrinter) PrintImg(img image.Image) error {
	// log.Debugf("label-img: %v", img)
	return lp.printImg(img, 0)
}

func (lp *LabelPrinter) printImg(img image.Image, rotate int) error {
	err := ql800_62.PrintLabel(lp.dev, img, rotate)
	if err != nil {
		return errors.Wrap(err, "fail to print img")
	}
	return nil
}
