package main

import (
	"bytes"
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"github.com/suapapa/thermal-station/input"
)

type ReceiptPrinter struct{}

func NewReceiptPrinter() *ReceiptPrinter {
	return &ReceiptPrinter{}
}

func (lp *ReceiptPrinter) PrintOrd(ord *input.Ord) error {
	log.Printf("receipt-ord: %v", ord)
	return nil
}

func (lp *ReceiptPrinter) PrintAddr(addr *input.Addr) error {
	log.Printf("receipt-addr: %v", addr)
	return nil
}

func (lp *ReceiptPrinter) PrintQR(content string) error {
	log.Printf("receipt-qr: %v", content)
	png, err := qrcode.Encode(content, qrcode.Medium, receiptMaxWidth)
	if err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}

	pngReader := bytes.NewReader(png)
	if err := receiptPrImage8bitDouble(pngReader); err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}
	cutPaper()
	return nil
}

func (lp *ReceiptPrinter) PrintImg(r io.Reader, dpi int) error {
	log.Printf("receipt-img: %v", r)
	switch dpi {
	case 200:
		if err := receiptPrImage24bitDouble(r); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	default:
		if err := receiptPrImage8bitDouble(r); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	}
	cutPaper()
	return nil
}
