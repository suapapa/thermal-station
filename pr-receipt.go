package main

import (
	"image"
	"log"

	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"github.com/suapapa/thermal-station/input"
	"github.com/suapapa/thermal-station/receipt"
)

type ReceiptPrinter struct {
	pr *receipt.Printer
}

func NewReceiptPrinter() *ReceiptPrinter {
	var pr *receipt.Printer
	if flagUsbDev != "" { // /dev/usb/lp0
		pr = receipt.NewUSBPrinter(flagUsbDev)
	} else {
		pr = receipt.NewSerialPrinter(flagSerialDev, flagSerialSpeed)
	}
	return &ReceiptPrinter{pr: pr}
}

func (p *ReceiptPrinter) PrintOrd(ord *input.Ord) error {
	log.Printf("receipt-ord: %v", ord)
	img, err := drawItems(receipt.MaxWidth, ord.ID, ord.Items)
	if err != nil {
		return errors.Wrap(err, "fail to print recipt ord")
	}
	switch flagDPI {
	case 200:
		if err := p.pr.PrintImage24bitDouble(img); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	default:
		if err := p.pr.PrintImage8bitDouble(img); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	}
	p.pr.CutPaper()
	return nil
}

func (p *ReceiptPrinter) PrintAddr(addr *input.Addr) error {
	log.Printf("receipt-addr: %v", addr)
	// TODO: TBD
	return nil
}

func (p *ReceiptPrinter) PrintQR(content string) error {
	log.Printf("receipt-qr: %v", content)
	qrc, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}

	if err := p.pr.PrintImage8bitDouble(qrc.Image(receipt.MaxWidth)); err != nil {
		return errors.Wrap(err, "fail to print recipt qr")
	}
	p.pr.CutPaper()
	return nil
}

func (p *ReceiptPrinter) PrintImg(img image.Image) error {
	log.Printf("receipt-img: %v, dpi=%d", img, flagDPI)
	switch flagDPI {
	case 200:
		if err := p.pr.PrintImage24bitDouble(img); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	default:
		if err := p.pr.PrintImage8bitDouble(img); err != nil {
			return errors.Wrap(err, "fail to print recipt img")
		}
	}
	p.pr.CutPaper()
	return nil
}
