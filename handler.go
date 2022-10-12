package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suapapa/thermal-station/input"
)

func ordHandler(c *gin.Context) {
	// printer := getPrinter(c.Param("printer"))
}

func addrHandler(c *gin.Context) {
	// printer := getPrinter(c.Param("printer"))
}

func imgHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("img")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	dpi, err := strconv.Atoi(c.Query("dpi"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	printer := getPrinter(c.Param("printer"))
	printer.PrintImg(file, dpi)
}

func qrHandler(c *gin.Context) {
	var input input.QR
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	printer := getPrinter(c.Param("printer"))
	printer.PrintQR(input.Content)
}