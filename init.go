package main

import (
	"flag"
	"image/color"

	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/image/font"
)

var (
	refreshCache bool
	limit        int
	cutAfter     int

	fontRegular font.Face
	fontBold    font.Face

	height         int
	width          int
	heightDays     int
	widthPerCase   int
	heightPerCase  int
	centerCaseDays int

	backgroundColor color.Color
	colorDayOn      color.Color
	colorDayOff     color.Color
	colorText       color.Color

	xMission  float64
	yMission  float64
	xSchedule float64
	ySchedule float64
	xOrigin   float64
	yOrigin   float64
	xTerminus float64
	yTerminus float64
)

func init() {
	initConfig()
	initFont()
	initVar()
	initDays()
}

func initConfig() {
	cfgFile := flag.String("config", "config.json", "config file")
	flag.Parse()

	viper.SetConfigFile(*cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func initFont() {
	var err error

	fontRegular, err = gg.LoadFontFace(viper.GetString("font.regular.path"), viper.GetFloat64("font.regular.size"))
	if err != nil {
		log.Fatal(err)
	}

	fontBold, err = gg.LoadFontFace(viper.GetString("font.bold.path"), viper.GetFloat64("font.bold.size"))
	if err != nil {
		log.Fatal(err)
	}
}

func initVar() {
	// size
	height = viper.GetInt("image.size.height")
	width = viper.GetInt("image.size.width")
	heightDays = viper.GetInt("image.position.days.y")
	widthPerCase = width / daysPerWeek
	heightPerCase = height - heightDays
	centerCaseDays = height - heightPerCase/2

	// color
	backgroundColor = color.RGBA{
		R: uint8(viper.GetInt("image.color.background.R")),
		G: uint8(viper.GetInt("image.color.background.G")),
		B: uint8(viper.GetInt("image.color.background.B")),
		A: uint8(viper.GetInt("image.color.background.A")),
	}
	colorDayOn = color.RGBA{
		R: uint8(viper.GetInt("image.color.day_on.R")),
		G: uint8(viper.GetInt("image.color.day_on.G")),
		B: uint8(viper.GetInt("image.color.day_on.B")),
		A: uint8(viper.GetInt("image.color.day_on.A")),
	}
	colorDayOff = color.RGBA{
		R: uint8(viper.GetInt("image.color.day_off.R")),
		G: uint8(viper.GetInt("image.color.day_off.G")),
		B: uint8(viper.GetInt("image.color.day_off.B")),
		A: uint8(viper.GetInt("image.color.day_off.A")),
	}
	colorText = color.RGBA{
		R: uint8(viper.GetInt("image.color.text.R")),
		G: uint8(viper.GetInt("image.color.text.G")),
		B: uint8(viper.GetInt("image.color.text.B")),
		A: uint8(viper.GetInt("image.color.text.A")),
	}

	// config
	refreshCache = viper.GetBool("cache.force_refresh")
	limit = viper.GetInt("text.limit")
	cutAfter = viper.GetInt("text.cut_after")

	// position
	xMission = viper.GetFloat64("image.position.mission.x")
	yMission = viper.GetFloat64("image.position.mission.y")

	xSchedule = viper.GetFloat64("image.position.schedule.x")
	ySchedule = viper.GetFloat64("image.position.schedule.y")

	xOrigin = viper.GetFloat64("image.position.origin.x")
	yOrigin = viper.GetFloat64("image.position.origin.y")

	xTerminus = viper.GetFloat64("image.position.terminus.x")
	yTerminus = viper.GetFloat64("image.position.terminus.y")
}

func initDays() {
	days = []Day{
		{Monday, "L", widthPerCase / 2},
		{Tuesday, "M", widthPerCase*1 + widthPerCase/2},
		{Wednesday, "M", widthPerCase*2 + widthPerCase/2},
		{Thursday, "J", widthPerCase*3 + widthPerCase/2},
		{Friday, "V", widthPerCase*4 + widthPerCase/2},
		{Saturday, "S", widthPerCase*5 + widthPerCase/2},
		{Sunday, "D", widthPerCase*6 + widthPerCase/2},
	}
}
