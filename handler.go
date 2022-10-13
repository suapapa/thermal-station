package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suapapa/thermal-station/input"
)

func ordHandler(c *gin.Context) {
	var ord input.Ord
	if err := c.ShouldBindJSON(&ord); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			},
		)
		return
	}
	getPrinter(c.Param("printer")).PrintOrd(&ord)
}

func addrHandler(c *gin.Context) {
	var addr input.Addr
	if err := c.ShouldBindJSON(&addr); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			},
		)
		return
	}
	getPrinter(c.Param("printer")).PrintAddr(&addr)
}

func imgHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("img")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// override dpi
	dpiStr := c.Query("dpi")
	if dpiStr != "" {
		flagDPI, err = strconv.Atoi(dpiStr)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status": http.StatusBadRequest,
					"error":  err.Error(),
				},
			)
			return
		}
	}

	img, _, err := image.Decode(file)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	getPrinter(c.Param("printer")).PrintImg(img)
}

func qrHandler(c *gin.Context) {
	var qr input.QR
	if err := c.ShouldBindJSON(&qr); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			},
		)
		return
	}
	getPrinter(c.Param("printer")).PrintQR(qr.Content)
}
