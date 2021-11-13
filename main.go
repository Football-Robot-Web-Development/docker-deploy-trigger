package main

import "operation-backend/service"

func main() {
	r := service.SetupRouter()
	defer r.Run()
}
