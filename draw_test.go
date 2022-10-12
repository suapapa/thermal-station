package main

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/suapapa/thermal-station/input"
	"github.com/suapapa/thermal-station/ql800_62"
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
	img, err := drawItems(ql800_62.MaxWidth, ord.ID, ord.Items)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw items"))
	}
	if err := saveImg2Png(img, "items.png"); err != nil {
		t.Error(err)
	}
}

func TestDrawAddressXXX(t *testing.T) {
	img, err := drawAddressFrom(ql800_62.MaxWidth, fromAddr)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_from.png"); err != nil {
		t.Error(err)
	}
	img, err = drawAddressTo(ql800_62.MaxWidth, toAddr)
	if err != nil {
		t.Error(errors.Wrap(err, "fail to draw address"))
	}
	if err := saveImg2Png(img, "addr_to.png"); err != nil {
		t.Error(err)
	}
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
