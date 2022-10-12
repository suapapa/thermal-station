package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tarm/serial"
)

var (
	programName                   = "thermal-station"
	buildStamp, gitHash, buildTag string

	printerDev    *bufio.Writer
	flagSerialDev string
	flagUsbDev    string
	flagAdvDu     int
)

func main() {
	log.Printf("%s-%s-%s(%s)", programName, buildTag, gitHash, buildStamp)

	flag.StringVar(&flagSerialDev, "s", "/dev/ttyACM0", "serial device")
	flag.StringVar(&flagUsbDev, "u", "", "if specify usb lp device -s will be ignored")
	flag.IntVar(&flagAdvDu, "a", 0, "ble advertisement duration")
	flag.Parse()

	if flagUsbDev != "" { // /dev/usb/lp0
		dev, err := os.OpenFile(flagUsbDev, os.O_RDWR, 0)
		printerDev = bufio.NewWriter(dev)
		if err != nil {
			log.Fatal(err)
		}
		defer dev.Close()
	} else {
		c := &serial.Config{Name: flagSerialDev, Baud: 9600}
		dev, err := serial.OpenPort(c)
		if err != nil {
			log.Fatal(err)
		}
		printerDev = bufio.NewWriter(dev)
		defer dev.Close()
	}

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	v1 := r.Group("v1")

	v1.POST("/:printer/ord", ordHandler)
	v1.POST("/:printer/addr", addrHandler)
	v1.POST("/:printer/img", imgHandler)
	v1.POST("/:printer/qr", qrHandler)

	go r.Run(":8080")

	stopC := make(chan interface{})
	<-stopC
}
