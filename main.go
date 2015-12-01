package main

import (
	"fmt"

	"github.com/bndr/gojenkins"
)

func main() {
	jenkinsAddr := "http://my-jenkins"
	viewName := "Local projects"

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
					fmt.Println(jobs[i].Name, " поломано")
					isOk = false
					break
				}
			}
			if isOk {
				fmt.Println("Всё хорошо")
				// Зажечь зелёный
			} else {
				// fmt.Println("Что-то сломалось")
				// Зажечь красный
			}
		}
	}
	if err != nil {
		fmt.Println("Что-то пошло не так")
		// Зажечь жёлтый
	}
}
