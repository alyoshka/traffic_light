package main

import (
	"fmt"

	"github.com/bndr/gojenkins"
)

func main() {
	jenkinsAddr := "http://my-jenkins"
	viewName := "my view"

	jenkins, err := gojenkins.CreateJenkins(jenkinsAddr).Init()
	if err != nil {
		fmt.Println("Failed to create jenkins: ", err)
	} else {
		view, err := jenkins.GetView(viewName)
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
			} else {
				// Turn red light on
			}
		}
	}
	if err != nil {
		fmt.Println("Somethin went wrong")
		// Turn yellow light on
	}
}
