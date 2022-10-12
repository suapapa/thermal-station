package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"github.com/suapapa/thermal-station/draw"
	"github.com/suapapa/thermal-station/input"
	"golang.org/x/image/font"
)

const (
	// font size
	fsFromAddr = 46
	fsFromName = 48
	fsToAddr   = 98
	fsToName   = 85
	fsOrd      = 60
)

func drawItems(maxWidth int, ordID int, items []*input.Item) (image.Image, error) {
	var ordLines []string
	lineSpacing := 5
	for _, item := range items {
		ordLines = append(ordLines, fmt.Sprintf("- %s x %dea", item.Name, item.Cnt))
	}

	mw := maxWidth
	var ordF font.Face
	var err error
	if ordF, err = draw.GetFont(fsOrd); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	mh := int(fsOrd+lineSpacing)*len(items) + fsOrd
	dc := gg.NewContext(mw, mh)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(ordF)
	var y float64
	for _, line := range ordLines {
		y += (fsOrd + float64(lineSpacing))
		dc.DrawStringAnchored(line, 5, y, 0, 0)
	}

	if ordF, err = draw.GetFont(20); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}
	dc.SetFontFace(ordF)
	dc.DrawStringAnchored(fmt.Sprintf("ord#%d", ordID), float64(mw), float64(mh), 1, -1)

	return dc.Image(), nil
}

func drawAddressFrom(maxWidth int, addr *input.Addr) (image.Image, error) {
	mw := maxWidth
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsFromAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	var addrLines []string
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line1)...)
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line2)...)

	img, err := drawAddress(addrLines, addr.Name, addr.PostNumber, fsFromAddr, fsFromName, mw, -1)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	return img, nil
}

func drawAddressTo(maxWidth int, addr *input.Addr) (image.Image, error) {
	mw := (maxWidth * 3) / 2
	var addrF font.Face
	var err error
	if addrF, err = draw.GetFont(fsToAddr); err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	var addrLines []string
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line1)...)
	addrLines = append(addrLines, draw.FitToLines(addrF, mw, addr.Line2)...)

	img, err := drawAddress(addrLines, addr.Name, addr.PostNumber, fsToAddr, fsToName, mw, maxWidth)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}

	// TODO: rotate 90
	return img, nil
}

func drawAddress(addrLines []string, name, pn string, addrFSize, nameFSize float64, width int, height int) (image.Image, error) {
	addrF, err := draw.GetFont(addrFSize)
	if err != nil {
		return nil, err
	}
	nameF, err := draw.GetFont(nameFSize)
	if err != nil {
		return nil, err
	}
	pnFSize := addrFSize * 0.8
	pnF, err := draw.GetFont(pnFSize)
	if err != nil {
		return nil, err
	}

	var y float64
	var varHeight bool
	if height < 0 {
		varHeight = true
	}

	if varHeight {
		height = int(addrFSize+5)*len(addrLines) + int(nameFSize+nameFSize*0.2+10) + int(pnFSize)
	}
	dc := gg.NewContext(width, height)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(color.Black)
	dc.SetFontFace(addrF)
	for _, line := range addrLines {
		y += (addrFSize + 5)
		dc.DrawStringAnchored(line, 5, y, 0, 0)
	}
	pn = fmt.Sprintf("우) %s", pn)
	if varHeight {
		y += (nameFSize + 5)
		dc.SetFontFace(nameF)
		dc.DrawStringAnchored(name, float64(width)-5, y, 1, 0)
		dc.SetFontFace(pnF)
		dc.DrawStringAnchored(pn, 5, y, 0, 0)
	} else {
		y = float64(height)
		dc.SetFontFace(pnF)
		// dc.DrawStringAnchored(pn, float64(width)-5, y, 1, -1)
		dc.DrawStringAnchored(pn, 5, y, 0, -0.5)
		y -= (pnFSize + 5)
		dc.SetFontFace(nameF)
		dc.DrawStringAnchored(name, float64(width)-5, y, 1, -1)
	}

	ordF, err := draw.GetFont(20)
	if err != nil {
		return nil, errors.Wrap(err, "fail to draw from")
	}
	dc.SetFontFace(ordF)
	// TODO: will be deprecated
	// dc.DrawStringAnchored(fmt.Sprintf("ord#%d", ordID), float64(width), float64(height), 1, -1)

	return dc.Image(), nil
}
