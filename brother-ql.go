package main

import (
	"image"
	"image/png"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/suapapa/thermal-station/input"
)

type Orient int

const (
	Landscape Orient = iota
	Portrait

	ql800MaxPix = 696 // 62mm endless
	tmpPngPath  = "/tmp/order_pic.png"
)

func labelItems(ordID int, items []*input.Item) error {
	img, err := drawItems(ordID, items)
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}
	if err = saveImg2Png(img, tmpPngPath); err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	err = exec.Command("sh", "-c", "brother_ql -b pyusb -p usb://04f9:209b -m QL-800 print -l 62 "+tmpPngPath).Run()
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	return nil
}

func labelAddr(ordID int, addr *input.Addr) error {
	img, err := drawAddressFrom(ordID, addr)
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}
	if err = saveImg2Png(img, tmpPngPath); err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	err = exec.Command("sh", "-c", "brother_ql -b pyusb -p usb://04f9:209b -m QL-800 print -l 62 "+tmpPngPath).Run()
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	return nil
}

func labelAddrVertical(ordID int, addr *input.Addr) error {
	img, err := drawAddressTo(ordID, addr)
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}
	if err = saveImg2Png(img, tmpPngPath); err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	err = exec.Command("sh", "-c", "brother_ql -b pyusb -p usb://04f9:209b -m QL-800 print -r 90 -l 62 "+tmpPngPath).Run()
	if err != nil {
		return errors.Wrap(err, "fail to print from")
	}

	return nil
}

func saveImg2Png(img image.Image, pngFN string) error {
	f, err := os.Create(pngFN)
	if err != nil {
		return errors.Wrap(err, "fail to savePNG")
	}
	defer f.Close()

	if err = png.Encode(f, img); err != nil {
		return errors.Wrap(err, "fail to savePNG")
	}

	return nil
}
