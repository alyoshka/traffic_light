package main

import (
	"flag"
	"fmt"
	
	"github.com/bndr/gojenkins"
	
	"github.com/alyoshka/traffic_light"
)

var jenkins = flag.String("jenkins", "localhost:8080", "url of jenkins")
var view = flag.String("view", "my view", "name of view")
var tty = flag.String("tty", "/dev/tty/ACM0", "tty where to write commands for arduino")


func main() {
	flag.Parse()
	
	trafficLight := tl.NewTrafficLightImpl(*tty)

	jenkins, err := gojenkins.CreateJenkins(*jenkins).Init()
	if err != nil {
		fmt.Println("Failed to create jenkins: ", err)
	} else {
		view, err := jenkins.GetView(*view)
		if err != nil {
			fmt.Println("Failed to get views: ", err)
		} else {
			jobs := view.GetJobs()
			isOk := true
			for i := range jobs {
				if jobs[i].Color == "red" {
					fmt.Println(jobs[i].Name, " is broken")
					isOk = false
					break
				}
			}
			if isOk {
				fmt.Println("Everything is OK")
				// Turn green light on
				trafficLight.Green()
			} else {
				// Turn red light on
				trafficLight.Red()
			}
		}
	}
	if err != nil {
		fmt.Println("Somethin went wrong")
		// Turn yellow light on
		trafficLight.Yellow()
	}
}
