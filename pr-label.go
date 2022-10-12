package main

import (
	"image"
	"log"

	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"github.com/suapapa/thermal-station/input"
	"github.com/suapapa/thermal-station/ql800_62"
)

type LabelPrinter struct{}

func NewLabelPrinter() *LabelPrinter {
	return &LabelPrinter{}
}

func (lp *LabelPrinter) PrintOrd(ord *input.Ord) error {
	log.Printf("label-ord: %v", ord)
	img, err := drawItems(ql800_62.MaxWidth, ord.ID, ord.Items)
	if err != nil {
		return errors.Wrap(err, "fail to print label ord")
	}
	err = ql800_62.PrintLabel(img)
	if err != nil {
		return errors.Wrap(err, "fail to print label ord")
	}
	return lp.printImg(img)
}

func (lp *LabelPrinter) PrintAddr(addr *input.Addr) error {
	log.Printf("label-addr: %v", addr)
	var img image.Image
	var err error
	switch addr.Vertical {
	case false: // From
		if img, err = drawAddressFrom(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
	case true: // To
		if img, err = drawAddressTo(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
	}
	return lp.printImg(img)
}

func (lp *LabelPrinter) PrintQR(content string) error {
	log.Printf("label-qr: %v", content)
	qrc, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}
	img := qrc.Image(ql800_62.MaxWidth)
	return lp.printImg(img)
}

func (lp *LabelPrinter) PrintImg(img image.Image, dpi int) error {
	log.Printf("label-img: %v", img)
	return lp.printImg(img)
}

func (lp *LabelPrinter) printImg(img image.Image) error {
	err := ql800_62.PrintLabel(img)
	if err != nil {
		return errors.Wrap(err, "fail to print img")
	}
	return nil
}
