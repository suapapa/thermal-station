package main

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/suapapa/thermal-station/input"
)

var (
	ord = &input.Ord{
		ID: 20220926,
		Items: []*input.Item{
			{Name: "panic-01", Cnt: 3},
			{Name: "defer-01", Cnt: 3},
			{Name: "ch-01", Cnt: 3},
		},
	}
	fromAddr = &input.Addr{
		Line1:      "경기도 성남시 분당구 판교역로 166",
		Line2:      "",
		Name:       "판교 아지트",
		PostNumber: "12345",
	}
	toAddr = &input.Addr{
		Line1:      "경기 성남시 분당구 판교역로 235 (에이치스퀘어 엔동)",
		Line2:      "7층",
		Name:       "카카오 엔터프라이즈",
		PostNumber: "12345",
		Vertical:   true,
	}
)

// func TestPrintOrder(t *testing.T) {
// 	ord := input.Ord{
// 		ID:   1234567890,
// 		From: ord.From,
// 		To:   ord.To,
// 	}
// 	je := json.NewEncoder(os.Stdout)
// 	je.SetIndent("", "  ")
// 	je.Encode(&ord)
// }

func TestDrawItems(t *testing.T) {
	img, err := drawItems(45, ord.Items)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw items"))
	}
	if err := saveImg2Png(img, "items.png"); err != nil {
		t.Error(err)
	}
}

func TestDrawAddressXXX(t *testing.T) {
	img, err := drawAddressFrom(45, fromAddr)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_from.png"); err != nil {
		t.Error(err)
	}
	img, err = drawAddressTo(ord.ID, toAddr)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_to.png"); err != nil {
		t.Error(err)
	}
}

func TestDrawAddress(t *testing.T) {
	from, to := fromAddr, toAddr
	img, err := drawAddress(
		ord.ID,
		[]string{from.Line1, from.Line2},
		from.Name,
		"",
		fsFromAddr, fsFromName,
		ql800MaxPix,
		-1,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "draw_from.png"); err != nil {
		t.Error(err)
	}

	img, err = drawAddress(
		ord.ID,
		[]string{to.Line1, to.Line2},
		to.Name,
		"",
		fsToAddr, fsToName,
		(ql800MaxPix*3)/2,
		ql800MaxPix,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "draw_to.png"); err != nil {
		t.Error(err)
	}
}
