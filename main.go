package main

//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

import "github.com/livanjimenez/env-monitor/cmd/routers"

func main() {
	routers.RootRouters()
}
