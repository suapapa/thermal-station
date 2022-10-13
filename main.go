package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	programName = "thermal-station"
	programVer  = "dev"

	flagDPI         int
	flagSerialDev   string // Serial flag will be ignored if flagUsbDev is set
	flagSerialSpeed int
	flagUsbDev      string

	flagLabelPrinterDev string

	flagFontPath string

	flagDebug bool
)

func main() {
	log.Infof("%s-%s", programName, programVer)

	flag.IntVar(&flagDPI, "dpi", 200, "receipt printer DPI (100 or 200)")
	flag.StringVar(&flagSerialDev, "s", "/dev/ttyACM0", "serial device")
	flag.IntVar(&flagSerialSpeed, "ss", 9600, "serial speed")
	flag.StringVar(&flagUsbDev, "u", "", "if specify usb lp device -s will be ignored")
	flag.StringVar(&flagFontPath, "f", "", "external font path to use")
	flag.StringVar(&flagLabelPrinterDev, "l", "/dev/usb/lp0", "labelPrinter dev")
	flag.BoolVar(&flagDebug, "debug", false, "print debug logs")
	flag.Parse()

	if flagDebug {
		log.Logger.SetLevel(logrus.DebugLevel)
	}

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := r.Group("v1")

	v1.POST("/:printer/ord", ordHandler)
	v1.POST("/:printer/addr", addrHandler)
	v1.POST("/:printer/img", imgHandler)
	v1.POST("/:printer/qr", qrHandler)

	go r.Run(":8080")

	ctx, cancelF := context.WithCancel(context.Background())
	defer cancelF()
	conf := Config{
		Host:     "homin.dev",
		Port:     9001,
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
		ClientID: fmt.Sprintf("%s-%s", programName, programVer),
	}
	go guestbook(ctx, conf)

	stopC := make(chan interface{})
	<-stopC
}
