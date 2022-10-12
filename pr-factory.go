package main

import (
	"io"
	"log"

	"github.com/suapapa/thermal-station/input"
)

type Printer interface {
	PrintOrd(*input.Ord) error
	PrintAddr(*input.Addr) error
	PrintQR(string) error
	PrintImg(io.Reader, int) error
}

func getPrinter(printerType string) Printer {
	switch printerType {
	case "receipt":
		return NewReceiptPrinter()
	case "label":
		return NewLabelPrinter()
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

func (lp *LogoutPrinter) PrintImg(r io.Reader, dpi int) error {
	log.Printf("img: %v, dpi=%d", r, dpi)
	return nil
}
