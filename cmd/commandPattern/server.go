package main

import "fmt"

type Server struct {
}

func (s Server) start() {
	fmt.Println("Start server")
}

func (s Server) shutdown() {
	fmt.Println("Shutdown server")
}
