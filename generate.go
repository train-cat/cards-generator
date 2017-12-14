package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"time"

	"github.com/fogleman/gg"
	"github.com/sirupsen/logrus"
)

var mutex = &sync.Mutex{}

func generateStoptimeCards(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	v := r.URL.Query()

	mission := v.Get("mission")
	origin := formatStation(v.Get("origin"))
	terminus := formatStation(v.Get("terminus"))
	schedule := v.Get("schedule")
	days, err := strconv.Atoi(v.Get("days"))

	logger := logrus.WithFields(logrus.Fields{
		"action": "generate",
		"type":   "stop_time",
		"query":  r.URL.RawQuery,
	})

	if mission == "" || origin == "" || terminus == "" || schedule == "" || days == 0 || days > 127 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"status\": \"bad request\"}"))

		logger.WithField("status", "bad_request").Warn("bad_request")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	dc := gg.NewContextForImage(getBaseImage(Days(days)))

	dc.SetColor(colorText)

	dc.SetFontFace(fontRegular)
	dc.DrawStringAnchored(mission, xMission, yMission, 1, 1)
	dc.DrawStringAnchored(schedule, xSchedule, ySchedule, 0, 1)

	dc.SetFontFace(fontBold)
	dc.DrawStringAnchored(origin, xOrigin, yOrigin, 0, 1)
	dc.DrawStringAnchored(terminus, xTerminus, yTerminus, 0, 1)

	if err := dc.EncodePNG(w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		logrus.WithField("status", "fail").Error(err.Error())
		return
	}

	logger.WithFields(logrus.Fields{
		"status":        "success",
		"generate_time": time.Since(t).Seconds(),
	}).Info("success")
}

func formatStation(name string) string {
	if len(name) > limit {
		return fmt.Sprintf("%s...", name[:cutAfter])
	}

	return name
}
