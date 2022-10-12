package main

import (
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/suapapa/thermal-station/input"
)

type LabelPrinter struct{}

func NewLabelPrinter() *LabelPrinter {
	return &LabelPrinter{}
}

func (lp *LabelPrinter) PrintOrd(ord *input.Ord) error {
	log.Printf("label-ord: %v", ord)
	if err := labelItems(ord.ID, ord.Items); err != nil {
		return errors.Wrap(err, "fail to print label ord")
	}
	return nil
}

func (lp *LabelPrinter) PrintAddr(addr *input.Addr) error {
	log.Printf("label-addr: %v", addr)
	switch addr.Vertical {
	case false: // From
		if err := labelAddr(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
	case true: // To
		if err := labelAddrVertical(-1, addr); err != nil {
			return errors.Wrap(err, "fail to print label addr")
		}
	}
	return nil
}

func (lp *LabelPrinter) PrintQR(content string) error {
	log.Printf("label-qr: %v", content)
	// png, err := qrcode.Encode(content, qrcode.Medium, receiptMaxWidth)
	// if err != nil {
	// 	return errors.Wrap(err, "fail to print recipt qr")
	// }

	// pngReader := bytes.NewReader(png)
	// if err := receiptPrImage8bitDouble(pngReader); err != nil {
	// 	return errors.Wrap(err, "fail to print recipt qr")
	// }
	// cutPaper()
	return nil
}

func (lp *LabelPrinter) PrintImg(r io.Reader, dpi int) error {
	log.Printf("label-img: %v", r)
	// switch dpi {
	// case 200:
	// 	if err := receiptPrImage24bitDouble(r); err != nil {
	// 		return errors.Wrap(err, "fail to print recipt img")
	// 	}
	// default:
	// 	if err := receiptPrImage8bitDouble(r); err != nil {
	// 		return errors.Wrap(err, "fail to print recipt img")
	// 	}
	// }
	// cutPaper()
	return nil
}
