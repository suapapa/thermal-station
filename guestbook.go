package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"github.com/suapapa/site-gb/msg"
	"github.com/suapapa/thermal-station/draw"
	"github.com/suapapa/thermal-station/receipt"
	"golang.org/x/image/font"
	"golang.org/x/sync/errgroup"
)

const (
	topicGB = "homin-dev/gb"
)

var (
	lastPork    = time.Now()
	maxWaitPork = 35 * time.Minute
	confSub     Config
)

func guestbook(ctx context.Context, conf Config) {
	confSub = conf
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(subF)
	eg.Go(checkPork)
	err := eg.Wait()
	log.Printf("eg stop: %v", err)
}

// ---

func subF() error {
	mqttC, err := connectBrokerByWSS(&confSub)
	if err != nil {
		return errors.Wrap(err, "fail to sub")
	}
	defer mqttC.Disconnect(1000)
	log.Println("sub: connected with MQTT broker")
	mqttC.Subscribe(topicGB, 1,
		func(mqttClient mqtt.Client, mqttMsg mqtt.Message) {
			topic, payload := mqttMsg.Topic(), mqttMsg.Payload()
			log.Printf("got %v from %s", string(payload), topic)

			switch topic {
			case "homin-dev/gb":
				lastPork = time.Now()
				if gb, err := getGBFromMsg(mqttMsg.Payload()); err != nil {
					log.Printf("err: %v", errors.Wrap(err, "fail in sub"))
				} else {
					if err := printToReceipt(gb); err != nil {
						log.Printf("err: %v", errors.Wrap(err, "fail in sub"))
					}
				}

			default:
				log.Printf("err: unknown topic %s", topic)
			}
		},
	)
	tk := time.NewTicker(10 * time.Second)
	defer tk.Stop()
	for range tk.C {
		if !mqttC.IsConnected() || !mqttC.IsConnectionOpen() {
			return errors.Wrap(err, "mqtt sub conn lost")
		}
	}
	return nil
}

func checkPork() error {
	tk := time.NewTicker(5 * time.Second)
	defer tk.Stop()
	for range tk.C {
		if time.Since(lastPork) > maxWaitPork {
			return fmt.Errorf("no porking over %v", maxWaitPork)
		}
	}
	return nil
}

func getGBFromMsg(msgBytes []byte) (*msg.GuestBook, error) {
	m := msg.Message{
		Data: &msg.GuestBook{}, // it is needed. if not data will be map[string]any
	}
	if err := json.Unmarshal(msgBytes, &m); err != nil {
		return nil, errors.Wrap(err, "fail to get gb from msg")
	}

	return m.GetGuestBook()
}

// 각 줄을 이미지로 만들어 출력
func printToReceipt(c *msg.GuestBook) error {
	mFF, err := getFont(48)
	if err != nil {
		return errors.Wrap(err, "fail to print")
	}
	lines := draw.FitToLines(mFF, receipt.MaxWidth, c.Content)
	if len(lines) == 0 {
		return fmt.Errorf("no content")
	}

	lines = append([]string{c.TimeStamp, c.From}, lines...)
	pr := NewReceiptPrinter()

	lines = append(lines, " ") // TODO: 마지막 줄이 잘려서 패딩 라인 붙임
	for _, l := range lines {
		img, err := draw.Txt2Img(mFF, receipt.MaxWidth, l)
		if err != nil {
			return errors.Wrap(err, "fail to print")
		}
		err = pr.PrintImg(img)
		if err != nil {
			return errors.Wrap(err, "fail to print")
		}
	}

	return nil
}

func getFont(size float64) (font.Face, error) {
	if flagFontPath != "" {
		data, err := os.ReadFile(flagFontPath)
		if err != nil {
			return nil, errors.Wrap(err, "fail to load font")
		}
		f, err := truetype.Parse(data)
		if err != nil {
			return nil, errors.Wrap(err, "fail to load font")
		}

		nface := truetype.NewFace(f, &truetype.Options{
			Size:    size,
			Hinting: font.HintingFull,
		})
		return nface, nil
	}

	return draw.GetFont(size)
}
