package main

import (
	"flag"
	"log"

	"github.com/bndr/gojenkins"

	"github.com/alyoshka/traffic_light/app/traffic_light"
)

var jenkins = flag.String("jenkins", "localhost:8080", "url of jenkins")
var view = flag.String("view", "my view", "name of view")
var tty = flag.String("tty", "/dev/tty/ACM0", "tty where to write commands for arduino")

func main() {
	flag.Parse()

	trafficLight := trafficlight.NewTrafficLightTTY(*tty)

	jenkins, err := gojenkins.CreateJenkins(*jenkins).Init()
	if err != nil {
		trafficLight.Yellow()
		log.Panic("Failed to init jenkins: ", err)
	}

	view, err := jenkins.GetView(*view)
	if err != nil {
		trafficLight.Yellow()
		log.Panic("Failed to get views: ", err)
	}

	jobs := view.GetJobs()
	isRunning := false
	isOk := true
	for i := range jobs {
		switch jobs[i].Color {
		case "blue_anime":
			isRunning = true
		case "red":
			isOk = false
			break
		}
	}
	switch {
	case !isOk:
		// something broken
		trafficLight.Red()
	case isRunning:
		// build in progress
		trafficLight.Blink()
	default:
		// everything is OK
		trafficLight.Green()
	}
}
