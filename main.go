package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	programName                   = "thermal-station"
	buildStamp, gitHash, buildTag string

	flagDPI         int
	flagSerialDev   string
	flagSerialSpeed int
	flagUsbDev      string
)

func main() {
	log.Printf("%s-%s-%s(%s)", programName, buildTag, gitHash, buildStamp)

	flag.IntVar(&flagDPI, "dpi", 200, "receipt printer DPI (100 or 200)")
	flag.StringVar(&flagSerialDev, "s", "/dev/ttyACM0", "serial device")
	flag.IntVar(&flagSerialSpeed, "ss", 9600, "serial speed")
	flag.StringVar(&flagUsbDev, "u", "", "if specify usb lp device -s will be ignored")
	flag.Parse()

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
