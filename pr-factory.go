package main

import (
	"image"
	"log"

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
		labelPrinter = NewLabelPrinter()
		return labelPrinter
	}
	return NewLogoutPrinter()
}

// ---

type LogoutPrinter struct{}

func NewLogoutPrinter() *LogoutPrinter {
	return &LogoutPrinter{}
}

func (lp *LogoutPrinter) PrintOrd(ord *input.Ord) error {
	log.Printf("ord: %v", ord)
	return nil
}

func (lp *LogoutPrinter) PrintAddr(addr *input.Addr) error {
	log.Printf("addr: %v", addr)
	return nil
}

func (lp *LogoutPrinter) PrintQR(content string) error {
	log.Printf("qr: %v", content)
	return nil
}

func (lp *LogoutPrinter) PrintImg(img image.Image) error {
	log.Printf("img: %v, dpi=%d", img, flagDPI)
	return nil
}
