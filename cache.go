package main

import (
	"image"

	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
)

var cacheImages = map[int]image.Image{}

func getBaseImage(week Days) image.Image {
	img, ok := cacheImages[int(week)]

	if !ok || refreshCache {
		img = generateBaseImage(week)

		log.Infof("miss cache")

		cacheImages[int(week)] = img
	}

	return img
}

func generateBaseImage(week Days) image.Image {
	dc := gg.NewContext(width, height)

	dc.SetColor(backgroundColor)
	dc.Clear()

	drawDays(dc, week)

	return dc.Image()
}

func drawDays(dc *gg.Context, week Days) {
	for _, d := range days {
		if week.HasFlag(d.Flag) {
			loadDayOn(dc)
		} else {
			loadDayOff(dc)
		}

		dc.DrawStringAnchored(d.Letter, float64(d.PosX), float64(centerCaseDays), 0.5, 0.5)
	}
}

func loadDayOn(dc *gg.Context) {
	dc.SetFontFace(fontBold)
	dc.SetColor(colorDayOn)
}

func loadDayOff(dc *gg.Context) {
	dc.SetFontFace(fontRegular)
	dc.SetColor(colorDayOff)
}
