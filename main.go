// main.go
package main

import (
	"ScheduleVM/controller"
	"ScheduleVM/model"
)

func main() {

	session := model.NewSession("signatures")
	server := controller.NewServer(session)

	server.Run()
}
