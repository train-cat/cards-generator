package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fogleman/gg"
)

func serveImage(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	mission := v.Get("mission")
	origin := formatStation(v.Get("origin"))
	terminus := formatStation(v.Get("terminus"))
	schedule := v.Get("schedule")
	days, err := strconv.Atoi(v.Get("days"))

	if mission == "" || origin == "" || terminus == "" || schedule == "" || days == 0 || days > 127 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"bad request\"}"))
		return
	}

	dc := gg.NewContextForImage(getBaseImage(Days(days)))

	dc.SetColor(colorText)

	dc.SetFontFace(fontRegular)
	dc.DrawStringAnchored(mission, xMission, yMission, 1, 1)
	dc.DrawStringAnchored(schedule, xSchedule, ySchedule, 0, 1)

	dc.SetFontFace(fontBold)
	dc.DrawStringAnchored(origin, xOrigin, yOrigin, 0, 1)
	dc.DrawStringAnchored(terminus, xTerminus, yTerminus, 0, 1)

	dc.EncodePNG(w)
}

func formatStation(name string) string {
	if len(name) > limit {
		return fmt.Sprintf("%s...", name[:cutAfter])
	}

	return name
}
