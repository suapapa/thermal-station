package main

import (
	"image"

	"github.com/suapapa/thermal-station/input"
)

type Printer interface {
	PrintOrd(*input.Ord) error
	PrintAddr(*input.Addr) error
	PrintQR(string) error
	PrintImg(image.Image) error
}

var (
	receiptPrinter, labelPrinter Printer
)

func getPrinter(printerType string) Printer {
	switch printerType {
	case "receipt":
		if receiptPrinter != nil {
			return receiptPrinter
		}
		receiptPrinter = NewReceiptPrinter()
		return receiptPrinter
	case "label":
		if labelPrinter != nil {
			return labelPrinter
		}
		labelPrinter = NewLabelPrinter(flagLabelPrinterDev)
		return labelPrinter
	}
	log.Errorf("unknown printer, %s", printerType)
	return NewLogoutPrinter()
}

// ---

type LogoutPrinter struct{}

func NewLogoutPrinter() *LogoutPrinter {
	return &LogoutPrinter{}
}

func (lp *LogoutPrinter) PrintOrd(ord *input.Ord) error {
	log.Infof("ord: %v", ord)
	return nil
}

func (lp *LogoutPrinter) PrintAddr(addr *input.Addr) error {
	log.Infof("addr: %v", addr)
	return nil
}

func (lp *LogoutPrinter) PrintQR(content string) error {
	log.Infof("qr: %v", content)
	return nil
}

func (lp *LogoutPrinter) PrintImg(img image.Image) error {
	log.Infof("img: %v, dpi=%d", img, flagDPI)
	return nil
}
